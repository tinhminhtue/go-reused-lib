package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	Moment struct {
		NamespaceId uint64 `json:"namespaceId"`
		DepotId     uint64 `json:"depotId"`
		TimelineId  uint64 `json:"timelineId"`
		BaseModel
		Begin    string   `json:"begin"`
		End      string   `json:"end"`
		Children []Moment `json:"children"`
	}

	CreateMomentFlowInput struct {
		Entity Moment `json:"entity"`
	}

	CreateMomentActInput struct {
		Entity Moment `json:"entity"`
	}
	// CreateMomentOutput contains the result for this sample
	CreateMomentOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateMomentFlowOutput struct {
		Output CreateMomentOutput `json:"output"`
	}
)

func (flowInput *CreateMomentFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *Moment) CheckValid() error {
	// check name space
	if input.NamespaceId == 0 {
		return errors.New("namespaceId is empty")
	}
	return nil
}

// Typical APIs request and response
type (
	ListMomentRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
