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

// checks if the ContentMetadataV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentMetadataV2{}

// ContentMetadataV2 struct for ContentMetadataV2
type ContentMetadataV2 struct {
	Variant string `json:"variant"`
}

type _ContentMetadataV2 ContentMetadataV2

// NewContentMetadataV2 instantiates a new ContentMetadataV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentMetadataV2(variant string) *ContentMetadataV2 {
	this := ContentMetadataV2{}
	this.Variant = variant
	return &this
}

// NewContentMetadataV2WithDefaults instantiates a new ContentMetadataV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentMetadataV2WithDefaults() *ContentMetadataV2 {
	this := ContentMetadataV2{}
	return &this
}

// GetVariant returns the Variant field value
func (o *ContentMetadataV2) GetVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ContentMetadataV2) GetVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ContentMetadataV2) SetVariant(v string) {
	o.Variant = v
}

func (o ContentMetadataV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentMetadataV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant"] = o.Variant
	return toSerialize, nil
}

func (o *ContentMetadataV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"variant",
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

	varContentMetadataV2 := _ContentMetadataV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContentMetadataV2)

	if err != nil {
		return err
	}

	*o = ContentMetadataV2(varContentMetadataV2)

	return err
}

type NullableContentMetadataV2 struct {
	value *ContentMetadataV2
	isSet bool
}

func (v NullableContentMetadataV2) Get() *ContentMetadataV2 {
	return v.value
}

func (v *NullableContentMetadataV2) Set(val *ContentMetadataV2) {
	v.value = val
	v.isSet = true
}

func (v NullableContentMetadataV2) IsSet() bool {
	return v.isSet
}

func (v *NullableContentMetadataV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentMetadataV2(val *ContentMetadataV2) *NullableContentMetadataV2 {
	return &NullableContentMetadataV2{value: val, isSet: true}
}

func (v NullableContentMetadataV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentMetadataV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


