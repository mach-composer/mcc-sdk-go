# SiteDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key for the site | 
**Name** | **string** | The name for the site | 
**Description** | Pointer to **string** | The description for the site | [optional] 

## Methods

### NewSiteDraft

`func NewSiteDraft(key string, name string, ) *SiteDraft`

NewSiteDraft instantiates a new SiteDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteDraftWithDefaults

`func NewSiteDraftWithDefaults() *SiteDraft`

NewSiteDraftWithDefaults instantiates a new SiteDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SiteDraft) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SiteDraft) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SiteDraft) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SiteDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteDraft) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SiteDraft) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteDraft) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteDraft) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteDraft) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


