package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	Timeline struct {
		NamespaceId uint64 `json:"namespaceId"`
		DepotId     uint64 `json:"depotId"`
		BaseModel
		Name     string     `json:"name"`
		Priority int        `json:"priority"`
		Children []Timeline `json:"children"`
	}

	CreateTimelineFlowInput struct {
		Entity Timeline `json:"entity"`
	}

	CreateTimelineActInput struct {
		Entity Timeline `json:"entity"`
	}

	// CreateTimelineOutput contains the result for this sample
	CreateTimelineOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateTimelineFlowOutput struct {
		Output CreateTimelineOutput `json:"output"`
	}
)

func (flowInput *CreateTimelineFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *Timeline) CheckValid() error {
	// check name space
	if input.NamespaceId == 0 {
		return errors.New("namespaceId is empty")
	}
	return nil
}

// Typical APIs request and response
type (
	ListTimelineRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
