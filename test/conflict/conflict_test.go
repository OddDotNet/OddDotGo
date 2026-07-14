// Package conflict is a regression guard for SREI-9047.
//
// OddDotGo must NOT generate its own copy of the OTLP protos. If it does, any Go
// binary that links both OddDotGo and the canonical go.opentelemetry.io/proto/otlp
// packages panics during init with:
//
//	panic: proto: file "opentelemetry/proto/common/v1/common.proto" is already registered
//	       previously from "github.com/OddDotNet/OddDotGo/gen/opentelemetry/proto/common/v1"
//	       currently from "go.opentelemetry.io/proto/otlp/common/v1"
//
// This test package imports BOTH trees in the same binary. If the regression ever
// returns, package init panics and the whole test binary fails before any test runs.
package conflict

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	// OddDotGo's query builder transitively pulls the generated odddotproto packages,
	// which now import the canonical OTLP Go types.
	"github.com/OddDotNet/OddDotGo/builder"

	// The canonical upstream OTLP common package — the other registrant of common.proto.
	otlpcommon "go.opentelemetry.io/proto/otlp/common/v1"
)

// TestWhenLinkingBuilderAndCanonicalOTLP_GivenSharedProtoFiles_ThenNoDuplicateRegistration
// proves both trees coexist in one binary with no duplicate proto-file registration.
func TestWhenLinkingBuilderAndCanonicalOTLP_GivenSharedProtoFiles_ThenNoDuplicateRegistration(t *testing.T) {
	// Reaching this line at all means package init did NOT panic with "already registered".

	// Exercise the OddDotGo side so the linker cannot drop it.
	if b := builder.NewLogQueryRequestBuilder(); b == nil {
		t.Fatal("NewLogQueryRequestBuilder returned nil")
	}

	// Exercise the canonical OTLP side.
	kv := &otlpcommon.KeyValue{Key: "service.name"}
	if kv.GetKey() != "service.name" {
		t.Fatalf("canonical otlp KeyValue not usable: got %q", kv.GetKey())
	}

	// The shared OTLP file must be registered exactly once and resolve to the
	// canonical Go package (otlpcommon), not a vendored OddDotGo copy.
	const path = "opentelemetry/proto/common/v1/common.proto"
	fd, err := protoregistry.GlobalFiles.FindFileByPath(path)
	if err != nil {
		t.Fatalf("FindFileByPath(%q) failed: %v", path, err)
	}
	if got := string(fd.Path()); got != path {
		t.Fatalf("unexpected file path: got %q want %q", got, path)
	}

	// The KeyValue message descriptor must come from that single registered file.
	md := kv.ProtoReflect().Descriptor()
	if got := md.ParentFile().Path(); got != path {
		t.Fatalf("KeyValue descriptor came from %q, want %q", got, path)
	}

	// And the global type registry must resolve the message name exactly once.
	fullName := protoreflect.FullName("opentelemetry.proto.common.v1.KeyValue")
	if _, err := protoregistry.GlobalTypes.FindMessageByName(fullName); err != nil {
		t.Fatalf("FindMessageByName(%q) failed: %v", fullName, err)
	}
}
