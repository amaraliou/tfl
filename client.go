package tfl

// Client ...
type Client struct {
	AppID     string
	AppSecret string
	// To add other fields
}

// NewClient ...
func NewClient(appID, appSecret string) (*Client, error) {
	return &Client{}, nil
}
