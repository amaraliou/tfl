package tfl

import (
	"fmt"
	"net/http"
	"time"
)

// Client ...
type Client struct {
	appID     string
	appSecret string
	baseURL   string
	http      *http.Client
}

// NewClient ...
func NewClient(appID, appSecret, baseURL string) (*Client, error) {
	return &Client{
		appID:     appID,
		appSecret: appSecret,
		baseURL:   baseURL,
		http:      &http.Client{Timeout: 300 * time.Second},
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

func (client *Client) getAuthString() string {
	return fmt.Sprintf("app_id=%s&app_key=%s", client.appID, client.appSecret)
}

func (client *Client) getRequestURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", client.baseURL, endpoint)
}
