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

// checks if the ReferenceMetadata3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceMetadata3{}

// ReferenceMetadata3 Only returned by the server when explicitly requested by the client and contains the following information:  - numCommitsAhead (number of commits ahead of the default branch)  - numCommitsBehind (number of commits behind the default branch)  - commitMetaOfHEAD (the commit metadata of the HEAD commit)  - commonAncestorHash (the hash of the common ancestor in relation to the default branch).  - numTotalCommits (the total number of commits in this reference). 
type ReferenceMetadata3 struct {
	NumCommitsAhead *int32 `json:"numCommitsAhead,omitempty"`
	NumCommitsBehind *int32 `json:"numCommitsBehind,omitempty"`
	CommitMetaOfHEAD *CommitMeta3 `json:"commitMetaOfHEAD,omitempty"`
	CommonAncestorHash *string `json:"commonAncestorHash,omitempty"`
	NumTotalCommits *int64 `json:"numTotalCommits,omitempty"`
}

// NewReferenceMetadata3 instantiates a new ReferenceMetadata3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceMetadata3() *ReferenceMetadata3 {
	this := ReferenceMetadata3{}
	return &this
}

// NewReferenceMetadata3WithDefaults instantiates a new ReferenceMetadata3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceMetadata3WithDefaults() *ReferenceMetadata3 {
	this := ReferenceMetadata3{}
	return &this
}

// GetNumCommitsAhead returns the NumCommitsAhead field value if set, zero value otherwise.
func (o *ReferenceMetadata3) GetNumCommitsAhead() int32 {
	if o == nil || IsNil(o.NumCommitsAhead) {
		var ret int32
		return ret
	}
	return *o.NumCommitsAhead
}

// GetNumCommitsAheadOk returns a tuple with the NumCommitsAhead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata3) GetNumCommitsAheadOk() (*int32, bool) {
	if o == nil || IsNil(o.NumCommitsAhead) {
		return nil, false
	}
	return o.NumCommitsAhead, true
}

// HasNumCommitsAhead returns a boolean if a field has been set.
func (o *ReferenceMetadata3) HasNumCommitsAhead() bool {
	if o != nil && !IsNil(o.NumCommitsAhead) {
		return true
	}

	return false
}

// SetNumCommitsAhead gets a reference to the given int32 and assigns it to the NumCommitsAhead field.
func (o *ReferenceMetadata3) SetNumCommitsAhead(v int32) {
	o.NumCommitsAhead = &v
}

// GetNumCommitsBehind returns the NumCommitsBehind field value if set, zero value otherwise.
func (o *ReferenceMetadata3) GetNumCommitsBehind() int32 {
	if o == nil || IsNil(o.NumCommitsBehind) {
		var ret int32
		return ret
	}
	return *o.NumCommitsBehind
}

// GetNumCommitsBehindOk returns a tuple with the NumCommitsBehind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata3) GetNumCommitsBehindOk() (*int32, bool) {
	if o == nil || IsNil(o.NumCommitsBehind) {
		return nil, false
	}
	return o.NumCommitsBehind, true
}

// HasNumCommitsBehind returns a boolean if a field has been set.
func (o *ReferenceMetadata3) HasNumCommitsBehind() bool {
	if o != nil && !IsNil(o.NumCommitsBehind) {
		return true
	}

	return false
}

// SetNumCommitsBehind gets a reference to the given int32 and assigns it to the NumCommitsBehind field.
func (o *ReferenceMetadata3) SetNumCommitsBehind(v int32) {
	o.NumCommitsBehind = &v
}

// GetCommitMetaOfHEAD returns the CommitMetaOfHEAD field value if set, zero value otherwise.
func (o *ReferenceMetadata3) GetCommitMetaOfHEAD() CommitMeta3 {
	if o == nil || IsNil(o.CommitMetaOfHEAD) {
		var ret CommitMeta3
		return ret
	}
	return *o.CommitMetaOfHEAD
}

// GetCommitMetaOfHEADOk returns a tuple with the CommitMetaOfHEAD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata3) GetCommitMetaOfHEADOk() (*CommitMeta3, bool) {
	if o == nil || IsNil(o.CommitMetaOfHEAD) {
		return nil, false
	}
	return o.CommitMetaOfHEAD, true
}

// HasCommitMetaOfHEAD returns a boolean if a field has been set.
func (o *ReferenceMetadata3) HasCommitMetaOfHEAD() bool {
	if o != nil && !IsNil(o.CommitMetaOfHEAD) {
		return true
	}

	return false
}

// SetCommitMetaOfHEAD gets a reference to the given CommitMeta3 and assigns it to the CommitMetaOfHEAD field.
func (o *ReferenceMetadata3) SetCommitMetaOfHEAD(v CommitMeta3) {
	o.CommitMetaOfHEAD = &v
}

// GetCommonAncestorHash returns the CommonAncestorHash field value if set, zero value otherwise.
func (o *ReferenceMetadata3) GetCommonAncestorHash() string {
	if o == nil || IsNil(o.CommonAncestorHash) {
		var ret string
		return ret
	}
	return *o.CommonAncestorHash
}

// GetCommonAncestorHashOk returns a tuple with the CommonAncestorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata3) GetCommonAncestorHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommonAncestorHash) {
		return nil, false
	}
	return o.CommonAncestorHash, true
}

// HasCommonAncestorHash returns a boolean if a field has been set.
func (o *ReferenceMetadata3) HasCommonAncestorHash() bool {
	if o != nil && !IsNil(o.CommonAncestorHash) {
		return true
	}

	return false
}

// SetCommonAncestorHash gets a reference to the given string and assigns it to the CommonAncestorHash field.
func (o *ReferenceMetadata3) SetCommonAncestorHash(v string) {
	o.CommonAncestorHash = &v
}

// GetNumTotalCommits returns the NumTotalCommits field value if set, zero value otherwise.
func (o *ReferenceMetadata3) GetNumTotalCommits() int64 {
	if o == nil || IsNil(o.NumTotalCommits) {
		var ret int64
		return ret
	}
	return *o.NumTotalCommits
}

// GetNumTotalCommitsOk returns a tuple with the NumTotalCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceMetadata3) GetNumTotalCommitsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumTotalCommits) {
		return nil, false
	}
	return o.NumTotalCommits, true
}

// HasNumTotalCommits returns a boolean if a field has been set.
func (o *ReferenceMetadata3) HasNumTotalCommits() bool {
	if o != nil && !IsNil(o.NumTotalCommits) {
		return true
	}

	return false
}

// SetNumTotalCommits gets a reference to the given int64 and assigns it to the NumTotalCommits field.
func (o *ReferenceMetadata3) SetNumTotalCommits(v int64) {
	o.NumTotalCommits = &v
}

func (o ReferenceMetadata3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceMetadata3) ToMap() (map[string]interface{}, error) {
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

type NullableReferenceMetadata3 struct {
	value *ReferenceMetadata3
	isSet bool
}

func (v NullableReferenceMetadata3) Get() *ReferenceMetadata3 {
	return v.value
}

func (v *NullableReferenceMetadata3) Set(val *ReferenceMetadata3) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceMetadata3) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceMetadata3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceMetadata3(val *ReferenceMetadata3) *NullableReferenceMetadata3 {
	return &NullableReferenceMetadata3{value: val, isSet: true}
}

func (v NullableReferenceMetadata3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceMetadata3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


