# TransportIncludedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**ServiceDto**](ServiceDto.md) |  | [optional] 
**SubState** | Pointer to [**SubStateDto**](SubStateDto.md) |  | [optional] 
**Customer** | Pointer to [**CustomerDto**](CustomerDto.md) |  | [optional] 
**Agency** | Pointer to [**AgencyDto**](AgencyDto.md) |  | [optional] 
**SearchFlags** | Pointer to [**SearchFlagsDto**](SearchFlagsDto.md) |  | [optional] 
**Handler** | Pointer to [**DriverDto**](DriverDto.md) |  | [optional] 
**Driver** | Pointer to [**DriverDto**](DriverDto.md) |  | [optional] 
**Vehicle** | Pointer to [**VehicleDto**](VehicleDto.md) |  | [optional] 
**Trailer** | Pointer to [**VehicleDto**](VehicleDto.md) |  | [optional] 

## Methods

### NewTransportIncludedDto

`func NewTransportIncludedDto() *TransportIncludedDto`

NewTransportIncludedDto instantiates a new TransportIncludedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportIncludedDtoWithDefaults

`func NewTransportIncludedDtoWithDefaults() *TransportIncludedDto`

NewTransportIncludedDtoWithDefaults instantiates a new TransportIncludedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *TransportIncludedDto) GetService() ServiceDto`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TransportIncludedDto) GetServiceOk() (*ServiceDto, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TransportIncludedDto) SetService(v ServiceDto)`

SetService sets Service field to given value.

### HasService

`func (o *TransportIncludedDto) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSubState

`func (o *TransportIncludedDto) GetSubState() SubStateDto`

GetSubState returns the SubState field if non-nil, zero value otherwise.

### GetSubStateOk

`func (o *TransportIncludedDto) GetSubStateOk() (*SubStateDto, bool)`

GetSubStateOk returns a tuple with the SubState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubState

`func (o *TransportIncludedDto) SetSubState(v SubStateDto)`

SetSubState sets SubState field to given value.

### HasSubState

`func (o *TransportIncludedDto) HasSubState() bool`

HasSubState returns a boolean if a field has been set.

### GetCustomer

`func (o *TransportIncludedDto) GetCustomer() CustomerDto`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *TransportIncludedDto) GetCustomerOk() (*CustomerDto, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *TransportIncludedDto) SetCustomer(v CustomerDto)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *TransportIncludedDto) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAgency

`func (o *TransportIncludedDto) GetAgency() AgencyDto`

GetAgency returns the Agency field if non-nil, zero value otherwise.

### GetAgencyOk

`func (o *TransportIncludedDto) GetAgencyOk() (*AgencyDto, bool)`

GetAgencyOk returns a tuple with the Agency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgency

`func (o *TransportIncludedDto) SetAgency(v AgencyDto)`

SetAgency sets Agency field to given value.

### HasAgency

`func (o *TransportIncludedDto) HasAgency() bool`

HasAgency returns a boolean if a field has been set.

### GetSearchFlags

`func (o *TransportIncludedDto) GetSearchFlags() SearchFlagsDto`

GetSearchFlags returns the SearchFlags field if non-nil, zero value otherwise.

### GetSearchFlagsOk

`func (o *TransportIncludedDto) GetSearchFlagsOk() (*SearchFlagsDto, bool)`

GetSearchFlagsOk returns a tuple with the SearchFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFlags

`func (o *TransportIncludedDto) SetSearchFlags(v SearchFlagsDto)`

SetSearchFlags sets SearchFlags field to given value.

### HasSearchFlags

`func (o *TransportIncludedDto) HasSearchFlags() bool`

HasSearchFlags returns a boolean if a field has been set.

### GetHandler

`func (o *TransportIncludedDto) GetHandler() DriverDto`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *TransportIncludedDto) GetHandlerOk() (*DriverDto, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *TransportIncludedDto) SetHandler(v DriverDto)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *TransportIncludedDto) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### GetDriver

`func (o *TransportIncludedDto) GetDriver() DriverDto`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *TransportIncludedDto) GetDriverOk() (*DriverDto, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *TransportIncludedDto) SetDriver(v DriverDto)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *TransportIncludedDto) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetVehicle

`func (o *TransportIncludedDto) GetVehicle() VehicleDto`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *TransportIncludedDto) GetVehicleOk() (*VehicleDto, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *TransportIncludedDto) SetVehicle(v VehicleDto)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *TransportIncludedDto) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetTrailer

`func (o *TransportIncludedDto) GetTrailer() VehicleDto`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *TransportIncludedDto) GetTrailerOk() (*VehicleDto, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *TransportIncludedDto) SetTrailer(v VehicleDto)`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *TransportIncludedDto) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


