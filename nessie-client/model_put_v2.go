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

// checks if the PutV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutV2{}

// PutV2 Used to add new content or to update existing content.  A new content object is created by populating the `value` field, the content-id in the content object must not be present (null).  A content object is updated by populating the `value` containing the correct content-id.  If the key for a content shall change (aka a rename), then use a `Delete` operation using the current (old) key and a `Put` operation using the new key with the `value` having the correct content-id. Both operations must happen in the same commit.  A content object can be replaced (think: `DROP TABLE xyz` + `CREATE TABLE xyz`) with a `Delete` operation and a `Put` operation for a content using a `value`representing a new content object, so without a content-id, in the same commit.
type PutV2 struct {
	Key GetMultipleContentsRequest1RequestedKeysInner `json:"key"`
	Content Content5 `json:"content"`
	Metadata []ContentMetadata1 `json:"metadata,omitempty"`
	Documentation *PutContentOperationForAContentKeyDocumentation `json:"documentation,omitempty"`
}

type _PutV2 PutV2

// NewPutV2 instantiates a new PutV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutV2(key GetMultipleContentsRequest1RequestedKeysInner, content Content5) *PutV2 {
	this := PutV2{}
	this.Key = key
	this.Content = content
	return &this
}

// NewPutV2WithDefaults instantiates a new PutV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutV2WithDefaults() *PutV2 {
	this := PutV2{}
	return &this
}

// GetKey returns the Key field value
func (o *PutV2) GetKey() GetMultipleContentsRequest1RequestedKeysInner {
	if o == nil {
		var ret GetMultipleContentsRequest1RequestedKeysInner
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PutV2) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PutV2) SetKey(v GetMultipleContentsRequest1RequestedKeysInner) {
	o.Key = v
}

// GetContent returns the Content field value
func (o *PutV2) GetContent() Content5 {
	if o == nil {
		var ret Content5
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *PutV2) GetContentOk() (*Content5, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *PutV2) SetContent(v Content5) {
	o.Content = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PutV2) GetMetadata() []ContentMetadata1 {
	if o == nil || IsNil(o.Metadata) {
		var ret []ContentMetadata1
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutV2) GetMetadataOk() ([]ContentMetadata1, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PutV2) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ContentMetadata1 and assigns it to the Metadata field.
func (o *PutV2) SetMetadata(v []ContentMetadata1) {
	o.Metadata = v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *PutV2) GetDocumentation() PutContentOperationForAContentKeyDocumentation {
	if o == nil || IsNil(o.Documentation) {
		var ret PutContentOperationForAContentKeyDocumentation
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutV2) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool) {
	if o == nil || IsNil(o.Documentation) {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *PutV2) HasDocumentation() bool {
	if o != nil && !IsNil(o.Documentation) {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given PutContentOperationForAContentKeyDocumentation and assigns it to the Documentation field.
func (o *PutV2) SetDocumentation(v PutContentOperationForAContentKeyDocumentation) {
	o.Documentation = &v
}

func (o PutV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["content"] = o.Content
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Documentation) {
		toSerialize["documentation"] = o.Documentation
	}
	return toSerialize, nil
}

func (o *PutV2) UnmarshalJSON(data []byte) (err error) {
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

	varPutV2 := _PutV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutV2)

	if err != nil {
		return err
	}

	*o = PutV2(varPutV2)

	return err
}

type NullablePutV2 struct {
	value *PutV2
	isSet bool
}

func (v NullablePutV2) Get() *PutV2 {
	return v.value
}

func (v *NullablePutV2) Set(val *PutV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePutV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePutV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutV2(val *PutV2) *NullablePutV2 {
	return &NullablePutV2{value: val, isSet: true}
}

func (v NullablePutV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


