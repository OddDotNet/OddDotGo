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

// NewSpanQueryRequestBuilder creates a builder with defaults: TakeFirst, a
// 30-second Wait, and no filters. See the package documentation for the Take
// and Duration contract.
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

// TakeFirst configures the query to return as soon as the first matching span
// is found, or when the Wait duration elapses if none is found. This is the
// default.
func (b *SpanQueryRequestBuilder) TakeFirst() *SpanQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the query to return as soon as count matching spans are
// found, or when the Wait duration elapses with fewer than count found.
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

// TakeAll configures the query to collect every matching span seen over the
// whole Wait duration. TakeAll never returns early — the query always blocks
// for the full duration — so prefer TakeFirst or TakeExact with a filter to
// return as soon as a specific span arrives.
func (b *SpanQueryRequestBuilder) TakeAll() *SpanQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the maximum duration the query blocks for matching spans. A value
// of zero or less selects the sink default of 30 seconds; it does not return
// immediately.
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
