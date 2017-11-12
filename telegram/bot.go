package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Bot struct {
	token  string
	client *http.Client
}

func New(token string) *Bot {
	return &Bot{
		token: token,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 5,
			},
			Timeout: time.Second * 31,
		},
	}
}

func (b *Bot) request(method string, payload, response interface{}) (err error) {
	var (
		data []byte
		res  *http.Response
	)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, method)
	data, err = json.Marshal(payload)
	fmt.Println(url, string(data))

	if err != nil {
		return
	}

	res, err = b.client.Post(url, "application/json; charset=utf8", bytes.NewReader(data))
	if err != nil {
		return
	}

	defer res.Body.Close()

	r := Response{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return
	}

	if !r.Ok {
		return &apiError{r.ErrorCode, r.ErrorDescription}
	}

	return json.Unmarshal(r.Result, &response)
}
