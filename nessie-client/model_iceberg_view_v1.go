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

// checks if the IcebergViewV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IcebergViewV1{}

// IcebergViewV1 struct for IcebergViewV1
type IcebergViewV1 struct {
	Id *string `json:"id,omitempty"`
	MetadataLocation string `json:"metadataLocation" validate:"regexp=\\\\S"`
	VersionId *int64 `json:"versionId,omitempty"`
	SchemaId *int32 `json:"schemaId,omitempty"`
	// Deprecated
	SqlText string `json:"sqlText" validate:"regexp=\\\\S"`
	// Deprecated
	Dialect *string `json:"dialect,omitempty"`
	// Deprecated
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type _IcebergViewV1 IcebergViewV1

// NewIcebergViewV1 instantiates a new IcebergViewV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcebergViewV1(metadataLocation string, sqlText string) *IcebergViewV1 {
	this := IcebergViewV1{}
	this.MetadataLocation = metadataLocation
	this.SqlText = sqlText
	return &this
}

// NewIcebergViewV1WithDefaults instantiates a new IcebergViewV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIcebergViewV1WithDefaults() *IcebergViewV1 {
	this := IcebergViewV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IcebergViewV1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergViewV1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IcebergViewV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IcebergViewV1) SetId(v string) {
	o.Id = &v
}

// GetMetadataLocation returns the MetadataLocation field value
func (o *IcebergViewV1) GetMetadataLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataLocation
}

// GetMetadataLocationOk returns a tuple with the MetadataLocation field value
// and a boolean to check if the value has been set.
func (o *IcebergViewV1) GetMetadataLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataLocation, true
}

// SetMetadataLocation sets field value
func (o *IcebergViewV1) SetMetadataLocation(v string) {
	o.MetadataLocation = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *IcebergViewV1) GetVersionId() int64 {
	if o == nil || IsNil(o.VersionId) {
		var ret int64
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergViewV1) GetVersionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *IcebergViewV1) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given int64 and assigns it to the VersionId field.
func (o *IcebergViewV1) SetVersionId(v int64) {
	o.VersionId = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *IcebergViewV1) GetSchemaId() int32 {
	if o == nil || IsNil(o.SchemaId) {
		var ret int32
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IcebergViewV1) GetSchemaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaId) {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *IcebergViewV1) HasSchemaId() bool {
	if o != nil && !IsNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given int32 and assigns it to the SchemaId field.
func (o *IcebergViewV1) SetSchemaId(v int32) {
	o.SchemaId = &v
}

// GetSqlText returns the SqlText field value
// Deprecated
func (o *IcebergViewV1) GetSqlText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *IcebergViewV1) GetSqlTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlText, true
}

// SetSqlText sets field value
// Deprecated
func (o *IcebergViewV1) SetSqlText(v string) {
	o.SqlText = v
}

// GetDialect returns the Dialect field value if set, zero value otherwise.
// Deprecated
func (o *IcebergViewV1) GetDialect() string {
	if o == nil || IsNil(o.Dialect) {
		var ret string
		return ret
	}
	return *o.Dialect
}

// GetDialectOk returns a tuple with the Dialect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *IcebergViewV1) GetDialectOk() (*string, bool) {
	if o == nil || IsNil(o.Dialect) {
		return nil, false
	}
	return o.Dialect, true
}

// HasDialect returns a boolean if a field has been set.
func (o *IcebergViewV1) HasDialect() bool {
	if o != nil && !IsNil(o.Dialect) {
		return true
	}

	return false
}

// SetDialect gets a reference to the given string and assigns it to the Dialect field.
// Deprecated
func (o *IcebergViewV1) SetDialect(v string) {
	o.Dialect = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
// Deprecated
func (o *IcebergViewV1) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *IcebergViewV1) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IcebergViewV1) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
// Deprecated
func (o *IcebergViewV1) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o IcebergViewV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IcebergViewV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["metadataLocation"] = o.MetadataLocation
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	if !IsNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	toSerialize["sqlText"] = o.SqlText
	if !IsNil(o.Dialect) {
		toSerialize["dialect"] = o.Dialect
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *IcebergViewV1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataLocation",
		"sqlText",
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

	varIcebergViewV1 := _IcebergViewV1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIcebergViewV1)

	if err != nil {
		return err
	}

	*o = IcebergViewV1(varIcebergViewV1)

	return err
}

type NullableIcebergViewV1 struct {
	value *IcebergViewV1
	isSet bool
}

func (v NullableIcebergViewV1) Get() *IcebergViewV1 {
	return v.value
}

func (v *NullableIcebergViewV1) Set(val *IcebergViewV1) {
	v.value = val
	v.isSet = true
}

func (v NullableIcebergViewV1) IsSet() bool {
	return v.isSet
}

func (v *NullableIcebergViewV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcebergViewV1(val *IcebergViewV1) *NullableIcebergViewV1 {
	return &NullableIcebergViewV1{value: val, isSet: true}
}

func (v NullableIcebergViewV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcebergViewV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


