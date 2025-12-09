# PickupDataOnStatusAppliedDto

## Properties

| Name                              | Type                  | Description                                         | Notes      |
| --------------------------------- | --------------------- | --------------------------------------------------- | ---------- |
| **Sign**                          | Pointer to **string** | Transport pickup sign                               | [optional] |
| **Comment**                       | Pointer to **string** | Transport pickup comment                            | [optional] |
| **RealizedArrivalOnSiteDateTime** | Pointer to **Time**   | Transport pickup realized arrival on site date time | [optional] |
| **RealizedLoadingStartDateTime**  | Pointer to **Time**   | Transport pickup realized loading start date time   | [optional] |
| **RealizedLoadingEndDateTime**    | Pointer to **Time**   | Transport pickup realized loading end date time     | [optional] |

## Methods

### NewPickupDataOnStatusAppliedDto

`func NewPickupDataOnStatusAppliedDto() *PickupDataOnStatusAppliedDto`

NewPickupDataOnStatusAppliedDto instantiates a new PickupDataOnStatusAppliedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupDataOnStatusAppliedDtoWithDefaults

`func NewPickupDataOnStatusAppliedDtoWithDefaults() *PickupDataOnStatusAppliedDto`

NewPickupDataOnStatusAppliedDtoWithDefaults instantiates a new PickupDataOnStatusAppliedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSign

`func (o *PickupDataOnStatusAppliedDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *PickupDataOnStatusAppliedDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *PickupDataOnStatusAppliedDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *PickupDataOnStatusAppliedDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetComment

`func (o *PickupDataOnStatusAppliedDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PickupDataOnStatusAppliedDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PickupDataOnStatusAppliedDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PickupDataOnStatusAppliedDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRealizedArrivalOnSiteDateTime

`func (o *PickupDataOnStatusAppliedDto) GetRealizedArrivalOnSiteDateTime() Time`

GetRealizedArrivalOnSiteDateTime returns the RealizedArrivalOnSiteDateTime field if non-nil, zero value otherwise.

### GetRealizedArrivalOnSiteDateTimeOk

`func (o *PickupDataOnStatusAppliedDto) GetRealizedArrivalOnSiteDateTimeOk() (*Time, bool)`

GetRealizedArrivalOnSiteDateTimeOk returns a tuple with the RealizedArrivalOnSiteDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedArrivalOnSiteDateTime

`func (o *PickupDataOnStatusAppliedDto) SetRealizedArrivalOnSiteDateTime(v Time)`

SetRealizedArrivalOnSiteDateTime sets RealizedArrivalOnSiteDateTime field to given value.

### HasRealizedArrivalOnSiteDateTime

`func (o *PickupDataOnStatusAppliedDto) HasRealizedArrivalOnSiteDateTime() bool`

HasRealizedArrivalOnSiteDateTime returns a boolean if a field has been set.

### GetRealizedLoadingStartDateTime

`func (o *PickupDataOnStatusAppliedDto) GetRealizedLoadingStartDateTime() Time`

GetRealizedLoadingStartDateTime returns the RealizedLoadingStartDateTime field if non-nil, zero value otherwise.

### GetRealizedLoadingStartDateTimeOk

`func (o *PickupDataOnStatusAppliedDto) GetRealizedLoadingStartDateTimeOk() (*Time, bool)`

GetRealizedLoadingStartDateTimeOk returns a tuple with the RealizedLoadingStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedLoadingStartDateTime

`func (o *PickupDataOnStatusAppliedDto) SetRealizedLoadingStartDateTime(v Time)`

SetRealizedLoadingStartDateTime sets RealizedLoadingStartDateTime field to given value.

### HasRealizedLoadingStartDateTime

`func (o *PickupDataOnStatusAppliedDto) HasRealizedLoadingStartDateTime() bool`

HasRealizedLoadingStartDateTime returns a boolean if a field has been set.

### GetRealizedLoadingEndDateTime

`func (o *PickupDataOnStatusAppliedDto) GetRealizedLoadingEndDateTime() Time`

GetRealizedLoadingEndDateTime returns the RealizedLoadingEndDateTime field if non-nil, zero value otherwise.

### GetRealizedLoadingEndDateTimeOk

`func (o *PickupDataOnStatusAppliedDto) GetRealizedLoadingEndDateTimeOk() (*Time, bool)`

GetRealizedLoadingEndDateTimeOk returns a tuple with the RealizedLoadingEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedLoadingEndDateTime

`func (o *PickupDataOnStatusAppliedDto) SetRealizedLoadingEndDateTime(v Time)`

SetRealizedLoadingEndDateTime sets RealizedLoadingEndDateTime field to given value.

### HasRealizedLoadingEndDateTime

`func (o *PickupDataOnStatusAppliedDto) HasRealizedLoadingEndDateTime() bool`

HasRealizedLoadingEndDateTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
