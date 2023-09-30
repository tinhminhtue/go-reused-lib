package client

import "yedda.io/yedda_temporal_libs/notify/inout"

//
// Proxy Activities is client API that call GRPC, HTTP to your Module server.
//
// Notify_* prefix mark this is an external activity for other module

func Notify_SendMessageActivity(input inout.StoplightActionReactInput) (output inout.StoplightActionReactOutput, err error) {
	// Write your HTTP API request here

	return
}
