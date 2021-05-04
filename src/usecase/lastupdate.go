package usecase

import (
	"encoding/json"
	"fmt"
	"go-warta/src/api_response"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

type LastUpdate struct {
	*sqlx.DB
}

func (l *LastUpdate) GetLastUpdateID() int64 {
	var last_update_id int64
	var last_update_at int64

	Row := l.DB.QueryRow("select last_update_id, last_update_at from update_record")
	err := Row.Scan(&last_update_id, &last_update_at)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return (last_update_id + 1)
}

func (l *LastUpdate) SetLastUpdateID(LastID int64) bool {
	if LastID == 0 {
		return false
	}

	now := time.Now()
	var err error

	Rows, err := l.DB.Query("select * from update_record")
	if err != nil {
		fmt.Println("On Select : " + err.Error())
		return false
	}

	defer Rows.Close()
	if Rows.Next() {
		_, err = l.DB.Exec("UPDATE update_record SET last_update_id = ?, last_update_at = ?", LastID, now.Unix())
		if err != nil {
			fmt.Println("On Update : " + err.Error())
			return false
		}
	} else {
		_, err = l.DB.Exec("INSERT INTO update_record (last_update_id, last_update_at) VALUES (?, ?)", LastID, now.Unix())
		if err != nil {
			fmt.Println("On Insert : " + err.Error())
			return false
		}
	}

	return err == nil
}

func (l *LastUpdate) GetUpdates() (api_response.GetUpdate, error) {
	Address := os.Getenv("BOT_ADDRESS")
	BotID := os.Getenv("BOT_ID")
	var err error
	var client = &http.Client{}
	var data api_response.GetUpdate

	lastID := l.GetLastUpdateID()
	Address = Address + BotID + fmt.Sprintf("/getUpdates?offset=%d", lastID)
	request, err := http.NewRequest("GET", Address, nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return data, err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	Consumer := &GetUpdateConsumer{
		Reponse: &data,
		DB:      l.DB,
	}

	go func(f *GetUpdateConsumer) {
		f.Consume()
	}(Consumer)

	return data, nil
}
