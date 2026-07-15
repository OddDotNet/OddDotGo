package builder

import (
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	logsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/logs/v1"
	resourcev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/resource/v1"
	otlplogsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

// WhereLogFilterConfigurator provides methods to add filters for log queries.
type WhereLogFilterConfigurator struct {
	filters []*logsv1.Where

	Resource             *WhereLogResourceFilterConfigurator
	InstrumentationScope *WhereLogInstrumentationScopeFilterConfigurator
}

func newWhereLogFilterConfigurator() *WhereLogFilterConfigurator {
	c := &WhereLogFilterConfigurator{
		filters: make([]*logsv1.Where, 0),
	}
	c.Resource = &WhereLogResourceFilterConfigurator{parent: c}
	c.InstrumentationScope = &WhereLogInstrumentationScopeFilterConfigurator{parent: c}
	return c
}

// AddTimeUnixNanoFilter adds a filter for the log time.
func (c *WhereLogFilterConfigurator) AddTimeUnixNanoFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_TimeUnixNano{
					TimeUnixNano: &commonv1.UInt64Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddObservedTimeUnixNanoFilter adds a filter for the observed time.
func (c *WhereLogFilterConfigurator) AddObservedTimeUnixNanoFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_ObservedTimeUnixNano{
					ObservedTimeUnixNano: &commonv1.UInt64Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddSeverityNumberFilter adds a filter for the severity number.
func (c *WhereLogFilterConfigurator) AddSeverityNumberFilter(compare otlplogsv1.SeverityNumber, compareAs commonv1.EnumCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_SeverityNumber{
					SeverityNumber: &logsv1.SeverityNumberProperty{
						CompareAs: compareAs,
						Compare:   compare,
					},
				},
			},
		},
	})
	return c
}

// AddSeverityTextFilter adds a filter for the severity text.
func (c *WhereLogFilterConfigurator) AddSeverityTextFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_SeverityText{
					SeverityText: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddStringBodyFilter adds a string body filter.
func (c *WhereLogFilterConfigurator) AddStringBodyFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Body{
					Body: &commonv1.AnyValueProperty{
						Value: &commonv1.AnyValueProperty_StringValue{
							StringValue: &commonv1.StringProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddIntBodyFilter adds an int body filter.
func (c *WhereLogFilterConfigurator) AddIntBodyFilter(compare int64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Body{
					Body: &commonv1.AnyValueProperty{
						Value: &commonv1.AnyValueProperty_IntValue{
							IntValue: &commonv1.Int64Property{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddDoubleBodyFilter adds a double body filter.
func (c *WhereLogFilterConfigurator) AddDoubleBodyFilter(compare float64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Body{
					Body: &commonv1.AnyValueProperty{
						Value: &commonv1.AnyValueProperty_DoubleValue{
							DoubleValue: &commonv1.DoubleProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddBoolBodyFilter adds a bool body filter.
func (c *WhereLogFilterConfigurator) AddBoolBodyFilter(compare bool, compareAs commonv1.BoolCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Body{
					Body: &commonv1.AnyValueProperty{
						Value: &commonv1.AnyValueProperty_BoolValue{
							BoolValue: &commonv1.BoolProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddFlagsFilter adds a filter for log flags.
func (c *WhereLogFilterConfigurator) AddFlagsFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Flags{
					Flags: &commonv1.UInt32Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddTraceIdFilter adds a filter for the trace ID.
func (c *WhereLogFilterConfigurator) AddTraceIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_TraceId{
					TraceId: &commonv1.ByteStringProperty{
						CompareAs: compareAs,
						Compare:   compare,
					},
				},
			},
		},
	})
	return c
}

// AddSpanIdFilter adds a filter for the span ID.
func (c *WhereLogFilterConfigurator) AddSpanIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_SpanId{
					SpanId: &commonv1.ByteStringProperty{
						CompareAs: compareAs,
						Compare:   compare,
					},
				},
			},
		},
	})
	return c
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count.
func (c *WhereLogFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_DroppedAttributesCount{
					DroppedAttributesCount: &commonv1.UInt32Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddStringAttributeFilter adds a string attribute filter.
func (c *WhereLogFilterConfigurator) AddStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_StringValue{
										StringValue: &commonv1.StringProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddIntAttributeFilter adds an int64 attribute filter.
func (c *WhereLogFilterConfigurator) AddIntAttributeFilter(key string, compare int64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_IntValue{
										IntValue: &commonv1.Int64Property{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddDoubleAttributeFilter adds a double attribute filter.
func (c *WhereLogFilterConfigurator) AddDoubleAttributeFilter(key string, compare float64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_DoubleValue{
										DoubleValue: &commonv1.DoubleProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddBoolAttributeFilter adds a bool attribute filter.
func (c *WhereLogFilterConfigurator) AddBoolAttributeFilter(key string, compare bool, compareAs commonv1.BoolCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_BoolValue{
										BoolValue: &commonv1.BoolProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddBytesAttributeFilter adds a byte array attribute filter.
func (c *WhereLogFilterConfigurator) AddBytesAttributeFilter(key string, compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereLogFilterConfigurator {
	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Property{
			Property: &logsv1.PropertyFilter{
				Value: &logsv1.PropertyFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_ByteStringValue{
										ByteStringValue: &commonv1.ByteStringProperty{
											CompareAs: compareAs,
											Compare:   compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c
}

// AddOrFilter adds an OR filter with the given conditions.
func (c *WhereLogFilterConfigurator) AddOrFilter(configure func(*WhereLogFilterConfigurator)) *WhereLogFilterConfigurator {
	orConfigurator := newWhereLogFilterConfigurator()
	configure(orConfigurator)

	c.filters = append(c.filters, &logsv1.Where{
		Value: &logsv1.Where_Or{
			Or: &logsv1.OrFilter{
				Filters: orConfigurator.filters,
			},
		},
	})
	return c
}

// WhereLogResourceFilterConfigurator provides methods to add filters for log resources.
type WhereLogResourceFilterConfigurator struct {
	parent *WhereLogFilterConfigurator
}

// AddSchemaUrlFilter adds a filter for the resource schema URL.
func (c *WhereLogResourceFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_ResourceSchemaUrl{
			ResourceSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count on resources.
func (c *WhereLogResourceFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_DroppedAttributesCount{
					DroppedAttributesCount: &commonv1.UInt32Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c.parent
}

// AddStringAttributeFilter adds a string attribute filter for resources.
func (c *WhereLogResourceFilterConfigurator) AddStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_StringValue{
										StringValue: &commonv1.StringProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddIntAttributeFilter adds an int64 attribute filter for resources.
func (c *WhereLogResourceFilterConfigurator) AddIntAttributeFilter(key string, compare int64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_IntValue{
										IntValue: &commonv1.Int64Property{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddDoubleAttributeFilter adds a double attribute filter for resources.
func (c *WhereLogResourceFilterConfigurator) AddDoubleAttributeFilter(key string, compare float64, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_DoubleValue{
										DoubleValue: &commonv1.DoubleProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddBoolAttributeFilter adds a bool attribute filter for resources.
func (c *WhereLogResourceFilterConfigurator) AddBoolAttributeFilter(key string, compare bool, compareAs commonv1.BoolCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_BoolValue{
										BoolValue: &commonv1.BoolProperty{
											CompareAs: compareAs,
											Compare:   &compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddBytesAttributeFilter adds a byte array attribute filter for resources.
func (c *WhereLogResourceFilterConfigurator) AddBytesAttributeFilter(key string, compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_Resource{
			Resource: &resourcev1.ResourceFilter{
				Value: &resourcev1.ResourceFilter_Attributes{
					Attributes: &commonv1.KeyValueListProperty{
						Values: []*commonv1.KeyValueProperty{
							{
								Key: key,
								Value: &commonv1.AnyValueProperty{
									Value: &commonv1.AnyValueProperty_ByteStringValue{
										ByteStringValue: &commonv1.ByteStringProperty{
											CompareAs: compareAs,
											Compare:   compare,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// WhereLogInstrumentationScopeFilterConfigurator provides methods to add filters for instrumentation scope.
type WhereLogInstrumentationScopeFilterConfigurator struct {
	parent *WhereLogFilterConfigurator
}

// AddNameFilter adds a filter for the instrumentation scope name.
func (c *WhereLogInstrumentationScopeFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_InstrumentationScope{
			InstrumentationScope: &commonv1.InstrumentationScopeFilter{
				Value: &commonv1.InstrumentationScopeFilter_Name{
					Name: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c.parent
}

// AddVersionFilter adds a filter for the instrumentation scope version.
func (c *WhereLogInstrumentationScopeFilterConfigurator) AddVersionFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_InstrumentationScope{
			InstrumentationScope: &commonv1.InstrumentationScopeFilter{
				Value: &commonv1.InstrumentationScopeFilter_Version{
					Version: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count on instrumentation scope.
func (c *WhereLogInstrumentationScopeFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_InstrumentationScope{
			InstrumentationScope: &commonv1.InstrumentationScopeFilter{
				Value: &commonv1.InstrumentationScopeFilter_DroppedAttributesCount{
					DroppedAttributesCount: &commonv1.UInt32Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c.parent
}

// AddSchemaUrlFilter adds a filter for the instrumentation scope schema URL.
func (c *WhereLogInstrumentationScopeFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereLogFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &logsv1.Where{
		Value: &logsv1.Where_InstrumentationScopeSchemaUrl{
			InstrumentationScopeSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}
