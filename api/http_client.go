package api

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

// BaseURL returns the base URL of the client
func (c *Client) BaseURL() string {
	return c.baseURL
}

// HTTPClient returns the HTTP client
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// Dummy is a test function to verify client functionality
func (c *Client) Dummy() string {
	return "Dummy reached"
}