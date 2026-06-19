# CreateCityCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Mandatory. | 
**PostCode** | Pointer to **string** |  | [optional] 
**SectorCode** | Pointer to **string** |  | [optional] 
**OperationZoneCode** | Pointer to **string** |  | [optional] 
**Geocoding** | Pointer to [**CreateCityCommandCreateCityGeocodingDto**](CreateCityCommandCreateCityGeocodingDto.md) |  | [optional] 
**ProvinceUid** | Pointer to **string** | The province is mandatory when the specified country has provinces, and must be null otherwise.    It is also mandatory if the subdivision is specified. | [optional] 
**ProvinceSubdivisionUid** | Pointer to **string** |  | [optional] 
**CountryCode** | **string** | Mandatory. | 
**AgencyCode** | Pointer to **string** |  | [optional] 
**IsAvailableForMissionEntry** | Pointer to **bool** | Indicates whether the entity is available for mission entry.    When not set, the city will be available for mission entry. | [optional] 

## Methods

### NewCreateCityCommand

`func NewCreateCityCommand(name string, countryCode string, ) *CreateCityCommand`

NewCreateCityCommand instantiates a new CreateCityCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCityCommandWithDefaults

`func NewCreateCityCommandWithDefaults() *CreateCityCommand`

NewCreateCityCommandWithDefaults instantiates a new CreateCityCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCityCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCityCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCityCommand) SetName(v string)`

SetName sets Name field to given value.


### GetPostCode

`func (o *CreateCityCommand) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *CreateCityCommand) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *CreateCityCommand) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *CreateCityCommand) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### GetSectorCode

`func (o *CreateCityCommand) GetSectorCode() string`

GetSectorCode returns the SectorCode field if non-nil, zero value otherwise.

### GetSectorCodeOk

`func (o *CreateCityCommand) GetSectorCodeOk() (*string, bool)`

GetSectorCodeOk returns a tuple with the SectorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorCode

`func (o *CreateCityCommand) SetSectorCode(v string)`

SetSectorCode sets SectorCode field to given value.

### HasSectorCode

`func (o *CreateCityCommand) HasSectorCode() bool`

HasSectorCode returns a boolean if a field has been set.

### GetOperationZoneCode

`func (o *CreateCityCommand) GetOperationZoneCode() string`

GetOperationZoneCode returns the OperationZoneCode field if non-nil, zero value otherwise.

### GetOperationZoneCodeOk

`func (o *CreateCityCommand) GetOperationZoneCodeOk() (*string, bool)`

GetOperationZoneCodeOk returns a tuple with the OperationZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationZoneCode

`func (o *CreateCityCommand) SetOperationZoneCode(v string)`

SetOperationZoneCode sets OperationZoneCode field to given value.

### HasOperationZoneCode

`func (o *CreateCityCommand) HasOperationZoneCode() bool`

HasOperationZoneCode returns a boolean if a field has been set.

### GetGeocoding

`func (o *CreateCityCommand) GetGeocoding() CreateCityCommandCreateCityGeocodingDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *CreateCityCommand) GetGeocodingOk() (*CreateCityCommandCreateCityGeocodingDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *CreateCityCommand) SetGeocoding(v CreateCityCommandCreateCityGeocodingDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *CreateCityCommand) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetProvinceUid

`func (o *CreateCityCommand) GetProvinceUid() string`

GetProvinceUid returns the ProvinceUid field if non-nil, zero value otherwise.

### GetProvinceUidOk

`func (o *CreateCityCommand) GetProvinceUidOk() (*string, bool)`

GetProvinceUidOk returns a tuple with the ProvinceUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceUid

`func (o *CreateCityCommand) SetProvinceUid(v string)`

SetProvinceUid sets ProvinceUid field to given value.

### HasProvinceUid

`func (o *CreateCityCommand) HasProvinceUid() bool`

HasProvinceUid returns a boolean if a field has been set.

### GetProvinceSubdivisionUid

`func (o *CreateCityCommand) GetProvinceSubdivisionUid() string`

GetProvinceSubdivisionUid returns the ProvinceSubdivisionUid field if non-nil, zero value otherwise.

### GetProvinceSubdivisionUidOk

`func (o *CreateCityCommand) GetProvinceSubdivisionUidOk() (*string, bool)`

GetProvinceSubdivisionUidOk returns a tuple with the ProvinceSubdivisionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvinceSubdivisionUid

`func (o *CreateCityCommand) SetProvinceSubdivisionUid(v string)`

SetProvinceSubdivisionUid sets ProvinceSubdivisionUid field to given value.

### HasProvinceSubdivisionUid

`func (o *CreateCityCommand) HasProvinceSubdivisionUid() bool`

HasProvinceSubdivisionUid returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreateCityCommand) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateCityCommand) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateCityCommand) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetAgencyCode

`func (o *CreateCityCommand) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CreateCityCommand) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CreateCityCommand) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *CreateCityCommand) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetIsAvailableForMissionEntry

`func (o *CreateCityCommand) GetIsAvailableForMissionEntry() bool`

GetIsAvailableForMissionEntry returns the IsAvailableForMissionEntry field if non-nil, zero value otherwise.

### GetIsAvailableForMissionEntryOk

`func (o *CreateCityCommand) GetIsAvailableForMissionEntryOk() (*bool, bool)`

GetIsAvailableForMissionEntryOk returns a tuple with the IsAvailableForMissionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForMissionEntry

`func (o *CreateCityCommand) SetIsAvailableForMissionEntry(v bool)`

SetIsAvailableForMissionEntry sets IsAvailableForMissionEntry field to given value.

### HasIsAvailableForMissionEntry

`func (o *CreateCityCommand) HasIsAvailableForMissionEntry() bool`

HasIsAvailableForMissionEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


