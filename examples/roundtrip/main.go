// Command roundtrip demonstrates a full OddDotNet push->query round-trip: it
// exports one span over OTLP to a running OddDotNet sink, then queries it back
// over the span query API using TakeFirst + a name filter + Wait so the query
// returns the instant the span is ingested.
//
// Run it against a sink reachable at -addr (default 127.0.0.1:4317):
//
//	go run . -addr 127.0.0.1:4317
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/OddDotNet/OddDotGo/builder"
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	tracev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/trace/v1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:4317", "OddDotNet OTLP + query gRPC address")
	flag.Parse()

	ctx := context.Background()
	spanName := fmt.Sprintf("roundtrip-%d", time.Now().UnixNano())

	if err := pushSpan(ctx, *addr, spanName); err != nil {
		log.Fatalf("cannot push span: %v", err)
	}

	found, err := querySpan(ctx, *addr, spanName)
	if err != nil {
		log.Fatalf("cannot query span: %v", err)
	}
	if !found {
		log.Fatalf("span %q not found within the wait window", spanName)
	}

	fmt.Printf("round-trip ok: span %q ingested and queried back\n", spanName)
}

// pushSpan exports a single named span to the sink's OTLP receiver. It uses a
// synchronous span processor so the span is exported on End without needing a
// batch flush interval.
func pushSpan(ctx context.Context, addr, spanName string) error {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(addr), otlptracegrpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot create OTLP exporter: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(resource.NewSchemaless(attribute.String("service.name", "odddotgo-roundtrip"))),
	)
	defer func() { _ = tp.Shutdown(ctx) }()

	_, span := tp.Tracer("odddotgo-roundtrip").Start(ctx, spanName)
	span.End()

	if err := tp.ForceFlush(ctx); err != nil {
		return fmt.Errorf("cannot flush span: %w", err)
	}

	return nil
}

// querySpan long-polls the sink for a span named spanName, returning true as
// soon as it arrives. TakeFirst + a name filter make the query return the
// instant the span is ingested, bounded by Wait — unlike TakeAll, which would
// always block for the full duration.
func querySpan(ctx context.Context, addr, spanName string) (bool, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, fmt.Errorf("cannot dial sink: %w", err)
	}
	defer func() { _ = conn.Close() }()

	req := builder.NewSpanQueryRequestBuilder().
		TakeFirst().
		Where(func(c *builder.WhereSpanFilterConfigurator) {
			c.AddNameFilter(spanName, commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Wait(30 * time.Second).
		Build()

	resp, err := tracev1.NewSpanQueryServiceClient(conn).Query(ctx, req)
	if err != nil {
		return false, fmt.Errorf("cannot query spans: %w", err)
	}

	return len(resp.GetSpans()) > 0, nil
}
