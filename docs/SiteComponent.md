# SiteComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Key** | **string** | key of the site component | 
**Name** | **string** | name of the site component | 
**SiteId** | **string** | id of the site | 
**ComponentId** | **string** | id of the component | 

## Methods

### NewSiteComponent

`func NewSiteComponent(id string, createdAt time.Time, key string, name string, siteId string, componentId string, ) *SiteComponent`

NewSiteComponent instantiates a new SiteComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteComponentWithDefaults

`func NewSiteComponentWithDefaults() *SiteComponent`

NewSiteComponentWithDefaults instantiates a new SiteComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteComponent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteComponent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteComponent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SiteComponent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SiteComponent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SiteComponent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetKey

`func (o *SiteComponent) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SiteComponent) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SiteComponent) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SiteComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteComponent) SetName(v string)`

SetName sets Name field to given value.


### GetSiteId

`func (o *SiteComponent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SiteComponent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SiteComponent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetComponentId

`func (o *SiteComponent) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *SiteComponent) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *SiteComponent) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


