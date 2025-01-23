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

// checks if the DeleteV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteV2{}

// DeleteV2 Used to delete an existing content key.  If the key for a content shall change (aka a rename), then use a `Delete` operation using the current (old) key and a `Put` operation using the new key with the current `Content` in the the `value` field. See `Put` operation.
type DeleteV2 struct {
	Key AddedContentKey `json:"key"`
}

type _DeleteV2 DeleteV2

// NewDeleteV2 instantiates a new DeleteV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteV2(key AddedContentKey) *DeleteV2 {
	this := DeleteV2{}
	this.Key = key
	return &this
}

// NewDeleteV2WithDefaults instantiates a new DeleteV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteV2WithDefaults() *DeleteV2 {
	this := DeleteV2{}
	return &this
}

// GetKey returns the Key field value
func (o *DeleteV2) GetKey() AddedContentKey {
	if o == nil {
		var ret AddedContentKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeleteV2) GetKeyOk() (*AddedContentKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeleteV2) SetKey(v AddedContentKey) {
	o.Key = v
}

func (o DeleteV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *DeleteV2) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteV2 := _DeleteV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteV2)

	if err != nil {
		return err
	}

	*o = DeleteV2(varDeleteV2)

	return err
}

type NullableDeleteV2 struct {
	value *DeleteV2
	isSet bool
}

func (v NullableDeleteV2) Get() *DeleteV2 {
	return v.value
}

func (v *NullableDeleteV2) Set(val *DeleteV2) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteV2) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteV2(val *DeleteV2) *NullableDeleteV2 {
	return &NullableDeleteV2{value: val, isSet: true}
}

func (v NullableDeleteV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


