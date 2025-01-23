/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CommitConsistencyV2 the model 'CommitConsistencyV2'
type CommitConsistencyV2 string


// All allowed values of CommitConsistencyV2 enum
var AllowedCommitConsistencyV2EnumValues = []CommitConsistencyV2{
	"NOT_CHECKED",
	"COMMIT_CONSISTENT",
	"COMMIT_CONTENT_INCONSISTENT",
	"COMMIT_INCONSISTENT",
}

func (v *CommitConsistencyV2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommitConsistencyV2(value)
	for _, existing := range AllowedCommitConsistencyV2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommitConsistencyV2", value)
}

// NewCommitConsistencyV2FromValue returns a pointer to a valid CommitConsistencyV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommitConsistencyV2FromValue(v string) (*CommitConsistencyV2, error) {
	ev := CommitConsistencyV2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommitConsistencyV2: valid values are %v", v, AllowedCommitConsistencyV2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommitConsistencyV2) IsValid() bool {
	for _, existing := range AllowedCommitConsistencyV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommitConsistency_V2 value
func (v CommitConsistencyV2) Ptr() *CommitConsistencyV2 {
	return &v
}

type NullableCommitConsistencyV2 struct {
	value *CommitConsistencyV2
	isSet bool
}

func (v NullableCommitConsistencyV2) Get() *CommitConsistencyV2 {
	return v.value
}

func (v *NullableCommitConsistencyV2) Set(val *CommitConsistencyV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitConsistencyV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitConsistencyV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitConsistencyV2(val *CommitConsistencyV2) *NullableCommitConsistencyV2 {
	return &NullableCommitConsistencyV2{value: val, isSet: true}
}

func (v NullableCommitConsistencyV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitConsistencyV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

