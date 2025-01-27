package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"github.com/pixisai/go-nessie/models"
)

// GetAllBranches returns all trees/branches from the Nessie API
func (c *Client) GetAllBranches() (*models.GetReferencesResponse, error) {
	url := fmt.Sprintf("%s/api/v2/trees", c.BaseURL())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for /api/v2/trees: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from GET /api/v2/trees: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /api/v2/trees failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response models.GetReferencesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse GetReferencesResponse: %w\nResponse body: %s", err, string(body))
	}

	return &response, nil
}

// GetBranch retrieves a specific branch from the Nessie API
func (c *Client) GetBranch(name string) (*models.Reference, error) {
	url := fmt.Sprintf("%s/api/v2/trees/%s", c.BaseURL(), name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for /api/v2/trees/%s: %w", name, err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from GET /api/v2/trees/%s: %w", name, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /api/v2/trees/%s failed with status %d: %s", name, resp.StatusCode, string(body))
	}

	var response struct {
		Reference models.Reference `json:"reference"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Reference response: %w\nResponse body: %s", err, string(body))
	}

	return &response.Reference, nil
}

// CreateBranch creates a new branch in the Nessie repository
func (c *Client) CreateBranch(name, sourceRef string) (*models.Reference, error) {
	// First, get the hash of the source branch
	source, err := c.GetBranch(sourceRef)
	if err != nil {
		return nil, fmt.Errorf("failed to get source branch '%s': %w", sourceRef, err)
	}

	// Build URL with query parameters
	baseURL := fmt.Sprintf("%s/api/v2/trees", c.BaseURL())
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := u.Query()
	q.Set("name", name)
	q.Set("type", "BRANCH")
	u.RawQuery = q.Encode()

	// Create request body
	reqBody := map[string]interface{}{
		"type": "BRANCH",
		"hash": source.Hash,
		"name": sourceRef,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request for /api/v2/trees: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute POST request to %s: %w\nRequest body: %s", u.String(), err, string(jsonBody))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from POST /api/v2/trees: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("POST /api/v2/trees failed with status %d: %s\nRequest body: %s", 
			resp.StatusCode, string(body), string(jsonBody))
	}

	var response struct {
		Reference models.Reference `json:"reference"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse Reference response: %w\nResponse body: %s", err, string(body))
	}

	return &response.Reference, nil
}

// DeleteBranch deletes a branch from the Nessie repository
func (c *Client) DeleteBranch(request *models.DeleteReferenceRequest) error {
	// Build URL with hash
	baseURL := fmt.Sprintf("%s/api/v2/trees/%s@%s", c.BaseURL(), request.Name, request.Hash)
	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}

	// Add type parameter
	q := u.Query()
	q.Set("type", "branch")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create DELETE request for /api/v2/trees/%s: %w", request.Name, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute DELETE request to %s: %w", u.String(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("DELETE %s failed with status %d: %s", 
			u.String(), resp.StatusCode, string(body))
	}

	return nil
}
