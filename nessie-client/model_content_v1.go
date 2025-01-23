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

// checks if the ContentV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV1{}

// ContentV1 struct for ContentV1
type ContentV1 struct {
	Content1AnyOf     *Content1AnyOf
	ContentAnyOf      *ContentAnyOf
	ContentAnyOf1      *ContentAnyOf1
	ContentAnyOf2     *ContentAnyOf2
	ContentAnyOf3     *ContentAnyOf3
	IcebergTableState *IcebergTableState
    DeltaLakeTableV1    *DeltaLakeTableV1 // Add this field
    IcebergTableV1      *IcebergTableV1   // Add this field
    IcebergViewV1       *IcebergViewV1    // Add this field
    UDFV1               *UDFV1            // Add this field
    NamespaceV1         *NamespaceV1      // Add this field
	IcebergTableState1 *IcebergTableState1
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ContentV1) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ContentV1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentV1) MarshalJSON() ([]byte, error) {
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

func (src ContentV1) ToMap() (map[string]interface{}, error) {
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

	// if src.MapmapOfStringAny != nil {
	// 	return src.MapmapOfStringAny.ToMap()
	// }

    return nil, nil // no data in anyOf schemas
}

type NullableContentV1 struct {
	value *ContentV1
	isSet bool
}

func (v NullableContentV1) Get() *ContentV1 {
	return v.value
}

func (v *NullableContentV1) Set(val *ContentV1) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV1) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV1(val *ContentV1) *NullableContentV1 {
	return &NullableContentV1{value: val, isSet: true}
}

func (v NullableContentV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


