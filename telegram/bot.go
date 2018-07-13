package telegram

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var (
	Endpoint string = "https://api.telegram.org/"
)

type Bot struct {
	token  string
	base   string
	client *http.Client
}

func New(token string) *Bot {
	return &Bot{
		token: token,
		base:  strings.Join([]string{Endpoint, "bot", token}, ""),
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

	url := strings.Join([]string{b.base, method}, "/")
	data, err = json.Marshal(payload)

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
		return &r
	}

	return json.Unmarshal(r.Result, &response)
}
