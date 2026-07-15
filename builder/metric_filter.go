package builder

import (
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	metricsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/metrics/v1"
	resourcev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/resource/v1"
	otlpmetricsv1 "go.opentelemetry.io/proto/otlp/metrics/v1"
)

// WhereMetricFilterConfigurator provides methods to add filters for metric queries.
type WhereMetricFilterConfigurator struct {
	filters []*metricsv1.Where

	Resource             *WhereMetricResourceFilterConfigurator
	InstrumentationScope *WhereMetricInstrumentationScopeFilterConfigurator
	Gauge                *WhereMetricGaugeFilterConfigurator
	Sum                  *WhereMetricSumFilterConfigurator
	Histogram            *WhereMetricHistogramFilterConfigurator
	ExponentialHistogram *WhereMetricExponentialHistogramFilterConfigurator
	Summary              *WhereMetricSummaryFilterConfigurator
}

func newWhereMetricFilterConfigurator() *WhereMetricFilterConfigurator {
	c := &WhereMetricFilterConfigurator{
		filters: make([]*metricsv1.Where, 0),
	}
	c.Resource = &WhereMetricResourceFilterConfigurator{parent: c}
	c.InstrumentationScope = &WhereMetricInstrumentationScopeFilterConfigurator{parent: c}
	c.Gauge = &WhereMetricGaugeFilterConfigurator{parent: c}
	c.Sum = &WhereMetricSumFilterConfigurator{parent: c}
	c.Histogram = &WhereMetricHistogramFilterConfigurator{parent: c}
	c.ExponentialHistogram = &WhereMetricExponentialHistogramFilterConfigurator{parent: c}
	c.Summary = &WhereMetricSummaryFilterConfigurator{parent: c}
	return c
}

// AddNameFilter adds a filter for the metric name.
func (c *WhereMetricFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.filters = append(c.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Name{
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

// AddDescriptionFilter adds a filter for the metric description.
func (c *WhereMetricFilterConfigurator) AddDescriptionFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.filters = append(c.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Description{
					Description: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddUnitFilter adds a filter for the metric unit.
func (c *WhereMetricFilterConfigurator) AddUnitFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.filters = append(c.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Unit{
					Unit: &commonv1.StringProperty{
						CompareAs: compareAs,
						Compare:   &compare,
					},
				},
			},
		},
	})
	return c
}

// AddOrFilter adds an OR filter with the given conditions.
func (c *WhereMetricFilterConfigurator) AddOrFilter(configure func(*WhereMetricFilterConfigurator)) *WhereMetricFilterConfigurator {
	orConfigurator := newWhereMetricFilterConfigurator()
	configure(orConfigurator)

	c.filters = append(c.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Or{
			Or: &metricsv1.OrFilter{
				Filters: orConfigurator.filters,
			},
		},
	})
	return c
}

// WhereMetricResourceFilterConfigurator provides methods to add filters for metric resources.
type WhereMetricResourceFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddSchemaUrlFilter adds a filter for the resource schema URL.
func (c *WhereMetricResourceFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_ResourceSchemaUrl{
			ResourceSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}

// AddDroppedAttributesCountFilter adds a filter for dropped attributes count on resources.
func (c *WhereMetricResourceFilterConfigurator) AddDroppedAttributesCountFilter(compare uint32, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Resource{
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
func (c *WhereMetricResourceFilterConfigurator) AddStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Resource{
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
func (c *WhereMetricResourceFilterConfigurator) AddIntAttributeFilter(key string, compare int64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Resource{
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

// WhereMetricInstrumentationScopeFilterConfigurator provides methods to add filters for instrumentation scope.
type WhereMetricInstrumentationScopeFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddNameFilter adds a filter for the instrumentation scope name.
func (c *WhereMetricInstrumentationScopeFilterConfigurator) AddNameFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_InstrumentationScope{
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
func (c *WhereMetricInstrumentationScopeFilterConfigurator) AddVersionFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_InstrumentationScope{
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

// AddSchemaUrlFilter adds a filter for the instrumentation scope schema URL.
func (c *WhereMetricInstrumentationScopeFilterConfigurator) AddSchemaUrlFilter(compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_InstrumentationScopeSchemaUrl{
			InstrumentationScopeSchemaUrl: &commonv1.StringProperty{
				CompareAs: compareAs,
				Compare:   &compare,
			},
		},
	})
	return c.parent
}

// WhereMetricGaugeFilterConfigurator provides methods to add filters for gauge metrics.
type WhereMetricGaugeFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddDataPointDoubleValueFilter adds a filter for gauge data point double value.
func (c *WhereMetricGaugeFilterConfigurator) AddDataPointDoubleValueFilter(compare float64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Gauge{
					Gauge: &metricsv1.GaugeFilter{
						Value: &metricsv1.GaugeFilter_DataPoint{
							DataPoint: &metricsv1.NumberDataPointFilter{
								Value: &metricsv1.NumberDataPointFilter_ValueAsDouble{
									ValueAsDouble: &commonv1.DoubleProperty{
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
	})
	return c.parent
}

// AddDataPointIntValueFilter adds a filter for gauge data point int value.
func (c *WhereMetricGaugeFilterConfigurator) AddDataPointIntValueFilter(compare int64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Gauge{
					Gauge: &metricsv1.GaugeFilter{
						Value: &metricsv1.GaugeFilter_DataPoint{
							DataPoint: &metricsv1.NumberDataPointFilter{
								Value: &metricsv1.NumberDataPointFilter_ValueAsInt{
									ValueAsInt: &commonv1.Int64Property{
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
	})
	return c.parent
}

// AddDataPointStringAttributeFilter adds a string attribute filter for gauge data points.
func (c *WhereMetricGaugeFilterConfigurator) AddDataPointStringAttributeFilter(key string, compare string, compareAs commonv1.StringCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Gauge{
					Gauge: &metricsv1.GaugeFilter{
						Value: &metricsv1.GaugeFilter_DataPoint{
							DataPoint: &metricsv1.NumberDataPointFilter{
								Value: &metricsv1.NumberDataPointFilter_Attributes{
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
					},
				},
			},
		},
	})
	return c.parent
}

// WhereMetricSumFilterConfigurator provides methods to add filters for sum metrics.
type WhereMetricSumFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddAggregationTemporalityFilter adds a filter for aggregation temporality.
func (c *WhereMetricSumFilterConfigurator) AddAggregationTemporalityFilter(compare otlpmetricsv1.AggregationTemporality, compareAs commonv1.EnumCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Sum{
					Sum: &metricsv1.SumFilter{
						Value: &metricsv1.SumFilter_AggregationTemporality{
							AggregationTemporality: &metricsv1.AggregationTemporalityProperty{
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

// AddIsMonotonicFilter adds a filter for is_monotonic.
func (c *WhereMetricSumFilterConfigurator) AddIsMonotonicFilter(compare bool, compareAs commonv1.BoolCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Sum{
					Sum: &metricsv1.SumFilter{
						Value: &metricsv1.SumFilter_IsMonotonic{
							IsMonotonic: &commonv1.BoolProperty{
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

// AddDataPointDoubleValueFilter adds a filter for sum data point double value.
func (c *WhereMetricSumFilterConfigurator) AddDataPointDoubleValueFilter(compare float64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Sum{
					Sum: &metricsv1.SumFilter{
						Value: &metricsv1.SumFilter_DataPoint{
							DataPoint: &metricsv1.NumberDataPointFilter{
								Value: &metricsv1.NumberDataPointFilter_ValueAsDouble{
									ValueAsDouble: &commonv1.DoubleProperty{
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
	})
	return c.parent
}

// AddDataPointIntValueFilter adds a filter for sum data point int value.
func (c *WhereMetricSumFilterConfigurator) AddDataPointIntValueFilter(compare int64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Sum{
					Sum: &metricsv1.SumFilter{
						Value: &metricsv1.SumFilter_DataPoint{
							DataPoint: &metricsv1.NumberDataPointFilter{
								Value: &metricsv1.NumberDataPointFilter_ValueAsInt{
									ValueAsInt: &commonv1.Int64Property{
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
	})
	return c.parent
}

// WhereMetricHistogramFilterConfigurator provides methods to add filters for histogram metrics.
type WhereMetricHistogramFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddAggregationTemporalityFilter adds a filter for aggregation temporality.
func (c *WhereMetricHistogramFilterConfigurator) AddAggregationTemporalityFilter(compare otlpmetricsv1.AggregationTemporality, compareAs commonv1.EnumCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Histogram{
					Histogram: &metricsv1.HistogramFilter{
						Value: &metricsv1.HistogramFilter_AggregationTemporality{
							AggregationTemporality: &metricsv1.AggregationTemporalityProperty{
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

// AddDataPointCountFilter adds a filter for histogram data point count.
func (c *WhereMetricHistogramFilterConfigurator) AddDataPointCountFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Histogram{
					Histogram: &metricsv1.HistogramFilter{
						Value: &metricsv1.HistogramFilter_DataPoint{
							DataPoint: &metricsv1.HistogramDataPointFilter{
								Value: &metricsv1.HistogramDataPointFilter_Count{
									Count: &commonv1.UInt64Property{
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
	})
	return c.parent
}

// AddDataPointSumFilter adds a filter for histogram data point sum.
func (c *WhereMetricHistogramFilterConfigurator) AddDataPointSumFilter(compare float64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Histogram{
					Histogram: &metricsv1.HistogramFilter{
						Value: &metricsv1.HistogramFilter_DataPoint{
							DataPoint: &metricsv1.HistogramDataPointFilter{
								Value: &metricsv1.HistogramDataPointFilter_Sum{
									Sum: &commonv1.DoubleProperty{
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
	})
	return c.parent
}

// WhereMetricExponentialHistogramFilterConfigurator provides methods to add filters for exponential histogram metrics.
type WhereMetricExponentialHistogramFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddAggregationTemporalityFilter adds a filter for aggregation temporality.
func (c *WhereMetricExponentialHistogramFilterConfigurator) AddAggregationTemporalityFilter(compare otlpmetricsv1.AggregationTemporality, compareAs commonv1.EnumCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_ExponentialHistogram{
					ExponentialHistogram: &metricsv1.ExponentialHistogramFilter{
						Value: &metricsv1.ExponentialHistogramFilter_AggregationTemporality{
							AggregationTemporality: &metricsv1.AggregationTemporalityProperty{
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

// AddDataPointCountFilter adds a filter for exponential histogram data point count.
func (c *WhereMetricExponentialHistogramFilterConfigurator) AddDataPointCountFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_ExponentialHistogram{
					ExponentialHistogram: &metricsv1.ExponentialHistogramFilter{
						Value: &metricsv1.ExponentialHistogramFilter_DataPoint{
							DataPoint: &metricsv1.ExponentialHistogramDataPointFilter{
								Value: &metricsv1.ExponentialHistogramDataPointFilter_Count{
									Count: &commonv1.UInt64Property{
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
	})
	return c.parent
}

// WhereMetricSummaryFilterConfigurator provides methods to add filters for summary metrics.
type WhereMetricSummaryFilterConfigurator struct {
	parent *WhereMetricFilterConfigurator
}

// AddDataPointCountFilter adds a filter for summary data point count.
func (c *WhereMetricSummaryFilterConfigurator) AddDataPointCountFilter(compare uint64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Summary{
					Summary: &metricsv1.SummaryFilter{
						Value: &metricsv1.SummaryFilter_DataPoint{
							DataPoint: &metricsv1.SummaryDataPointFilter{
								Value: &metricsv1.SummaryDataPointFilter_Count{
									Count: &commonv1.UInt64Property{
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
	})
	return c.parent
}

// AddDataPointSumFilter adds a filter for summary data point sum.
func (c *WhereMetricSummaryFilterConfigurator) AddDataPointSumFilter(compare float64, compareAs commonv1.NumberCompareAsType) *WhereMetricFilterConfigurator {
	c.parent.filters = append(c.parent.filters, &metricsv1.Where{
		Value: &metricsv1.Where_Property{
			Property: &metricsv1.PropertyFilter{
				Value: &metricsv1.PropertyFilter_Summary{
					Summary: &metricsv1.SummaryFilter{
						Value: &metricsv1.SummaryFilter_DataPoint{
							DataPoint: &metricsv1.SummaryDataPointFilter{
								Value: &metricsv1.SummaryDataPointFilter_Sum{
									Sum: &commonv1.DoubleProperty{
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
	})
	return c.parent
}
