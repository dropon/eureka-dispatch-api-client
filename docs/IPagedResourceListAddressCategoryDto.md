# IPagedResourceListAddressCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AddressCategoryDto**](AddressCategoryDto.md) |  | [optional] 
**Paging** | Pointer to [**PagingInformations**](PagingInformations.md) |  | [optional] 

## Methods

### NewIPagedResourceListAddressCategoryDto

`func NewIPagedResourceListAddressCategoryDto() *IPagedResourceListAddressCategoryDto`

NewIPagedResourceListAddressCategoryDto instantiates a new IPagedResourceListAddressCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPagedResourceListAddressCategoryDtoWithDefaults

`func NewIPagedResourceListAddressCategoryDtoWithDefaults() *IPagedResourceListAddressCategoryDto`

NewIPagedResourceListAddressCategoryDtoWithDefaults instantiates a new IPagedResourceListAddressCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IPagedResourceListAddressCategoryDto) GetData() []AddressCategoryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IPagedResourceListAddressCategoryDto) GetDataOk() (*[]AddressCategoryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IPagedResourceListAddressCategoryDto) SetData(v []AddressCategoryDto)`

SetData sets Data field to given value.

### HasData

`func (o *IPagedResourceListAddressCategoryDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *IPagedResourceListAddressCategoryDto) GetPaging() PagingInformations`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *IPagedResourceListAddressCategoryDto) GetPagingOk() (*PagingInformations, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *IPagedResourceListAddressCategoryDto) SetPaging(v PagingInformations)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *IPagedResourceListAddressCategoryDto) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


