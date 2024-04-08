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
)

// checks if the JSONPatchRequestMoveCopy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSONPatchRequestMoveCopy{}

// JSONPatchRequestMoveCopy struct for JSONPatchRequestMoveCopy
type JSONPatchRequestMoveCopy struct {
	// A JSON Pointer path.
	Path string `json:"path"`
	// The operation to perform.
	Op string `json:"op"`
}

type _JSONPatchRequestMoveCopy JSONPatchRequestMoveCopy

// NewJSONPatchRequestMoveCopy instantiates a new JSONPatchRequestMoveCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONPatchRequestMoveCopy(path string, op string) *JSONPatchRequestMoveCopy {
	this := JSONPatchRequestMoveCopy{}
	this.Path = path
	this.Op = op
	return &this
}

// NewJSONPatchRequestMoveCopyWithDefaults instantiates a new JSONPatchRequestMoveCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONPatchRequestMoveCopyWithDefaults() *JSONPatchRequestMoveCopy {
	this := JSONPatchRequestMoveCopy{}
	return &this
}

// GetPath returns the Path field value
func (o *JSONPatchRequestMoveCopy) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JSONPatchRequestMoveCopy) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JSONPatchRequestMoveCopy) SetPath(v string) {
	o.Path = v
}

// GetOp returns the Op field value
func (o *JSONPatchRequestMoveCopy) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JSONPatchRequestMoveCopy) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JSONPatchRequestMoveCopy) SetOp(v string) {
	o.Op = v
}

func (o JSONPatchRequestMoveCopy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSONPatchRequestMoveCopy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["op"] = o.Op
	return toSerialize, nil
}

func (o *JSONPatchRequestMoveCopy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"op",
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

	varJSONPatchRequestMoveCopy := _JSONPatchRequestMoveCopy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJSONPatchRequestMoveCopy)

	if err != nil {
		return err
	}

	*o = JSONPatchRequestMoveCopy(varJSONPatchRequestMoveCopy)

	return err
}

type NullableJSONPatchRequestMoveCopy struct {
	value *JSONPatchRequestMoveCopy
	isSet bool
}

func (v NullableJSONPatchRequestMoveCopy) Get() *JSONPatchRequestMoveCopy {
	return v.value
}

func (v *NullableJSONPatchRequestMoveCopy) Set(val *JSONPatchRequestMoveCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONPatchRequestMoveCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONPatchRequestMoveCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONPatchRequestMoveCopy(val *JSONPatchRequestMoveCopy) *NullableJSONPatchRequestMoveCopy {
	return &NullableJSONPatchRequestMoveCopy{value: val, isSet: true}
}

func (v NullableJSONPatchRequestMoveCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONPatchRequestMoveCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}