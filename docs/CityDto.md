# CityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PostCode** | Pointer to **string** |  | [optional] 
**SectorCode** | Pointer to **string** |  | [optional] 
**Geocoding** | Pointer to [**GeocodingDto**](GeocodingDto.md) |  | [optional] 
**Country** | Pointer to [**CountryDto**](CountryDto.md) |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**IsAvailableForMissionEntry** | Pointer to **bool** |  | [optional] 

## Methods

### NewCityDto

`func NewCityDto() *CityDto`

NewCityDto instantiates a new CityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCityDtoWithDefaults

`func NewCityDtoWithDefaults() *CityDto`

NewCityDtoWithDefaults instantiates a new CityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityId

`func (o *CityDto) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *CityDto) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *CityDto) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *CityDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetName

`func (o *CityDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CityDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CityDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CityDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostCode

`func (o *CityDto) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *CityDto) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *CityDto) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *CityDto) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### GetSectorCode

`func (o *CityDto) GetSectorCode() string`

GetSectorCode returns the SectorCode field if non-nil, zero value otherwise.

### GetSectorCodeOk

`func (o *CityDto) GetSectorCodeOk() (*string, bool)`

GetSectorCodeOk returns a tuple with the SectorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorCode

`func (o *CityDto) SetSectorCode(v string)`

SetSectorCode sets SectorCode field to given value.

### HasSectorCode

`func (o *CityDto) HasSectorCode() bool`

HasSectorCode returns a boolean if a field has been set.

### GetGeocoding

`func (o *CityDto) GetGeocoding() GeocodingDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *CityDto) GetGeocodingOk() (*GeocodingDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *CityDto) SetGeocoding(v GeocodingDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *CityDto) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetCountry

`func (o *CityDto) GetCountry() CountryDto`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CityDto) GetCountryOk() (*CountryDto, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CityDto) SetCountry(v CountryDto)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CityDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAgencyCode

`func (o *CityDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CityDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CityDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *CityDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetIsAvailableForMissionEntry

`func (o *CityDto) GetIsAvailableForMissionEntry() bool`

GetIsAvailableForMissionEntry returns the IsAvailableForMissionEntry field if non-nil, zero value otherwise.

### GetIsAvailableForMissionEntryOk

`func (o *CityDto) GetIsAvailableForMissionEntryOk() (*bool, bool)`

GetIsAvailableForMissionEntryOk returns a tuple with the IsAvailableForMissionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForMissionEntry

`func (o *CityDto) SetIsAvailableForMissionEntry(v bool)`

SetIsAvailableForMissionEntry sets IsAvailableForMissionEntry field to given value.

### HasIsAvailableForMissionEntry

`func (o *CityDto) HasIsAvailableForMissionEntry() bool`

HasIsAvailableForMissionEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


