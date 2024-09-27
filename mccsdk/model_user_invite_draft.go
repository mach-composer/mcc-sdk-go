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

// checks if the UserInviteDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInviteDraft{}

// UserInviteDraft struct for UserInviteDraft
type UserInviteDraft struct {
	AcceptedAt   NullableTime   `json:"accepted_at,omitempty"`
	CreatedBy    string         `json:"created_by"`
	Email        string         `json:"email"`
	Organization string         `json:"organization"`
	Project      NullableString `json:"project,omitempty"`
	Role         string         `json:"role"`
}

type _UserInviteDraft UserInviteDraft

// NewUserInviteDraft instantiates a new UserInviteDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInviteDraft(createdBy string, email string, organization string, role string) *UserInviteDraft {
	this := UserInviteDraft{}
	this.CreatedBy = createdBy
	this.Email = email
	this.Organization = organization
	this.Role = role
	return &this
}

// NewUserInviteDraftWithDefaults instantiates a new UserInviteDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInviteDraftWithDefaults() *UserInviteDraft {
	this := UserInviteDraft{}
	return &this
}

// GetAcceptedAt returns the AcceptedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInviteDraft) GetAcceptedAt() time.Time {
	if o == nil || IsNil(o.AcceptedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AcceptedAt.Get()
}

// GetAcceptedAtOk returns a tuple with the AcceptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInviteDraft) GetAcceptedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedAt.Get(), o.AcceptedAt.IsSet()
}

// HasAcceptedAt returns a boolean if a field has been set.
func (o *UserInviteDraft) HasAcceptedAt() bool {
	if o != nil && o.AcceptedAt.IsSet() {
		return true
	}

	return false
}

// SetAcceptedAt gets a reference to the given NullableTime and assigns it to the AcceptedAt field.
func (o *UserInviteDraft) SetAcceptedAt(v time.Time) {
	o.AcceptedAt.Set(&v)
}

// SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil
func (o *UserInviteDraft) SetAcceptedAtNil() {
	o.AcceptedAt.Set(nil)
}

// UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
func (o *UserInviteDraft) UnsetAcceptedAt() {
	o.AcceptedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value
func (o *UserInviteDraft) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *UserInviteDraft) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *UserInviteDraft) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEmail returns the Email field value
func (o *UserInviteDraft) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserInviteDraft) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserInviteDraft) SetEmail(v string) {
	o.Email = v
}

// GetOrganization returns the Organization field value
func (o *UserInviteDraft) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *UserInviteDraft) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *UserInviteDraft) SetOrganization(v string) {
	o.Organization = v
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInviteDraft) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInviteDraft) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *UserInviteDraft) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *UserInviteDraft) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *UserInviteDraft) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *UserInviteDraft) UnsetProject() {
	o.Project.Unset()
}

// GetRole returns the Role field value
func (o *UserInviteDraft) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserInviteDraft) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserInviteDraft) SetRole(v string) {
	o.Role = v
}

func (o UserInviteDraft) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInviteDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptedAt.IsSet() {
		toSerialize["accepted_at"] = o.AcceptedAt.Get()
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["email"] = o.Email
	toSerialize["organization"] = o.Organization
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *UserInviteDraft) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_by",
		"email",
		"organization",
		"role",
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

	varUserInviteDraft := _UserInviteDraft{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserInviteDraft)

	if err != nil {
		return err
	}

	*o = UserInviteDraft(varUserInviteDraft)

	return err
}

type NullableUserInviteDraft struct {
	value *UserInviteDraft
	isSet bool
}

func (v NullableUserInviteDraft) Get() *UserInviteDraft {
	return v.value
}

func (v *NullableUserInviteDraft) Set(val *UserInviteDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInviteDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInviteDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInviteDraft(val *UserInviteDraft) *NullableUserInviteDraft {
	return &NullableUserInviteDraft{value: val, isSet: true}
}

func (v NullableUserInviteDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInviteDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}