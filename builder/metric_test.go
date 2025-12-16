package builder

import (
	"testing"
	"time"

	commonv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/common/v1"
	metricsv1 "github.com/OddDotNet/OddDotGo/gen/odddotproto/proto/metrics/v1"
	otlpmetricsv1 "github.com/OddDotNet/OddDotGo/gen/opentelemetry/proto/metrics/v1"
)

func TestMetricQueryRequestBuilder_Defaults(t *testing.T) {
	builder := NewMetricQueryRequestBuilder()
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

func TestMetricQueryRequestBuilder_TakeFirst(t *testing.T) {
	request := NewMetricQueryRequestBuilder().TakeFirst().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeFirst); !ok {
		t.Error("Take should be TakeFirst")
	}
}

func TestMetricQueryRequestBuilder_TakeExact(t *testing.T) {
	request := NewMetricQueryRequestBuilder().TakeExact(5).Build()

	takeExact, ok := request.Take.Value.(*commonv1.Take_TakeExact)
	if !ok {
		t.Fatal("Take should be TakeExact")
	}
	if takeExact.TakeExact.Count != 5 {
		t.Errorf("TakeExact count should be 5, got %d", takeExact.TakeExact.Count)
	}
}

func TestMetricQueryRequestBuilder_TakeAll(t *testing.T) {
	request := NewMetricQueryRequestBuilder().TakeAll().Build()

	if _, ok := request.Take.Value.(*commonv1.Take_TakeAll); !ok {
		t.Error("Take should be TakeAll")
	}
}

func TestMetricQueryRequestBuilder_Wait(t *testing.T) {
	request := NewMetricQueryRequestBuilder().Wait(5 * time.Second).Build()

	if request.Duration.Milliseconds != 5000 {
		t.Errorf("Duration should be 5000ms, got %d", request.Duration.Milliseconds)
	}
}

func TestMetricQueryRequestBuilder_Where_AddNameFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.AddNameFilter("http.server.duration", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	name, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Name)
	if !ok {
		t.Fatal("Property filter should be Name")
	}

	if *name.Name.Compare != "http.server.duration" {
		t.Errorf("Expected name 'http.server.duration', got '%s'", *name.Name.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_AddDescriptionFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.AddDescriptionFilter("request duration", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_CONTAINS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	desc, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Description)
	if !ok {
		t.Fatal("Property filter should be Description")
	}

	if *desc.Description.Compare != "request duration" {
		t.Errorf("Expected 'request duration', got '%s'", *desc.Description.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_AddUnitFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.AddUnitFilter("ms", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	unit, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Unit)
	if !ok {
		t.Fatal("Property filter should be Unit")
	}

	if *unit.Unit.Compare != "ms" {
		t.Errorf("Expected 'ms', got '%s'", *unit.Unit.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_Gauge_AddDataPointDoubleValueFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.Gauge.AddDataPointDoubleValueFilter(100.5, commonv1.NumberCompareAsType_NUMBER_COMPARE_AS_TYPE_GREATER_THAN)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	gauge, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Gauge)
	if !ok {
		t.Fatal("Property filter should be Gauge")
	}

	dataPoint, ok := gauge.Gauge.Value.(*metricsv1.GaugeFilter_DataPoint)
	if !ok {
		t.Fatal("Gauge filter should be DataPoint")
	}

	value, ok := dataPoint.DataPoint.Value.(*metricsv1.NumberDataPointFilter_ValueAsDouble)
	if !ok {
		t.Fatal("DataPoint filter should be ValueAsDouble")
	}

	if *value.ValueAsDouble.Compare != 100.5 {
		t.Errorf("Expected 100.5, got %f", *value.ValueAsDouble.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_Sum_AddAggregationTemporalityFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.Sum.AddAggregationTemporalityFilter(
				otlpmetricsv1.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE,
				commonv1.EnumCompareAsType_ENUM_COMPARE_AS_TYPE_EQUALS,
			)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	sum, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Sum)
	if !ok {
		t.Fatal("Property filter should be Sum")
	}

	agg, ok := sum.Sum.Value.(*metricsv1.SumFilter_AggregationTemporality)
	if !ok {
		t.Fatal("Sum filter should be AggregationTemporality")
	}

	if agg.AggregationTemporality.Compare != otlpmetricsv1.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE {
		t.Errorf("Expected CUMULATIVE, got %v", agg.AggregationTemporality.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_Sum_AddIsMonotonicFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.Sum.AddIsMonotonicFilter(true, commonv1.BoolCompareAsType_BOOL_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	sum, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Sum)
	if !ok {
		t.Fatal("Property filter should be Sum")
	}

	monotonic, ok := sum.Sum.Value.(*metricsv1.SumFilter_IsMonotonic)
	if !ok {
		t.Fatal("Sum filter should be IsMonotonic")
	}

	if *monotonic.IsMonotonic.Compare != true {
		t.Error("Expected true")
	}
}

func TestMetricQueryRequestBuilder_Where_Histogram_AddDataPointCountFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.Histogram.AddDataPointCountFilter(100, commonv1.NumberCompareAsType_NUMBER_COMPARE_AS_TYPE_GREATER_THAN_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	prop, ok := request.Filters[0].Value.(*metricsv1.Where_Property)
	if !ok {
		t.Fatal("Filter should be a Property filter")
	}

	histogram, ok := prop.Property.Value.(*metricsv1.PropertyFilter_Histogram)
	if !ok {
		t.Fatal("Property filter should be Histogram")
	}

	dataPoint, ok := histogram.Histogram.Value.(*metricsv1.HistogramFilter_DataPoint)
	if !ok {
		t.Fatal("Histogram filter should be DataPoint")
	}

	count, ok := dataPoint.DataPoint.Value.(*metricsv1.HistogramDataPointFilter_Count)
	if !ok {
		t.Fatal("DataPoint filter should be Count")
	}

	if *count.Count.Compare != 100 {
		t.Errorf("Expected 100, got %d", *count.Count.Compare)
	}
}

func TestMetricQueryRequestBuilder_Where_AddOrFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.AddOrFilter(func(or *WhereMetricFilterConfigurator) {
				or.AddNameFilter("metric-a", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
				or.AddNameFilter("metric-b", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
			})
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	orFilter, ok := request.Filters[0].Value.(*metricsv1.Where_Or)
	if !ok {
		t.Fatal("Filter should be an Or filter")
	}

	if len(orFilter.Or.Filters) != 2 {
		t.Fatalf("Expected 2 filters in Or, got %d", len(orFilter.Or.Filters))
	}
}

func TestMetricQueryRequestBuilder_Where_InstrumentationScope_AddNameFilter(t *testing.T) {
	request := NewMetricQueryRequestBuilder().
		Where(func(f *WhereMetricFilterConfigurator) {
			f.InstrumentationScope.AddNameFilter("otel-library", commonv1.StringCompareAsType_STRING_COMPARE_AS_TYPE_EQUALS)
		}).
		Build()

	if len(request.Filters) != 1 {
		t.Fatalf("Expected 1 filter, got %d", len(request.Filters))
	}

	scope, ok := request.Filters[0].Value.(*metricsv1.Where_InstrumentationScope)
	if !ok {
		t.Fatal("Filter should be an InstrumentationScope filter")
	}

	name, ok := scope.InstrumentationScope.Value.(*commonv1.InstrumentationScopeFilter_Name)
	if !ok {
		t.Fatal("InstrumentationScope filter should be Name")
	}

	if *name.Name.Compare != "otel-library" {
		t.Errorf("Expected 'otel-library', got '%s'", *name.Name.Compare)
	}
}

func TestMetricQueryRequestBuilder_Chaining(t *testing.T) {
	builder := NewMetricQueryRequestBuilder()

	result := builder.
		TakeFirst().
		TakeExact(10).
		TakeAll().
		Wait(time.Second).
		Where(func(f *WhereMetricFilterConfigurator) {})

	if result != builder {
		t.Error("Methods should return the same builder instance for chaining")
	}
}
