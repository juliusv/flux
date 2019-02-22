package plan

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
)

// PhysicalPlanner performs transforms a logical plan to a physical plan,
// by applying any registered physical rules.
type PhysicalPlanner interface {
	Plan(lplan *PlanSpec) (*PlanSpec, error)
}

// NewPhysicalPlanner creates a new physical plan with the specified options.
// The new plan will be configured to apply any physical rules that have been registered.
func NewPhysicalPlanner(options ...PhysicalOption) PhysicalPlanner {
	pp := &physicalPlanner{
		heuristicPlanner:   newHeuristicPlanner(),
		defaultMemoryLimit: math.MaxInt64,
	}

	rules := make([]Rule, len(ruleNameToPhysicalRule))
	i := 0
	for _, v := range ruleNameToPhysicalRule {
		rules[i] = v
		i++
	}

	pp.addRules(rules...)

	pp.addRules(physicalConverterRule{})

	// Options may add or remove rules, so process them after we've
	// added registered rules.
	for _, opt := range options {
		opt.apply(pp)
	}

	return pp
}

func (pp *physicalPlanner) Plan(spec *PlanSpec) (*PlanSpec, error) {
	transformedSpec, err := pp.heuristicPlanner.Plan(spec)
	if err != nil {
		return nil, err
	}

	// Compute time bounds for nodes in the plan
	if err := transformedSpec.BottomUpWalk(ComputeBounds); err != nil {
		return nil, err
	}

	// Ensure that the plan is valid
	if !pp.disableValidation {
		err := transformedSpec.CheckIntegrity()
		if err != nil {
			return nil, err
		}

		err = validatePhysicalPlan(transformedSpec)
		if err != nil {
			return nil, err
		}
	}

	// Update memory quota
	if transformedSpec.Resources.MemoryBytesQuota == 0 {
		transformedSpec.Resources.MemoryBytesQuota = pp.defaultMemoryLimit
	}

	// Update concurrency quota
	if transformedSpec.Resources.ConcurrencyQuota == 0 {
		transformedSpec.Resources.ConcurrencyQuota = len(transformedSpec.Roots)
	}

	return transformedSpec, nil
}

func validatePhysicalPlan(plan *PlanSpec) error {
	err := plan.BottomUpWalk(func(pn PlanNode) error {
		if validator, ok := pn.ProcedureSpec().(PostPhysicalValidator); ok {
			return validator.PostPhysicalValidate(pn.ID())
		}

		if _, ok := pn.(*PhysicalPlanNode); !ok {
			return errors.Errorf("logical node \"%v\" could not be converted to a physical node", pn.ID())
		}

		return nil
	})
	return err
}

type physicalPlanner struct {
	*heuristicPlanner
	defaultMemoryLimit int64
	disableValidation  bool
}

// PhysicalOption is an option to configure the behavior of the physical plan.
type PhysicalOption interface {
	apply(*physicalPlanner)
}

type physicalOption func(*physicalPlanner)

func (opt physicalOption) apply(p *physicalPlanner) {
	opt(p)
}

// WithDefaultMemoryLimit sets the default memory limit for plans generated by the plan.
// If the query spec explicitly sets a memory limit, that limit is used instead of the default.
func WithDefaultMemoryLimit(memBytes int64) PhysicalOption {
	return physicalOption(func(p *physicalPlanner) {
		p.defaultMemoryLimit = memBytes
	})
}

// OnlyPhysicalRules produces a physical plan option that forces only a particular set of rules to be applied.
func OnlyPhysicalRules(rules ...Rule) PhysicalOption {
	return physicalOption(func(pp *physicalPlanner) {
		pp.clearRules()
		pp.addRules(rules...)
	})
}

// Disables validation in the physical planner
func DisableValidation() PhysicalOption {
	return physicalOption(func(p *physicalPlanner) {
		p.disableValidation = true
	})
}

// physicalConverterRule rewrites logical nodes that have a ProcedureSpec that implements
// PhysicalProcedureSpec as a physical node.  For operations that have a 1:1 relationship
// between their physical and logical operations, this is the default behavior.
type physicalConverterRule struct {
}

func (physicalConverterRule) Name() string {
	return "physicalConverterRule"
}

func (physicalConverterRule) Pattern() Pattern {
	return Any()
}

func (physicalConverterRule) Rewrite(pn PlanNode) (PlanNode, bool, error) {
	if _, ok := pn.(*PhysicalPlanNode); ok {
		// Already converted
		return pn, false, nil
	}

	ln := pn.(*LogicalPlanNode)
	pspec, ok := ln.Spec.(PhysicalProcedureSpec)
	if !ok {
		// A different rule will do the conversion
		return pn, false, nil
	}

	newNode := PhysicalPlanNode{
		bounds: ln.bounds,
		id:     "phys_" + ln.id,
		Spec:   pspec,
	}

	ReplaceNode(pn, &newNode)

	return &newNode, true, nil
}

// PhysicalProcedureSpec is similar to its logical counterpart but must provide a method to determine cost.
type PhysicalProcedureSpec interface {
	Kind() ProcedureKind
	Copy() ProcedureSpec
	Cost(inStats []Statistics) (cost Cost, outStats Statistics)
}

// PhysicalPlanNode represents a physical operation in a plan.
type PhysicalPlanNode struct {
	edges
	bounds
	id   NodeID
	Spec PhysicalProcedureSpec

	// The attributes required from inputs to this node
	RequiredAttrs []PhysicalAttributes

	// The attributes provided to consumers of this node's output
	OutputAttrs PhysicalAttributes
}

// ID returns a human-readable id for this plan node.
func (ppn *PhysicalPlanNode) ID() NodeID {
	return ppn.id
}

// ProcedureSpec returns the procedure spec for this plan node.
func (ppn *PhysicalPlanNode) ProcedureSpec() ProcedureSpec {
	return ppn.Spec
}

func (ppn *PhysicalPlanNode) ReplaceSpec(newSpec ProcedureSpec) error {
	physSpec, ok := newSpec.(PhysicalProcedureSpec)
	if !ok {
		return fmt.Errorf("couldn't replace ProcedureSpec for physical plan node \"%v\"", ppn.ID())
	}

	ppn.Spec = physSpec
	return nil
}

// Kind returns the procedure kind for this plan node.
func (ppn *PhysicalPlanNode) Kind() ProcedureKind {
	return ppn.Spec.Kind()
}

func (ppn *PhysicalPlanNode) ShallowCopy() PlanNode {
	newNode := new(PhysicalPlanNode)
	newNode.edges = ppn.edges.shallowCopy()
	newNode.id = ppn.id + "_copy"
	// TODO: the type assertion below... is it needed?
	newNode.Spec = ppn.Spec.Copy().(PhysicalProcedureSpec)
	return newNode
}

// Cost provides the self-cost (i.e., does not include the cost of its predecessors) for
// this plan node.  Caller must provide statistics of predecessors to this node.
func (ppn *PhysicalPlanNode) Cost(inStats []Statistics) (cost Cost, outStats Statistics) {
	return ppn.Spec.Cost(inStats)
}

// PhysicalAttributes encapsulates sny physical attributes of the result produced
// by a physical plan node, such as collation, etc.
type PhysicalAttributes struct {
}

// CreatePhysicalNode creates a single physical plan node from a procedure spec.
// The newly created physical node has no incoming or outgoing edges.
func CreatePhysicalNode(id NodeID, spec PhysicalProcedureSpec) *PhysicalPlanNode {
	return &PhysicalPlanNode{
		id:   id,
		Spec: spec,
	}
}

// PostPhysicalValidator provides an interface that can be implemented by PhysicalProcedureSpecs for any
// validation checks to be performed post-physical planning.
type PostPhysicalValidator interface {
	PostPhysicalValidate(id NodeID) error
}
