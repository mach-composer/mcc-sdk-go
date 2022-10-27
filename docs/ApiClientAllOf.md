# ApiClientAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description about the api client | [optional] 
**ClientId** | **string** | the client id | 
**ClientSecret** | **string** | the client id | 
**Scope** | **[]string** | Scope | 

## Methods

### NewApiClientAllOf

`func NewApiClientAllOf(clientId string, clientSecret string, scope []string, ) *ApiClientAllOf`

NewApiClientAllOf instantiates a new ApiClientAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientAllOfWithDefaults

`func NewApiClientAllOfWithDefaults() *ApiClientAllOf`

NewApiClientAllOfWithDefaults instantiates a new ApiClientAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApiClientAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiClientAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiClientAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiClientAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientId

`func (o *ApiClientAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiClientAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiClientAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ApiClientAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiClientAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiClientAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetScope

`func (o *ApiClientAllOf) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiClientAllOf) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiClientAllOf) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


