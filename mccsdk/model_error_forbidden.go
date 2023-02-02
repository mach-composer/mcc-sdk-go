/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.   All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"encoding/json"
)

// ErrorForbidden struct for ErrorForbidden
type ErrorForbidden struct {
	Status      *int    `json:"status,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Description *string `json:"description,omitempty"`
	Message     *string `json:"message,omitempty"`
}

// NewErrorForbidden instantiates a new ErrorForbidden object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorForbidden() *ErrorForbidden {
	this := ErrorForbidden{}
	return &this
}

// NewErrorForbiddenWithDefaults instantiates a new ErrorForbidden object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorForbiddenWithDefaults() *ErrorForbidden {
	this := ErrorForbidden{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ErrorForbidden) GetStatus() int {
	if o == nil || o.Status == nil {
		var ret int
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorForbidden) GetStatusOk() (*int, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ErrorForbidden) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int and assigns it to the Status field.
func (o *ErrorForbidden) SetStatus(v int) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ErrorForbidden) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorForbidden) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ErrorForbidden) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ErrorForbidden) SetSummary(v string) {
	o.Summary = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorForbidden) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorForbidden) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorForbidden) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorForbidden) SetDescription(v string) {
	o.Description = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorForbidden) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorForbidden) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorForbidden) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorForbidden) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorForbidden) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableErrorForbidden struct {
	value *ErrorForbidden
	isSet bool
}

func (v NullableErrorForbidden) Get() *ErrorForbidden {
	return v.value
}

func (v *NullableErrorForbidden) Set(val *ErrorForbidden) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorForbidden) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorForbidden) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorForbidden(val *ErrorForbidden) *NullableErrorForbidden {
	return &NullableErrorForbidden{value: val, isSet: true}
}

func (v NullableErrorForbidden) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorForbidden) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
