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

// ComponentVersionCommits struct for ComponentVersionCommits
type ComponentVersionCommits struct {
	Commits []CommitData `json:"commits"`
}

// NewComponentVersionCommits instantiates a new ComponentVersionCommits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionCommits(commits []CommitData) *ComponentVersionCommits {
	this := ComponentVersionCommits{}
	this.Commits = commits
	return &this
}

// NewComponentVersionCommitsWithDefaults instantiates a new ComponentVersionCommits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentVersionCommitsWithDefaults() *ComponentVersionCommits {
	this := ComponentVersionCommits{}
	return &this
}

// GetCommits returns the Commits field value
func (o *ComponentVersionCommits) GetCommits() []CommitData {
	if o == nil {
		var ret []CommitData
		return ret
	}

	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value
// and a boolean to check if the value has been set.
func (o *ComponentVersionCommits) GetCommitsOk() ([]CommitData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commits, true
}

// SetCommits sets field value
func (o *ComponentVersionCommits) SetCommits(v []CommitData) {
	o.Commits = v
}

func (o ComponentVersionCommits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commits"] = o.Commits
	}
	return json.Marshal(toSerialize)
}

type NullableComponentVersionCommits struct {
	value *ComponentVersionCommits
	isSet bool
}

func (v NullableComponentVersionCommits) Get() *ComponentVersionCommits {
	return v.value
}

func (v *NullableComponentVersionCommits) Set(val *ComponentVersionCommits) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentVersionCommits) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentVersionCommits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentVersionCommits(val *ComponentVersionCommits) *NullableComponentVersionCommits {
	return &NullableComponentVersionCommits{value: val, isSet: true}
}

func (v NullableComponentVersionCommits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentVersionCommits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
