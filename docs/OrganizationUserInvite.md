# OrganizationUserInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] 
**Email** | **string** | E-mail address of the user | 
**ProjectKey** | Pointer to **string** | Key of the project to invite user to | [optional] 
**Role** | Pointer to **string** | Role for the user | [optional] 

## Methods

### NewOrganizationUserInvite

`func NewOrganizationUserInvite(email string, ) *OrganizationUserInvite`

NewOrganizationUserInvite instantiates a new OrganizationUserInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserInviteWithDefaults

`func NewOrganizationUserInviteWithDefaults() *OrganizationUserInvite`

NewOrganizationUserInviteWithDefaults instantiates a new OrganizationUserInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OrganizationUserInvite) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationUserInvite) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationUserInvite) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganizationUserInvite) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationUserInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUserInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUserInvite) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProjectKey

`func (o *OrganizationUserInvite) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *OrganizationUserInvite) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *OrganizationUserInvite) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *OrganizationUserInvite) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationUserInvite) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationUserInvite) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationUserInvite) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationUserInvite) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


