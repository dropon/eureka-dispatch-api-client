# MissionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Mission&#39;s internal unique identifier | [optional] 
**MissionNumber** | Pointer to **int32** | Mission&#39;s public identifier | [optional] 
**TrackId** | Pointer to **string** | Mission&#39;s tracking number | [optional] 
**IsReadyToBill** | Pointer to **bool** |  | [optional] 
**IsUniqueOrder** | Pointer to **bool** |  | [optional] 
**Transports** | Pointer to [**CappedCollectionDtoTransportDto**](CappedCollectionDtoTransportDto.md) |  | [optional] 

## Methods

### NewMissionDto

`func NewMissionDto() *MissionDto`

NewMissionDto instantiates a new MissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionDtoWithDefaults

`func NewMissionDtoWithDefaults() *MissionDto`

NewMissionDtoWithDefaults instantiates a new MissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *MissionDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MissionDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MissionDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MissionDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetMissionNumber

`func (o *MissionDto) GetMissionNumber() int32`

GetMissionNumber returns the MissionNumber field if non-nil, zero value otherwise.

### GetMissionNumberOk

`func (o *MissionDto) GetMissionNumberOk() (*int32, bool)`

GetMissionNumberOk returns a tuple with the MissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionNumber

`func (o *MissionDto) SetMissionNumber(v int32)`

SetMissionNumber sets MissionNumber field to given value.

### HasMissionNumber

`func (o *MissionDto) HasMissionNumber() bool`

HasMissionNumber returns a boolean if a field has been set.

### GetTrackId

`func (o *MissionDto) GetTrackId() string`

GetTrackId returns the TrackId field if non-nil, zero value otherwise.

### GetTrackIdOk

`func (o *MissionDto) GetTrackIdOk() (*string, bool)`

GetTrackIdOk returns a tuple with the TrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackId

`func (o *MissionDto) SetTrackId(v string)`

SetTrackId sets TrackId field to given value.

### HasTrackId

`func (o *MissionDto) HasTrackId() bool`

HasTrackId returns a boolean if a field has been set.

### GetIsReadyToBill

`func (o *MissionDto) GetIsReadyToBill() bool`

GetIsReadyToBill returns the IsReadyToBill field if non-nil, zero value otherwise.

### GetIsReadyToBillOk

`func (o *MissionDto) GetIsReadyToBillOk() (*bool, bool)`

GetIsReadyToBillOk returns a tuple with the IsReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadyToBill

`func (o *MissionDto) SetIsReadyToBill(v bool)`

SetIsReadyToBill sets IsReadyToBill field to given value.

### HasIsReadyToBill

`func (o *MissionDto) HasIsReadyToBill() bool`

HasIsReadyToBill returns a boolean if a field has been set.

### GetIsUniqueOrder

`func (o *MissionDto) GetIsUniqueOrder() bool`

GetIsUniqueOrder returns the IsUniqueOrder field if non-nil, zero value otherwise.

### GetIsUniqueOrderOk

`func (o *MissionDto) GetIsUniqueOrderOk() (*bool, bool)`

GetIsUniqueOrderOk returns a tuple with the IsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniqueOrder

`func (o *MissionDto) SetIsUniqueOrder(v bool)`

SetIsUniqueOrder sets IsUniqueOrder field to given value.

### HasIsUniqueOrder

`func (o *MissionDto) HasIsUniqueOrder() bool`

HasIsUniqueOrder returns a boolean if a field has been set.

### GetTransports

`func (o *MissionDto) GetTransports() CappedCollectionDtoTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *MissionDto) GetTransportsOk() (*CappedCollectionDtoTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *MissionDto) SetTransports(v CappedCollectionDtoTransportDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *MissionDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


