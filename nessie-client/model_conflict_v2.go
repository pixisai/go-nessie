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

// checks if the ConflictV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictV2{}

// ConflictV2 struct for ConflictV2
type ConflictV2 struct {
	ConflictType interface{} `json:"conflictType,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// NewConflictV2 instantiates a new ConflictV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictV2() *ConflictV2 {
	this := ConflictV2{}
	return &this
}

// NewConflictV2WithDefaults instantiates a new ConflictV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictV2WithDefaults() *ConflictV2 {
	this := ConflictV2{}
	return &this
}

// GetConflictType returns the ConflictType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConflictV2) GetConflictType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ConflictType
}

// GetConflictTypeOk returns a tuple with the ConflictType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConflictV2) GetConflictTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ConflictType) {
		return nil, false
	}
	return &o.ConflictType, true
}

// HasConflictType returns a boolean if a field has been set.
func (o *ConflictV2) HasConflictType() bool {
	if o != nil && !IsNil(o.ConflictType) {
		return true
	}

	return false
}

// SetConflictType gets a reference to the given interface{} and assigns it to the ConflictType field.
func (o *ConflictV2) SetConflictType(v interface{}) {
	o.ConflictType = v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConflictV2) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConflictV2) GetKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConflictV2) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *ConflictV2) SetKey(v interface{}) {
	o.Key = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConflictV2) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConflictV2) GetMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConflictV2) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *ConflictV2) SetMessage(v interface{}) {
	o.Message = v
}

func (o ConflictV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConflictV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConflictType != nil {
		toSerialize["conflictType"] = o.ConflictType
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableConflictV2 struct {
	value *ConflictV2
	isSet bool
}

func (v NullableConflictV2) Get() *ConflictV2 {
	return v.value
}

func (v *NullableConflictV2) Set(val *ConflictV2) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictV2) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictV2(val *ConflictV2) *NullableConflictV2 {
	return &NullableConflictV2{value: val, isSet: true}
}

func (v NullableConflictV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


