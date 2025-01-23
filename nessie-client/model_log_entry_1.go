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

// checks if the LogEntry1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogEntry1{}

// LogEntry1 struct for LogEntry1
type LogEntry1 struct {
	CommitMeta CommitMeta3 `json:"commitMeta"`
	AdditionalParents []string `json:"additionalParents,omitempty"`
	ParentCommitHash *string `json:"parentCommitHash,omitempty"`
	Operations []Operation1 `json:"operations,omitempty"`
}

type _LogEntry1 LogEntry1

// NewLogEntry1 instantiates a new LogEntry1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEntry1(commitMeta CommitMeta3) *LogEntry1 {
	this := LogEntry1{}
	this.CommitMeta = commitMeta
	return &this
}

// NewLogEntry1WithDefaults instantiates a new LogEntry1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEntry1WithDefaults() *LogEntry1 {
	this := LogEntry1{}
	return &this
}

// GetCommitMeta returns the CommitMeta field value
func (o *LogEntry1) GetCommitMeta() CommitMeta3 {
	if o == nil {
		var ret CommitMeta3
		return ret
	}

	return o.CommitMeta
}

// GetCommitMetaOk returns a tuple with the CommitMeta field value
// and a boolean to check if the value has been set.
func (o *LogEntry1) GetCommitMetaOk() (*CommitMeta3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitMeta, true
}

// SetCommitMeta sets field value
func (o *LogEntry1) SetCommitMeta(v CommitMeta3) {
	o.CommitMeta = v
}

// GetAdditionalParents returns the AdditionalParents field value if set, zero value otherwise.
func (o *LogEntry1) GetAdditionalParents() []string {
	if o == nil || IsNil(o.AdditionalParents) {
		var ret []string
		return ret
	}
	return o.AdditionalParents
}

// GetAdditionalParentsOk returns a tuple with the AdditionalParents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry1) GetAdditionalParentsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalParents) {
		return nil, false
	}
	return o.AdditionalParents, true
}

// HasAdditionalParents returns a boolean if a field has been set.
func (o *LogEntry1) HasAdditionalParents() bool {
	if o != nil && !IsNil(o.AdditionalParents) {
		return true
	}

	return false
}

// SetAdditionalParents gets a reference to the given []string and assigns it to the AdditionalParents field.
func (o *LogEntry1) SetAdditionalParents(v []string) {
	o.AdditionalParents = v
}

// GetParentCommitHash returns the ParentCommitHash field value if set, zero value otherwise.
func (o *LogEntry1) GetParentCommitHash() string {
	if o == nil || IsNil(o.ParentCommitHash) {
		var ret string
		return ret
	}
	return *o.ParentCommitHash
}

// GetParentCommitHashOk returns a tuple with the ParentCommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry1) GetParentCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCommitHash) {
		return nil, false
	}
	return o.ParentCommitHash, true
}

// HasParentCommitHash returns a boolean if a field has been set.
func (o *LogEntry1) HasParentCommitHash() bool {
	if o != nil && !IsNil(o.ParentCommitHash) {
		return true
	}

	return false
}

// SetParentCommitHash gets a reference to the given string and assigns it to the ParentCommitHash field.
func (o *LogEntry1) SetParentCommitHash(v string) {
	o.ParentCommitHash = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *LogEntry1) GetOperations() []Operation1 {
	if o == nil || IsNil(o.Operations) {
		var ret []Operation1
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogEntry1) GetOperationsOk() ([]Operation1, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *LogEntry1) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []Operation1 and assigns it to the Operations field.
func (o *LogEntry1) SetOperations(v []Operation1) {
	o.Operations = v
}

func (o LogEntry1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogEntry1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commitMeta"] = o.CommitMeta
	if !IsNil(o.AdditionalParents) {
		toSerialize["additionalParents"] = o.AdditionalParents
	}
	if !IsNil(o.ParentCommitHash) {
		toSerialize["parentCommitHash"] = o.ParentCommitHash
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	return toSerialize, nil
}

func (o *LogEntry1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commitMeta",
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

	varLogEntry1 := _LogEntry1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogEntry1)

	if err != nil {
		return err
	}

	*o = LogEntry1(varLogEntry1)

	return err
}

type NullableLogEntry1 struct {
	value *LogEntry1
	isSet bool
}

func (v NullableLogEntry1) Get() *LogEntry1 {
	return v.value
}

func (v *NullableLogEntry1) Set(val *LogEntry1) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEntry1) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEntry1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEntry1(val *LogEntry1) *NullableLogEntry1 {
	return &NullableLogEntry1{value: val, isSet: true}
}

func (v NullableLogEntry1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEntry1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


