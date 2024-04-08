/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.  All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CommitData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitData{}

// CommitData struct for CommitData
type CommitData struct {
	Commit    string           `json:"commit"`
	Parents   []string         `json:"parents"`
	Subject   string           `json:"subject"`
	Author    CommitDataAuthor `json:"author"`
	Committer CommitDataAuthor `json:"committer"`
	CreatedAt *time.Time       `json:"created_at,omitempty"`
}

type _CommitData CommitData

// NewCommitData instantiates a new CommitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitData(commit string, parents []string, subject string, author CommitDataAuthor, committer CommitDataAuthor) *CommitData {
	this := CommitData{}
	this.Commit = commit
	this.Parents = parents
	this.Subject = subject
	this.Author = author
	this.Committer = committer
	return &this
}

// NewCommitDataWithDefaults instantiates a new CommitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDataWithDefaults() *CommitData {
	this := CommitData{}
	return &this
}

// GetCommit returns the Commit field value
func (o *CommitData) GetCommit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *CommitData) GetCommitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *CommitData) SetCommit(v string) {
	o.Commit = v
}

// GetParents returns the Parents field value
func (o *CommitData) GetParents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value
// and a boolean to check if the value has been set.
func (o *CommitData) GetParentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parents, true
}

// SetParents sets field value
func (o *CommitData) SetParents(v []string) {
	o.Parents = v
}

// GetSubject returns the Subject field value
func (o *CommitData) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CommitData) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CommitData) SetSubject(v string) {
	o.Subject = v
}

// GetAuthor returns the Author field value
func (o *CommitData) GetAuthor() CommitDataAuthor {
	if o == nil {
		var ret CommitDataAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *CommitData) GetAuthorOk() (*CommitDataAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *CommitData) SetAuthor(v CommitDataAuthor) {
	o.Author = v
}

// GetCommitter returns the Committer field value
func (o *CommitData) GetCommitter() CommitDataAuthor {
	if o == nil {
		var ret CommitDataAuthor
		return ret
	}

	return o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value
// and a boolean to check if the value has been set.
func (o *CommitData) GetCommitterOk() (*CommitDataAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Committer, true
}

// SetCommitter sets field value
func (o *CommitData) SetCommitter(v CommitDataAuthor) {
	o.Committer = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CommitData) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CommitData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CommitData) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o CommitData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commit"] = o.Commit
	toSerialize["parents"] = o.Parents
	toSerialize["subject"] = o.Subject
	toSerialize["author"] = o.Author
	toSerialize["committer"] = o.Committer
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *CommitData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commit",
		"parents",
		"subject",
		"author",
		"committer",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommitData := _CommitData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommitData)

	if err != nil {
		return err
	}

	*o = CommitData(varCommitData)

	return err
}

type NullableCommitData struct {
	value *CommitData
	isSet bool
}

func (v NullableCommitData) Get() *CommitData {
	return v.value
}

func (v *NullableCommitData) Set(val *CommitData) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitData) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitData(val *CommitData) *NullableCommitData {
	return &NullableCommitData{value: val, isSet: true}
}

func (v NullableCommitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
