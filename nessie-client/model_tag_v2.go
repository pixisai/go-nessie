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

// checks if the TagV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagV2{}

// TagV2 struct for TagV2
type TagV2 struct {
	Name string `json:"name" validate:"regexp=^(?:[A-Za-z](?:(?:(?![.][.])[A-Za-z0-9.\\/_-])*[A-Za-z0-9_-])?)|-$"`
	Metadata *ReferenceMetadata3 `json:"metadata,omitempty"`
	Hash *string `json:"hash,omitempty" validate:"regexp=^([0-9a-fA-F]{8,64})?((?:([~*^])([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}[.][0-9]{1,9}Z|([0-9]+)))*)$"`
}

type _TagV2 TagV2

// NewTagV2 instantiates a new TagV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagV2(name string) *TagV2 {
	this := TagV2{}
	this.Name = name
	return &this
}

// NewTagV2WithDefaults instantiates a new TagV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagV2WithDefaults() *TagV2 {
	this := TagV2{}
	return &this
}

// GetName returns the Name field value
func (o *TagV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagV2) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TagV2) GetMetadata() ReferenceMetadata3 {
	if o == nil || IsNil(o.Metadata) {
		var ret ReferenceMetadata3
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagV2) GetMetadataOk() (*ReferenceMetadata3, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TagV2) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ReferenceMetadata3 and assigns it to the Metadata field.
func (o *TagV2) SetMetadata(v ReferenceMetadata3) {
	o.Metadata = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *TagV2) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagV2) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *TagV2) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *TagV2) SetHash(v string) {
	o.Hash = &v
}

func (o TagV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

func (o *TagV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTagV2 := _TagV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagV2)

	if err != nil {
		return err
	}

	*o = TagV2(varTagV2)

	return err
}

type NullableTagV2 struct {
	value *TagV2
	isSet bool
}

func (v NullableTagV2) Get() *TagV2 {
	return v.value
}

func (v *NullableTagV2) Set(val *TagV2) {
	v.value = val
	v.isSet = true
}

func (v NullableTagV2) IsSet() bool {
	return v.isSet
}

func (v *NullableTagV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagV2(val *TagV2) *NullableTagV2 {
	return &NullableTagV2{value: val, isSet: true}
}

func (v NullableTagV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


