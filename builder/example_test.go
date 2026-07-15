package builder_test

import (
	"fmt"
	"time"

	"github.com/OddDotNet/OddDotGo/builder"
	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
)

// ExampleSpanQueryRequestBuilder builds a query that waits for a single span
// with a specific name and returns as soon as it arrives (or after 30s if it
// never does). Pairing TakeFirst with a filter lets the sink's long-poll return
// the instant the span is ingested, instead of TakeAll always blocking for the
// full Wait duration.
func ExampleSpanQueryRequestBuilder() {
	req := builder.NewSpanQueryRequestBuilder().
		TakeFirst().
		Where(func(c *builder.WhereSpanFilterConfigurator) {
			c.AddNameFilter("my-span", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Wait(30 * time.Second).
		Build()

	fmt.Println(req.GetTake().GetTakeFirst() != nil)
	fmt.Println(req.GetDuration().GetMilliseconds())
	// Output:
	// true
	// 30000
}
