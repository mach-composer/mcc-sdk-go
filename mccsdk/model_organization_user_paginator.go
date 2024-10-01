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

// checks if the OrganizationUserPaginator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationUserPaginator{}

// OrganizationUserPaginator struct for OrganizationUserPaginator
type OrganizationUserPaginator struct {
	// Number of items in the current page
	Count *int32 `json:"count,omitempty"`
	// Total number of items found
	Total   *int64             `json:"total,omitempty"`
	Offset  *int32             `json:"offset,omitempty"`
	Limit   *int32             `json:"limit,omitempty"`
	Results []OrganizationUser `json:"results,omitempty"`
}

// NewOrganizationUserPaginator instantiates a new OrganizationUserPaginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserPaginator() *OrganizationUserPaginator {
	this := OrganizationUserPaginator{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// NewOrganizationUserPaginatorWithDefaults instantiates a new OrganizationUserPaginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserPaginatorWithDefaults() *OrganizationUserPaginator {
	this := OrganizationUserPaginator{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OrganizationUserPaginator) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OrganizationUserPaginator) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OrganizationUserPaginator) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrganizationUserPaginator) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrganizationUserPaginator) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *OrganizationUserPaginator) SetTotal(v int64) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *OrganizationUserPaginator) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *OrganizationUserPaginator) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *OrganizationUserPaginator) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *OrganizationUserPaginator) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *OrganizationUserPaginator) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *OrganizationUserPaginator) SetLimit(v int32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OrganizationUserPaginator) GetResults() []OrganizationUser {
	if o == nil || IsNil(o.Results) {
		var ret []OrganizationUser
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUserPaginator) GetResultsOk() ([]OrganizationUser, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OrganizationUserPaginator) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrganizationUser and assigns it to the Results field.
func (o *OrganizationUserPaginator) SetResults(v []OrganizationUser) {
	o.Results = v
}

func (o OrganizationUserPaginator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationUserPaginator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
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
