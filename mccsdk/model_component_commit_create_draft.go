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

// checks if the ComponentCommitCreateDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentCommitCreateDraft{}

// ComponentCommitCreateDraft struct for ComponentCommitCreateDraft
type ComponentCommitCreateDraft struct {
	Commits              []CommitDraft `json:"commits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComponentCommitCreateDraft ComponentCommitCreateDraft

// NewComponentCommitCreateDraft instantiates a new ComponentCommitCreateDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentCommitCreateDraft() *ComponentCommitCreateDraft {
	this := ComponentCommitCreateDraft{}
	return &this
}

// NewComponentCommitCreateDraftWithDefaults instantiates a new ComponentCommitCreateDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentCommitCreateDraftWithDefaults() *ComponentCommitCreateDraft {
	this := ComponentCommitCreateDraft{}
	return &this
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *ComponentCommitCreateDraft) GetCommits() []CommitDraft {
	if o == nil || IsNil(o.Commits) {
		var ret []CommitDraft
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentCommitCreateDraft) GetCommitsOk() ([]CommitDraft, bool) {
	if o == nil || IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *ComponentCommitCreateDraft) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []CommitDraft and assigns it to the Commits field.
func (o *ComponentCommitCreateDraft) SetCommits(v []CommitDraft) {
	o.Commits = v
}

func (o ComponentCommitCreateDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentCommitCreateDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commits) {
		toSerialize["commits"] = o.Commits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComponentCommitCreateDraft) UnmarshalJSON(data []byte) (err error) {
	varComponentCommitCreateDraft := _ComponentCommitCreateDraft{}

	err = json.Unmarshal(data, &varComponentCommitCreateDraft)

	if err != nil {
		return err
	}

	*o = ComponentCommitCreateDraft(varComponentCommitCreateDraft)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComponentCommitCreateDraft struct {
	value *ComponentCommitCreateDraft
	isSet bool
}

func (v NullableComponentCommitCreateDraft) Get() *ComponentCommitCreateDraft {
	return v.value
}

func (v *NullableComponentCommitCreateDraft) Set(val *ComponentCommitCreateDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentCommitCreateDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentCommitCreateDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentCommitCreateDraft(val *ComponentCommitCreateDraft) *NullableComponentCommitCreateDraft {
	return &NullableComponentCommitCreateDraft{value: val, isSet: true}
}

func (v NullableComponentCommitCreateDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentCommitCreateDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
