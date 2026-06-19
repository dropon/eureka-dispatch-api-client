# LocalizedLabelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewLocalizedLabelDto

`func NewLocalizedLabelDto() *LocalizedLabelDto`

NewLocalizedLabelDto instantiates a new LocalizedLabelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalizedLabelDtoWithDefaults

`func NewLocalizedLabelDtoWithDefaults() *LocalizedLabelDto`

NewLocalizedLabelDtoWithDefaults instantiates a new LocalizedLabelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *LocalizedLabelDto) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LocalizedLabelDto) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LocalizedLabelDto) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *LocalizedLabelDto) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetValue

`func (o *LocalizedLabelDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LocalizedLabelDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LocalizedLabelDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LocalizedLabelDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


