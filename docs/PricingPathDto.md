# PricingPathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingPathId** | Pointer to **int32** |  | [optional] 
**PriceCatalogCode** | Pointer to **string** |  | [optional] 
**ServiceCode** | Pointer to **string** |  | [optional] 
**CustomerCode** | Pointer to **string** |  | [optional] 
**SubcontractorCode** | Pointer to **string** |  | [optional] 
**DistanceKm** | Pointer to **float64** | Distance of the pricing path in kilometers | [optional] 
**WithReturn** | Pointer to **bool** |  | [optional] 
**PickupDelayTimeMode** | Pointer to **string** |  | [optional] 
**PickupDurationTotalMinutes** | Pointer to **int32** |  | [optional] 
**PickupDelayDays** | Pointer to **int32** |  | [optional] 
**PickupDelayTime** | Pointer to **string** |  | [optional] 
**DeliveryDelayTimeMode** | Pointer to **string** |  | [optional] 
**DeliveryDurationTotalMinutes** | Pointer to **int32** |  | [optional] 
**DeliveryDelayDays** | Pointer to **int32** |  | [optional] 
**DeliveryDelayTime** | Pointer to **string** |  | [optional] 
**DeliveryTimeSlotValueMinutes** | Pointer to **int32** |  | [optional] 
**DeliveryPoint** | Pointer to [**PricingPointDto**](PricingPointDto.md) |  | [optional] 
**PickupPoint** | Pointer to [**PricingPointDto**](PricingPointDto.md) |  | [optional] 
**SubServices** | Pointer to [**CappedCollectionDtoPricingPathSubServiceDto**](CappedCollectionDtoPricingPathSubServiceDto.md) |  | [optional] 
**SubServiceAdditionalUnits** | Pointer to [**CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto**](CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto.md) |  | [optional] 
**SubServiceUnitPricingSlots** | Pointer to [**CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto**](CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto.md) |  | [optional] 
**UnitPricingSlotEquivalences** | Pointer to [**CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto**](CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto.md) |  | [optional] 

## Methods

### NewPricingPathDto

`func NewPricingPathDto() *PricingPathDto`

NewPricingPathDto instantiates a new PricingPathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPathDtoWithDefaults

`func NewPricingPathDtoWithDefaults() *PricingPathDto`

NewPricingPathDtoWithDefaults instantiates a new PricingPathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingPathId

`func (o *PricingPathDto) GetPricingPathId() int32`

GetPricingPathId returns the PricingPathId field if non-nil, zero value otherwise.

### GetPricingPathIdOk

`func (o *PricingPathDto) GetPricingPathIdOk() (*int32, bool)`

GetPricingPathIdOk returns a tuple with the PricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingPathId

`func (o *PricingPathDto) SetPricingPathId(v int32)`

SetPricingPathId sets PricingPathId field to given value.

### HasPricingPathId

`func (o *PricingPathDto) HasPricingPathId() bool`

HasPricingPathId returns a boolean if a field has been set.

### GetPriceCatalogCode

`func (o *PricingPathDto) GetPriceCatalogCode() string`

GetPriceCatalogCode returns the PriceCatalogCode field if non-nil, zero value otherwise.

### GetPriceCatalogCodeOk

`func (o *PricingPathDto) GetPriceCatalogCodeOk() (*string, bool)`

GetPriceCatalogCodeOk returns a tuple with the PriceCatalogCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCatalogCode

`func (o *PricingPathDto) SetPriceCatalogCode(v string)`

SetPriceCatalogCode sets PriceCatalogCode field to given value.

### HasPriceCatalogCode

`func (o *PricingPathDto) HasPriceCatalogCode() bool`

HasPriceCatalogCode returns a boolean if a field has been set.

### GetServiceCode

`func (o *PricingPathDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *PricingPathDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *PricingPathDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *PricingPathDto) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetCustomerCode

`func (o *PricingPathDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *PricingPathDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *PricingPathDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *PricingPathDto) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetSubcontractorCode

`func (o *PricingPathDto) GetSubcontractorCode() string`

GetSubcontractorCode returns the SubcontractorCode field if non-nil, zero value otherwise.

### GetSubcontractorCodeOk

`func (o *PricingPathDto) GetSubcontractorCodeOk() (*string, bool)`

GetSubcontractorCodeOk returns a tuple with the SubcontractorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcontractorCode

`func (o *PricingPathDto) SetSubcontractorCode(v string)`

SetSubcontractorCode sets SubcontractorCode field to given value.

### HasSubcontractorCode

`func (o *PricingPathDto) HasSubcontractorCode() bool`

HasSubcontractorCode returns a boolean if a field has been set.

### GetDistanceKm

`func (o *PricingPathDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *PricingPathDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *PricingPathDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *PricingPathDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetWithReturn

`func (o *PricingPathDto) GetWithReturn() bool`

GetWithReturn returns the WithReturn field if non-nil, zero value otherwise.

### GetWithReturnOk

`func (o *PricingPathDto) GetWithReturnOk() (*bool, bool)`

GetWithReturnOk returns a tuple with the WithReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithReturn

`func (o *PricingPathDto) SetWithReturn(v bool)`

SetWithReturn sets WithReturn field to given value.

### HasWithReturn

`func (o *PricingPathDto) HasWithReturn() bool`

HasWithReturn returns a boolean if a field has been set.

### GetPickupDelayTimeMode

`func (o *PricingPathDto) GetPickupDelayTimeMode() string`

GetPickupDelayTimeMode returns the PickupDelayTimeMode field if non-nil, zero value otherwise.

### GetPickupDelayTimeModeOk

`func (o *PricingPathDto) GetPickupDelayTimeModeOk() (*string, bool)`

GetPickupDelayTimeModeOk returns a tuple with the PickupDelayTimeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelayTimeMode

`func (o *PricingPathDto) SetPickupDelayTimeMode(v string)`

SetPickupDelayTimeMode sets PickupDelayTimeMode field to given value.

### HasPickupDelayTimeMode

`func (o *PricingPathDto) HasPickupDelayTimeMode() bool`

HasPickupDelayTimeMode returns a boolean if a field has been set.

### GetPickupDurationTotalMinutes

`func (o *PricingPathDto) GetPickupDurationTotalMinutes() int32`

GetPickupDurationTotalMinutes returns the PickupDurationTotalMinutes field if non-nil, zero value otherwise.

### GetPickupDurationTotalMinutesOk

`func (o *PricingPathDto) GetPickupDurationTotalMinutesOk() (*int32, bool)`

GetPickupDurationTotalMinutesOk returns a tuple with the PickupDurationTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDurationTotalMinutes

`func (o *PricingPathDto) SetPickupDurationTotalMinutes(v int32)`

SetPickupDurationTotalMinutes sets PickupDurationTotalMinutes field to given value.

### HasPickupDurationTotalMinutes

`func (o *PricingPathDto) HasPickupDurationTotalMinutes() bool`

HasPickupDurationTotalMinutes returns a boolean if a field has been set.

### GetPickupDelayDays

`func (o *PricingPathDto) GetPickupDelayDays() int32`

GetPickupDelayDays returns the PickupDelayDays field if non-nil, zero value otherwise.

### GetPickupDelayDaysOk

`func (o *PricingPathDto) GetPickupDelayDaysOk() (*int32, bool)`

GetPickupDelayDaysOk returns a tuple with the PickupDelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelayDays

`func (o *PricingPathDto) SetPickupDelayDays(v int32)`

SetPickupDelayDays sets PickupDelayDays field to given value.

### HasPickupDelayDays

`func (o *PricingPathDto) HasPickupDelayDays() bool`

HasPickupDelayDays returns a boolean if a field has been set.

### GetPickupDelayTime

`func (o *PricingPathDto) GetPickupDelayTime() string`

GetPickupDelayTime returns the PickupDelayTime field if non-nil, zero value otherwise.

### GetPickupDelayTimeOk

`func (o *PricingPathDto) GetPickupDelayTimeOk() (*string, bool)`

GetPickupDelayTimeOk returns a tuple with the PickupDelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupDelayTime

`func (o *PricingPathDto) SetPickupDelayTime(v string)`

SetPickupDelayTime sets PickupDelayTime field to given value.

### HasPickupDelayTime

`func (o *PricingPathDto) HasPickupDelayTime() bool`

HasPickupDelayTime returns a boolean if a field has been set.

### GetDeliveryDelayTimeMode

`func (o *PricingPathDto) GetDeliveryDelayTimeMode() string`

GetDeliveryDelayTimeMode returns the DeliveryDelayTimeMode field if non-nil, zero value otherwise.

### GetDeliveryDelayTimeModeOk

`func (o *PricingPathDto) GetDeliveryDelayTimeModeOk() (*string, bool)`

GetDeliveryDelayTimeModeOk returns a tuple with the DeliveryDelayTimeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDelayTimeMode

`func (o *PricingPathDto) SetDeliveryDelayTimeMode(v string)`

SetDeliveryDelayTimeMode sets DeliveryDelayTimeMode field to given value.

### HasDeliveryDelayTimeMode

`func (o *PricingPathDto) HasDeliveryDelayTimeMode() bool`

HasDeliveryDelayTimeMode returns a boolean if a field has been set.

### GetDeliveryDurationTotalMinutes

`func (o *PricingPathDto) GetDeliveryDurationTotalMinutes() int32`

GetDeliveryDurationTotalMinutes returns the DeliveryDurationTotalMinutes field if non-nil, zero value otherwise.

### GetDeliveryDurationTotalMinutesOk

`func (o *PricingPathDto) GetDeliveryDurationTotalMinutesOk() (*int32, bool)`

GetDeliveryDurationTotalMinutesOk returns a tuple with the DeliveryDurationTotalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDurationTotalMinutes

`func (o *PricingPathDto) SetDeliveryDurationTotalMinutes(v int32)`

SetDeliveryDurationTotalMinutes sets DeliveryDurationTotalMinutes field to given value.

### HasDeliveryDurationTotalMinutes

`func (o *PricingPathDto) HasDeliveryDurationTotalMinutes() bool`

HasDeliveryDurationTotalMinutes returns a boolean if a field has been set.

### GetDeliveryDelayDays

`func (o *PricingPathDto) GetDeliveryDelayDays() int32`

GetDeliveryDelayDays returns the DeliveryDelayDays field if non-nil, zero value otherwise.

### GetDeliveryDelayDaysOk

`func (o *PricingPathDto) GetDeliveryDelayDaysOk() (*int32, bool)`

GetDeliveryDelayDaysOk returns a tuple with the DeliveryDelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDelayDays

`func (o *PricingPathDto) SetDeliveryDelayDays(v int32)`

SetDeliveryDelayDays sets DeliveryDelayDays field to given value.

### HasDeliveryDelayDays

`func (o *PricingPathDto) HasDeliveryDelayDays() bool`

HasDeliveryDelayDays returns a boolean if a field has been set.

### GetDeliveryDelayTime

`func (o *PricingPathDto) GetDeliveryDelayTime() string`

GetDeliveryDelayTime returns the DeliveryDelayTime field if non-nil, zero value otherwise.

### GetDeliveryDelayTimeOk

`func (o *PricingPathDto) GetDeliveryDelayTimeOk() (*string, bool)`

GetDeliveryDelayTimeOk returns a tuple with the DeliveryDelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDelayTime

`func (o *PricingPathDto) SetDeliveryDelayTime(v string)`

SetDeliveryDelayTime sets DeliveryDelayTime field to given value.

### HasDeliveryDelayTime

`func (o *PricingPathDto) HasDeliveryDelayTime() bool`

HasDeliveryDelayTime returns a boolean if a field has been set.

### GetDeliveryTimeSlotValueMinutes

`func (o *PricingPathDto) GetDeliveryTimeSlotValueMinutes() int32`

GetDeliveryTimeSlotValueMinutes returns the DeliveryTimeSlotValueMinutes field if non-nil, zero value otherwise.

### GetDeliveryTimeSlotValueMinutesOk

`func (o *PricingPathDto) GetDeliveryTimeSlotValueMinutesOk() (*int32, bool)`

GetDeliveryTimeSlotValueMinutesOk returns a tuple with the DeliveryTimeSlotValueMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeSlotValueMinutes

`func (o *PricingPathDto) SetDeliveryTimeSlotValueMinutes(v int32)`

SetDeliveryTimeSlotValueMinutes sets DeliveryTimeSlotValueMinutes field to given value.

### HasDeliveryTimeSlotValueMinutes

`func (o *PricingPathDto) HasDeliveryTimeSlotValueMinutes() bool`

HasDeliveryTimeSlotValueMinutes returns a boolean if a field has been set.

### GetDeliveryPoint

`func (o *PricingPathDto) GetDeliveryPoint() PricingPointDto`

GetDeliveryPoint returns the DeliveryPoint field if non-nil, zero value otherwise.

### GetDeliveryPointOk

`func (o *PricingPathDto) GetDeliveryPointOk() (*PricingPointDto, bool)`

GetDeliveryPointOk returns a tuple with the DeliveryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPoint

`func (o *PricingPathDto) SetDeliveryPoint(v PricingPointDto)`

SetDeliveryPoint sets DeliveryPoint field to given value.

### HasDeliveryPoint

`func (o *PricingPathDto) HasDeliveryPoint() bool`

HasDeliveryPoint returns a boolean if a field has been set.

### GetPickupPoint

`func (o *PricingPathDto) GetPickupPoint() PricingPointDto`

GetPickupPoint returns the PickupPoint field if non-nil, zero value otherwise.

### GetPickupPointOk

`func (o *PricingPathDto) GetPickupPointOk() (*PricingPointDto, bool)`

GetPickupPointOk returns a tuple with the PickupPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupPoint

`func (o *PricingPathDto) SetPickupPoint(v PricingPointDto)`

SetPickupPoint sets PickupPoint field to given value.

### HasPickupPoint

`func (o *PricingPathDto) HasPickupPoint() bool`

HasPickupPoint returns a boolean if a field has been set.

### GetSubServices

`func (o *PricingPathDto) GetSubServices() CappedCollectionDtoPricingPathSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *PricingPathDto) GetSubServicesOk() (*CappedCollectionDtoPricingPathSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *PricingPathDto) SetSubServices(v CappedCollectionDtoPricingPathSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *PricingPathDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetSubServiceAdditionalUnits

`func (o *PricingPathDto) GetSubServiceAdditionalUnits() CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto`

GetSubServiceAdditionalUnits returns the SubServiceAdditionalUnits field if non-nil, zero value otherwise.

### GetSubServiceAdditionalUnitsOk

`func (o *PricingPathDto) GetSubServiceAdditionalUnitsOk() (*CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto, bool)`

GetSubServiceAdditionalUnitsOk returns a tuple with the SubServiceAdditionalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServiceAdditionalUnits

`func (o *PricingPathDto) SetSubServiceAdditionalUnits(v CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto)`

SetSubServiceAdditionalUnits sets SubServiceAdditionalUnits field to given value.

### HasSubServiceAdditionalUnits

`func (o *PricingPathDto) HasSubServiceAdditionalUnits() bool`

HasSubServiceAdditionalUnits returns a boolean if a field has been set.

### GetSubServiceUnitPricingSlots

`func (o *PricingPathDto) GetSubServiceUnitPricingSlots() CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto`

GetSubServiceUnitPricingSlots returns the SubServiceUnitPricingSlots field if non-nil, zero value otherwise.

### GetSubServiceUnitPricingSlotsOk

`func (o *PricingPathDto) GetSubServiceUnitPricingSlotsOk() (*CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto, bool)`

GetSubServiceUnitPricingSlotsOk returns a tuple with the SubServiceUnitPricingSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServiceUnitPricingSlots

`func (o *PricingPathDto) SetSubServiceUnitPricingSlots(v CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto)`

SetSubServiceUnitPricingSlots sets SubServiceUnitPricingSlots field to given value.

### HasSubServiceUnitPricingSlots

`func (o *PricingPathDto) HasSubServiceUnitPricingSlots() bool`

HasSubServiceUnitPricingSlots returns a boolean if a field has been set.

### GetUnitPricingSlotEquivalences

`func (o *PricingPathDto) GetUnitPricingSlotEquivalences() CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto`

GetUnitPricingSlotEquivalences returns the UnitPricingSlotEquivalences field if non-nil, zero value otherwise.

### GetUnitPricingSlotEquivalencesOk

`func (o *PricingPathDto) GetUnitPricingSlotEquivalencesOk() (*CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto, bool)`

GetUnitPricingSlotEquivalencesOk returns a tuple with the UnitPricingSlotEquivalences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricingSlotEquivalences

`func (o *PricingPathDto) SetUnitPricingSlotEquivalences(v CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto)`

SetUnitPricingSlotEquivalences sets UnitPricingSlotEquivalences field to given value.

### HasUnitPricingSlotEquivalences

`func (o *PricingPathDto) HasUnitPricingSlotEquivalences() bool`

HasUnitPricingSlotEquivalences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


