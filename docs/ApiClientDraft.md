# ApiClientDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**Scope** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LastUsedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewApiClientDraft

`func NewApiClientDraft(clientId string, clientSecret string, ) *ApiClientDraft`

NewApiClientDraft instantiates a new ApiClientDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientDraftWithDefaults

`func NewApiClientDraftWithDefaults() *ApiClientDraft`

NewApiClientDraftWithDefaults instantiates a new ApiClientDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ApiClientDraft) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiClientDraft) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiClientDraft) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ApiClientDraft) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiClientDraft) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiClientDraft) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


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

### GetLastUsedAt

`func (o *ApiClientDraft) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiClientDraft) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiClientDraft) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiClientDraft) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ApiClientDraft) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ApiClientDraft) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


