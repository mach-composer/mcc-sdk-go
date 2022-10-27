/*
MCC Public API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"encoding/json"
)

// ComponentPaginatorAllOf struct for ComponentPaginatorAllOf
type ComponentPaginatorAllOf struct {
	Results []Component `json:"results,omitempty"`
}

// NewComponentPaginatorAllOf instantiates a new ComponentPaginatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentPaginatorAllOf() *ComponentPaginatorAllOf {
	this := ComponentPaginatorAllOf{}
	return &this
}

// NewComponentPaginatorAllOfWithDefaults instantiates a new ComponentPaginatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentPaginatorAllOfWithDefaults() *ComponentPaginatorAllOf {
	this := ComponentPaginatorAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ComponentPaginatorAllOf) GetResults() []Component {
	if o == nil || o.Results == nil {
		var ret []Component
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPaginatorAllOf) GetResultsOk() ([]Component, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ComponentPaginatorAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Component and assigns it to the Results field.
func (o *ComponentPaginatorAllOf) SetResults(v []Component) {
	o.Results = v
}

func (o ComponentPaginatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableComponentPaginatorAllOf struct {
	value *ComponentPaginatorAllOf
	isSet bool
}

func (v NullableComponentPaginatorAllOf) Get() *ComponentPaginatorAllOf {
	return v.value
}

func (v *NullableComponentPaginatorAllOf) Set(val *ComponentPaginatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentPaginatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentPaginatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentPaginatorAllOf(val *ComponentPaginatorAllOf) *NullableComponentPaginatorAllOf {
	return &NullableComponentPaginatorAllOf{value: val, isSet: true}
}

func (v NullableComponentPaginatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentPaginatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}