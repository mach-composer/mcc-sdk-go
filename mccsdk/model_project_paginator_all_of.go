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

// ProjectPaginatorAllOf struct for ProjectPaginatorAllOf
type ProjectPaginatorAllOf struct {
	Results []Project `json:"results,omitempty"`
}

// NewProjectPaginatorAllOf instantiates a new ProjectPaginatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPaginatorAllOf() *ProjectPaginatorAllOf {
	this := ProjectPaginatorAllOf{}
	return &this
}

// NewProjectPaginatorAllOfWithDefaults instantiates a new ProjectPaginatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPaginatorAllOfWithDefaults() *ProjectPaginatorAllOf {
	this := ProjectPaginatorAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProjectPaginatorAllOf) GetResults() []Project {
	if o == nil || o.Results == nil {
		var ret []Project
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginatorAllOf) GetResultsOk() ([]Project, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProjectPaginatorAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Project and assigns it to the Results field.
func (o *ProjectPaginatorAllOf) SetResults(v []Project) {
	o.Results = v
}

func (o ProjectPaginatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableProjectPaginatorAllOf struct {
	value *ProjectPaginatorAllOf
	isSet bool
}

func (v NullableProjectPaginatorAllOf) Get() *ProjectPaginatorAllOf {
	return v.value
}

func (v *NullableProjectPaginatorAllOf) Set(val *ProjectPaginatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPaginatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPaginatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPaginatorAllOf(val *ProjectPaginatorAllOf) *NullableProjectPaginatorAllOf {
	return &NullableProjectPaginatorAllOf{value: val, isSet: true}
}

func (v NullableProjectPaginatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPaginatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}