package paystack_client

import "net/http"

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) (*Client, error) {
	return &Client{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}, nil
}
