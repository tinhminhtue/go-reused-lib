package inout

import (
	"errors"
	"time"
)

// / You can change the field inside structs. It got default value when older version is used.

type (
	StoplightActionReactInput struct {
		Notification string `json:"notification"`
		Action       Action `json:"action"`
	}

	Action struct {
		Actions []string `json:"actions"`
	}

	StoplightActionReactFlowInput struct {
		Input StoplightActionReactInput `json:"input"`
	}

	// StoplightActionReactOutput contains the result for this sample
	StoplightActionReactOutput struct {
		SendChat         bool   `json:"send_chat"`
		ConversationId   int    `json:"conversation_id"`
		NotificationPath string `json:"notification_path"`
		Content          string `json:"content"`
	}

	StoplightActionReactFlowOutput struct {
		Output StoplightActionReactOutput `json:"output"`
	}
)

func (input *StoplightActionReactFlowInput) CheckValid() error {
	// To check check flow input here, return true if valid
	return nil
}

func (input *StoplightActionReactInput) CheckValid() error {
	if input.Notification == "" {
		return errors.New("Missing notification")
	}
	if len(input.Action.Actions) == 0 {
		return errors.New("Missing actions")
	}
	return nil
}

// etl worker notify
type (
	EtlWorkerNotifyInput struct {
		IdCompany string     `json:"id_company"`
		IdModule  int        `json:"id_module"`
		Data      DataNotify `json:"data"`
		MetaBoard MetaBoard  `json:"meta_board"`
	}

	MetaBoard struct {
		ExpiredNotificationTemplateBody  string `json:"expired_notification_template_body"`
		ExpiredNotificationTemplateTitle string `json:"expired_notification_template_title"`
	}

	DataNotify struct {
		IdNotify    int        `json:"id_notify"`
		AreaType    int        `json:"area_type"`
		AlertType   int        `json:"alert_type"`
		IdStation   int        `json:"id_station"`
		DocumentId  string     `json:"document_id"`
		ActionType  string     `json:"action_type"`
		Priority    int        `json:"priority"`
		Author      int        `json:"author"`
		TimeExpired *time.Time `json:"time_expired"`
		Module      string     `json:"module"`
		AlertAction string     `json:"alert_action"`
		ActionData  string     `json:"action_data"`
	}

	EtlWorkerNotifyFlowInput struct {
		Input EtlWorkerNotifyInput `json:"input"`
	}

	EtlWorkerNotifyFlowOutput struct {
	}

	EtlWorkerNotifyFirebaseOutput struct {
		SendNotification        bool `json:"send_notification"`
		SendOverdueNotification bool `json:"send_overdue_notification"`
		EtlWorkerNotifySendNotificationInput
		SendChat bool `json:"send_chat"`
		EtlWorkerNotifySendChatInput
	}

	EtlWorkerNotifySendNotificationInput struct {
		Body        BodyNotification        `json:"body"`
		BodyOverdue BodyOverdueNotification `json:"body_overdue"`
		Headers     InfoHeaderNotification  `json:"headers"`
	}

	EtlWorkerNotifySendChatInput struct {
		ConversationId   int    `json:"conversation_id"`
		NotificationPath string `json:"notification_path"`
		Content          string `json:"content"`
	}

	Property struct {
		Priority    int        `json:"priority" firestore:"priority"`
		ActionType  string     `json:"action_type" firestore:"current_state"`
		IsExpired   bool       `json:"is_expired" firestore:"is_expired"`
		TimeExpired *time.Time `json:"time_expired" firestore:"time_expired,omitempty"`
	}

	ActionHistory struct {
		ActionType         string      `json:"action_type" firestore:"action_type"`
		SubActions         []SubAction `json:"sub_actions,omitempty" firestore:"sub_actions,omitempty"`
		Actions            []string    `json:"actions,omitempty" firestore:"actions,omitempty"`
		ActionMode         int         `json:"action_mode,omitempty" firestore:"action_mode,omitempty"`
		Author             string      `json:"author" firestore:"author"`
		Color              string      `json:"color" firestore:"color"`
		Comment            string      `json:"comment" firestore:"comment"`
		Label              string      `json:"label" firestore:"label"`
		Text               string      `json:"text" firestore:"text"`
		Attachment         string      `json:"attachment" firestore:"attachment,omitempty"`
		AttachmentMimetype string      `json:"attachment_mimetype" firestore:"attachment_mimetype,omitempty"`
		TimeCreated        time.Time   `json:"time_created" firestore:"time_created"`
	}

	BodyOverdueNotification struct {
		Apns         Apns         `json:"apns"`
		Data         DataOverdue  `json:"data"`
		Notification Notification `json:"notification"`
		Priority     string       `json:"priority"`
	}

	DataOverdue struct {
		IdModule    string `json:"id_module"`
		IdCompany   string `json:"id_company"`
		Module      string `json:"module"`
		Path        string `json:"path"`
		Author      string `json:"author"`
		ClickAction string `json:"click_action"`
	}

	InfoHeaderDataReportNotification struct {
		MessageKey  string   `json:"message_key"`
		Action      string   `json:"action"`
		Recipients  []string `json:"recipients"`
		ServiceFrom string   `json:"service_from"`
	}

	BodyDataReportNotification struct {
		Apns         Apns         `json:"apns"`
		Data         DataReport   `json:"data"`
		Notification Notification `json:"notification"`
		Priority     string       `json:"priority"`
	}

	DataReport struct {
		Action      string `json:"action"`
		Author      string `json:"author"`
		ClickAction string `json:"click_action"`
		En          string `json:"en"`
		IdCompany   string `json:"id_company"`
		Module      string `json:"module"`
		Path        string `json:"path"`
	}

	BodyNotification struct {
		Apns         Apns         `json:"apns"`
		Data         Data         `json:"data"`
		Notification Notification `json:"notification"`
		Priority     string       `json:"priority"`
	}

	Apns struct {
		Payload Payload `json:"payload"`
	}

	Payload struct {
		Aps Aps `json:"aps"`
	}

	Aps struct {
		MutableContent int `json:"mutable-content"`
	}

	Data struct {
		IdAction      string `json:"id_action"`
		IdModule      string `json:"id_module"`
		IdCompany     string `json:"id_company"`
		Module        string `json:"module"`
		Path          string `json:"path"`
		Author        string `json:"author"`
		ClickAction   string `json:"click_action"`
		NotiType      string `json:"noti_type"`
		ActionHistory string `json:"action_history"`
	}

	Notification struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	InfoHeaderNotification struct {
		Action      string   `json:"action"`
		Recipients  []string `json:"recipients"`
		ServiceFrom string   `json:"service_from"`
	}

	AlertAction struct {
		Id                    int                   `json:"id"`
		Properties            PropertyConfig        `json:"properties"`
		ActionHistoryTemplate ActionHistoryTemplate `json:"action_history_template"`
		InputProperties       []InputProperty       `json:"input_properties"`
		ActionType            string                `json:"action_type"`
	}

	PropertyConfig struct {
		Icon                 string               `json:"icon"`
		Name                 string               `json:"name"`
		NotificationTemplate NotificationTemplate `json:"notification_template"`
	}

	NotificationTemplate struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	ActionHistoryTemplate struct {
		Color string `json:"color"`
		Label string `json:"label"`
		Text  string `json:"text"`
	}

	InputProperty struct {
		AllowInput    bool           `json:"allow_input"`
		Display       string         `json:"display"`
		PropertyName  string         `json:"property_name"`
		PropertyType  string         `json:"property_type"`
		Options       []string       `json:"options"`
		ActionOptions []ActionOption `json:"action_options"`
	}

	ActionOption struct {
		Id           int    `json:"id"`
		PropertyName string `json:"property_name"`
		Caption      string `json:"caption"`
		Active       int    `json:"active"`
		IdSubAction  int    `json:"id_sub_action"`
		Meta         Meta   `json:"meta"`
	}

	Meta struct {
		NotificationTemplateBody  *string `json:"notification_template_body"`
		NotificationTemplateTitle string  `json:"notification_template_title"`
		Recipients                string  `json:"recipients"`
	}

	ActionBody struct {
		IdAction   int        `json:"id_action"`
		IdModule   int        `json:"id_module"`
		ActionData ActionData `json:"action_data"`
	}

	ActionData struct {
		Comment            string        `json:"comment"`
		Actions            []string      `json:"actions"`
		SubActions         []SubAction   `json:"sub_actions"`
		Attachment         string        `json:"attachment"`
		AttachmentMimetype string        `json:"attachment_mimetype"`
		ActionHistory      ActionHistory `json:"action_history"`
	}

	SubAction struct {
		Id          int    `json:"id" firestore:"id"`
		Caption     string `json:"caption" firestore:"caption"`
		IdSubAction int    `json:"id_sub_action" firestore:"id_sub_action"`
	}

	DataResponseUsers struct {
		Data DataUsers `json:"data"`
	}

	DataUsers struct {
		Users []User `json:"users"`
	}

	DataResponseIdUsers struct {
		Data DataIdUsers `json:"data"`
	}

	DataIdUsers struct {
		IdUsers []int `json:"id_users"`
	}

	User struct {
		IdUser   int    `json:"id_user"`
		LastName string `json:"last_name"`
		Name     string `json:"name"`
	}

	ResponseErr struct {
		Err string `json:"err"`
	}

	EtlWorkerNotifyOverdueFlowInput struct {
		Input EtlWorkerNotifyOverdueInput `json:"input"`
	}
	EtlWorkerNotifyOverdueInput struct {
		IdCompany string            `json:"id_company"`
		Data      DataNotifyOverdue `json:"data"`
	}
	DataNotifyOverdue struct {
		Error string `json:"error"`
	}
	EtlWorkerNotifyOverdueFlowOutput struct {
	}
	EtlWorkerNotifyOverdueOutput struct {
		SendChat bool `json:"send_chat"`
		EtlWorkerNotifySendChatInput
	}

	DataReportFlowInput struct {
		DataReportInput
	}
	DataReportInput struct {
		Report  Report `json:"report"`
		Related string `json:"related"`
	}
	NewReport struct {
		Name      string      `json:"name" firestore:"name"`
		DateFrom  string      `json:"date_from" firestore:"date_from"`
		DateTo    string      `json:"date_to" firestore:"date_to"`
		Data      interface{} `json:"data" firestore:"data"`
		Files     interface{} `json:"files" firestore:"files"`
		TimeBegin string      `json:"time_begin" firestore:"time_begin"`
		TimeEnd   string      `json:"time_end" firestore:"time_end"`
		IdCompany int         `json:"id_company" firestore:"id_company"`
	}
	Report struct {
		Name       string      `json:"name"`
		Datefrom   string      `json:"datefrom"`
		Dateto     string      `json:"dateto"`
		Data       interface{} `json:"data"`
		Files      interface{} `json:"files"`
		Hourfrom   string      `json:"hourfrom"`
		Hourto     string      `json:"hourto"`
		IdCustomer int         `json:"id_customer"`
	}
	InsertMessage struct {
		Txt         map[string]string `json:"txt" firestore:"txt"`
		Recipients  []string          `json:"recipients" firestore:"recipients"`
		Report      NewReport         `json:"report" firestore:"report"`
		IdCompany   int               `json:"id_company" firestore:"id_company"`
		TimeCreated time.Time         `json:"time_created" firestore:"time_created"`
	}
	DataReportOutput struct {
		SendNotification bool `json:"send_notification"`
		DataReportSendNotificationInput
		Path string `json:"path"`
	}
	DataReportFlowOutput struct {
		Path string `json:"path"`
	}

	DataReportSendNotificationInput struct {
		Body    BodyDataReportNotification       `json:"body"`
		Headers InfoHeaderDataReportNotification `json:"headers"`
	}
)

func (input *EtlWorkerNotifyFlowInput) CheckValid() error {
	return nil
}
func (input *EtlWorkerNotifyInput) CheckValid() error {
	return nil
}
func (input *EtlWorkerNotifyOverdueFlowInput) CheckValid() error {
	return nil
}
func (input *EtlWorkerNotifyOverdueInput) CheckValid() error {
	return nil
}
func (input *EtlWorkerNotifySendNotificationInput) CheckValid() error {
	return nil
}
func (input *DataReportFlowInput) CheckValid() error {
	return nil
}
func (input *DataReportInput) CheckValid() error {
	return nil
}
func (input *DataReportSendNotificationInput) CheckValid() error {
	return nil
}
