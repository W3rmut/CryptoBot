package data

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type UpdateResponse struct {
	Result []Update
}

type MessageRequest struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
