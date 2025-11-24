# PackageLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLineId** | Pointer to **int32** |  | [optional] 
**Nature** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**DimensionUnit** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **float64** |  | [optional] 
**Width** | Pointer to **float64** |  | [optional] 
**Length** | Pointer to **float64** |  | [optional] 
**BarCode** | Pointer to **string** |  | [optional] 
**PackageCode** | Pointer to **string** |  | [optional] 
**PackageHistoryCode** | Pointer to **string** |  | [optional] 
**PackageSizes** | Pointer to [**CappedCollectionDtoPackageLineSizeDto**](CappedCollectionDtoPackageLineSizeDto.md) |  | [optional] 
**References** | Pointer to [**CappedCollectionDtoPackageLineReferenceValueDto**](CappedCollectionDtoPackageLineReferenceValueDto.md) |  | [optional] 

## Methods

### NewPackageLineDto

`func NewPackageLineDto() *PackageLineDto`

NewPackageLineDto instantiates a new PackageLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLineDtoWithDefaults

`func NewPackageLineDtoWithDefaults() *PackageLineDto`

NewPackageLineDtoWithDefaults instantiates a new PackageLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLineId

`func (o *PackageLineDto) GetPackageLineId() int32`

GetPackageLineId returns the PackageLineId field if non-nil, zero value otherwise.

### GetPackageLineIdOk

`func (o *PackageLineDto) GetPackageLineIdOk() (*int32, bool)`

GetPackageLineIdOk returns a tuple with the PackageLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineId

`func (o *PackageLineDto) SetPackageLineId(v int32)`

SetPackageLineId sets PackageLineId field to given value.

### HasPackageLineId

`func (o *PackageLineDto) HasPackageLineId() bool`

HasPackageLineId returns a boolean if a field has been set.

### GetNature

`func (o *PackageLineDto) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *PackageLineDto) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *PackageLineDto) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *PackageLineDto) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetQuantity

`func (o *PackageLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PackageLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PackageLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PackageLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *PackageLineDto) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *PackageLineDto) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *PackageLineDto) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *PackageLineDto) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetHeight

`func (o *PackageLineDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PackageLineDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PackageLineDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PackageLineDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *PackageLineDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PackageLineDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PackageLineDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PackageLineDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *PackageLineDto) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PackageLineDto) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PackageLineDto) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *PackageLineDto) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetBarCode

`func (o *PackageLineDto) GetBarCode() string`

GetBarCode returns the BarCode field if non-nil, zero value otherwise.

### GetBarCodeOk

`func (o *PackageLineDto) GetBarCodeOk() (*string, bool)`

GetBarCodeOk returns a tuple with the BarCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode

`func (o *PackageLineDto) SetBarCode(v string)`

SetBarCode sets BarCode field to given value.

### HasBarCode

`func (o *PackageLineDto) HasBarCode() bool`

HasBarCode returns a boolean if a field has been set.

### GetPackageCode

`func (o *PackageLineDto) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *PackageLineDto) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *PackageLineDto) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *PackageLineDto) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetPackageHistoryCode

`func (o *PackageLineDto) GetPackageHistoryCode() string`

GetPackageHistoryCode returns the PackageHistoryCode field if non-nil, zero value otherwise.

### GetPackageHistoryCodeOk

`func (o *PackageLineDto) GetPackageHistoryCodeOk() (*string, bool)`

GetPackageHistoryCodeOk returns a tuple with the PackageHistoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageHistoryCode

`func (o *PackageLineDto) SetPackageHistoryCode(v string)`

SetPackageHistoryCode sets PackageHistoryCode field to given value.

### HasPackageHistoryCode

`func (o *PackageLineDto) HasPackageHistoryCode() bool`

HasPackageHistoryCode returns a boolean if a field has been set.

### GetPackageSizes

`func (o *PackageLineDto) GetPackageSizes() CappedCollectionDtoPackageLineSizeDto`

GetPackageSizes returns the PackageSizes field if non-nil, zero value otherwise.

### GetPackageSizesOk

`func (o *PackageLineDto) GetPackageSizesOk() (*CappedCollectionDtoPackageLineSizeDto, bool)`

GetPackageSizesOk returns a tuple with the PackageSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSizes

`func (o *PackageLineDto) SetPackageSizes(v CappedCollectionDtoPackageLineSizeDto)`

SetPackageSizes sets PackageSizes field to given value.

### HasPackageSizes

`func (o *PackageLineDto) HasPackageSizes() bool`

HasPackageSizes returns a boolean if a field has been set.

### GetReferences

`func (o *PackageLineDto) GetReferences() CappedCollectionDtoPackageLineReferenceValueDto`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PackageLineDto) GetReferencesOk() (*CappedCollectionDtoPackageLineReferenceValueDto, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PackageLineDto) SetReferences(v CappedCollectionDtoPackageLineReferenceValueDto)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PackageLineDto) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


