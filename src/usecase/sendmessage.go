package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-warta/src/api_param"
	"go-warta/src/api_response"
	"io/ioutil"
	"net/http"
)

type SendMessage struct {
	Address string
	BotID   string
}

func (s *SendMessage) Send(Payload *api_param.SendMessage) error {
	var reqBody, _ = json.Marshal(Payload)
	s.Address = s.Address + s.BotID + "/sendMessage"
	request, err := http.Post(s.Address, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer request.Body.Close()
	res, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	var Response api_response.SendMessage
	if err = json.Unmarshal(res, &Response); err != nil {
		fmt.Println("Parse response failed")
		return err
	}

	if !Response.Ok {
		return errors.New(Response.Description)
	}

	return nil
}
