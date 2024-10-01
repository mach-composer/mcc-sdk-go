# PatchedApiClientDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LastUsedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPatchedApiClientDraft

`func NewPatchedApiClientDraft() *PatchedApiClientDraft`

NewPatchedApiClientDraft instantiates a new PatchedApiClientDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApiClientDraftWithDefaults

`func NewPatchedApiClientDraftWithDefaults() *PatchedApiClientDraft`

NewPatchedApiClientDraftWithDefaults instantiates a new PatchedApiClientDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PatchedApiClientDraft) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PatchedApiClientDraft) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PatchedApiClientDraft) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PatchedApiClientDraft) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PatchedApiClientDraft) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PatchedApiClientDraft) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PatchedApiClientDraft) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PatchedApiClientDraft) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScope

`func (o *PatchedApiClientDraft) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PatchedApiClientDraft) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PatchedApiClientDraft) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PatchedApiClientDraft) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedApiClientDraft) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedApiClientDraft) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedApiClientDraft) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedApiClientDraft) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *PatchedApiClientDraft) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *PatchedApiClientDraft) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *PatchedApiClientDraft) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *PatchedApiClientDraft) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *PatchedApiClientDraft) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *PatchedApiClientDraft) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


