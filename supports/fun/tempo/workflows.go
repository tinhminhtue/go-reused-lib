package tempo

import (
	"errors"

	"github.com/tinhminhtue/go-reused-lib/supports/idgen"
	idgenwire "github.com/tinhminhtue/go-reused-lib/supports/idgen/idgen_wire"
	"go.temporal.io/sdk/workflow"
)

func SideEffectNewUUID(ctx workflow.Context) (uint64, error) {
	sideEffectID := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
		id, er := idgenwire.IdGen().Generate()
		errString := ""
		if er != nil {
			errString = er.Error()
		}
		return idgen.IdSideEffect{
			Id:  id,
			Err: errString,
		}
	})
	var ids idgen.IdSideEffect
	err := sideEffectID.Get(&ids)
	if err != nil {
		return 0,
			err
	}
	if ids.Err != "" {
		return 0,
			errors.New(ids.Err)
	}
	return ids.Id, nil
}
