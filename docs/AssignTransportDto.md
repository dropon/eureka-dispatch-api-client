# AssignTransportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverId** | Pointer to **int32** | Identifier of the driver to assign to the transport.    When the driver is unset, the transport is un-assigned.    When the driver is updated, the transport is assigned to this driver. | [optional] 
**VehicleUid** | Pointer to **string** | Vehicle&#39;s unique identifier | [optional] 
**TrailerUid** | Pointer to **string** | Trailer&#39;s unique identifier | [optional] 
**HandlerId** | Pointer to **int32** | Identifier of the handler | [optional] 
**ContractorAgentId** | Pointer to **int32** | Identifier of the contractor agent.    It can be provided only when the driver is a subcontractor or a subcontractor employee. | [optional] 
**ContractorAgentName** | Pointer to **string** | Name of the contractor agent.    To be provided when the contractor does not exist in the system (free text).    It can be provided only when the driver is a subcontractor or a subcontractor employee. | [optional] 
**ApiBehavior** | Pointer to [**AssignTransportDtoApiBehaviorDto**](AssignTransportDtoApiBehaviorDto.md) |  | [optional] 

## Methods

### NewAssignTransportDto

`func NewAssignTransportDto() *AssignTransportDto`

NewAssignTransportDto instantiates a new AssignTransportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignTransportDtoWithDefaults

`func NewAssignTransportDtoWithDefaults() *AssignTransportDto`

NewAssignTransportDtoWithDefaults instantiates a new AssignTransportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverId

`func (o *AssignTransportDto) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *AssignTransportDto) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *AssignTransportDto) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *AssignTransportDto) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetVehicleUid

`func (o *AssignTransportDto) GetVehicleUid() string`

GetVehicleUid returns the VehicleUid field if non-nil, zero value otherwise.

### GetVehicleUidOk

`func (o *AssignTransportDto) GetVehicleUidOk() (*string, bool)`

GetVehicleUidOk returns a tuple with the VehicleUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleUid

`func (o *AssignTransportDto) SetVehicleUid(v string)`

SetVehicleUid sets VehicleUid field to given value.

### HasVehicleUid

`func (o *AssignTransportDto) HasVehicleUid() bool`

HasVehicleUid returns a boolean if a field has been set.

### GetTrailerUid

`func (o *AssignTransportDto) GetTrailerUid() string`

GetTrailerUid returns the TrailerUid field if non-nil, zero value otherwise.

### GetTrailerUidOk

`func (o *AssignTransportDto) GetTrailerUidOk() (*string, bool)`

GetTrailerUidOk returns a tuple with the TrailerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerUid

`func (o *AssignTransportDto) SetTrailerUid(v string)`

SetTrailerUid sets TrailerUid field to given value.

### HasTrailerUid

`func (o *AssignTransportDto) HasTrailerUid() bool`

HasTrailerUid returns a boolean if a field has been set.

### GetHandlerId

`func (o *AssignTransportDto) GetHandlerId() int32`

GetHandlerId returns the HandlerId field if non-nil, zero value otherwise.

### GetHandlerIdOk

`func (o *AssignTransportDto) GetHandlerIdOk() (*int32, bool)`

GetHandlerIdOk returns a tuple with the HandlerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerId

`func (o *AssignTransportDto) SetHandlerId(v int32)`

SetHandlerId sets HandlerId field to given value.

### HasHandlerId

`func (o *AssignTransportDto) HasHandlerId() bool`

HasHandlerId returns a boolean if a field has been set.

### GetContractorAgentId

`func (o *AssignTransportDto) GetContractorAgentId() int32`

GetContractorAgentId returns the ContractorAgentId field if non-nil, zero value otherwise.

### GetContractorAgentIdOk

`func (o *AssignTransportDto) GetContractorAgentIdOk() (*int32, bool)`

GetContractorAgentIdOk returns a tuple with the ContractorAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorAgentId

`func (o *AssignTransportDto) SetContractorAgentId(v int32)`

SetContractorAgentId sets ContractorAgentId field to given value.

### HasContractorAgentId

`func (o *AssignTransportDto) HasContractorAgentId() bool`

HasContractorAgentId returns a boolean if a field has been set.

### GetContractorAgentName

`func (o *AssignTransportDto) GetContractorAgentName() string`

GetContractorAgentName returns the ContractorAgentName field if non-nil, zero value otherwise.

### GetContractorAgentNameOk

`func (o *AssignTransportDto) GetContractorAgentNameOk() (*string, bool)`

GetContractorAgentNameOk returns a tuple with the ContractorAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorAgentName

`func (o *AssignTransportDto) SetContractorAgentName(v string)`

SetContractorAgentName sets ContractorAgentName field to given value.

### HasContractorAgentName

`func (o *AssignTransportDto) HasContractorAgentName() bool`

HasContractorAgentName returns a boolean if a field has been set.

### GetApiBehavior

`func (o *AssignTransportDto) GetApiBehavior() AssignTransportDtoApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *AssignTransportDto) GetApiBehaviorOk() (*AssignTransportDtoApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *AssignTransportDto) SetApiBehavior(v AssignTransportDtoApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *AssignTransportDto) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


