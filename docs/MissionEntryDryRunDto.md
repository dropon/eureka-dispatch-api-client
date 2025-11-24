# MissionEntryDryRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReadyToBill** | Pointer to **bool** |  | [optional] 
**IsUniqueOrder** | Pointer to **bool** |  | [optional] 
**Transports** | Pointer to [**[]TransportEntryDryRunDto**](TransportEntryDryRunDto.md) |  | [optional] 
**NonBlockingErrors** | Pointer to [**[]NonBlockingErrorDto**](NonBlockingErrorDto.md) |  | [optional] 

## Methods

### NewMissionEntryDryRunDto

`func NewMissionEntryDryRunDto() *MissionEntryDryRunDto`

NewMissionEntryDryRunDto instantiates a new MissionEntryDryRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionEntryDryRunDtoWithDefaults

`func NewMissionEntryDryRunDtoWithDefaults() *MissionEntryDryRunDto`

NewMissionEntryDryRunDtoWithDefaults instantiates a new MissionEntryDryRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReadyToBill

`func (o *MissionEntryDryRunDto) GetIsReadyToBill() bool`

GetIsReadyToBill returns the IsReadyToBill field if non-nil, zero value otherwise.

### GetIsReadyToBillOk

`func (o *MissionEntryDryRunDto) GetIsReadyToBillOk() (*bool, bool)`

GetIsReadyToBillOk returns a tuple with the IsReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadyToBill

`func (o *MissionEntryDryRunDto) SetIsReadyToBill(v bool)`

SetIsReadyToBill sets IsReadyToBill field to given value.

### HasIsReadyToBill

`func (o *MissionEntryDryRunDto) HasIsReadyToBill() bool`

HasIsReadyToBill returns a boolean if a field has been set.

### GetIsUniqueOrder

`func (o *MissionEntryDryRunDto) GetIsUniqueOrder() bool`

GetIsUniqueOrder returns the IsUniqueOrder field if non-nil, zero value otherwise.

### GetIsUniqueOrderOk

`func (o *MissionEntryDryRunDto) GetIsUniqueOrderOk() (*bool, bool)`

GetIsUniqueOrderOk returns a tuple with the IsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniqueOrder

`func (o *MissionEntryDryRunDto) SetIsUniqueOrder(v bool)`

SetIsUniqueOrder sets IsUniqueOrder field to given value.

### HasIsUniqueOrder

`func (o *MissionEntryDryRunDto) HasIsUniqueOrder() bool`

HasIsUniqueOrder returns a boolean if a field has been set.

### GetTransports

`func (o *MissionEntryDryRunDto) GetTransports() []TransportEntryDryRunDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *MissionEntryDryRunDto) GetTransportsOk() (*[]TransportEntryDryRunDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *MissionEntryDryRunDto) SetTransports(v []TransportEntryDryRunDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *MissionEntryDryRunDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetNonBlockingErrors

`func (o *MissionEntryDryRunDto) GetNonBlockingErrors() []NonBlockingErrorDto`

GetNonBlockingErrors returns the NonBlockingErrors field if non-nil, zero value otherwise.

### GetNonBlockingErrorsOk

`func (o *MissionEntryDryRunDto) GetNonBlockingErrorsOk() (*[]NonBlockingErrorDto, bool)`

GetNonBlockingErrorsOk returns a tuple with the NonBlockingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBlockingErrors

`func (o *MissionEntryDryRunDto) SetNonBlockingErrors(v []NonBlockingErrorDto)`

SetNonBlockingErrors sets NonBlockingErrors field to given value.

### HasNonBlockingErrors

`func (o *MissionEntryDryRunDto) HasNonBlockingErrors() bool`

HasNonBlockingErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


