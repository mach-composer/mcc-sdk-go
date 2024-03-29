/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.  All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"encoding/json"
)

// ComponentAllOf struct for ComponentAllOf
type ComponentAllOf struct {
	// key of the component
	Key *string `json:"key,omitempty"`
}

// NewComponentAllOf instantiates a new ComponentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentAllOf() *ComponentAllOf {
	this := ComponentAllOf{}
	return &this
}

// NewComponentAllOfWithDefaults instantiates a new ComponentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentAllOfWithDefaults() *ComponentAllOf {
	this := ComponentAllOf{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ComponentAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ComponentAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ComponentAllOf) SetKey(v string) {
	o.Key = &v
}

func (o ComponentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableComponentAllOf struct {
	value *ComponentAllOf
	isSet bool
}

func (v NullableComponentAllOf) Get() *ComponentAllOf {
	return v.value
}

func (v *NullableComponentAllOf) Set(val *ComponentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentAllOf(val *ComponentAllOf) *NullableComponentAllOf {
	return &NullableComponentAllOf{value: val, isSet: true}
}

func (v NullableComponentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
