# JSONPatchRequestAddReplaceTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer path. | 
**Value** | **interface{}** | The value to add, replace or test. | 
**Op** | **string** | The operation to perform. | 

## Methods

### NewJSONPatchRequestAddReplaceTest

`func NewJSONPatchRequestAddReplaceTest(path string, value interface{}, op string, ) *JSONPatchRequestAddReplaceTest`

NewJSONPatchRequestAddReplaceTest instantiates a new JSONPatchRequestAddReplaceTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchRequestAddReplaceTestWithDefaults

`func NewJSONPatchRequestAddReplaceTestWithDefaults() *JSONPatchRequestAddReplaceTest`

NewJSONPatchRequestAddReplaceTestWithDefaults instantiates a new JSONPatchRequestAddReplaceTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JSONPatchRequestAddReplaceTest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchRequestAddReplaceTest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchRequestAddReplaceTest) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *JSONPatchRequestAddReplaceTest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JSONPatchRequestAddReplaceTest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JSONPatchRequestAddReplaceTest) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JSONPatchRequestAddReplaceTest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JSONPatchRequestAddReplaceTest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOp

`func (o *JSONPatchRequestAddReplaceTest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchRequestAddReplaceTest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchRequestAddReplaceTest) SetOp(v string)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


