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

// checks if the PutContentOperationForAContentKey4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutContentOperationForAContentKey4{}

// PutContentOperationForAContentKey4 Used to add new content or to update existing content.  A new content object is created by populating the `value` field, the content-id in the content object must not be present (null).  A content object is updated by populating the `value` containing the correct content-id.  If the key for a content shall change (aka a rename), then use a `Delete` operation using the current (old) key and a `Put` operation using the new key with the `value` having the correct content-id. Both operations must happen in the same commit.  A content object can be replaced (think: `DROP TABLE xyz` + `CREATE TABLE xyz`) with a `Delete` operation and a `Put` operation for a content using a `value`representing a new content object, so without a content-id, in the same commit.
type PutContentOperationForAContentKey4 struct {
	Key GetMultipleContentsRequest1RequestedKeysInner `json:"key"`
	Content Content8 `json:"content"`
	// Deprecated
	ExpectedContent *Content9 `json:"expectedContent,omitempty"`
}

type _PutContentOperationForAContentKey4 PutContentOperationForAContentKey4

// NewPutContentOperationForAContentKey4 instantiates a new PutContentOperationForAContentKey4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutContentOperationForAContentKey4(key GetMultipleContentsRequest1RequestedKeysInner, content Content8) *PutContentOperationForAContentKey4 {
	this := PutContentOperationForAContentKey4{}
	this.Key = key
	this.Content = content
	return &this
}

// NewPutContentOperationForAContentKey4WithDefaults instantiates a new PutContentOperationForAContentKey4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutContentOperationForAContentKey4WithDefaults() *PutContentOperationForAContentKey4 {
	this := PutContentOperationForAContentKey4{}
	return &this
}

// GetKey returns the Key field value
func (o *PutContentOperationForAContentKey4) GetKey() GetMultipleContentsRequest1RequestedKeysInner {
	if o == nil {
		var ret GetMultipleContentsRequest1RequestedKeysInner
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PutContentOperationForAContentKey4) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PutContentOperationForAContentKey4) SetKey(v GetMultipleContentsRequest1RequestedKeysInner) {
	o.Key = v
}

// GetContent returns the Content field value
func (o *PutContentOperationForAContentKey4) GetContent() Content8 {
	if o == nil {
		var ret Content8
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PutContentOperationForAContentKey4) GetContentOk() (*Content8, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PutContentOperationForAContentKey4) SetContent(v Content8) {
	o.Content = v
}

// GetExpectedContent returns the ExpectedContent field value if set, zero value otherwise.
// Deprecated
func (o *PutContentOperationForAContentKey4) GetExpectedContent() Content9 {
	if o == nil || IsNil(o.ExpectedContent) {
		var ret Content9
		return ret
	}
	return *o.ExpectedContent
}

// GetExpectedContentOk returns a tuple with the ExpectedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PutContentOperationForAContentKey4) GetExpectedContentOk() (*Content9, bool) {
	if o == nil || IsNil(o.ExpectedContent) {
		return nil, false
	}
	return o.ExpectedContent, true
}

// HasExpectedContent returns a boolean if a field has been set.
func (o *PutContentOperationForAContentKey4) HasExpectedContent() bool {
	if o != nil && !IsNil(o.ExpectedContent) {
		return true
	}

	return false
}

// SetExpectedContent gets a reference to the given Content9 and assigns it to the ExpectedContent field.
// Deprecated
func (o *PutContentOperationForAContentKey4) SetExpectedContent(v Content9) {
	o.ExpectedContent = &v
}

func (o PutContentOperationForAContentKey4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutContentOperationForAContentKey4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["content"] = o.Content
	if !IsNil(o.ExpectedContent) {
		toSerialize["expectedContent"] = o.ExpectedContent
	}
	return toSerialize, nil
}

func (o *PutContentOperationForAContentKey4) UnmarshalJSON(data []byte) (err error) {
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

	varPutContentOperationForAContentKey4 := _PutContentOperationForAContentKey4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutContentOperationForAContentKey4)

	if err != nil {
		return err
	}

	*o = PutContentOperationForAContentKey4(varPutContentOperationForAContentKey4)

	return err
}

type NullablePutContentOperationForAContentKey4 struct {
	value *PutContentOperationForAContentKey4
	isSet bool
}

func (v NullablePutContentOperationForAContentKey4) Get() *PutContentOperationForAContentKey4 {
	return v.value
}

func (v *NullablePutContentOperationForAContentKey4) Set(val *PutContentOperationForAContentKey4) {
	v.value = val
	v.isSet = true
}

func (v NullablePutContentOperationForAContentKey4) IsSet() bool {
	return v.isSet
}

func (v *NullablePutContentOperationForAContentKey4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutContentOperationForAContentKey4(val *PutContentOperationForAContentKey4) *NullablePutContentOperationForAContentKey4 {
	return &NullablePutContentOperationForAContentKey4{value: val, isSet: true}
}

func (v NullablePutContentOperationForAContentKey4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutContentOperationForAContentKey4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


