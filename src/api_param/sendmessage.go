package api_param

type SendMessage struct {
	ChatID      int64        `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard *[][]InlineButton `json:"inline_keyboard"`
}

type InlineButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
