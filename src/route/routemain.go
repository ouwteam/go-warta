package route

import (
	"encoding/json"
	"fmt"
	"go-warta/src/api_response"
	"go-warta/src/entity"
	"go-warta/src/helper"
	"net/http"
	"net/url"
	"os"
)

type RouteMain struct{}

func NewRouteMain() *RouteMain {
	return &RouteMain{}
}

func (route *RouteMain) HandleMain(rw http.ResponseWriter, r *http.Request) {
	mResponse := helper.NewResponse(r, rw)
	mResponse.ResponseOK([]byte("OK"))
}

func (route *RouteMain) HandleSendMessage(rw http.ResponseWriter, r *http.Request) {
	Address := os.Getenv("BOT_ADDRESS")
	BotID := os.Getenv("BOT_ID")
	mResponse := helper.NewResponse(r, rw)
	chat_id := r.URL.Query().Get("chat_id")
	text := r.URL.Query().Get("text")
	var Output entity.Response
	var err error
	var client = &http.Client{}
	var data api_response.GetMe
	var payload = url.Values{}

	payload.Set("chat_id", chat_id)
	payload.Set("text", text)

	Address = Address + BotID + "/sendMessage"
	fmt.Println(Address)
	request, err := http.NewRequest("POST", Address, nil)
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
