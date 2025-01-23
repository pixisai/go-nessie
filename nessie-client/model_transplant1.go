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

// checks if the Transplant1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transplant1{}

// Transplant1 struct for Transplant1
type Transplant1 struct {
	FromRefName string `json:"fromRefName" validate:"regexp=^(?:[A-Za-z](?:(?:(?![.][.])[A-Za-z0-9.\\/_-])*[A-Za-z0-9_-])?)|-$"`
	HashesToTransplant []string `json:"hashesToTransplant"`
	KeyMergeModes []MergeOperationKeyMergeModesInner `json:"keyMergeModes,omitempty"`
	DefaultKeyMergeMode *string `json:"defaultKeyMergeMode,omitempty"`
	DryRun *bool `json:"dryRun,omitempty"`
	FetchAdditionalInfo *bool `json:"fetchAdditionalInfo,omitempty"`
	ReturnConflictAsResult *bool `json:"returnConflictAsResult,omitempty"`
}

type _Transplant1 Transplant1

// NewTransplant1 instantiates a new Transplant1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransplant1(fromRefName string, hashesToTransplant []string) *Transplant1 {
	this := Transplant1{}
	this.FromRefName = fromRefName
	this.HashesToTransplant = hashesToTransplant
	return &this
}

// NewTransplant1WithDefaults instantiates a new Transplant1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransplant1WithDefaults() *Transplant1 {
	this := Transplant1{}
	return &this
}

// GetFromRefName returns the FromRefName field value
func (o *Transplant1) GetFromRefName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromRefName
}

// GetFromRefNameOk returns a tuple with the FromRefName field value
// and a boolean to check if the value has been set.
func (o *Transplant1) GetFromRefNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromRefName, true
}

// SetFromRefName sets field value
func (o *Transplant1) SetFromRefName(v string) {
	o.FromRefName = v
}

// GetHashesToTransplant returns the HashesToTransplant field value
func (o *Transplant1) GetHashesToTransplant() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HashesToTransplant
}

// GetHashesToTransplantOk returns a tuple with the HashesToTransplant field value
// and a boolean to check if the value has been set.
func (o *Transplant1) GetHashesToTransplantOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashesToTransplant, true
}

// SetHashesToTransplant sets field value
func (o *Transplant1) SetHashesToTransplant(v []string) {
	o.HashesToTransplant = v
}

// GetKeyMergeModes returns the KeyMergeModes field value if set, zero value otherwise.
func (o *Transplant1) GetKeyMergeModes() []MergeOperationKeyMergeModesInner {
	if o == nil || IsNil(o.KeyMergeModes) {
		var ret []MergeOperationKeyMergeModesInner
		return ret
	}
	return o.KeyMergeModes
}

// GetKeyMergeModesOk returns a tuple with the KeyMergeModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transplant1) GetKeyMergeModesOk() ([]MergeOperationKeyMergeModesInner, bool) {
	if o == nil || IsNil(o.KeyMergeModes) {
		return nil, false
	}
	return o.KeyMergeModes, true
}

// HasKeyMergeModes returns a boolean if a field has been set.
func (o *Transplant1) HasKeyMergeModes() bool {
	if o != nil && !IsNil(o.KeyMergeModes) {
		return true
	}

	return false
}

// SetKeyMergeModes gets a reference to the given []MergeOperationKeyMergeModesInner and assigns it to the KeyMergeModes field.
func (o *Transplant1) SetKeyMergeModes(v []MergeOperationKeyMergeModesInner) {
	o.KeyMergeModes = v
}

// GetDefaultKeyMergeMode returns the DefaultKeyMergeMode field value if set, zero value otherwise.
func (o *Transplant1) GetDefaultKeyMergeMode() string {
	if o == nil || IsNil(o.DefaultKeyMergeMode) {
		var ret string
		return ret
	}
	return *o.DefaultKeyMergeMode
}

// GetDefaultKeyMergeModeOk returns a tuple with the DefaultKeyMergeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transplant1) GetDefaultKeyMergeModeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultKeyMergeMode) {
		return nil, false
	}
	return o.DefaultKeyMergeMode, true
}

// HasDefaultKeyMergeMode returns a boolean if a field has been set.
func (o *Transplant1) HasDefaultKeyMergeMode() bool {
	if o != nil && !IsNil(o.DefaultKeyMergeMode) {
		return true
	}

	return false
}

// SetDefaultKeyMergeMode gets a reference to the given string and assigns it to the DefaultKeyMergeMode field.
func (o *Transplant1) SetDefaultKeyMergeMode(v string) {
	o.DefaultKeyMergeMode = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *Transplant1) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transplant1) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *Transplant1) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *Transplant1) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFetchAdditionalInfo returns the FetchAdditionalInfo field value if set, zero value otherwise.
func (o *Transplant1) GetFetchAdditionalInfo() bool {
	if o == nil || IsNil(o.FetchAdditionalInfo) {
		var ret bool
		return ret
	}
	return *o.FetchAdditionalInfo
}

// GetFetchAdditionalInfoOk returns a tuple with the FetchAdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transplant1) GetFetchAdditionalInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.FetchAdditionalInfo) {
		return nil, false
	}
	return o.FetchAdditionalInfo, true
}

// HasFetchAdditionalInfo returns a boolean if a field has been set.
func (o *Transplant1) HasFetchAdditionalInfo() bool {
	if o != nil && !IsNil(o.FetchAdditionalInfo) {
		return true
	}

	return false
}

// SetFetchAdditionalInfo gets a reference to the given bool and assigns it to the FetchAdditionalInfo field.
func (o *Transplant1) SetFetchAdditionalInfo(v bool) {
	o.FetchAdditionalInfo = &v
}

// GetReturnConflictAsResult returns the ReturnConflictAsResult field value if set, zero value otherwise.
func (o *Transplant1) GetReturnConflictAsResult() bool {
	if o == nil || IsNil(o.ReturnConflictAsResult) {
		var ret bool
		return ret
	}
	return *o.ReturnConflictAsResult
}

// GetReturnConflictAsResultOk returns a tuple with the ReturnConflictAsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transplant1) GetReturnConflictAsResultOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnConflictAsResult) {
		return nil, false
	}
	return o.ReturnConflictAsResult, true
}

// HasReturnConflictAsResult returns a boolean if a field has been set.
func (o *Transplant1) HasReturnConflictAsResult() bool {
	if o != nil && !IsNil(o.ReturnConflictAsResult) {
		return true
	}

	return false
}

// SetReturnConflictAsResult gets a reference to the given bool and assigns it to the ReturnConflictAsResult field.
func (o *Transplant1) SetReturnConflictAsResult(v bool) {
	o.ReturnConflictAsResult = &v
}

func (o Transplant1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transplant1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromRefName"] = o.FromRefName
	toSerialize["hashesToTransplant"] = o.HashesToTransplant
	if !IsNil(o.KeyMergeModes) {
		toSerialize["keyMergeModes"] = o.KeyMergeModes
	}
	if !IsNil(o.DefaultKeyMergeMode) {
		toSerialize["defaultKeyMergeMode"] = o.DefaultKeyMergeMode
	}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !IsNil(o.FetchAdditionalInfo) {
		toSerialize["fetchAdditionalInfo"] = o.FetchAdditionalInfo
	}
	if !IsNil(o.ReturnConflictAsResult) {
		toSerialize["returnConflictAsResult"] = o.ReturnConflictAsResult
	}
	return toSerialize, nil
}

func (o *Transplant1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromRefName",
		"hashesToTransplant",
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

	varTransplant1 := _Transplant1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransplant1)

	if err != nil {
		return err
	}

	*o = Transplant1(varTransplant1)

	return err
}

type NullableTransplant1 struct {
	value *Transplant1
	isSet bool
}

func (v NullableTransplant1) Get() *Transplant1 {
	return v.value
}

func (v *NullableTransplant1) Set(val *Transplant1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransplant1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransplant1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransplant1(val *Transplant1) *NullableTransplant1 {
	return &NullableTransplant1{value: val, isSet: true}
}

func (v NullableTransplant1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransplant1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


