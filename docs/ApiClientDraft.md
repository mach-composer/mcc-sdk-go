# ApiClientDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description about the api client | [optional] 
**Scope** | Pointer to **[]string** | Scope | [optional] 

## Methods

### NewApiClientDraft

`func NewApiClientDraft() *ApiClientDraft`

NewApiClientDraft instantiates a new ApiClientDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientDraftWithDefaults

`func NewApiClientDraftWithDefaults() *ApiClientDraft`

NewApiClientDraftWithDefaults instantiates a new ApiClientDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiClientDraft) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiClientDraft) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiClientDraft) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiClientDraft) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *ApiClientDraft) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiClientDraft) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiClientDraft) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiClientDraft) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


