package builder

import (
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	logsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/logs/v1"
)

// LogQueryRequestBuilder provides a fluent API for creating LogQueryRequest objects.
type LogQueryRequestBuilder struct {
	request      *logsv1.LogQueryRequest
	configurator *WhereLogFilterConfigurator
}

// NewLogQueryRequestBuilder creates a builder with defaults: TakeFirst, a
// 30-second Wait, and no filters. See the package documentation for the Take
// and Duration contract.
func NewLogQueryRequestBuilder() *LogQueryRequestBuilder {
	return &LogQueryRequestBuilder{
		request: &logsv1.LogQueryRequest{
			Take: &commonv1.Take{
				Value: &commonv1.Take_TakeFirst{
					TakeFirst: &commonv1.TakeFirst{},
				},
			},
			Duration: &commonv1.Duration{
				Milliseconds: defaultDurationMilliseconds,
			},
		},
		configurator: newWhereLogFilterConfigurator(),
	}
}

// TakeFirst configures the query to return as soon as the first matching log
// is found, or when the Wait duration elapses if none is found. This is the
// default.
func (b *LogQueryRequestBuilder) TakeFirst() *LogQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the query to return as soon as count matching logs are
// found, or when the Wait duration elapses with fewer than count found.
func (b *LogQueryRequestBuilder) TakeExact(count int32) *LogQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeExact{
			TakeExact: &commonv1.TakeExact{
				Count: count,
			},
		},
	}
	return b
}

// TakeAll configures the query to collect every matching log seen over the
// whole Wait duration. TakeAll never returns early — the query always blocks
// for the full duration — so prefer TakeFirst or TakeExact with a filter to
// return as soon as a specific log arrives.
func (b *LogQueryRequestBuilder) TakeAll() *LogQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the maximum duration the query blocks for matching logs. A value
// of zero or less selects the sink default of 30 seconds; it does not return
// immediately.
func (b *LogQueryRequestBuilder) Wait(d time.Duration) *LogQueryRequestBuilder {
	ms := int32(d.Milliseconds())
	if ms < 0 {
		ms = 0
	}
	b.request.Duration = &commonv1.Duration{
		Milliseconds: ms,
	}
	return b
}

// Where configures filters for the log query using the provided function.
func (b *LogQueryRequestBuilder) Where(configure func(*WhereLogFilterConfigurator)) *LogQueryRequestBuilder {
	configure(b.configurator)
	return b
}

// Build returns the configured LogQueryRequest.
func (b *LogQueryRequestBuilder) Build() *logsv1.LogQueryRequest {
	b.request.Filters = append(b.request.Filters, b.configurator.filters...)
	return b.request
}
