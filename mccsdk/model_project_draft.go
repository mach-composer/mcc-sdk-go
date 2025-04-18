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
	"fmt"
)

// checks if the ProjectDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDraft{}

// ProjectDraft struct for ProjectDraft
type ProjectDraft struct {
	// The project key (must be unique)
	Key string `json:"key"`
	// The name of the project
	Name string `json:"name"`
	// The description of the project
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectDraft ProjectDraft

// NewProjectDraft instantiates a new ProjectDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDraft(key string, name string) *ProjectDraft {
	this := ProjectDraft{}
	this.Key = key
	this.Name = name
	return &this
}

// NewProjectDraftWithDefaults instantiates a new ProjectDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDraftWithDefaults() *ProjectDraft {
	this := ProjectDraft{}
	return &this
}

// GetKey returns the Key field value
func (o *ProjectDraft) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProjectDraft) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProjectDraft) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ProjectDraft) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectDraft) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectDraft) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectDraft) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDraft) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectDraft) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectDraft) SetDescription(v string) {
	o.Description = &v
}

func (o ProjectDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectDraft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
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

	varProjectDraft := _ProjectDraft{}

	err = json.Unmarshal(data, &varProjectDraft)

	if err != nil {
		return err
	}

	*o = ProjectDraft(varProjectDraft)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectDraft struct {
	value *ProjectDraft
	isSet bool
}

func (v NullableProjectDraft) Get() *ProjectDraft {
	return v.value
}

func (v *NullableProjectDraft) Set(val *ProjectDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDraft(val *ProjectDraft) *NullableProjectDraft {
	return &NullableProjectDraft{value: val, isSet: true}
}

func (v NullableProjectDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
