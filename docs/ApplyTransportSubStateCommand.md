# ApplyTransportSubStateCommand

## Properties

| Name                 | Type                   | Description                                                                                                                                  | Notes      |
| -------------------- | ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **TransportUid**     | **string**             | Mandatory. Transport unique identifier                                                                                                       |
| **SubStateCode**     | **string**             | Mandatory. Sub state code to apply                                                                                                           |
| **Comment**          | Pointer to **string**  | Free text comment                                                                                                                            | [optional] |
| **ActionDateTime**   | Pointer to **Time**    | Date and time of the action, with unspecified time zone. The time corresponds to time zone of the location where the action must be applied. | [optional] |
| **Latitude**         | Pointer to **float64** | Latitude of the position where the action is applied. This number must be between -90 and 90.                                                | [optional] |
| **Longitude**        | Pointer to **float64** | Longitude of the position where the action is applied. This number must be between -180 and 180.                                             | [optional] |
| **PositionDateTime** | Pointer to **Time**    | Date and time corresponding to when the position is set.                                                                                     | [optional] |

## Methods

### NewApplyTransportSubStateCommand

`func NewApplyTransportSubStateCommand(transportUid string, subStateCode string, ) *ApplyTransportSubStateCommand`

NewApplyTransportSubStateCommand instantiates a new ApplyTransportSubStateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTransportSubStateCommandWithDefaults

`func NewApplyTransportSubStateCommandWithDefaults() *ApplyTransportSubStateCommand`

NewApplyTransportSubStateCommandWithDefaults instantiates a new ApplyTransportSubStateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *ApplyTransportSubStateCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *ApplyTransportSubStateCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *ApplyTransportSubStateCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.

### GetSubStateCode

`func (o *ApplyTransportSubStateCommand) GetSubStateCode() string`

GetSubStateCode returns the SubStateCode field if non-nil, zero value otherwise.

### GetSubStateCodeOk

`func (o *ApplyTransportSubStateCommand) GetSubStateCodeOk() (*string, bool)`

GetSubStateCodeOk returns a tuple with the SubStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStateCode

`func (o *ApplyTransportSubStateCommand) SetSubStateCode(v string)`

SetSubStateCode sets SubStateCode field to given value.

### GetComment

`func (o *ApplyTransportSubStateCommand) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApplyTransportSubStateCommand) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApplyTransportSubStateCommand) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApplyTransportSubStateCommand) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetActionDateTime

`func (o *ApplyTransportSubStateCommand) GetActionDateTime() Time`

GetActionDateTime returns the ActionDateTime field if non-nil, zero value otherwise.

### GetActionDateTimeOk

`func (o *ApplyTransportSubStateCommand) GetActionDateTimeOk() (*Time, bool)`

GetActionDateTimeOk returns a tuple with the ActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDateTime

`func (o *ApplyTransportSubStateCommand) SetActionDateTime(v Time)`

SetActionDateTime sets ActionDateTime field to given value.

### HasActionDateTime

`func (o *ApplyTransportSubStateCommand) HasActionDateTime() bool`

HasActionDateTime returns a boolean if a field has been set.

### GetLatitude

`func (o *ApplyTransportSubStateCommand) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ApplyTransportSubStateCommand) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ApplyTransportSubStateCommand) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ApplyTransportSubStateCommand) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *ApplyTransportSubStateCommand) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ApplyTransportSubStateCommand) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ApplyTransportSubStateCommand) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ApplyTransportSubStateCommand) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetPositionDateTime

`func (o *ApplyTransportSubStateCommand) GetPositionDateTime() Time`

GetPositionDateTime returns the PositionDateTime field if non-nil, zero value otherwise.

### GetPositionDateTimeOk

`func (o *ApplyTransportSubStateCommand) GetPositionDateTimeOk() (*Time, bool)`

GetPositionDateTimeOk returns a tuple with the PositionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionDateTime

`func (o *ApplyTransportSubStateCommand) SetPositionDateTime(v Time)`

SetPositionDateTime sets PositionDateTime field to given value.

### HasPositionDateTime

`func (o *ApplyTransportSubStateCommand) HasPositionDateTime() bool`

HasPositionDateTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
