# CommitData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | **string** |  | 
**Parents** | Pointer to **[]string** |  | [optional] 
**Subject** | **string** |  | 
**Author** | [**CommitDataAuthor**](CommitDataAuthor.md) |  | 
**Committer** | [**CommitDataAuthor**](CommitDataAuthor.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCommitData

`func NewCommitData(commit string, subject string, author CommitDataAuthor, committer CommitDataAuthor, ) *CommitData`

NewCommitData instantiates a new CommitData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDataWithDefaults

`func NewCommitDataWithDefaults() *CommitData`

NewCommitDataWithDefaults instantiates a new CommitData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *CommitData) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitData) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitData) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetParents

`func (o *CommitData) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *CommitData) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *CommitData) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *CommitData) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetSubject

`func (o *CommitData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommitData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommitData) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetAuthor

`func (o *CommitData) GetAuthor() CommitDataAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitData) GetAuthorOk() (*CommitDataAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitData) SetAuthor(v CommitDataAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *CommitData) GetCommitter() CommitDataAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitData) GetCommitterOk() (*CommitDataAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitData) SetCommitter(v CommitDataAuthor)`

SetCommitter sets Committer field to given value.


### GetCreatedAt

`func (o *CommitData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommitData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommitData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommitData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


