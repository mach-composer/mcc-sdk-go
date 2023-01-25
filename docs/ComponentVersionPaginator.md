# ComponentVersionPaginator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | Number of items in the current page | 
**Total** | **float32** | Total number of items found | 
**Offset** | **float32** |  | [default to 0]
**Limit** | **float32** |  | [default to 20]
**Results** | [**[]ComponentVersion**](ComponentVersion.md) |  | 

## Methods

### NewComponentVersionPaginator

`func NewComponentVersionPaginator(count float32, total float32, offset float32, limit float32, results []ComponentVersion, ) *ComponentVersionPaginator`

NewComponentVersionPaginator instantiates a new ComponentVersionPaginator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionPaginatorWithDefaults

`func NewComponentVersionPaginatorWithDefaults() *ComponentVersionPaginator`

NewComponentVersionPaginatorWithDefaults instantiates a new ComponentVersionPaginator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ComponentVersionPaginator) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComponentVersionPaginator) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComponentVersionPaginator) SetCount(v float32)`

SetCount sets Count field to given value.


### GetTotal

`func (o *ComponentVersionPaginator) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ComponentVersionPaginator) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ComponentVersionPaginator) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ComponentVersionPaginator) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ComponentVersionPaginator) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ComponentVersionPaginator) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ComponentVersionPaginator) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ComponentVersionPaginator) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ComponentVersionPaginator) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetResults

`func (o *ComponentVersionPaginator) GetResults() []ComponentVersion`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ComponentVersionPaginator) GetResultsOk() (*[]ComponentVersion, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ComponentVersionPaginator) SetResults(v []ComponentVersion)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


