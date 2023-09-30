package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	Fragment struct {
		NamespaceId uint64 `json:"namespaceId"`
		DepotId     uint64 `json:"depotId"`
		TimelineId  uint64 `json:"timelineId"`
		MomentId    uint64 `json:"momentId"`
		BaseModel
		Begin    string     `json:"begin"`
		End      string     `json:"end"`
		Children []Fragment `json:"children"`
	}

	CreateFragmentFlowInput struct {
		Entity Fragment `json:"entity"`
	}

	CreateFragmentActInput struct {
		Entity Fragment `json:"entity"`
	}
	// CreateFragmentOutput contains the result for this sample
	CreateFragmentOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateFragmentFlowOutput struct {
		Output CreateFragmentOutput `json:"output"`
	}
)

func (flowInput *CreateFragmentFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *Fragment) CheckValid() error {
	// check name space
	if input.NamespaceId == 0 {
		return errors.New("namespaceId is empty")
	}
	return nil
}

// Typical APIs request and response
type (
	ListFragmentRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
