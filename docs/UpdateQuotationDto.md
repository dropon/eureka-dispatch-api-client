# UpdateQuotationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | Pointer to [**[]MergeableTransportDto**](MergeableTransportDto.md) |  | [optional] 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**ApiBehavior** | Pointer to [**UpdateQuotationDtoApiBehaviorDto**](UpdateQuotationDtoApiBehaviorDto.md) |  | [optional] 

## Methods

### NewUpdateQuotationDto

`func NewUpdateQuotationDto() *UpdateQuotationDto`

NewUpdateQuotationDto instantiates a new UpdateQuotationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQuotationDtoWithDefaults

`func NewUpdateQuotationDtoWithDefaults() *UpdateQuotationDto`

NewUpdateQuotationDtoWithDefaults instantiates a new UpdateQuotationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *UpdateQuotationDto) GetTransports() []MergeableTransportDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *UpdateQuotationDto) GetTransportsOk() (*[]MergeableTransportDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *UpdateQuotationDto) SetTransports(v []MergeableTransportDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *UpdateQuotationDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetIsFinalized

`func (o *UpdateQuotationDto) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *UpdateQuotationDto) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *UpdateQuotationDto) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *UpdateQuotationDto) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetApiBehavior

`func (o *UpdateQuotationDto) GetApiBehavior() UpdateQuotationDtoApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *UpdateQuotationDto) GetApiBehaviorOk() (*UpdateQuotationDtoApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *UpdateQuotationDto) SetApiBehavior(v UpdateQuotationDtoApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *UpdateQuotationDto) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


