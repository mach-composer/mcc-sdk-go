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
)

// PatchRequestInner - struct for PatchRequestInner
type PatchRequestInner struct {
	JSONPatchRequestAddReplaceTest *JSONPatchRequestAddReplaceTest
	JSONPatchRequestMoveCopy       *JSONPatchRequestMoveCopy
	JSONPatchRequestRemove         *JSONPatchRequestRemove
}

// JSONPatchRequestAddReplaceTestAsPatchRequestInner is a convenience function that returns JSONPatchRequestAddReplaceTest wrapped in PatchRequestInner
func JSONPatchRequestAddReplaceTestAsPatchRequestInner(v *JSONPatchRequestAddReplaceTest) PatchRequestInner {
	return PatchRequestInner{
		JSONPatchRequestAddReplaceTest: v,
	}
}

// JSONPatchRequestMoveCopyAsPatchRequestInner is a convenience function that returns JSONPatchRequestMoveCopy wrapped in PatchRequestInner
func JSONPatchRequestMoveCopyAsPatchRequestInner(v *JSONPatchRequestMoveCopy) PatchRequestInner {
	return PatchRequestInner{
		JSONPatchRequestMoveCopy: v,
	}
}

// JSONPatchRequestRemoveAsPatchRequestInner is a convenience function that returns JSONPatchRequestRemove wrapped in PatchRequestInner
func JSONPatchRequestRemoveAsPatchRequestInner(v *JSONPatchRequestRemove) PatchRequestInner {
	return PatchRequestInner{
		JSONPatchRequestRemove: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchRequestInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into JSONPatchRequestAddReplaceTest
	err = newStrictDecoder(data).Decode(&dst.JSONPatchRequestAddReplaceTest)
	if err == nil {
		jsonJSONPatchRequestAddReplaceTest, _ := json.Marshal(dst.JSONPatchRequestAddReplaceTest)
		if string(jsonJSONPatchRequestAddReplaceTest) == "{}" { // empty struct
			dst.JSONPatchRequestAddReplaceTest = nil
		} else {
			match++
		}
	} else {
		dst.JSONPatchRequestAddReplaceTest = nil
	}

	// try to unmarshal data into JSONPatchRequestMoveCopy
	err = newStrictDecoder(data).Decode(&dst.JSONPatchRequestMoveCopy)
	if err == nil {
		jsonJSONPatchRequestMoveCopy, _ := json.Marshal(dst.JSONPatchRequestMoveCopy)
		if string(jsonJSONPatchRequestMoveCopy) == "{}" { // empty struct
			dst.JSONPatchRequestMoveCopy = nil
		} else {
			match++
		}
	} else {
		dst.JSONPatchRequestMoveCopy = nil
	}

	// try to unmarshal data into JSONPatchRequestRemove
	err = newStrictDecoder(data).Decode(&dst.JSONPatchRequestRemove)
	if err == nil {
		jsonJSONPatchRequestRemove, _ := json.Marshal(dst.JSONPatchRequestRemove)
		if string(jsonJSONPatchRequestRemove) == "{}" { // empty struct
			dst.JSONPatchRequestRemove = nil
		} else {
			match++
		}
	} else {
		dst.JSONPatchRequestRemove = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.JSONPatchRequestAddReplaceTest = nil
		dst.JSONPatchRequestMoveCopy = nil
		dst.JSONPatchRequestRemove = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchRequestInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchRequestInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchRequestInner) MarshalJSON() ([]byte, error) {
	if src.JSONPatchRequestAddReplaceTest != nil {
		return json.Marshal(&src.JSONPatchRequestAddReplaceTest)
	}

	if src.JSONPatchRequestMoveCopy != nil {
		return json.Marshal(&src.JSONPatchRequestMoveCopy)
	}

	if src.JSONPatchRequestRemove != nil {
		return json.Marshal(&src.JSONPatchRequestRemove)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchRequestInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.JSONPatchRequestAddReplaceTest != nil {
		return obj.JSONPatchRequestAddReplaceTest
	}

	if obj.JSONPatchRequestMoveCopy != nil {
		return obj.JSONPatchRequestMoveCopy
	}

	if obj.JSONPatchRequestRemove != nil {
		return obj.JSONPatchRequestRemove
	}

	// all schemas are nil
	return nil
}

type NullablePatchRequestInner struct {
	value *PatchRequestInner
	isSet bool
}

func (v NullablePatchRequestInner) Get() *PatchRequestInner {
	return v.value
}

func (v *NullablePatchRequestInner) Set(val *PatchRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRequestInner(val *PatchRequestInner) *NullablePatchRequestInner {
	return &NullablePatchRequestInner{value: val, isSet: true}
}

func (v NullablePatchRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
