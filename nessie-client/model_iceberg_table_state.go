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

// checks if the IcebergTableState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IcebergTableState{}

// IcebergTableState Represents the state of an Iceberg table in Nessie. An Iceberg table is globally identified via its unique 'Content.id'.  A Nessie commit-operation, performed via 'TreeApi.commitMultipleOperations',for Iceberg consists of a 'Operation.Put' with an 'IcebergTable' as in the 'content' field and the previous value of 'IcebergTable' in the 'expectedContent' field.
type IcebergTableState struct {
	Id *string `json:"id,omitempty"`
	MetadataLocation string `json:"metadataLocation" validate:"regexp=\\\\S"`
	SnapshotId *int64 `json:"snapshotId,omitempty"`
	SchemaId *int32 `json:"schemaId,omitempty"`
	SpecId *int32 `json:"specId,omitempty"`
	SortOrderId *int32 `json:"sortOrderId,omitempty"`
	// Deprecated
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _IcebergTableState IcebergTableState

// NewIcebergTableState instantiates a new IcebergTableState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcebergTableState(metadataLocation string) *IcebergTableState {
	this := IcebergTableState{}
	this.MetadataLocation = metadataLocation
	return &this
}

// NewIcebergTableStateWithDefaults instantiates a new IcebergTableState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIcebergTableStateWithDefaults() *IcebergTableState {
	this := IcebergTableState{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IcebergTableState) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IcebergTableState) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IcebergTableState) SetId(v string) {
	o.Id = &v
}

// GetMetadataLocation returns the MetadataLocation field value
func (o *IcebergTableState) GetMetadataLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataLocation
}

// GetMetadataLocationOk returns a tuple with the MetadataLocation field value
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetMetadataLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataLocation, true
}

// SetMetadataLocation sets field value
func (o *IcebergTableState) SetMetadataLocation(v string) {
	o.MetadataLocation = v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *IcebergTableState) GetSnapshotId() int64 {
	if o == nil || IsNil(o.SnapshotId) {
		var ret int64
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetSnapshotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *IcebergTableState) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given int64 and assigns it to the SnapshotId field.
func (o *IcebergTableState) SetSnapshotId(v int64) {
	o.SnapshotId = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *IcebergTableState) GetSchemaId() int32 {
	if o == nil || IsNil(o.SchemaId) {
		var ret int32
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetSchemaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaId) {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *IcebergTableState) HasSchemaId() bool {
	if o != nil && !IsNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given int32 and assigns it to the SchemaId field.
func (o *IcebergTableState) SetSchemaId(v int32) {
	o.SchemaId = &v
}

// GetSpecId returns the SpecId field value if set, zero value otherwise.
func (o *IcebergTableState) GetSpecId() int32 {
	if o == nil || IsNil(o.SpecId) {
		var ret int32
		return ret
	}
	return *o.SpecId
}

// GetSpecIdOk returns a tuple with the SpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetSpecIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SpecId) {
		return nil, false
	}
	return o.SpecId, true
}

// HasSpecId returns a boolean if a field has been set.
func (o *IcebergTableState) HasSpecId() bool {
	if o != nil && !IsNil(o.SpecId) {
		return true
	}

	return false
}

// SetSpecId gets a reference to the given int32 and assigns it to the SpecId field.
func (o *IcebergTableState) SetSpecId(v int32) {
	o.SpecId = &v
}

// GetSortOrderId returns the SortOrderId field value if set, zero value otherwise.
func (o *IcebergTableState) GetSortOrderId() int32 {
	if o == nil || IsNil(o.SortOrderId) {
		var ret int32
		return ret
	}
	return *o.SortOrderId
}

// GetSortOrderIdOk returns a tuple with the SortOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergTableState) GetSortOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrderId) {
		return nil, false
	}
	return o.SortOrderId, true
}

// HasSortOrderId returns a boolean if a field has been set.
func (o *IcebergTableState) HasSortOrderId() bool {
	if o != nil && !IsNil(o.SortOrderId) {
		return true
	}

	return false
}

// SetSortOrderId gets a reference to the given int32 and assigns it to the SortOrderId field.
func (o *IcebergTableState) SetSortOrderId(v int32) {
	o.SortOrderId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
// Deprecated
func (o *IcebergTableState) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *IcebergTableState) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IcebergTableState) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
// Deprecated
func (o *IcebergTableState) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o IcebergTableState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IcebergTableState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["metadataLocation"] = o.MetadataLocation
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !IsNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if !IsNil(o.SpecId) {
		toSerialize["specId"] = o.SpecId
	}
	if !IsNil(o.SortOrderId) {
		toSerialize["sortOrderId"] = o.SortOrderId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *IcebergTableState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataLocation",
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

	varIcebergTableState := _IcebergTableState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIcebergTableState)

	if err != nil {
		return err
	}

	*o = IcebergTableState(varIcebergTableState)

	return err
}

type NullableIcebergTableState struct {
	value *IcebergTableState
	isSet bool
}

func (v NullableIcebergTableState) Get() *IcebergTableState {
	return v.value
}

func (v *NullableIcebergTableState) Set(val *IcebergTableState) {
	v.value = val
	v.isSet = true
}

func (v NullableIcebergTableState) IsSet() bool {
	return v.isSet
}

func (v *NullableIcebergTableState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcebergTableState(val *IcebergTableState) *NullableIcebergTableState {
	return &NullableIcebergTableState{value: val, isSet: true}
}

func (v NullableIcebergTableState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcebergTableState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


