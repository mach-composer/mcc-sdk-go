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

// checks if the ComponentVersionDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentVersionDraft{}

// ComponentVersionDraft struct for ComponentVersionDraft
type ComponentVersionDraft struct {
	// version of the component
	Version string `json:"version"`
	// branch of the version
	Branch *string `json:"branch,omitempty"`
}

type _ComponentVersionDraft ComponentVersionDraft

// NewComponentVersionDraft instantiates a new ComponentVersionDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionDraft(version string) *ComponentVersionDraft {
	this := ComponentVersionDraft{}
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

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ComponentVersionDraft) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionDraft) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ComponentVersionDraft) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ComponentVersionDraft) SetBranch(v string) {
	o.Branch = &v
}

func (o ComponentVersionDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentVersionDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	return toSerialize, nil
}

func (o *ComponentVersionDraft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
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

	varComponentVersionDraft := _ComponentVersionDraft{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentVersionDraft)

	if err != nil {
		return err
	}

	*o = ComponentVersionDraft(varComponentVersionDraft)

	return err
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
