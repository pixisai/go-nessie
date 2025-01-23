/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReferenceMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceMetadata{}

// ReferenceMetadata Only returned by the server when explicitly requested by the client and contains the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference). 
type ReferenceMetadata struct {
	NumCommitsAhead *int32 `json:"numCommitsAhead,omitempty"`
	NumCommitsBehind *int32 `json:"numCommitsBehind,omitempty"`
	CommitMetaOfHEAD *CommitMeta3 `json:"commitMetaOfHEAD,omitempty"`
	CommonAncestorHash *string `json:"commonAncestorHash,omitempty"`
	NumTotalCommits *int64 `json:"numTotalCommits,omitempty"`
}

// NewReferenceMetadata instantiates a new ReferenceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceMetadata() *ReferenceMetadata {
	this := ReferenceMetadata{}
	return &this
}

// NewReferenceMetadataWithDefaults instantiates a new ReferenceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceMetadataWithDefaults() *ReferenceMetadata {
	this := ReferenceMetadata{}
	return &this
}

// GetNumCommitsAhead returns the NumCommitsAhead field value if set, zero value otherwise.
func (o *ReferenceMetadata) GetNumCommitsAhead() int32 {
	if o == nil || IsNil(o.NumCommitsAhead) {
		var ret int32
		return ret
	}
	return *o.NumCommitsAhead
}

// GetNumCommitsAheadOk returns a tuple with the NumCommitsAhead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata) GetNumCommitsAheadOk() (*int32, bool) {
	if o == nil || IsNil(o.NumCommitsAhead) {
		return nil, false
	}
	return o.NumCommitsAhead, true
}

// HasNumCommitsAhead returns a boolean if a field has been set.
func (o *ReferenceMetadata) HasNumCommitsAhead() bool {
	if o != nil && !IsNil(o.NumCommitsAhead) {
		return true
	}

	return false
}

// SetNumCommitsAhead gets a reference to the given int32 and assigns it to the NumCommitsAhead field.
func (o *ReferenceMetadata) SetNumCommitsAhead(v int32) {
	o.NumCommitsAhead = &v
}

// GetNumCommitsBehind returns the NumCommitsBehind field value if set, zero value otherwise.
func (o *ReferenceMetadata) GetNumCommitsBehind() int32 {
	if o == nil || IsNil(o.NumCommitsBehind) {
		var ret int32
		return ret
	}
	return *o.NumCommitsBehind
}

// GetNumCommitsBehindOk returns a tuple with the NumCommitsBehind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata) GetNumCommitsBehindOk() (*int32, bool) {
	if o == nil || IsNil(o.NumCommitsBehind) {
		return nil, false
	}
	return o.NumCommitsBehind, true
}

// HasNumCommitsBehind returns a boolean if a field has been set.
func (o *ReferenceMetadata) HasNumCommitsBehind() bool {
	if o != nil && !IsNil(o.NumCommitsBehind) {
		return true
	}

	return false
}

// SetNumCommitsBehind gets a reference to the given int32 and assigns it to the NumCommitsBehind field.
func (o *ReferenceMetadata) SetNumCommitsBehind(v int32) {
	o.NumCommitsBehind = &v
}

// GetCommitMetaOfHEAD returns the CommitMetaOfHEAD field value if set, zero value otherwise.
func (o *ReferenceMetadata) GetCommitMetaOfHEAD() CommitMeta3 {
	if o == nil || IsNil(o.CommitMetaOfHEAD) {
		var ret CommitMeta3
		return ret
	}
	return *o.CommitMetaOfHEAD
}

// GetCommitMetaOfHEADOk returns a tuple with the CommitMetaOfHEAD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata) GetCommitMetaOfHEADOk() (*CommitMeta3, bool) {
	if o == nil || IsNil(o.CommitMetaOfHEAD) {
		return nil, false
	}
	return o.CommitMetaOfHEAD, true
}

// HasCommitMetaOfHEAD returns a boolean if a field has been set.
func (o *ReferenceMetadata) HasCommitMetaOfHEAD() bool {
	if o != nil && !IsNil(o.CommitMetaOfHEAD) {
		return true
	}

	return false
}

// SetCommitMetaOfHEAD gets a reference to the given CommitMeta3 and assigns it to the CommitMetaOfHEAD field.
func (o *ReferenceMetadata) SetCommitMetaOfHEAD(v CommitMeta3) {
	o.CommitMetaOfHEAD = &v
}

// GetCommonAncestorHash returns the CommonAncestorHash field value if set, zero value otherwise.
func (o *ReferenceMetadata) GetCommonAncestorHash() string {
	if o == nil || IsNil(o.CommonAncestorHash) {
		var ret string
		return ret
	}
	return *o.CommonAncestorHash
}

// GetCommonAncestorHashOk returns a tuple with the CommonAncestorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata) GetCommonAncestorHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommonAncestorHash) {
		return nil, false
	}
	return o.CommonAncestorHash, true
}

// HasCommonAncestorHash returns a boolean if a field has been set.
func (o *ReferenceMetadata) HasCommonAncestorHash() bool {
	if o != nil && !IsNil(o.CommonAncestorHash) {
		return true
	}

	return false
}

// SetCommonAncestorHash gets a reference to the given string and assigns it to the CommonAncestorHash field.
func (o *ReferenceMetadata) SetCommonAncestorHash(v string) {
	o.CommonAncestorHash = &v
}

// GetNumTotalCommits returns the NumTotalCommits field value if set, zero value otherwise.
func (o *ReferenceMetadata) GetNumTotalCommits() int64 {
	if o == nil || IsNil(o.NumTotalCommits) {
		var ret int64
		return ret
	}
	return *o.NumTotalCommits
}

// GetNumTotalCommitsOk returns a tuple with the NumTotalCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata) GetNumTotalCommitsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumTotalCommits) {
		return nil, false
	}
	return o.NumTotalCommits, true
}

// HasNumTotalCommits returns a boolean if a field has been set.
func (o *ReferenceMetadata) HasNumTotalCommits() bool {
	if o != nil && !IsNil(o.NumTotalCommits) {
		return true
	}

	return false
}

// SetNumTotalCommits gets a reference to the given int64 and assigns it to the NumTotalCommits field.
func (o *ReferenceMetadata) SetNumTotalCommits(v int64) {
	o.NumTotalCommits = &v
}

func (o ReferenceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumCommitsAhead) {
		toSerialize["numCommitsAhead"] = o.NumCommitsAhead
	}
	if !IsNil(o.NumCommitsBehind) {
		toSerialize["numCommitsBehind"] = o.NumCommitsBehind
	}
	if !IsNil(o.CommitMetaOfHEAD) {
		toSerialize["commitMetaOfHEAD"] = o.CommitMetaOfHEAD
	}
	if !IsNil(o.CommonAncestorHash) {
		toSerialize["commonAncestorHash"] = o.CommonAncestorHash
	}
	if !IsNil(o.NumTotalCommits) {
		toSerialize["numTotalCommits"] = o.NumTotalCommits
	}
	return toSerialize, nil
}

type NullableReferenceMetadata struct {
	value *ReferenceMetadata
	isSet bool
}

func (v NullableReferenceMetadata) Get() *ReferenceMetadata {
	return v.value
}

func (v *NullableReferenceMetadata) Set(val *ReferenceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceMetadata(val *ReferenceMetadata) *NullableReferenceMetadata {
	return &NullableReferenceMetadata{value: val, isSet: true}
}

func (v NullableReferenceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


