package telegram

// GetUpdatesRequest https://core.telegram.org/bots/api#getupdates
type GetUpdatesRequest struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (b *Bot) GetUpdates(message *GetUpdatesRequest) (updates []*Update, err error) {
	err = b.request("getUpdates", message, &updates)
	return
}
