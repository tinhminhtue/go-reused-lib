package inout

type NotifyMessageBody struct {
	Message Message `json:"message,omitempty"`
}

type Videos struct {
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Duration int    `json:"duration,omitempty"`
	TypeShow int    `json:"type_show,omitempty"`
	KeyVideo string `json:"key_video,omitempty"`
	KeyImg   string `json:"key_img,omitempty"`
	KeyThumb string `json:"key_thumb,omitempty"`
}
type Files struct {
	FileName    string `json:"file_name,omitempty"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	FileType    string `json:"file_type,omitempty"`
}
type Reply struct {
	MessageID int `json:"message_id,omitempty"`
}
type Forward struct {
	NotificationPath string `json:"notification_path,omitempty"`
	MessageID        int    `json:"message_id,omitempty"`
	ConversationID   int    `json:"conversation_id,omitempty"`
	DateTime         string `json:"date_time,omitempty"`
}
type Message struct {
	Content string   `json:"content,omitempty"`
	Videos  []Videos `json:"videos,omitempty"`
	Files   []Files  `json:"files,omitempty"`
	Reply   *Reply   `json:"reply,omitempty"`
	Forward *Forward `json:"forward,omitempty"`
}
