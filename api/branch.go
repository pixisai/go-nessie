package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/pixisai/go-nessie/models"
)

// GetAllTrees returns all trees/branches from the Nessie API
func (c *Client) GetAllBranches() (*models.GetReferencesResponse, error) {
	url := fmt.Sprintf("%s/api/v2/trees", c.BaseURL())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var response models.GetReferencesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}
