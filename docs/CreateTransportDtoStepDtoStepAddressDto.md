# CreateTransportDtoStepDtoStepAddressDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **int32** | When an address id is provided, step address information is filled with address data,  but address line 1 and 2 can be overridden. | [optional] 
**CityId** | Pointer to **int32** | When a city id is provided, the country, post code and city name are pre-filled with city data. | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PostCode** | Pointer to **string** |  | [optional] 
**CityName** | Pointer to **string** |  | [optional] 
**SectorCode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StreetNumber** | Pointer to **string** |  | [optional] 
**StreetName** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Longitude** | Pointer to **float64** | This number must be between -180 and 180. | [optional] 
**Latitude** | Pointer to **float64** | This number must be between -90 and 90. | [optional] 
**AddressComplement1** | Pointer to **string** |  | [optional] 
**AddressComplement2** | Pointer to **string** |  | [optional] 
**AddressComplement3** | Pointer to **string** |  | [optional] 
**AddressComplement4** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTransportDtoStepDtoStepAddressDto

`func NewCreateTransportDtoStepDtoStepAddressDto() *CreateTransportDtoStepDtoStepAddressDto`

NewCreateTransportDtoStepDtoStepAddressDto instantiates a new CreateTransportDtoStepDtoStepAddressDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoStepDtoStepAddressDtoWithDefaults

`func NewCreateTransportDtoStepDtoStepAddressDtoWithDefaults() *CreateTransportDtoStepDtoStepAddressDto`

NewCreateTransportDtoStepDtoStepAddressDtoWithDefaults instantiates a new CreateTransportDtoStepDtoStepAddressDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetCityId

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPostCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetPostCode() string`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetPostCodeOk() (*string, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetPostCode(v string)`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### GetCityName

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetCityName(v string)`

SetCityName sets CityName field to given value.

### HasCityName

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasCityName() bool`

HasCityName returns a boolean if a field has been set.

### GetSectorCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetSectorCode() string`

GetSectorCode returns the SectorCode field if non-nil, zero value otherwise.

### GetSectorCodeOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetSectorCodeOk() (*string, bool)`

GetSectorCodeOk returns a tuple with the SectorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetSectorCode(v string)`

SetSectorCode sets SectorCode field to given value.

### HasSectorCode

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasSectorCode() bool`

HasSectorCode returns a boolean if a field has been set.

### GetName

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetNumber

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetLongitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetAddressComplement1

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement1() string`

GetAddressComplement1 returns the AddressComplement1 field if non-nil, zero value otherwise.

### GetAddressComplement1Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement1Ok() (*string, bool)`

GetAddressComplement1Ok returns a tuple with the AddressComplement1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement1

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressComplement1(v string)`

SetAddressComplement1 sets AddressComplement1 field to given value.

### HasAddressComplement1

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressComplement1() bool`

HasAddressComplement1 returns a boolean if a field has been set.

### GetAddressComplement2

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement2() string`

GetAddressComplement2 returns the AddressComplement2 field if non-nil, zero value otherwise.

### GetAddressComplement2Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement2Ok() (*string, bool)`

GetAddressComplement2Ok returns a tuple with the AddressComplement2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement2

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressComplement2(v string)`

SetAddressComplement2 sets AddressComplement2 field to given value.

### HasAddressComplement2

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressComplement2() bool`

HasAddressComplement2 returns a boolean if a field has been set.

### GetAddressComplement3

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement3() string`

GetAddressComplement3 returns the AddressComplement3 field if non-nil, zero value otherwise.

### GetAddressComplement3Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement3Ok() (*string, bool)`

GetAddressComplement3Ok returns a tuple with the AddressComplement3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement3

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressComplement3(v string)`

SetAddressComplement3 sets AddressComplement3 field to given value.

### HasAddressComplement3

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressComplement3() bool`

HasAddressComplement3 returns a boolean if a field has been set.

### GetAddressComplement4

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement4() string`

GetAddressComplement4 returns the AddressComplement4 field if non-nil, zero value otherwise.

### GetAddressComplement4Ok

`func (o *CreateTransportDtoStepDtoStepAddressDto) GetAddressComplement4Ok() (*string, bool)`

GetAddressComplement4Ok returns a tuple with the AddressComplement4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement4

`func (o *CreateTransportDtoStepDtoStepAddressDto) SetAddressComplement4(v string)`

SetAddressComplement4 sets AddressComplement4 field to given value.

### HasAddressComplement4

`func (o *CreateTransportDtoStepDtoStepAddressDto) HasAddressComplement4() bool`

HasAddressComplement4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


