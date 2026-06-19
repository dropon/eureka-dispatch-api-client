# CreateCountryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Mandatory. Country iso code in ISO 3166-1 alpha-2 format.  Must contain only ASCII letters (A-Z, a-z), digits (0-9), underscores (_), or hyphens (-). | 
**Label** | Pointer to **string** | Country default label.  When this value is provided, it overrides the localized labels for user&#39;s and global settings&#39; locale. | [optional] 
**LocalizedLabels** | Pointer to [**[]LocalizedLabelDto**](LocalizedLabelDto.md) | Country label localizations | [optional] 
**Iso3Code** | Pointer to **string** | Country iso code in ISO 3166-1 alpha-3 format. | [optional] 
**DefaultVatCode** | Pointer to **string** |  | [optional] 
**CountryFamilyCode** | Pointer to **string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCountryCommand

`func NewCreateCountryCommand(code string, ) *CreateCountryCommand`

NewCreateCountryCommand instantiates a new CreateCountryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCountryCommandWithDefaults

`func NewCreateCountryCommandWithDefaults() *CreateCountryCommand`

NewCreateCountryCommandWithDefaults instantiates a new CreateCountryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateCountryCommand) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCountryCommand) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCountryCommand) SetCode(v string)`

SetCode sets Code field to given value.


### GetLabel

`func (o *CreateCountryCommand) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateCountryCommand) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateCountryCommand) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateCountryCommand) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocalizedLabels

`func (o *CreateCountryCommand) GetLocalizedLabels() []LocalizedLabelDto`

GetLocalizedLabels returns the LocalizedLabels field if non-nil, zero value otherwise.

### GetLocalizedLabelsOk

`func (o *CreateCountryCommand) GetLocalizedLabelsOk() (*[]LocalizedLabelDto, bool)`

GetLocalizedLabelsOk returns a tuple with the LocalizedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedLabels

`func (o *CreateCountryCommand) SetLocalizedLabels(v []LocalizedLabelDto)`

SetLocalizedLabels sets LocalizedLabels field to given value.

### HasLocalizedLabels

`func (o *CreateCountryCommand) HasLocalizedLabels() bool`

HasLocalizedLabels returns a boolean if a field has been set.

### GetIso3Code

`func (o *CreateCountryCommand) GetIso3Code() string`

GetIso3Code returns the Iso3Code field if non-nil, zero value otherwise.

### GetIso3CodeOk

`func (o *CreateCountryCommand) GetIso3CodeOk() (*string, bool)`

GetIso3CodeOk returns a tuple with the Iso3Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3Code

`func (o *CreateCountryCommand) SetIso3Code(v string)`

SetIso3Code sets Iso3Code field to given value.

### HasIso3Code

`func (o *CreateCountryCommand) HasIso3Code() bool`

HasIso3Code returns a boolean if a field has been set.

### GetDefaultVatCode

`func (o *CreateCountryCommand) GetDefaultVatCode() string`

GetDefaultVatCode returns the DefaultVatCode field if non-nil, zero value otherwise.

### GetDefaultVatCodeOk

`func (o *CreateCountryCommand) GetDefaultVatCodeOk() (*string, bool)`

GetDefaultVatCodeOk returns a tuple with the DefaultVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVatCode

`func (o *CreateCountryCommand) SetDefaultVatCode(v string)`

SetDefaultVatCode sets DefaultVatCode field to given value.

### HasDefaultVatCode

`func (o *CreateCountryCommand) HasDefaultVatCode() bool`

HasDefaultVatCode returns a boolean if a field has been set.

### GetCountryFamilyCode

`func (o *CreateCountryCommand) GetCountryFamilyCode() string`

GetCountryFamilyCode returns the CountryFamilyCode field if non-nil, zero value otherwise.

### GetCountryFamilyCodeOk

`func (o *CreateCountryCommand) GetCountryFamilyCodeOk() (*string, bool)`

GetCountryFamilyCodeOk returns a tuple with the CountryFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryFamilyCode

`func (o *CreateCountryCommand) SetCountryFamilyCode(v string)`

SetCountryFamilyCode sets CountryFamilyCode field to given value.

### HasCountryFamilyCode

`func (o *CreateCountryCommand) HasCountryFamilyCode() bool`

HasCountryFamilyCode returns a boolean if a field has been set.

### GetLanguageCode

`func (o *CreateCountryCommand) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *CreateCountryCommand) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *CreateCountryCommand) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *CreateCountryCommand) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


