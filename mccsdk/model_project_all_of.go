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

// ProjectAllOf struct for ProjectAllOf
type ProjectAllOf struct {
	// The organization key (must be unique)
	Key string `json:"key"`
	// The name of the organization
	Name string `json:"name"`
	// description about the api client
	Description *string `json:"description,omitempty"`
}

// NewProjectAllOf instantiates a new ProjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAllOf(key string, name string) *ProjectAllOf {
	this := ProjectAllOf{}
	this.Key = key
	this.Name = name
	return &this
}

// NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAllOfWithDefaults() *ProjectAllOf {
	this := ProjectAllOf{}
	return &this
}

// GetKey returns the Key field value
func (o *ProjectAllOf) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProjectAllOf) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ProjectAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectAllOf) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectAllOf) SetDescription(v string) {
	o.Description = &v
}

func (o ProjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableProjectAllOf struct {
	value *ProjectAllOf
	isSet bool
}

func (v NullableProjectAllOf) Get() *ProjectAllOf {
	return v.value
}

func (v *NullableProjectAllOf) Set(val *ProjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAllOf(val *ProjectAllOf) *NullableProjectAllOf {
	return &NullableProjectAllOf{value: val, isSet: true}
}

func (v NullableProjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
