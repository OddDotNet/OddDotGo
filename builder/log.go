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

// NewLogQueryRequestBuilder creates a new builder with defaults:
// Take: TakeFirst, Duration: 30 seconds, no filters.
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

// TakeFirst configures the request to take the first matching log.
func (b *LogQueryRequestBuilder) TakeFirst() *LogQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeFirst{
			TakeFirst: &commonv1.TakeFirst{},
		},
	}
	return b
}

// TakeExact configures the request to take exactly count matching logs.
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

// TakeAll configures the request to take all matching logs within the duration.
func (b *LogQueryRequestBuilder) TakeAll() *LogQueryRequestBuilder {
	b.request.Take = &commonv1.Take{
		Value: &commonv1.Take_TakeAll{
			TakeAll: &commonv1.TakeAll{},
		},
	}
	return b
}

// Wait sets the duration to wait for matching logs.
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
