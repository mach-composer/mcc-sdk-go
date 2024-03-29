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

// OrganizationUserPaginator struct for OrganizationUserPaginator
type OrganizationUserPaginator struct {
	// Number of items in the current page
	Count float32 `json:"count"`
	// Total number of items found
	Total   float32            `json:"total"`
	Offset  float32            `json:"offset"`
	Limit   float32            `json:"limit"`
	Results []OrganizationUser `json:"results"`
}

// NewOrganizationUserPaginator instantiates a new OrganizationUserPaginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserPaginator(count float32, total float32, offset float32, limit float32, results []OrganizationUser) *OrganizationUserPaginator {
	this := OrganizationUserPaginator{}
	this.Count = count
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	this.Results = results
	return &this
}

// NewOrganizationUserPaginatorWithDefaults instantiates a new OrganizationUserPaginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserPaginatorWithDefaults() *OrganizationUserPaginator {
	this := OrganizationUserPaginator{}
	var offset float32 = 0
	this.Offset = offset
	var limit float32 = 20
	this.Limit = limit
	return &this
}

// GetCount returns the Count field value
func (o *OrganizationUserPaginator) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *OrganizationUserPaginator) SetCount(v float32) {
	o.Count = v
}

// GetTotal returns the Total field value
func (o *OrganizationUserPaginator) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OrganizationUserPaginator) SetTotal(v float32) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *OrganizationUserPaginator) GetOffset() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *OrganizationUserPaginator) SetOffset(v float32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *OrganizationUserPaginator) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *OrganizationUserPaginator) SetLimit(v float32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *OrganizationUserPaginator) GetResults() []OrganizationUser {
	if o == nil {
		var ret []OrganizationUser
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetResultsOk() ([]OrganizationUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *OrganizationUserPaginator) SetResults(v []OrganizationUser) {
	o.Results = v
}

func (o OrganizationUserPaginator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserPaginator struct {
	value *OrganizationUserPaginator
	isSet bool
}

func (v NullableOrganizationUserPaginator) Get() *OrganizationUserPaginator {
	return v.value
}

func (v *NullableOrganizationUserPaginator) Set(val *OrganizationUserPaginator) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserPaginator) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserPaginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserPaginator(val *OrganizationUserPaginator) *NullableOrganizationUserPaginator {
	return &NullableOrganizationUserPaginator{value: val, isSet: true}
}

func (v NullableOrganizationUserPaginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserPaginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
