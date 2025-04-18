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

// checks if the ErrorUnauthorized type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorUnauthorized{}

// ErrorUnauthorized Unauthorized
type ErrorUnauthorized struct {
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorUnauthorized ErrorUnauthorized

// NewErrorUnauthorized instantiates a new ErrorUnauthorized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorUnauthorized() *ErrorUnauthorized {
	this := ErrorUnauthorized{}
	return &this
}

// NewErrorUnauthorizedWithDefaults instantiates a new ErrorUnauthorized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorUnauthorizedWithDefaults() *ErrorUnauthorized {
	this := ErrorUnauthorized{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorUnauthorized) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorUnauthorized) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorUnauthorized) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorUnauthorized) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorUnauthorized) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorUnauthorized) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorUnauthorized) UnmarshalJSON(data []byte) (err error) {
	varErrorUnauthorized := _ErrorUnauthorized{}

	err = json.Unmarshal(data, &varErrorUnauthorized)

	if err != nil {
		return err
	}

	*o = ErrorUnauthorized(varErrorUnauthorized)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorUnauthorized struct {
	value *ErrorUnauthorized
	isSet bool
}

func (v NullableErrorUnauthorized) Get() *ErrorUnauthorized {
	return v.value
}

func (v *NullableErrorUnauthorized) Set(val *ErrorUnauthorized) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorUnauthorized) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorUnauthorized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorUnauthorized(val *ErrorUnauthorized) *NullableErrorUnauthorized {
	return &NullableErrorUnauthorized{value: val, isSet: true}
}

func (v NullableErrorUnauthorized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorUnauthorized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
