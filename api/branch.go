package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"github.com/pixisai/go-nessie/models"
	"github.com/pixisai/go-nessie/utils"
)

// GetAllBranches returns all trees/branches from the Nessie API
func (c *Client) GetAllBranches() (*models.GetReferencesResponse, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL(), utils.APIBasePath)
	fmt.Println("Hello, from Nessie - nix flake update success")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for /%s: %w", utils.APIBasePath, err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GEEEEEET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from GET /%s: %w", utils.APIBasePath, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /%s failed with status %d: %s", utils.APIBasePath, resp.StatusCode, string(body))
	}

	var response models.GetReferencesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse GetReferencesResponse: %w\nResponse body: %s", err, string(body))
	}

	return &response, nil
}

// GetBranch retrieves a specific branch from the Nessie API
func (c *Client) GetBranch(name string) (*models.Reference, error) {
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL(), utils.APIBasePath, name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for /%s/%s: %w", utils.APIBasePath, name, err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from GET /%s/%s: %w", utils.APIBasePath, name, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /%s/%s failed with status %d: %s", utils.APIBasePath, name, resp.StatusCode, string(body))
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
	baseURL := fmt.Sprintf("%s/%s", c.BaseURL(), utils.APIBasePath)
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
		return nil, fmt.Errorf("failed to create POST request for /%s: %w", utils.APIBasePath, err)
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
		return nil, fmt.Errorf("failed to read response body from POST /%s: %w", utils.APIBasePath, err)
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
	baseURL := fmt.Sprintf("%s/%s/%s@%s", c.BaseURL(), utils.APIBasePath, request.Name, request.Hash)
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
		return fmt.Errorf("failed to create DELETE request for /%s/%s: %w", utils.APIBasePath, request.Name, err)
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

// CommitChange commits a change to a branch
func (c *Client) CommitChange(branchName, hash string, operations []models.Operation, message string) (*models.CommitResponse, error) {
	// Build URL with branch and hash
	baseURL := fmt.Sprintf("%s/api/v2/trees/%s@%s/history/commit", 
		c.BaseURL(), branchName, hash)

	// Create commit request with metadata
	reqBody := map[string]interface{}{
		"operations": operations,
		"commitMeta": map[string]interface{}{
			"message":    message,
			"author":     "nessie-user",
			"authorTime": time.Now().Format(time.RFC3339),
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal commit request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request for commit: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute POST request to %s: %w\nRequest body: %s", 
			baseURL, err, string(jsonBody))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read commit response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("commit failed with status %d: %s\nRequest body: %s", 
			resp.StatusCode, string(body), string(jsonBody))
	}

	// Print response for debugging
	fmt.Printf("Commit response: %s\n", string(body))

	// Parse the full commit response
	var response models.CommitResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse commit response: %w\nResponse body: %s", 
			err, string(body))
	}

	// Verify we got a valid reference back
	if response.TargetBranch.Hash == "" {
		return nil, fmt.Errorf("commit succeeded but returned empty hash. Response: %s", string(body))
	}

	return &response, nil
}

// GetContent retrieves content at a specific key in a branch
func (c *Client) GetContent(branchName, hash string, key models.ContentKey) (*models.Content, error) {
	// Build URL with branch, hash, and key
	elements := strings.Join(key.Elements, "/")
	baseURL := fmt.Sprintf("%s/api/v2/trees/%s@%s/contents/%s", 
		c.BaseURL(), branchName, hash, elements)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for content: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", baseURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read content response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get content failed with status %d: %s", 
			resp.StatusCode, string(body))
	}

	var content models.Content
	if err := json.Unmarshal(body, &content); err != nil {
		return nil, fmt.Errorf("failed to parse content response: %w\nResponse body: %s", 
			err, string(body))
	}

	return &content, nil
}
