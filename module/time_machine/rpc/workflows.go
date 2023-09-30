package rpc

import (
	"github.com/pborman/uuid"
	defined "github.com/tinhminhtue/go-reused-lib/module/time_machine/defined"
	"github.com/tinhminhtue/go-reused-lib/module/time_machine/inout"
	"go.temporal.io/sdk/workflow"
)

//
// Rpc package is optional
// This package for Child workflow starter template. It help your client call your workflow precisely
// https://docs.temporal.io/go/spawn-a-child-workflow-execution
//

func SendChatChildWorkflow_v1(ctx workflow.Context, request inout.CreateNamespaceFlowInput) (response inout.CreateNamespaceFlowOutput, err error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{
		WorkflowID: defined.QueueDefault + "_" + uuid.New(),
		TaskQueue:  defined.QueueDefault,
		// ...
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	err = workflow.ExecuteChildWorkflow(ctx, "SendChatWorkflow_v1", request).Get(ctx, &response)
	if err != nil {
		logger.Error("SendChatWorkflow_v1 child execution failure.", "Error", err)
		return
	}

	logger.Info("SendChatWorkflow_v1 child execution completed.")
	return
}
