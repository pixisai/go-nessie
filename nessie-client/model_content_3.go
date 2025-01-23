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

// checks if the Content3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Content3{}

// Content3 struct for Content3
type Content3 struct {
	ContentAnyOf *ContentAnyOf
	ContentAnyOf1 *ContentAnyOf1
	ContentAnyOf2 *ContentAnyOf2
	ContentAnyOf3 *ContentAnyOf3
	IcebergTableState *IcebergTableState
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Content3) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DELTA_LAKE_TABLE'
	if jsonDict["type"] == "DELTA_LAKE_TABLE" {
		// try to unmarshal JSON data into DeltaLakeTableV1
		err = json.Unmarshal(data, &dst.DeltaLakeTableV1);
		if err == nil {
			jsonDeltaLakeTableV1, _ := json.Marshal(dst.DeltaLakeTableV1)
			if string(jsonDeltaLakeTableV1) == "{}" { // empty struct
				dst.DeltaLakeTableV1 = nil
			} else {
				return nil // data stored in dst.DeltaLakeTableV1, return on the first match
			}
		} else {
			dst.DeltaLakeTableV1 = nil
		}
	}

	// check if the discriminator value is 'ICEBERG_TABLE'
	if jsonDict["type"] == "ICEBERG_TABLE" {
		// try to unmarshal JSON data into IcebergTableV1
		err = json.Unmarshal(data, &dst.IcebergTableV1);
		if err == nil {
			jsonIcebergTableV1, _ := json.Marshal(dst.IcebergTableV1)
			if string(jsonIcebergTableV1) == "{}" { // empty struct
				dst.IcebergTableV1 = nil
			} else {
				return nil // data stored in dst.IcebergTableV1, return on the first match
			}
		} else {
			dst.IcebergTableV1 = nil
		}
	}

	// check if the discriminator value is 'ICEBERG_VIEW'
	if jsonDict["type"] == "ICEBERG_VIEW" {
		// try to unmarshal JSON data into IcebergViewV1
		err = json.Unmarshal(data, &dst.IcebergViewV1);
		if err == nil {
			jsonIcebergViewV1, _ := json.Marshal(dst.IcebergViewV1)
			if string(jsonIcebergViewV1) == "{}" { // empty struct
				dst.IcebergViewV1 = nil
			} else {
				return nil // data stored in dst.IcebergViewV1, return on the first match
			}
		} else {
			dst.IcebergViewV1 = nil
		}
	}

	// check if the discriminator value is 'NAMESPACE'
	if jsonDict["type"] == "NAMESPACE" {
		// try to unmarshal JSON data into NamespaceV1
		err = json.Unmarshal(data, &dst.NamespaceV1);
		if err == nil {
			jsonNamespaceV1, _ := json.Marshal(dst.NamespaceV1)
			if string(jsonNamespaceV1) == "{}" { // empty struct
				dst.NamespaceV1 = nil
			} else {
				return nil // data stored in dst.NamespaceV1, return on the first match
			}
		} else {
			dst.NamespaceV1 = nil
		}
	}

	// check if the discriminator value is 'UDF'
	if jsonDict["type"] == "UDF" {
		// try to unmarshal JSON data into UDFV1
		err = json.Unmarshal(data, &dst.UDFV1);
		if err == nil {
			jsonUDFV1, _ := json.Marshal(dst.UDFV1)
			if string(jsonUDFV1) == "{}" { // empty struct
				dst.UDFV1 = nil
			} else {
				return nil // data stored in dst.UDFV1, return on the first match
			}
		} else {
			dst.UDFV1 = nil
		}
	}

	// check if the discriminator value is 'Content_anyOf'
	if jsonDict["type"] == "Content_anyOf" {
		// try to unmarshal JSON data into ContentAnyOf
		err = json.Unmarshal(data, &dst.ContentAnyOf);
		if err == nil {
			jsonContentAnyOf, _ := json.Marshal(dst.ContentAnyOf)
			if string(jsonContentAnyOf) == "{}" { // empty struct
				dst.ContentAnyOf = nil
			} else {
				return nil // data stored in dst.ContentAnyOf, return on the first match
			}
		} else {
			dst.ContentAnyOf = nil
		}
	}

	// check if the discriminator value is 'Content_anyOf_1'
	if jsonDict["type"] == "Content_anyOf_1" {
		// try to unmarshal JSON data into ContentAnyOf1
		err = json.Unmarshal(data, &dst.ContentAnyOf1);
		if err == nil {
			jsonContentAnyOf1, _ := json.Marshal(dst.ContentAnyOf1)
			if string(jsonContentAnyOf1) == "{}" { // empty struct
				dst.ContentAnyOf1 = nil
			} else {
				return nil // data stored in dst.ContentAnyOf1, return on the first match
			}
		} else {
			dst.ContentAnyOf1 = nil
		}
	}

	// check if the discriminator value is 'Content_anyOf_2'
	if jsonDict["type"] == "Content_anyOf_2" {
		// try to unmarshal JSON data into ContentAnyOf2
		err = json.Unmarshal(data, &dst.ContentAnyOf2);
		if err == nil {
			jsonContentAnyOf2, _ := json.Marshal(dst.ContentAnyOf2)
			if string(jsonContentAnyOf2) == "{}" { // empty struct
				dst.ContentAnyOf2 = nil
			} else {
				return nil // data stored in dst.ContentAnyOf2, return on the first match
			}
		} else {
			dst.ContentAnyOf2 = nil
		}
	}

	// check if the discriminator value is 'Content_anyOf_3'
	if jsonDict["type"] == "Content_anyOf_3" {
		// try to unmarshal JSON data into ContentAnyOf3
		err = json.Unmarshal(data, &dst.ContentAnyOf3);
		if err == nil {
			jsonContentAnyOf3, _ := json.Marshal(dst.ContentAnyOf3)
			if string(jsonContentAnyOf3) == "{}" { // empty struct
				dst.ContentAnyOf3 = nil
			} else {
				return nil // data stored in dst.ContentAnyOf3, return on the first match
			}
		} else {
			dst.ContentAnyOf3 = nil
		}
	}

	// check if the discriminator value is 'Iceberg_table_state'
	if jsonDict["type"] == "Iceberg_table_state" {
		// try to unmarshal JSON data into IcebergTableState
		err = json.Unmarshal(data, &dst.IcebergTableState);
		if err == nil {
			jsonIcebergTableState, _ := json.Marshal(dst.IcebergTableState)
			if string(jsonIcebergTableState) == "{}" { // empty struct
				dst.IcebergTableState = nil
			} else {
				return nil // data stored in dst.IcebergTableState, return on the first match
			}
		} else {
			dst.IcebergTableState = nil
		}
	}

	// check if the discriminator value is ''
	if jsonDict["type"] == "" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// try to unmarshal JSON data into ContentAnyOf
	err = json.Unmarshal(data, &dst.ContentAnyOf);
	if err == nil {
		jsonContentAnyOf, _ := json.Marshal(dst.ContentAnyOf)
		if string(jsonContentAnyOf) == "{}" { // empty struct
			dst.ContentAnyOf = nil
		} else {
			return nil // data stored in dst.ContentAnyOf, return on the first match
		}
	} else {
		dst.ContentAnyOf = nil
	}

	// try to unmarshal JSON data into ContentAnyOf1
	err = json.Unmarshal(data, &dst.ContentAnyOf1);
	if err == nil {
		jsonContentAnyOf1, _ := json.Marshal(dst.ContentAnyOf1)
		if string(jsonContentAnyOf1) == "{}" { // empty struct
			dst.ContentAnyOf1 = nil
		} else {
			return nil // data stored in dst.ContentAnyOf1, return on the first match
		}
	} else {
		dst.ContentAnyOf1 = nil
	}

	// try to unmarshal JSON data into ContentAnyOf2
	err = json.Unmarshal(data, &dst.ContentAnyOf2);
	if err == nil {
		jsonContentAnyOf2, _ := json.Marshal(dst.ContentAnyOf2)
		if string(jsonContentAnyOf2) == "{}" { // empty struct
			dst.ContentAnyOf2 = nil
		} else {
			return nil // data stored in dst.ContentAnyOf2, return on the first match
		}
	} else {
		dst.ContentAnyOf2 = nil
	}

	// try to unmarshal JSON data into ContentAnyOf3
	err = json.Unmarshal(data, &dst.ContentAnyOf3);
	if err == nil {
		jsonContentAnyOf3, _ := json.Marshal(dst.ContentAnyOf3)
		if string(jsonContentAnyOf3) == "{}" { // empty struct
			dst.ContentAnyOf3 = nil
		} else {
			return nil // data stored in dst.ContentAnyOf3, return on the first match
		}
	} else {
		dst.ContentAnyOf3 = nil
	}

	// try to unmarshal JSON data into IcebergTableState
	err = json.Unmarshal(data, &dst.IcebergTableState);
	if err == nil {
		jsonIcebergTableState, _ := json.Marshal(dst.IcebergTableState)
		if string(jsonIcebergTableState) == "{}" { // empty struct
			dst.IcebergTableState = nil
		} else {
			return nil // data stored in dst.IcebergTableState, return on the first match
		}
	} else {
		dst.IcebergTableState = nil
	}

	// try to unmarshal JSON data into MapmapOfStringAny
	err = json.Unmarshal(data, &dst.MapmapOfStringAny);
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			return nil // data stored in dst.MapmapOfStringAny, return on the first match
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Content3)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Content3) MarshalJSON() ([]byte, error) {
	if src.ContentAnyOf != nil {
		return json.Marshal(&src.ContentAnyOf)
	}

	if src.ContentAnyOf1 != nil {
		return json.Marshal(&src.ContentAnyOf1)
	}

	if src.ContentAnyOf2 != nil {
		return json.Marshal(&src.ContentAnyOf2)
	}

	if src.ContentAnyOf3 != nil {
		return json.Marshal(&src.ContentAnyOf3)
	}

	if src.IcebergTableState != nil {
		return json.Marshal(&src.IcebergTableState)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}

func (src Content3) ToMap() (map[string]interface{}, error) {
	if src.ContentAnyOf != nil {
		return src.ContentAnyOf.ToMap()
	}

	if src.ContentAnyOf1 != nil {
		return src.ContentAnyOf1.ToMap()
	}

	if src.ContentAnyOf2 != nil {
		return src.ContentAnyOf2.ToMap()
	}

	if src.ContentAnyOf3 != nil {
		return src.ContentAnyOf3.ToMap()
	}

	if src.IcebergTableState != nil {
		return src.IcebergTableState.ToMap()
	}

	if src.MapmapOfStringAny != nil {
		return src.MapmapOfStringAny.ToMap()
	}

    return nil, nil // no data in anyOf schemas
}

type NullableContent3 struct {
	value *Content3
	isSet bool
}

func (v NullableContent3) Get() *Content3 {
	return v.value
}

func (v *NullableContent3) Set(val *Content3) {
	v.value = val
	v.isSet = true
}

func (v NullableContent3) IsSet() bool {
	return v.isSet
}

func (v *NullableContent3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent3(val *Content3) *NullableContent3 {
	return &NullableContent3{value: val, isSet: true}
}

func (v NullableContent3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


