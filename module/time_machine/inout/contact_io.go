package inout

// / You can change the field inside structs. It got default value when older version is used.

type (
	ContactIO struct {
		UserId       uint64         `json:"userId"`
		UserIdTarget uint64         `json:"userIdTarget"`
		Status       int8           `json:"status"`
		IntroRequest map[string]any `json:"introRequest"`
		IntroReply   map[string]any `json:"introReply"`
	}

	// For simple case, entity of activity and workflow can be the same
	CreateContactActInput struct {
		Entity ContactIO `json:"entity"`
		Params Params    `json:"params"`
	}

	CreateContactFlowInput struct {
		Entity ContactIO `json:"entity"`
		Params Params    `json:"params"`
	}

	// CreateContactActOutput contains the result for this sample
	CreateContactActOutput struct {
		Data IdData `json:"data"`
		Err  string `json:"err"`
	}

	// For simple case, output data of activity and workflow can be the same
	CreateContactFlowOutput struct {
		Data IdData `json:"data"`
		Err  string `json:"err"`
	}
)

func (flowInput *CreateContactFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *ContactIO) CheckValid() error {
	return nil
}

// CreateContactFlowInput HasAuthenticated interface
func (input *CreateContactFlowInput) BuildAuthUri() string {
	// this is sub uri parse from entity, example: contact/1234 with action is "edit"
	// for create, it should be empty
	return ""
}

// Typical APIs request and response
type (
	ListContactRequest struct {
		Filter struct {
			Tenant   string `json:"tenant"`
			DataType string `json:"dataType"`
		} `json:"filter"`
	}

	CreateContactRequest struct {
		Entity ContactIO `json:"entity"`
	}

	CreateContactResponse struct {
		Data struct {
			ID uint64 `json:"id"`
		} `json:"data"`
		Err string `json:"err"`
	}
)
