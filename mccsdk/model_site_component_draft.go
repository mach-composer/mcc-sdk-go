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

// checks if the SiteComponentDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteComponentDraft{}

// SiteComponentDraft struct for SiteComponentDraft
type SiteComponentDraft struct {
	// key of the site component
	Key string `json:"key"`
	// name of the site component
	Name string `json:"name"`
	// key of the site
	SiteKey string `json:"site_key"`
	// key of the component
	ComponentKey string `json:"component_key"`
}

type _SiteComponentDraft SiteComponentDraft

// NewSiteComponentDraft instantiates a new SiteComponentDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteComponentDraft(key string, name string, siteKey string, componentKey string) *SiteComponentDraft {
	this := SiteComponentDraft{}
	this.Key = key
	this.Name = name
	this.SiteKey = siteKey
	this.ComponentKey = componentKey
	return &this
}

// NewSiteComponentDraftWithDefaults instantiates a new SiteComponentDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteComponentDraftWithDefaults() *SiteComponentDraft {
	this := SiteComponentDraft{}
	return &this
}

// GetKey returns the Key field value
func (o *SiteComponentDraft) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SiteComponentDraft) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SiteComponentDraft) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *SiteComponentDraft) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteComponentDraft) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteComponentDraft) SetName(v string) {
	o.Name = v
}

// GetSiteKey returns the SiteKey field value
func (o *SiteComponentDraft) GetSiteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteKey
}

// GetSiteKeyOk returns a tuple with the SiteKey field value
// and a boolean to check if the value has been set.
func (o *SiteComponentDraft) GetSiteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteKey, true
}

// SetSiteKey sets field value
func (o *SiteComponentDraft) SetSiteKey(v string) {
	o.SiteKey = v
}

// GetComponentKey returns the ComponentKey field value
func (o *SiteComponentDraft) GetComponentKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentKey
}

// GetComponentKeyOk returns a tuple with the ComponentKey field value
// and a boolean to check if the value has been set.
func (o *SiteComponentDraft) GetComponentKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentKey, true
}

// SetComponentKey sets field value
func (o *SiteComponentDraft) SetComponentKey(v string) {
	o.ComponentKey = v
}

func (o SiteComponentDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteComponentDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["site_key"] = o.SiteKey
	toSerialize["component_key"] = o.ComponentKey
	return toSerialize, nil
}

func (o *SiteComponentDraft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
		"site_key",
		"component_key",
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

	varSiteComponentDraft := _SiteComponentDraft{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSiteComponentDraft)

	if err != nil {
		return err
	}

	*o = SiteComponentDraft(varSiteComponentDraft)

	return err
}

type NullableSiteComponentDraft struct {
	value *SiteComponentDraft
	isSet bool
}

func (v NullableSiteComponentDraft) Get() *SiteComponentDraft {
	return v.value
}

func (v *NullableSiteComponentDraft) Set(val *SiteComponentDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteComponentDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteComponentDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteComponentDraft(val *SiteComponentDraft) *NullableSiteComponentDraft {
	return &NullableSiteComponentDraft{value: val, isSet: true}
}

func (v NullableSiteComponentDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteComponentDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}