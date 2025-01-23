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

// checks if the ReferenceCutoffPolicyV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceCutoffPolicyV2{}

// ReferenceCutoffPolicyV2 Cutoff policies per reference names. Supplied as a ref-name-pattern=policy tuple. Reference name patterns are regular expressions.
type ReferenceCutoffPolicyV2 struct {
	// Reference name patterns as a regular expressions.
	ReferenceNamePattern *string `json:"referenceNamePattern,omitempty"`
	// Policies can be one of: - number of commits as an integer value - a duration (see java.time.Duration) - an ISO instant - 'NONE', means everything's considered as live
	Policy *string `json:"policy,omitempty"`
}

// NewReferenceCutoffPolicyV2 instantiates a new ReferenceCutoffPolicyV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceCutoffPolicyV2() *ReferenceCutoffPolicyV2 {
	this := ReferenceCutoffPolicyV2{}
	return &this
}

// NewReferenceCutoffPolicyV2WithDefaults instantiates a new ReferenceCutoffPolicyV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceCutoffPolicyV2WithDefaults() *ReferenceCutoffPolicyV2 {
	this := ReferenceCutoffPolicyV2{}
	return &this
}

// GetReferenceNamePattern returns the ReferenceNamePattern field value if set, zero value otherwise.
func (o *ReferenceCutoffPolicyV2) GetReferenceNamePattern() string {
	if o == nil || IsNil(o.ReferenceNamePattern) {
		var ret string
		return ret
	}
	return *o.ReferenceNamePattern
}

// GetReferenceNamePatternOk returns a tuple with the ReferenceNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceCutoffPolicyV2) GetReferenceNamePatternOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNamePattern) {
		return nil, false
	}
	return o.ReferenceNamePattern, true
}

// HasReferenceNamePattern returns a boolean if a field has been set.
func (o *ReferenceCutoffPolicyV2) HasReferenceNamePattern() bool {
	if o != nil && !IsNil(o.ReferenceNamePattern) {
		return true
	}

	return false
}

// SetReferenceNamePattern gets a reference to the given string and assigns it to the ReferenceNamePattern field.
func (o *ReferenceCutoffPolicyV2) SetReferenceNamePattern(v string) {
	o.ReferenceNamePattern = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *ReferenceCutoffPolicyV2) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceCutoffPolicyV2) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *ReferenceCutoffPolicyV2) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *ReferenceCutoffPolicyV2) SetPolicy(v string) {
	o.Policy = &v
}

func (o ReferenceCutoffPolicyV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceCutoffPolicyV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceNamePattern) {
		toSerialize["referenceNamePattern"] = o.ReferenceNamePattern
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullableReferenceCutoffPolicyV2 struct {
	value *ReferenceCutoffPolicyV2
	isSet bool
}

func (v NullableReferenceCutoffPolicyV2) Get() *ReferenceCutoffPolicyV2 {
	return v.value
}

func (v *NullableReferenceCutoffPolicyV2) Set(val *ReferenceCutoffPolicyV2) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceCutoffPolicyV2) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceCutoffPolicyV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceCutoffPolicyV2(val *ReferenceCutoffPolicyV2) *NullableReferenceCutoffPolicyV2 {
	return &NullableReferenceCutoffPolicyV2{value: val, isSet: true}
}

func (v NullableReferenceCutoffPolicyV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceCutoffPolicyV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


