package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	Depot struct {
		NamespaceId uint64 `json:"namespaceId"`
		BaseModel
		Name     string  `json:"name"`
		Children []Depot `json:"children"`
	}

	CreateDepotActInput struct {
		Entity Depot `json:"entity"`
	}

	CreateDepotFlowInput struct {
		Entity Depot `json:"entity"`
	}

	// CreateDepotOutput contains the result for this sample
	CreateDepotOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateDepotFlowOutput struct {
		Output CreateDepotOutput `json:"output"`
	}
)

func (flowInput *CreateDepotFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *Depot) CheckValid() error {
	// check name space
	if input.NamespaceId == 0 {
		return errors.New("namespaceId is empty")
	}
	return nil
}

// Typical APIs request and response
type (
	ListDepotRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
