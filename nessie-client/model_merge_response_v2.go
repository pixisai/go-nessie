/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MergeResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeResponseV2{}

// MergeResponseV2 struct for MergeResponseV2
type MergeResponseV2 struct {
	ResultantTargetHash *string `json:"resultantTargetHash,omitempty"`
	CommonAncestor *string `json:"commonAncestor,omitempty"`
	TargetBranch *string `json:"targetBranch,omitempty"`
	EffectiveTargetHash *string `json:"effectiveTargetHash,omitempty"`
	ExpectedHash *string `json:"expectedHash,omitempty"`
	Details []MergePerContentKeyDetails1 `json:"details,omitempty"`
}

// NewMergeResponseV2 instantiates a new MergeResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeResponseV2() *MergeResponseV2 {
	this := MergeResponseV2{}
	return &this
}

// NewMergeResponseV2WithDefaults instantiates a new MergeResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeResponseV2WithDefaults() *MergeResponseV2 {
	this := MergeResponseV2{}
	return &this
}

// GetResultantTargetHash returns the ResultantTargetHash field value if set, zero value otherwise.
func (o *MergeResponseV2) GetResultantTargetHash() string {
	if o == nil || IsNil(o.ResultantTargetHash) {
		var ret string
		return ret
	}
	return *o.ResultantTargetHash
}

// GetResultantTargetHashOk returns a tuple with the ResultantTargetHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetResultantTargetHashOk() (*string, bool) {
	if o == nil || IsNil(o.ResultantTargetHash) {
		return nil, false
	}
	return o.ResultantTargetHash, true
}

// HasResultantTargetHash returns a boolean if a field has been set.
func (o *MergeResponseV2) HasResultantTargetHash() bool {
	if o != nil && !IsNil(o.ResultantTargetHash) {
		return true
	}

	return false
}

// SetResultantTargetHash gets a reference to the given string and assigns it to the ResultantTargetHash field.
func (o *MergeResponseV2) SetResultantTargetHash(v string) {
	o.ResultantTargetHash = &v
}

// GetCommonAncestor returns the CommonAncestor field value if set, zero value otherwise.
func (o *MergeResponseV2) GetCommonAncestor() string {
	if o == nil || IsNil(o.CommonAncestor) {
		var ret string
		return ret
	}
	return *o.CommonAncestor
}

// GetCommonAncestorOk returns a tuple with the CommonAncestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetCommonAncestorOk() (*string, bool) {
	if o == nil || IsNil(o.CommonAncestor) {
		return nil, false
	}
	return o.CommonAncestor, true
}

// HasCommonAncestor returns a boolean if a field has been set.
func (o *MergeResponseV2) HasCommonAncestor() bool {
	if o != nil && !IsNil(o.CommonAncestor) {
		return true
	}

	return false
}

// SetCommonAncestor gets a reference to the given string and assigns it to the CommonAncestor field.
func (o *MergeResponseV2) SetCommonAncestor(v string) {
	o.CommonAncestor = &v
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise.
func (o *MergeResponseV2) GetTargetBranch() string {
	if o == nil || IsNil(o.TargetBranch) {
		var ret string
		return ret
	}
	return *o.TargetBranch
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetTargetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.TargetBranch) {
		return nil, false
	}
	return o.TargetBranch, true
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *MergeResponseV2) HasTargetBranch() bool {
	if o != nil && !IsNil(o.TargetBranch) {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given string and assigns it to the TargetBranch field.
func (o *MergeResponseV2) SetTargetBranch(v string) {
	o.TargetBranch = &v
}

// GetEffectiveTargetHash returns the EffectiveTargetHash field value if set, zero value otherwise.
func (o *MergeResponseV2) GetEffectiveTargetHash() string {
	if o == nil || IsNil(o.EffectiveTargetHash) {
		var ret string
		return ret
	}
	return *o.EffectiveTargetHash
}

// GetEffectiveTargetHashOk returns a tuple with the EffectiveTargetHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetEffectiveTargetHashOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveTargetHash) {
		return nil, false
	}
	return o.EffectiveTargetHash, true
}

// HasEffectiveTargetHash returns a boolean if a field has been set.
func (o *MergeResponseV2) HasEffectiveTargetHash() bool {
	if o != nil && !IsNil(o.EffectiveTargetHash) {
		return true
	}

	return false
}

// SetEffectiveTargetHash gets a reference to the given string and assigns it to the EffectiveTargetHash field.
func (o *MergeResponseV2) SetEffectiveTargetHash(v string) {
	o.EffectiveTargetHash = &v
}

// GetExpectedHash returns the ExpectedHash field value if set, zero value otherwise.
func (o *MergeResponseV2) GetExpectedHash() string {
	if o == nil || IsNil(o.ExpectedHash) {
		var ret string
		return ret
	}
	return *o.ExpectedHash
}

// GetExpectedHashOk returns a tuple with the ExpectedHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetExpectedHashOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedHash) {
		return nil, false
	}
	return o.ExpectedHash, true
}

// HasExpectedHash returns a boolean if a field has been set.
func (o *MergeResponseV2) HasExpectedHash() bool {
	if o != nil && !IsNil(o.ExpectedHash) {
		return true
	}

	return false
}

// SetExpectedHash gets a reference to the given string and assigns it to the ExpectedHash field.
func (o *MergeResponseV2) SetExpectedHash(v string) {
	o.ExpectedHash = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MergeResponseV2) GetDetails() []MergePerContentKeyDetails1 {
	if o == nil || IsNil(o.Details) {
		var ret []MergePerContentKeyDetails1
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeResponseV2) GetDetailsOk() ([]MergePerContentKeyDetails1, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MergeResponseV2) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []MergePerContentKeyDetails1 and assigns it to the Details field.
func (o *MergeResponseV2) SetDetails(v []MergePerContentKeyDetails1) {
	o.Details = v
}

func (o MergeResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultantTargetHash) {
		toSerialize["resultantTargetHash"] = o.ResultantTargetHash
	}
	if !IsNil(o.CommonAncestor) {
		toSerialize["commonAncestor"] = o.CommonAncestor
	}
	if !IsNil(o.TargetBranch) {
		toSerialize["targetBranch"] = o.TargetBranch
	}
	if !IsNil(o.EffectiveTargetHash) {
		toSerialize["effectiveTargetHash"] = o.EffectiveTargetHash
	}
	if !IsNil(o.ExpectedHash) {
		toSerialize["expectedHash"] = o.ExpectedHash
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableMergeResponseV2 struct {
	value *MergeResponseV2
	isSet bool
}

func (v NullableMergeResponseV2) Get() *MergeResponseV2 {
	return v.value
}

func (v *NullableMergeResponseV2) Set(val *MergeResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeResponseV2(val *MergeResponseV2) *NullableMergeResponseV2 {
	return &NullableMergeResponseV2{value: val, isSet: true}
}

func (v NullableMergeResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


