# ComponentVersionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** | key of the component | 
**Version** | **string** | version of the component | 
**Branch** | Pointer to **string** | branch of the version | [optional] 

## Methods

### NewComponentVersionAllOf

`func NewComponentVersionAllOf(component string, version string, ) *ComponentVersionAllOf`

NewComponentVersionAllOf instantiates a new ComponentVersionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionAllOfWithDefaults

`func NewComponentVersionAllOfWithDefaults() *ComponentVersionAllOf`

NewComponentVersionAllOfWithDefaults instantiates a new ComponentVersionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ComponentVersionAllOf) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentVersionAllOf) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentVersionAllOf) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVersion

`func (o *ComponentVersionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBranch

`func (o *ComponentVersionAllOf) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ComponentVersionAllOf) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ComponentVersionAllOf) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ComponentVersionAllOf) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


