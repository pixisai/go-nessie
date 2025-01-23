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

// checks if the DeltaLakeTableV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeltaLakeTableV2{}

// DeltaLakeTableV2 struct for DeltaLakeTableV2
type DeltaLakeTableV2 struct {
	Id *string `json:"id,omitempty"`
	MetadataLocationHistory []string `json:"metadataLocationHistory"`
	CheckpointLocationHistory []string `json:"checkpointLocationHistory"`
	LastCheckpoint *string `json:"lastCheckpoint,omitempty"`
}

type _DeltaLakeTableV2 DeltaLakeTableV2

// NewDeltaLakeTableV2 instantiates a new DeltaLakeTableV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeltaLakeTableV2(metadataLocationHistory []string, checkpointLocationHistory []string) *DeltaLakeTableV2 {
	this := DeltaLakeTableV2{}
	this.MetadataLocationHistory = metadataLocationHistory
	this.CheckpointLocationHistory = checkpointLocationHistory
	return &this
}

// NewDeltaLakeTableV2WithDefaults instantiates a new DeltaLakeTableV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeltaLakeTableV2WithDefaults() *DeltaLakeTableV2 {
	this := DeltaLakeTableV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeltaLakeTableV2) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeltaLakeTableV2) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeltaLakeTableV2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeltaLakeTableV2) SetId(v string) {
	o.Id = &v
}

// GetMetadataLocationHistory returns the MetadataLocationHistory field value
func (o *DeltaLakeTableV2) GetMetadataLocationHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MetadataLocationHistory
}

// GetMetadataLocationHistoryOk returns a tuple with the MetadataLocationHistory field value
// and a boolean to check if the value has been set.
func (o *DeltaLakeTableV2) GetMetadataLocationHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataLocationHistory, true
}

// SetMetadataLocationHistory sets field value
func (o *DeltaLakeTableV2) SetMetadataLocationHistory(v []string) {
	o.MetadataLocationHistory = v
}

// GetCheckpointLocationHistory returns the CheckpointLocationHistory field value
func (o *DeltaLakeTableV2) GetCheckpointLocationHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CheckpointLocationHistory
}

// GetCheckpointLocationHistoryOk returns a tuple with the CheckpointLocationHistory field value
// and a boolean to check if the value has been set.
func (o *DeltaLakeTableV2) GetCheckpointLocationHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckpointLocationHistory, true
}

// SetCheckpointLocationHistory sets field value
func (o *DeltaLakeTableV2) SetCheckpointLocationHistory(v []string) {
	o.CheckpointLocationHistory = v
}

// GetLastCheckpoint returns the LastCheckpoint field value if set, zero value otherwise.
func (o *DeltaLakeTableV2) GetLastCheckpoint() string {
	if o == nil || IsNil(o.LastCheckpoint) {
		var ret string
		return ret
	}
	return *o.LastCheckpoint
}

// GetLastCheckpointOk returns a tuple with the LastCheckpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeltaLakeTableV2) GetLastCheckpointOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckpoint) {
		return nil, false
	}
	return o.LastCheckpoint, true
}

// HasLastCheckpoint returns a boolean if a field has been set.
func (o *DeltaLakeTableV2) HasLastCheckpoint() bool {
	if o != nil && !IsNil(o.LastCheckpoint) {
		return true
	}

	return false
}

// SetLastCheckpoint gets a reference to the given string and assigns it to the LastCheckpoint field.
func (o *DeltaLakeTableV2) SetLastCheckpoint(v string) {
	o.LastCheckpoint = &v
}

func (o DeltaLakeTableV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeltaLakeTableV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["metadataLocationHistory"] = o.MetadataLocationHistory
	toSerialize["checkpointLocationHistory"] = o.CheckpointLocationHistory
	if !IsNil(o.LastCheckpoint) {
		toSerialize["lastCheckpoint"] = o.LastCheckpoint
	}
	return toSerialize, nil
}

func (o *DeltaLakeTableV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataLocationHistory",
		"checkpointLocationHistory",
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

	varDeltaLakeTableV2 := _DeltaLakeTableV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeltaLakeTableV2)

	if err != nil {
		return err
	}

	*o = DeltaLakeTableV2(varDeltaLakeTableV2)

	return err
}

type NullableDeltaLakeTableV2 struct {
	value *DeltaLakeTableV2
	isSet bool
}

func (v NullableDeltaLakeTableV2) Get() *DeltaLakeTableV2 {
	return v.value
}

func (v *NullableDeltaLakeTableV2) Set(val *DeltaLakeTableV2) {
	v.value = val
	v.isSet = true
}

func (v NullableDeltaLakeTableV2) IsSet() bool {
	return v.isSet
}

func (v *NullableDeltaLakeTableV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeltaLakeTableV2(val *DeltaLakeTableV2) *NullableDeltaLakeTableV2 {
	return &NullableDeltaLakeTableV2{value: val, isSet: true}
}

func (v NullableDeltaLakeTableV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeltaLakeTableV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


