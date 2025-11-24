# TransportEntryDryRunDtoPackingDtoPackageLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nature** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**DimensionUnit** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **float64** |  | [optional] 
**Width** | Pointer to **float64** |  | [optional] 
**Length** | Pointer to **float64** |  | [optional] 
**BarCode** | Pointer to **string** |  | [optional] 
**PackageCode** | Pointer to **string** |  | [optional] 
**PackageSizes** | Pointer to [**[]TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto**](TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto.md) |  | [optional] 
**References** | Pointer to [**[]TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto**](TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto.md) |  | [optional] 

## Methods

### NewTransportEntryDryRunDtoPackingDtoPackageLineDto

`func NewTransportEntryDryRunDtoPackingDtoPackageLineDto() *TransportEntryDryRunDtoPackingDtoPackageLineDto`

NewTransportEntryDryRunDtoPackingDtoPackageLineDto instantiates a new TransportEntryDryRunDtoPackingDtoPackageLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEntryDryRunDtoPackingDtoPackageLineDtoWithDefaults

`func NewTransportEntryDryRunDtoPackingDtoPackageLineDtoWithDefaults() *TransportEntryDryRunDtoPackingDtoPackageLineDto`

NewTransportEntryDryRunDtoPackingDtoPackageLineDtoWithDefaults instantiates a new TransportEntryDryRunDtoPackingDtoPackageLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNature

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetQuantity

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetHeight

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetBarCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetBarCode() string`

GetBarCode returns the BarCode field if non-nil, zero value otherwise.

### GetBarCodeOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetBarCodeOk() (*string, bool)`

GetBarCodeOk returns a tuple with the BarCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetBarCode(v string)`

SetBarCode sets BarCode field to given value.

### HasBarCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasBarCode() bool`

HasBarCode returns a boolean if a field has been set.

### GetPackageCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetPackageSizes

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetPackageSizes() []TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto`

GetPackageSizes returns the PackageSizes field if non-nil, zero value otherwise.

### GetPackageSizesOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetPackageSizesOk() (*[]TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto, bool)`

GetPackageSizesOk returns a tuple with the PackageSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSizes

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetPackageSizes(v []TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto)`

SetPackageSizes sets PackageSizes field to given value.

### HasPackageSizes

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasPackageSizes() bool`

HasPackageSizes returns a boolean if a field has been set.

### GetReferences

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetReferences() []TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) GetReferencesOk() (*[]TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) SetReferences(v []TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto)`

SetReferences sets References field to given value.

### HasReferences

`func (o *TransportEntryDryRunDtoPackingDtoPackageLineDto) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


