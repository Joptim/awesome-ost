package client

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	addr    string
	timeout int
}

func (c *Client) Tracks() (map[string][]string, error) {
	endpoint := fmt.Sprintf("%s/tracks/", c.addr)
	ctx := context.WithTimeout(
		context.Background(),
		time.Duration(c.timeout)*time.Second,
	)
	req, err := http.NewRequestWithContext("GET", endpoint, nil)
	if err != nil {

	}
}

func (c *Client) AddTrack(id string) error {
	return nil
}

func (c *Client) RemoveTrack(id string) error {
	return nil
}

func (c *Client) RemoveAll() error {
	return nil
}

func New(addr string, timeout int) *Client {
	return &Client{
		addr:    addr,
		timeout: timeout,
	}
}
