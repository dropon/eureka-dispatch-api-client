# CreateTransportDtoPackingDtoPackageLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatureId** | Pointer to **int32** | Identifier of the package nature.  Mandatory when Nature is not provided. | [optional] 
**Nature** | Pointer to **string** | Package nature name.  Mandatory when NatureId is not provided.    The access to this property is rejected if the user has at least one of the following restriction(s):  Disable free entry in package nature. | [optional] 
**Quantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**DimensionUnit** | Pointer to **string** | Allowed dimension units are m, dm, cm and mm. | [optional] 
**Height** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**Width** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**Length** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**BarCode** | Pointer to **string** |  | [optional] 
**PackageCode** | Pointer to **string** |  | [optional] 
**PackageSizes** | Pointer to [**[]CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto**](CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto.md) |  | [optional] 
**References** | Pointer to [**[]CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto**](CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto.md) |  | [optional] 

## Methods

### NewCreateTransportDtoPackingDtoPackageLineDto

`func NewCreateTransportDtoPackingDtoPackageLineDto() *CreateTransportDtoPackingDtoPackageLineDto`

NewCreateTransportDtoPackingDtoPackageLineDto instantiates a new CreateTransportDtoPackingDtoPackageLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoPackingDtoPackageLineDtoWithDefaults

`func NewCreateTransportDtoPackingDtoPackageLineDtoWithDefaults() *CreateTransportDtoPackingDtoPackageLineDto`

NewCreateTransportDtoPackingDtoPackageLineDtoWithDefaults instantiates a new CreateTransportDtoPackingDtoPackageLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNatureId

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetNatureId() int32`

GetNatureId returns the NatureId field if non-nil, zero value otherwise.

### GetNatureIdOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetNatureIdOk() (*int32, bool)`

GetNatureIdOk returns a tuple with the NatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatureId

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetNatureId(v int32)`

SetNatureId sets NatureId field to given value.

### HasNatureId

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasNatureId() bool`

HasNatureId returns a boolean if a field has been set.

### GetNature

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetQuantity

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetHeight

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetBarCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetBarCode() string`

GetBarCode returns the BarCode field if non-nil, zero value otherwise.

### GetBarCodeOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetBarCodeOk() (*string, bool)`

GetBarCodeOk returns a tuple with the BarCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetBarCode(v string)`

SetBarCode sets BarCode field to given value.

### HasBarCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasBarCode() bool`

HasBarCode returns a boolean if a field has been set.

### GetPackageCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetPackageSizes

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetPackageSizes() []CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto`

GetPackageSizes returns the PackageSizes field if non-nil, zero value otherwise.

### GetPackageSizesOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetPackageSizesOk() (*[]CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto, bool)`

GetPackageSizesOk returns a tuple with the PackageSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSizes

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetPackageSizes(v []CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto)`

SetPackageSizes sets PackageSizes field to given value.

### HasPackageSizes

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasPackageSizes() bool`

HasPackageSizes returns a boolean if a field has been set.

### GetReferences

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetReferences() []CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CreateTransportDtoPackingDtoPackageLineDto) GetReferencesOk() (*[]CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CreateTransportDtoPackingDtoPackageLineDto) SetReferences(v []CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto)`

SetReferences sets References field to given value.

### HasReferences

`func (o *CreateTransportDtoPackingDtoPackageLineDto) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


