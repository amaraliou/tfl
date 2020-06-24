package tfl

import (
	"net/http"
	"time"
)

// Client ...
type Client struct {
	appID     string
	appSecret string
	baseURL   string
	client    *http.Client
}

// NewClient ...
func NewClient(appID, appSecret, baseURL string) (*Client, error) {
	return &Client{
		appID:     appID,
		appSecret: appSecret,
		baseURL:   baseURL,
		client:    &http.Client{Timeout: 300 * time.Second},
	}, nil
}

// GetAppID ...
func (client *Client) GetAppID() string {
	return client.appID
}

// GetAppSecret ...
func (client *Client) GetAppSecret() string {
	return client.appSecret
}

// GetBaseURL ...
func (client *Client) GetBaseURL() string {
	return client.baseURL
}
