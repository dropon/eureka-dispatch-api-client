# CappedCollectionDtoTransportHistoryAnomalyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TransportHistoryAnomalyDto**](TransportHistoryAnomalyDto.md) | List of items included in the capped collection | [optional] 
**HasAdditionalResult** | Pointer to **bool** | Indicates if the included data was partially returned because capped collection size is reached.  In this case, additional result can be retrieved from another route using pagination. | [optional] 

## Methods

### NewCappedCollectionDtoTransportHistoryAnomalyDto

`func NewCappedCollectionDtoTransportHistoryAnomalyDto() *CappedCollectionDtoTransportHistoryAnomalyDto`

NewCappedCollectionDtoTransportHistoryAnomalyDto instantiates a new CappedCollectionDtoTransportHistoryAnomalyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCappedCollectionDtoTransportHistoryAnomalyDtoWithDefaults

`func NewCappedCollectionDtoTransportHistoryAnomalyDtoWithDefaults() *CappedCollectionDtoTransportHistoryAnomalyDto`

NewCappedCollectionDtoTransportHistoryAnomalyDtoWithDefaults instantiates a new CappedCollectionDtoTransportHistoryAnomalyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) GetData() []TransportHistoryAnomalyDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) GetDataOk() (*[]TransportHistoryAnomalyDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) SetData(v []TransportHistoryAnomalyDto)`

SetData sets Data field to given value.

### HasData

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasAdditionalResult

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) GetHasAdditionalResult() bool`

GetHasAdditionalResult returns the HasAdditionalResult field if non-nil, zero value otherwise.

### GetHasAdditionalResultOk

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) GetHasAdditionalResultOk() (*bool, bool)`

GetHasAdditionalResultOk returns a tuple with the HasAdditionalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalResult

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) SetHasAdditionalResult(v bool)`

SetHasAdditionalResult sets HasAdditionalResult field to given value.

### HasHasAdditionalResult

`func (o *CappedCollectionDtoTransportHistoryAnomalyDto) HasHasAdditionalResult() bool`

HasHasAdditionalResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


