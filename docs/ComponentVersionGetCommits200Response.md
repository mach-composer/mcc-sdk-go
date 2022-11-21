# ComponentVersionGetCommits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]CommitData**](CommitData.md) |  | [optional] 

## Methods

### NewComponentVersionGetCommits200Response

`func NewComponentVersionGetCommits200Response() *ComponentVersionGetCommits200Response`

NewComponentVersionGetCommits200Response instantiates a new ComponentVersionGetCommits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionGetCommits200ResponseWithDefaults

`func NewComponentVersionGetCommits200ResponseWithDefaults() *ComponentVersionGetCommits200Response`

NewComponentVersionGetCommits200ResponseWithDefaults instantiates a new ComponentVersionGetCommits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *ComponentVersionGetCommits200Response) GetCommits() []CommitData`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *ComponentVersionGetCommits200Response) GetCommitsOk() (*[]CommitData, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *ComponentVersionGetCommits200Response) SetCommits(v []CommitData)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *ComponentVersionGetCommits200Response) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


