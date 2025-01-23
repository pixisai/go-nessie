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

// checks if the GetAllReferences200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllReferences200Response{}

// GetAllReferences200Response struct for GetAllReferences200Response
type GetAllReferences200Response struct {
	HasMore *bool `json:"hasMore,omitempty"`
	Token *string `json:"token,omitempty"`
	References []Reference1 `json:"references"`
}

type _GetAllReferences200Response GetAllReferences200Response

// NewGetAllReferences200Response instantiates a new GetAllReferences200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllReferences200Response(references []Reference1) *GetAllReferences200Response {
	this := GetAllReferences200Response{}
	this.References = references
	return &this
}

// NewGetAllReferences200ResponseWithDefaults instantiates a new GetAllReferences200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllReferences200ResponseWithDefaults() *GetAllReferences200Response {
	this := GetAllReferences200Response{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *GetAllReferences200Response) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllReferences200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *GetAllReferences200Response) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *GetAllReferences200Response) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetAllReferences200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllReferences200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetAllReferences200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetAllReferences200Response) SetToken(v string) {
	o.Token = &v
}

// GetReferences returns the References field value
func (o *GetAllReferences200Response) GetReferences() []Reference1 {
	if o == nil {
		var ret []Reference1
		return ret
	}

	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *GetAllReferences200Response) GetReferencesOk() ([]Reference1, bool) {
	if o == nil {
		return nil, false
	}
	return o.References, true
}

// SetReferences sets field value
func (o *GetAllReferences200Response) SetReferences(v []Reference1) {
	o.References = v
}

func (o GetAllReferences200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllReferences200Response) ToMap() (map[string]interface{}, error) {
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

func (o *GetAllReferences200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetAllReferences200Response := _GetAllReferences200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAllReferences200Response)

	if err != nil {
		return err
	}

	*o = GetAllReferences200Response(varGetAllReferences200Response)

	return err
}

type NullableGetAllReferences200Response struct {
	value *GetAllReferences200Response
	isSet bool
}

func (v NullableGetAllReferences200Response) Get() *GetAllReferences200Response {
	return v.value
}

func (v *NullableGetAllReferences200Response) Set(val *GetAllReferences200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllReferences200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllReferences200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllReferences200Response(val *GetAllReferences200Response) *NullableGetAllReferences200Response {
	return &NullableGetAllReferences200Response{value: val, isSet: true}
}

func (v NullableGetAllReferences200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllReferences200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


