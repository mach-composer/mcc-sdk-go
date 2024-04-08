# PatchRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer path. | 
**Value** | **map[string]interface{}** | The value to add, replace or test. | 
**Op** | **string** | The operation to perform. | 

## Methods

### NewPatchRequestInner

`func NewPatchRequestInner(path string, value map[string]interface{}, op string, ) *PatchRequestInner`

NewPatchRequestInner instantiates a new PatchRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestInnerWithDefaults

`func NewPatchRequestInnerWithDefaults() *PatchRequestInner`

NewPatchRequestInnerWithDefaults instantiates a new PatchRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PatchRequestInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchRequestInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchRequestInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchRequestInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchRequestInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchRequestInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetOp

`func (o *PatchRequestInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchRequestInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchRequestInner) SetOp(v string)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


