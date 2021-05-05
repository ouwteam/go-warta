package usecase

import (
	"fmt"
	"go-warta/src/api_param"
	"go-warta/src/api_response"
	"go-warta/src/entity"
	"os"

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
	- kirim balikan setelah callback di consume
	*/
	var LastID int64
	for _, item := range *c.Reponse.Result {
		if item.CallbackQuery != nil {
			fmt.Println("a callback")
		}

		if item.Message != nil {
			if item.Message.Text == "P" || item.Message.Text == "/start" {
				chat_id := item.Message.Chat.ID
				username := item.Message.Chat.Username

				fmt.Println("step 1 :")
				// step 1. Register the user
				Tx, err := c.Beginx()
				if err != nil {
					fmt.Println(err.Error())
					return false
				}

				Rows, err := c.DB.Queryx("select * from users where username = ?", username)
				if err != nil {
					fmt.Println(err.Error())
					return false
				}

				defer Rows.Close()
				if Rows.Next() {
					_, err = Tx.Exec("update users set chat_id = ? where username = ?", chat_id, username)
					if err != nil {
						fmt.Println(err.Error())
						Tx.Rollback()
					}
				} else {
					_, err = Tx.Exec("insert into users (chat_id, username) values (?, ?)", chat_id, username)
					if err != nil {
						fmt.Println(err.Error())
						Tx.Rollback()
					}
				}

				Tx.Commit()

				// step 2. Send to choose channels
				Rows, err = c.DB.Queryx("select * from channels")
				if err != nil {
					fmt.Println(err.Error())
					return false
				}

				defer Rows.Close()

				var Payload = &api_param.SendMessage{
					ChatID: int64(chat_id),
					Text:   "Choose channel",
					ReplyMarkup: &api_param.ReplyMarkup{
						InlineKeyboard: &[][]api_param.InlineButton{},
					},
				}

				mInlineButton := []api_param.InlineButton{}
				for Rows.Next() {
					var Channel entity.Channel
					err = Rows.StructScan(&Channel)
					if err != nil {
						fmt.Println(err.Error())
						return false
					}

					inlineButton := &api_param.InlineButton{
						Text:         Channel.Name,
						CallbackData: Channel.Code,
					}

					mInlineButton = append(mInlineButton, *inlineButton)
				}

				mInlineKeyboard := &[][]api_param.InlineButton{
					mInlineButton,
				}
				Payload.ReplyMarkup.InlineKeyboard = mInlineKeyboard

				Address := os.Getenv("BOT_ADDRESS")
				BotID := os.Getenv("BOT_ID")
				Sender := &SendMessage{
					Address: Address,
					BotID:   BotID,
				}

				err = Sender.Send(Payload)
				if err != nil {
					fmt.Println(err.Error())
					return false
				}
			}
		}

		LastID = int64(item.UpdateID)
	}

	uc := &LastUpdate{
		DB: c.DB,
	}

	return uc.SetLastUpdateID(LastID)
}
