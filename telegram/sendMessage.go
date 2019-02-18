package telegram

// SendMessage https://core.telegram.org/bots/api#sendmessage
type SendMessage struct {
	ChatID                int64        `json:"chat_id"`
	Text                  string       `json:"text"`
	ParseMode             string       `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool         `json:"disable_notification,omitempty"`
	ReplyToMessageID      int64        `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

func (b *Bot) SendMessage(message *SendMessage) (resp *Message, err error) {
	err = b.request("sendMessage", message, &resp)
	return
}
