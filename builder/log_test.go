package builder

import (
	"testing"
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	logsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/logs/v1"
	otlplogsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

func TestLogQueryRequestBuilder_Defaults(t *testing.T) {
	builder := NewLogQueryRequestBuilder()
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

func TestLogQueryRequestBuilder_TakeFirst(t *testing.T) {
	request := NewLogQueryRequestBuilder().TakeFirst().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeFirst); !ok {
		t.Error("Take should be TakeFirst")
	}
}

func TestLogQueryRequestBuilder_TakeExact(t *testing.T) {
	request := NewLogQueryRequestBuilder().TakeExact(10).Build()

	takeExact, ok := request.Take.Value.(*commonv1.Take_TakeExact)
	if !ok {
		t.Fatal("Take should be TakeExact")
	}
	if takeExact.TakeExact.Count != 10 {
		t.Errorf("TakeExact count should be 10, got %d", takeExact.TakeExact.Count)
	}
}

func TestLogQueryRequestBuilder_TakeAll(t *testing.T) {
	request := NewLogQueryRequestBuilder().TakeAll().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeAll); !ok {
		t.Error("Take should be TakeAll")
	}
}

func TestLogQueryRequestBuilder_Wait(t *testing.T) {
	request := NewLogQueryRequestBuilder().Wait(3 * time.Second).Build()

	if request.Duration.Milliseconds != 3000 {
		t.Errorf("Duration should be 3000ms, got %d", request.Duration.Milliseconds)
	}
}

func TestLogQueryRequestBuilder_Where_AddSeverityNumberFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddSeverityNumberFilter(otlplogsv1.SeverityNumber_SEVERITY_NUMBER_ERROR, commonv1.EnumCompareAsType_ENUM_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*logsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	severity, ok := prop.Property.Value.(*logsv1.PropertyFilter_SeverityNumber)
	if !ok {
		t.Fatal("Property filter should be SeverityNumber")
	}

	if severity.SeverityNumber.Compare != otlplogsv1.SeverityNumber_SEVERITY_NUMBER_ERROR {
		t.Errorf("Expected SEVERITY_NUMBER_ERROR, got %v", severity.SeverityNumber.Compare)
	}
}

func TestLogQueryRequestBuilder_Where_AddSeverityTextFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddSeverityTextFilter("ERROR", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*logsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	severity, ok := prop.Property.Value.(*logsv1.PropertyFilter_SeverityText)
	if !ok {
		t.Fatal("Property filter should be SeverityText")
	}

	if *severity.SeverityText.Compare != "ERROR" {
		t.Errorf("Expected 'ERROR', got '%s'", *severity.SeverityText.Compare)
	}
}

func TestLogQueryRequestBuilder_Where_AddStringBodyFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddStringBodyFilter("error occurred", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_CONTAINS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*logsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	body, ok := prop.Property.Value.(*logsv1.PropertyFilter_Body)
	if !ok {
		t.Fatal("Property filter should be Body")
	}

	strVal, ok := body.Body.Value.(*commonv1.AnyValueProperty_StringValue)
	if !ok {
		t.Fatal("Body should be StringValue")
	}

	if *strVal.StringValue.Compare != "error occurred" {
		t.Errorf("Expected 'error occurred', got '%s'", *strVal.StringValue.Compare)
	}
}

func TestLogQueryRequestBuilder_Where_AddTraceIdFilter(t *testing.T) {
	traceId := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10}

	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddTraceIdFilter(traceId, commonv1.ByteStringCompareAsType_BYTE_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*logsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	trace, ok := prop.Property.Value.(*logsv1.PropertyFilter_TraceId)
	if !ok {
		t.Fatal("Property filter should be TraceId")
	}

	if len(trace.TraceId.Compare) != 16 {
		t.Errorf("TraceId should have 16 bytes, got %d", len(trace.TraceId.Compare))
	}
}

func TestLogQueryRequestBuilder_Where_AddStringAttributeFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddStringAttributeFilter("log.source", "application", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*logsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	attrs, ok := prop.Property.Value.(*logsv1.PropertyFilter_Attributes)
	if !ok {
		t.Fatal("Property filter should be Attributes")
	}

	if attrs.Attributes.Values[0].Key != "log.source" {
		t.Errorf("Expected key 'log.source', got '%s'", attrs.Attributes.Values[0].Key)
	}
}

func TestLogQueryRequestBuilder_Where_AddOrFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.AddOrFilter(func(or *WhereLogFilterConfigurator) {
				or.AddSeverityTextFilter("ERROR", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
				or.AddSeverityTextFilter("WARN", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
			})
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	orFilter, ok := request.Filters[0].Value.(*logsv1.Where_Or)
	if !ok {
		t.Fatal("Filter should be an Or filter")
	}

	if len(orFilter.Or.Filters) != 2 {
		t.Fatalf("Expected 2 filters in Or, got %d", len(orFilter.Or.Filters))
	}
}

func TestLogQueryRequestBuilder_Where_InstrumentationScope_AddNameFilter(t *testing.T) {
	request := NewLogQueryRequestBuilder().
		Where(func(f *WhereLogFilterConfigurator) {
			f.InstrumentationScope.AddNameFilter("my-logger", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	scope, ok := request.Filters[0].Value.(*logsv1.Where_InstrumentationScope)
	if !ok {
		t.Fatal("Filter should be an InstrumentationScope filter")
	}

	name, ok := scope.InstrumentationScope.Value.(*commonv1.InstrumentationScopeFilter_Name)
	if !ok {
		t.Fatal("InstrumentationScope filter should be Name")
	}

	if *name.Name.Compare != "my-logger" {
		t.Errorf("Expected 'my-logger', got '%s'", *name.Name.Compare)
	}
}

func TestLogQueryRequestBuilder_Chaining(t *testing.T) {
	builder := NewLogQueryRequestBuilder()

	result := builder.
		TakeFirst().
		TakeExact(10).
		TakeAll().
		Wait(time.Second).
		Where(func(f *WhereLogFilterConfigurator) {})

	if result != builder {
		t.Error("Methods should return the same builder instance for chaining")
	}
}
