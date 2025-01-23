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

// checks if the Content4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Content4{}

// Content4 struct for Content4
type Content4 struct {
	Content1AnyOf *Content1AnyOf
	Content1AnyOf1 *Content1AnyOf1
	Content1AnyOf2 *Content1AnyOf2
	Content1AnyOf3 *Content1AnyOf3
	IcebergTableState *IcebergTableState
	DeltaLakeTableV1 *DeltaLakeTableV1
	IcebergTableV1 *IcebergTableV1
	IcebergViewV1 *IcebergViewV1
	NamespaceV1 *NamespaceV1
	UDFV1 *UDFV1
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Content4) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'Content_1_anyOf'
	if jsonDict["type"] == "Content_1_anyOf" {
		// try to unmarshal JSON data into Content1AnyOf
		err = json.Unmarshal(data, &dst.Content1AnyOf);
		if err == nil {
			jsonContent1AnyOf, _ := json.Marshal(dst.Content1AnyOf)
			if string(jsonContent1AnyOf) == "{}" { // empty struct
				dst.Content1AnyOf = nil
			} else {
				return nil // data stored in dst.Content1AnyOf, return on the first match
			}
		} else {
			dst.Content1AnyOf = nil
		}
	}

	// check if the discriminator value is 'Content_1_anyOf_1'
	if jsonDict["type"] == "Content_1_anyOf_1" {
		// try to unmarshal JSON data into Content1AnyOf1
		err = json.Unmarshal(data, &dst.Content1AnyOf1);
		if err == nil {
			jsonContent1AnyOf1, _ := json.Marshal(dst.Content1AnyOf1)
			if string(jsonContent1AnyOf1) == "{}" { // empty struct
				dst.Content1AnyOf1 = nil
			} else {
				return nil // data stored in dst.Content1AnyOf1, return on the first match
			}
		} else {
			dst.Content1AnyOf1 = nil
		}
	}

	// check if the discriminator value is 'Content_1_anyOf_2'
	if jsonDict["type"] == "Content_1_anyOf_2" {
		// try to unmarshal JSON data into Content1AnyOf2
		err = json.Unmarshal(data, &dst.Content1AnyOf2);
		if err == nil {
			jsonContent1AnyOf2, _ := json.Marshal(dst.Content1AnyOf2)
			if string(jsonContent1AnyOf2) == "{}" { // empty struct
				dst.Content1AnyOf2 = nil
			} else {
				return nil // data stored in dst.Content1AnyOf2, return on the first match
			}
		} else {
			dst.Content1AnyOf2 = nil
		}
	}

	// check if the discriminator value is 'Content_1_anyOf_3'
	if jsonDict["type"] == "Content_1_anyOf_3" {
		// try to unmarshal JSON data into Content1AnyOf3
		err = json.Unmarshal(data, &dst.Content1AnyOf3);
		if err == nil {
			jsonContent1AnyOf3, _ := json.Marshal(dst.Content1AnyOf3)
			if string(jsonContent1AnyOf3) == "{}" { // empty struct
				dst.Content1AnyOf3 = nil
			} else {
				return nil // data stored in dst.Content1AnyOf3, return on the first match
			}
		} else {
			dst.Content1AnyOf3 = nil
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
	// if jsonDict["type"] == "" {
	// 	// try to unmarshal JSON data into ERRORUNKNOWN
	// 	err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
	// 	if err == nil {
	// 		jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
	// 		if string(jsonERRORUNKNOWN) == "{}" { // empty struct
	// 			dst.ERRORUNKNOWN = nil
	// 		} else {
	// 			return nil // data stored in dst.ERRORUNKNOWN, return on the first match
	// 		}
	// 	} else {
	// 		dst.ERRORUNKNOWN = nil
	// 	}
	// }

	// try to unmarshal JSON data into Content1AnyOf
	err = json.Unmarshal(data, &dst.Content1AnyOf);
	if err == nil {
		jsonContent1AnyOf, _ := json.Marshal(dst.Content1AnyOf)
		if string(jsonContent1AnyOf) == "{}" { // empty struct
			dst.Content1AnyOf = nil
		} else {
			return nil // data stored in dst.Content1AnyOf, return on the first match
		}
	} else {
		dst.Content1AnyOf = nil
	}

	// try to unmarshal JSON data into Content1AnyOf1
	err = json.Unmarshal(data, &dst.Content1AnyOf1);
	if err == nil {
		jsonContent1AnyOf1, _ := json.Marshal(dst.Content1AnyOf1)
		if string(jsonContent1AnyOf1) == "{}" { // empty struct
			dst.Content1AnyOf1 = nil
		} else {
			return nil // data stored in dst.Content1AnyOf1, return on the first match
		}
	} else {
		dst.Content1AnyOf1 = nil
	}

	// try to unmarshal JSON data into Content1AnyOf2
	err = json.Unmarshal(data, &dst.Content1AnyOf2);
	if err == nil {
		jsonContent1AnyOf2, _ := json.Marshal(dst.Content1AnyOf2)
		if string(jsonContent1AnyOf2) == "{}" { // empty struct
			dst.Content1AnyOf2 = nil
		} else {
			return nil // data stored in dst.Content1AnyOf2, return on the first match
		}
	} else {
		dst.Content1AnyOf2 = nil
	}

	// try to unmarshal JSON data into Content1AnyOf3
	err = json.Unmarshal(data, &dst.Content1AnyOf3);
	if err == nil {
		jsonContent1AnyOf3, _ := json.Marshal(dst.Content1AnyOf3)
		if string(jsonContent1AnyOf3) == "{}" { // empty struct
			dst.Content1AnyOf3 = nil
		} else {
			return nil // data stored in dst.Content1AnyOf3, return on the first match
		}
	} else {
		dst.Content1AnyOf3 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Content4)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Content4) MarshalJSON() ([]byte, error) {
	if src.Content1AnyOf != nil {
		return json.Marshal(&src.Content1AnyOf)
	}

	if src.Content1AnyOf1 != nil {
		return json.Marshal(&src.Content1AnyOf1)
	}

	if src.Content1AnyOf2 != nil {
		return json.Marshal(&src.Content1AnyOf2)
	}

	if src.Content1AnyOf3 != nil {
		return json.Marshal(&src.Content1AnyOf3)
	}

	if src.IcebergTableState != nil {
		return json.Marshal(&src.IcebergTableState)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}

func (src Content4) ToMap() (map[string]interface{}, error) {
	if src.Content1AnyOf != nil {
		return src.Content1AnyOf.ToMap()
	}

	if src.Content1AnyOf1 != nil {
		return src.Content1AnyOf1.ToMap()
	}

	if src.Content1AnyOf2 != nil {
		return src.Content1AnyOf2.ToMap()
	}

	if src.Content1AnyOf3 != nil {
		return src.Content1AnyOf3.ToMap()
	}

	if src.IcebergTableState != nil {
		return src.IcebergTableState.ToMap()
	}

	// if src.MapmapOfStringAny != nil {
	// 	return src.MapmapOfStringAny.ToMap()
	// }

    return nil, nil // no data in anyOf schemas
}

type NullableContent4 struct {
	value *Content4
	isSet bool
}

func (v NullableContent4) Get() *Content4 {
	return v.value
}

func (v *NullableContent4) Set(val *Content4) {
	v.value = val
	v.isSet = true
}

func (v NullableContent4) IsSet() bool {
	return v.isSet
}

func (v *NullableContent4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent4(val *Content4) *NullableContent4 {
	return &NullableContent4{value: val, isSet: true}
}

func (v NullableContent4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


