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

// checks if the ReferenceHistoryState1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceHistoryState1{}

// ReferenceHistoryState1 Consistency status of the current HEAD commit.
type ReferenceHistoryState1 struct {
	// Nessie commit ID.
	CommitHash *string `json:"commitHash,omitempty"`
	// Consistency status of the commit.
	CommitConsistency *string `json:"commitConsistency,omitempty"`
	Meta *CommitMeta4 `json:"meta,omitempty"`
}

// NewReferenceHistoryState1 instantiates a new ReferenceHistoryState1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceHistoryState1() *ReferenceHistoryState1 {
	this := ReferenceHistoryState1{}
	return &this
}

// NewReferenceHistoryState1WithDefaults instantiates a new ReferenceHistoryState1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceHistoryState1WithDefaults() *ReferenceHistoryState1 {
	this := ReferenceHistoryState1{}
	return &this
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *ReferenceHistoryState1) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash) {
		var ret string
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceHistoryState1) GetCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *ReferenceHistoryState1) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given string and assigns it to the CommitHash field.
func (o *ReferenceHistoryState1) SetCommitHash(v string) {
	o.CommitHash = &v
}

// GetCommitConsistency returns the CommitConsistency field value if set, zero value otherwise.
func (o *ReferenceHistoryState1) GetCommitConsistency() string {
	if o == nil || IsNil(o.CommitConsistency) {
		var ret string
		return ret
	}
	return *o.CommitConsistency
}

// GetCommitConsistencyOk returns a tuple with the CommitConsistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceHistoryState1) GetCommitConsistencyOk() (*string, bool) {
	if o == nil || IsNil(o.CommitConsistency) {
		return nil, false
	}
	return o.CommitConsistency, true
}

// HasCommitConsistency returns a boolean if a field has been set.
func (o *ReferenceHistoryState1) HasCommitConsistency() bool {
	if o != nil && !IsNil(o.CommitConsistency) {
		return true
	}

	return false
}

// SetCommitConsistency gets a reference to the given string and assigns it to the CommitConsistency field.
func (o *ReferenceHistoryState1) SetCommitConsistency(v string) {
	o.CommitConsistency = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReferenceHistoryState1) GetMeta() CommitMeta4 {
	if o == nil || IsNil(o.Meta) {
		var ret CommitMeta4
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceHistoryState1) GetMetaOk() (*CommitMeta4, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReferenceHistoryState1) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CommitMeta4 and assigns it to the Meta field.
func (o *ReferenceHistoryState1) SetMeta(v CommitMeta4) {
	o.Meta = &v
}

func (o ReferenceHistoryState1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceHistoryState1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommitHash) {
		toSerialize["commitHash"] = o.CommitHash
	}
	if !IsNil(o.CommitConsistency) {
		toSerialize["commitConsistency"] = o.CommitConsistency
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableReferenceHistoryState1 struct {
	value *ReferenceHistoryState1
	isSet bool
}

func (v NullableReferenceHistoryState1) Get() *ReferenceHistoryState1 {
	return v.value
}

func (v *NullableReferenceHistoryState1) Set(val *ReferenceHistoryState1) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceHistoryState1) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceHistoryState1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceHistoryState1(val *ReferenceHistoryState1) *NullableReferenceHistoryState1 {
	return &NullableReferenceHistoryState1{value: val, isSet: true}
}

func (v NullableReferenceHistoryState1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceHistoryState1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


