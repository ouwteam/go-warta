package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-warta/src/api_param"
	"go-warta/src/api_response"
	"go-warta/src/entity"
	"go-warta/src/helper"
	"go-warta/src/usecase"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	chat_id, _ := strconv.Atoi(r.URL.Query().Get("chat_id"))
	text := r.URL.Query().Get("text")
	var Output entity.Response
	var err error

	var Payload = &api_param.SendMessage{
		ChatID: int64(chat_id),
		Text:   text,
		ReplyMarkup: &api_param.ReplyMarkup{
			InlineKeyboard: &[][]api_param.InlineButton{
				{
					api_param.InlineButton{Text: "test 1", CallbackData: "test 1"},
					api_param.InlineButton{Text: "test 2", CallbackData: "test 2"},
				},
			},
		},
	}

	var reqBody, _ = json.Marshal(Payload)
	Address = Address + BotID + "/sendMessage"
	request, err := http.Post(Address, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		Output = entity.Response{
			Info:    false,
			Message: err.Error(),
			Content: nil,
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseBadRequest(b)
		return
	}

	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		Output = entity.Response{
			Info:    false,
			Message: err.Error(),
			Content: nil,
		}

		b, _ := json.Marshal(Output)
		mResponse.ResponseBadRequest(b)
		return
	}

	Output = entity.Response{
		Info:    true,
		Message: "",
		Content: string(body),
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
