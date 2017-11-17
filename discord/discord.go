package discord

import (
	"net/http"
	"time"
)

type Discord struct {
	webhook string
	client  *http.Client
}

func New() *Discord {
	return &Discord{
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
	}
}
