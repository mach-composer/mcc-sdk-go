# CommitDataAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Email** | **string** |  | 
**Date** | **time.Time** |  | 

## Methods

### NewCommitDataAuthor

`func NewCommitDataAuthor(name string, email string, date time.Time, ) *CommitDataAuthor`

NewCommitDataAuthor instantiates a new CommitDataAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDataAuthorWithDefaults

`func NewCommitDataAuthorWithDefaults() *CommitDataAuthor`

NewCommitDataAuthorWithDefaults instantiates a new CommitDataAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommitDataAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommitDataAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommitDataAuthor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CommitDataAuthor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CommitDataAuthor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CommitDataAuthor) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDate

`func (o *CommitDataAuthor) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CommitDataAuthor) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CommitDataAuthor) SetDate(v time.Time)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


