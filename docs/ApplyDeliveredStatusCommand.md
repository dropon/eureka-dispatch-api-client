# ApplyDeliveredStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportUid** | **string** | Mandatory. Transport&#39;s unique identifier | 
**TransportReceipt** | Pointer to **string** | Transport receipt | [optional] 
**DeliverySign** | Pointer to **string** | Transport delivery sign | [optional] 
**DeliveryComment** | Pointer to **string** | Transport delivery comment | [optional] 
**RealizedDateTime** | Pointer to **time.Time** | Transport delivery realized date time | [optional] 
**RealizedArrivalOnSiteDateTime** | Pointer to **time.Time** | Transport delivery realized arrival on site date time | [optional] 
**RealizedUnloadingStartDateTime** | Pointer to **time.Time** | Transport delivery realized unloading start date time | [optional] 
**RealizedUnloadingEndDateTime** | Pointer to **time.Time** | Transport delivery realized unloading end date time | [optional] 
**PickupData** | Pointer to [**PickupDataOnStatusAppliedDto**](PickupDataOnStatusAppliedDto.md) |  | [optional] 

## Methods

### NewApplyDeliveredStatusCommand

`func NewApplyDeliveredStatusCommand(transportUid string, ) *ApplyDeliveredStatusCommand`

NewApplyDeliveredStatusCommand instantiates a new ApplyDeliveredStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyDeliveredStatusCommandWithDefaults

`func NewApplyDeliveredStatusCommandWithDefaults() *ApplyDeliveredStatusCommand`

NewApplyDeliveredStatusCommandWithDefaults instantiates a new ApplyDeliveredStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *ApplyDeliveredStatusCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *ApplyDeliveredStatusCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *ApplyDeliveredStatusCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.


### GetTransportReceipt

`func (o *ApplyDeliveredStatusCommand) GetTransportReceipt() string`

GetTransportReceipt returns the TransportReceipt field if non-nil, zero value otherwise.

### GetTransportReceiptOk

`func (o *ApplyDeliveredStatusCommand) GetTransportReceiptOk() (*string, bool)`

GetTransportReceiptOk returns a tuple with the TransportReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReceipt

`func (o *ApplyDeliveredStatusCommand) SetTransportReceipt(v string)`

SetTransportReceipt sets TransportReceipt field to given value.

### HasTransportReceipt

`func (o *ApplyDeliveredStatusCommand) HasTransportReceipt() bool`

HasTransportReceipt returns a boolean if a field has been set.

### GetDeliverySign

`func (o *ApplyDeliveredStatusCommand) GetDeliverySign() string`

GetDeliverySign returns the DeliverySign field if non-nil, zero value otherwise.

### GetDeliverySignOk

`func (o *ApplyDeliveredStatusCommand) GetDeliverySignOk() (*string, bool)`

GetDeliverySignOk returns a tuple with the DeliverySign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySign

`func (o *ApplyDeliveredStatusCommand) SetDeliverySign(v string)`

SetDeliverySign sets DeliverySign field to given value.

### HasDeliverySign

`func (o *ApplyDeliveredStatusCommand) HasDeliverySign() bool`

HasDeliverySign returns a boolean if a field has been set.

### GetDeliveryComment

`func (o *ApplyDeliveredStatusCommand) GetDeliveryComment() string`

GetDeliveryComment returns the DeliveryComment field if non-nil, zero value otherwise.

### GetDeliveryCommentOk

`func (o *ApplyDeliveredStatusCommand) GetDeliveryCommentOk() (*string, bool)`

GetDeliveryCommentOk returns a tuple with the DeliveryComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryComment

`func (o *ApplyDeliveredStatusCommand) SetDeliveryComment(v string)`

SetDeliveryComment sets DeliveryComment field to given value.

### HasDeliveryComment

`func (o *ApplyDeliveredStatusCommand) HasDeliveryComment() bool`

HasDeliveryComment returns a boolean if a field has been set.

### GetRealizedDateTime

`func (o *ApplyDeliveredStatusCommand) GetRealizedDateTime() time.Time`

GetRealizedDateTime returns the RealizedDateTime field if non-nil, zero value otherwise.

### GetRealizedDateTimeOk

`func (o *ApplyDeliveredStatusCommand) GetRealizedDateTimeOk() (*time.Time, bool)`

GetRealizedDateTimeOk returns a tuple with the RealizedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedDateTime

`func (o *ApplyDeliveredStatusCommand) SetRealizedDateTime(v time.Time)`

SetRealizedDateTime sets RealizedDateTime field to given value.

### HasRealizedDateTime

`func (o *ApplyDeliveredStatusCommand) HasRealizedDateTime() bool`

HasRealizedDateTime returns a boolean if a field has been set.

### GetRealizedArrivalOnSiteDateTime

`func (o *ApplyDeliveredStatusCommand) GetRealizedArrivalOnSiteDateTime() time.Time`

GetRealizedArrivalOnSiteDateTime returns the RealizedArrivalOnSiteDateTime field if non-nil, zero value otherwise.

### GetRealizedArrivalOnSiteDateTimeOk

`func (o *ApplyDeliveredStatusCommand) GetRealizedArrivalOnSiteDateTimeOk() (*time.Time, bool)`

GetRealizedArrivalOnSiteDateTimeOk returns a tuple with the RealizedArrivalOnSiteDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedArrivalOnSiteDateTime

`func (o *ApplyDeliveredStatusCommand) SetRealizedArrivalOnSiteDateTime(v time.Time)`

SetRealizedArrivalOnSiteDateTime sets RealizedArrivalOnSiteDateTime field to given value.

### HasRealizedArrivalOnSiteDateTime

`func (o *ApplyDeliveredStatusCommand) HasRealizedArrivalOnSiteDateTime() bool`

HasRealizedArrivalOnSiteDateTime returns a boolean if a field has been set.

### GetRealizedUnloadingStartDateTime

`func (o *ApplyDeliveredStatusCommand) GetRealizedUnloadingStartDateTime() time.Time`

GetRealizedUnloadingStartDateTime returns the RealizedUnloadingStartDateTime field if non-nil, zero value otherwise.

### GetRealizedUnloadingStartDateTimeOk

`func (o *ApplyDeliveredStatusCommand) GetRealizedUnloadingStartDateTimeOk() (*time.Time, bool)`

GetRealizedUnloadingStartDateTimeOk returns a tuple with the RealizedUnloadingStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedUnloadingStartDateTime

`func (o *ApplyDeliveredStatusCommand) SetRealizedUnloadingStartDateTime(v time.Time)`

SetRealizedUnloadingStartDateTime sets RealizedUnloadingStartDateTime field to given value.

### HasRealizedUnloadingStartDateTime

`func (o *ApplyDeliveredStatusCommand) HasRealizedUnloadingStartDateTime() bool`

HasRealizedUnloadingStartDateTime returns a boolean if a field has been set.

### GetRealizedUnloadingEndDateTime

`func (o *ApplyDeliveredStatusCommand) GetRealizedUnloadingEndDateTime() time.Time`

GetRealizedUnloadingEndDateTime returns the RealizedUnloadingEndDateTime field if non-nil, zero value otherwise.

### GetRealizedUnloadingEndDateTimeOk

`func (o *ApplyDeliveredStatusCommand) GetRealizedUnloadingEndDateTimeOk() (*time.Time, bool)`

GetRealizedUnloadingEndDateTimeOk returns a tuple with the RealizedUnloadingEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedUnloadingEndDateTime

`func (o *ApplyDeliveredStatusCommand) SetRealizedUnloadingEndDateTime(v time.Time)`

SetRealizedUnloadingEndDateTime sets RealizedUnloadingEndDateTime field to given value.

### HasRealizedUnloadingEndDateTime

`func (o *ApplyDeliveredStatusCommand) HasRealizedUnloadingEndDateTime() bool`

HasRealizedUnloadingEndDateTime returns a boolean if a field has been set.

### GetPickupData

`func (o *ApplyDeliveredStatusCommand) GetPickupData() PickupDataOnStatusAppliedDto`

GetPickupData returns the PickupData field if non-nil, zero value otherwise.

### GetPickupDataOk

`func (o *ApplyDeliveredStatusCommand) GetPickupDataOk() (*PickupDataOnStatusAppliedDto, bool)`

GetPickupDataOk returns a tuple with the PickupData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupData

`func (o *ApplyDeliveredStatusCommand) SetPickupData(v PickupDataOnStatusAppliedDto)`

SetPickupData sets PickupData field to given value.

### HasPickupData

`func (o *ApplyDeliveredStatusCommand) HasPickupData() bool`

HasPickupData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


