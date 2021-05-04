package api_response

type GetUpdate struct {
	Ok     bool      `json:"ok"`
	Result *[]Result `json:"result"`
}

type Result struct {
	UpdateID      int            `json:"update_id"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	Message       *Message       `json:"message,omitempty"`
}

type CallbackQuery struct {
	ID           string   `json:"id"`
	From         *From    `json:"from"`
	Message      *Message `json:"message"`
	ChatInstance string   `json:"chat_instance"`
	Data         string   `json:"data"`
}

type Message struct {
	MessageID   int          `json:"message_id"`
	From        *From        `json:"from"`
	Chat        *Chat        `json:"chat"`
	Date        int          `json:"date"`
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]struct {
		Text         string `json:"text"`
		CallbackData string `json:"callback_data"`
	} `json:"inline_keyboard"`
}
