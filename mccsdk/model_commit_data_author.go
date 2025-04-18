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
	"fmt"
	"time"
)

// checks if the CommitDataAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitDataAuthor{}

// CommitDataAuthor struct for CommitDataAuthor
type CommitDataAuthor struct {
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Date                 time.Time `json:"date"`
	AdditionalProperties map[string]interface{}
}

type _CommitDataAuthor CommitDataAuthor

// NewCommitDataAuthor instantiates a new CommitDataAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitDataAuthor(name string, email string, date time.Time) *CommitDataAuthor {
	this := CommitDataAuthor{}
	this.Name = name
	this.Email = email
	this.Date = date
	return &this
}

// NewCommitDataAuthorWithDefaults instantiates a new CommitDataAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDataAuthorWithDefaults() *CommitDataAuthor {
	this := CommitDataAuthor{}
	return &this
}

// GetName returns the Name field value
func (o *CommitDataAuthor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommitDataAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CommitDataAuthor) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CommitDataAuthor) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CommitDataAuthor) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CommitDataAuthor) SetEmail(v string) {
	o.Email = v
}

// GetDate returns the Date field value
func (o *CommitDataAuthor) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *CommitDataAuthor) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *CommitDataAuthor) SetDate(v time.Time) {
	o.Date = v
}

func (o CommitDataAuthor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitDataAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["date"] = o.Date

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommitDataAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"email",
		"date",
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

	varCommitDataAuthor := _CommitDataAuthor{}

	err = json.Unmarshal(data, &varCommitDataAuthor)

	if err != nil {
		return err
	}

	*o = CommitDataAuthor(varCommitDataAuthor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommitDataAuthor struct {
	value *CommitDataAuthor
	isSet bool
}

func (v NullableCommitDataAuthor) Get() *CommitDataAuthor {
	return v.value
}

func (v *NullableCommitDataAuthor) Set(val *CommitDataAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitDataAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitDataAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitDataAuthor(val *CommitDataAuthor) *NullableCommitDataAuthor {
	return &NullableCommitDataAuthor{value: val, isSet: true}
}

func (v NullableCommitDataAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitDataAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
