# CreateMissionResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissionNumber** | Pointer to **int32** |  | [optional] 
**MissionUid** | Pointer to **string** |  | [optional] 
**Transports** | Pointer to [**[]CreateTransportResultDto**](CreateTransportResultDto.md) |  | [optional] 
**NonBlockingErrors** | Pointer to [**[]NonBlockingErrorDto**](NonBlockingErrorDto.md) |  | [optional] 

## Methods

### NewCreateMissionResultDto

`func NewCreateMissionResultDto() *CreateMissionResultDto`

NewCreateMissionResultDto instantiates a new CreateMissionResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMissionResultDtoWithDefaults

`func NewCreateMissionResultDtoWithDefaults() *CreateMissionResultDto`

NewCreateMissionResultDtoWithDefaults instantiates a new CreateMissionResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissionNumber

`func (o *CreateMissionResultDto) GetMissionNumber() int32`

GetMissionNumber returns the MissionNumber field if non-nil, zero value otherwise.

### GetMissionNumberOk

`func (o *CreateMissionResultDto) GetMissionNumberOk() (*int32, bool)`

GetMissionNumberOk returns a tuple with the MissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionNumber

`func (o *CreateMissionResultDto) SetMissionNumber(v int32)`

SetMissionNumber sets MissionNumber field to given value.

### HasMissionNumber

`func (o *CreateMissionResultDto) HasMissionNumber() bool`

HasMissionNumber returns a boolean if a field has been set.

### GetMissionUid

`func (o *CreateMissionResultDto) GetMissionUid() string`

GetMissionUid returns the MissionUid field if non-nil, zero value otherwise.

### GetMissionUidOk

`func (o *CreateMissionResultDto) GetMissionUidOk() (*string, bool)`

GetMissionUidOk returns a tuple with the MissionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionUid

`func (o *CreateMissionResultDto) SetMissionUid(v string)`

SetMissionUid sets MissionUid field to given value.

### HasMissionUid

`func (o *CreateMissionResultDto) HasMissionUid() bool`

HasMissionUid returns a boolean if a field has been set.

### GetTransports

`func (o *CreateMissionResultDto) GetTransports() []CreateTransportResultDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateMissionResultDto) GetTransportsOk() (*[]CreateTransportResultDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateMissionResultDto) SetTransports(v []CreateTransportResultDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *CreateMissionResultDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetNonBlockingErrors

`func (o *CreateMissionResultDto) GetNonBlockingErrors() []NonBlockingErrorDto`

GetNonBlockingErrors returns the NonBlockingErrors field if non-nil, zero value otherwise.

### GetNonBlockingErrorsOk

`func (o *CreateMissionResultDto) GetNonBlockingErrorsOk() (*[]NonBlockingErrorDto, bool)`

GetNonBlockingErrorsOk returns a tuple with the NonBlockingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBlockingErrors

`func (o *CreateMissionResultDto) SetNonBlockingErrors(v []NonBlockingErrorDto)`

SetNonBlockingErrors sets NonBlockingErrors field to given value.

### HasNonBlockingErrors

`func (o *CreateMissionResultDto) HasNonBlockingErrors() bool`

HasNonBlockingErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


