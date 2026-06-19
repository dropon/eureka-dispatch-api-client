# TransportHistoryLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportHistoryLineId** | Pointer to **int32** |  | [optional] 
**InsertionDateTime** | Pointer to **time.Time** |  | [optional] 
**ActionDateTime** | Pointer to **time.Time** |  | [optional] 
**TransportState** | Pointer to **string** |  | [optional] 
**SubStateCode** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**LatitudeLocation** | Pointer to **float64** |  | [optional] 
**LongitudeLocation** | Pointer to **float64** |  | [optional] 
**LocationDateTime** | Pointer to **time.Time** |  | [optional] 
**Anomalies** | Pointer to [**CappedCollectionDtoTransportHistoryAnomalyDto**](CappedCollectionDtoTransportHistoryAnomalyDto.md) |  | [optional] 

## Methods

### NewTransportHistoryLineDto

`func NewTransportHistoryLineDto() *TransportHistoryLineDto`

NewTransportHistoryLineDto instantiates a new TransportHistoryLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportHistoryLineDtoWithDefaults

`func NewTransportHistoryLineDtoWithDefaults() *TransportHistoryLineDto`

NewTransportHistoryLineDtoWithDefaults instantiates a new TransportHistoryLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportHistoryLineId

`func (o *TransportHistoryLineDto) GetTransportHistoryLineId() int32`

GetTransportHistoryLineId returns the TransportHistoryLineId field if non-nil, zero value otherwise.

### GetTransportHistoryLineIdOk

`func (o *TransportHistoryLineDto) GetTransportHistoryLineIdOk() (*int32, bool)`

GetTransportHistoryLineIdOk returns a tuple with the TransportHistoryLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportHistoryLineId

`func (o *TransportHistoryLineDto) SetTransportHistoryLineId(v int32)`

SetTransportHistoryLineId sets TransportHistoryLineId field to given value.

### HasTransportHistoryLineId

`func (o *TransportHistoryLineDto) HasTransportHistoryLineId() bool`

HasTransportHistoryLineId returns a boolean if a field has been set.

### GetInsertionDateTime

`func (o *TransportHistoryLineDto) GetInsertionDateTime() time.Time`

GetInsertionDateTime returns the InsertionDateTime field if non-nil, zero value otherwise.

### GetInsertionDateTimeOk

`func (o *TransportHistoryLineDto) GetInsertionDateTimeOk() (*time.Time, bool)`

GetInsertionDateTimeOk returns a tuple with the InsertionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertionDateTime

`func (o *TransportHistoryLineDto) SetInsertionDateTime(v time.Time)`

SetInsertionDateTime sets InsertionDateTime field to given value.

### HasInsertionDateTime

`func (o *TransportHistoryLineDto) HasInsertionDateTime() bool`

HasInsertionDateTime returns a boolean if a field has been set.

### GetActionDateTime

`func (o *TransportHistoryLineDto) GetActionDateTime() time.Time`

GetActionDateTime returns the ActionDateTime field if non-nil, zero value otherwise.

### GetActionDateTimeOk

`func (o *TransportHistoryLineDto) GetActionDateTimeOk() (*time.Time, bool)`

GetActionDateTimeOk returns a tuple with the ActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDateTime

`func (o *TransportHistoryLineDto) SetActionDateTime(v time.Time)`

SetActionDateTime sets ActionDateTime field to given value.

### HasActionDateTime

`func (o *TransportHistoryLineDto) HasActionDateTime() bool`

HasActionDateTime returns a boolean if a field has been set.

### GetTransportState

`func (o *TransportHistoryLineDto) GetTransportState() string`

GetTransportState returns the TransportState field if non-nil, zero value otherwise.

### GetTransportStateOk

`func (o *TransportHistoryLineDto) GetTransportStateOk() (*string, bool)`

GetTransportStateOk returns a tuple with the TransportState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportState

`func (o *TransportHistoryLineDto) SetTransportState(v string)`

SetTransportState sets TransportState field to given value.

### HasTransportState

`func (o *TransportHistoryLineDto) HasTransportState() bool`

HasTransportState returns a boolean if a field has been set.

### GetSubStateCode

`func (o *TransportHistoryLineDto) GetSubStateCode() string`

GetSubStateCode returns the SubStateCode field if non-nil, zero value otherwise.

### GetSubStateCodeOk

`func (o *TransportHistoryLineDto) GetSubStateCodeOk() (*string, bool)`

GetSubStateCodeOk returns a tuple with the SubStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStateCode

`func (o *TransportHistoryLineDto) SetSubStateCode(v string)`

SetSubStateCode sets SubStateCode field to given value.

### HasSubStateCode

`func (o *TransportHistoryLineDto) HasSubStateCode() bool`

HasSubStateCode returns a boolean if a field has been set.

### GetComment

`func (o *TransportHistoryLineDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransportHistoryLineDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransportHistoryLineDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransportHistoryLineDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLatitudeLocation

`func (o *TransportHistoryLineDto) GetLatitudeLocation() float64`

GetLatitudeLocation returns the LatitudeLocation field if non-nil, zero value otherwise.

### GetLatitudeLocationOk

`func (o *TransportHistoryLineDto) GetLatitudeLocationOk() (*float64, bool)`

GetLatitudeLocationOk returns a tuple with the LatitudeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitudeLocation

`func (o *TransportHistoryLineDto) SetLatitudeLocation(v float64)`

SetLatitudeLocation sets LatitudeLocation field to given value.

### HasLatitudeLocation

`func (o *TransportHistoryLineDto) HasLatitudeLocation() bool`

HasLatitudeLocation returns a boolean if a field has been set.

### GetLongitudeLocation

`func (o *TransportHistoryLineDto) GetLongitudeLocation() float64`

GetLongitudeLocation returns the LongitudeLocation field if non-nil, zero value otherwise.

### GetLongitudeLocationOk

`func (o *TransportHistoryLineDto) GetLongitudeLocationOk() (*float64, bool)`

GetLongitudeLocationOk returns a tuple with the LongitudeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitudeLocation

`func (o *TransportHistoryLineDto) SetLongitudeLocation(v float64)`

SetLongitudeLocation sets LongitudeLocation field to given value.

### HasLongitudeLocation

`func (o *TransportHistoryLineDto) HasLongitudeLocation() bool`

HasLongitudeLocation returns a boolean if a field has been set.

### GetLocationDateTime

`func (o *TransportHistoryLineDto) GetLocationDateTime() time.Time`

GetLocationDateTime returns the LocationDateTime field if non-nil, zero value otherwise.

### GetLocationDateTimeOk

`func (o *TransportHistoryLineDto) GetLocationDateTimeOk() (*time.Time, bool)`

GetLocationDateTimeOk returns a tuple with the LocationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDateTime

`func (o *TransportHistoryLineDto) SetLocationDateTime(v time.Time)`

SetLocationDateTime sets LocationDateTime field to given value.

### HasLocationDateTime

`func (o *TransportHistoryLineDto) HasLocationDateTime() bool`

HasLocationDateTime returns a boolean if a field has been set.

### GetAnomalies

`func (o *TransportHistoryLineDto) GetAnomalies() CappedCollectionDtoTransportHistoryAnomalyDto`

GetAnomalies returns the Anomalies field if non-nil, zero value otherwise.

### GetAnomaliesOk

`func (o *TransportHistoryLineDto) GetAnomaliesOk() (*CappedCollectionDtoTransportHistoryAnomalyDto, bool)`

GetAnomaliesOk returns a tuple with the Anomalies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalies

`func (o *TransportHistoryLineDto) SetAnomalies(v CappedCollectionDtoTransportHistoryAnomalyDto)`

SetAnomalies sets Anomalies field to given value.

### HasAnomalies

`func (o *TransportHistoryLineDto) HasAnomalies() bool`

HasAnomalies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


