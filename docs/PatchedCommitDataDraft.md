# PatchedCommitDataDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to **string** |  | [optional] 
**Parents** | Pointer to **[]string** |  | [optional] 
**Author** | Pointer to [**CommitDataAuthorDraft**](CommitDataAuthorDraft.md) |  | [optional] 
**Committer** | Pointer to [**CommitDataAuthorDraft**](CommitDataAuthorDraft.md) |  | [optional] 

## Methods

### NewPatchedCommitDataDraft

`func NewPatchedCommitDataDraft() *PatchedCommitDataDraft`

NewPatchedCommitDataDraft instantiates a new PatchedCommitDataDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCommitDataDraftWithDefaults

`func NewPatchedCommitDataDraftWithDefaults() *PatchedCommitDataDraft`

NewPatchedCommitDataDraftWithDefaults instantiates a new PatchedCommitDataDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *PatchedCommitDataDraft) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PatchedCommitDataDraft) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PatchedCommitDataDraft) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PatchedCommitDataDraft) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetCommit

`func (o *PatchedCommitDataDraft) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PatchedCommitDataDraft) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PatchedCommitDataDraft) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PatchedCommitDataDraft) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetParents

`func (o *PatchedCommitDataDraft) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *PatchedCommitDataDraft) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *PatchedCommitDataDraft) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *PatchedCommitDataDraft) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetAuthor

`func (o *PatchedCommitDataDraft) GetAuthor() CommitDataAuthorDraft`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PatchedCommitDataDraft) GetAuthorOk() (*CommitDataAuthorDraft, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PatchedCommitDataDraft) SetAuthor(v CommitDataAuthorDraft)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PatchedCommitDataDraft) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitter

`func (o *PatchedCommitDataDraft) GetCommitter() CommitDataAuthorDraft`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *PatchedCommitDataDraft) GetCommitterOk() (*CommitDataAuthorDraft, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *PatchedCommitDataDraft) SetCommitter(v CommitDataAuthorDraft)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *PatchedCommitDataDraft) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


