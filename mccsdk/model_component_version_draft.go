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

// ComponentVersionDraft struct for ComponentVersionDraft
type ComponentVersionDraft struct {
	Component string `json:"component"`
	Version   string `json:"version"`
}

// NewComponentVersionDraft instantiates a new ComponentVersionDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionDraft(component string, version string) *ComponentVersionDraft {
	this := ComponentVersionDraft{}
	this.Component = component
	this.Version = version
	return &this
}

// NewComponentVersionDraftWithDefaults instantiates a new ComponentVersionDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentVersionDraftWithDefaults() *ComponentVersionDraft {
	this := ComponentVersionDraft{}
	return &this
}

// GetComponent returns the Component field value
func (o *ComponentVersionDraft) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ComponentVersionDraft) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ComponentVersionDraft) SetComponent(v string) {
	o.Component = v
}

// GetVersion returns the Version field value
func (o *ComponentVersionDraft) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ComponentVersionDraft) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ComponentVersionDraft) SetVersion(v string) {
	o.Version = v
}

func (o ComponentVersionDraft) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableComponentVersionDraft struct {
	value *ComponentVersionDraft
	isSet bool
}

func (v NullableComponentVersionDraft) Get() *ComponentVersionDraft {
	return v.value
}

func (v *NullableComponentVersionDraft) Set(val *ComponentVersionDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentVersionDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentVersionDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentVersionDraft(val *ComponentVersionDraft) *NullableComponentVersionDraft {
	return &NullableComponentVersionDraft{value: val, isSet: true}
}

func (v NullableComponentVersionDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentVersionDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
