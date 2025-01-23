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

// checks if the ContentV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentV2{}

// ContentV2 struct for ContentV2
type ContentV2 struct {
	Content1AnyOf *Content1AnyOf
	ContentAnyOf *ContentAnyOf
	ContentAnyOf2 *ContentAnyOf2
	ContentAnyOf3 *ContentAnyOf3
	IcebergTableState1 *IcebergTableState1
	MapmapOfStringAny *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ContentV2) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DELTA_LAKE_TABLE'
	if jsonDict["type"] == "DELTA_LAKE_TABLE" {
		// try to unmarshal JSON data into DeltaLakeTableV2
		err = json.Unmarshal(data, &dst.DeltaLakeTableV2);
		if err == nil {
			jsonDeltaLakeTableV2, _ := json.Marshal(dst.DeltaLakeTableV2)
			if string(jsonDeltaLakeTableV2) == "{}" { // empty struct
				dst.DeltaLakeTableV2 = nil
			} else {
				return nil // data stored in dst.DeltaLakeTableV2, return on the first match
			}
		} else {
			dst.DeltaLakeTableV2 = nil
		}
	}

	// check if the discriminator value is 'ICEBERG_TABLE'
	if jsonDict["type"] == "ICEBERG_TABLE" {
		// try to unmarshal JSON data into IcebergTableV2
		err = json.Unmarshal(data, &dst.IcebergTableV2);
		if err == nil {
			jsonIcebergTableV2, _ := json.Marshal(dst.IcebergTableV2)
			if string(jsonIcebergTableV2) == "{}" { // empty struct
				dst.IcebergTableV2 = nil
			} else {
				return nil // data stored in dst.IcebergTableV2, return on the first match
			}
		} else {
			dst.IcebergTableV2 = nil
		}
	}

	// check if the discriminator value is 'ICEBERG_VIEW'
	if jsonDict["type"] == "ICEBERG_VIEW" {
		// try to unmarshal JSON data into IcebergViewV2
		err = json.Unmarshal(data, &dst.IcebergViewV2);
		if err == nil {
			jsonIcebergViewV2, _ := json.Marshal(dst.IcebergViewV2)
			if string(jsonIcebergViewV2) == "{}" { // empty struct
				dst.IcebergViewV2 = nil
			} else {
				return nil // data stored in dst.IcebergViewV2, return on the first match
			}
		} else {
			dst.IcebergViewV2 = nil
		}
	}

	// check if the discriminator value is 'NAMESPACE'
	if jsonDict["type"] == "NAMESPACE" {
		// try to unmarshal JSON data into NamespaceV2
		err = json.Unmarshal(data, &dst.NamespaceV2);
		if err == nil {
			jsonNamespaceV2, _ := json.Marshal(dst.NamespaceV2)
			if string(jsonNamespaceV2) == "{}" { // empty struct
				dst.NamespaceV2 = nil
			} else {
				return nil // data stored in dst.NamespaceV2, return on the first match
			}
		} else {
			dst.NamespaceV2 = nil
		}
	}

	// check if the discriminator value is 'UDF'
	if jsonDict["type"] == "UDF" {
		// try to unmarshal JSON data into UDFV2
		err = json.Unmarshal(data, &dst.UDFV2);
		if err == nil {
			jsonUDFV2, _ := json.Marshal(dst.UDFV2)
			if string(jsonUDFV2) == "{}" { // empty struct
				dst.UDFV2 = nil
			} else {
				return nil // data stored in dst.UDFV2, return on the first match
			}
		} else {
			dst.UDFV2 = nil
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

	// check if the discriminator value is 'Iceberg_table_state_1'
	if jsonDict["type"] == "Iceberg_table_state_1" {
		// try to unmarshal JSON data into IcebergTableState1
		err = json.Unmarshal(data, &dst.IcebergTableState1);
		if err == nil {
			jsonIcebergTableState1, _ := json.Marshal(dst.IcebergTableState1)
			if string(jsonIcebergTableState1) == "{}" { // empty struct
				dst.IcebergTableState1 = nil
			} else {
				return nil // data stored in dst.IcebergTableState1, return on the first match
			}
		} else {
			dst.IcebergTableState1 = nil
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

	// try to unmarshal JSON data into IcebergTableState1
	err = json.Unmarshal(data, &dst.IcebergTableState1);
	if err == nil {
		jsonIcebergTableState1, _ := json.Marshal(dst.IcebergTableState1)
		if string(jsonIcebergTableState1) == "{}" { // empty struct
			dst.IcebergTableState1 = nil
		} else {
			return nil // data stored in dst.IcebergTableState1, return on the first match
		}
	} else {
		dst.IcebergTableState1 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ContentV2)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContentV2) MarshalJSON() ([]byte, error) {
	if src.Content1AnyOf != nil {
		return json.Marshal(&src.Content1AnyOf)
	}

	if src.ContentAnyOf != nil {
		return json.Marshal(&src.ContentAnyOf)
	}

	if src.ContentAnyOf2 != nil {
		return json.Marshal(&src.ContentAnyOf2)
	}

	if src.ContentAnyOf3 != nil {
		return json.Marshal(&src.ContentAnyOf3)
	}

	if src.IcebergTableState1 != nil {
		return json.Marshal(&src.IcebergTableState1)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in anyOf schemas
}

func (src ContentV2) ToMap() (map[string]interface{}, error) {
	if src.Content1AnyOf != nil {
		return src.Content1AnyOf.ToMap()
	}

	if src.ContentAnyOf != nil {
		return src.ContentAnyOf.ToMap()
	}

	if src.ContentAnyOf2 != nil {
		return src.ContentAnyOf2.ToMap()
	}

	if src.ContentAnyOf3 != nil {
		return src.ContentAnyOf3.ToMap()
	}

	if src.IcebergTableState1 != nil {
		return src.IcebergTableState1.ToMap()
	}

	if src.MapmapOfStringAny != nil {
		return src.MapmapOfStringAny.ToMap()
	}

    return nil, nil // no data in anyOf schemas
}

type NullableContentV2 struct {
	value *ContentV2
	isSet bool
}

func (v NullableContentV2) Get() *ContentV2 {
	return v.value
}

func (v *NullableContentV2) Set(val *ContentV2) {
	v.value = val
	v.isSet = true
}

func (v NullableContentV2) IsSet() bool {
	return v.isSet
}

func (v *NullableContentV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentV2(val *ContentV2) *NullableContentV2 {
	return &NullableContentV2{value: val, isSet: true}
}

func (v NullableContentV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


