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
	"time"
)

// OrganizationUser struct for OrganizationUser
type OrganizationUser struct {
	ClientId  *string    `json:"client_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Name of the user
	Name *string `json:"name,omitempty"`
	// E-mail address of the user
	Email string `json:"email"`
	// The avatar profile image url of the user
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// Key of the organization
	OrganizationKey *string `json:"organization_key,omitempty"`
	// Name of the organization
	OrganizationName   *string                         `json:"organization_name,omitempty"`
	OrganizationScopes []string                        `json:"organization_scopes,omitempty"`
	Projects           []OrganizationUserProjectsInner `json:"projects,omitempty"`
}

// NewOrganizationUser instantiates a new OrganizationUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUser(email string) *OrganizationUser {
	this := OrganizationUser{}
	this.Email = email
	return &this
}

// NewOrganizationUserWithDefaults instantiates a new OrganizationUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUserWithDefaults() *OrganizationUser {
	this := OrganizationUser{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OrganizationUser) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OrganizationUser) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OrganizationUser) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganizationUser) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganizationUser) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganizationUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationUser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationUser) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value
func (o *OrganizationUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrganizationUser) SetEmail(v string) {
	o.Email = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *OrganizationUser) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *OrganizationUser) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *OrganizationUser) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetOrganizationKey returns the OrganizationKey field value if set, zero value otherwise.
func (o *OrganizationUser) GetOrganizationKey() string {
	if o == nil || o.OrganizationKey == nil {
		var ret string
		return ret
	}
	return *o.OrganizationKey
}

// GetOrganizationKeyOk returns a tuple with the OrganizationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetOrganizationKeyOk() (*string, bool) {
	if o == nil || o.OrganizationKey == nil {
		return nil, false
	}
	return o.OrganizationKey, true
}

// HasOrganizationKey returns a boolean if a field has been set.
func (o *OrganizationUser) HasOrganizationKey() bool {
	if o != nil && o.OrganizationKey != nil {
		return true
	}

	return false
}

// SetOrganizationKey gets a reference to the given string and assigns it to the OrganizationKey field.
func (o *OrganizationUser) SetOrganizationKey(v string) {
	o.OrganizationKey = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *OrganizationUser) GetOrganizationName() string {
	if o == nil || o.OrganizationName == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetOrganizationNameOk() (*string, bool) {
	if o == nil || o.OrganizationName == nil {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *OrganizationUser) HasOrganizationName() bool {
	if o != nil && o.OrganizationName != nil {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *OrganizationUser) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetOrganizationScopes returns the OrganizationScopes field value if set, zero value otherwise.
func (o *OrganizationUser) GetOrganizationScopes() []string {
	if o == nil || o.OrganizationScopes == nil {
		var ret []string
		return ret
	}
	return o.OrganizationScopes
}

// GetOrganizationScopesOk returns a tuple with the OrganizationScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetOrganizationScopesOk() ([]string, bool) {
	if o == nil || o.OrganizationScopes == nil {
		return nil, false
	}
	return o.OrganizationScopes, true
}

// HasOrganizationScopes returns a boolean if a field has been set.
func (o *OrganizationUser) HasOrganizationScopes() bool {
	if o != nil && o.OrganizationScopes != nil {
		return true
	}

	return false
}

// SetOrganizationScopes gets a reference to the given []string and assigns it to the OrganizationScopes field.
func (o *OrganizationUser) SetOrganizationScopes(v []string) {
	o.OrganizationScopes = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *OrganizationUser) GetProjects() []OrganizationUserProjectsInner {
	if o == nil || o.Projects == nil {
		var ret []OrganizationUserProjectsInner
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUser) GetProjectsOk() ([]OrganizationUserProjectsInner, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *OrganizationUser) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []OrganizationUserProjectsInner and assigns it to the Projects field.
func (o *OrganizationUser) SetProjects(v []OrganizationUserProjectsInner) {
	o.Projects = v
}

func (o OrganizationUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if o.OrganizationKey != nil {
		toSerialize["organization_key"] = o.OrganizationKey
	}
	if o.OrganizationName != nil {
		toSerialize["organization_name"] = o.OrganizationName
	}
	if o.OrganizationScopes != nil {
		toSerialize["organization_scopes"] = o.OrganizationScopes
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationUser struct {
	value *OrganizationUser
	isSet bool
}

func (v NullableOrganizationUser) Get() *OrganizationUser {
	return v.value
}

func (v *NullableOrganizationUser) Set(val *OrganizationUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUser(val *OrganizationUser) *NullableOrganizationUser {
	return &NullableOrganizationUser{value: val, isSet: true}
}

func (v NullableOrganizationUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
