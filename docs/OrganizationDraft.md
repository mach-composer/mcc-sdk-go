# OrganizationDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The organization key (must be unique) | 
**Name** | **string** | The name of the organization | 

## Methods

### NewOrganizationDraft

`func NewOrganizationDraft(key string, name string, ) *OrganizationDraft`

NewOrganizationDraft instantiates a new OrganizationDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDraftWithDefaults

`func NewOrganizationDraftWithDefaults() *OrganizationDraft`

NewOrganizationDraftWithDefaults instantiates a new OrganizationDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OrganizationDraft) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OrganizationDraft) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OrganizationDraft) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *OrganizationDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDraft) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


