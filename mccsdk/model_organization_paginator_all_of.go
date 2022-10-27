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

// OrganizationPaginatorAllOf struct for OrganizationPaginatorAllOf
type OrganizationPaginatorAllOf struct {
	Results []Organization `json:"results,omitempty"`
}

// NewOrganizationPaginatorAllOf instantiates a new OrganizationPaginatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPaginatorAllOf() *OrganizationPaginatorAllOf {
	this := OrganizationPaginatorAllOf{}
	return &this
}

// NewOrganizationPaginatorAllOfWithDefaults instantiates a new OrganizationPaginatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPaginatorAllOfWithDefaults() *OrganizationPaginatorAllOf {
	this := OrganizationPaginatorAllOf{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OrganizationPaginatorAllOf) GetResults() []Organization {
	if o == nil || o.Results == nil {
		var ret []Organization
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationPaginatorAllOf) GetResultsOk() ([]Organization, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OrganizationPaginatorAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Organization and assigns it to the Results field.
func (o *OrganizationPaginatorAllOf) SetResults(v []Organization) {
	o.Results = v
}

func (o OrganizationPaginatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationPaginatorAllOf struct {
	value *OrganizationPaginatorAllOf
	isSet bool
}

func (v NullableOrganizationPaginatorAllOf) Get() *OrganizationPaginatorAllOf {
	return v.value
}

func (v *NullableOrganizationPaginatorAllOf) Set(val *OrganizationPaginatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPaginatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPaginatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPaginatorAllOf(val *OrganizationPaginatorAllOf) *NullableOrganizationPaginatorAllOf {
	return &NullableOrganizationPaginatorAllOf{value: val, isSet: true}
}

func (v NullableOrganizationPaginatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPaginatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}