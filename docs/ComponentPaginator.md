# ComponentPaginator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of items in the current page | [optional] 
**Total** | Pointer to **int64** | Total number of items found | [optional] 
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 10]
**Results** | Pointer to [**[]Component**](Component.md) |  | [optional] 

## Methods

### NewComponentPaginator

`func NewComponentPaginator() *ComponentPaginator`

NewComponentPaginator instantiates a new ComponentPaginator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPaginatorWithDefaults

`func NewComponentPaginatorWithDefaults() *ComponentPaginator`

NewComponentPaginatorWithDefaults instantiates a new ComponentPaginator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ComponentPaginator) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComponentPaginator) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComponentPaginator) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ComponentPaginator) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *ComponentPaginator) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ComponentPaginator) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ComponentPaginator) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ComponentPaginator) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *ComponentPaginator) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ComponentPaginator) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ComponentPaginator) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ComponentPaginator) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ComponentPaginator) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ComponentPaginator) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ComponentPaginator) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ComponentPaginator) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetResults

`func (o *ComponentPaginator) GetResults() []Component`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ComponentPaginator) GetResultsOk() (*[]Component, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ComponentPaginator) SetResults(v []Component)`

SetResults sets Results field to given value.

### HasResults

`func (o *ComponentPaginator) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


