package usecase

import (
	"bytes"
	"encoding/json"
	"go-warta/src/api_param"
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
	_, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	return nil
}
