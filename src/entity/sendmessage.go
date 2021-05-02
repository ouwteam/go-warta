package entity

type SendMessage struct {
	ChatID                   int64          `json:"chat_id"`
	Text                     string         `json:"text"`
	ParseMode                string         `json:"parse_mode"`
	Entities                 *MessageEntity `json:"entities"`
	DisableWebPagePreview    bool           `json:"disable_web_page_preview"`
	DisableNotification      bool           `json:"disable_notification"`
	ReplyToMessageID         int64          `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool           `json:"allow_sending_without_reply"`
}
