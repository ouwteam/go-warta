package entity

type MessageEntity struct {
	Type     int64  `json:"type"`
	Offset   int64  `json:"offset"`
	Length   int64  `json:"length"`
	Url      int64  `json:"url"`
	User     *User  `json:"user"`
	Language string `json:"language"`
}
