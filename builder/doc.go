// Package builder provides fluent builders for constructing OddDotNet signal
// query requests (spans, logs, and metrics).
//
// # Query contract: Take and Duration
//
// A query blocks on the sink until either its Take target is met or its
// Duration elapses, whichever happens first. When the query starts, signals
// already buffered by the sink are evaluated immediately; the query then waits
// for newly-arriving signals until it is satisfied or times out.
//
//   - TakeFirst (the default) returns as soon as one matching signal is found.
//   - TakeExact(n) returns as soon as n matching signals are found.
//   - TakeAll has no target count, so it never returns early: the query always
//     blocks for the full Duration and returns every match seen in that window.
//
// Duration is set with Wait. A Duration of zero or less — including never
// calling Wait — selects the sink default of 30 seconds; it does not mean
// "return immediately". Builders start with a 30-second default.
//
// # Waiting for a specific signal
//
// To wait for a particular signal and return the instant it arrives, use
// TakeFirst (or TakeExact) together with a filter — not TakeAll, which would
// always block for the whole Duration:
//
//	req := builder.NewSpanQueryRequestBuilder().
//		TakeFirst().
//		Where(func(c *builder.WhereSpanFilterConfigurator) {
//			c.AddNameFilter(spanName, commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
//		}).
//		Wait(30 * time.Second).
//		Build()
//
//	resp, err := client.Query(ctx, req) // returns as soon as the span arrives
package builder
