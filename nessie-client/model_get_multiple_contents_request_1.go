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

// checks if the GetMultipleContentsRequest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMultipleContentsRequest1{}

// GetMultipleContentsRequest1 struct for GetMultipleContentsRequest1
type GetMultipleContentsRequest1 struct {
	RequestedKeys []GetMultipleContentsRequest1RequestedKeysInner `json:"requestedKeys"`
}

type _GetMultipleContentsRequest1 GetMultipleContentsRequest1

// NewGetMultipleContentsRequest1 instantiates a new GetMultipleContentsRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMultipleContentsRequest1(requestedKeys []GetMultipleContentsRequest1RequestedKeysInner) *GetMultipleContentsRequest1 {
	this := GetMultipleContentsRequest1{}
	this.RequestedKeys = requestedKeys
	return &this
}

// NewGetMultipleContentsRequest1WithDefaults instantiates a new GetMultipleContentsRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMultipleContentsRequest1WithDefaults() *GetMultipleContentsRequest1 {
	this := GetMultipleContentsRequest1{}
	return &this
}

// GetRequestedKeys returns the RequestedKeys field value
func (o *GetMultipleContentsRequest1) GetRequestedKeys() []GetMultipleContentsRequest1RequestedKeysInner {
	if o == nil {
		var ret []GetMultipleContentsRequest1RequestedKeysInner
		return ret
	}

	return o.RequestedKeys
}

// GetRequestedKeysOk returns a tuple with the RequestedKeys field value
// and a boolean to check if the value has been set.
func (o *GetMultipleContentsRequest1) GetRequestedKeysOk() ([]GetMultipleContentsRequest1RequestedKeysInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedKeys, true
}

// SetRequestedKeys sets field value
func (o *GetMultipleContentsRequest1) SetRequestedKeys(v []GetMultipleContentsRequest1RequestedKeysInner) {
	o.RequestedKeys = v
}

func (o GetMultipleContentsRequest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMultipleContentsRequest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestedKeys"] = o.RequestedKeys
	return toSerialize, nil
}

func (o *GetMultipleContentsRequest1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestedKeys",
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

	varGetMultipleContentsRequest1 := _GetMultipleContentsRequest1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMultipleContentsRequest1)

	if err != nil {
		return err
	}

	*o = GetMultipleContentsRequest1(varGetMultipleContentsRequest1)

	return err
}

type NullableGetMultipleContentsRequest1 struct {
	value *GetMultipleContentsRequest1
	isSet bool
}

func (v NullableGetMultipleContentsRequest1) Get() *GetMultipleContentsRequest1 {
	return v.value
}

func (v *NullableGetMultipleContentsRequest1) Set(val *GetMultipleContentsRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultipleContentsRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultipleContentsRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultipleContentsRequest1(val *GetMultipleContentsRequest1) *NullableGetMultipleContentsRequest1 {
	return &NullableGetMultipleContentsRequest1{value: val, isSet: true}
}

func (v NullableGetMultipleContentsRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMultipleContentsRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


