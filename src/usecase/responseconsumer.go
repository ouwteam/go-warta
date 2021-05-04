package usecase

import (
	"fmt"
	"go-warta/src/api_response"

	"github.com/jmoiron/sqlx"
)

type GetUpdateConsumer struct {
	Reponse *api_response.GetUpdate
	*sqlx.DB
}

func (c *GetUpdateConsumer) Consume() bool {
	if !c.Reponse.Ok {
		return false
	}

	/**
	TODO :
	- implementasikan offset
	- kirim balikan setelah message/callback di consume
	- message /start simpan chat_id as UID, simpan nama juga sebagai username jika username tersedia update chat_id
	*/
	var LastID int64
	for _, item := range *c.Reponse.Result {
		if item.CallbackQuery != nil {
			fmt.Println("a callback")
		}

		if item.Message != nil {
			fmt.Println("a message")
		}

		LastID = int64(item.UpdateID)
	}

	uc := &LastUpdate{
		DB: c.DB,
	}

	return uc.SetLastUpdateID(LastID)
}
