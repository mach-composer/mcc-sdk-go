# ComponentVersionDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | version of the component | 
**Branch** | Pointer to **string** | branch of the version | [optional] 

## Methods

### NewComponentVersionDraft

`func NewComponentVersionDraft(version string, ) *ComponentVersionDraft`

NewComponentVersionDraft instantiates a new ComponentVersionDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionDraftWithDefaults

`func NewComponentVersionDraftWithDefaults() *ComponentVersionDraft`

NewComponentVersionDraftWithDefaults instantiates a new ComponentVersionDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ComponentVersionDraft) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersionDraft) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersionDraft) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBranch

`func (o *ComponentVersionDraft) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ComponentVersionDraft) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ComponentVersionDraft) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ComponentVersionDraft) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


