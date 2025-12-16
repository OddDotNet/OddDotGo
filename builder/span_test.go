package builder

import (
	"testing"
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	resourcev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/resource/v1"
	tracev1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/trace/v1"
	otlptracev1 "github.com/OddDotNet/OddDotGo/gen/opentelemetry/proto/trace/v1"
)

func TestSpanQueryRequestBuilder_Defaults(t *testing.T) {
	builder := NewSpanQueryRequestBuilder()
	request := builder.Build()

	if request.Take == nil {
		t.Fatal("Take should not be nil")
	}
	if _, ok := request.Take.Value.(*commonv1.Take_TakeFirst); !ok {
		t.Error("Default Take should be TakeFirst")
	}
	if request.Duration == nil {
		t.Fatal("Duration should not be nil")
	}
	if request.Duration.Milliseconds != 30000 {
		t.Errorf("Default duration should be 30000ms, got %d", request.Duration.Milliseconds)
	}
}

func TestSpanQueryRequestBuilder_TakeFirst(t *testing.T) {
	request := NewSpanQueryRequestBuilder().TakeFirst().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeFirst); !ok {
		t.Error("Take should be TakeFirst")
	}
}

func TestSpanQueryRequestBuilder_TakeExact(t *testing.T) {
	request := NewSpanQueryRequestBuilder().TakeExact(5).Build()

	takeExact, ok := request.Take.Value.(*commonv1.Take_TakeExact)
	if !ok {
		t.Fatal("Take should be TakeExact")
	}
	if takeExact.TakeExact.Count != 5 {
		t.Errorf("TakeExact count should be 5, got %d", takeExact.TakeExact.Count)
	}
}

func TestSpanQueryRequestBuilder_TakeAll(t *testing.T) {
	request := NewSpanQueryRequestBuilder().TakeAll().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeAll); !ok {
		t.Error("Take should be TakeAll")
	}
}

func TestSpanQueryRequestBuilder_Wait(t *testing.T) {
	request := NewSpanQueryRequestBuilder().Wait(5 * time.Second).Build()

	if request.Duration.Milliseconds != 5000 {
		t.Errorf("Duration should be 5000ms, got %d", request.Duration.Milliseconds)
	}
}

func TestSpanQueryRequestBuilder_Wait_NegativeDuration(t *testing.T) {
	request := NewSpanQueryRequestBuilder().Wait(-5 * time.Second).Build()

	if request.Duration.Milliseconds != 0 {
		t.Errorf("Negative duration should result in 0ms, got %d", request.Duration.Milliseconds)
	}
}

func TestSpanQueryRequestBuilder_Where_AddNameFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.AddNameFilter("test-span", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*tracev1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	name, ok := prop.Property.Value.(*tracev1.PropertyFilter_Name)
	if !ok {
		t.Fatal("Property filter should be Name")
	}

	if *name.Name.Compare != "test-span" {
		t.Errorf("Expected name 'test-span', got '%s'", *name.Name.Compare)
	}
	if name.Name.CompareAs != commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS {
		t.Error("Expected EQUALS comparison type")
	}
}

func TestSpanQueryRequestBuilder_Where_AddKindFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.AddKindFilter(otlptracev1.Span_SPAN_KIND_SERVER, commonv1.EnumCompareAsType_ENUM_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*tracev1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	kind, ok := prop.Property.Value.(*tracev1.PropertyFilter_Kind)
	if !ok {
		t.Fatal("Property filter should be Kind")
	}

	if kind.Kind.Compare != otlptracev1.Span_SPAN_KIND_SERVER {
		t.Errorf("Expected SPAN_KIND_SERVER, got %v", kind.Kind.Compare)
	}
}

func TestSpanQueryRequestBuilder_Where_AddStringAttributeFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.AddStringAttributeFilter("http.method", "GET", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*tracev1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	attrs, ok := prop.Property.Value.(*tracev1.PropertyFilter_Attributes)
	if !ok {
		t.Fatal("Property filter should be Attributes")
	}

	if len(attrs.Attributes.Values) != 1 {
		t.Fatalf("Expected 1 attribute, got %d", len(attrs.Attributes.Values))
	}

	if attrs.Attributes.Values[0].Key != "http.method" {
		t.Errorf("Expected key 'http.method', got '%s'", attrs.Attributes.Values[0].Key)
	}

	strVal, ok := attrs.Attributes.Values[0].Value.Value.(*commonv1.AnyValueProperty_StringValue)
	if !ok {
		t.Fatal("Attribute value should be StringValue")
	}

	if *strVal.StringValue.Compare != "GET" {
		t.Errorf("Expected value 'GET', got '%s'", *strVal.StringValue.Compare)
	}
}

func TestSpanQueryRequestBuilder_Where_AddOrFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.AddOrFilter(func(or *WhereSpanFilterConfigurator) {
				or.AddNameFilter("span-a", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
				or.AddNameFilter("span-b", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
			})
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	orFilter, ok := request.Filters[0].Value.(*tracev1.Where_Or)
	if !ok {
		t.Fatal("Filter should be an Or filter")
	}

	if len(orFilter.Or.Filters) != 2 {
		t.Fatalf("Expected 2 filters in Or, got %d", len(orFilter.Or.Filters))
	}
}

func TestSpanQueryRequestBuilder_Where_Resource_AddStringAttributeFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.Resource.AddStringAttributeFilter("service.name", "my-service", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	resource, ok := request.Filters[0].Value.(*tracev1.Where_Resource)
	if !ok {
		t.Fatal("Filter should be a Resource filter")
	}

	attrs, ok := resource.Resource.Value.(*resourcev1.ResourceFilter_Attributes)
	if !ok {
		t.Fatal("Resource filter should be Attributes")
	}

	if attrs.Attributes.Values[0].Key != "service.name" {
		t.Errorf("Expected key 'service.name', got '%s'", attrs.Attributes.Values[0].Key)
	}
}

func TestSpanQueryRequestBuilder_Where_Status_AddCodeFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.Status.AddCodeFilter(otlptracev1.Status_STATUS_CODE_ERROR, commonv1.EnumCompareAsType_ENUM_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*tracev1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	status, ok := prop.Property.Value.(*tracev1.PropertyFilter_Status)
	if !ok {
		t.Fatal("Property filter should be Status")
	}

	code, ok := status.Status.Value.(*tracev1.StatusFilter_Code)
	if !ok {
		t.Fatal("Status filter should be Code")
	}

	if code.Code.Compare != otlptracev1.Status_STATUS_CODE_ERROR {
		t.Errorf("Expected STATUS_CODE_ERROR, got %v", code.Code.Compare)
	}
}

func TestSpanQueryRequestBuilder_Where_InstrumentationScope_AddNameFilter(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		Where(func(f *WhereSpanFilterConfigurator) {
			f.InstrumentationScope.AddNameFilter("my-library", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	scope, ok := request.Filters[0].Value.(*tracev1.Where_InstrumentationScope)
	if !ok {
		t.Fatal("Filter should be an InstrumentationScope filter")
	}

	name, ok := scope.InstrumentationScope.Value.(*commonv1.InstrumentationScopeFilter_Name)
	if !ok {
		t.Fatal("InstrumentationScope filter should be Name")
	}

	if *name.Name.Compare != "my-library" {
		t.Errorf("Expected 'my-library', got '%s'", *name.Name.Compare)
	}
}

func TestSpanQueryRequestBuilder_MultipleFilters(t *testing.T) {
	request := NewSpanQueryRequestBuilder().
		TakeAll().
		Wait(10 * time.Second).
		Where(func(f *WhereSpanFilterConfigurator) {
			f.AddNameFilter("test-span", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_CONTAINS)
			f.AddKindFilter(otlptracev1.Span_SPAN_KIND_CLIENT, commonv1.EnumCompareAsType_ENUM_COMPARE_AS_TYPE_EQUALS)
			f.Resource.AddStringAttributeFilter("service.name", "test-service", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeAll); !ok {
		t.Error("Take should be TakeAll")
	}

	if request.Duration.Milliseconds != 10000 {
		t.Errorf("Duration should be 10000ms, got %d", request.Duration.Milliseconds)
	}

	if len(request.Filters) != 3 {
		t.Fatalf("Expected 3 filters, got %d", len(request.Filters))
	}
}

func TestSpanQueryRequestBuilder_Chaining(t *testing.T) {
	// Test that all methods return the builder for chaining
	builder := NewSpanQueryRequestBuilder()

	result := builder.
		TakeFirst().
		TakeExact(10).
		TakeAll().
		Wait(time.Second).
		Where(func(f *WhereSpanFilterConfigurator) {})

	if result != builder {
		t.Error("Methods should return the same builder instance for chaining")
	}
}
