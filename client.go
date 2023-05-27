package tcmd

import (
	"errors"

	"github.com/gempir/go-twitch-irc/v4"
)

// Client encapsulates a client that interacting with Twitch IRC and tcmd
type Client struct {
	t      *twitch.Client
	prefix string
}

// ClientOption provides a variadic option for configuring the client
type ClientOption func(c *Client) error

// WithPrefix configures the prefix used to trigger commands
func WithPrefix(prefix string) ClientOption {
	return func(c *Client) error {
		return c.SetPrefix(prefix)
	}
}

// NewClient returns an instance of the tcmd client
func NewClient(t *twitch.Client, opts ...ClientOption) (*Client, error) {
	c := &Client{
		t:      t,
		prefix: "!", // Set default prefix
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// GetPrefix returns the current prefix used to trigger commands
func (c *Client) GetPrefix() string {
	return c.prefix
}

// SetPrefix sets the prefix used to trigger commands
func (c *Client) SetPrefix(prefix string) error {
	if prefix == "" {
		return errors.New("missing prefix")
	}

	c.prefix = prefix
	return nil
}
