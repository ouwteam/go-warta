package entity

type Response struct {
	Info    bool        `json:"info"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
