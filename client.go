package tcmd

import (
	"github.com/gempir/go-twitch-irc/v4"
)

type Client struct {
	t *twitch.Client
}

func NewClient(t *twitch.Client) *Client {
	return &Client{
		t: t,
	}
}
