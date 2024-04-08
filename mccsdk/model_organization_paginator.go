/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.  All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OrganizationPaginator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationPaginator{}

// OrganizationPaginator struct for OrganizationPaginator
type OrganizationPaginator struct {
	// Number of items in the current page
	Count int32 `json:"count"`
	// Total number of items found
	Total   int64          `json:"total"`
	Offset  int32          `json:"offset"`
	Limit   int32          `json:"limit"`
	Results []Organization `json:"results"`
}

type _OrganizationPaginator OrganizationPaginator

// NewOrganizationPaginator instantiates a new OrganizationPaginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationPaginator(count int32, total int64, offset int32, limit int32, results []Organization) *OrganizationPaginator {
	this := OrganizationPaginator{}
	this.Count = count
	this.Total = total
	this.Offset = offset
	this.Limit = limit
	this.Results = results
	return &this
}

// NewOrganizationPaginatorWithDefaults instantiates a new OrganizationPaginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationPaginatorWithDefaults() *OrganizationPaginator {
	this := OrganizationPaginator{}
	var offset int32 = 0
	this.Offset = offset
	var limit int32 = 20
	this.Limit = limit
	return &this
}

// GetCount returns the Count field value
func (o *OrganizationPaginator) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrganizationPaginator) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *OrganizationPaginator) SetCount(v int32) {
	o.Count = v
}

// GetTotal returns the Total field value
func (o *OrganizationPaginator) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OrganizationPaginator) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OrganizationPaginator) SetTotal(v int64) {
	o.Total = v
}

// GetOffset returns the Offset field value
func (o *OrganizationPaginator) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *OrganizationPaginator) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *OrganizationPaginator) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *OrganizationPaginator) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *OrganizationPaginator) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *OrganizationPaginator) SetLimit(v int32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *OrganizationPaginator) GetResults() []Organization {
	if o == nil {
		var ret []Organization
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OrganizationPaginator) GetResultsOk() ([]Organization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *OrganizationPaginator) SetResults(v []Organization) {
	o.Results = v
}

func (o OrganizationPaginator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationPaginator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total"] = o.Total
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *OrganizationPaginator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"total",
		"offset",
		"limit",
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationPaginator := _OrganizationPaginator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationPaginator)

	if err != nil {
		return err
	}

	*o = OrganizationPaginator(varOrganizationPaginator)

	return err
}

type NullableOrganizationPaginator struct {
	value *OrganizationPaginator
	isSet bool
}

func (v NullableOrganizationPaginator) Get() *OrganizationPaginator {
	return v.value
}

func (v *NullableOrganizationPaginator) Set(val *OrganizationPaginator) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationPaginator) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationPaginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationPaginator(val *OrganizationPaginator) *NullableOrganizationPaginator {
	return &NullableOrganizationPaginator{value: val, isSet: true}
}

func (v NullableOrganizationPaginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationPaginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
