# Paginator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | Number of items in the current page | 
**Total** | **float32** | Total number of items found | 
**Offset** | **float32** |  | [default to 0]
**Limit** | **float32** |  | [default to 20]

## Methods

### NewPaginator

`func NewPaginator(count float32, total float32, offset float32, limit float32, ) *Paginator`

NewPaginator instantiates a new Paginator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatorWithDefaults

`func NewPaginatorWithDefaults() *Paginator`

NewPaginatorWithDefaults instantiates a new Paginator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Paginator) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Paginator) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Paginator) SetCount(v float32)`

SetCount sets Count field to given value.


### GetTotal

`func (o *Paginator) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Paginator) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Paginator) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *Paginator) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Paginator) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Paginator) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *Paginator) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Paginator) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Paginator) SetLimit(v float32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


