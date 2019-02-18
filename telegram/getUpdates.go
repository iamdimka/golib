package telegram

// GetUpdates https://core.telegram.org/bots/api#getupdates
type GetUpdates struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// GetUpdates https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(message *GetUpdates) (updates []*Update, err error) {
	err = b.request("getUpdates", message, updates)
	return
}
