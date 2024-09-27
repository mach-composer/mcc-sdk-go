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

// checks if the PatchedComponentDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedComponentDraft{}

// PatchedComponentDraft struct for PatchedComponentDraft
type PatchedComponentDraft struct {
	// key of the component
	Key *string `json:"key,omitempty"`
	// short description of the component
	Description *string `json:"description,omitempty"`
	// name of the component
	Name *string `json:"name,omitempty"`
}

// NewPatchedComponentDraft instantiates a new PatchedComponentDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedComponentDraft() *PatchedComponentDraft {
	this := PatchedComponentDraft{}
	return &this
}

// NewPatchedComponentDraftWithDefaults instantiates a new PatchedComponentDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedComponentDraftWithDefaults() *PatchedComponentDraft {
	this := PatchedComponentDraft{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PatchedComponentDraft) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedComponentDraft) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PatchedComponentDraft) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PatchedComponentDraft) SetKey(v string) {
	o.Key = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedComponentDraft) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedComponentDraft) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedComponentDraft) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedComponentDraft) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedComponentDraft) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedComponentDraft) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedComponentDraft) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedComponentDraft) SetName(v string) {
	o.Name = &v
}

func (o PatchedComponentDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedComponentDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePatchedComponentDraft struct {
	value *PatchedComponentDraft
	isSet bool
}

func (v NullablePatchedComponentDraft) Get() *PatchedComponentDraft {
	return v.value
}

func (v *NullablePatchedComponentDraft) Set(val *PatchedComponentDraft) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedComponentDraft) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedComponentDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedComponentDraft(val *PatchedComponentDraft) *NullablePatchedComponentDraft {
	return &NullablePatchedComponentDraft{value: val, isSet: true}
}

func (v NullablePatchedComponentDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedComponentDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}