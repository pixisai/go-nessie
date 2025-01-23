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

// checks if the RepositoryConfigV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryConfigV2{}

// RepositoryConfigV2 struct for RepositoryConfigV2
type RepositoryConfigV2 struct {
	// The default cutoff policy. Policies can be one of: - number of commits as an integer value - a duration (see java.time.Duration) - an ISO instant - 'NONE', means everything's considered as live
	DefaultCutoffPolicy *string `json:"defaultCutoffPolicy,omitempty" validate:"regexp=NONE|^[1-9]\\\\d{0,10}|([-+]?)P(?:([-+]?[0-9]+)D)?(T(?:([-+]?[0-9]+)H)?(?:([-+]?[0-9]+)M)?(?:([-+]?[0-9]+)(?:[.,]([0-9]{0,9}))?S)?)?|^(?:[1-9]\\\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\\\d|2[0-3]):[0-5]\\\\d:[0-5]\\\\d(?:\\\\.\\\\d{1,9})?(?:Z|[+-][01]\\\\d:[0-5]\\\\d)$"`
	PerRefCutoffPolicies []ReferencesCutoffPolicy `json:"perRefCutoffPolicies,omitempty"`
	// Files that have been created after 'gc-start-time - new-files-grace-period' are not being deleted.
	NewFilesGracePeriod *string `json:"newFilesGracePeriod,omitempty"`
	ExpectedFileCountPerContent *int32 `json:"expectedFileCountPerContent,omitempty"`
}

// NewRepositoryConfigV2 instantiates a new RepositoryConfigV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryConfigV2() *RepositoryConfigV2 {
	this := RepositoryConfigV2{}
	return &this
}

// NewRepositoryConfigV2WithDefaults instantiates a new RepositoryConfigV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryConfigV2WithDefaults() *RepositoryConfigV2 {
	this := RepositoryConfigV2{}
	return &this
}

// GetDefaultCutoffPolicy returns the DefaultCutoffPolicy field value if set, zero value otherwise.
func (o *RepositoryConfigV2) GetDefaultCutoffPolicy() string {
	if o == nil || IsNil(o.DefaultCutoffPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultCutoffPolicy
}

// GetDefaultCutoffPolicyOk returns a tuple with the DefaultCutoffPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryConfigV2) GetDefaultCutoffPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultCutoffPolicy) {
		return nil, false
	}
	return o.DefaultCutoffPolicy, true
}

// HasDefaultCutoffPolicy returns a boolean if a field has been set.
func (o *RepositoryConfigV2) HasDefaultCutoffPolicy() bool {
	if o != nil && !IsNil(o.DefaultCutoffPolicy) {
		return true
	}

	return false
}

// SetDefaultCutoffPolicy gets a reference to the given string and assigns it to the DefaultCutoffPolicy field.
func (o *RepositoryConfigV2) SetDefaultCutoffPolicy(v string) {
	o.DefaultCutoffPolicy = &v
}

// GetPerRefCutoffPolicies returns the PerRefCutoffPolicies field value if set, zero value otherwise.
func (o *RepositoryConfigV2) GetPerRefCutoffPolicies() []ReferencesCutoffPolicy {
	if o == nil || IsNil(o.PerRefCutoffPolicies) {
		var ret []ReferencesCutoffPolicy
		return ret
	}
	return o.PerRefCutoffPolicies
}

// GetPerRefCutoffPoliciesOk returns a tuple with the PerRefCutoffPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryConfigV2) GetPerRefCutoffPoliciesOk() ([]ReferencesCutoffPolicy, bool) {
	if o == nil || IsNil(o.PerRefCutoffPolicies) {
		return nil, false
	}
	return o.PerRefCutoffPolicies, true
}

// HasPerRefCutoffPolicies returns a boolean if a field has been set.
func (o *RepositoryConfigV2) HasPerRefCutoffPolicies() bool {
	if o != nil && !IsNil(o.PerRefCutoffPolicies) {
		return true
	}

	return false
}

// SetPerRefCutoffPolicies gets a reference to the given []ReferencesCutoffPolicy and assigns it to the PerRefCutoffPolicies field.
func (o *RepositoryConfigV2) SetPerRefCutoffPolicies(v []ReferencesCutoffPolicy) {
	o.PerRefCutoffPolicies = v
}

// GetNewFilesGracePeriod returns the NewFilesGracePeriod field value if set, zero value otherwise.
func (o *RepositoryConfigV2) GetNewFilesGracePeriod() string {
	if o == nil || IsNil(o.NewFilesGracePeriod) {
		var ret string
		return ret
	}
	return *o.NewFilesGracePeriod
}

// GetNewFilesGracePeriodOk returns a tuple with the NewFilesGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryConfigV2) GetNewFilesGracePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.NewFilesGracePeriod) {
		return nil, false
	}
	return o.NewFilesGracePeriod, true
}

// HasNewFilesGracePeriod returns a boolean if a field has been set.
func (o *RepositoryConfigV2) HasNewFilesGracePeriod() bool {
	if o != nil && !IsNil(o.NewFilesGracePeriod) {
		return true
	}

	return false
}

// SetNewFilesGracePeriod gets a reference to the given string and assigns it to the NewFilesGracePeriod field.
func (o *RepositoryConfigV2) SetNewFilesGracePeriod(v string) {
	o.NewFilesGracePeriod = &v
}

// GetExpectedFileCountPerContent returns the ExpectedFileCountPerContent field value if set, zero value otherwise.
func (o *RepositoryConfigV2) GetExpectedFileCountPerContent() int32 {
	if o == nil || IsNil(o.ExpectedFileCountPerContent) {
		var ret int32
		return ret
	}
	return *o.ExpectedFileCountPerContent
}

// GetExpectedFileCountPerContentOk returns a tuple with the ExpectedFileCountPerContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryConfigV2) GetExpectedFileCountPerContentOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpectedFileCountPerContent) {
		return nil, false
	}
	return o.ExpectedFileCountPerContent, true
}

// HasExpectedFileCountPerContent returns a boolean if a field has been set.
func (o *RepositoryConfigV2) HasExpectedFileCountPerContent() bool {
	if o != nil && !IsNil(o.ExpectedFileCountPerContent) {
		return true
	}

	return false
}

// SetExpectedFileCountPerContent gets a reference to the given int32 and assigns it to the ExpectedFileCountPerContent field.
func (o *RepositoryConfigV2) SetExpectedFileCountPerContent(v int32) {
	o.ExpectedFileCountPerContent = &v
}

func (o RepositoryConfigV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryConfigV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultCutoffPolicy) {
		toSerialize["defaultCutoffPolicy"] = o.DefaultCutoffPolicy
	}
	if !IsNil(o.PerRefCutoffPolicies) {
		toSerialize["perRefCutoffPolicies"] = o.PerRefCutoffPolicies
	}
	if !IsNil(o.NewFilesGracePeriod) {
		toSerialize["newFilesGracePeriod"] = o.NewFilesGracePeriod
	}
	if !IsNil(o.ExpectedFileCountPerContent) {
		toSerialize["expectedFileCountPerContent"] = o.ExpectedFileCountPerContent
	}
	return toSerialize, nil
}

type NullableRepositoryConfigV2 struct {
	value *RepositoryConfigV2
	isSet bool
}

func (v NullableRepositoryConfigV2) Get() *RepositoryConfigV2 {
	return v.value
}

func (v *NullableRepositoryConfigV2) Set(val *RepositoryConfigV2) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryConfigV2) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryConfigV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryConfigV2(val *RepositoryConfigV2) *NullableRepositoryConfigV2 {
	return &NullableRepositoryConfigV2{value: val, isSet: true}
}

func (v NullableRepositoryConfigV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryConfigV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


