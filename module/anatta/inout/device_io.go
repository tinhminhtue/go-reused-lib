package inout

// / You can change the field inside structs. It got default value when older version is used.

type (
	DeviceIO struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		LocationId   int    `json:"location_id"`
		RegisterDate string `json:"register_date"`
		ActiveStatus string `json:"active_status"`
	}

	// For simple case, entity of activity and workflow can be the same
	CreateDeviceActInput struct {
		Entity DeviceIO `json:"entity"`
		Params Params   `json:"params"`
	}

	// CreateDeviceActOutput contains the result for this sample
	CreateDeviceActOutput struct {
		Data IdData `json:"data"`
		Err  string `json:"err"`
	}

	CreateDeviceFlowInput struct {
		CreateDeviceActInput
	}

	// For simple case, output data of activity and workflow can be the same
	CreateDeviceFlowOutput struct {
		CreateDeviceActOutput
	}
)

type (
	ListDeviceActInput struct {
		Entity DeviceIO `json:"entity"`
		Params Params   `json:"params"`
	}

	// ListDeviceActOutput contains the result for this sample
	ListDeviceActOutput struct {
		Data []DeviceIO `json:"data"`
		Err  string     `json:"err"`
	}

	// For simple case, output data of activity and workflow can be the same
	ListDeviceFlowInput struct {
		ListDeviceActInput
	}

	ListDeviceFlowOutput struct {
		ListDeviceActOutput
	}
)

// Implement IReply interface
func (output *CreateDeviceFlowOutput) GetErr() string {
	return output.Err
}

// Implement IReply interface
func (output *ListDeviceFlowOutput) GetErr() string {
	return output.Err
}

func (flowInput *CreateDeviceFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return flowInput.Entity.CheckValid()
	// return nil
}

func (input *DeviceIO) CheckValid() error {
	return nil
}

// CreateDeviceFlowInput HasAuthenticated interface
func (input *CreateDeviceFlowInput) BuildAuthUri() string {
	// this is sub uri parse from entity, example: device/id_1234 with action in header is "edit"
	// for create, it should be empty because it's new entity
	// for edit or delete, it should be "id_1234". "device" is the path of this api
	return ""
}
