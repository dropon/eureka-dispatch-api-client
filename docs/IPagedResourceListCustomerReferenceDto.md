# IPagedResourceListCustomerReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CustomerReferenceDto**](CustomerReferenceDto.md) |  | [optional] 
**Paging** | Pointer to [**PagingInformations**](PagingInformations.md) |  | [optional] 

## Methods

### NewIPagedResourceListCustomerReferenceDto

`func NewIPagedResourceListCustomerReferenceDto() *IPagedResourceListCustomerReferenceDto`

NewIPagedResourceListCustomerReferenceDto instantiates a new IPagedResourceListCustomerReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPagedResourceListCustomerReferenceDtoWithDefaults

`func NewIPagedResourceListCustomerReferenceDtoWithDefaults() *IPagedResourceListCustomerReferenceDto`

NewIPagedResourceListCustomerReferenceDtoWithDefaults instantiates a new IPagedResourceListCustomerReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IPagedResourceListCustomerReferenceDto) GetData() []CustomerReferenceDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IPagedResourceListCustomerReferenceDto) GetDataOk() (*[]CustomerReferenceDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IPagedResourceListCustomerReferenceDto) SetData(v []CustomerReferenceDto)`

SetData sets Data field to given value.

### HasData

`func (o *IPagedResourceListCustomerReferenceDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *IPagedResourceListCustomerReferenceDto) GetPaging() PagingInformations`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *IPagedResourceListCustomerReferenceDto) GetPagingOk() (*PagingInformations, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *IPagedResourceListCustomerReferenceDto) SetPaging(v PagingInformations)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *IPagedResourceListCustomerReferenceDto) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


