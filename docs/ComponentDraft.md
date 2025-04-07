# ComponentDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The component key (must be unique) | 
**Name** | **string** | The name of the component | 
**Description** | Pointer to **string** | The description of the component | [optional] 

## Methods

### NewComponentDraft

`func NewComponentDraft(key string, name string, ) *ComponentDraft`

NewComponentDraft instantiates a new ComponentDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentDraftWithDefaults

`func NewComponentDraftWithDefaults() *ComponentDraft`

NewComponentDraftWithDefaults instantiates a new ComponentDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ComponentDraft) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComponentDraft) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComponentDraft) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ComponentDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentDraft) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComponentDraft) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentDraft) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentDraft) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComponentDraft) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


