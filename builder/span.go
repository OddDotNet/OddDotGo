package builder

import (
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	tracev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/trace/v1"
)

const defaultDurationMilliseconds = 30000

// SpanQueryRequestBuilder provides a fluent API for creating SpanQueryRequest objects.
type SpanQueryRequestBuilder struct {
	request      *tracev1.SpanQueryRequest
	configurator *WhereSpanFilterConfigurator
}

// NewSpanQueryRequestBuilder creates a new builder with defaults:
// Take: TakeFirst, Duration: 30 seconds, no filters.
func NewSpanQueryRequestBuilder() *SpanQueryRequestBuilder {
	return &SpanQueryRequestBuilder{
		request: &tracev1.SpanQueryRequest{
			Take: &commonv1.Take{
				Value: &commonv1.Take_TakeFirst{
					TakeFirst: &commonv1.TakeFirst{},
				},
			},
			Duration: &commonv1.Duration{
				Milliseconds: defaultDurationMilliseconds,
			},
		},
		configurator: newWhereSpanFilterConfigurator(),
	}
}

// TakeFirst configures the request to take the first matching span.
func (b *SpanQueryRequestBuilder) TakeFirst() *SpanQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the request to take exactly count matching spans.
func (b *SpanQueryRequestBuilder) TakeExact(count int32) *SpanQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeExact{
			TakeExact: &commonv1.TakeExact{
				Count: count,
			},
		},
	}
	return b
}

// TakeAll configures the request to take all matching spans within the duration.
func (b *SpanQueryRequestBuilder) TakeAll() *SpanQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the duration to wait for matching spans.
func (b *SpanQueryRequestBuilder) Wait(d time.Duration) *SpanQueryRequestBuilder {
	ms := int32(d.Milliseconds())
	if ms < 0 {
		ms = 0
	}
	b.request.Duration = &commonv1.Duration{
		Milliseconds: ms,
	}
	return b
}

// Where configures filters for the span query using the provided function.
func (b *SpanQueryRequestBuilder) Where(configure func(*WhereSpanFilterConfigurator)) *SpanQueryRequestBuilder {
	configure(b.configurator)
	return b
}

// Build returns the configured SpanQueryRequest.
func (b *SpanQueryRequestBuilder) Build() *tracev1.SpanQueryRequest {
	b.request.Filters = append(b.request.Filters, b.configurator.filters...)
	return b.request
}
