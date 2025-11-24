# CreateTransportDtoPackingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**[]CreateTransportDtoPackingDtoPackageLineDto**](CreateTransportDtoPackingDtoPackageLineDto.md) |  | [optional] 
**TotalBulkSizes** | Pointer to [**[]CreateTransportDtoPackingDtoTotalBulkSizeDto**](CreateTransportDtoPackingDtoTotalBulkSizeDto.md) |  | [optional] 

## Methods

### NewCreateTransportDtoPackingDto

`func NewCreateTransportDtoPackingDto() *CreateTransportDtoPackingDto`

NewCreateTransportDtoPackingDto instantiates a new CreateTransportDtoPackingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoPackingDtoWithDefaults

`func NewCreateTransportDtoPackingDtoWithDefaults() *CreateTransportDtoPackingDto`

NewCreateTransportDtoPackingDtoWithDefaults instantiates a new CreateTransportDtoPackingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *CreateTransportDtoPackingDto) GetLines() []CreateTransportDtoPackingDtoPackageLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *CreateTransportDtoPackingDto) GetLinesOk() (*[]CreateTransportDtoPackingDtoPackageLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *CreateTransportDtoPackingDto) SetLines(v []CreateTransportDtoPackingDtoPackageLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *CreateTransportDtoPackingDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetTotalBulkSizes

`func (o *CreateTransportDtoPackingDto) GetTotalBulkSizes() []CreateTransportDtoPackingDtoTotalBulkSizeDto`

GetTotalBulkSizes returns the TotalBulkSizes field if non-nil, zero value otherwise.

### GetTotalBulkSizesOk

`func (o *CreateTransportDtoPackingDto) GetTotalBulkSizesOk() (*[]CreateTransportDtoPackingDtoTotalBulkSizeDto, bool)`

GetTotalBulkSizesOk returns a tuple with the TotalBulkSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBulkSizes

`func (o *CreateTransportDtoPackingDto) SetTotalBulkSizes(v []CreateTransportDtoPackingDtoTotalBulkSizeDto)`

SetTotalBulkSizes sets TotalBulkSizes field to given value.

### HasTotalBulkSizes

`func (o *CreateTransportDtoPackingDto) HasTotalBulkSizes() bool`

HasTotalBulkSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


