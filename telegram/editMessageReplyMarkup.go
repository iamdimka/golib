package telegram

// EditMessageReplyMarkup https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkup struct {
	ChatID          int64        `json:"chat_id"`
	MessageID       int64        `json:"message_id,omitempty"`
	InlineMessageID string       `json:"inline_message_id,omitempty"`
	ReplyMarkup     *ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(message *EditMessageReplyMarkup) (err error) {
	err = b.request("editMessageReplyMarkup", message, nil)
	return
}
