/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// Operation1 - Describes an operation to be performed against one content object.  The Nessie backend will validate the correctness of the operations.
type Operation1 struct {
	CommitResponseAddedContentsInner *CommitResponseAddedContentsInner
	DeleteContentOperationForAContentKey *DeleteContentOperationForAContentKey
	PutContentOperationForAContentKey *PutContentOperationForAContentKey
}

// CommitResponseAddedContentsInnerAsOperation1 is a convenience function that returns CommitResponseAddedContentsInner wrapped in Operation1
func CommitResponseAddedContentsInnerAsOperation1(v *CommitResponseAddedContentsInner) Operation1 {
	return Operation1{
		CommitResponseAddedContentsInner: v,
	}
}

// DeleteContentOperationForAContentKeyAsOperation1 is a convenience function that returns DeleteContentOperationForAContentKey wrapped in Operation1
func DeleteContentOperationForAContentKeyAsOperation1(v *DeleteContentOperationForAContentKey) Operation1 {
	return Operation1{
		DeleteContentOperationForAContentKey: v,
	}
}

// PutContentOperationForAContentKeyAsOperation1 is a convenience function that returns PutContentOperationForAContentKey wrapped in Operation1
func PutContentOperationForAContentKeyAsOperation1(v *PutContentOperationForAContentKey) Operation1 {
	return Operation1{
		PutContentOperationForAContentKey: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Operation1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CommitResponseAddedContentsInner
	err = newStrictDecoder(data).Decode(&dst.CommitResponseAddedContentsInner)
	if err == nil {
		jsonCommitResponseAddedContentsInner, _ := json.Marshal(dst.CommitResponseAddedContentsInner)
		if string(jsonCommitResponseAddedContentsInner) == "{}" { // empty struct
			dst.CommitResponseAddedContentsInner = nil
		} else {
			if err = validator.Validate(dst.CommitResponseAddedContentsInner); err != nil {
				dst.CommitResponseAddedContentsInner = nil
			} else {
				match++
			}
		}
	} else {
		dst.CommitResponseAddedContentsInner = nil
	}

	// try to unmarshal data into DeleteContentOperationForAContentKey
	err = newStrictDecoder(data).Decode(&dst.DeleteContentOperationForAContentKey)
	if err == nil {
		jsonDeleteContentOperationForAContentKey, _ := json.Marshal(dst.DeleteContentOperationForAContentKey)
		if string(jsonDeleteContentOperationForAContentKey) == "{}" { // empty struct
			dst.DeleteContentOperationForAContentKey = nil
		} else {
			if err = validator.Validate(dst.DeleteContentOperationForAContentKey); err != nil {
				dst.DeleteContentOperationForAContentKey = nil
			} else {
				match++
			}
		}
	} else {
		dst.DeleteContentOperationForAContentKey = nil
	}

	// try to unmarshal data into PutContentOperationForAContentKey
	err = newStrictDecoder(data).Decode(&dst.PutContentOperationForAContentKey)
	if err == nil {
		jsonPutContentOperationForAContentKey, _ := json.Marshal(dst.PutContentOperationForAContentKey)
		if string(jsonPutContentOperationForAContentKey) == "{}" { // empty struct
			dst.PutContentOperationForAContentKey = nil
		} else {
			if err = validator.Validate(dst.PutContentOperationForAContentKey); err != nil {
				dst.PutContentOperationForAContentKey = nil
			} else {
				match++
			}
		}
	} else {
		dst.PutContentOperationForAContentKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CommitResponseAddedContentsInner = nil
		dst.DeleteContentOperationForAContentKey = nil
		dst.PutContentOperationForAContentKey = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Operation1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Operation1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Operation1) MarshalJSON() ([]byte, error) {
	if src.CommitResponseAddedContentsInner != nil {
		return json.Marshal(&src.CommitResponseAddedContentsInner)
	}

	if src.DeleteContentOperationForAContentKey != nil {
		return json.Marshal(&src.DeleteContentOperationForAContentKey)
	}

	if src.PutContentOperationForAContentKey != nil {
		return json.Marshal(&src.PutContentOperationForAContentKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Operation1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CommitResponseAddedContentsInner != nil {
		return obj.CommitResponseAddedContentsInner
	}

	if obj.DeleteContentOperationForAContentKey != nil {
		return obj.DeleteContentOperationForAContentKey
	}

	if obj.PutContentOperationForAContentKey != nil {
		return obj.PutContentOperationForAContentKey
	}

	// all schemas are nil
	return nil
}

type NullableOperation1 struct {
	value *Operation1
	isSet bool
}

func (v NullableOperation1) Get() *Operation1 {
	return v.value
}

func (v *NullableOperation1) Set(val *Operation1) {
	v.value = val
	v.isSet = true
}

func (v NullableOperation1) IsSet() bool {
	return v.isSet
}

func (v *NullableOperation1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperation1(val *Operation1) *NullableOperation1 {
	return &NullableOperation1{value: val, isSet: true}
}

func (v NullableOperation1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperation1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


