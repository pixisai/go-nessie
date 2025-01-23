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

// checks if the ContentKeyV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKeyV2{}

// ContentKeyV2 struct for ContentKeyV2
type ContentKeyV2 struct {
	Elements []string `json:"elements"`
}

type _ContentKeyV2 ContentKeyV2

// NewContentKeyV2 instantiates a new ContentKeyV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKeyV2(elements []string) *ContentKeyV2 {
	this := ContentKeyV2{}
	this.Elements = elements
	return &this
}

// NewContentKeyV2WithDefaults instantiates a new ContentKeyV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKeyV2WithDefaults() *ContentKeyV2 {
	this := ContentKeyV2{}
	return &this
}

// GetElements returns the Elements field value
func (o *ContentKeyV2) GetElements() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value
// and a boolean to check if the value has been set.
func (o *ContentKeyV2) GetElementsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elements, true
}

// SetElements sets field value
func (o *ContentKeyV2) SetElements(v []string) {
	o.Elements = v
}

func (o ContentKeyV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKeyV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["elements"] = o.Elements
	return toSerialize, nil
}

func (o *ContentKeyV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"elements",
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

	varContentKeyV2 := _ContentKeyV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentKeyV2)

	if err != nil {
		return err
	}

	*o = ContentKeyV2(varContentKeyV2)

	return err
}

type NullableContentKeyV2 struct {
	value *ContentKeyV2
	isSet bool
}

func (v NullableContentKeyV2) Get() *ContentKeyV2 {
	return v.value
}

func (v *NullableContentKeyV2) Set(val *ContentKeyV2) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKeyV2) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKeyV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKeyV2(val *ContentKeyV2) *NullableContentKeyV2 {
	return &NullableContentKeyV2{value: val, isSet: true}
}

func (v NullableContentKeyV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKeyV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


