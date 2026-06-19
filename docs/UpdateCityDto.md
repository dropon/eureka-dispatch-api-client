# UpdateCityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PostCode** | Pointer to **string** |  | [optional] 
**SectorCode** | Pointer to **string** |  | [optional] 
**OperationZoneCode** | Pointer to **string** |  | [optional] 
**Geocoding** | Pointer to [**UpdateCityDtoUpdateCityGeocodingDto**](UpdateCityDtoUpdateCityGeocodingDto.md) |  | [optional] 
**ProvinceUid** | Pointer to **string** |  | [optional] 
**ProvinceSubdivisionUid** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**IsAvailableForMissionEntry** | Pointer to **bool** | Indicates whether the entity is available for mission entry.    When not set, the city will be available for mission entry. | [optional] 

## Methods

### NewUpdateCityDto

`func NewUpdateCityDto() *UpdateCityDto`

NewUpdateCityDto instantiates a new UpdateCityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCityDtoWithDefaults

`func NewUpdateCityDtoWithDefaults() *UpdateCityDto`

NewUpdateCityDtoWithDefaults instantiates a new UpdateCityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCityDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCityDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCityDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCityDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostCode

`func (o *UpdateCityDto) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *UpdateCityDto) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *UpdateCityDto) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *UpdateCityDto) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### GetSectorCode

`func (o *UpdateCityDto) GetSectorCode() string`

GetSectorCode returns the SectorCode field if non-nil, zero value otherwise.

### GetSectorCodeOk

`func (o *UpdateCityDto) GetSectorCodeOk() (*string, bool)`

GetSectorCodeOk returns a tuple with the SectorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorCode

`func (o *UpdateCityDto) SetSectorCode(v string)`

SetSectorCode sets SectorCode field to given value.

### HasSectorCode

`func (o *UpdateCityDto) HasSectorCode() bool`

HasSectorCode returns a boolean if a field has been set.

### GetOperationZoneCode

`func (o *UpdateCityDto) GetOperationZoneCode() string`

GetOperationZoneCode returns the OperationZoneCode field if non-nil, zero value otherwise.

### GetOperationZoneCodeOk

`func (o *UpdateCityDto) GetOperationZoneCodeOk() (*string, bool)`

GetOperationZoneCodeOk returns a tuple with the OperationZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationZoneCode

`func (o *UpdateCityDto) SetOperationZoneCode(v string)`

SetOperationZoneCode sets OperationZoneCode field to given value.

### HasOperationZoneCode

`func (o *UpdateCityDto) HasOperationZoneCode() bool`

HasOperationZoneCode returns a boolean if a field has been set.

### GetGeocoding

`func (o *UpdateCityDto) GetGeocoding() UpdateCityDtoUpdateCityGeocodingDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *UpdateCityDto) GetGeocodingOk() (*UpdateCityDtoUpdateCityGeocodingDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *UpdateCityDto) SetGeocoding(v UpdateCityDtoUpdateCityGeocodingDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *UpdateCityDto) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetProvinceUid

`func (o *UpdateCityDto) GetProvinceUid() string`

GetProvinceUid returns the ProvinceUid field if non-nil, zero value otherwise.

### GetProvinceUidOk

`func (o *UpdateCityDto) GetProvinceUidOk() (*string, bool)`

GetProvinceUidOk returns a tuple with the ProvinceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceUid

`func (o *UpdateCityDto) SetProvinceUid(v string)`

SetProvinceUid sets ProvinceUid field to given value.

### HasProvinceUid

`func (o *UpdateCityDto) HasProvinceUid() bool`

HasProvinceUid returns a boolean if a field has been set.

### GetProvinceSubdivisionUid

`func (o *UpdateCityDto) GetProvinceSubdivisionUid() string`

GetProvinceSubdivisionUid returns the ProvinceSubdivisionUid field if non-nil, zero value otherwise.

### GetProvinceSubdivisionUidOk

`func (o *UpdateCityDto) GetProvinceSubdivisionUidOk() (*string, bool)`

GetProvinceSubdivisionUidOk returns a tuple with the ProvinceSubdivisionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceSubdivisionUid

`func (o *UpdateCityDto) SetProvinceSubdivisionUid(v string)`

SetProvinceSubdivisionUid sets ProvinceSubdivisionUid field to given value.

### HasProvinceSubdivisionUid

`func (o *UpdateCityDto) HasProvinceSubdivisionUid() bool`

HasProvinceSubdivisionUid returns a boolean if a field has been set.

### GetCountryCode

`func (o *UpdateCityDto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UpdateCityDto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UpdateCityDto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UpdateCityDto) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *UpdateCityDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *UpdateCityDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *UpdateCityDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *UpdateCityDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetIsAvailableForMissionEntry

`func (o *UpdateCityDto) GetIsAvailableForMissionEntry() bool`

GetIsAvailableForMissionEntry returns the IsAvailableForMissionEntry field if non-nil, zero value otherwise.

### GetIsAvailableForMissionEntryOk

`func (o *UpdateCityDto) GetIsAvailableForMissionEntryOk() (*bool, bool)`

GetIsAvailableForMissionEntryOk returns a tuple with the IsAvailableForMissionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForMissionEntry

`func (o *UpdateCityDto) SetIsAvailableForMissionEntry(v bool)`

SetIsAvailableForMissionEntry sets IsAvailableForMissionEntry field to given value.

### HasIsAvailableForMissionEntry

`func (o *UpdateCityDto) HasIsAvailableForMissionEntry() bool`

HasIsAvailableForMissionEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


