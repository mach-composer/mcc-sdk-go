# OrganizationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | Name of the user | [optional] 
**Email** | **string** | E-mail address of the user | 
**AvatarUrl** | Pointer to **string** | The avatar profile image url of the user | [optional] 
**OrganizationKey** | Pointer to **string** | Key of the organization | [optional] 
**OrganizationName** | Pointer to **string** | Name of the organization | [optional] 
**OrganizationScopes** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to [**[]OrganizationUserProjectsInner**](OrganizationUserProjectsInner.md) |  | [optional] 

## Methods

### NewOrganizationUser

`func NewOrganizationUser(email string, ) *OrganizationUser`

NewOrganizationUser instantiates a new OrganizationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserWithDefaults

`func NewOrganizationUserWithDefaults() *OrganizationUser`

NewOrganizationUserWithDefaults instantiates a new OrganizationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OrganizationUser) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OrganizationUser) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OrganizationUser) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OrganizationUser) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

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

### HasCreatedAt

`func (o *OrganizationUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasName

`func (o *OrganizationUser) HasName() bool`

HasName returns a boolean if a field has been set.

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

### GetOrganizationKey

`func (o *OrganizationUser) GetOrganizationKey() string`

GetOrganizationKey returns the OrganizationKey field if non-nil, zero value otherwise.

### GetOrganizationKeyOk

`func (o *OrganizationUser) GetOrganizationKeyOk() (*string, bool)`

GetOrganizationKeyOk returns a tuple with the OrganizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationKey

`func (o *OrganizationUser) SetOrganizationKey(v string)`

SetOrganizationKey sets OrganizationKey field to given value.

### HasOrganizationKey

`func (o *OrganizationUser) HasOrganizationKey() bool`

HasOrganizationKey returns a boolean if a field has been set.

### GetOrganizationName

`func (o *OrganizationUser) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *OrganizationUser) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *OrganizationUser) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *OrganizationUser) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetOrganizationScopes

`func (o *OrganizationUser) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *OrganizationUser) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *OrganizationUser) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.

### HasOrganizationScopes

`func (o *OrganizationUser) HasOrganizationScopes() bool`

HasOrganizationScopes returns a boolean if a field has been set.

### GetProjects

`func (o *OrganizationUser) GetProjects() []OrganizationUserProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationUser) GetProjectsOk() (*[]OrganizationUserProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationUser) SetProjects(v []OrganizationUserProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationUser) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


