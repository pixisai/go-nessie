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

// checks if the ReferencesResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferencesResponseV2{}

// ReferencesResponseV2 struct for ReferencesResponseV2
type ReferencesResponseV2 struct {
	HasMore *bool `json:"hasMore,omitempty"`
	Token *string `json:"token,omitempty"`
	References []Reference2 `json:"references"`
}

type _ReferencesResponseV2 ReferencesResponseV2

// NewReferencesResponseV2 instantiates a new ReferencesResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencesResponseV2(references []Reference2) *ReferencesResponseV2 {
	this := ReferencesResponseV2{}
	this.References = references
	return &this
}

// NewReferencesResponseV2WithDefaults instantiates a new ReferencesResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencesResponseV2WithDefaults() *ReferencesResponseV2 {
	this := ReferencesResponseV2{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ReferencesResponseV2) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesResponseV2) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ReferencesResponseV2) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ReferencesResponseV2) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ReferencesResponseV2) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencesResponseV2) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ReferencesResponseV2) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ReferencesResponseV2) SetToken(v string) {
	o.Token = &v
}

// GetReferences returns the References field value
func (o *ReferencesResponseV2) GetReferences() []Reference2 {
	if o == nil {
		var ret []Reference2
		return ret
	}

	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *ReferencesResponseV2) GetReferencesOk() ([]Reference2, bool) {
	if o == nil {
		return nil, false
	}
	return o.References, true
}

// SetReferences sets field value
func (o *ReferencesResponseV2) SetReferences(v []Reference2) {
	o.References = v
}

func (o ReferencesResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencesResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["references"] = o.References
	return toSerialize, nil
}

func (o *ReferencesResponseV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"references",
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

	varReferencesResponseV2 := _ReferencesResponseV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReferencesResponseV2)

	if err != nil {
		return err
	}

	*o = ReferencesResponseV2(varReferencesResponseV2)

	return err
}

type NullableReferencesResponseV2 struct {
	value *ReferencesResponseV2
	isSet bool
}

func (v NullableReferencesResponseV2) Get() *ReferencesResponseV2 {
	return v.value
}

func (v *NullableReferencesResponseV2) Set(val *ReferencesResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencesResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencesResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencesResponseV2(val *ReferencesResponseV2) *NullableReferencesResponseV2 {
	return &NullableReferencesResponseV2{value: val, isSet: true}
}

func (v NullableReferencesResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencesResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


