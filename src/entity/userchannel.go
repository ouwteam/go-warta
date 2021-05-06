package entity

type UserChannel struct {
	ChatID      int64  `json:"chat_id" db:"chat_id"`
	ChannelCode string `json:"channel_code" db:"channel_code"`
}
