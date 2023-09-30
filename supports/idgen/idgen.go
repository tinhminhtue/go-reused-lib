package idgen

// generate id interface
type IdGen interface {
	// generate id
	Generate() (id uint64, err error)
	GenerateForNs(namespaceID uint64) (id uint64, err error)
}

type IdSideEffect struct {
	Id  uint64
	Err string
}
