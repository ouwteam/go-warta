package api_response

import "go-warta/src/entity"

type GetMe struct {
	Ok     bool        `json:"ok"`
	Result entity.User `json:"result"`
}
