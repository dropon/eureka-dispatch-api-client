# SimulateMissionCreationDryRunCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | [**[]CreateTransportDto**](CreateTransportDto.md) | Mandatory. Transports included in the mission | 
**IsReadyToBill** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**SimulateMissionCreationDryRunCommandApiBehaviorDto**](SimulateMissionCreationDryRunCommandApiBehaviorDto.md) |  | [optional] 

## Methods

### NewSimulateMissionCreationDryRunCommand

`func NewSimulateMissionCreationDryRunCommand(transports []CreateTransportDto, ) *SimulateMissionCreationDryRunCommand`

NewSimulateMissionCreationDryRunCommand instantiates a new SimulateMissionCreationDryRunCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateMissionCreationDryRunCommandWithDefaults

`func NewSimulateMissionCreationDryRunCommandWithDefaults() *SimulateMissionCreationDryRunCommand`

NewSimulateMissionCreationDryRunCommandWithDefaults instantiates a new SimulateMissionCreationDryRunCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *SimulateMissionCreationDryRunCommand) GetTransports() []CreateTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *SimulateMissionCreationDryRunCommand) GetTransportsOk() (*[]CreateTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *SimulateMissionCreationDryRunCommand) SetTransports(v []CreateTransportDto)`

SetTransports sets Transports field to given value.


### GetIsReadyToBill

`func (o *SimulateMissionCreationDryRunCommand) GetIsReadyToBill() bool`

GetIsReadyToBill returns the IsReadyToBill field if non-nil, zero value otherwise.

### GetIsReadyToBillOk

`func (o *SimulateMissionCreationDryRunCommand) GetIsReadyToBillOk() (*bool, bool)`

GetIsReadyToBillOk returns a tuple with the IsReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadyToBill

`func (o *SimulateMissionCreationDryRunCommand) SetIsReadyToBill(v bool)`

SetIsReadyToBill sets IsReadyToBill field to given value.

### HasIsReadyToBill

`func (o *SimulateMissionCreationDryRunCommand) HasIsReadyToBill() bool`

HasIsReadyToBill returns a boolean if a field has been set.

### GetApiBehavior

`func (o *SimulateMissionCreationDryRunCommand) GetApiBehavior() SimulateMissionCreationDryRunCommandApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *SimulateMissionCreationDryRunCommand) GetApiBehaviorOk() (*SimulateMissionCreationDryRunCommandApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *SimulateMissionCreationDryRunCommand) SetApiBehavior(v SimulateMissionCreationDryRunCommandApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *SimulateMissionCreationDryRunCommand) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


