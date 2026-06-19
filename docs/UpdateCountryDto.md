# UpdateCountryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Country default label.  When this value is provided, it overrides the localized labels for user&#39;s and global settings&#39; locale. | [optional] 
**LocalizedLabels** | Pointer to [**[]LocalizedLabelDto**](LocalizedLabelDto.md) | Country label localizations.  When this property is set, the full collection is replaced with the provided value. | [optional] 
**Iso3Code** | Pointer to **string** | Country iso code in ISO 3166-1 alpha-3 format. | [optional] 
**DefaultVatCode** | Pointer to **string** |  | [optional] 
**CountryFamilyCode** | Pointer to **string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateCountryDto

`func NewUpdateCountryDto() *UpdateCountryDto`

NewUpdateCountryDto instantiates a new UpdateCountryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCountryDtoWithDefaults

`func NewUpdateCountryDtoWithDefaults() *UpdateCountryDto`

NewUpdateCountryDtoWithDefaults instantiates a new UpdateCountryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateCountryDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateCountryDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateCountryDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateCountryDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocalizedLabels

`func (o *UpdateCountryDto) GetLocalizedLabels() []LocalizedLabelDto`

GetLocalizedLabels returns the LocalizedLabels field if non-nil, zero value otherwise.

### GetLocalizedLabelsOk

`func (o *UpdateCountryDto) GetLocalizedLabelsOk() (*[]LocalizedLabelDto, bool)`

GetLocalizedLabelsOk returns a tuple with the LocalizedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedLabels

`func (o *UpdateCountryDto) SetLocalizedLabels(v []LocalizedLabelDto)`

SetLocalizedLabels sets LocalizedLabels field to given value.

### HasLocalizedLabels

`func (o *UpdateCountryDto) HasLocalizedLabels() bool`

HasLocalizedLabels returns a boolean if a field has been set.

### GetIso3Code

`func (o *UpdateCountryDto) GetIso3Code() string`

GetIso3Code returns the Iso3Code field if non-nil, zero value otherwise.

### GetIso3CodeOk

`func (o *UpdateCountryDto) GetIso3CodeOk() (*string, bool)`

GetIso3CodeOk returns a tuple with the Iso3Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3Code

`func (o *UpdateCountryDto) SetIso3Code(v string)`

SetIso3Code sets Iso3Code field to given value.

### HasIso3Code

`func (o *UpdateCountryDto) HasIso3Code() bool`

HasIso3Code returns a boolean if a field has been set.

### GetDefaultVatCode

`func (o *UpdateCountryDto) GetDefaultVatCode() string`

GetDefaultVatCode returns the DefaultVatCode field if non-nil, zero value otherwise.

### GetDefaultVatCodeOk

`func (o *UpdateCountryDto) GetDefaultVatCodeOk() (*string, bool)`

GetDefaultVatCodeOk returns a tuple with the DefaultVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVatCode

`func (o *UpdateCountryDto) SetDefaultVatCode(v string)`

SetDefaultVatCode sets DefaultVatCode field to given value.

### HasDefaultVatCode

`func (o *UpdateCountryDto) HasDefaultVatCode() bool`

HasDefaultVatCode returns a boolean if a field has been set.

### GetCountryFamilyCode

`func (o *UpdateCountryDto) GetCountryFamilyCode() string`

GetCountryFamilyCode returns the CountryFamilyCode field if non-nil, zero value otherwise.

### GetCountryFamilyCodeOk

`func (o *UpdateCountryDto) GetCountryFamilyCodeOk() (*string, bool)`

GetCountryFamilyCodeOk returns a tuple with the CountryFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryFamilyCode

`func (o *UpdateCountryDto) SetCountryFamilyCode(v string)`

SetCountryFamilyCode sets CountryFamilyCode field to given value.

### HasCountryFamilyCode

`func (o *UpdateCountryDto) HasCountryFamilyCode() bool`

HasCountryFamilyCode returns a boolean if a field has been set.

### GetLanguageCode

`func (o *UpdateCountryDto) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *UpdateCountryDto) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *UpdateCountryDto) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *UpdateCountryDto) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


