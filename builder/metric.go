package builder

import (
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	metricsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/metrics/v1"
)

// MetricQueryRequestBuilder provides a fluent API for creating MetricQueryRequest objects.
type MetricQueryRequestBuilder struct {
	request      *metricsv1.MetricQueryRequest
	configurator *WhereMetricFilterConfigurator
}

// NewMetricQueryRequestBuilder creates a builder with defaults: TakeFirst, a
// 30-second Wait, and no filters. See the package documentation for the Take
// and Duration contract.
func NewMetricQueryRequestBuilder() *MetricQueryRequestBuilder {
	return &MetricQueryRequestBuilder{
		request: &metricsv1.MetricQueryRequest{
			Take: &commonv1.Take{
				Value: &commonv1.Take_TakeFirst{
					TakeFirst: &commonv1.TakeFirst{},
				},
			},
			Duration: &commonv1.Duration{
				Milliseconds: defaultDurationMilliseconds,
			},
		},
		configurator: newWhereMetricFilterConfigurator(),
	}
}

// TakeFirst configures the query to return as soon as the first matching metric
// is found, or when the Wait duration elapses if none is found. This is the
// default.
func (b *MetricQueryRequestBuilder) TakeFirst() *MetricQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the query to return as soon as count matching metrics
// are found, or when the Wait duration elapses with fewer than count found.
func (b *MetricQueryRequestBuilder) TakeExact(count int32) *MetricQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeExact{
			TakeExact: &commonv1.TakeExact{
				Count: count,
			},
		},
	}
	return b
}

// TakeAll configures the query to collect every matching metric seen over the
// whole Wait duration. TakeAll never returns early — the query always blocks
// for the full duration — so prefer TakeFirst or TakeExact with a filter to
// return as soon as a specific metric arrives.
func (b *MetricQueryRequestBuilder) TakeAll() *MetricQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the maximum duration the query blocks for matching metrics. A value
// of zero or less selects the sink default of 30 seconds; it does not return
// immediately.
func (b *MetricQueryRequestBuilder) Wait(d time.Duration) *MetricQueryRequestBuilder {
	ms := int32(d.Milliseconds())
	if ms < 0 {
		ms = 0
	}
	b.request.Duration = &commonv1.Duration{
		Milliseconds: ms,
	}
	return b
}

// Where configures filters for the metric query using the provided function.
func (b *MetricQueryRequestBuilder) Where(configure func(*WhereMetricFilterConfigurator)) *MetricQueryRequestBuilder {
	configure(b.configurator)
	return b
}

// Build returns the configured MetricQueryRequest.
func (b *MetricQueryRequestBuilder) Build() *metricsv1.MetricQueryRequest {
	b.request.Filters = append(b.request.Filters, b.configurator.filters...)
	return b.request
}
