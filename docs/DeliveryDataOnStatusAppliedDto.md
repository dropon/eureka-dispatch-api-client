# DeliveryDataOnStatusAppliedDto

## Properties

| Name                               | Type                  | Description                                           | Notes      |
| ---------------------------------- | --------------------- | ----------------------------------------------------- | ---------- |
| **Sign**                           | Pointer to **string** | Transport delivery sign                               | [optional] |
| **Comment**                        | Pointer to **string** | Transport delivery comment                            | [optional] |
| **RealizedArrivalOnSiteDateTime**  | Pointer to **Time**   | Transport delivery realized arrival on site date time | [optional] |
| **RealizedUnloadingStartDateTime** | Pointer to **Time**   | Transport delivery realized unloading start date time | [optional] |
| **RealizedUnloadingEndDateTime**   | Pointer to **Time**   | Transport delivery realized unloading end date time   | [optional] |

## Methods

### NewDeliveryDataOnStatusAppliedDto

`func NewDeliveryDataOnStatusAppliedDto() *DeliveryDataOnStatusAppliedDto`

NewDeliveryDataOnStatusAppliedDto instantiates a new DeliveryDataOnStatusAppliedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryDataOnStatusAppliedDtoWithDefaults

`func NewDeliveryDataOnStatusAppliedDtoWithDefaults() *DeliveryDataOnStatusAppliedDto`

NewDeliveryDataOnStatusAppliedDtoWithDefaults instantiates a new DeliveryDataOnStatusAppliedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSign

`func (o *DeliveryDataOnStatusAppliedDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *DeliveryDataOnStatusAppliedDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *DeliveryDataOnStatusAppliedDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *DeliveryDataOnStatusAppliedDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetComment

`func (o *DeliveryDataOnStatusAppliedDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DeliveryDataOnStatusAppliedDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DeliveryDataOnStatusAppliedDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DeliveryDataOnStatusAppliedDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRealizedArrivalOnSiteDateTime

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedArrivalOnSiteDateTime() Time`

GetRealizedArrivalOnSiteDateTime returns the RealizedArrivalOnSiteDateTime field if non-nil, zero value otherwise.

### GetRealizedArrivalOnSiteDateTimeOk

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedArrivalOnSiteDateTimeOk() (*Time, bool)`

GetRealizedArrivalOnSiteDateTimeOk returns a tuple with the RealizedArrivalOnSiteDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedArrivalOnSiteDateTime

`func (o *DeliveryDataOnStatusAppliedDto) SetRealizedArrivalOnSiteDateTime(v Time)`

SetRealizedArrivalOnSiteDateTime sets RealizedArrivalOnSiteDateTime field to given value.

### HasRealizedArrivalOnSiteDateTime

`func (o *DeliveryDataOnStatusAppliedDto) HasRealizedArrivalOnSiteDateTime() bool`

HasRealizedArrivalOnSiteDateTime returns a boolean if a field has been set.

### GetRealizedUnloadingStartDateTime

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedUnloadingStartDateTime() Time`

GetRealizedUnloadingStartDateTime returns the RealizedUnloadingStartDateTime field if non-nil, zero value otherwise.

### GetRealizedUnloadingStartDateTimeOk

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedUnloadingStartDateTimeOk() (*Time, bool)`

GetRealizedUnloadingStartDateTimeOk returns a tuple with the RealizedUnloadingStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedUnloadingStartDateTime

`func (o *DeliveryDataOnStatusAppliedDto) SetRealizedUnloadingStartDateTime(v Time)`

SetRealizedUnloadingStartDateTime sets RealizedUnloadingStartDateTime field to given value.

### HasRealizedUnloadingStartDateTime

`func (o *DeliveryDataOnStatusAppliedDto) HasRealizedUnloadingStartDateTime() bool`

HasRealizedUnloadingStartDateTime returns a boolean if a field has been set.

### GetRealizedUnloadingEndDateTime

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedUnloadingEndDateTime() Time`

GetRealizedUnloadingEndDateTime returns the RealizedUnloadingEndDateTime field if non-nil, zero value otherwise.

### GetRealizedUnloadingEndDateTimeOk

`func (o *DeliveryDataOnStatusAppliedDto) GetRealizedUnloadingEndDateTimeOk() (*Time, bool)`

GetRealizedUnloadingEndDateTimeOk returns a tuple with the RealizedUnloadingEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedUnloadingEndDateTime

`func (o *DeliveryDataOnStatusAppliedDto) SetRealizedUnloadingEndDateTime(v Time)`

SetRealizedUnloadingEndDateTime sets RealizedUnloadingEndDateTime field to given value.

### HasRealizedUnloadingEndDateTime

`func (o *DeliveryDataOnStatusAppliedDto) HasRealizedUnloadingEndDateTime() bool`

HasRealizedUnloadingEndDateTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
