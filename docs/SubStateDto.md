# SubStateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**BackgroundColor** | Pointer to **string** | Background color in ARGB hexadecimal format, example : #FFFF4500 | [optional] 
**TextColor** | Pointer to **string** | Text color in ARGB hexadecimal format, example : #FFFF4500 | [optional] 

## Methods

### NewSubStateDto

`func NewSubStateDto() *SubStateDto`

NewSubStateDto instantiates a new SubStateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubStateDtoWithDefaults

`func NewSubStateDtoWithDefaults() *SubStateDto`

NewSubStateDtoWithDefaults instantiates a new SubStateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SubStateDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubStateDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubStateDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubStateDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *SubStateDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SubStateDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SubStateDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SubStateDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsEnabled

`func (o *SubStateDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SubStateDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SubStateDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SubStateDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *SubStateDto) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *SubStateDto) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *SubStateDto) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *SubStateDto) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetTextColor

`func (o *SubStateDto) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *SubStateDto) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *SubStateDto) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *SubStateDto) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


