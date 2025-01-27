package groot

import (
	"net/http"
	"time"
)

// Client represents a Nessie API client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Nessie API client
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}