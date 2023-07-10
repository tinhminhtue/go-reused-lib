package gclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ory/viper"
	log "github.com/sirupsen/logrus"
	proxy "github.com/tinhminhtue/go-reused-lib/nats/proto"
	"github.com/tinhminhtue/go-reused-lib/otelc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
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
	ctx, cancel := context.WithTimeout(ctx, defaultTimeOutLimit*time.Second)
	defer cancel()
	md := metadata.New(map[string]string{"x-subject": uri})

	// Mark child force otel if this key exist
	if _, ok := ctx.Value(otelc.ForceSampleKey).(struct{}); ok {
		md.Set("otelc", "force")
	}

	ctx = metadata.NewOutgoingContext(ctx, md)

	log.Info("Test open telemetry")
	ctx, span := otel.Tracer(viper.GetString("service.name")).Start(ctx, "ExternalRequestBytes")
	defer span.End()

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
