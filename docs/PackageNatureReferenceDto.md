# PackageNatureReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceIndex** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 
**IsValueChoiceForced** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to [**CappedCollectionDtoPackageNatureReferenceValueDto**](CappedCollectionDtoPackageNatureReferenceValueDto.md) |  | [optional] 

## Methods

### NewPackageNatureReferenceDto

`func NewPackageNatureReferenceDto() *PackageNatureReferenceDto`

NewPackageNatureReferenceDto instantiates a new PackageNatureReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageNatureReferenceDtoWithDefaults

`func NewPackageNatureReferenceDtoWithDefaults() *PackageNatureReferenceDto`

NewPackageNatureReferenceDtoWithDefaults instantiates a new PackageNatureReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceIndex

`func (o *PackageNatureReferenceDto) GetReferenceIndex() int32`

GetReferenceIndex returns the ReferenceIndex field if non-nil, zero value otherwise.

### GetReferenceIndexOk

`func (o *PackageNatureReferenceDto) GetReferenceIndexOk() (*int32, bool)`

GetReferenceIndexOk returns a tuple with the ReferenceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIndex

`func (o *PackageNatureReferenceDto) SetReferenceIndex(v int32)`

SetReferenceIndex sets ReferenceIndex field to given value.

### HasReferenceIndex

`func (o *PackageNatureReferenceDto) HasReferenceIndex() bool`

HasReferenceIndex returns a boolean if a field has been set.

### GetLabel

`func (o *PackageNatureReferenceDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PackageNatureReferenceDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PackageNatureReferenceDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PackageNatureReferenceDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PackageNatureReferenceDto) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PackageNatureReferenceDto) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PackageNatureReferenceDto) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PackageNatureReferenceDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsMandatory

`func (o *PackageNatureReferenceDto) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *PackageNatureReferenceDto) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *PackageNatureReferenceDto) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *PackageNatureReferenceDto) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.

### GetIsValueChoiceForced

`func (o *PackageNatureReferenceDto) GetIsValueChoiceForced() bool`

GetIsValueChoiceForced returns the IsValueChoiceForced field if non-nil, zero value otherwise.

### GetIsValueChoiceForcedOk

`func (o *PackageNatureReferenceDto) GetIsValueChoiceForcedOk() (*bool, bool)`

GetIsValueChoiceForcedOk returns a tuple with the IsValueChoiceForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueChoiceForced

`func (o *PackageNatureReferenceDto) SetIsValueChoiceForced(v bool)`

SetIsValueChoiceForced sets IsValueChoiceForced field to given value.

### HasIsValueChoiceForced

`func (o *PackageNatureReferenceDto) HasIsValueChoiceForced() bool`

HasIsValueChoiceForced returns a boolean if a field has been set.

### GetValues

`func (o *PackageNatureReferenceDto) GetValues() CappedCollectionDtoPackageNatureReferenceValueDto`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PackageNatureReferenceDto) GetValuesOk() (*CappedCollectionDtoPackageNatureReferenceValueDto, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PackageNatureReferenceDto) SetValues(v CappedCollectionDtoPackageNatureReferenceValueDto)`

SetValues sets Values field to given value.

### HasValues

`func (o *PackageNatureReferenceDto) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


