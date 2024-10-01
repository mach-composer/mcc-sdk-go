# CommitDataDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** |  | 
**Commit** | **string** |  | 
**Parents** | Pointer to **[]string** |  | [optional] 
**Author** | [**CommitDataAuthorDraft**](CommitDataAuthorDraft.md) |  | 
**Committer** | [**CommitDataAuthorDraft**](CommitDataAuthorDraft.md) |  | 

## Methods

### NewCommitDataDraft

`func NewCommitDataDraft(subject string, commit string, author CommitDataAuthorDraft, committer CommitDataAuthorDraft, ) *CommitDataDraft`

NewCommitDataDraft instantiates a new CommitDataDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDataDraftWithDefaults

`func NewCommitDataDraftWithDefaults() *CommitDataDraft`

NewCommitDataDraftWithDefaults instantiates a new CommitDataDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CommitDataDraft) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommitDataDraft) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommitDataDraft) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCommit

`func (o *CommitDataDraft) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitDataDraft) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitDataDraft) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetParents

`func (o *CommitDataDraft) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *CommitDataDraft) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *CommitDataDraft) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *CommitDataDraft) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetAuthor

`func (o *CommitDataDraft) GetAuthor() CommitDataAuthorDraft`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitDataDraft) GetAuthorOk() (*CommitDataAuthorDraft, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitDataDraft) SetAuthor(v CommitDataAuthorDraft)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *CommitDataDraft) GetCommitter() CommitDataAuthorDraft`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitDataDraft) GetCommitterOk() (*CommitDataAuthorDraft, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitDataDraft) SetCommitter(v CommitDataAuthorDraft)`

SetCommitter sets Committer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


