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
	Type      string   `json:"type"`               // BRANCH or TAG
	Name      string   `json:"name"`               // Name of the reference to create
	Reference string   `json:"reference"`          // Source reference name
	Hash      string   `json:"hash,omitempty"`     // Hash to create reference at
	Metadata  Metadata `json:"metadata,omitempty"` // Optional metadata
}

// DeleteReferenceRequest represents the request to delete a reference
type DeleteReferenceRequest struct {
	Name string `json:"name"`           // Name of the reference to delete
	Hash string `json:"hash,omitempty"` // Expected hash of the reference
}

// ContentKey represents a key in the content store
type ContentKey struct {
	Elements []string `json:"elements"`
}

// Content represents the content of an operation
type Content struct {
	Type             string           `json:"type"`
	ContentId        string           `json:"contentId,omitempty"`
	MetadataLocation string           `json:"metadataLocation,omitempty"`
	SnapshotId       int64            `json:"snapshotId,omitempty"`
	SchemaId         int64            `json:"schemaId,omitempty"`
	SpecId           int64            `json:"specId,omitempty"`
	SortOrderId      int64            `json:"sortOrderId,omitempty"`
	IcebergTable     *IcebergTable    `json:"icebergTable,omitempty"`
	Schema           *Schema          `json:"schema"`
	PartitionSpec    []PartitionField `json:"partitionSpec,omitempty"`
}

// Operation represents a single operation in a commit
type Operation struct {
	Type    string     `json:"type"` // PUT, DELETE
	Key     ContentKey `json:"key"`
	Content *Content   `json:"content,omitempty"`
}

// CommitResponse represents the response from a commit operation
type CommitResponse struct {
	TargetBranch  Reference `json:"targetBranch"`
	AddedContents []struct {
		Key struct {
			Elements []string `json:"elements"`
		} `json:"key"`
		ContentId string `json:"contentId"`
	} `json:"addedContents"`
}

type Field struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Nullable bool   `json:"nullable"`
}

// Schema represents the schema definition for an Iceberg table
type Schema struct {
	SchemaId int64   `json:"schemaId"`
	Fields   []Field `json:"fields"`
}

type PartitionField struct {
	SourceId int    `json:"sourceId"`
	Name     string `json:"name"`
}

type IcebergTable struct {
	SchemaId         int64             `json:"schemaId"`
	SpecId           int64             `json:"specId"`
	MetadataLocation string            `json:"metadataLocation"`
	Schema           *Schema           `json:"schema"`
	PartitionSpec    []PartitionField  `json:"partitionSpec,omitempty"`
	CurrentSnapshot  *Snapshot         `json:"currentSnapshot,omitempty"`
	Snapshots        []Snapshot        `json:"snapshots,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
	LastUpdatedMs    int64             `json:"lastUpdatedMs,omitempty"`
	LastColumnId     int32             `json:"lastColumnId,omitempty"`
}

// Snapshot represents a point-in-time state of a table
type Snapshot struct {
	SnapshotId       int64           `json:"snapshotId"`
	ParentSnapshotId *int64          `json:"parentSnapshotId,omitempty"`
	SequenceNumber   int64           `json:"sequenceNumber"`
	TimestampMs      int64           `json:"timestampMs"`
	ManifestList     string          `json:"manifestList"`
	Summary          SnapshotSummary `json:"summary"`
	SchemaId         int64           `json:"schemaId"`
}

// SnapshotSummary contains summary statistics for a snapshot
type SnapshotSummary struct {
	Operation    string            `json:"operation"`
	Added        int64             `json:"added"`
	Deleted      int64             `json:"deleted"`
	Changed      int64             `json:"changed"`
	TotalRecords int64             `json:"totalRecords"`
	Properties   map[string]string `json:"properties,omitempty"`
}

type TableMeatadata struct {
	Content  			*Content    `json:"content"`
	EffectiveReference  *Reference  `json:"effectiveReference"`
}

// TableDetails represents detailed information about a table
type TableDetails struct {
	Table     *IcebergTable `json:"table"`
	Reference *Reference    `json:"reference"`
}
