package route

import (
	"encoding/json"
	"fmt"
	"go-warta/src/api_param"
	"go-warta/src/api_response"
	"go-warta/src/entity"
	"go-warta/src/helper"
	"go-warta/src/usecase"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
)

type RouteMain struct {
	*sqlx.DB
}

func NewRouteMain(db *sqlx.DB) *RouteMain {
	return &RouteMain{
		DB: db,
	}
}

func (route *RouteMain) HandleMain(rw http.ResponseWriter, r *http.Request) {
	mResponse := helper.NewResponse(r, rw)
	mResponse.ResponseOK([]byte("OK"))
}

func (route *RouteMain) HandleSendMessage(rw http.ResponseWriter, r *http.Request) {
	Address := os.Getenv("BOT_ADDRESS")
	BotID := os.Getenv("BOT_ID")
	mResponse := helper.NewResponse(r, rw)
	text := r.URL.Query().Get("text")
	channel_code := r.URL.Query().Get("channel")

	Rows, err := route.DB.Queryx("select * from user_channels where channel_code = ?", channel_code)
	if err != nil {
		Output := &entity.Response{
			Info:    false,
			Message: "no user channel found for the channel",
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseBadRequest(b)
		return
	}

	defer Rows.Close()
	for Rows.Next() {
		var userChannel entity.UserChannel
		Rows.StructScan(&userChannel)

		go func(mUserChannel *entity.UserChannel) {
			var Payload = &api_param.SendMessage{
				ChatID: mUserChannel.ChatID,
				Text:   text,
			}

			Sender := &usecase.SendMessage{
				Address: Address,
				BotID:   BotID,
			}

			err = Sender.Send(Payload)
			if err != nil {
				fmt.Println(err.Error())
			}
		}(&userChannel)
	}

	Output := &entity.Response{
		Info:    true,
		Message: "Added to queue",
	}

	b, _ := json.Marshal(Output)
	mResponse.ResponseOK(b)
}

func (route *RouteMain) HandleBotInfo(rw http.ResponseWriter, r *http.Request) {
	Address := os.Getenv("BOT_ADDRESS")
	BotID := os.Getenv("BOT_ID")
	mResponse := helper.NewResponse(r, rw)
	var Output entity.Response
	var err error
	var client = &http.Client{}
	var data api_response.GetMe

	Address = Address + BotID + "/getMe"
	fmt.Println(Address)
	request, err := http.NewRequest("GET", Address, nil)
	if err != nil {
		Output = entity.Response{
			Info:    false,
			Message: err.Error(),
			Content: nil,
		}
	}

	response, err := client.Do(request)
	if err != nil {
		Output = entity.Response{
			Info:    false,
			Message: err.Error(),
			Content: nil,
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseOK(b)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		Output = entity.Response{
			Info:    false,
			Message: response.Status,
			Content: nil,
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseBadRequest(b)
		return
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		Output = entity.Response{
			Info:    false,
			Message: err.Error(),
			Content: nil,
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseOK(b)
		return
	}

	Output = entity.Response{
		Info:    true,
		Message: "",
		Content: data.Result,
	}

	b, _ := json.Marshal(Output)
	mResponse.ResponseOK(b)
}

func (route *RouteMain) HandleGetUpdate(rw http.ResponseWriter, r *http.Request) {
	ucLastUpdate := &usecase.LastUpdate{
		DB: route.DB,
	}
	mResponse := helper.NewResponse(r, rw)
	Result, err := ucLastUpdate.GetUpdates()
	if err != nil {
		Output, _ := json.Marshal(&entity.Response{
			Info:    false,
			Message: err.Error(),
		})

		mResponse.ResponseBadRequest(Output)
		return
	}

	Output, _ := json.Marshal(&entity.Response{
		Info:    true,
		Content: Result,
	})
	mResponse.ResponseOK(Output)
}
