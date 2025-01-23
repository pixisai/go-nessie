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

// checks if the Operations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Operations{}

// Operations struct for Operations
type Operations struct {
	CommitMeta CommitMeta2 `json:"commitMeta"`
	Operations []Operation1 `json:"operations"`
}

type _Operations Operations

// NewOperations instantiates a new Operations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperations(commitMeta CommitMeta2, operations []Operation1) *Operations {
	this := Operations{}
	this.CommitMeta = commitMeta
	this.Operations = operations
	return &this
}

// NewOperationsWithDefaults instantiates a new Operations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationsWithDefaults() *Operations {
	this := Operations{}
	return &this
}

// GetCommitMeta returns the CommitMeta field value
func (o *Operations) GetCommitMeta() CommitMeta2 {
	if o == nil {
		var ret CommitMeta2
		return ret
	}

	return o.CommitMeta
}

// GetCommitMetaOk returns a tuple with the CommitMeta field value
// and a boolean to check if the value has been set.
func (o *Operations) GetCommitMetaOk() (*CommitMeta2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitMeta, true
}

// SetCommitMeta sets field value
func (o *Operations) SetCommitMeta(v CommitMeta2) {
	o.CommitMeta = v
}

// GetOperations returns the Operations field value
func (o *Operations) GetOperations() []Operation1 {
	if o == nil {
		var ret []Operation1
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *Operations) GetOperationsOk() ([]Operation1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *Operations) SetOperations(v []Operation1) {
	o.Operations = v
}

func (o Operations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Operations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commitMeta"] = o.CommitMeta
	toSerialize["operations"] = o.Operations
	return toSerialize, nil
}

func (o *Operations) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commitMeta",
		"operations",
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

	varOperations := _Operations{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperations)

	if err != nil {
		return err
	}

	*o = Operations(varOperations)

	return err
}

type NullableOperations struct {
	value *Operations
	isSet bool
}

func (v NullableOperations) Get() *Operations {
	return v.value
}

func (v *NullableOperations) Set(val *Operations) {
	v.value = val
	v.isSet = true
}

func (v NullableOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperations(val *Operations) *NullableOperations {
	return &NullableOperations{value: val, isSet: true}
}

func (v NullableOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


