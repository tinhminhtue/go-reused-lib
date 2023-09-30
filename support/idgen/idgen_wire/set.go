package idgenwire

import (
	"github.com/google/wire"
	"github.com/tinhminhtue/go-reused-lib/support/idgen"
)

var IdGenSet = wire.NewSet(
	// wire.Value()
	// conf.ProvideConfig,
	// drivers.ProvideGormDriver,
	// drivers.DriverSet,
	idgen.ProvideIdGenImpSonyflake,
	wire.Bind(new(idgen.IdGen), new(*idgen.IdGenImpSonyflake)),
)
