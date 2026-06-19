# CustomerIncludedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationAddress** | Pointer to [**AddressDto**](AddressDto.md) |  | [optional] 
**PaymentTerm** | Pointer to [**PaymentTermDto**](PaymentTermDto.md) |  | [optional] 

## Methods

### NewCustomerIncludedDto

`func NewCustomerIncludedDto() *CustomerIncludedDto`

NewCustomerIncludedDto instantiates a new CustomerIncludedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerIncludedDtoWithDefaults

`func NewCustomerIncludedDtoWithDefaults() *CustomerIncludedDto`

NewCustomerIncludedDtoWithDefaults instantiates a new CustomerIncludedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationAddress

`func (o *CustomerIncludedDto) GetOperationAddress() AddressDto`

GetOperationAddress returns the OperationAddress field if non-nil, zero value otherwise.

### GetOperationAddressOk

`func (o *CustomerIncludedDto) GetOperationAddressOk() (*AddressDto, bool)`

GetOperationAddressOk returns a tuple with the OperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddress

`func (o *CustomerIncludedDto) SetOperationAddress(v AddressDto)`

SetOperationAddress sets OperationAddress field to given value.

### HasOperationAddress

`func (o *CustomerIncludedDto) HasOperationAddress() bool`

HasOperationAddress returns a boolean if a field has been set.

### GetPaymentTerm

`func (o *CustomerIncludedDto) GetPaymentTerm() PaymentTermDto`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *CustomerIncludedDto) GetPaymentTermOk() (*PaymentTermDto, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *CustomerIncludedDto) SetPaymentTerm(v PaymentTermDto)`

SetPaymentTerm sets PaymentTerm field to given value.

### HasPaymentTerm

`func (o *CustomerIncludedDto) HasPaymentTerm() bool`

HasPaymentTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


