# SimulateQuotationCreationDryRunCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | [**[]CreateTransportDto**](CreateTransportDto.md) | Mandatory. Transports included in the mission | 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**SimulateQuotationCreationDryRunCommandApiBehaviorDto**](SimulateQuotationCreationDryRunCommandApiBehaviorDto.md) |  | [optional] 

## Methods

### NewSimulateQuotationCreationDryRunCommand

`func NewSimulateQuotationCreationDryRunCommand(transports []CreateTransportDto, ) *SimulateQuotationCreationDryRunCommand`

NewSimulateQuotationCreationDryRunCommand instantiates a new SimulateQuotationCreationDryRunCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateQuotationCreationDryRunCommandWithDefaults

`func NewSimulateQuotationCreationDryRunCommandWithDefaults() *SimulateQuotationCreationDryRunCommand`

NewSimulateQuotationCreationDryRunCommandWithDefaults instantiates a new SimulateQuotationCreationDryRunCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *SimulateQuotationCreationDryRunCommand) GetTransports() []CreateTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *SimulateQuotationCreationDryRunCommand) GetTransportsOk() (*[]CreateTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *SimulateQuotationCreationDryRunCommand) SetTransports(v []CreateTransportDto)`

SetTransports sets Transports field to given value.


### GetIsFinalized

`func (o *SimulateQuotationCreationDryRunCommand) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *SimulateQuotationCreationDryRunCommand) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *SimulateQuotationCreationDryRunCommand) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *SimulateQuotationCreationDryRunCommand) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetApiBehavior

`func (o *SimulateQuotationCreationDryRunCommand) GetApiBehavior() SimulateQuotationCreationDryRunCommandApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *SimulateQuotationCreationDryRunCommand) GetApiBehaviorOk() (*SimulateQuotationCreationDryRunCommandApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *SimulateQuotationCreationDryRunCommand) SetApiBehavior(v SimulateQuotationCreationDryRunCommandApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *SimulateQuotationCreationDryRunCommand) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


