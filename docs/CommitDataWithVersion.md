# CommitDataWithVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | **string** |  | 
**Parents** | **[]string** |  | 
**Subject** | **string** |  | 
**Author** | [**CommitDataAuthor**](CommitDataAuthor.md) |  | 
**Committer** | [**CommitDataAuthor**](CommitDataAuthor.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 

## Methods

### NewCommitDataWithVersion

`func NewCommitDataWithVersion(commit string, parents []string, subject string, author CommitDataAuthor, committer CommitDataAuthor, ) *CommitDataWithVersion`

NewCommitDataWithVersion instantiates a new CommitDataWithVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDataWithVersionWithDefaults

`func NewCommitDataWithVersionWithDefaults() *CommitDataWithVersion`

NewCommitDataWithVersionWithDefaults instantiates a new CommitDataWithVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *CommitDataWithVersion) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitDataWithVersion) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitDataWithVersion) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetParents

`func (o *CommitDataWithVersion) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *CommitDataWithVersion) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *CommitDataWithVersion) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetSubject

`func (o *CommitDataWithVersion) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommitDataWithVersion) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommitDataWithVersion) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetAuthor

`func (o *CommitDataWithVersion) GetAuthor() CommitDataAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitDataWithVersion) GetAuthorOk() (*CommitDataAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitDataWithVersion) SetAuthor(v CommitDataAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *CommitDataWithVersion) GetCommitter() CommitDataAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitDataWithVersion) GetCommitterOk() (*CommitDataAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitDataWithVersion) SetCommitter(v CommitDataAuthor)`

SetCommitter sets Committer field to given value.


### GetCreatedAt

`func (o *CommitDataWithVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommitDataWithVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommitDataWithVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommitDataWithVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CommitDataWithVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CommitDataWithVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CommitDataWithVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CommitDataWithVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBranch

`func (o *CommitDataWithVersion) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CommitDataWithVersion) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CommitDataWithVersion) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CommitDataWithVersion) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


