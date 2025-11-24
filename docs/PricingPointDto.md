# PricingPointDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **int32** |  | [optional] 
**CityId** | Pointer to **int32** |  | [optional] 
**SectorCode** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**AddressDto**](AddressDto.md) |  | [optional] 
**City** | Pointer to [**CityDto**](CityDto.md) |  | [optional] 

## Methods

### NewPricingPointDto

`func NewPricingPointDto() *PricingPointDto`

NewPricingPointDto instantiates a new PricingPointDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPointDtoWithDefaults

`func NewPricingPointDtoWithDefaults() *PricingPointDto`

NewPricingPointDtoWithDefaults instantiates a new PricingPointDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *PricingPointDto) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *PricingPointDto) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *PricingPointDto) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *PricingPointDto) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetCityId

`func (o *PricingPointDto) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *PricingPointDto) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *PricingPointDto) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *PricingPointDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetSectorCode

`func (o *PricingPointDto) GetSectorCode() string`

GetSectorCode returns the SectorCode field if non-nil, zero value otherwise.

### GetSectorCodeOk

`func (o *PricingPointDto) GetSectorCodeOk() (*string, bool)`

GetSectorCodeOk returns a tuple with the SectorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorCode

`func (o *PricingPointDto) SetSectorCode(v string)`

SetSectorCode sets SectorCode field to given value.

### HasSectorCode

`func (o *PricingPointDto) HasSectorCode() bool`

HasSectorCode returns a boolean if a field has been set.

### GetAddress

`func (o *PricingPointDto) GetAddress() AddressDto`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PricingPointDto) GetAddressOk() (*AddressDto, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PricingPointDto) SetAddress(v AddressDto)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PricingPointDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *PricingPointDto) GetCity() CityDto`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PricingPointDto) GetCityOk() (*CityDto, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PricingPointDto) SetCity(v CityDto)`

SetCity sets City field to given value.

### HasCity

`func (o *PricingPointDto) HasCity() bool`

HasCity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


