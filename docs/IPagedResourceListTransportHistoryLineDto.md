# IPagedResourceListTransportHistoryLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TransportHistoryLineDto**](TransportHistoryLineDto.md) |  | [optional] 
**Paging** | Pointer to [**PagingInformations**](PagingInformations.md) |  | [optional] 

## Methods

### NewIPagedResourceListTransportHistoryLineDto

`func NewIPagedResourceListTransportHistoryLineDto() *IPagedResourceListTransportHistoryLineDto`

NewIPagedResourceListTransportHistoryLineDto instantiates a new IPagedResourceListTransportHistoryLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPagedResourceListTransportHistoryLineDtoWithDefaults

`func NewIPagedResourceListTransportHistoryLineDtoWithDefaults() *IPagedResourceListTransportHistoryLineDto`

NewIPagedResourceListTransportHistoryLineDtoWithDefaults instantiates a new IPagedResourceListTransportHistoryLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IPagedResourceListTransportHistoryLineDto) GetData() []TransportHistoryLineDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IPagedResourceListTransportHistoryLineDto) GetDataOk() (*[]TransportHistoryLineDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IPagedResourceListTransportHistoryLineDto) SetData(v []TransportHistoryLineDto)`

SetData sets Data field to given value.

### HasData

`func (o *IPagedResourceListTransportHistoryLineDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *IPagedResourceListTransportHistoryLineDto) GetPaging() PagingInformations`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *IPagedResourceListTransportHistoryLineDto) GetPagingOk() (*PagingInformations, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *IPagedResourceListTransportHistoryLineDto) SetPaging(v PagingInformations)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *IPagedResourceListTransportHistoryLineDto) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


