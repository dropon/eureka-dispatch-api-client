# IPagedResourceListCompanyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CompanyDto**](CompanyDto.md) |  | [optional] 
**Paging** | Pointer to [**PagingInformations**](PagingInformations.md) |  | [optional] 

## Methods

### NewIPagedResourceListCompanyDto

`func NewIPagedResourceListCompanyDto() *IPagedResourceListCompanyDto`

NewIPagedResourceListCompanyDto instantiates a new IPagedResourceListCompanyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPagedResourceListCompanyDtoWithDefaults

`func NewIPagedResourceListCompanyDtoWithDefaults() *IPagedResourceListCompanyDto`

NewIPagedResourceListCompanyDtoWithDefaults instantiates a new IPagedResourceListCompanyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IPagedResourceListCompanyDto) GetData() []CompanyDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IPagedResourceListCompanyDto) GetDataOk() (*[]CompanyDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IPagedResourceListCompanyDto) SetData(v []CompanyDto)`

SetData sets Data field to given value.

### HasData

`func (o *IPagedResourceListCompanyDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *IPagedResourceListCompanyDto) GetPaging() PagingInformations`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *IPagedResourceListCompanyDto) GetPagingOk() (*PagingInformations, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *IPagedResourceListCompanyDto) SetPaging(v PagingInformations)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *IPagedResourceListCompanyDto) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


