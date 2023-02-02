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

// ComponentPaginator struct for ComponentPaginator
type ComponentPaginator struct {
	// Number of items in the current page
	Count float32 `json:"count"`
	// Total number of items found
	Total   float32     `json:"total"`
	Offset  float32     `json:"offset"`
	Limit   float32     `json:"limit"`
	Results []Component `json:"results"`
}

// NewComponentPaginator instantiates a new ComponentPaginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentPaginator(count float32, total float32, offset float32, limit float32, results []Component) *ComponentPaginator {
	this := ComponentPaginator{}
	this.Count = count
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	this.Results = results
	return &this
}

// NewComponentPaginatorWithDefaults instantiates a new ComponentPaginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentPaginatorWithDefaults() *ComponentPaginator {
	this := ComponentPaginator{}
	var offset float32 = 0
	this.Offset = offset
	var limit float32 = 20
	this.Limit = limit
	return &this
}

// GetCount returns the Count field value
func (o *ComponentPaginator) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ComponentPaginator) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ComponentPaginator) SetCount(v float32) {
	o.Count = v
}

// GetTotal returns the Total field value
func (o *ComponentPaginator) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ComponentPaginator) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ComponentPaginator) SetTotal(v float32) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *ComponentPaginator) GetOffset() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ComponentPaginator) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ComponentPaginator) SetOffset(v float32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *ComponentPaginator) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ComponentPaginator) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ComponentPaginator) SetLimit(v float32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *ComponentPaginator) GetResults() []Component {
	if o == nil {
		var ret []Component
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ComponentPaginator) GetResultsOk() ([]Component, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ComponentPaginator) SetResults(v []Component) {
	o.Results = v
}

func (o ComponentPaginator) MarshalJSON() ([]byte, error) {
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

type NullableComponentPaginator struct {
	value *ComponentPaginator
	isSet bool
}

func (v NullableComponentPaginator) Get() *ComponentPaginator {
	return v.value
}

func (v *NullableComponentPaginator) Set(val *ComponentPaginator) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentPaginator) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentPaginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentPaginator(val *ComponentPaginator) *NullableComponentPaginator {
	return &NullableComponentPaginator{value: val, isSet: true}
}

func (v NullableComponentPaginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentPaginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
