// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package idgenwire

import (
	"github.com/tinhminhtue/go-reused-lib/support/idgen"
)

// Injectors from wire.go:

func IdGen() idgen.IdGen {
	idGenImpSonyflake := idgen.ProvideIdGenImpSonyflake()
	return idGenImpSonyflake
}
