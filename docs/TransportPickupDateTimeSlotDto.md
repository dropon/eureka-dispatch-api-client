# TransportPickupDateTimeSlotDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadingStart** | Pointer to **time.Time** |  | [optional] 
**LoadingEnd** | Pointer to **time.Time** |  | [optional] 
**PlannedStart** | Pointer to [**SlotDto**](SlotDto.md) |  | [optional] 
**PlannedEnd** | Pointer to [**SlotDto**](SlotDto.md) |  | [optional] 
**ArrivalOnSite** | Pointer to **time.Time** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTransportPickupDateTimeSlotDto

`func NewTransportPickupDateTimeSlotDto() *TransportPickupDateTimeSlotDto`

NewTransportPickupDateTimeSlotDto instantiates a new TransportPickupDateTimeSlotDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPickupDateTimeSlotDtoWithDefaults

`func NewTransportPickupDateTimeSlotDtoWithDefaults() *TransportPickupDateTimeSlotDto`

NewTransportPickupDateTimeSlotDtoWithDefaults instantiates a new TransportPickupDateTimeSlotDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadingStart

`func (o *TransportPickupDateTimeSlotDto) GetLoadingStart() time.Time`

GetLoadingStart returns the LoadingStart field if non-nil, zero value otherwise.

### GetLoadingStartOk

`func (o *TransportPickupDateTimeSlotDto) GetLoadingStartOk() (*time.Time, bool)`

GetLoadingStartOk returns a tuple with the LoadingStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingStart

`func (o *TransportPickupDateTimeSlotDto) SetLoadingStart(v time.Time)`

SetLoadingStart sets LoadingStart field to given value.

### HasLoadingStart

`func (o *TransportPickupDateTimeSlotDto) HasLoadingStart() bool`

HasLoadingStart returns a boolean if a field has been set.

### GetLoadingEnd

`func (o *TransportPickupDateTimeSlotDto) GetLoadingEnd() time.Time`

GetLoadingEnd returns the LoadingEnd field if non-nil, zero value otherwise.

### GetLoadingEndOk

`func (o *TransportPickupDateTimeSlotDto) GetLoadingEndOk() (*time.Time, bool)`

GetLoadingEndOk returns a tuple with the LoadingEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingEnd

`func (o *TransportPickupDateTimeSlotDto) SetLoadingEnd(v time.Time)`

SetLoadingEnd sets LoadingEnd field to given value.

### HasLoadingEnd

`func (o *TransportPickupDateTimeSlotDto) HasLoadingEnd() bool`

HasLoadingEnd returns a boolean if a field has been set.

### GetPlannedStart

`func (o *TransportPickupDateTimeSlotDto) GetPlannedStart() SlotDto`

GetPlannedStart returns the PlannedStart field if non-nil, zero value otherwise.

### GetPlannedStartOk

`func (o *TransportPickupDateTimeSlotDto) GetPlannedStartOk() (*SlotDto, bool)`

GetPlannedStartOk returns a tuple with the PlannedStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStart

`func (o *TransportPickupDateTimeSlotDto) SetPlannedStart(v SlotDto)`

SetPlannedStart sets PlannedStart field to given value.

### HasPlannedStart

`func (o *TransportPickupDateTimeSlotDto) HasPlannedStart() bool`

HasPlannedStart returns a boolean if a field has been set.

### GetPlannedEnd

`func (o *TransportPickupDateTimeSlotDto) GetPlannedEnd() SlotDto`

GetPlannedEnd returns the PlannedEnd field if non-nil, zero value otherwise.

### GetPlannedEndOk

`func (o *TransportPickupDateTimeSlotDto) GetPlannedEndOk() (*SlotDto, bool)`

GetPlannedEndOk returns a tuple with the PlannedEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEnd

`func (o *TransportPickupDateTimeSlotDto) SetPlannedEnd(v SlotDto)`

SetPlannedEnd sets PlannedEnd field to given value.

### HasPlannedEnd

`func (o *TransportPickupDateTimeSlotDto) HasPlannedEnd() bool`

HasPlannedEnd returns a boolean if a field has been set.

### GetArrivalOnSite

`func (o *TransportPickupDateTimeSlotDto) GetArrivalOnSite() time.Time`

GetArrivalOnSite returns the ArrivalOnSite field if non-nil, zero value otherwise.

### GetArrivalOnSiteOk

`func (o *TransportPickupDateTimeSlotDto) GetArrivalOnSiteOk() (*time.Time, bool)`

GetArrivalOnSiteOk returns a tuple with the ArrivalOnSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalOnSite

`func (o *TransportPickupDateTimeSlotDto) SetArrivalOnSite(v time.Time)`

SetArrivalOnSite sets ArrivalOnSite field to given value.

### HasArrivalOnSite

`func (o *TransportPickupDateTimeSlotDto) HasArrivalOnSite() bool`

HasArrivalOnSite returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *TransportPickupDateTimeSlotDto) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *TransportPickupDateTimeSlotDto) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *TransportPickupDateTimeSlotDto) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *TransportPickupDateTimeSlotDto) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


