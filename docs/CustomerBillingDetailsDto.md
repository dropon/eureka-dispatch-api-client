# CustomerBillingDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankDetails** | Pointer to [**CustomerBillingDetailsDtoBankDetailsDto**](CustomerBillingDetailsDtoBankDetailsDto.md) |  | [optional] 
**AdministrativeInfo** | Pointer to [**CustomerBillingDetailsDtoAdministrativeInfoDto**](CustomerBillingDetailsDtoAdministrativeInfoDto.md) |  | [optional] 

## Methods

### NewCustomerBillingDetailsDto

`func NewCustomerBillingDetailsDto() *CustomerBillingDetailsDto`

NewCustomerBillingDetailsDto instantiates a new CustomerBillingDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingDetailsDtoWithDefaults

`func NewCustomerBillingDetailsDtoWithDefaults() *CustomerBillingDetailsDto`

NewCustomerBillingDetailsDtoWithDefaults instantiates a new CustomerBillingDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankDetails

`func (o *CustomerBillingDetailsDto) GetBankDetails() CustomerBillingDetailsDtoBankDetailsDto`

GetBankDetails returns the BankDetails field if non-nil, zero value otherwise.

### GetBankDetailsOk

`func (o *CustomerBillingDetailsDto) GetBankDetailsOk() (*CustomerBillingDetailsDtoBankDetailsDto, bool)`

GetBankDetailsOk returns a tuple with the BankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankDetails

`func (o *CustomerBillingDetailsDto) SetBankDetails(v CustomerBillingDetailsDtoBankDetailsDto)`

SetBankDetails sets BankDetails field to given value.

### HasBankDetails

`func (o *CustomerBillingDetailsDto) HasBankDetails() bool`

HasBankDetails returns a boolean if a field has been set.

### GetAdministrativeInfo

`func (o *CustomerBillingDetailsDto) GetAdministrativeInfo() CustomerBillingDetailsDtoAdministrativeInfoDto`

GetAdministrativeInfo returns the AdministrativeInfo field if non-nil, zero value otherwise.

### GetAdministrativeInfoOk

`func (o *CustomerBillingDetailsDto) GetAdministrativeInfoOk() (*CustomerBillingDetailsDtoAdministrativeInfoDto, bool)`

GetAdministrativeInfoOk returns a tuple with the AdministrativeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeInfo

`func (o *CustomerBillingDetailsDto) SetAdministrativeInfo(v CustomerBillingDetailsDtoAdministrativeInfoDto)`

SetAdministrativeInfo sets AdministrativeInfo field to given value.

### HasAdministrativeInfo

`func (o *CustomerBillingDetailsDto) HasAdministrativeInfo() bool`

HasAdministrativeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


