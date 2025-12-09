# ApplyTerminatedStatusCommand

## Properties

| Name                 | Type                                                                               | Description                              | Notes      |
| -------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------- | ---------- |
| **TransportUid**     | **string**                                                                         | Mandatory. Transport unique identifier   |
| **TransportReceipt** | Pointer to **string**                                                              | Transport receipt                        | [optional] |
| **RealizedDateTime** | Pointer to **Time**                                                                | Transport termination realized date time | [optional] |
| **PickupData**       | Pointer to [**PickupDataOnStatusAppliedDto**](PickupDataOnStatusAppliedDto.md)     |                                          | [optional] |
| **DeliveryData**     | Pointer to [**DeliveryDataOnStatusAppliedDto**](DeliveryDataOnStatusAppliedDto.md) |                                          | [optional] |

## Methods

### NewApplyTerminatedStatusCommand

`func NewApplyTerminatedStatusCommand(transportUid string, ) *ApplyTerminatedStatusCommand`

NewApplyTerminatedStatusCommand instantiates a new ApplyTerminatedStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTerminatedStatusCommandWithDefaults

`func NewApplyTerminatedStatusCommandWithDefaults() *ApplyTerminatedStatusCommand`

NewApplyTerminatedStatusCommandWithDefaults instantiates a new ApplyTerminatedStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *ApplyTerminatedStatusCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *ApplyTerminatedStatusCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *ApplyTerminatedStatusCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.

### GetTransportReceipt

`func (o *ApplyTerminatedStatusCommand) GetTransportReceipt() string`

GetTransportReceipt returns the TransportReceipt field if non-nil, zero value otherwise.

### GetTransportReceiptOk

`func (o *ApplyTerminatedStatusCommand) GetTransportReceiptOk() (*string, bool)`

GetTransportReceiptOk returns a tuple with the TransportReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReceipt

`func (o *ApplyTerminatedStatusCommand) SetTransportReceipt(v string)`

SetTransportReceipt sets TransportReceipt field to given value.

### HasTransportReceipt

`func (o *ApplyTerminatedStatusCommand) HasTransportReceipt() bool`

HasTransportReceipt returns a boolean if a field has been set.

### GetRealizedDateTime

`func (o *ApplyTerminatedStatusCommand) GetRealizedDateTime() Time`

GetRealizedDateTime returns the RealizedDateTime field if non-nil, zero value otherwise.

### GetRealizedDateTimeOk

`func (o *ApplyTerminatedStatusCommand) GetRealizedDateTimeOk() (*Time, bool)`

GetRealizedDateTimeOk returns a tuple with the RealizedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedDateTime

`func (o *ApplyTerminatedStatusCommand) SetRealizedDateTime(v Time)`

SetRealizedDateTime sets RealizedDateTime field to given value.

### HasRealizedDateTime

`func (o *ApplyTerminatedStatusCommand) HasRealizedDateTime() bool`

HasRealizedDateTime returns a boolean if a field has been set.

### GetPickupData

`func (o *ApplyTerminatedStatusCommand) GetPickupData() PickupDataOnStatusAppliedDto`

GetPickupData returns the PickupData field if non-nil, zero value otherwise.

### GetPickupDataOk

`func (o *ApplyTerminatedStatusCommand) GetPickupDataOk() (*PickupDataOnStatusAppliedDto, bool)`

GetPickupDataOk returns a tuple with the PickupData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupData

`func (o *ApplyTerminatedStatusCommand) SetPickupData(v PickupDataOnStatusAppliedDto)`

SetPickupData sets PickupData field to given value.

### HasPickupData

`func (o *ApplyTerminatedStatusCommand) HasPickupData() bool`

HasPickupData returns a boolean if a field has been set.

### GetDeliveryData

`func (o *ApplyTerminatedStatusCommand) GetDeliveryData() DeliveryDataOnStatusAppliedDto`

GetDeliveryData returns the DeliveryData field if non-nil, zero value otherwise.

### GetDeliveryDataOk

`func (o *ApplyTerminatedStatusCommand) GetDeliveryDataOk() (*DeliveryDataOnStatusAppliedDto, bool)`

GetDeliveryDataOk returns a tuple with the DeliveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryData

`func (o *ApplyTerminatedStatusCommand) SetDeliveryData(v DeliveryDataOnStatusAppliedDto)`

SetDeliveryData sets DeliveryData field to given value.

### HasDeliveryData

`func (o *ApplyTerminatedStatusCommand) HasDeliveryData() bool`

HasDeliveryData returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
