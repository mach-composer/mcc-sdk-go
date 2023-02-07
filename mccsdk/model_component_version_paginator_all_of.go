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

// ComponentVersionPaginatorAllOf struct for ComponentVersionPaginatorAllOf
type ComponentVersionPaginatorAllOf struct {
	Results []ComponentVersion `json:"results,omitempty"`
}

// NewComponentVersionPaginatorAllOf instantiates a new ComponentVersionPaginatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionPaginatorAllOf() *ComponentVersionPaginatorAllOf {
	this := ComponentVersionPaginatorAllOf{}
	return &this
}

// NewComponentVersionPaginatorAllOfWithDefaults instantiates a new ComponentVersionPaginatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentVersionPaginatorAllOfWithDefaults() *ComponentVersionPaginatorAllOf {
	this := ComponentVersionPaginatorAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ComponentVersionPaginatorAllOf) GetResults() []ComponentVersion {
	if o == nil || o.Results == nil {
		var ret []ComponentVersion
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionPaginatorAllOf) GetResultsOk() ([]ComponentVersion, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ComponentVersionPaginatorAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ComponentVersion and assigns it to the Results field.
func (o *ComponentVersionPaginatorAllOf) SetResults(v []ComponentVersion) {
	o.Results = v
}

func (o ComponentVersionPaginatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableComponentVersionPaginatorAllOf struct {
	value *ComponentVersionPaginatorAllOf
	isSet bool
}

func (v NullableComponentVersionPaginatorAllOf) Get() *ComponentVersionPaginatorAllOf {
	return v.value
}

func (v *NullableComponentVersionPaginatorAllOf) Set(val *ComponentVersionPaginatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentVersionPaginatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentVersionPaginatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentVersionPaginatorAllOf(val *ComponentVersionPaginatorAllOf) *NullableComponentVersionPaginatorAllOf {
	return &NullableComponentVersionPaginatorAllOf{value: val, isSet: true}
}

func (v NullableComponentVersionPaginatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentVersionPaginatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
