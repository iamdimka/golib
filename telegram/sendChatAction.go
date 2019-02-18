package telegram

// SendChatAction https://core.telegram.org/bots/api#sendchataction
type SendChatAction struct {
	ChatID int64  `json:"chat_id"`
	Action string `json:"action"`
}

// SendChatAction https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(payload *SendChatAction) error {
	return b.request("sendChatAction", payload, nil)
}
