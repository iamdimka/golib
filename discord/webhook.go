package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Webhook struct {
	Content   string   `json:"content,omitempty"`
	Username  string   `json:"username,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`
	TTS       bool     `json:"tts,omitempty"`
	Embeds    []*Embed `json:"embeds,omitempty"`
}

func (d *Discord) SetWebHook(webhook string) {
	d.webhook = webhook
}

func (d *Discord) Webhook(payload *Webhook) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	res, err := d.client.Post(d.webhook, "application/json; charset=utf-8", bytes.NewReader(data))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	m := json.RawMessage{}
	err = json.NewDecoder(res.Body).Decode(m)
	if err != nil {
		return err
	}

	fmt.Println(string(m))
	return nil
}
