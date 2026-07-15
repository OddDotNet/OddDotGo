# OddDotGo

Go client and query builders for [OddDotNet](https://github.com/OddDotNet/OddDotNet),
the OpenTelemetry test harness. Push spans, logs, and metrics to a running
OddDotNet sink over OTLP, then query them back to assert on what your service
emitted.

```bash
go get github.com/OddDotNet/OddDotGo
```

## Quickstart

Build a query with the fluent builders in [`builder`](./builder) and send it with
the generated gRPC client in [`gen`](./gen):

```go
import (
	"github.com/OddDotNet/OddDotGo/builder"
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	tracev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/trace/v1"
)

req := builder.NewSpanQueryRequestBuilder().
	TakeFirst().
	Where(func(c *builder.WhereSpanFilterConfigurator) {
		c.AddNameFilter("my-span", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
	}).
	Wait(30 * time.Second).
	Build()

resp, err := tracev1.NewSpanQueryServiceClient(conn).Query(ctx, req)
```

A full push→query round-trip lives in [`examples/roundtrip`](./examples/roundtrip).

## The query contract: Take and Duration

A query **blocks on the sink** until either its `Take` target is met or its
`Duration` elapses, whichever happens first. When a query starts, signals
already buffered by the sink are evaluated immediately; the query then waits for
newly-arriving signals until it is satisfied or times out.

| `Take` | Returns |
| --- | --- |
| `TakeFirst` (default) | as soon as **1** matching signal is found, else at `Duration` |
| `TakeExact(n)` | as soon as **n** matching signals are found, else at `Duration` |
| `TakeAll` | **never early** — always blocks the full `Duration`, returning every match seen in that window |

`Duration` is set with `Wait`. **A `Duration` of zero or less — including never
calling `Wait` — selects the sink default of 30 seconds; it does not mean
"return immediately."** Builders start with a 30-second default.

### Wait for a specific signal — use `TakeFirst` + a filter

To wait for one signal and return the instant it arrives, filter for it and take
the first match. The sink's long-poll returns as soon as the signal is ingested:

```go
req := builder.NewSpanQueryRequestBuilder().
	TakeFirst().
	Where(func(c *builder.WhereSpanFilterConfigurator) {
		c.AddNameFilter(spanName, commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
	}).
	Wait(30 * time.Second). // upper bound; returns early on arrival
	Build()
```

**Anti-pattern:** `TakeAll().Wait(30 * time.Second)` combined with client-side
retry polling. `TakeAll` has no target count, so it *always* blocks for the full
`Wait` duration — a single call takes the whole 30s even if the signal arrived
immediately, and wrapping it in a retry loop just stacks waits. Reach for
`TakeFirst`/`TakeExact` + a filter instead.

## Development

Generated code under [`gen`](./gen) is produced from the
[OddDotProto](https://github.com/OddDotNet/OddDotProto) submodule via
[`buf`](https://buf.build):

```bash
make generate   # rm -rf gen && buf generate
make build
make test
make lint
```
