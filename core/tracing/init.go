package tracing

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Init returns an instance of Jaeger Tracer.
// func Init(service string) (opentracing.Tracer, io.Closer) {
// 	cfg, err := config.FromEnv()
// 	if err != nil {
// 		panic(fmt.Sprintf("ERROR: failed to read config from env vars: %v\n", err))
// 	}
// 	cfg.ServiceName = service
// 	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
// 	if err != nil {
// 		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
// 	}
// 	return tracer, closer
// }

// NewFileExporter returns a console exporter.
func NewFileExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func NewGrpcExporter(ctx context.Context) (trace.SpanProcessor, error) {
	//HTTP
	// client := otlptracehttp.NewClient()
	// exporter, err := otlptrace.New(ctx, client)
	// if err != nil {
	// 	return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
	// }

	/// GRPC

	// If the OpenTelemetry Collector is running on a local cluster (minikube or
	// microk8s), it should be accessible through the NodePort service at the
	// `localhost:30080` endpoint. Otherwise, replace `localhost` with the
	// endpoint of your cluster. If you run the app inside k8s, then you can
	// probably connect directly to the service through dns.
	conn, err := grpc.DialContext(ctx, "localhost:4317",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}
	bsp := trace.NewBatchSpanProcessor(traceExporter)
	return bsp, nil
}
