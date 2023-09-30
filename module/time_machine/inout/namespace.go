package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	Namespace struct {
		BaseModel
		Tenant   string `json:"tenant"`
		DataType string `json:"dataType"`
	}

	CreateNamespaceActInput struct {
		Entity Namespace `json:"entity"`
	}

	CreateNamespaceFlowInput struct {
		Entity Namespace `json:"entity"`
	}

	// CreateNamespaceOutput contains the result for this sample
	CreateNamespaceOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateNamespaceFlowOutput struct {
		Output CreateNamespaceOutput `json:"output"`
	}
)

func (flowInput *CreateNamespaceFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *Namespace) CheckValid() error {
	if input.Tenant == "" {
		return errors.New("meta.tenant is empty")
	}
	if input.DataType == "" {
		return errors.New("meta.dataType is empty")
	}
	return nil
}

// Typical APIs request and response
type (
	ListNamespaceRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
