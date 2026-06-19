# CustomerBillingDetailsDtoBankDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankCode** | Pointer to **string** |  | [optional] 
**BranchCode** | Pointer to **string** | Branch code (French: Guichet) | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**BankKey** | Pointer to **string** |  | [optional] 
**BankDomiciliation** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**BicCode** | Pointer to **string** |  | [optional] 
**SepaCreditorIdentifier** | Pointer to **string** | SEPA creditor identifier (French: ICS) | [optional] 
**SepaMandateReference** | Pointer to **string** | SEPA Unique Mandate Reference (French: RUM) | [optional] 

## Methods

### NewCustomerBillingDetailsDtoBankDetailsDto

`func NewCustomerBillingDetailsDtoBankDetailsDto() *CustomerBillingDetailsDtoBankDetailsDto`

NewCustomerBillingDetailsDtoBankDetailsDto instantiates a new CustomerBillingDetailsDtoBankDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingDetailsDtoBankDetailsDtoWithDefaults

`func NewCustomerBillingDetailsDtoBankDetailsDtoWithDefaults() *CustomerBillingDetailsDtoBankDetailsDto`

NewCustomerBillingDetailsDtoBankDetailsDtoWithDefaults instantiates a new CustomerBillingDetailsDtoBankDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBranchCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetAccountNumber

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBankKey

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankKey() string`

GetBankKey returns the BankKey field if non-nil, zero value otherwise.

### GetBankKeyOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankKeyOk() (*string, bool)`

GetBankKeyOk returns a tuple with the BankKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankKey

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetBankKey(v string)`

SetBankKey sets BankKey field to given value.

### HasBankKey

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasBankKey() bool`

HasBankKey returns a boolean if a field has been set.

### GetBankDomiciliation

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankDomiciliation() string`

GetBankDomiciliation returns the BankDomiciliation field if non-nil, zero value otherwise.

### GetBankDomiciliationOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBankDomiciliationOk() (*string, bool)`

GetBankDomiciliationOk returns a tuple with the BankDomiciliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankDomiciliation

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetBankDomiciliation(v string)`

SetBankDomiciliation sets BankDomiciliation field to given value.

### HasBankDomiciliation

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasBankDomiciliation() bool`

HasBankDomiciliation returns a boolean if a field has been set.

### GetIban

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetBicCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBicCode() string`

GetBicCode returns the BicCode field if non-nil, zero value otherwise.

### GetBicCodeOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetBicCodeOk() (*string, bool)`

GetBicCodeOk returns a tuple with the BicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBicCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetBicCode(v string)`

SetBicCode sets BicCode field to given value.

### HasBicCode

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasBicCode() bool`

HasBicCode returns a boolean if a field has been set.

### GetSepaCreditorIdentifier

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetSepaCreditorIdentifier() string`

GetSepaCreditorIdentifier returns the SepaCreditorIdentifier field if non-nil, zero value otherwise.

### GetSepaCreditorIdentifierOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetSepaCreditorIdentifierOk() (*string, bool)`

GetSepaCreditorIdentifierOk returns a tuple with the SepaCreditorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCreditorIdentifier

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetSepaCreditorIdentifier(v string)`

SetSepaCreditorIdentifier sets SepaCreditorIdentifier field to given value.

### HasSepaCreditorIdentifier

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasSepaCreditorIdentifier() bool`

HasSepaCreditorIdentifier returns a boolean if a field has been set.

### GetSepaMandateReference

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetSepaMandateReference() string`

GetSepaMandateReference returns the SepaMandateReference field if non-nil, zero value otherwise.

### GetSepaMandateReferenceOk

`func (o *CustomerBillingDetailsDtoBankDetailsDto) GetSepaMandateReferenceOk() (*string, bool)`

GetSepaMandateReferenceOk returns a tuple with the SepaMandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaMandateReference

`func (o *CustomerBillingDetailsDtoBankDetailsDto) SetSepaMandateReference(v string)`

SetSepaMandateReference sets SepaMandateReference field to given value.

### HasSepaMandateReference

`func (o *CustomerBillingDetailsDtoBankDetailsDto) HasSepaMandateReference() bool`

HasSepaMandateReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


