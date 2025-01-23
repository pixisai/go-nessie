/*
Nessie API

Transactional Catalog for Data Lakes  * Git-inspired data version control * Cross-table transactions and visibility * Works with Apache Iceberg tables

API version: 0.102.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the NessieConfiguration2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NessieConfiguration2{}

// NessieConfiguration2 Configuration object to tell a client how a server is configured.
type NessieConfiguration2 struct {
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	MinSupportedApiVersion *int32 `json:"minSupportedApiVersion,omitempty"`
	MaxSupportedApiVersion *int32 `json:"maxSupportedApiVersion,omitempty"`
	ActualApiVersion *int32 `json:"actualApiVersion,omitempty"`
	// Semver version representing the behavior of the Nessie server.  Additional functionality might be added to Nessie servers within a \"spec major version\" in a non-breaking way. Clients are encouraged to check the spec version when using such added functionality.
	SpecVersion *string `json:"specVersion,omitempty"`
	NoAncestorHash *string `json:"noAncestorHash,omitempty"`
	RepositoryCreationTimestamp *time.Time `json:"repositoryCreationTimestamp,omitempty"`
	OldestPossibleCommitTimestamp *time.Time `json:"oldestPossibleCommitTimestamp,omitempty"`
	AdditionalPropertiesField map[string]string `json:"additionalProperties,omitempty"`
}

// NewNessieConfiguration2 instantiates a new NessieConfiguration2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNessieConfiguration2() *NessieConfiguration2 {
	this := NessieConfiguration2{}
	return &this
}

// NewNessieConfiguration2WithDefaults instantiates a new NessieConfiguration2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNessieConfiguration2WithDefaults() *NessieConfiguration2 {
	this := NessieConfiguration2{}
	return &this
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *NessieConfiguration2) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetMinSupportedApiVersion returns the MinSupportedApiVersion field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetMinSupportedApiVersion() int32 {
	if o == nil || IsNil(o.MinSupportedApiVersion) {
		var ret int32
		return ret
	}
	return *o.MinSupportedApiVersion
}

// GetMinSupportedApiVersionOk returns a tuple with the MinSupportedApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetMinSupportedApiVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSupportedApiVersion) {
		return nil, false
	}
	return o.MinSupportedApiVersion, true
}

// HasMinSupportedApiVersion returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasMinSupportedApiVersion() bool {
	if o != nil && !IsNil(o.MinSupportedApiVersion) {
		return true
	}

	return false
}

// SetMinSupportedApiVersion gets a reference to the given int32 and assigns it to the MinSupportedApiVersion field.
func (o *NessieConfiguration2) SetMinSupportedApiVersion(v int32) {
	o.MinSupportedApiVersion = &v
}

// GetMaxSupportedApiVersion returns the MaxSupportedApiVersion field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetMaxSupportedApiVersion() int32 {
	if o == nil || IsNil(o.MaxSupportedApiVersion) {
		var ret int32
		return ret
	}
	return *o.MaxSupportedApiVersion
}

// GetMaxSupportedApiVersionOk returns a tuple with the MaxSupportedApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetMaxSupportedApiVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSupportedApiVersion) {
		return nil, false
	}
	return o.MaxSupportedApiVersion, true
}

// HasMaxSupportedApiVersion returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasMaxSupportedApiVersion() bool {
	if o != nil && !IsNil(o.MaxSupportedApiVersion) {
		return true
	}

	return false
}

// SetMaxSupportedApiVersion gets a reference to the given int32 and assigns it to the MaxSupportedApiVersion field.
func (o *NessieConfiguration2) SetMaxSupportedApiVersion(v int32) {
	o.MaxSupportedApiVersion = &v
}

// GetActualApiVersion returns the ActualApiVersion field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetActualApiVersion() int32 {
	if o == nil || IsNil(o.ActualApiVersion) {
		var ret int32
		return ret
	}
	return *o.ActualApiVersion
}

// GetActualApiVersionOk returns a tuple with the ActualApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetActualApiVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.ActualApiVersion) {
		return nil, false
	}
	return o.ActualApiVersion, true
}

// HasActualApiVersion returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasActualApiVersion() bool {
	if o != nil && !IsNil(o.ActualApiVersion) {
		return true
	}

	return false
}

// SetActualApiVersion gets a reference to the given int32 and assigns it to the ActualApiVersion field.
func (o *NessieConfiguration2) SetActualApiVersion(v int32) {
	o.ActualApiVersion = &v
}

// GetSpecVersion returns the SpecVersion field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetSpecVersion() string {
	if o == nil || IsNil(o.SpecVersion) {
		var ret string
		return ret
	}
	return *o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetSpecVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SpecVersion) {
		return nil, false
	}
	return o.SpecVersion, true
}

// HasSpecVersion returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasSpecVersion() bool {
	if o != nil && !IsNil(o.SpecVersion) {
		return true
	}

	return false
}

// SetSpecVersion gets a reference to the given string and assigns it to the SpecVersion field.
func (o *NessieConfiguration2) SetSpecVersion(v string) {
	o.SpecVersion = &v
}

// GetNoAncestorHash returns the NoAncestorHash field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetNoAncestorHash() string {
	if o == nil || IsNil(o.NoAncestorHash) {
		var ret string
		return ret
	}
	return *o.NoAncestorHash
}

// GetNoAncestorHashOk returns a tuple with the NoAncestorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetNoAncestorHashOk() (*string, bool) {
	if o == nil || IsNil(o.NoAncestorHash) {
		return nil, false
	}
	return o.NoAncestorHash, true
}

// HasNoAncestorHash returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasNoAncestorHash() bool {
	if o != nil && !IsNil(o.NoAncestorHash) {
		return true
	}

	return false
}

// SetNoAncestorHash gets a reference to the given string and assigns it to the NoAncestorHash field.
func (o *NessieConfiguration2) SetNoAncestorHash(v string) {
	o.NoAncestorHash = &v
}

// GetRepositoryCreationTimestamp returns the RepositoryCreationTimestamp field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetRepositoryCreationTimestamp() time.Time {
	if o == nil || IsNil(o.RepositoryCreationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.RepositoryCreationTimestamp
}

// GetRepositoryCreationTimestampOk returns a tuple with the RepositoryCreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetRepositoryCreationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RepositoryCreationTimestamp) {
		return nil, false
	}
	return o.RepositoryCreationTimestamp, true
}

// HasRepositoryCreationTimestamp returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasRepositoryCreationTimestamp() bool {
	if o != nil && !IsNil(o.RepositoryCreationTimestamp) {
		return true
	}

	return false
}

// SetRepositoryCreationTimestamp gets a reference to the given time.Time and assigns it to the RepositoryCreationTimestamp field.
func (o *NessieConfiguration2) SetRepositoryCreationTimestamp(v time.Time) {
	o.RepositoryCreationTimestamp = &v
}

// GetOldestPossibleCommitTimestamp returns the OldestPossibleCommitTimestamp field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetOldestPossibleCommitTimestamp() time.Time {
	if o == nil || IsNil(o.OldestPossibleCommitTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.OldestPossibleCommitTimestamp
}

// GetOldestPossibleCommitTimestampOk returns a tuple with the OldestPossibleCommitTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetOldestPossibleCommitTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OldestPossibleCommitTimestamp) {
		return nil, false
	}
	return o.OldestPossibleCommitTimestamp, true
}

// HasOldestPossibleCommitTimestamp returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasOldestPossibleCommitTimestamp() bool {
	if o != nil && !IsNil(o.OldestPossibleCommitTimestamp) {
		return true
	}

	return false
}

// SetOldestPossibleCommitTimestamp gets a reference to the given time.Time and assigns it to the OldestPossibleCommitTimestamp field.
func (o *NessieConfiguration2) SetOldestPossibleCommitTimestamp(v time.Time) {
	o.OldestPossibleCommitTimestamp = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *NessieConfiguration2) GetAdditionalPropertiesField() map[string]string {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		var ret map[string]string
		return ret
	}
	return o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NessieConfiguration2) GetAdditionalPropertiesFieldOk() (map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalPropertiesField) {
		return map[string]string{}, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *NessieConfiguration2) HasAdditionalPropertiesField() bool {
	if o != nil && !IsNil(o.AdditionalPropertiesField) {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]string and assigns it to the AdditionalPropertiesField field.
func (o *NessieConfiguration2) SetAdditionalPropertiesField(v map[string]string) {
	o.AdditionalPropertiesField = v
}

func (o NessieConfiguration2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NessieConfiguration2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultBranch) {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if !IsNil(o.MinSupportedApiVersion) {
		toSerialize["minSupportedApiVersion"] = o.MinSupportedApiVersion
	}
	if !IsNil(o.MaxSupportedApiVersion) {
		toSerialize["maxSupportedApiVersion"] = o.MaxSupportedApiVersion
	}
	if !IsNil(o.ActualApiVersion) {
		toSerialize["actualApiVersion"] = o.ActualApiVersion
	}
	if !IsNil(o.SpecVersion) {
		toSerialize["specVersion"] = o.SpecVersion
	}
	if !IsNil(o.NoAncestorHash) {
		toSerialize["noAncestorHash"] = o.NoAncestorHash
	}
	if !IsNil(o.RepositoryCreationTimestamp) {
		toSerialize["repositoryCreationTimestamp"] = o.RepositoryCreationTimestamp
	}
	if !IsNil(o.OldestPossibleCommitTimestamp) {
		toSerialize["oldestPossibleCommitTimestamp"] = o.OldestPossibleCommitTimestamp
	}
	if !IsNil(o.AdditionalPropertiesField) {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return toSerialize, nil
}

type NullableNessieConfiguration2 struct {
	value *NessieConfiguration2
	isSet bool
}

func (v NullableNessieConfiguration2) Get() *NessieConfiguration2 {
	return v.value
}

func (v *NullableNessieConfiguration2) Set(val *NessieConfiguration2) {
	v.value = val
	v.isSet = true
}

func (v NullableNessieConfiguration2) IsSet() bool {
	return v.isSet
}

func (v *NullableNessieConfiguration2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNessieConfiguration2(val *NessieConfiguration2) *NullableNessieConfiguration2 {
	return &NullableNessieConfiguration2{value: val, isSet: true}
}

func (v NullableNessieConfiguration2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNessieConfiguration2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


