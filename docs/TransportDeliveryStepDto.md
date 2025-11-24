# TransportDeliveryStepDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepDate** | Pointer to [**TransportDeliveryDateTimeSlotDto**](TransportDeliveryDateTimeSlotDto.md) |  | [optional] 
**StepId** | Pointer to **int32** |  | [optional] 
**Sign** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**StepAddress** | Pointer to [**TransportStepAddressDto**](TransportStepAddressDto.md) |  | [optional] 
**StepContact** | Pointer to [**TransportStepContactDto**](TransportStepContactDto.md) |  | [optional] 
**StepConstraints** | Pointer to [**TransportStepConstraintsDto**](TransportStepConstraintsDto.md) |  | [optional] 

## Methods

### NewTransportDeliveryStepDto

`func NewTransportDeliveryStepDto() *TransportDeliveryStepDto`

NewTransportDeliveryStepDto instantiates a new TransportDeliveryStepDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportDeliveryStepDtoWithDefaults

`func NewTransportDeliveryStepDtoWithDefaults() *TransportDeliveryStepDto`

NewTransportDeliveryStepDtoWithDefaults instantiates a new TransportDeliveryStepDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepDate

`func (o *TransportDeliveryStepDto) GetStepDate() TransportDeliveryDateTimeSlotDto`

GetStepDate returns the StepDate field if non-nil, zero value otherwise.

### GetStepDateOk

`func (o *TransportDeliveryStepDto) GetStepDateOk() (*TransportDeliveryDateTimeSlotDto, bool)`

GetStepDateOk returns a tuple with the StepDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepDate

`func (o *TransportDeliveryStepDto) SetStepDate(v TransportDeliveryDateTimeSlotDto)`

SetStepDate sets StepDate field to given value.

### HasStepDate

`func (o *TransportDeliveryStepDto) HasStepDate() bool`

HasStepDate returns a boolean if a field has been set.

### GetStepId

`func (o *TransportDeliveryStepDto) GetStepId() int32`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *TransportDeliveryStepDto) GetStepIdOk() (*int32, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *TransportDeliveryStepDto) SetStepId(v int32)`

SetStepId sets StepId field to given value.

### HasStepId

`func (o *TransportDeliveryStepDto) HasStepId() bool`

HasStepId returns a boolean if a field has been set.

### GetSign

`func (o *TransportDeliveryStepDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *TransportDeliveryStepDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *TransportDeliveryStepDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *TransportDeliveryStepDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetComment

`func (o *TransportDeliveryStepDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransportDeliveryStepDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransportDeliveryStepDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransportDeliveryStepDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStepAddress

`func (o *TransportDeliveryStepDto) GetStepAddress() TransportStepAddressDto`

GetStepAddress returns the StepAddress field if non-nil, zero value otherwise.

### GetStepAddressOk

`func (o *TransportDeliveryStepDto) GetStepAddressOk() (*TransportStepAddressDto, bool)`

GetStepAddressOk returns a tuple with the StepAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepAddress

`func (o *TransportDeliveryStepDto) SetStepAddress(v TransportStepAddressDto)`

SetStepAddress sets StepAddress field to given value.

### HasStepAddress

`func (o *TransportDeliveryStepDto) HasStepAddress() bool`

HasStepAddress returns a boolean if a field has been set.

### GetStepContact

`func (o *TransportDeliveryStepDto) GetStepContact() TransportStepContactDto`

GetStepContact returns the StepContact field if non-nil, zero value otherwise.

### GetStepContactOk

`func (o *TransportDeliveryStepDto) GetStepContactOk() (*TransportStepContactDto, bool)`

GetStepContactOk returns a tuple with the StepContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepContact

`func (o *TransportDeliveryStepDto) SetStepContact(v TransportStepContactDto)`

SetStepContact sets StepContact field to given value.

### HasStepContact

`func (o *TransportDeliveryStepDto) HasStepContact() bool`

HasStepContact returns a boolean if a field has been set.

### GetStepConstraints

`func (o *TransportDeliveryStepDto) GetStepConstraints() TransportStepConstraintsDto`

GetStepConstraints returns the StepConstraints field if non-nil, zero value otherwise.

### GetStepConstraintsOk

`func (o *TransportDeliveryStepDto) GetStepConstraintsOk() (*TransportStepConstraintsDto, bool)`

GetStepConstraintsOk returns a tuple with the StepConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepConstraints

`func (o *TransportDeliveryStepDto) SetStepConstraints(v TransportStepConstraintsDto)`

SetStepConstraints sets StepConstraints field to given value.

### HasStepConstraints

`func (o *TransportDeliveryStepDto) HasStepConstraints() bool`

HasStepConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


