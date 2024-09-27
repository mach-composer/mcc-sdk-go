# OrganizationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**LastModifiedAt** | **time.Time** |  | [readonly] 
**AccountStatus** | **string** |  | 
**Name** | **string** |  | 
**Email** | **string** |  | 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsStaff** | Pointer to **bool** |  | [optional] 
**IsSuperuser** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganizationUser

`func NewOrganizationUser(id string, createdAt time.Time, lastModifiedAt time.Time, accountStatus string, name string, email string, ) *OrganizationUser`

NewOrganizationUser instantiates a new OrganizationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserWithDefaults

`func NewOrganizationUserWithDefaults() *OrganizationUser`

NewOrganizationUserWithDefaults instantiates a new OrganizationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationUser) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OrganizationUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastModifiedAt

`func (o *OrganizationUser) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *OrganizationUser) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *OrganizationUser) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.


### GetAccountStatus

`func (o *OrganizationUser) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *OrganizationUser) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *OrganizationUser) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.


### GetName

`func (o *OrganizationUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationUser) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *OrganizationUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAvatarUrl

`func (o *OrganizationUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *OrganizationUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *OrganizationUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *OrganizationUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *OrganizationUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *OrganizationUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *OrganizationUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *OrganizationUser) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsStaff

`func (o *OrganizationUser) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *OrganizationUser) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *OrganizationUser) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *OrganizationUser) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsSuperuser

`func (o *OrganizationUser) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *OrganizationUser) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *OrganizationUser) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *OrganizationUser) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


