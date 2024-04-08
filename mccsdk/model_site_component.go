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
	"time"
)

// checks if the SiteComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteComponent{}

// SiteComponent struct for SiteComponent
type SiteComponent struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// key of the site component
	Key string `json:"key"`
	// name of the site component
	Name string `json:"name"`
	// id of the site
	SiteId string `json:"site_id"`
	// id of the component
	ComponentId string `json:"component_id"`
}

type _SiteComponent SiteComponent

// NewSiteComponent instantiates a new SiteComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteComponent(id string, createdAt time.Time, key string, name string, siteId string, componentId string) *SiteComponent {
	this := SiteComponent{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Key = key
	this.Name = name
	this.SiteId = siteId
	this.ComponentId = componentId
	return &this
}

// NewSiteComponentWithDefaults instantiates a new SiteComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteComponentWithDefaults() *SiteComponent {
	this := SiteComponent{}
	return &this
}

// GetId returns the Id field value
func (o *SiteComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SiteComponent) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SiteComponent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SiteComponent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetKey returns the Key field value
func (o *SiteComponent) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SiteComponent) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *SiteComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteComponent) SetName(v string) {
	o.Name = v
}

// GetSiteId returns the SiteId field value
func (o *SiteComponent) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *SiteComponent) SetSiteId(v string) {
	o.SiteId = v
}

// GetComponentId returns the ComponentId field value
func (o *SiteComponent) GetComponentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *SiteComponent) GetComponentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *SiteComponent) SetComponentId(v string) {
	o.ComponentId = v
}

func (o SiteComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["site_id"] = o.SiteId
	toSerialize["component_id"] = o.ComponentId
	return toSerialize, nil
}

func (o *SiteComponent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"key",
		"name",
		"site_id",
		"component_id",
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

	varSiteComponent := _SiteComponent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSiteComponent)

	if err != nil {
		return err
	}

	*o = SiteComponent(varSiteComponent)

	return err
}

type NullableSiteComponent struct {
	value *SiteComponent
	isSet bool
}

func (v NullableSiteComponent) Get() *SiteComponent {
	return v.value
}

func (v *NullableSiteComponent) Set(val *SiteComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteComponent(val *SiteComponent) *NullableSiteComponent {
	return &NullableSiteComponent{value: val, isSet: true}
}

func (v NullableSiteComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}