# CommitDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | [**CommitAuthorDraft**](CommitAuthorDraft.md) |  | 
**Committer** | [**CommitAuthorDraft**](CommitAuthorDraft.md) |  | 
**Commit** | **string** |  | 
**Parents** | **[]string** |  | 
**Subject** | **string** |  | 

## Methods

### NewCommitDraft

`func NewCommitDraft(author CommitAuthorDraft, committer CommitAuthorDraft, commit string, parents []string, subject string, ) *CommitDraft`

NewCommitDraft instantiates a new CommitDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDraftWithDefaults

`func NewCommitDraftWithDefaults() *CommitDraft`

NewCommitDraftWithDefaults instantiates a new CommitDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CommitDraft) GetAuthor() CommitAuthorDraft`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitDraft) GetAuthorOk() (*CommitAuthorDraft, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitDraft) SetAuthor(v CommitAuthorDraft)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *CommitDraft) GetCommitter() CommitAuthorDraft`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitDraft) GetCommitterOk() (*CommitAuthorDraft, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitDraft) SetCommitter(v CommitAuthorDraft)`

SetCommitter sets Committer field to given value.


### GetCommit

`func (o *CommitDraft) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitDraft) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitDraft) SetCommit(v string)`

SetCommit sets Commit field to given value.


### GetParents

`func (o *CommitDraft) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *CommitDraft) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *CommitDraft) SetParents(v []string)`

SetParents sets Parents field to given value.


### GetSubject

`func (o *CommitDraft) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommitDraft) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommitDraft) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


