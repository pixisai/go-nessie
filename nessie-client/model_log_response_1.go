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

// checks if the LogResponse1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogResponse1{}

// LogResponse1 struct for LogResponse1
type LogResponse1 struct {
	HasMore *bool `json:"hasMore,omitempty"`
	Token *string `json:"token,omitempty"`
	LogEntries []LogEntry1 `json:"logEntries"`
}

type _LogResponse1 LogResponse1

// NewLogResponse1 instantiates a new LogResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResponse1(logEntries []LogEntry1) *LogResponse1 {
	this := LogResponse1{}
	this.LogEntries = logEntries
	return &this
}

// NewLogResponse1WithDefaults instantiates a new LogResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResponse1WithDefaults() *LogResponse1 {
	this := LogResponse1{}
	return &this
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *LogResponse1) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse1) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *LogResponse1) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *LogResponse1) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LogResponse1) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse1) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LogResponse1) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LogResponse1) SetToken(v string) {
	o.Token = &v
}

// GetLogEntries returns the LogEntries field value
func (o *LogResponse1) GetLogEntries() []LogEntry1 {
	if o == nil {
		var ret []LogEntry1
		return ret
	}

	return o.LogEntries
}

// GetLogEntriesOk returns a tuple with the LogEntries field value
// and a boolean to check if the value has been set.
func (o *LogResponse1) GetLogEntriesOk() ([]LogEntry1, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogEntries, true
}

// SetLogEntries sets field value
func (o *LogResponse1) SetLogEntries(v []LogEntry1) {
	o.LogEntries = v
}

func (o LogResponse1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["logEntries"] = o.LogEntries
	return toSerialize, nil
}

func (o *LogResponse1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logEntries",
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

	varLogResponse1 := _LogResponse1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogResponse1)

	if err != nil {
		return err
	}

	*o = LogResponse1(varLogResponse1)

	return err
}

type NullableLogResponse1 struct {
	value *LogResponse1
	isSet bool
}

func (v NullableLogResponse1) Get() *LogResponse1 {
	return v.value
}

func (v *NullableLogResponse1) Set(val *LogResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResponse1(val *LogResponse1) *NullableLogResponse1 {
	return &NullableLogResponse1{value: val, isSet: true}
}

func (v NullableLogResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


