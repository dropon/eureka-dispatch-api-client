# CustomParameterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomParameterId** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 
**IsFreeEntry** | Pointer to **bool** |  | [optional] 
**IsAutocompleteEnabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**OperatorsVisibility** | Pointer to **string** |  | [optional] 
**OrderersVisibility** | Pointer to **string** |  | [optional] 
**DriversVisibility** | Pointer to **string** |  | [optional] 
**SatisfactionSurveyVisibility** | Pointer to **string** |  | [optional] 
**CustomParameterValues** | Pointer to [**CappedCollectionDtoCustomParameterValueDto**](CappedCollectionDtoCustomParameterValueDto.md) |  | [optional] 

## Methods

### NewCustomParameterDto

`func NewCustomParameterDto() *CustomParameterDto`

NewCustomParameterDto instantiates a new CustomParameterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomParameterDtoWithDefaults

`func NewCustomParameterDtoWithDefaults() *CustomParameterDto`

NewCustomParameterDtoWithDefaults instantiates a new CustomParameterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomParameterId

`func (o *CustomParameterDto) GetCustomParameterId() int32`

GetCustomParameterId returns the CustomParameterId field if non-nil, zero value otherwise.

### GetCustomParameterIdOk

`func (o *CustomParameterDto) GetCustomParameterIdOk() (*int32, bool)`

GetCustomParameterIdOk returns a tuple with the CustomParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameterId

`func (o *CustomParameterDto) SetCustomParameterId(v int32)`

SetCustomParameterId sets CustomParameterId field to given value.

### HasCustomParameterId

`func (o *CustomParameterDto) HasCustomParameterId() bool`

HasCustomParameterId returns a boolean if a field has been set.

### GetCode

`func (o *CustomParameterDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomParameterDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomParameterDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomParameterDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *CustomParameterDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomParameterDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomParameterDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomParameterDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDefaultValue

`func (o *CustomParameterDto) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CustomParameterDto) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CustomParameterDto) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *CustomParameterDto) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetIsMandatory

`func (o *CustomParameterDto) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *CustomParameterDto) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *CustomParameterDto) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *CustomParameterDto) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.

### GetIsFreeEntry

`func (o *CustomParameterDto) GetIsFreeEntry() bool`

GetIsFreeEntry returns the IsFreeEntry field if non-nil, zero value otherwise.

### GetIsFreeEntryOk

`func (o *CustomParameterDto) GetIsFreeEntryOk() (*bool, bool)`

GetIsFreeEntryOk returns a tuple with the IsFreeEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeEntry

`func (o *CustomParameterDto) SetIsFreeEntry(v bool)`

SetIsFreeEntry sets IsFreeEntry field to given value.

### HasIsFreeEntry

`func (o *CustomParameterDto) HasIsFreeEntry() bool`

HasIsFreeEntry returns a boolean if a field has been set.

### GetIsAutocompleteEnabled

`func (o *CustomParameterDto) GetIsAutocompleteEnabled() bool`

GetIsAutocompleteEnabled returns the IsAutocompleteEnabled field if non-nil, zero value otherwise.

### GetIsAutocompleteEnabledOk

`func (o *CustomParameterDto) GetIsAutocompleteEnabledOk() (*bool, bool)`

GetIsAutocompleteEnabledOk returns a tuple with the IsAutocompleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutocompleteEnabled

`func (o *CustomParameterDto) SetIsAutocompleteEnabled(v bool)`

SetIsAutocompleteEnabled sets IsAutocompleteEnabled field to given value.

### HasIsAutocompleteEnabled

`func (o *CustomParameterDto) HasIsAutocompleteEnabled() bool`

HasIsAutocompleteEnabled returns a boolean if a field has been set.

### GetType

`func (o *CustomParameterDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomParameterDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomParameterDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomParameterDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperatorsVisibility

`func (o *CustomParameterDto) GetOperatorsVisibility() string`

GetOperatorsVisibility returns the OperatorsVisibility field if non-nil, zero value otherwise.

### GetOperatorsVisibilityOk

`func (o *CustomParameterDto) GetOperatorsVisibilityOk() (*string, bool)`

GetOperatorsVisibilityOk returns a tuple with the OperatorsVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorsVisibility

`func (o *CustomParameterDto) SetOperatorsVisibility(v string)`

SetOperatorsVisibility sets OperatorsVisibility field to given value.

### HasOperatorsVisibility

`func (o *CustomParameterDto) HasOperatorsVisibility() bool`

HasOperatorsVisibility returns a boolean if a field has been set.

### GetOrderersVisibility

`func (o *CustomParameterDto) GetOrderersVisibility() string`

GetOrderersVisibility returns the OrderersVisibility field if non-nil, zero value otherwise.

### GetOrderersVisibilityOk

`func (o *CustomParameterDto) GetOrderersVisibilityOk() (*string, bool)`

GetOrderersVisibilityOk returns a tuple with the OrderersVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderersVisibility

`func (o *CustomParameterDto) SetOrderersVisibility(v string)`

SetOrderersVisibility sets OrderersVisibility field to given value.

### HasOrderersVisibility

`func (o *CustomParameterDto) HasOrderersVisibility() bool`

HasOrderersVisibility returns a boolean if a field has been set.

### GetDriversVisibility

`func (o *CustomParameterDto) GetDriversVisibility() string`

GetDriversVisibility returns the DriversVisibility field if non-nil, zero value otherwise.

### GetDriversVisibilityOk

`func (o *CustomParameterDto) GetDriversVisibilityOk() (*string, bool)`

GetDriversVisibilityOk returns a tuple with the DriversVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriversVisibility

`func (o *CustomParameterDto) SetDriversVisibility(v string)`

SetDriversVisibility sets DriversVisibility field to given value.

### HasDriversVisibility

`func (o *CustomParameterDto) HasDriversVisibility() bool`

HasDriversVisibility returns a boolean if a field has been set.

### GetSatisfactionSurveyVisibility

`func (o *CustomParameterDto) GetSatisfactionSurveyVisibility() string`

GetSatisfactionSurveyVisibility returns the SatisfactionSurveyVisibility field if non-nil, zero value otherwise.

### GetSatisfactionSurveyVisibilityOk

`func (o *CustomParameterDto) GetSatisfactionSurveyVisibilityOk() (*string, bool)`

GetSatisfactionSurveyVisibilityOk returns a tuple with the SatisfactionSurveyVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatisfactionSurveyVisibility

`func (o *CustomParameterDto) SetSatisfactionSurveyVisibility(v string)`

SetSatisfactionSurveyVisibility sets SatisfactionSurveyVisibility field to given value.

### HasSatisfactionSurveyVisibility

`func (o *CustomParameterDto) HasSatisfactionSurveyVisibility() bool`

HasSatisfactionSurveyVisibility returns a boolean if a field has been set.

### GetCustomParameterValues

`func (o *CustomParameterDto) GetCustomParameterValues() CappedCollectionDtoCustomParameterValueDto`

GetCustomParameterValues returns the CustomParameterValues field if non-nil, zero value otherwise.

### GetCustomParameterValuesOk

`func (o *CustomParameterDto) GetCustomParameterValuesOk() (*CappedCollectionDtoCustomParameterValueDto, bool)`

GetCustomParameterValuesOk returns a tuple with the CustomParameterValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameterValues

`func (o *CustomParameterDto) SetCustomParameterValues(v CappedCollectionDtoCustomParameterValueDto)`

SetCustomParameterValues sets CustomParameterValues field to given value.

### HasCustomParameterValues

`func (o *CustomParameterDto) HasCustomParameterValues() bool`

HasCustomParameterValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


