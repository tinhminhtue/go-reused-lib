package otelc

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	traceapi "go.opentelemetry.io/otel/trace"
)

type key int

const ForceSampleKey key = iota

type sampler struct {
	delegate trace.Sampler
}

func NewCustomSampler(delegate trace.Sampler) trace.Sampler {
	return &sampler{delegate: delegate}
}

func (s *sampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	if _, ok := p.ParentContext.Value(ForceSampleKey).(struct{}); ok {
		return trace.SamplingResult{
			Decision:   trace.RecordAndSample,
			Attributes: []attribute.KeyValue{},
			Tracestate: traceapi.SpanContextFromContext(p.ParentContext).TraceState(),
		}
	}
	return s.delegate.ShouldSample(p)
}

func (s *sampler) Description() string {
	return fmt.Sprintf("my-sampler{%s}", s.delegate.Description())
}

// func main() {
// 	s := New(trace.ParentBased(trace.TraceIDRatioBased(0.5)))
// 	tracer := trace.NewTracerProvider(trace.WithSampler(s)).Tracer("my-app")

// 	ctx := context.WithValue(context.Background(), ForceSampleKey, struct{}{})
// 	_, span := tracer.Start(ctx, "span name")
// 	fmt.Println(span.SpanContext().IsSampled())
// }
