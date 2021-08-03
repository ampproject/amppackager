package client

import (
	lwApi "github.com/liquidweb/go-lwApi"
)

// Client provides the HTTP backend.
type Client struct {
	config     *lwApi.LWAPIConfig
	httpClient *lwApi.Client
}

// NewClient returns a prepared API client.
func NewClient(config *lwApi.LWAPIConfig) (*Client, error) {
	client := &Client{}
	httpClient, err := lwApi.New(config)
	if err == nil {
		client.config = config
		client.httpClient = httpClient
	}

	return client, err
}
