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

// OrganizationDraft struct for OrganizationDraft
type OrganizationDraft struct {
	// The organization key (must be unique)
	Key string `json:"key"`
	// The name of the organization
	Name string `json:"name"`
}

// NewOrganizationDraft instantiates a new OrganizationDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDraft(key string, name string) *OrganizationDraft {
	this := OrganizationDraft{}
	this.Key = key
	this.Name = name
	return &this
}

// NewOrganizationDraftWithDefaults instantiates a new OrganizationDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDraftWithDefaults() *OrganizationDraft {
	this := OrganizationDraft{}
	return &this
}

// GetKey returns the Key field value
func (o *OrganizationDraft) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *OrganizationDraft) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *OrganizationDraft) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *OrganizationDraft) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationDraft) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationDraft) SetName(v string) {
	o.Name = v
}

func (o OrganizationDraft) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationDraft struct {
	value *OrganizationDraft
	isSet bool
}

func (v NullableOrganizationDraft) Get() *OrganizationDraft {
	return v.value
}

func (v *NullableOrganizationDraft) Set(val *OrganizationDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDraft(val *OrganizationDraft) *NullableOrganizationDraft {
	return &NullableOrganizationDraft{value: val, isSet: true}
}

func (v NullableOrganizationDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
