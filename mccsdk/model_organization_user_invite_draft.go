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

// OrganizationUserInviteDraft struct for OrganizationUserInviteDraft
type OrganizationUserInviteDraft struct {
	// E-mail address of the user
	Email string `json:"email"`
}

// NewOrganizationUserInviteDraft instantiates a new OrganizationUserInviteDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUserInviteDraft(email string) *OrganizationUserInviteDraft {
	this := OrganizationUserInviteDraft{}
	this.Email = email
	return &this
}

// NewOrganizationUserInviteDraftWithDefaults instantiates a new OrganizationUserInviteDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserInviteDraftWithDefaults() *OrganizationUserInviteDraft {
	this := OrganizationUserInviteDraft{}
	return &this
}

// GetEmail returns the Email field value
func (o *OrganizationUserInviteDraft) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrganizationUserInviteDraft) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrganizationUserInviteDraft) SetEmail(v string) {
	o.Email = v
}

func (o OrganizationUserInviteDraft) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUserInviteDraft struct {
	value *OrganizationUserInviteDraft
	isSet bool
}

func (v NullableOrganizationUserInviteDraft) Get() *OrganizationUserInviteDraft {
	return v.value
}

func (v *NullableOrganizationUserInviteDraft) Set(val *OrganizationUserInviteDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUserInviteDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUserInviteDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUserInviteDraft(val *OrganizationUserInviteDraft) *NullableOrganizationUserInviteDraft {
	return &NullableOrganizationUserInviteDraft{value: val, isSet: true}
}

func (v NullableOrganizationUserInviteDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUserInviteDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
