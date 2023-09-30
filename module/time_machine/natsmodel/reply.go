package natsmodel

import (
	"context"
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

func Reply(ctx context.Context) {
	// Use the env variable if running in the container, otherwise use the default.
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	url = "127.0.0.1:4111" // DEV: hardcoded to use leaf node
	fmt.Println("Nats URL: ", url)

	// Create an unauthenticated connection to NATS.
	nc, _ := nats.Connect(url)
	defer nc.Drain()

	// In addition to vanilla publish-request, NATS supports request-reply
	// interactions as well. Under the covers, this is just an optimized
	// pair of publish-subscribe operations.
	// The _request handler_ is just a subscription that _responds_ to a message
	// sent to it. This kind of subscription is called a _service_.
	// For this example, we can use the built-in asynchronous
	// subscription in the Go SDK.
	sub, _ := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		// Parse out the second token in the subject (everything after greet.)
		// and use it as part of the response message.
		name := msg.Subject[6:]
		// decode byte here

		msg.Respond([]byte("hello, " + name))
	})

	<-ctx.Done()
	sub.Unsubscribe()
}
