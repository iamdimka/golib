package telegram

type Handler interface {
	HandleTelegramUpdate(*Update)
}

type HandlerFunc func(*Update)

func (fn HandlerFunc) HandleTelegramUpdate(up *Update) {
	fn(up)
}

func (b *Bot) HandleUpdates(offset int, h Handler) (err error) {
	req := &GetUpdatesRequest{
		Offset:  offset,
		Timeout: 30,
	}

	updates := []*Update{}

	for {
		err = b.request("getUpdates", req, &updates)
		if err != nil {
			return
		}

		for _, up := range updates {
			h.HandleTelegramUpdate(up)
			req.Offset = up.UpdateID + 1
		}
	}

	return
}
