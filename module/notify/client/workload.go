package client

import "yedda.io/yedda_temporal_libs/notify/inout"

//
// Workload Activities is shared for other Module to use directly in their worker.
//
// Notify_* prefix mark this is an external activity for other module

func Notify_EncryptEndToEndActivity(input inout.StoplightActionReactInput) (output inout.StoplightActionReactOutput) {
	// Write your CPU bound encrypt here

	return output
}
