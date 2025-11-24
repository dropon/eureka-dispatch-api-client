# UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeAction** | Pointer to **string** | Merge action to perform : add, update, remove. | [optional] 
**UnitCode** | Pointer to **string** | This property is an entity identifier. | [optional] 
**Value** | Pointer to **float64** | This number can have up to 3 decimals. | [optional] 
**IsVolumeLineTotalValue** | Pointer to **bool** | Indicates if the volume corresponds to the total volume of this package.  In this case, the volume is used (without multiplying by the quantity) when the total size is computed. | [optional] 
**IsEquivalenceEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto

`func NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto() *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto`

NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto instantiates a new UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults

`func NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults() *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto`

NewUpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDtoWithDefaults instantiates a new UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetUnitCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsVolumeLineTotalValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsVolumeLineTotalValue() bool`

GetIsVolumeLineTotalValue returns the IsVolumeLineTotalValue field if non-nil, zero value otherwise.

### GetIsVolumeLineTotalValueOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsVolumeLineTotalValueOk() (*bool, bool)`

GetIsVolumeLineTotalValueOk returns a tuple with the IsVolumeLineTotalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVolumeLineTotalValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) SetIsVolumeLineTotalValue(v bool)`

SetIsVolumeLineTotalValue sets IsVolumeLineTotalValue field to given value.

### HasIsVolumeLineTotalValue

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) HasIsVolumeLineTotalValue() bool`

HasIsVolumeLineTotalValue returns a boolean if a field has been set.

### GetIsEquivalenceEnabled

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsEquivalenceEnabled() bool`

GetIsEquivalenceEnabled returns the IsEquivalenceEnabled field if non-nil, zero value otherwise.

### GetIsEquivalenceEnabledOk

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) GetIsEquivalenceEnabledOk() (*bool, bool)`

GetIsEquivalenceEnabledOk returns a tuple with the IsEquivalenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEquivalenceEnabled

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) SetIsEquivalenceEnabled(v bool)`

SetIsEquivalenceEnabled sets IsEquivalenceEnabled field to given value.

### HasIsEquivalenceEnabled

`func (o *UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto) HasIsEquivalenceEnabled() bool`

HasIsEquivalenceEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


