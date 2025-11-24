# TransportPackingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to [**CappedCollectionDtoPackageLineDto**](CappedCollectionDtoPackageLineDto.md) |  | [optional] 
**TotalBulkSizes** | Pointer to [**CappedCollectionDtoPackagesTotalBulkSizeDto**](CappedCollectionDtoPackagesTotalBulkSizeDto.md) |  | [optional] 

## Methods

### NewTransportPackingDto

`func NewTransportPackingDto() *TransportPackingDto`

NewTransportPackingDto instantiates a new TransportPackingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPackingDtoWithDefaults

`func NewTransportPackingDtoWithDefaults() *TransportPackingDto`

NewTransportPackingDtoWithDefaults instantiates a new TransportPackingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *TransportPackingDto) GetLines() CappedCollectionDtoPackageLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *TransportPackingDto) GetLinesOk() (*CappedCollectionDtoPackageLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *TransportPackingDto) SetLines(v CappedCollectionDtoPackageLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *TransportPackingDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetTotalBulkSizes

`func (o *TransportPackingDto) GetTotalBulkSizes() CappedCollectionDtoPackagesTotalBulkSizeDto`

GetTotalBulkSizes returns the TotalBulkSizes field if non-nil, zero value otherwise.

### GetTotalBulkSizesOk

`func (o *TransportPackingDto) GetTotalBulkSizesOk() (*CappedCollectionDtoPackagesTotalBulkSizeDto, bool)`

GetTotalBulkSizesOk returns a tuple with the TotalBulkSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBulkSizes

`func (o *TransportPackingDto) SetTotalBulkSizes(v CappedCollectionDtoPackagesTotalBulkSizeDto)`

SetTotalBulkSizes sets TotalBulkSizes field to given value.

### HasTotalBulkSizes

`func (o *TransportPackingDto) HasTotalBulkSizes() bool`

HasTotalBulkSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


