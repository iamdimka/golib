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

func (d *Discord) Textf(format string, args ...interface{}) error {
	return d.Text(fmt.Sprintf(format, args...))
}

func (d *Discord) Text(content string) error {
	return d.Webhook(&Webhook{
		Content: content,
	})
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

	fmt.Printf("%s\n", m)
	return nil
}
