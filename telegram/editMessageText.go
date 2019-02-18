package telegram

// EditMessageText https://core.telegram.org/bots/api#editmessagetext
type EditMessageText struct {
	ChatID                int64        `json:"chat_id"`
	MessageID             int64        `json:"message_id,omitempty"`
	InlineMessageID       string       `json:"inline_message_id,omitempty"`
	Text                  string       `json:"text,omitempty"`
	ParseMode             string       `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool         `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageText https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(message *EditMessageText) (resp *Message, err error) {
	err = b.request("editMessageText", message, resp)
	return
}
