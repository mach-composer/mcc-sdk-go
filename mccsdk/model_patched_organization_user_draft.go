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

// checks if the PatchedOrganizationUserDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedOrganizationUserDraft{}

// PatchedOrganizationUserDraft struct for PatchedOrganizationUserDraft
type PatchedOrganizationUserDraft struct {
	AccountStatus        *string `json:"account_status,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Email                *string `json:"email,omitempty"`
	AvatarUrl            *string `json:"avatar_url,omitempty"`
	IsActive             *bool   `json:"is_active,omitempty"`
	IsStaff              *bool   `json:"is_staff,omitempty"`
	IsSuperuser          *bool   `json:"is_superuser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedOrganizationUserDraft PatchedOrganizationUserDraft

// NewPatchedOrganizationUserDraft instantiates a new PatchedOrganizationUserDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOrganizationUserDraft() *PatchedOrganizationUserDraft {
	this := PatchedOrganizationUserDraft{}
	return &this
}

// NewPatchedOrganizationUserDraftWithDefaults instantiates a new PatchedOrganizationUserDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOrganizationUserDraftWithDefaults() *PatchedOrganizationUserDraft {
	this := PatchedOrganizationUserDraft{}
	return &this
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetAccountStatus() string {
	if o == nil || IsNil(o.AccountStatus) {
		var ret string
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetAccountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasAccountStatus() bool {
	if o != nil && !IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given string and assigns it to the AccountStatus field.
func (o *PatchedOrganizationUserDraft) SetAccountStatus(v string) {
	o.AccountStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOrganizationUserDraft) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedOrganizationUserDraft) SetEmail(v string) {
	o.Email = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *PatchedOrganizationUserDraft) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedOrganizationUserDraft) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetIsStaff() bool {
	if o == nil || IsNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetIsStaffOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStaff) {
		return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasIsStaff() bool {
	if o != nil && !IsNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *PatchedOrganizationUserDraft) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *PatchedOrganizationUserDraft) GetIsSuperuser() bool {
	if o == nil || IsNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOrganizationUserDraft) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperuser) {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *PatchedOrganizationUserDraft) HasIsSuperuser() bool {
	if o != nil && !IsNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *PatchedOrganizationUserDraft) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

func (o PatchedOrganizationUserDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedOrganizationUserDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountStatus) {
		toSerialize["account_status"] = o.AccountStatus
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsStaff) {
		toSerialize["is_staff"] = o.IsStaff
	}
	if !IsNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedOrganizationUserDraft) UnmarshalJSON(data []byte) (err error) {
	varPatchedOrganizationUserDraft := _PatchedOrganizationUserDraft{}

	err = json.Unmarshal(data, &varPatchedOrganizationUserDraft)

	if err != nil {
		return err
	}

	*o = PatchedOrganizationUserDraft(varPatchedOrganizationUserDraft)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_status")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_staff")
		delete(additionalProperties, "is_superuser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedOrganizationUserDraft struct {
	value *PatchedOrganizationUserDraft
	isSet bool
}

func (v NullablePatchedOrganizationUserDraft) Get() *PatchedOrganizationUserDraft {
	return v.value
}

func (v *NullablePatchedOrganizationUserDraft) Set(val *PatchedOrganizationUserDraft) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOrganizationUserDraft) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOrganizationUserDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOrganizationUserDraft(val *PatchedOrganizationUserDraft) *NullablePatchedOrganizationUserDraft {
	return &NullablePatchedOrganizationUserDraft{value: val, isSet: true}
}

func (v NullablePatchedOrganizationUserDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOrganizationUserDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
