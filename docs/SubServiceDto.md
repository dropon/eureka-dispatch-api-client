# SubServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**VatCode** | Pointer to **string** |  | [optional] 
**QuantityMustBeSelected** | Pointer to **bool** |  | [optional] 
**SubServiceQuantities** | Pointer to [**CappedCollectionDtoSubServiceQuantityDto**](CappedCollectionDtoSubServiceQuantityDto.md) |  | [optional] 

## Methods

### NewSubServiceDto

`func NewSubServiceDto() *SubServiceDto`

NewSubServiceDto instantiates a new SubServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubServiceDtoWithDefaults

`func NewSubServiceDtoWithDefaults() *SubServiceDto`

NewSubServiceDtoWithDefaults instantiates a new SubServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SubServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubServiceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubServiceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *SubServiceDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SubServiceDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SubServiceDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SubServiceDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SubServiceDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SubServiceDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SubServiceDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SubServiceDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetUnitCode

`func (o *SubServiceDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *SubServiceDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *SubServiceDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *SubServiceDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetComment

`func (o *SubServiceDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SubServiceDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SubServiceDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SubServiceDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetVatCode

`func (o *SubServiceDto) GetVatCode() string`

GetVatCode returns the VatCode field if non-nil, zero value otherwise.

### GetVatCodeOk

`func (o *SubServiceDto) GetVatCodeOk() (*string, bool)`

GetVatCodeOk returns a tuple with the VatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCode

`func (o *SubServiceDto) SetVatCode(v string)`

SetVatCode sets VatCode field to given value.

### HasVatCode

`func (o *SubServiceDto) HasVatCode() bool`

HasVatCode returns a boolean if a field has been set.

### GetQuantityMustBeSelected

`func (o *SubServiceDto) GetQuantityMustBeSelected() bool`

GetQuantityMustBeSelected returns the QuantityMustBeSelected field if non-nil, zero value otherwise.

### GetQuantityMustBeSelectedOk

`func (o *SubServiceDto) GetQuantityMustBeSelectedOk() (*bool, bool)`

GetQuantityMustBeSelectedOk returns a tuple with the QuantityMustBeSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMustBeSelected

`func (o *SubServiceDto) SetQuantityMustBeSelected(v bool)`

SetQuantityMustBeSelected sets QuantityMustBeSelected field to given value.

### HasQuantityMustBeSelected

`func (o *SubServiceDto) HasQuantityMustBeSelected() bool`

HasQuantityMustBeSelected returns a boolean if a field has been set.

### GetSubServiceQuantities

`func (o *SubServiceDto) GetSubServiceQuantities() CappedCollectionDtoSubServiceQuantityDto`

GetSubServiceQuantities returns the SubServiceQuantities field if non-nil, zero value otherwise.

### GetSubServiceQuantitiesOk

`func (o *SubServiceDto) GetSubServiceQuantitiesOk() (*CappedCollectionDtoSubServiceQuantityDto, bool)`

GetSubServiceQuantitiesOk returns a tuple with the SubServiceQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServiceQuantities

`func (o *SubServiceDto) SetSubServiceQuantities(v CappedCollectionDtoSubServiceQuantityDto)`

SetSubServiceQuantities sets SubServiceQuantities field to given value.

### HasSubServiceQuantities

`func (o *SubServiceDto) HasSubServiceQuantities() bool`

HasSubServiceQuantities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


