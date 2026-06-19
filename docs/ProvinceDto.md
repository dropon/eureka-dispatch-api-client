# ProvinceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryDto**](CountryDto.md) |  | [optional] 

## Methods

### NewProvinceDto

`func NewProvinceDto() *ProvinceDto`

NewProvinceDto instantiates a new ProvinceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvinceDtoWithDefaults

`func NewProvinceDtoWithDefaults() *ProvinceDto`

NewProvinceDtoWithDefaults instantiates a new ProvinceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ProvinceDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProvinceDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProvinceDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProvinceDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *ProvinceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProvinceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProvinceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProvinceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *ProvinceDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProvinceDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProvinceDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProvinceDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCountry

`func (o *ProvinceDto) GetCountry() CountryDto`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProvinceDto) GetCountryOk() (*CountryDto, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProvinceDto) SetCountry(v CountryDto)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProvinceDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


