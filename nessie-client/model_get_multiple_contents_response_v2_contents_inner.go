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

// checks if the GetMultipleContentsResponseV2ContentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMultipleContentsResponseV2ContentsInner{}

// GetMultipleContentsResponseV2ContentsInner struct for GetMultipleContentsResponseV2ContentsInner
type GetMultipleContentsResponseV2ContentsInner struct {
	Key AddedContentKey `json:"key"`
	Content Content1 `json:"content"`
	Documentation *ContentResponseV2Documentation `json:"documentation,omitempty"`
}

type _GetMultipleContentsResponseV2ContentsInner GetMultipleContentsResponseV2ContentsInner

// NewGetMultipleContentsResponseV2ContentsInner instantiates a new GetMultipleContentsResponseV2ContentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMultipleContentsResponseV2ContentsInner(key AddedContentKey, content Content1) *GetMultipleContentsResponseV2ContentsInner {
	this := GetMultipleContentsResponseV2ContentsInner{}
	this.Key = key
	this.Content = content
	return &this
}

// NewGetMultipleContentsResponseV2ContentsInnerWithDefaults instantiates a new GetMultipleContentsResponseV2ContentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMultipleContentsResponseV2ContentsInnerWithDefaults() *GetMultipleContentsResponseV2ContentsInner {
	this := GetMultipleContentsResponseV2ContentsInner{}
	return &this
}

// GetKey returns the Key field value
func (o *GetMultipleContentsResponseV2ContentsInner) GetKey() AddedContentKey {
	if o == nil {
		var ret AddedContentKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetMultipleContentsResponseV2ContentsInner) GetKeyOk() (*AddedContentKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetMultipleContentsResponseV2ContentsInner) SetKey(v AddedContentKey) {
	o.Key = v
}

// GetContent returns the Content field value
func (o *GetMultipleContentsResponseV2ContentsInner) GetContent() Content1 {
	if o == nil {
		var ret Content1
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *GetMultipleContentsResponseV2ContentsInner) GetContentOk() (*Content1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *GetMultipleContentsResponseV2ContentsInner) SetContent(v Content1) {
	o.Content = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *GetMultipleContentsResponseV2ContentsInner) GetDocumentation() ContentResponseV2Documentation {
	if o == nil || IsNil(o.Documentation) {
		var ret ContentResponseV2Documentation
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleContentsResponseV2ContentsInner) GetDocumentationOk() (*ContentResponseV2Documentation, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *GetMultipleContentsResponseV2ContentsInner) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given ContentResponseV2Documentation and assigns it to the Documentation field.
func (o *GetMultipleContentsResponseV2ContentsInner) SetDocumentation(v ContentResponseV2Documentation) {
	o.Documentation = &v
}

func (o GetMultipleContentsResponseV2ContentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMultipleContentsResponseV2ContentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["content"] = o.Content
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	return toSerialize, nil
}

func (o *GetMultipleContentsResponseV2ContentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"content",
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

	varGetMultipleContentsResponseV2ContentsInner := _GetMultipleContentsResponseV2ContentsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMultipleContentsResponseV2ContentsInner)

	if err != nil {
		return err
	}

	*o = GetMultipleContentsResponseV2ContentsInner(varGetMultipleContentsResponseV2ContentsInner)

	return err
}

type NullableGetMultipleContentsResponseV2ContentsInner struct {
	value *GetMultipleContentsResponseV2ContentsInner
	isSet bool
}

func (v NullableGetMultipleContentsResponseV2ContentsInner) Get() *GetMultipleContentsResponseV2ContentsInner {
	return v.value
}

func (v *NullableGetMultipleContentsResponseV2ContentsInner) Set(val *GetMultipleContentsResponseV2ContentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMultipleContentsResponseV2ContentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMultipleContentsResponseV2ContentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMultipleContentsResponseV2ContentsInner(val *GetMultipleContentsResponseV2ContentsInner) *NullableGetMultipleContentsResponseV2ContentsInner {
	return &NullableGetMultipleContentsResponseV2ContentsInner{value: val, isSet: true}
}

func (v NullableGetMultipleContentsResponseV2ContentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMultipleContentsResponseV2ContentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


