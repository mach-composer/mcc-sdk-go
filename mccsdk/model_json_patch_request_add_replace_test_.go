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

// checks if the JSONPatchRequestAddReplaceTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSONPatchRequestAddReplaceTest{}

// JSONPatchRequestAddReplaceTest struct for JSONPatchRequestAddReplaceTest
type JSONPatchRequestAddReplaceTest struct {
	// A JSON Pointer path.
	Path string `json:"path"`
	// The value to add, replace or test.
	Value interface{} `json:"value"`
	// The operation to perform.
	Op string `json:"op"`
}

type _JSONPatchRequestAddReplaceTest JSONPatchRequestAddReplaceTest

// NewJSONPatchRequestAddReplaceTest instantiates a new JSONPatchRequestAddReplaceTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONPatchRequestAddReplaceTest(path string, value interface{}, op string) *JSONPatchRequestAddReplaceTest {
	this := JSONPatchRequestAddReplaceTest{}
	this.Path = path
	this.Value = value
	this.Op = op
	return &this
}

// NewJSONPatchRequestAddReplaceTestWithDefaults instantiates a new JSONPatchRequestAddReplaceTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONPatchRequestAddReplaceTestWithDefaults() *JSONPatchRequestAddReplaceTest {
	this := JSONPatchRequestAddReplaceTest{}
	return &this
}

// GetPath returns the Path field value
func (o *JSONPatchRequestAddReplaceTest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JSONPatchRequestAddReplaceTest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JSONPatchRequestAddReplaceTest) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JSONPatchRequestAddReplaceTest) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSONPatchRequestAddReplaceTest) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *JSONPatchRequestAddReplaceTest) SetValue(v interface{}) {
	o.Value = v
}

// GetOp returns the Op field value
func (o *JSONPatchRequestAddReplaceTest) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JSONPatchRequestAddReplaceTest) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JSONPatchRequestAddReplaceTest) SetOp(v string) {
	o.Op = v
}

func (o JSONPatchRequestAddReplaceTest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSONPatchRequestAddReplaceTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["op"] = o.Op
	return toSerialize, nil
}

func (o *JSONPatchRequestAddReplaceTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"value",
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

	varJSONPatchRequestAddReplaceTest := _JSONPatchRequestAddReplaceTest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJSONPatchRequestAddReplaceTest)

	if err != nil {
		return err
	}

	*o = JSONPatchRequestAddReplaceTest(varJSONPatchRequestAddReplaceTest)

	return err
}

type NullableJSONPatchRequestAddReplaceTest struct {
	value *JSONPatchRequestAddReplaceTest
	isSet bool
}

func (v NullableJSONPatchRequestAddReplaceTest) Get() *JSONPatchRequestAddReplaceTest {
	return v.value
}

func (v *NullableJSONPatchRequestAddReplaceTest) Set(val *JSONPatchRequestAddReplaceTest) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONPatchRequestAddReplaceTest) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONPatchRequestAddReplaceTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONPatchRequestAddReplaceTest(val *JSONPatchRequestAddReplaceTest) *NullableJSONPatchRequestAddReplaceTest {
	return &NullableJSONPatchRequestAddReplaceTest{value: val, isSet: true}
}

func (v NullableJSONPatchRequestAddReplaceTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONPatchRequestAddReplaceTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
