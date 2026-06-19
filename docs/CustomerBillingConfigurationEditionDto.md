# CustomerBillingConfigurationEditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBillable** | Pointer to **bool** | Indicate whether the customer is billable. If set to null, the value defaults to true. | [optional] 
**IsVatExempt** | Pointer to **bool** |  | [optional] 
**PaymentTermCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerBillingConfigurationEditionDto

`func NewCustomerBillingConfigurationEditionDto() *CustomerBillingConfigurationEditionDto`

NewCustomerBillingConfigurationEditionDto instantiates a new CustomerBillingConfigurationEditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingConfigurationEditionDtoWithDefaults

`func NewCustomerBillingConfigurationEditionDtoWithDefaults() *CustomerBillingConfigurationEditionDto`

NewCustomerBillingConfigurationEditionDtoWithDefaults instantiates a new CustomerBillingConfigurationEditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBillable

`func (o *CustomerBillingConfigurationEditionDto) GetIsBillable() bool`

GetIsBillable returns the IsBillable field if non-nil, zero value otherwise.

### GetIsBillableOk

`func (o *CustomerBillingConfigurationEditionDto) GetIsBillableOk() (*bool, bool)`

GetIsBillableOk returns a tuple with the IsBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillable

`func (o *CustomerBillingConfigurationEditionDto) SetIsBillable(v bool)`

SetIsBillable sets IsBillable field to given value.

### HasIsBillable

`func (o *CustomerBillingConfigurationEditionDto) HasIsBillable() bool`

HasIsBillable returns a boolean if a field has been set.

### GetIsVatExempt

`func (o *CustomerBillingConfigurationEditionDto) GetIsVatExempt() bool`

GetIsVatExempt returns the IsVatExempt field if non-nil, zero value otherwise.

### GetIsVatExemptOk

`func (o *CustomerBillingConfigurationEditionDto) GetIsVatExemptOk() (*bool, bool)`

GetIsVatExemptOk returns a tuple with the IsVatExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVatExempt

`func (o *CustomerBillingConfigurationEditionDto) SetIsVatExempt(v bool)`

SetIsVatExempt sets IsVatExempt field to given value.

### HasIsVatExempt

`func (o *CustomerBillingConfigurationEditionDto) HasIsVatExempt() bool`

HasIsVatExempt returns a boolean if a field has been set.

### GetPaymentTermCode

`func (o *CustomerBillingConfigurationEditionDto) GetPaymentTermCode() string`

GetPaymentTermCode returns the PaymentTermCode field if non-nil, zero value otherwise.

### GetPaymentTermCodeOk

`func (o *CustomerBillingConfigurationEditionDto) GetPaymentTermCodeOk() (*string, bool)`

GetPaymentTermCodeOk returns a tuple with the PaymentTermCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermCode

`func (o *CustomerBillingConfigurationEditionDto) SetPaymentTermCode(v string)`

SetPaymentTermCode sets PaymentTermCode field to given value.

### HasPaymentTermCode

`func (o *CustomerBillingConfigurationEditionDto) HasPaymentTermCode() bool`

HasPaymentTermCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


