# OrganizationUserInviteDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | E-mail address of the user | 
**ProjectKey** | Pointer to **string** | Key of the project to invite user to | [optional] 
**Role** | Pointer to **string** | Role for the user | [optional] 

## Methods

### NewOrganizationUserInviteDraft

`func NewOrganizationUserInviteDraft(email string, ) *OrganizationUserInviteDraft`

NewOrganizationUserInviteDraft instantiates a new OrganizationUserInviteDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserInviteDraftWithDefaults

`func NewOrganizationUserInviteDraftWithDefaults() *OrganizationUserInviteDraft`

NewOrganizationUserInviteDraftWithDefaults instantiates a new OrganizationUserInviteDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationUserInviteDraft) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUserInviteDraft) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUserInviteDraft) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProjectKey

`func (o *OrganizationUserInviteDraft) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *OrganizationUserInviteDraft) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *OrganizationUserInviteDraft) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *OrganizationUserInviteDraft) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetRole

`func (o *OrganizationUserInviteDraft) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationUserInviteDraft) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationUserInviteDraft) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationUserInviteDraft) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


