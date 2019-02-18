package telegram

// DeleteMessage https://core.telegram.org/bots/api#deletemessage
type DeleteMessage struct {
	ChatID                int64        `json:"chat_id"`
	MessageID             int64        `json:"message_id,omitempty"`
}

// DeleteMessage https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(message *DeleteMessage) (resp *bool, err error) {
	err = b.request("deleteMessage", message, resp)
	return
}
