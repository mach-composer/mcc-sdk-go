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

// OrganizationUserInviteData struct for OrganizationUserInviteData
type OrganizationUserInviteData struct {
	Id           string                                 `json:"id"`
	CreatedBy    string                                 `json:"created_by"`
	Organization OrganizationUserInviteDataOrganization `json:"organization"`
}

// NewOrganizationUserInviteData instantiates a new OrganizationUserInviteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserInviteData(id string, createdBy string, organization OrganizationUserInviteDataOrganization) *OrganizationUserInviteData {
	this := OrganizationUserInviteData{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Organization = organization
	return &this
}

// NewOrganizationUserInviteDataWithDefaults instantiates a new OrganizationUserInviteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserInviteDataWithDefaults() *OrganizationUserInviteData {
	this := OrganizationUserInviteData{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationUserInviteData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserInviteData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationUserInviteData) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *OrganizationUserInviteData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserInviteData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *OrganizationUserInviteData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetOrganization returns the Organization field value
func (o *OrganizationUserInviteData) GetOrganization() OrganizationUserInviteDataOrganization {
	if o == nil {
		var ret OrganizationUserInviteDataOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserInviteData) GetOrganizationOk() (*OrganizationUserInviteDataOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *OrganizationUserInviteData) SetOrganization(v OrganizationUserInviteDataOrganization) {
	o.Organization = v
}

func (o OrganizationUserInviteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserInviteData struct {
	value *OrganizationUserInviteData
	isSet bool
}

func (v NullableOrganizationUserInviteData) Get() *OrganizationUserInviteData {
	return v.value
}

func (v *NullableOrganizationUserInviteData) Set(val *OrganizationUserInviteData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserInviteData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserInviteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserInviteData(val *OrganizationUserInviteData) *NullableOrganizationUserInviteData {
	return &NullableOrganizationUserInviteData{value: val, isSet: true}
}

func (v NullableOrganizationUserInviteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserInviteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
