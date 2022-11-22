# ComponentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Component** | **string** | key of the component | 
**Version** | **string** | version of the component | 
**Branch** | Pointer to **string** | branch of the version | [optional] 

## Methods

### NewComponentVersion

`func NewComponentVersion(createdAt time.Time, component string, version string, ) *ComponentVersion`

NewComponentVersion instantiates a new ComponentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionWithDefaults

`func NewComponentVersionWithDefaults() *ComponentVersion`

NewComponentVersionWithDefaults instantiates a new ComponentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ComponentVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetComponent

`func (o *ComponentVersion) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentVersion) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentVersion) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVersion

`func (o *ComponentVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBranch

`func (o *ComponentVersion) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ComponentVersion) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ComponentVersion) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ComponentVersion) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


