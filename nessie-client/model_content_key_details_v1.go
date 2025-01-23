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

// checks if the ContentKeyDetailsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKeyDetailsV1{}

// ContentKeyDetailsV1 struct for ContentKeyDetailsV1
type ContentKeyDetailsV1 struct {
	Key *GetMultipleContentsRequest1RequestedKeysInner `json:"key,omitempty"`
	MergeBehavior *string `json:"mergeBehavior,omitempty"`
}

// NewContentKeyDetailsV1 instantiates a new ContentKeyDetailsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKeyDetailsV1() *ContentKeyDetailsV1 {
	this := ContentKeyDetailsV1{}
	return &this
}

// NewContentKeyDetailsV1WithDefaults instantiates a new ContentKeyDetailsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKeyDetailsV1WithDefaults() *ContentKeyDetailsV1 {
	this := ContentKeyDetailsV1{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ContentKeyDetailsV1) GetKey() GetMultipleContentsRequest1RequestedKeysInner {
	if o == nil || IsNil(o.Key) {
		var ret GetMultipleContentsRequest1RequestedKeysInner
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKeyDetailsV1) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ContentKeyDetailsV1) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given GetMultipleContentsRequest1RequestedKeysInner and assigns it to the Key field.
func (o *ContentKeyDetailsV1) SetKey(v GetMultipleContentsRequest1RequestedKeysInner) {
	o.Key = &v
}

// GetMergeBehavior returns the MergeBehavior field value if set, zero value otherwise.
func (o *ContentKeyDetailsV1) GetMergeBehavior() string {
	if o == nil || IsNil(o.MergeBehavior) {
		var ret string
		return ret
	}
	return *o.MergeBehavior
}

// GetMergeBehaviorOk returns a tuple with the MergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKeyDetailsV1) GetMergeBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.MergeBehavior) {
		return nil, false
	}
	return o.MergeBehavior, true
}

// HasMergeBehavior returns a boolean if a field has been set.
func (o *ContentKeyDetailsV1) HasMergeBehavior() bool {
	if o != nil && !IsNil(o.MergeBehavior) {
		return true
	}

	return false
}

// SetMergeBehavior gets a reference to the given string and assigns it to the MergeBehavior field.
func (o *ContentKeyDetailsV1) SetMergeBehavior(v string) {
	o.MergeBehavior = &v
}

func (o ContentKeyDetailsV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKeyDetailsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.MergeBehavior) {
		toSerialize["mergeBehavior"] = o.MergeBehavior
	}
	return toSerialize, nil
}

type NullableContentKeyDetailsV1 struct {
	value *ContentKeyDetailsV1
	isSet bool
}

func (v NullableContentKeyDetailsV1) Get() *ContentKeyDetailsV1 {
	return v.value
}

func (v *NullableContentKeyDetailsV1) Set(val *ContentKeyDetailsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKeyDetailsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKeyDetailsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKeyDetailsV1(val *ContentKeyDetailsV1) *NullableContentKeyDetailsV1 {
	return &NullableContentKeyDetailsV1{value: val, isSet: true}
}

func (v NullableContentKeyDetailsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKeyDetailsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


