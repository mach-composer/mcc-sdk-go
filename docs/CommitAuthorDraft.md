# CommitAuthorDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Date** | **time.Time** |  | 

## Methods

### NewCommitAuthorDraft

`func NewCommitAuthorDraft(name string, email string, date time.Time, ) *CommitAuthorDraft`

NewCommitAuthorDraft instantiates a new CommitAuthorDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitAuthorDraftWithDefaults

`func NewCommitAuthorDraftWithDefaults() *CommitAuthorDraft`

NewCommitAuthorDraftWithDefaults instantiates a new CommitAuthorDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommitAuthorDraft) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommitAuthorDraft) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommitAuthorDraft) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CommitAuthorDraft) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommitAuthorDraft) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommitAuthorDraft) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDate

`func (o *CommitAuthorDraft) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommitAuthorDraft) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommitAuthorDraft) SetDate(v time.Time)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


