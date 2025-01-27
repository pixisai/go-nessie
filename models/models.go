package models

import "time"

// CommitMetaData represents commit metadata
type CommitMetaData struct {
	Author      string    `json:"author,omitempty"`
	AuthorTime  time.Time `json:"authorTime,omitempty"`
	Committer   string    `json:"committer,omitempty"`
	CommitTime  time.Time `json:"commitTime,omitempty"`
	Message     string    `json:"message,omitempty"`
	SignedOffBy string    `json:"signedOffBy,omitempty"`
}

// Metadata represents the metadata associated with a reference
type Metadata struct {
	CommitMetaData CommitMetaData    `json:"commitMetaData,omitempty"`
	Properties     map[string]string `json:"properties,omitempty"`
}

// Reference represents a Nessie reference (branch or tag)
type Reference struct {
	Type        string    `json:"type"` // BRANCH or TAG
	Name        string    `json:"name"`
	Hash        string    `json:"hash"`
	Metadata    Metadata  `json:"metadata,omitempty"`
	CreatedTime time.Time `json:"createdTime,omitempty"`
}

// GetReferencesResponse represents the response from the getAllReferences endpoint
type GetReferencesResponse struct {
	Token      string      `json:"token,omitempty"` // Token for the next page
	References []Reference `json:"references"`      // List of references
	HasMore    bool        `json:"hasMore"`         // Indicates if there are more results
}

// CreateReferenceRequest represents the request to create a new reference
type CreateReferenceRequest struct {
	Type        string   `json:"type"`              // BRANCH or TAG
	Name        string   `json:"name"`              // Name of the reference to create
	Reference   string   `json:"reference"`         // Source reference name
	Hash        string   `json:"hash,omitempty"`    // Hash to create reference at
	Metadata    Metadata `json:"metadata,omitempty"` // Optional metadata
}

// DeleteReferenceRequest represents the request to delete a reference
type DeleteReferenceRequest struct {
	Name string `json:"name"`           // Name of the reference to delete
	Hash string `json:"hash,omitempty"` // Expected hash of the reference
}
