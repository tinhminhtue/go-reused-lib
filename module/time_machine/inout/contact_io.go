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

// Implement IReply interface
func (output *CreateContactFlowOutput) GetErr() string {
	return output.Err
}

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
	// this is sub uri parse from entity, example: contact/id_1234 with action in header is "edit"
	// for create, it should be empty because it's new entity
	// for edit or delete, it should be "id_1234". "contact" is the path of this api
	return ""
}

// Deprecated: use CreateContactFlowInput instead, the request response moved to flow input output named.
type (
//	ListContactRequest struct {
//		Filter struct {
//			Tenant   string `json:"tenant"`
//			DataType string `json:"dataType"`
//		} `json:"filter"`
//	}
)
