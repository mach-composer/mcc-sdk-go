# OrganizationUserInviteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**Organization** | [**OrganizationUserInviteDataOrganization**](OrganizationUserInviteDataOrganization.md) |  | 

## Methods

### NewOrganizationUserInviteData

`func NewOrganizationUserInviteData(id string, createdBy string, organization OrganizationUserInviteDataOrganization, ) *OrganizationUserInviteData`

NewOrganizationUserInviteData instantiates a new OrganizationUserInviteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserInviteDataWithDefaults

`func NewOrganizationUserInviteDataWithDefaults() *OrganizationUserInviteData`

NewOrganizationUserInviteDataWithDefaults instantiates a new OrganizationUserInviteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationUserInviteData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationUserInviteData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationUserInviteData) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *OrganizationUserInviteData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationUserInviteData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationUserInviteData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetOrganization

`func (o *OrganizationUserInviteData) GetOrganization() OrganizationUserInviteDataOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationUserInviteData) GetOrganizationOk() (*OrganizationUserInviteDataOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationUserInviteData) SetOrganization(v OrganizationUserInviteDataOrganization)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


