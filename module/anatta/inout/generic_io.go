package inout

type BaseModel struct {
	Cls  string                 `json:"cls"` // only one class now: 'Namespace'
	ID   uint64                 `json:"id"`
	Ver  string                 `json:"ver"` // header got this ver duplicated for Auth
	Meta map[string]interface{} `json:"meta"`
}

type GenericResponse[E any] struct {
	Data E      `json:"data"`
	Err  string `json:"err"`
}

type IdEntity = uint64
type StrEntity = string

type IdData struct {
	ID IdEntity `json:"id"`
}

type HasAuthenticated interface {
	BuildAuthUri() string
}

type IReply interface {
	GetErr() string
}
