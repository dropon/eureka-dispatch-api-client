# DriverIncludedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVehicle** | Pointer to [**VehicleDto**](VehicleDto.md) |  | [optional] 
**DefaultVehicle** | Pointer to [**VehicleDto**](VehicleDto.md) |  | [optional] 
**Address** | Pointer to [**AddressDto**](AddressDto.md) |  | [optional] 

## Methods

### NewDriverIncludedDto

`func NewDriverIncludedDto() *DriverIncludedDto`

NewDriverIncludedDto instantiates a new DriverIncludedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverIncludedDtoWithDefaults

`func NewDriverIncludedDtoWithDefaults() *DriverIncludedDto`

NewDriverIncludedDtoWithDefaults instantiates a new DriverIncludedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVehicle

`func (o *DriverIncludedDto) GetCurrentVehicle() VehicleDto`

GetCurrentVehicle returns the CurrentVehicle field if non-nil, zero value otherwise.

### GetCurrentVehicleOk

`func (o *DriverIncludedDto) GetCurrentVehicleOk() (*VehicleDto, bool)`

GetCurrentVehicleOk returns a tuple with the CurrentVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVehicle

`func (o *DriverIncludedDto) SetCurrentVehicle(v VehicleDto)`

SetCurrentVehicle sets CurrentVehicle field to given value.

### HasCurrentVehicle

`func (o *DriverIncludedDto) HasCurrentVehicle() bool`

HasCurrentVehicle returns a boolean if a field has been set.

### GetDefaultVehicle

`func (o *DriverIncludedDto) GetDefaultVehicle() VehicleDto`

GetDefaultVehicle returns the DefaultVehicle field if non-nil, zero value otherwise.

### GetDefaultVehicleOk

`func (o *DriverIncludedDto) GetDefaultVehicleOk() (*VehicleDto, bool)`

GetDefaultVehicleOk returns a tuple with the DefaultVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVehicle

`func (o *DriverIncludedDto) SetDefaultVehicle(v VehicleDto)`

SetDefaultVehicle sets DefaultVehicle field to given value.

### HasDefaultVehicle

`func (o *DriverIncludedDto) HasDefaultVehicle() bool`

HasDefaultVehicle returns a boolean if a field has been set.

### GetAddress

`func (o *DriverIncludedDto) GetAddress() AddressDto`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DriverIncludedDto) GetAddressOk() (*AddressDto, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DriverIncludedDto) SetAddress(v AddressDto)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DriverIncludedDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


