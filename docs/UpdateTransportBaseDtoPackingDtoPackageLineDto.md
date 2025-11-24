# UpdateTransportBaseDtoPackingDtoPackageLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeAction** | Pointer to **string** | Merge action to perform : add, update, remove. | [optional] 
**PackageLineId** | Pointer to **int32** | This property is an entity identifier. | [optional] 
**NatureId** | Pointer to **int32** |  | [optional] 
**Nature** | Pointer to **string** | The access to this property is rejected if the user has at least one of the following restriction(s):  Disable free entry in package nature. | [optional] 
**Quantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**DimensionUnit** | Pointer to **string** | Allowed dimension units are m, dm, cm and mm. | [optional] 
**Height** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**Width** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**Length** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**BarCode** | Pointer to **string** |  | [optional] 
**PackageCode** | Pointer to **string** |  | [optional] 
**PackageSizes** | Pointer to [**[]UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto**](UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto.md) |  | [optional] 
**References** | Pointer to [**[]UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto**](UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto.md) |  | [optional] 

## Methods

### NewUpdateTransportBaseDtoPackingDtoPackageLineDto

`func NewUpdateTransportBaseDtoPackingDtoPackageLineDto() *UpdateTransportBaseDtoPackingDtoPackageLineDto`

NewUpdateTransportBaseDtoPackingDtoPackageLineDto instantiates a new UpdateTransportBaseDtoPackingDtoPackageLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransportBaseDtoPackingDtoPackageLineDtoWithDefaults

`func NewUpdateTransportBaseDtoPackingDtoPackageLineDtoWithDefaults() *UpdateTransportBaseDtoPackingDtoPackageLineDto`

NewUpdateTransportBaseDtoPackingDtoPackageLineDtoWithDefaults instantiates a new UpdateTransportBaseDtoPackingDtoPackageLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetPackageLineId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageLineId() int32`

GetPackageLineId returns the PackageLineId field if non-nil, zero value otherwise.

### GetPackageLineIdOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageLineIdOk() (*int32, bool)`

GetPackageLineIdOk returns a tuple with the PackageLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetPackageLineId(v int32)`

SetPackageLineId sets PackageLineId field to given value.

### HasPackageLineId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasPackageLineId() bool`

HasPackageLineId returns a boolean if a field has been set.

### GetNatureId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetNatureId() int32`

GetNatureId returns the NatureId field if non-nil, zero value otherwise.

### GetNatureIdOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetNatureIdOk() (*int32, bool)`

GetNatureIdOk returns a tuple with the NatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatureId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetNatureId(v int32)`

SetNatureId sets NatureId field to given value.

### HasNatureId

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasNatureId() bool`

HasNatureId returns a boolean if a field has been set.

### GetNature

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetNature() string`

GetNature returns the Nature field if non-nil, zero value otherwise.

### GetNatureOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetNatureOk() (*string, bool)`

GetNatureOk returns a tuple with the Nature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNature

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetNature(v string)`

SetNature sets Nature field to given value.

### HasNature

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasNature() bool`

HasNature returns a boolean if a field has been set.

### GetQuantity

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetHeight

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetBarCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetBarCode() string`

GetBarCode returns the BarCode field if non-nil, zero value otherwise.

### GetBarCodeOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetBarCodeOk() (*string, bool)`

GetBarCodeOk returns a tuple with the BarCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetBarCode(v string)`

SetBarCode sets BarCode field to given value.

### HasBarCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasBarCode() bool`

HasBarCode returns a boolean if a field has been set.

### GetPackageCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetPackageSizes

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageSizes() []UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto`

GetPackageSizes returns the PackageSizes field if non-nil, zero value otherwise.

### GetPackageSizesOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetPackageSizesOk() (*[]UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto, bool)`

GetPackageSizesOk returns a tuple with the PackageSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSizes

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetPackageSizes(v []UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto)`

SetPackageSizes sets PackageSizes field to given value.

### HasPackageSizes

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasPackageSizes() bool`

HasPackageSizes returns a boolean if a field has been set.

### GetReferences

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetReferences() []UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) GetReferencesOk() (*[]UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) SetReferences(v []UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto)`

SetReferences sets References field to given value.

### HasReferences

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDto) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


