// DO NOT EDIT: This file is autogenerated via the builtin command.

package promql

import (
	flux "github.com/influxdata/flux"
	ast "github.com/influxdata/flux/ast"
)

func init() {
	flux.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 2,
					Line:   108,
				},
				File:   "promql.flux",
				Source: "package promql\n\n// changes() implements functionality equivalent to PromQL's changes() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#changes\nbuiltin changes\n\n// promqlDayOfMonth() implements functionality equivalent to PromQL's day_of_month() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month\nbuiltin promqlDayOfMonth\n\n// promqlDayOfWeek() implements functionality equivalent to PromQL's day_of_week() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week\nbuiltin promqlDayOfWeek\n\n// promqlDaysInMonth() implements functionality equivalent to PromQL's days_in_month() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month\nbuiltin promqlDaysInMonth\n\n// emptyTable() returns an empty table, which is used as a helper function to implement\n// PromQL's time() and vector() functions:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#time\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#vector\nbuiltin emptyTable\n\n// extrapolatedRate() is a helper function that calculates extrapolated rates over\n// counters and is used to implement PromQL's rate(), delta(), and increase() functions.\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#rate\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#increase\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#delta\nbuiltin extrapolatedRate\n\n// holtWinters() implements functionality equivalent to PromQL's holt_winters()\n// function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#holt_winters\nbuiltin holtWinters\n\n// promqlHour() implements functionality equivalent to PromQL's hour() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#hour\nbuiltin promqlHour\n\n// instantRate() is a helper function that calculates instant rates over\n// counters and is used to implement PromQL's irate() and idelta() functions.\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#irate\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta\nbuiltin instantRate\n\n// labelReplace implements functionality equivalent to PromQL's label_replace() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace\nbuiltin labelReplace\n\n// linearRegression implements linear regression functionality required to implement\n// PromQL's deriv() and predict_linear() functions:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear\nbuiltin linearRegression\n\n// minute() implements functionality equivalent to PromQL's minute() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#minute\nbuiltin promqlMinute\n\n// month() implements functionality equivalent to PromQL's month() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#month\nbuiltin promqlMonth\n\n// promHistogramQuantile() implements functionality equivalent to PromQL's\n// histogram_quantile() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile\nbuiltin promHistogramQuantile\n\n// resets() implements functionality equivalent to PromQL's resets() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#resets\nbuiltin resets\n\n// timestamp() implements functionality equivalent to PromQL's timestamp() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp\nbuiltin timestamp\n\n// year() implements functionality equivalent to PromQL's year() function:\n//\n// https://prometheus.io/docs/prometheus/latest/querying/functions/#year\nbuiltin promqlYear\n\n// hack to simulate an imported promql package\npromql = {\n  promqlDayOfMonth:promqlDayOfMonth,\n  promqlDayOfWeek:promqlDayOfWeek,\n  promqlDaysInMonth:promqlDaysInMonth,\n  promqlHour:promqlHour,\n  promqlMinute:promqlMinute,\n  promqlMonth:promqlMonth,\n  promqlYear:promqlYear,\n}",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   6,
					},
					File:   "promql.flux",
					Source: "builtin changes",
					Start: ast.Position{
						Column: 1,
						Line:   6,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 16,
							Line:   6,
						},
						File:   "promql.flux",
						Source: "changes",
						Start: ast.Position{
							Column: 9,
							Line:   6,
						},
					},
				},
				Name: "changes",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 25,
						Line:   11,
					},
					File:   "promql.flux",
					Source: "builtin promqlDayOfMonth",
					Start: ast.Position{
						Column: 1,
						Line:   11,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 25,
							Line:   11,
						},
						File:   "promql.flux",
						Source: "promqlDayOfMonth",
						Start: ast.Position{
							Column: 9,
							Line:   11,
						},
					},
				},
				Name: "promqlDayOfMonth",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 24,
						Line:   16,
					},
					File:   "promql.flux",
					Source: "builtin promqlDayOfWeek",
					Start: ast.Position{
						Column: 1,
						Line:   16,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 24,
							Line:   16,
						},
						File:   "promql.flux",
						Source: "promqlDayOfWeek",
						Start: ast.Position{
							Column: 9,
							Line:   16,
						},
					},
				},
				Name: "promqlDayOfWeek",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 26,
						Line:   21,
					},
					File:   "promql.flux",
					Source: "builtin promqlDaysInMonth",
					Start: ast.Position{
						Column: 1,
						Line:   21,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 26,
							Line:   21,
						},
						File:   "promql.flux",
						Source: "promqlDaysInMonth",
						Start: ast.Position{
							Column: 9,
							Line:   21,
						},
					},
				},
				Name: "promqlDaysInMonth",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 19,
						Line:   28,
					},
					File:   "promql.flux",
					Source: "builtin emptyTable",
					Start: ast.Position{
						Column: 1,
						Line:   28,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 19,
							Line:   28,
						},
						File:   "promql.flux",
						Source: "emptyTable",
						Start: ast.Position{
							Column: 9,
							Line:   28,
						},
					},
				},
				Name: "emptyTable",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 25,
						Line:   36,
					},
					File:   "promql.flux",
					Source: "builtin extrapolatedRate",
					Start: ast.Position{
						Column: 1,
						Line:   36,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 25,
							Line:   36,
						},
						File:   "promql.flux",
						Source: "extrapolatedRate",
						Start: ast.Position{
							Column: 9,
							Line:   36,
						},
					},
				},
				Name: "extrapolatedRate",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 20,
						Line:   42,
					},
					File:   "promql.flux",
					Source: "builtin holtWinters",
					Start: ast.Position{
						Column: 1,
						Line:   42,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 20,
							Line:   42,
						},
						File:   "promql.flux",
						Source: "holtWinters",
						Start: ast.Position{
							Column: 9,
							Line:   42,
						},
					},
				},
				Name: "holtWinters",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 19,
						Line:   47,
					},
					File:   "promql.flux",
					Source: "builtin promqlHour",
					Start: ast.Position{
						Column: 1,
						Line:   47,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 19,
							Line:   47,
						},
						File:   "promql.flux",
						Source: "promqlHour",
						Start: ast.Position{
							Column: 9,
							Line:   47,
						},
					},
				},
				Name: "promqlHour",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 20,
						Line:   54,
					},
					File:   "promql.flux",
					Source: "builtin instantRate",
					Start: ast.Position{
						Column: 1,
						Line:   54,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 20,
							Line:   54,
						},
						File:   "promql.flux",
						Source: "instantRate",
						Start: ast.Position{
							Column: 9,
							Line:   54,
						},
					},
				},
				Name: "instantRate",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 21,
						Line:   59,
					},
					File:   "promql.flux",
					Source: "builtin labelReplace",
					Start: ast.Position{
						Column: 1,
						Line:   59,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 21,
							Line:   59,
						},
						File:   "promql.flux",
						Source: "labelReplace",
						Start: ast.Position{
							Column: 9,
							Line:   59,
						},
					},
				},
				Name: "labelReplace",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 25,
						Line:   66,
					},
					File:   "promql.flux",
					Source: "builtin linearRegression",
					Start: ast.Position{
						Column: 1,
						Line:   66,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 25,
							Line:   66,
						},
						File:   "promql.flux",
						Source: "linearRegression",
						Start: ast.Position{
							Column: 9,
							Line:   66,
						},
					},
				},
				Name: "linearRegression",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 21,
						Line:   71,
					},
					File:   "promql.flux",
					Source: "builtin promqlMinute",
					Start: ast.Position{
						Column: 1,
						Line:   71,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 21,
							Line:   71,
						},
						File:   "promql.flux",
						Source: "promqlMinute",
						Start: ast.Position{
							Column: 9,
							Line:   71,
						},
					},
				},
				Name: "promqlMinute",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 20,
						Line:   76,
					},
					File:   "promql.flux",
					Source: "builtin promqlMonth",
					Start: ast.Position{
						Column: 1,
						Line:   76,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 20,
							Line:   76,
						},
						File:   "promql.flux",
						Source: "promqlMonth",
						Start: ast.Position{
							Column: 9,
							Line:   76,
						},
					},
				},
				Name: "promqlMonth",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 30,
						Line:   82,
					},
					File:   "promql.flux",
					Source: "builtin promHistogramQuantile",
					Start: ast.Position{
						Column: 1,
						Line:   82,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 30,
							Line:   82,
						},
						File:   "promql.flux",
						Source: "promHistogramQuantile",
						Start: ast.Position{
							Column: 9,
							Line:   82,
						},
					},
				},
				Name: "promHistogramQuantile",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 15,
						Line:   87,
					},
					File:   "promql.flux",
					Source: "builtin resets",
					Start: ast.Position{
						Column: 1,
						Line:   87,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 15,
							Line:   87,
						},
						File:   "promql.flux",
						Source: "resets",
						Start: ast.Position{
							Column: 9,
							Line:   87,
						},
					},
				},
				Name: "resets",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 18,
						Line:   92,
					},
					File:   "promql.flux",
					Source: "builtin timestamp",
					Start: ast.Position{
						Column: 1,
						Line:   92,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 18,
							Line:   92,
						},
						File:   "promql.flux",
						Source: "timestamp",
						Start: ast.Position{
							Column: 9,
							Line:   92,
						},
					},
				},
				Name: "timestamp",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 19,
						Line:   97,
					},
					File:   "promql.flux",
					Source: "builtin promqlYear",
					Start: ast.Position{
						Column: 1,
						Line:   97,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 19,
							Line:   97,
						},
						File:   "promql.flux",
						Source: "promqlYear",
						Start: ast.Position{
							Column: 9,
							Line:   97,
						},
					},
				},
				Name: "promqlYear",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   108,
					},
					File:   "promql.flux",
					Source: "promql = {\n  promqlDayOfMonth:promqlDayOfMonth,\n  promqlDayOfWeek:promqlDayOfWeek,\n  promqlDaysInMonth:promqlDaysInMonth,\n  promqlHour:promqlHour,\n  promqlMinute:promqlMinute,\n  promqlMonth:promqlMonth,\n  promqlYear:promqlYear,\n}",
					Start: ast.Position{
						Column: 1,
						Line:   100,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 7,
							Line:   100,
						},
						File:   "promql.flux",
						Source: "promql",
						Start: ast.Position{
							Column: 1,
							Line:   100,
						},
					},
				},
				Name: "promql",
			},
			Init: &ast.ObjectExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   108,
						},
						File:   "promql.flux",
						Source: "{\n  promqlDayOfMonth:promqlDayOfMonth,\n  promqlDayOfWeek:promqlDayOfWeek,\n  promqlDaysInMonth:promqlDaysInMonth,\n  promqlHour:promqlHour,\n  promqlMinute:promqlMinute,\n  promqlMonth:promqlMonth,\n  promqlYear:promqlYear,\n}",
						Start: ast.Position{
							Column: 10,
							Line:   100,
						},
					},
				},
				Properties: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 36,
								Line:   101,
							},
							File:   "promql.flux",
							Source: "promqlDayOfMonth:promqlDayOfMonth",
							Start: ast.Position{
								Column: 3,
								Line:   101,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 19,
									Line:   101,
								},
								File:   "promql.flux",
								Source: "promqlDayOfMonth",
								Start: ast.Position{
									Column: 3,
									Line:   101,
								},
							},
						},
						Name: "promqlDayOfMonth",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 36,
									Line:   101,
								},
								File:   "promql.flux",
								Source: "promqlDayOfMonth",
								Start: ast.Position{
									Column: 20,
									Line:   101,
								},
							},
						},
						Name: "promqlDayOfMonth",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 34,
								Line:   102,
							},
							File:   "promql.flux",
							Source: "promqlDayOfWeek:promqlDayOfWeek",
							Start: ast.Position{
								Column: 3,
								Line:   102,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 18,
									Line:   102,
								},
								File:   "promql.flux",
								Source: "promqlDayOfWeek",
								Start: ast.Position{
									Column: 3,
									Line:   102,
								},
							},
						},
						Name: "promqlDayOfWeek",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 34,
									Line:   102,
								},
								File:   "promql.flux",
								Source: "promqlDayOfWeek",
								Start: ast.Position{
									Column: 19,
									Line:   102,
								},
							},
						},
						Name: "promqlDayOfWeek",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 38,
								Line:   103,
							},
							File:   "promql.flux",
							Source: "promqlDaysInMonth:promqlDaysInMonth",
							Start: ast.Position{
								Column: 3,
								Line:   103,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 20,
									Line:   103,
								},
								File:   "promql.flux",
								Source: "promqlDaysInMonth",
								Start: ast.Position{
									Column: 3,
									Line:   103,
								},
							},
						},
						Name: "promqlDaysInMonth",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   103,
								},
								File:   "promql.flux",
								Source: "promqlDaysInMonth",
								Start: ast.Position{
									Column: 21,
									Line:   103,
								},
							},
						},
						Name: "promqlDaysInMonth",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 24,
								Line:   104,
							},
							File:   "promql.flux",
							Source: "promqlHour:promqlHour",
							Start: ast.Position{
								Column: 3,
								Line:   104,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 13,
									Line:   104,
								},
								File:   "promql.flux",
								Source: "promqlHour",
								Start: ast.Position{
									Column: 3,
									Line:   104,
								},
							},
						},
						Name: "promqlHour",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 24,
									Line:   104,
								},
								File:   "promql.flux",
								Source: "promqlHour",
								Start: ast.Position{
									Column: 14,
									Line:   104,
								},
							},
						},
						Name: "promqlHour",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 28,
								Line:   105,
							},
							File:   "promql.flux",
							Source: "promqlMinute:promqlMinute",
							Start: ast.Position{
								Column: 3,
								Line:   105,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 15,
									Line:   105,
								},
								File:   "promql.flux",
								Source: "promqlMinute",
								Start: ast.Position{
									Column: 3,
									Line:   105,
								},
							},
						},
						Name: "promqlMinute",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 28,
									Line:   105,
								},
								File:   "promql.flux",
								Source: "promqlMinute",
								Start: ast.Position{
									Column: 16,
									Line:   105,
								},
							},
						},
						Name: "promqlMinute",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 26,
								Line:   106,
							},
							File:   "promql.flux",
							Source: "promqlMonth:promqlMonth",
							Start: ast.Position{
								Column: 3,
								Line:   106,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 14,
									Line:   106,
								},
								File:   "promql.flux",
								Source: "promqlMonth",
								Start: ast.Position{
									Column: 3,
									Line:   106,
								},
							},
						},
						Name: "promqlMonth",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 26,
									Line:   106,
								},
								File:   "promql.flux",
								Source: "promqlMonth",
								Start: ast.Position{
									Column: 15,
									Line:   106,
								},
							},
						},
						Name: "promqlMonth",
					},
				}, &ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 24,
								Line:   107,
							},
							File:   "promql.flux",
							Source: "promqlYear:promqlYear",
							Start: ast.Position{
								Column: 3,
								Line:   107,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 13,
									Line:   107,
								},
								File:   "promql.flux",
								Source: "promqlYear",
								Start: ast.Position{
									Column: 3,
									Line:   107,
								},
							},
						},
						Name: "promqlYear",
					},
					Value: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 24,
									Line:   107,
								},
								File:   "promql.flux",
								Source: "promqlYear",
								Start: ast.Position{
									Column: 14,
									Line:   107,
								},
							},
						},
						Name: "promqlYear",
					},
				}},
				With: nil,
			},
		}},
		Imports: nil,
		Name:    "promql.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 15,
						Line:   1,
					},
					File:   "promql.flux",
					Source: "package promql",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 15,
							Line:   1,
						},
						File:   "promql.flux",
						Source: "promql",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "promql",
			},
		},
	}},
	Package: "promql",
	Path:    "promql",
}
