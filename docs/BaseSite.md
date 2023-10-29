# BaseSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key for the site | 
**Name** | **string** | The name for the site | 
**Description** | Pointer to **string** | The description for the site | [optional] 

## Methods

### NewBaseSite

`func NewBaseSite(key string, name string, ) *BaseSite`

NewBaseSite instantiates a new BaseSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSiteWithDefaults

`func NewBaseSiteWithDefaults() *BaseSite`

NewBaseSiteWithDefaults instantiates a new BaseSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BaseSite) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BaseSite) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BaseSite) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *BaseSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseSite) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BaseSite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseSite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseSite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseSite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


