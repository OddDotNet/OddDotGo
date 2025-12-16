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

// NewMetricQueryRequestBuilder creates a new builder with defaults:
// Take: TakeFirst, Duration: 30 seconds, no filters.
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

// TakeFirst configures the request to take the first matching metric.
func (b *MetricQueryRequestBuilder) TakeFirst() *MetricQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the request to take exactly count matching metrics.
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

// TakeAll configures the request to take all matching metrics within the duration.
func (b *MetricQueryRequestBuilder) TakeAll() *MetricQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the duration to wait for matching metrics.
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
