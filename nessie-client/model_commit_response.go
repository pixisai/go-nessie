/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitResponse{}

// CommitResponse struct for CommitResponse
type CommitResponse struct {
	TargetBranch Branch3 `json:"targetBranch"`
	AddedContents []CommitResponseAddedContentsInner `json:"addedContents,omitempty"`
}

type _CommitResponse CommitResponse

// NewCommitResponse instantiates a new CommitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitResponse(targetBranch Branch3) *CommitResponse {
	this := CommitResponse{}
	this.TargetBranch = targetBranch
	return &this
}

// NewCommitResponseWithDefaults instantiates a new CommitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitResponseWithDefaults() *CommitResponse {
	this := CommitResponse{}
	return &this
}

// GetTargetBranch returns the TargetBranch field value
func (o *CommitResponse) GetTargetBranch() Branch3 {
	if o == nil {
		var ret Branch3
		return ret
	}

	return o.TargetBranch
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetTargetBranchOk() (*Branch3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetBranch, true
}

// SetTargetBranch sets field value
func (o *CommitResponse) SetTargetBranch(v Branch3) {
	o.TargetBranch = v
}

// GetAddedContents returns the AddedContents field value if set, zero value otherwise.
func (o *CommitResponse) GetAddedContents() []CommitResponseAddedContentsInner {
	if o == nil || IsNil(o.AddedContents) {
		var ret []CommitResponseAddedContentsInner
		return ret
	}
	return o.AddedContents
}

// GetAddedContentsOk returns a tuple with the AddedContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitResponse) GetAddedContentsOk() ([]CommitResponseAddedContentsInner, bool) {
	if o == nil || IsNil(o.AddedContents) {
		return nil, false
	}
	return o.AddedContents, true
}

// HasAddedContents returns a boolean if a field has been set.
func (o *CommitResponse) HasAddedContents() bool {
	if o != nil && !IsNil(o.AddedContents) {
		return true
	}

	return false
}

// SetAddedContents gets a reference to the given []CommitResponseAddedContentsInner and assigns it to the AddedContents field.
func (o *CommitResponse) SetAddedContents(v []CommitResponseAddedContentsInner) {
	o.AddedContents = v
}

func (o CommitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetBranch"] = o.TargetBranch
	if !IsNil(o.AddedContents) {
		toSerialize["addedContents"] = o.AddedContents
	}
	return toSerialize, nil
}

func (o *CommitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetBranch",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommitResponse := _CommitResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommitResponse)

	if err != nil {
		return err
	}

	*o = CommitResponse(varCommitResponse)

	return err
}

type NullableCommitResponse struct {
	value *CommitResponse
	isSet bool
}

func (v NullableCommitResponse) Get() *CommitResponse {
	return v.value
}

func (v *NullableCommitResponse) Set(val *CommitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitResponse(val *CommitResponse) *NullableCommitResponse {
	return &NullableCommitResponse{value: val, isSet: true}
}

func (v NullableCommitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


