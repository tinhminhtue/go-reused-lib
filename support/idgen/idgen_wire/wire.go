//go:build wireinject
// +build wireinject

package idgenwire

import (
	"github.com/google/wire"
	"github.com/tinhminhtue/go-reused-lib/support/idgen"
)

func IdGen() idgen.IdGen {
	wire.Build(IdGenSet)
	return nil
}
