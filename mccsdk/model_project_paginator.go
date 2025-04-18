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

// checks if the ProjectPaginator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectPaginator{}

// ProjectPaginator struct for ProjectPaginator
type ProjectPaginator struct {
	// Number of items in the current page
	Count *int32 `json:"count,omitempty"`
	// Total number of items found
	Total                *int64    `json:"total,omitempty"`
	Offset               *int32    `json:"offset,omitempty"`
	Limit                *int32    `json:"limit,omitempty"`
	Results              []Project `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectPaginator ProjectPaginator

// NewProjectPaginator instantiates a new ProjectPaginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPaginator() *ProjectPaginator {
	this := ProjectPaginator{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// NewProjectPaginatorWithDefaults instantiates a new ProjectPaginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPaginatorWithDefaults() *ProjectPaginator {
	this := ProjectPaginator{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 10
	this.Limit = &limit
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProjectPaginator) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginator) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProjectPaginator) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProjectPaginator) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectPaginator) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginator) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectPaginator) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ProjectPaginator) SetTotal(v int64) {
	o.Total = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ProjectPaginator) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginator) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ProjectPaginator) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ProjectPaginator) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProjectPaginator) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginator) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProjectPaginator) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ProjectPaginator) SetLimit(v int32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProjectPaginator) GetResults() []Project {
	if o == nil || IsNil(o.Results) {
		var ret []Project
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPaginator) GetResultsOk() ([]Project, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProjectPaginator) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Project and assigns it to the Results field.
func (o *ProjectPaginator) SetResults(v []Project) {
	o.Results = v
}

func (o ProjectPaginator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectPaginator) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectPaginator) UnmarshalJSON(data []byte) (err error) {
	varProjectPaginator := _ProjectPaginator{}

	err = json.Unmarshal(data, &varProjectPaginator)

	if err != nil {
		return err
	}

	*o = ProjectPaginator(varProjectPaginator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "total")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectPaginator struct {
	value *ProjectPaginator
	isSet bool
}

func (v NullableProjectPaginator) Get() *ProjectPaginator {
	return v.value
}

func (v *NullableProjectPaginator) Set(val *ProjectPaginator) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPaginator) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPaginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPaginator(val *ProjectPaginator) *NullableProjectPaginator {
	return &NullableProjectPaginator{value: val, isSet: true}
}

func (v NullableProjectPaginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPaginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
