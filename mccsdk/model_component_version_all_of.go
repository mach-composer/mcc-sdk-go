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

// ComponentVersionAllOf struct for ComponentVersionAllOf
type ComponentVersionAllOf struct {
	// key of the component
	Component string `json:"component"`
	// version of the component
	Version string `json:"version"`
	// branch of the version
	Branch *string `json:"branch,omitempty"`
}

// NewComponentVersionAllOf instantiates a new ComponentVersionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionAllOf(component string, version string) *ComponentVersionAllOf {
	this := ComponentVersionAllOf{}
	this.Component = component
	this.Version = version
	return &this
}

// NewComponentVersionAllOfWithDefaults instantiates a new ComponentVersionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentVersionAllOfWithDefaults() *ComponentVersionAllOf {
	this := ComponentVersionAllOf{}
	return &this
}

// GetComponent returns the Component field value
func (o *ComponentVersionAllOf) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ComponentVersionAllOf) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ComponentVersionAllOf) SetComponent(v string) {
	o.Component = v
}

// GetVersion returns the Version field value
func (o *ComponentVersionAllOf) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ComponentVersionAllOf) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ComponentVersionAllOf) SetVersion(v string) {
	o.Version = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ComponentVersionAllOf) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionAllOf) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ComponentVersionAllOf) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ComponentVersionAllOf) SetBranch(v string) {
	o.Branch = &v
}

func (o ComponentVersionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	return json.Marshal(toSerialize)
}

type NullableComponentVersionAllOf struct {
	value *ComponentVersionAllOf
	isSet bool
}

func (v NullableComponentVersionAllOf) Get() *ComponentVersionAllOf {
	return v.value
}

func (v *NullableComponentVersionAllOf) Set(val *ComponentVersionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentVersionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentVersionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentVersionAllOf(val *ComponentVersionAllOf) *NullableComponentVersionAllOf {
	return &NullableComponentVersionAllOf{value: val, isSet: true}
}

func (v NullableComponentVersionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentVersionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
