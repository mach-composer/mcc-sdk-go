/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.  All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"encoding/json"
)

// CommitDataDraft struct for CommitDataDraft
type CommitDataDraft struct {
	Commit    string           `json:"commit"`
	Parents   []string         `json:"parents,omitempty"`
	Subject   string           `json:"subject"`
	Author    CommitDataAuthor `json:"author"`
	Committer CommitDataAuthor `json:"committer"`
}

// NewCommitDataDraft instantiates a new CommitDataDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitDataDraft(commit string, subject string, author CommitDataAuthor, committer CommitDataAuthor) *CommitDataDraft {
	this := CommitDataDraft{}
	this.Commit = commit
	this.Subject = subject
	this.Author = author
	this.Committer = committer
	return &this
}

// NewCommitDataDraftWithDefaults instantiates a new CommitDataDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDataDraftWithDefaults() *CommitDataDraft {
	this := CommitDataDraft{}
	return &this
}

// GetCommit returns the Commit field value
func (o *CommitDataDraft) GetCommit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *CommitDataDraft) GetCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *CommitDataDraft) SetCommit(v string) {
	o.Commit = v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *CommitDataDraft) GetParents() []string {
	if o == nil || o.Parents == nil {
		var ret []string
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDataDraft) GetParentsOk() ([]string, bool) {
	if o == nil || o.Parents == nil {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *CommitDataDraft) HasParents() bool {
	if o != nil && o.Parents != nil {
		return true
	}

	return false
}

// SetParents gets a reference to the given []string and assigns it to the Parents field.
func (o *CommitDataDraft) SetParents(v []string) {
	o.Parents = v
}

// GetSubject returns the Subject field value
func (o *CommitDataDraft) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CommitDataDraft) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CommitDataDraft) SetSubject(v string) {
	o.Subject = v
}

// GetAuthor returns the Author field value
func (o *CommitDataDraft) GetAuthor() CommitDataAuthor {
	if o == nil {
		var ret CommitDataAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *CommitDataDraft) GetAuthorOk() (*CommitDataAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *CommitDataDraft) SetAuthor(v CommitDataAuthor) {
	o.Author = v
}

// GetCommitter returns the Committer field value
func (o *CommitDataDraft) GetCommitter() CommitDataAuthor {
	if o == nil {
		var ret CommitDataAuthor
		return ret
	}

	return o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value
// and a boolean to check if the value has been set.
func (o *CommitDataDraft) GetCommitterOk() (*CommitDataAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Committer, true
}

// SetCommitter sets field value
func (o *CommitDataDraft) SetCommitter(v CommitDataAuthor) {
	o.Committer = v
}

func (o CommitDataDraft) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if o.Parents != nil {
		toSerialize["parents"] = o.Parents
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["author"] = o.Author
	}
	if true {
		toSerialize["committer"] = o.Committer
	}
	return json.Marshal(toSerialize)
}

type NullableCommitDataDraft struct {
	value *CommitDataDraft
	isSet bool
}

func (v NullableCommitDataDraft) Get() *CommitDataDraft {
	return v.value
}

func (v *NullableCommitDataDraft) Set(val *CommitDataDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitDataDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitDataDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitDataDraft(val *CommitDataDraft) *NullableCommitDataDraft {
	return &NullableCommitDataDraft{value: val, isSet: true}
}

func (v NullableCommitDataDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitDataDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
