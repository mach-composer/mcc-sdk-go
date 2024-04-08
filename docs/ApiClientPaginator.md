# ApiClientPaginator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of items in the current page | 
**Total** | **int64** | Total number of items found | 
**Offset** | **int32** |  | [default to 0]
**Limit** | **int32** |  | [default to 20]
**Results** | [**[]ApiClient**](ApiClient.md) |  | 

## Methods

### NewApiClientPaginator

`func NewApiClientPaginator(count int32, total int64, offset int32, limit int32, results []ApiClient, ) *ApiClientPaginator`

NewApiClientPaginator instantiates a new ApiClientPaginator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientPaginatorWithDefaults

`func NewApiClientPaginatorWithDefaults() *ApiClientPaginator`

NewApiClientPaginatorWithDefaults instantiates a new ApiClientPaginator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApiClientPaginator) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiClientPaginator) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiClientPaginator) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotal

`func (o *ApiClientPaginator) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiClientPaginator) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiClientPaginator) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ApiClientPaginator) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ApiClientPaginator) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ApiClientPaginator) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ApiClientPaginator) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ApiClientPaginator) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ApiClientPaginator) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *ApiClientPaginator) GetResults() []ApiClient`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiClientPaginator) GetResultsOk() (*[]ApiClient, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiClientPaginator) SetResults(v []ApiClient)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


