# QuotationEntryDryRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUniqueOrder** | Pointer to **bool** |  | [optional] 
**IsFinalized** | Pointer to **bool** |  | [optional] 
**Transports** | Pointer to [**[]TransportEntryDryRunDto**](TransportEntryDryRunDto.md) |  | [optional] 
**NonBlockingErrors** | Pointer to [**[]NonBlockingErrorDto**](NonBlockingErrorDto.md) |  | [optional] 

## Methods

### NewQuotationEntryDryRunDto

`func NewQuotationEntryDryRunDto() *QuotationEntryDryRunDto`

NewQuotationEntryDryRunDto instantiates a new QuotationEntryDryRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationEntryDryRunDtoWithDefaults

`func NewQuotationEntryDryRunDtoWithDefaults() *QuotationEntryDryRunDto`

NewQuotationEntryDryRunDtoWithDefaults instantiates a new QuotationEntryDryRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUniqueOrder

`func (o *QuotationEntryDryRunDto) GetIsUniqueOrder() bool`

GetIsUniqueOrder returns the IsUniqueOrder field if non-nil, zero value otherwise.

### GetIsUniqueOrderOk

`func (o *QuotationEntryDryRunDto) GetIsUniqueOrderOk() (*bool, bool)`

GetIsUniqueOrderOk returns a tuple with the IsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUniqueOrder

`func (o *QuotationEntryDryRunDto) SetIsUniqueOrder(v bool)`

SetIsUniqueOrder sets IsUniqueOrder field to given value.

### HasIsUniqueOrder

`func (o *QuotationEntryDryRunDto) HasIsUniqueOrder() bool`

HasIsUniqueOrder returns a boolean if a field has been set.

### GetIsFinalized

`func (o *QuotationEntryDryRunDto) GetIsFinalized() bool`

GetIsFinalized returns the IsFinalized field if non-nil, zero value otherwise.

### GetIsFinalizedOk

`func (o *QuotationEntryDryRunDto) GetIsFinalizedOk() (*bool, bool)`

GetIsFinalizedOk returns a tuple with the IsFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalized

`func (o *QuotationEntryDryRunDto) SetIsFinalized(v bool)`

SetIsFinalized sets IsFinalized field to given value.

### HasIsFinalized

`func (o *QuotationEntryDryRunDto) HasIsFinalized() bool`

HasIsFinalized returns a boolean if a field has been set.

### GetTransports

`func (o *QuotationEntryDryRunDto) GetTransports() []TransportEntryDryRunDto`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *QuotationEntryDryRunDto) GetTransportsOk() (*[]TransportEntryDryRunDto, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *QuotationEntryDryRunDto) SetTransports(v []TransportEntryDryRunDto)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *QuotationEntryDryRunDto) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### GetNonBlockingErrors

`func (o *QuotationEntryDryRunDto) GetNonBlockingErrors() []NonBlockingErrorDto`

GetNonBlockingErrors returns the NonBlockingErrors field if non-nil, zero value otherwise.

### GetNonBlockingErrorsOk

`func (o *QuotationEntryDryRunDto) GetNonBlockingErrorsOk() (*[]NonBlockingErrorDto, bool)`

GetNonBlockingErrorsOk returns a tuple with the NonBlockingErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBlockingErrors

`func (o *QuotationEntryDryRunDto) SetNonBlockingErrors(v []NonBlockingErrorDto)`

SetNonBlockingErrors sets NonBlockingErrors field to given value.

### HasNonBlockingErrors

`func (o *QuotationEntryDryRunDto) HasNonBlockingErrors() bool`

HasNonBlockingErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


