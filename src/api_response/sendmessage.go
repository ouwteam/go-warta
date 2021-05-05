package api_response

type SendMessage struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
