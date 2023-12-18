package inout

import "errors"

// / You can change the field inside structs. It got default value when older version is used.

type (
	ContactIO struct {
		UserId       int64          `json:"userId"`
		UserIdTarget int64          `json:"userIdTarget"`
		Status       int8           `json:"status"`
		IntroRequest map[string]any `json:"introRequest"`
		IntroReply   map[string]any `json:"introReply"`
	}

	CreateContactActInput struct {
		Entity ContactIO `json:"entity"`
	}

	CreateContactFlowInput struct {
		Entity ContactIO `json:"entity"`
	}

	// CreateContactOutput contains the result for this sample
	CreateContactOutput struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}

	CreateContactFlowOutput struct {
		Output CreateContactOutput `json:"output"`
	}
)

func (flowInput *CreateContactFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *ContactIO) CheckValid() error {
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
	ListContactRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}
)
