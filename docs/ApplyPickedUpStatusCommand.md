# ApplyPickedUpStatusCommand

## Properties

| Name                              | Type                  | Description                                         | Notes      |
| --------------------------------- | --------------------- | --------------------------------------------------- | ---------- |
| **TransportUid**                  | **string**            | Mandatory. Transport unique identifier              |
| **TransportReceipt**              | Pointer to **string** | Transport receipt                                   | [optional] |
| **PickupSign**                    | Pointer to **string** | Transport pickup sign                               | [optional] |
| **PickupComment**                 | Pointer to **string** | Transport pickup comment                            | [optional] |
| **RealizedDateTime**              | Pointer to **Time**   | Transport pickup realized date time                 | [optional] |
| **RealizedArrivalOnSiteDateTime** | Pointer to **Time**   | Transport pickup realized arrival on site date time | [optional] |
| **RealizedLoadingStartDateTime**  | Pointer to **Time**   | Transport pickup realized loading start date time   | [optional] |
| **RealizedLoadingEndDateTime**    | Pointer to **Time**   | Transport pickup realized loading end date time     | [optional] |

## Methods

### NewApplyPickedUpStatusCommand

`func NewApplyPickedUpStatusCommand(transportUid string, ) *ApplyPickedUpStatusCommand`

NewApplyPickedUpStatusCommand instantiates a new ApplyPickedUpStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyPickedUpStatusCommandWithDefaults

`func NewApplyPickedUpStatusCommandWithDefaults() *ApplyPickedUpStatusCommand`

NewApplyPickedUpStatusCommandWithDefaults instantiates a new ApplyPickedUpStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *ApplyPickedUpStatusCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *ApplyPickedUpStatusCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *ApplyPickedUpStatusCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.

### GetTransportReceipt

`func (o *ApplyPickedUpStatusCommand) GetTransportReceipt() string`

GetTransportReceipt returns the TransportReceipt field if non-nil, zero value otherwise.

### GetTransportReceiptOk

`func (o *ApplyPickedUpStatusCommand) GetTransportReceiptOk() (*string, bool)`

GetTransportReceiptOk returns a tuple with the TransportReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReceipt

`func (o *ApplyPickedUpStatusCommand) SetTransportReceipt(v string)`

SetTransportReceipt sets TransportReceipt field to given value.

### HasTransportReceipt

`func (o *ApplyPickedUpStatusCommand) HasTransportReceipt() bool`

HasTransportReceipt returns a boolean if a field has been set.

### GetPickupSign

`func (o *ApplyPickedUpStatusCommand) GetPickupSign() string`

GetPickupSign returns the PickupSign field if non-nil, zero value otherwise.

### GetPickupSignOk

`func (o *ApplyPickedUpStatusCommand) GetPickupSignOk() (*string, bool)`

GetPickupSignOk returns a tuple with the PickupSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSign

`func (o *ApplyPickedUpStatusCommand) SetPickupSign(v string)`

SetPickupSign sets PickupSign field to given value.

### HasPickupSign

`func (o *ApplyPickedUpStatusCommand) HasPickupSign() bool`

HasPickupSign returns a boolean if a field has been set.

### GetPickupComment

`func (o *ApplyPickedUpStatusCommand) GetPickupComment() string`

GetPickupComment returns the PickupComment field if non-nil, zero value otherwise.

### GetPickupCommentOk

`func (o *ApplyPickedUpStatusCommand) GetPickupCommentOk() (*string, bool)`

GetPickupCommentOk returns a tuple with the PickupComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupComment

`func (o *ApplyPickedUpStatusCommand) SetPickupComment(v string)`

SetPickupComment sets PickupComment field to given value.

### HasPickupComment

`func (o *ApplyPickedUpStatusCommand) HasPickupComment() bool`

HasPickupComment returns a boolean if a field has been set.

### GetRealizedDateTime

`func (o *ApplyPickedUpStatusCommand) GetRealizedDateTime() Time`

GetRealizedDateTime returns the RealizedDateTime field if non-nil, zero value otherwise.

### GetRealizedDateTimeOk

`func (o *ApplyPickedUpStatusCommand) GetRealizedDateTimeOk() (*Time, bool)`

GetRealizedDateTimeOk returns a tuple with the RealizedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedDateTime

`func (o *ApplyPickedUpStatusCommand) SetRealizedDateTime(v Time)`

SetRealizedDateTime sets RealizedDateTime field to given value.

### HasRealizedDateTime

`func (o *ApplyPickedUpStatusCommand) HasRealizedDateTime() bool`

HasRealizedDateTime returns a boolean if a field has been set.

### GetRealizedArrivalOnSiteDateTime

`func (o *ApplyPickedUpStatusCommand) GetRealizedArrivalOnSiteDateTime() Time`

GetRealizedArrivalOnSiteDateTime returns the RealizedArrivalOnSiteDateTime field if non-nil, zero value otherwise.

### GetRealizedArrivalOnSiteDateTimeOk

`func (o *ApplyPickedUpStatusCommand) GetRealizedArrivalOnSiteDateTimeOk() (*Time, bool)`

GetRealizedArrivalOnSiteDateTimeOk returns a tuple with the RealizedArrivalOnSiteDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedArrivalOnSiteDateTime

`func (o *ApplyPickedUpStatusCommand) SetRealizedArrivalOnSiteDateTime(v Time)`

SetRealizedArrivalOnSiteDateTime sets RealizedArrivalOnSiteDateTime field to given value.

### HasRealizedArrivalOnSiteDateTime

`func (o *ApplyPickedUpStatusCommand) HasRealizedArrivalOnSiteDateTime() bool`

HasRealizedArrivalOnSiteDateTime returns a boolean if a field has been set.

### GetRealizedLoadingStartDateTime

`func (o *ApplyPickedUpStatusCommand) GetRealizedLoadingStartDateTime() Time`

GetRealizedLoadingStartDateTime returns the RealizedLoadingStartDateTime field if non-nil, zero value otherwise.

### GetRealizedLoadingStartDateTimeOk

`func (o *ApplyPickedUpStatusCommand) GetRealizedLoadingStartDateTimeOk() (*Time, bool)`

GetRealizedLoadingStartDateTimeOk returns a tuple with the RealizedLoadingStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedLoadingStartDateTime

`func (o *ApplyPickedUpStatusCommand) SetRealizedLoadingStartDateTime(v Time)`

SetRealizedLoadingStartDateTime sets RealizedLoadingStartDateTime field to given value.

### HasRealizedLoadingStartDateTime

`func (o *ApplyPickedUpStatusCommand) HasRealizedLoadingStartDateTime() bool`

HasRealizedLoadingStartDateTime returns a boolean if a field has been set.

### GetRealizedLoadingEndDateTime

`func (o *ApplyPickedUpStatusCommand) GetRealizedLoadingEndDateTime() Time`

GetRealizedLoadingEndDateTime returns the RealizedLoadingEndDateTime field if non-nil, zero value otherwise.

### GetRealizedLoadingEndDateTimeOk

`func (o *ApplyPickedUpStatusCommand) GetRealizedLoadingEndDateTimeOk() (*Time, bool)`

GetRealizedLoadingEndDateTimeOk returns a tuple with the RealizedLoadingEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedLoadingEndDateTime

`func (o *ApplyPickedUpStatusCommand) SetRealizedLoadingEndDateTime(v Time)`

SetRealizedLoadingEndDateTime sets RealizedLoadingEndDateTime field to given value.

### HasRealizedLoadingEndDateTime

`func (o *ApplyPickedUpStatusCommand) HasRealizedLoadingEndDateTime() bool`

HasRealizedLoadingEndDateTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
