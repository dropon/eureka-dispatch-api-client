# UpdateMissionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | Pointer to [**[]MergeableTransportDto**](MergeableTransportDto.md) |  | [optional] 
**IsReadyToBill** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**UpdateMissionDtoApiBehaviorDto**](UpdateMissionDtoApiBehaviorDto.md) |  | [optional] 

## Methods

### NewUpdateMissionDto

`func NewUpdateMissionDto() *UpdateMissionDto`

NewUpdateMissionDto instantiates a new UpdateMissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMissionDtoWithDefaults

`func NewUpdateMissionDtoWithDefaults() *UpdateMissionDto`

NewUpdateMissionDtoWithDefaults instantiates a new UpdateMissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *UpdateMissionDto) GetTransports() []MergeableTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *UpdateMissionDto) GetTransportsOk() (*[]MergeableTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *UpdateMissionDto) SetTransports(v []MergeableTransportDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *UpdateMissionDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetIsReadyToBill

`func (o *UpdateMissionDto) GetIsReadyToBill() bool`

GetIsReadyToBill returns the IsReadyToBill field if non-nil, zero value otherwise.

### GetIsReadyToBillOk

`func (o *UpdateMissionDto) GetIsReadyToBillOk() (*bool, bool)`

GetIsReadyToBillOk returns a tuple with the IsReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadyToBill

`func (o *UpdateMissionDto) SetIsReadyToBill(v bool)`

SetIsReadyToBill sets IsReadyToBill field to given value.

### HasIsReadyToBill

`func (o *UpdateMissionDto) HasIsReadyToBill() bool`

HasIsReadyToBill returns a boolean if a field has been set.

### GetApiBehavior

`func (o *UpdateMissionDto) GetApiBehavior() UpdateMissionDtoApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *UpdateMissionDto) GetApiBehaviorOk() (*UpdateMissionDtoApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *UpdateMissionDto) SetApiBehavior(v UpdateMissionDtoApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *UpdateMissionDto) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


