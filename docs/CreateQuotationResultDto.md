# CreateQuotationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotationNumber** | Pointer to **int32** |  | [optional] 
**QuotationUid** | Pointer to **string** |  | [optional] 
**Transports** | Pointer to [**[]CreateTransportResultDto**](CreateTransportResultDto.md) |  | [optional] 
**NonBlockingErrors** | Pointer to [**[]NonBlockingErrorDto**](NonBlockingErrorDto.md) |  | [optional] 

## Methods

### NewCreateQuotationResultDto

`func NewCreateQuotationResultDto() *CreateQuotationResultDto`

NewCreateQuotationResultDto instantiates a new CreateQuotationResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuotationResultDtoWithDefaults

`func NewCreateQuotationResultDtoWithDefaults() *CreateQuotationResultDto`

NewCreateQuotationResultDtoWithDefaults instantiates a new CreateQuotationResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotationNumber

`func (o *CreateQuotationResultDto) GetQuotationNumber() int32`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *CreateQuotationResultDto) GetQuotationNumberOk() (*int32, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *CreateQuotationResultDto) SetQuotationNumber(v int32)`

SetQuotationNumber sets QuotationNumber field to given value.

### HasQuotationNumber

`func (o *CreateQuotationResultDto) HasQuotationNumber() bool`

HasQuotationNumber returns a boolean if a field has been set.

### GetQuotationUid

`func (o *CreateQuotationResultDto) GetQuotationUid() string`

GetQuotationUid returns the QuotationUid field if non-nil, zero value otherwise.

### GetQuotationUidOk

`func (o *CreateQuotationResultDto) GetQuotationUidOk() (*string, bool)`

GetQuotationUidOk returns a tuple with the QuotationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationUid

`func (o *CreateQuotationResultDto) SetQuotationUid(v string)`

SetQuotationUid sets QuotationUid field to given value.

### HasQuotationUid

`func (o *CreateQuotationResultDto) HasQuotationUid() bool`

HasQuotationUid returns a boolean if a field has been set.

### GetTransports

`func (o *CreateQuotationResultDto) GetTransports() []CreateTransportResultDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateQuotationResultDto) GetTransportsOk() (*[]CreateTransportResultDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateQuotationResultDto) SetTransports(v []CreateTransportResultDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *CreateQuotationResultDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetNonBlockingErrors

`func (o *CreateQuotationResultDto) GetNonBlockingErrors() []NonBlockingErrorDto`

GetNonBlockingErrors returns the NonBlockingErrors field if non-nil, zero value otherwise.

### GetNonBlockingErrorsOk

`func (o *CreateQuotationResultDto) GetNonBlockingErrorsOk() (*[]NonBlockingErrorDto, bool)`

GetNonBlockingErrorsOk returns a tuple with the NonBlockingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBlockingErrors

`func (o *CreateQuotationResultDto) SetNonBlockingErrors(v []NonBlockingErrorDto)`

SetNonBlockingErrors sets NonBlockingErrors field to given value.

### HasNonBlockingErrors

`func (o *CreateQuotationResultDto) HasNonBlockingErrors() bool`

HasNonBlockingErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


