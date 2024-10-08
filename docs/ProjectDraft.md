# ProjectDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The project key (must be unique) | 
**Name** | **string** | The name of the project | 
**Description** | Pointer to **string** | The description of the project | [optional] 

## Methods

### NewProjectDraft

`func NewProjectDraft(key string, name string, ) *ProjectDraft`

NewProjectDraft instantiates a new ProjectDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDraftWithDefaults

`func NewProjectDraftWithDefaults() *ProjectDraft`

NewProjectDraftWithDefaults instantiates a new ProjectDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ProjectDraft) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectDraft) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectDraft) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ProjectDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDraft) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProjectDraft) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDraft) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDraft) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDraft) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


