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
	"time"
)

// CommitDataAllOf struct for CommitDataAllOf
type CommitDataAllOf struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewCommitDataAllOf instantiates a new CommitDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitDataAllOf() *CommitDataAllOf {
	this := CommitDataAllOf{}
	return &this
}

// NewCommitDataAllOfWithDefaults instantiates a new CommitDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDataAllOfWithDefaults() *CommitDataAllOf {
	this := CommitDataAllOf{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommitDataAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDataAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommitDataAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CommitDataAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o CommitDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCommitDataAllOf struct {
	value *CommitDataAllOf
	isSet bool
}

func (v NullableCommitDataAllOf) Get() *CommitDataAllOf {
	return v.value
}

func (v *NullableCommitDataAllOf) Set(val *CommitDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitDataAllOf(val *CommitDataAllOf) *NullableCommitDataAllOf {
	return &NullableCommitDataAllOf{value: val, isSet: true}
}

func (v NullableCommitDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
