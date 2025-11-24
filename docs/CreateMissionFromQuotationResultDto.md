# CreateMissionFromQuotationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissionNumber** | Pointer to **int32** |  | [optional] 
**MissionUid** | Pointer to **string** |  | [optional] 
**Transports** | Pointer to [**[]CreateTransportResultDto**](CreateTransportResultDto.md) |  | [optional] 

## Methods

### NewCreateMissionFromQuotationResultDto

`func NewCreateMissionFromQuotationResultDto() *CreateMissionFromQuotationResultDto`

NewCreateMissionFromQuotationResultDto instantiates a new CreateMissionFromQuotationResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMissionFromQuotationResultDtoWithDefaults

`func NewCreateMissionFromQuotationResultDtoWithDefaults() *CreateMissionFromQuotationResultDto`

NewCreateMissionFromQuotationResultDtoWithDefaults instantiates a new CreateMissionFromQuotationResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissionNumber

`func (o *CreateMissionFromQuotationResultDto) GetMissionNumber() int32`

GetMissionNumber returns the MissionNumber field if non-nil, zero value otherwise.

### GetMissionNumberOk

`func (o *CreateMissionFromQuotationResultDto) GetMissionNumberOk() (*int32, bool)`

GetMissionNumberOk returns a tuple with the MissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionNumber

`func (o *CreateMissionFromQuotationResultDto) SetMissionNumber(v int32)`

SetMissionNumber sets MissionNumber field to given value.

### HasMissionNumber

`func (o *CreateMissionFromQuotationResultDto) HasMissionNumber() bool`

HasMissionNumber returns a boolean if a field has been set.

### GetMissionUid

`func (o *CreateMissionFromQuotationResultDto) GetMissionUid() string`

GetMissionUid returns the MissionUid field if non-nil, zero value otherwise.

### GetMissionUidOk

`func (o *CreateMissionFromQuotationResultDto) GetMissionUidOk() (*string, bool)`

GetMissionUidOk returns a tuple with the MissionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionUid

`func (o *CreateMissionFromQuotationResultDto) SetMissionUid(v string)`

SetMissionUid sets MissionUid field to given value.

### HasMissionUid

`func (o *CreateMissionFromQuotationResultDto) HasMissionUid() bool`

HasMissionUid returns a boolean if a field has been set.

### GetTransports

`func (o *CreateMissionFromQuotationResultDto) GetTransports() []CreateTransportResultDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateMissionFromQuotationResultDto) GetTransportsOk() (*[]CreateTransportResultDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateMissionFromQuotationResultDto) SetTransports(v []CreateTransportResultDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *CreateMissionFromQuotationResultDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


