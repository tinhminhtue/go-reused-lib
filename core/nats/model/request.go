package natsmodel

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func Request() {
	// Use the env variable if running in the container, otherwise use the default.
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	// Note: both request and response use leafnode, the remote note only for cluster load balancing.
	// url = "127.0.0.1:4111" // DEV: hardcoded to use leaf node;
	// fmt.Println("Nats URL: ", url)

	// Create an unauthenticated connection to NATS.
	nc, err := nats.Connect(url, nats.Token("s3cr3t"))

	// Now we can use the built-in `Request` method to do the service request.
	// We simply pass a nil body since that is being used right now. In addition,
	// we need to specify a timeout since with a request we are _waiting_ for the
	// reply and we likely don't want to wait forever.

	if nc != nil {
		fmt.Println("Client connected nats status: ", nc.IsConnected())
	} else {
		fmt.Println(err)
	}

	defer nc.Close()

	// encode bytes here
	// exampleBytes, err := json.Marshal(example)
	// if err != nil {
	// 	print(err)
	// 	return
	// }
	rep, _ := nc.Request("greet.joe", nil, time.Second)
	if rep != nil {
		fmt.Println(string(rep.Data))
	}

	rep, _ = nc.Request("greet.sue", nil, time.Second)
	if rep != nil {
		fmt.Println(string(rep.Data))
	}
}
