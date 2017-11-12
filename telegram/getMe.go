package telegram

// GetMe https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (user *User, err error) {
	err = b.request("getMe", nil, &user)
	return
}
