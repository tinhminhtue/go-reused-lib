package gclient

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tinhminhtue/go-reused-lib/core/ct"
	"github.com/tinhminhtue/go-reused-lib/core/logt"
	proxy "github.com/tinhminhtue/go-reused-lib/core/nats/proto"
	"github.com/tinhminhtue/go-reused-lib/core/otelc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const defaultTimeOutLimit = 5 // second

var conn *grpc.ClientConn
var grpcServerAddr string
var client proxy.ProxyLocalClient

// Call once when start, should exit program if return nil or retry..
func InitExternalClient(addr string) proxy.ProxyLocalClient {
	// Open tracing
	grpcServerAddr = addr
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	conn = con
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil
	}
	client = proxy.NewProxyLocalClient(conn)
	return client
}

func CloseConn() {
	if conn != nil {
		conn.Close()
	}
}

func getClient() proxy.ProxyLocalClient {
	if conn == nil {
		log.Errorln("Grpc connection is nil, retrying to create connection")

		newCon := InitExternalClient(grpcServerAddr)
		if newCon == nil {
			log.Errorln("Retry but fail to connect: ", grpcServerAddr)
			return nil
		} else {
			c := proxy.NewProxyLocalClient(conn)
			return c
		}
	} else {
		// TODO: check state and reconnect

		if conn.GetState() >= connectivity.TransientFailure {
			conn = nil
			client = nil
			log.Warningln("Retrying because Grpc connection state failed: ", conn.GetState())
			return getClient()
		}
		if client == nil {
			client = proxy.NewProxyLocalClient(conn)
		}
		return client

	}

}

func ExternalRequestBytes(ctx context.Context, uri string, bytes []byte) ([]byte, error) {
	c := getClient()
	if c == nil {
		return nil, fmt.Errorf("proxy client is nil")
	}
	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(ctx, defaultTimeOutLimit*time.Second)
	// defer cancel()
	ver := "v1"
	ctxVerVal := ctx.Value(ct.CtxVerKey)
	if ctxVerVal != nil {
		ver = ctxVerVal.(string)
	}
	md := metadata.New(map[string]string{"x-subject": uri, "x-version": ver})

	// Mark child force otel if this key exist

	var span trace.Span
	ctx, span = otel.Tracer(viper.GetString("service.name")).Start(ctx, "ExternalRequestBytes")
	logt.Info(span, "TraceID: ", span.SpanContext().TraceID().String(), "; SpanID: ", span.SpanContext().SpanID().String())
	isForced := false
	if forceSampleMeta, ok := ctx.Value(otelc.ForceSampleKey).(otelc.ForceSampleMeta); ok {
		isForced = true
		span.SetAttributes(attribute.String("previousErr", forceSampleMeta.PreviousErr))
	}
	span.SetAttributes(attribute.Bool("isForced", isForced))
	// add event (as log)
	logt.Info(span, "Acquiring lock")

	// API for status Error
	// span.SetStatus(codes.Error, "operationThatCouldFail failed")
	// RecordError
	// span.RecordError(err)
	defer span.End()
	// send force tracing to child service (manual isn't need when using auto instrument lib)
	// md.Set("otelc", span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String())

	ctx = metadata.NewOutgoingContext(ctx, md)

	r, err := c.ProxyNats(ctx, &proxy.ProxyRequest{Data: bytes})
	if err != nil {
		return nil, err
	}
	return r.GetData(), nil
}

func SendExternalRequest[T any](ctx context.Context, subject string, model any) (resp T, err error) {
	bytes, err := json.Marshal(model)
	if err != nil {
		return resp, err
	}
	data, err := ExternalRequestBytes(ctx, subject, bytes)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func SendExternalRequestBytes(ctx context.Context, subject string, bytes []byte) ([]byte, error) {
	return ExternalRequestBytes(ctx, subject, bytes)
}
