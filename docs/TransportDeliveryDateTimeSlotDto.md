# TransportDeliveryDateTimeSlotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnloadingStart** | Pointer to **time.Time** |  | [optional] 
**UnloadingEnd** | Pointer to **time.Time** |  | [optional] 
**PlannedStart** | Pointer to [**SlotDto**](SlotDto.md) |  | [optional] 
**PlannedEnd** | Pointer to [**SlotDto**](SlotDto.md) |  | [optional] 
**ArrivalOnSite** | Pointer to **time.Time** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTransportDeliveryDateTimeSlotDto

`func NewTransportDeliveryDateTimeSlotDto() *TransportDeliveryDateTimeSlotDto`

NewTransportDeliveryDateTimeSlotDto instantiates a new TransportDeliveryDateTimeSlotDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportDeliveryDateTimeSlotDtoWithDefaults

`func NewTransportDeliveryDateTimeSlotDtoWithDefaults() *TransportDeliveryDateTimeSlotDto`

NewTransportDeliveryDateTimeSlotDtoWithDefaults instantiates a new TransportDeliveryDateTimeSlotDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnloadingStart

`func (o *TransportDeliveryDateTimeSlotDto) GetUnloadingStart() time.Time`

GetUnloadingStart returns the UnloadingStart field if non-nil, zero value otherwise.

### GetUnloadingStartOk

`func (o *TransportDeliveryDateTimeSlotDto) GetUnloadingStartOk() (*time.Time, bool)`

GetUnloadingStartOk returns a tuple with the UnloadingStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnloadingStart

`func (o *TransportDeliveryDateTimeSlotDto) SetUnloadingStart(v time.Time)`

SetUnloadingStart sets UnloadingStart field to given value.

### HasUnloadingStart

`func (o *TransportDeliveryDateTimeSlotDto) HasUnloadingStart() bool`

HasUnloadingStart returns a boolean if a field has been set.

### GetUnloadingEnd

`func (o *TransportDeliveryDateTimeSlotDto) GetUnloadingEnd() time.Time`

GetUnloadingEnd returns the UnloadingEnd field if non-nil, zero value otherwise.

### GetUnloadingEndOk

`func (o *TransportDeliveryDateTimeSlotDto) GetUnloadingEndOk() (*time.Time, bool)`

GetUnloadingEndOk returns a tuple with the UnloadingEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnloadingEnd

`func (o *TransportDeliveryDateTimeSlotDto) SetUnloadingEnd(v time.Time)`

SetUnloadingEnd sets UnloadingEnd field to given value.

### HasUnloadingEnd

`func (o *TransportDeliveryDateTimeSlotDto) HasUnloadingEnd() bool`

HasUnloadingEnd returns a boolean if a field has been set.

### GetPlannedStart

`func (o *TransportDeliveryDateTimeSlotDto) GetPlannedStart() SlotDto`

GetPlannedStart returns the PlannedStart field if non-nil, zero value otherwise.

### GetPlannedStartOk

`func (o *TransportDeliveryDateTimeSlotDto) GetPlannedStartOk() (*SlotDto, bool)`

GetPlannedStartOk returns a tuple with the PlannedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStart

`func (o *TransportDeliveryDateTimeSlotDto) SetPlannedStart(v SlotDto)`

SetPlannedStart sets PlannedStart field to given value.

### HasPlannedStart

`func (o *TransportDeliveryDateTimeSlotDto) HasPlannedStart() bool`

HasPlannedStart returns a boolean if a field has been set.

### GetPlannedEnd

`func (o *TransportDeliveryDateTimeSlotDto) GetPlannedEnd() SlotDto`

GetPlannedEnd returns the PlannedEnd field if non-nil, zero value otherwise.

### GetPlannedEndOk

`func (o *TransportDeliveryDateTimeSlotDto) GetPlannedEndOk() (*SlotDto, bool)`

GetPlannedEndOk returns a tuple with the PlannedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEnd

`func (o *TransportDeliveryDateTimeSlotDto) SetPlannedEnd(v SlotDto)`

SetPlannedEnd sets PlannedEnd field to given value.

### HasPlannedEnd

`func (o *TransportDeliveryDateTimeSlotDto) HasPlannedEnd() bool`

HasPlannedEnd returns a boolean if a field has been set.

### GetArrivalOnSite

`func (o *TransportDeliveryDateTimeSlotDto) GetArrivalOnSite() time.Time`

GetArrivalOnSite returns the ArrivalOnSite field if non-nil, zero value otherwise.

### GetArrivalOnSiteOk

`func (o *TransportDeliveryDateTimeSlotDto) GetArrivalOnSiteOk() (*time.Time, bool)`

GetArrivalOnSiteOk returns a tuple with the ArrivalOnSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalOnSite

`func (o *TransportDeliveryDateTimeSlotDto) SetArrivalOnSite(v time.Time)`

SetArrivalOnSite sets ArrivalOnSite field to given value.

### HasArrivalOnSite

`func (o *TransportDeliveryDateTimeSlotDto) HasArrivalOnSite() bool`

HasArrivalOnSite returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *TransportDeliveryDateTimeSlotDto) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *TransportDeliveryDateTimeSlotDto) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *TransportDeliveryDateTimeSlotDto) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *TransportDeliveryDateTimeSlotDto) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


