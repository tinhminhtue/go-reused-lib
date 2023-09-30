package rpc

import (
	"github.com/pborman/uuid"
	"go.temporal.io/sdk/workflow"
	defined "yedda.io/yedda_temporal_libs/notify/defined"
	"yedda.io/yedda_temporal_libs/notify/inout"
)

//
// Rpc package is optional
// This package for Child workflow starter template. It help your client call your workflow precisely
// https://docs.temporal.io/go/spawn-a-child-workflow-execution
//

func ExampleChildWorkflow(ctx workflow.Context, request inout.WriteMessageRequest) (response inout.WriteMessageResponse, err error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: defined.QueueDefault + "_" + uuid.New(),
		TaskQueue:  defined.QueueDefault,
		// ...
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	err = workflow.ExecuteChildWorkflow(ctx, "ExampleChildWorkflow", request).Get(ctx, &response)
	if err != nil {
		logger.Error("ExampleChildWorkflow child execution failure.", "Error", err)
		return
	}

	logger.Info("ExampleChildWorkflow child execution completed.")
	return
}
