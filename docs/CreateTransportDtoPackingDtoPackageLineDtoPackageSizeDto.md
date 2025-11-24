# CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitCode** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** | This number can have up to 3 decimals. | [optional] 
**IsVolumeLineTotalValue** | Pointer to **bool** | Indicates if the volume corresponds to the total volume of this package.  In this case, the volume is used (without multiplying by the quantity) when the total size is computed. | [optional] 
**IsEquivalenceEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto

`func NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto() *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto`

NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto instantiates a new CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults

`func NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults() *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto`

NewCreateTransportDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults instantiates a new CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitCode

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsVolumeLineTotalValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsVolumeLineTotalValue() bool`

GetIsVolumeLineTotalValue returns the IsVolumeLineTotalValue field if non-nil, zero value otherwise.

### GetIsVolumeLineTotalValueOk

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsVolumeLineTotalValueOk() (*bool, bool)`

GetIsVolumeLineTotalValueOk returns a tuple with the IsVolumeLineTotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVolumeLineTotalValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) SetIsVolumeLineTotalValue(v bool)`

SetIsVolumeLineTotalValue sets IsVolumeLineTotalValue field to given value.

### HasIsVolumeLineTotalValue

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) HasIsVolumeLineTotalValue() bool`

HasIsVolumeLineTotalValue returns a boolean if a field has been set.

### GetIsEquivalenceEnabled

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsEquivalenceEnabled() bool`

GetIsEquivalenceEnabled returns the IsEquivalenceEnabled field if non-nil, zero value otherwise.

### GetIsEquivalenceEnabledOk

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsEquivalenceEnabledOk() (*bool, bool)`

GetIsEquivalenceEnabledOk returns a tuple with the IsEquivalenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEquivalenceEnabled

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) SetIsEquivalenceEnabled(v bool)`

SetIsEquivalenceEnabled sets IsEquivalenceEnabled field to given value.

### HasIsEquivalenceEnabled

`func (o *CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto) HasIsEquivalenceEnabled() bool`

HasIsEquivalenceEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


