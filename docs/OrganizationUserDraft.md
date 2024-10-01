# OrganizationUserDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountStatus** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsStaff** | Pointer to **bool** |  | [optional] 
**IsSuperuser** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganizationUserDraft

`func NewOrganizationUserDraft(accountStatus string, name string, email string, ) *OrganizationUserDraft`

NewOrganizationUserDraft instantiates a new OrganizationUserDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserDraftWithDefaults

`func NewOrganizationUserDraftWithDefaults() *OrganizationUserDraft`

NewOrganizationUserDraftWithDefaults instantiates a new OrganizationUserDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountStatus

`func (o *OrganizationUserDraft) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *OrganizationUserDraft) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *OrganizationUserDraft) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.


### GetName

`func (o *OrganizationUserDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationUserDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationUserDraft) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *OrganizationUserDraft) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUserDraft) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUserDraft) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *OrganizationUserDraft) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *OrganizationUserDraft) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *OrganizationUserDraft) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *OrganizationUserDraft) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *OrganizationUserDraft) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *OrganizationUserDraft) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *OrganizationUserDraft) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *OrganizationUserDraft) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsStaff

`func (o *OrganizationUserDraft) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *OrganizationUserDraft) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *OrganizationUserDraft) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *OrganizationUserDraft) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *OrganizationUserDraft) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *OrganizationUserDraft) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *OrganizationUserDraft) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *OrganizationUserDraft) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


