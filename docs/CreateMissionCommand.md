# CreateMissionCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | [**[]CreateTransportDto**](CreateTransportDto.md) | Mandatory. Transports included in the mission | 
**IsReadyToBill** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**CreateMissionCommandApiBehaviorDto**](CreateMissionCommandApiBehaviorDto.md) |  | [optional] 

## Methods

### NewCreateMissionCommand

`func NewCreateMissionCommand(transports []CreateTransportDto, ) *CreateMissionCommand`

NewCreateMissionCommand instantiates a new CreateMissionCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMissionCommandWithDefaults

`func NewCreateMissionCommandWithDefaults() *CreateMissionCommand`

NewCreateMissionCommandWithDefaults instantiates a new CreateMissionCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *CreateMissionCommand) GetTransports() []CreateTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateMissionCommand) GetTransportsOk() (*[]CreateTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateMissionCommand) SetTransports(v []CreateTransportDto)`

SetTransports sets Transports field to given value.


### GetIsReadyToBill

`func (o *CreateMissionCommand) GetIsReadyToBill() bool`

GetIsReadyToBill returns the IsReadyToBill field if non-nil, zero value otherwise.

### GetIsReadyToBillOk

`func (o *CreateMissionCommand) GetIsReadyToBillOk() (*bool, bool)`

GetIsReadyToBillOk returns a tuple with the IsReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadyToBill

`func (o *CreateMissionCommand) SetIsReadyToBill(v bool)`

SetIsReadyToBill sets IsReadyToBill field to given value.

### HasIsReadyToBill

`func (o *CreateMissionCommand) HasIsReadyToBill() bool`

HasIsReadyToBill returns a boolean if a field has been set.

### GetApiBehavior

`func (o *CreateMissionCommand) GetApiBehavior() CreateMissionCommandApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *CreateMissionCommand) GetApiBehaviorOk() (*CreateMissionCommandApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *CreateMissionCommand) SetApiBehavior(v CreateMissionCommandApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *CreateMissionCommand) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


