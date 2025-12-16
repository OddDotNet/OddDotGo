package builder

import (
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	resourcev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/resource/v1"
	tracev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/trace/v1"
	otlptracev1 "github.com/OddDotNet/OddDotGo/gen/opentelemetry/proto/trace/v1"
)

// WhereSpanFilterConfigurator provides methods to add filters for span queries.
type WhereSpanFilterConfigurator struct {
	filters []*tracev1.Where

	Event                *WhereSpanEventFilterConfigurator
	Link                 *WhereSpanLinkFilterConfigurator
	Resource             *WhereSpanResourceFilterConfigurator
	Status               *WhereSpanStatusFilterConfigurator
	InstrumentationScope *WhereSpanInstrumentationScopeFilterConfigurator
}

func newWhereSpanFilterConfigurator() *WhereSpanFilterConfigurator {
	c := &WhereSpanFilterConfigurator{
		filters: make([]*tracev1.Where, 0),
	}
	c.Event = &WhereSpanEventFilterConfigurator{parent: c}
	c.Link = &WhereSpanLinkFilterConfigurator{parent: c}
	c.Resource = &WhereSpanResourceFilterConfigurator{parent: c}
	c.Status = &WhereSpanStatusFilterConfigurator{parent: c}
	c.InstrumentationScope = &WhereSpanInstrumentationScopeFilterConfigurator{parent: c}
	return c
}

// AddNameFilter adds a filter for the span name.
func (c *WhereSpanFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Name{
					Name: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddKindFilter adds a filter for the span kind.
func (c *WhereSpanFilterConfigurator) AddKindFilter(compare otlptracev1.Span_SpanKind, compareAs commonv1.EnumCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Kind{
					Kind: &tracev1.SpanKindProperty{
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
func (c *WhereSpanFilterConfigurator) AddSpanIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_SpanId{
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

// AddTraceIdFilter adds a filter for the trace ID.
func (c *WhereSpanFilterConfigurator) AddTraceIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_TraceId{
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

// AddParentSpanIdFilter adds a filter for the parent span ID.
func (c *WhereSpanFilterConfigurator) AddParentSpanIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_ParentSpanId{
					ParentSpanId: &commonv1.ByteStringProperty{
						CompareAs: compareAs,
						Compare:   compare,
					},
				},
			},
		},
	})
	return c
}

// AddTraceStateFilter adds a filter for the trace state.
func (c *WhereSpanFilterConfigurator) AddTraceStateFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_TraceState{
					TraceState: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddStartTimeUnixNanoFilter adds a filter for the start time.
func (c *WhereSpanFilterConfigurator) AddStartTimeUnixNanoFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_StartTimeUnixNano{
					StartTimeUnixNano: &commonv1.UInt64Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddEndTimeUnixNanoFilter adds a filter for the end time.
func (c *WhereSpanFilterConfigurator) AddEndTimeUnixNanoFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_EndTimeUnixNano{
					EndTimeUnixNano: &commonv1.UInt64Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddFlagsFilter adds a filter for span flags.
func (c *WhereSpanFilterConfigurator) AddFlagsFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Flags{
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

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count.
func (c *WhereSpanFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_DroppedAttributesCount{
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

// AddDroppedEventsCountFilter adds a filter for dropped events count.
func (c *WhereSpanFilterConfigurator) AddDroppedEventsCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_DroppedEventsCount{
					DroppedEventsCount: &commonv1.UInt32Property{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddDroppedLinksCountFilter adds a filter for dropped links count.
func (c *WhereSpanFilterConfigurator) AddDroppedLinksCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_DroppedLinksCount{
					DroppedLinksCount: &commonv1.UInt32Property{
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
func (c *WhereSpanFilterConfigurator) AddStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Attributes{
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
func (c *WhereSpanFilterConfigurator) AddIntAttributeFilter(key string, compare int64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Attributes{
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
func (c *WhereSpanFilterConfigurator) AddDoubleAttributeFilter(key string, compare float64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Attributes{
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
func (c *WhereSpanFilterConfigurator) AddBoolAttributeFilter(key string, compare bool, compareAs commonv1.BoolCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Attributes{
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
func (c *WhereSpanFilterConfigurator) AddBytesAttributeFilter(key string, compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Attributes{
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
func (c *WhereSpanFilterConfigurator) AddOrFilter(configure func(*WhereSpanFilterConfigurator)) *WhereSpanFilterConfigurator {
	orConfigurator := newWhereSpanFilterConfigurator()
	configure(orConfigurator)

	c.filters = append(c.filters, &tracev1.Where{
		Value: &tracev1.Where_Or{
			Or: &tracev1.OrFilter{
				Filters: orConfigurator.filters,
			},
		},
	})
	return c
}

// WhereSpanEventFilterConfigurator provides methods to add filters for span events.
type WhereSpanEventFilterConfigurator struct {
	parent *WhereSpanFilterConfigurator
}

// AddNameFilter adds a filter for the event name.
func (c *WhereSpanEventFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Event{
					Event: &tracev1.EventFilter{
						Value: &tracev1.EventFilter_Name{
							Name: &commonv1.StringProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddTimeUnixNanoFilter adds a filter for the event time.
func (c *WhereSpanEventFilterConfigurator) AddTimeUnixNanoFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Event{
					Event: &tracev1.EventFilter{
						Value: &tracev1.EventFilter_TimeUnixNano{
							TimeUnixNano: &commonv1.UInt64Property{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count.
func (c *WhereSpanEventFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Event{
					Event: &tracev1.EventFilter{
						Value: &tracev1.EventFilter_DroppedAttributesCount{
							DroppedAttributesCount: &commonv1.UInt32Property{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// WhereSpanLinkFilterConfigurator provides methods to add filters for span links.
type WhereSpanLinkFilterConfigurator struct {
	parent *WhereSpanFilterConfigurator
}

// AddTraceIdFilter adds a filter for the linked trace ID.
func (c *WhereSpanLinkFilterConfigurator) AddTraceIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Link{
					Link: &tracev1.LinkFilter{
						Value: &tracev1.LinkFilter_TraceId{
							TraceId: &commonv1.ByteStringProperty{
								CompareAs: compareAs,
								Compare:   compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddSpanIdFilter adds a filter for the linked span ID.
func (c *WhereSpanLinkFilterConfigurator) AddSpanIdFilter(compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Link{
					Link: &tracev1.LinkFilter{
						Value: &tracev1.LinkFilter_SpanId{
							SpanId: &commonv1.ByteStringProperty{
								CompareAs: compareAs,
								Compare:   compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddTraceStateFilter adds a filter for the linked trace state.
func (c *WhereSpanLinkFilterConfigurator) AddTraceStateFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Link{
					Link: &tracev1.LinkFilter{
						Value: &tracev1.LinkFilter_TraceState{
							TraceState: &commonv1.StringProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddFlagsFilter adds a filter for the link flags.
func (c *WhereSpanLinkFilterConfigurator) AddFlagsFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Link{
					Link: &tracev1.LinkFilter{
						Value: &tracev1.LinkFilter_Flags{
							Flags: &commonv1.UInt32Property{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count on links.
func (c *WhereSpanLinkFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Link{
					Link: &tracev1.LinkFilter{
						Value: &tracev1.LinkFilter_DroppedAttributesCount{
							DroppedAttributesCount: &commonv1.UInt32Property{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// WhereSpanResourceFilterConfigurator provides methods to add filters for span resources.
type WhereSpanResourceFilterConfigurator struct {
	parent *WhereSpanFilterConfigurator
}

// AddSchemaUrlFilter adds a filter for the resource schema URL.
func (c *WhereSpanResourceFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_ResourceSchemaUrl{
			ResourceSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count on resources.
func (c *WhereSpanResourceFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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
func (c *WhereSpanResourceFilterConfigurator) AddStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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
func (c *WhereSpanResourceFilterConfigurator) AddIntAttributeFilter(key string, compare int64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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
func (c *WhereSpanResourceFilterConfigurator) AddDoubleAttributeFilter(key string, compare float64, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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
func (c *WhereSpanResourceFilterConfigurator) AddBoolAttributeFilter(key string, compare bool, compareAs commonv1.BoolCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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
func (c *WhereSpanResourceFilterConfigurator) AddBytesAttributeFilter(key string, compare []byte, compareAs commonv1.ByteStringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Resource{
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

// WhereSpanStatusFilterConfigurator provides methods to add filters for span status.
type WhereSpanStatusFilterConfigurator struct {
	parent *WhereSpanFilterConfigurator
}

// AddMessageFilter adds a filter for the status message.
func (c *WhereSpanStatusFilterConfigurator) AddMessageFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Status{
					Status: &tracev1.StatusFilter{
						Value: &tracev1.StatusFilter_Message{
							Message: &commonv1.StringProperty{
								CompareAs: compareAs,
								Compare:   &compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// AddCodeFilter adds a filter for the status code.
func (c *WhereSpanStatusFilterConfigurator) AddCodeFilter(compare otlptracev1.Status_StatusCode, compareAs commonv1.EnumCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_Property{
			Property: &tracev1.PropertyFilter{
				Value: &tracev1.PropertyFilter_Status{
					Status: &tracev1.StatusFilter{
						Value: &tracev1.StatusFilter_Code{
							Code: &tracev1.SpanStatusCodeProperty{
								CompareAs: compareAs,
								Compare:   compare,
							},
						},
					},
				},
			},
		},
	})
	return c.parent
}

// WhereSpanInstrumentationScopeFilterConfigurator provides methods to add filters for instrumentation scope.
type WhereSpanInstrumentationScopeFilterConfigurator struct {
	parent *WhereSpanFilterConfigurator
}

// AddNameFilter adds a filter for the instrumentation scope name.
func (c *WhereSpanInstrumentationScopeFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_InstrumentationScope{
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
func (c *WhereSpanInstrumentationScopeFilterConfigurator) AddVersionFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_InstrumentationScope{
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
func (c *WhereSpanInstrumentationScopeFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_InstrumentationScope{
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
func (c *WhereSpanInstrumentationScopeFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereSpanFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &tracev1.Where{
		Value: &tracev1.Where_InstrumentationScopeSchemaUrl{
			InstrumentationScopeSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}
