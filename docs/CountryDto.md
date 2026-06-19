# CountryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Country iso code in ISO 3166-1 alpha-2 format. | [optional] 
**Label** | Pointer to **string** | Country displayed label. | [optional] 
**LocalizedLabels** | Pointer to [**[]LocalizedLabelDto**](LocalizedLabelDto.md) | Country label localizations | [optional] 
**Iso3Code** | Pointer to **string** | Country iso code in ISO 3166-1 alpha-3 format. | [optional] 
**DefaultVatCode** | Pointer to **string** |  | [optional] 
**CountryFamilyCode** | Pointer to **string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCountryDto

`func NewCountryDto() *CountryDto`

NewCountryDto instantiates a new CountryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryDtoWithDefaults

`func NewCountryDtoWithDefaults() *CountryDto`

NewCountryDtoWithDefaults instantiates a new CountryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CountryDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CountryDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *CountryDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CountryDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CountryDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CountryDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocalizedLabels

`func (o *CountryDto) GetLocalizedLabels() []LocalizedLabelDto`

GetLocalizedLabels returns the LocalizedLabels field if non-nil, zero value otherwise.

### GetLocalizedLabelsOk

`func (o *CountryDto) GetLocalizedLabelsOk() (*[]LocalizedLabelDto, bool)`

GetLocalizedLabelsOk returns a tuple with the LocalizedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedLabels

`func (o *CountryDto) SetLocalizedLabels(v []LocalizedLabelDto)`

SetLocalizedLabels sets LocalizedLabels field to given value.

### HasLocalizedLabels

`func (o *CountryDto) HasLocalizedLabels() bool`

HasLocalizedLabels returns a boolean if a field has been set.

### GetIso3Code

`func (o *CountryDto) GetIso3Code() string`

GetIso3Code returns the Iso3Code field if non-nil, zero value otherwise.

### GetIso3CodeOk

`func (o *CountryDto) GetIso3CodeOk() (*string, bool)`

GetIso3CodeOk returns a tuple with the Iso3Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3Code

`func (o *CountryDto) SetIso3Code(v string)`

SetIso3Code sets Iso3Code field to given value.

### HasIso3Code

`func (o *CountryDto) HasIso3Code() bool`

HasIso3Code returns a boolean if a field has been set.

### GetDefaultVatCode

`func (o *CountryDto) GetDefaultVatCode() string`

GetDefaultVatCode returns the DefaultVatCode field if non-nil, zero value otherwise.

### GetDefaultVatCodeOk

`func (o *CountryDto) GetDefaultVatCodeOk() (*string, bool)`

GetDefaultVatCodeOk returns a tuple with the DefaultVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVatCode

`func (o *CountryDto) SetDefaultVatCode(v string)`

SetDefaultVatCode sets DefaultVatCode field to given value.

### HasDefaultVatCode

`func (o *CountryDto) HasDefaultVatCode() bool`

HasDefaultVatCode returns a boolean if a field has been set.

### GetCountryFamilyCode

`func (o *CountryDto) GetCountryFamilyCode() string`

GetCountryFamilyCode returns the CountryFamilyCode field if non-nil, zero value otherwise.

### GetCountryFamilyCodeOk

`func (o *CountryDto) GetCountryFamilyCodeOk() (*string, bool)`

GetCountryFamilyCodeOk returns a tuple with the CountryFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryFamilyCode

`func (o *CountryDto) SetCountryFamilyCode(v string)`

SetCountryFamilyCode sets CountryFamilyCode field to given value.

### HasCountryFamilyCode

`func (o *CountryDto) HasCountryFamilyCode() bool`

HasCountryFamilyCode returns a boolean if a field has been set.

### GetLanguageCode

`func (o *CountryDto) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *CountryDto) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *CountryDto) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *CountryDto) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


