package idgen

import (
	"sync"

	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
	"github.com/spf13/viper"
	"github.com/tinhminhtue/go-reused-lib/core/defined"
)

// Singleton IdGenImpSonyflake
var IdGenImpSonyflakeSingleton *IdGenImpSonyflake
var once sync.Once

// Implentation of IdGen interface by sonyflake
type IdGenImpSonyflake struct {
	Sf *sonyflake.Sonyflake
}

// local machine id func
var machineIDFn = func() (uint16, error) {
	return 10, nil
}

// Provider singleton of IdGenImpSonyflake
func ProvideIdGenImpSonyflake() *IdGenImpSonyflake {
	once.Do(func() {
		var st sonyflake.Settings
		// Check for local machine id
		if viper.GetString(defined.USE_LOCAL_MACHINE_ID) == "true" {
			st.MachineID = machineIDFn
		} else {
			st.MachineID = awsutil.AmazonEC2MachineID
		}

		sf := sonyflake.NewSonyflake(st)
		IdGenImpSonyflakeSingleton = &IdGenImpSonyflake{
			Sf: sf,
		}
	})
	return IdGenImpSonyflakeSingleton
}

// Aws get machine id
func GetMachineID() (uint16, error) {
	return awsutil.AmazonEC2MachineID()
}

// generate id
func (i *IdGenImpSonyflake) Generate() (id uint64, err error) {
	// gen id using sony flakes
	id, err = i.Sf.NextID()
	return
}

// generate id for namespace
func (i *IdGenImpSonyflake) GenerateForNs(namespace uint64) (id uint64, err error) {
	// gen id using sony flakes
	id, err = i.Sf.NextID()
	return
}
