# QuotationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Quotation&#39;s internal unique identifier | [optional] 
**QuotationNumber** | Pointer to **int32** | Quotation&#39;s public identifier | [optional] 
**TrackId** | Pointer to **string** | Quotation&#39;s tracking number | [optional] 
**IsUniqueOrder** | Pointer to **bool** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**Transports** | Pointer to [**CappedCollectionDtoTransportDto**](CappedCollectionDtoTransportDto.md) |  | [optional] 

## Methods

### NewQuotationDto

`func NewQuotationDto() *QuotationDto`

NewQuotationDto instantiates a new QuotationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationDtoWithDefaults

`func NewQuotationDtoWithDefaults() *QuotationDto`

NewQuotationDtoWithDefaults instantiates a new QuotationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *QuotationDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *QuotationDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *QuotationDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *QuotationDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetQuotationNumber

`func (o *QuotationDto) GetQuotationNumber() int32`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *QuotationDto) GetQuotationNumberOk() (*int32, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *QuotationDto) SetQuotationNumber(v int32)`

SetQuotationNumber sets QuotationNumber field to given value.

### HasQuotationNumber

`func (o *QuotationDto) HasQuotationNumber() bool`

HasQuotationNumber returns a boolean if a field has been set.

### GetTrackId

`func (o *QuotationDto) GetTrackId() string`

GetTrackId returns the TrackId field if non-nil, zero value otherwise.

### GetTrackIdOk

`func (o *QuotationDto) GetTrackIdOk() (*string, bool)`

GetTrackIdOk returns a tuple with the TrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackId

`func (o *QuotationDto) SetTrackId(v string)`

SetTrackId sets TrackId field to given value.

### HasTrackId

`func (o *QuotationDto) HasTrackId() bool`

HasTrackId returns a boolean if a field has been set.

### GetIsUniqueOrder

`func (o *QuotationDto) GetIsUniqueOrder() bool`

GetIsUniqueOrder returns the IsUniqueOrder field if non-nil, zero value otherwise.

### GetIsUniqueOrderOk

`func (o *QuotationDto) GetIsUniqueOrderOk() (*bool, bool)`

GetIsUniqueOrderOk returns a tuple with the IsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniqueOrder

`func (o *QuotationDto) SetIsUniqueOrder(v bool)`

SetIsUniqueOrder sets IsUniqueOrder field to given value.

### HasIsUniqueOrder

`func (o *QuotationDto) HasIsUniqueOrder() bool`

HasIsUniqueOrder returns a boolean if a field has been set.

### GetIsArchived

`func (o *QuotationDto) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *QuotationDto) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *QuotationDto) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *QuotationDto) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsFinalized

`func (o *QuotationDto) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *QuotationDto) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *QuotationDto) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *QuotationDto) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetTransports

`func (o *QuotationDto) GetTransports() CappedCollectionDtoTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *QuotationDto) GetTransportsOk() (*CappedCollectionDtoTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *QuotationDto) SetTransports(v CappedCollectionDtoTransportDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *QuotationDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


