# UserInviteDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedBy** | **string** |  | 
**Email** | **string** |  | 
**Organization** | **string** |  | 
**Project** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewUserInviteDraft

`func NewUserInviteDraft(createdBy string, email string, organization string, role string, ) *UserInviteDraft`

NewUserInviteDraft instantiates a new UserInviteDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInviteDraftWithDefaults

`func NewUserInviteDraftWithDefaults() *UserInviteDraft`

NewUserInviteDraftWithDefaults instantiates a new UserInviteDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *UserInviteDraft) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *UserInviteDraft) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *UserInviteDraft) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *UserInviteDraft) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *UserInviteDraft) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *UserInviteDraft) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetCreatedBy

`func (o *UserInviteDraft) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserInviteDraft) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserInviteDraft) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetEmail

`func (o *UserInviteDraft) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInviteDraft) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInviteDraft) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetOrganization

`func (o *UserInviteDraft) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserInviteDraft) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserInviteDraft) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetProject

`func (o *UserInviteDraft) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UserInviteDraft) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UserInviteDraft) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *UserInviteDraft) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *UserInviteDraft) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *UserInviteDraft) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetRole

`func (o *UserInviteDraft) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserInviteDraft) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserInviteDraft) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


