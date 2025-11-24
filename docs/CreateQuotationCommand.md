# CreateQuotationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | [**[]CreateTransportDto**](CreateTransportDto.md) | Mandatory. Transports included in the mission | 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**CreateQuotationCommandApiBehaviorDto**](CreateQuotationCommandApiBehaviorDto.md) |  | [optional] 

## Methods

### NewCreateQuotationCommand

`func NewCreateQuotationCommand(transports []CreateTransportDto, ) *CreateQuotationCommand`

NewCreateQuotationCommand instantiates a new CreateQuotationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuotationCommandWithDefaults

`func NewCreateQuotationCommandWithDefaults() *CreateQuotationCommand`

NewCreateQuotationCommandWithDefaults instantiates a new CreateQuotationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *CreateQuotationCommand) GetTransports() []CreateTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateQuotationCommand) GetTransportsOk() (*[]CreateTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateQuotationCommand) SetTransports(v []CreateTransportDto)`

SetTransports sets Transports field to given value.


### GetIsFinalized

`func (o *CreateQuotationCommand) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *CreateQuotationCommand) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *CreateQuotationCommand) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *CreateQuotationCommand) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetApiBehavior

`func (o *CreateQuotationCommand) GetApiBehavior() CreateQuotationCommandApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *CreateQuotationCommand) GetApiBehaviorOk() (*CreateQuotationCommandApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *CreateQuotationCommand) SetApiBehavior(v CreateQuotationCommandApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *CreateQuotationCommand) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


