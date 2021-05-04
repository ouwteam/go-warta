package usecase

import (
	"fmt"
	"go-warta/src/api_response"
)

type GetUpdateConsumer struct {
	Reponse *api_response.GetUpdate
}

func (c *GetUpdateConsumer) Consume() bool {
	if !c.Reponse.Ok {
		return false
	}

	/**
	TODO :
	- simpan last update_id ke db
	- implementasikan offset
	- kirim balikan setelah message/callback di consume
	- message /start simpan chat_id as UID, simpan nama juga sebagai username jika username tersedia update chat_id
	*/
	for _, item := range *c.Reponse.Result {
		if item.CallbackQuery != nil {
			fmt.Println("a callback")
		}

		if item.Message != nil {
			fmt.Println("a message")
		}
	}

	return true
}
