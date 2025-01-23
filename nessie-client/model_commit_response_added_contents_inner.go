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

// checks if the CommitResponseAddedContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitResponseAddedContentsInner{}

// CommitResponseAddedContentsInner struct for CommitResponseAddedContentsInner
type CommitResponseAddedContentsInner struct {
	Key AddedContentKey `json:"key"`
}

type _CommitResponseAddedContentsInner CommitResponseAddedContentsInner

// NewCommitResponseAddedContentsInner instantiates a new CommitResponseAddedContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitResponseAddedContentsInner(key AddedContentKey) *CommitResponseAddedContentsInner {
	this := CommitResponseAddedContentsInner{}
	this.Key = key
	return &this
}

// NewCommitResponseAddedContentsInnerWithDefaults instantiates a new CommitResponseAddedContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitResponseAddedContentsInnerWithDefaults() *CommitResponseAddedContentsInner {
	this := CommitResponseAddedContentsInner{}
	return &this
}

// GetKey returns the Key field value
func (o *CommitResponseAddedContentsInner) GetKey() AddedContentKey {
	if o == nil {
		var ret AddedContentKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CommitResponseAddedContentsInner) GetKeyOk() (*AddedContentKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CommitResponseAddedContentsInner) SetKey(v AddedContentKey) {
	o.Key = v
}

func (o CommitResponseAddedContentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitResponseAddedContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *CommitResponseAddedContentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varCommitResponseAddedContentsInner := _CommitResponseAddedContentsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommitResponseAddedContentsInner)

	if err != nil {
		return err
	}

	*o = CommitResponseAddedContentsInner(varCommitResponseAddedContentsInner)

	return err
}

type NullableCommitResponseAddedContentsInner struct {
	value *CommitResponseAddedContentsInner
	isSet bool
}

func (v NullableCommitResponseAddedContentsInner) Get() *CommitResponseAddedContentsInner {
	return v.value
}

func (v *NullableCommitResponseAddedContentsInner) Set(val *CommitResponseAddedContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitResponseAddedContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitResponseAddedContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitResponseAddedContentsInner(val *CommitResponseAddedContentsInner) *NullableCommitResponseAddedContentsInner {
	return &NullableCommitResponseAddedContentsInner{value: val, isSet: true}
}

func (v NullableCommitResponseAddedContentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitResponseAddedContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


