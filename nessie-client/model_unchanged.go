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

// checks if the Unchanged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Unchanged{}

// Unchanged struct for Unchanged
type Unchanged struct {
	Key GetMultipleContentsRequest1RequestedKeysInner `json:"key"`
}

type _Unchanged Unchanged

// NewUnchanged instantiates a new Unchanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnchanged(key GetMultipleContentsRequest1RequestedKeysInner) *Unchanged {
	this := Unchanged{}
	this.Key = key
	return &this
}

// NewUnchangedWithDefaults instantiates a new Unchanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnchangedWithDefaults() *Unchanged {
	this := Unchanged{}
	return &this
}

// GetKey returns the Key field value
func (o *Unchanged) GetKey() GetMultipleContentsRequest1RequestedKeysInner {
	if o == nil {
		var ret GetMultipleContentsRequest1RequestedKeysInner
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Unchanged) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Unchanged) SetKey(v GetMultipleContentsRequest1RequestedKeysInner) {
	o.Key = v
}

func (o Unchanged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Unchanged) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

func (o *Unchanged) UnmarshalJSON(data []byte) (err error) {
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

	varUnchanged := _Unchanged{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnchanged)

	if err != nil {
		return err
	}

	*o = Unchanged(varUnchanged)

	return err
}

type NullableUnchanged struct {
	value *Unchanged
	isSet bool
}

func (v NullableUnchanged) Get() *Unchanged {
	return v.value
}

func (v *NullableUnchanged) Set(val *Unchanged) {
	v.value = val
	v.isSet = true
}

func (v NullableUnchanged) IsSet() bool {
	return v.isSet
}

func (v *NullableUnchanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnchanged(val *Unchanged) *NullableUnchanged {
	return &NullableUnchanged{value: val, isSet: true}
}

func (v NullableUnchanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnchanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


