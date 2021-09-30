package tracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
// If no server option detected, it will become the root span
func StartSpanFromRequest(operationName string, r *http.Request) opentracing.Span {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := Extract(tracer, r)
	return tracer.StartSpan(operationName, ext.RPCServerOption(spanCtx))
}
