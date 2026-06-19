# CreateCustomerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Must contain only ASCII letters (A-Z, a-z), digits (0-9), underscores (_), or hyphens (-). | [optional] 
**AgencyCode** | **string** | Mandatory. | 
**CustomerFamilyCode** | Pointer to **string** |  | [optional] 
**UseAgencyAddressInsteadOfOperationAddress** | Pointer to **bool** | Indicates whether the agency address should be used instead of the operation address.    If true, the provided operation address will be ignored and the agency address will be used instead. | [optional] 
**OperationAddress** | Pointer to [**OperationAddressEditionDto**](OperationAddressEditionDto.md) |  | [optional] 
**OperationalConfiguration** | Pointer to [**CustomerOperationalConfigurationEditionDto**](CustomerOperationalConfigurationEditionDto.md) |  | [optional] 
**BillingAddress** | Pointer to [**CustomerBillingAddressEditionDto**](CustomerBillingAddressEditionDto.md) |  | [optional] 
**BillingConfiguration** | Pointer to [**CustomerBillingConfigurationEditionDto**](CustomerBillingConfigurationEditionDto.md) |  | [optional] 
**BillingDetails** | Pointer to [**CustomerBillingDetailsEditionDto**](CustomerBillingDetailsEditionDto.md) |  | [optional] 

## Methods

### NewCreateCustomerCommand

`func NewCreateCustomerCommand(agencyCode string, ) *CreateCustomerCommand`

NewCreateCustomerCommand instantiates a new CreateCustomerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerCommandWithDefaults

`func NewCreateCustomerCommandWithDefaults() *CreateCustomerCommand`

NewCreateCustomerCommandWithDefaults instantiates a new CreateCustomerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateCustomerCommand) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCustomerCommand) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCustomerCommand) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateCustomerCommand) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *CreateCustomerCommand) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CreateCustomerCommand) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CreateCustomerCommand) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.


### GetCustomerFamilyCode

`func (o *CreateCustomerCommand) GetCustomerFamilyCode() string`

GetCustomerFamilyCode returns the CustomerFamilyCode field if non-nil, zero value otherwise.

### GetCustomerFamilyCodeOk

`func (o *CreateCustomerCommand) GetCustomerFamilyCodeOk() (*string, bool)`

GetCustomerFamilyCodeOk returns a tuple with the CustomerFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFamilyCode

`func (o *CreateCustomerCommand) SetCustomerFamilyCode(v string)`

SetCustomerFamilyCode sets CustomerFamilyCode field to given value.

### HasCustomerFamilyCode

`func (o *CreateCustomerCommand) HasCustomerFamilyCode() bool`

HasCustomerFamilyCode returns a boolean if a field has been set.

### GetUseAgencyAddressInsteadOfOperationAddress

`func (o *CreateCustomerCommand) GetUseAgencyAddressInsteadOfOperationAddress() bool`

GetUseAgencyAddressInsteadOfOperationAddress returns the UseAgencyAddressInsteadOfOperationAddress field if non-nil, zero value otherwise.

### GetUseAgencyAddressInsteadOfOperationAddressOk

`func (o *CreateCustomerCommand) GetUseAgencyAddressInsteadOfOperationAddressOk() (*bool, bool)`

GetUseAgencyAddressInsteadOfOperationAddressOk returns a tuple with the UseAgencyAddressInsteadOfOperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgencyAddressInsteadOfOperationAddress

`func (o *CreateCustomerCommand) SetUseAgencyAddressInsteadOfOperationAddress(v bool)`

SetUseAgencyAddressInsteadOfOperationAddress sets UseAgencyAddressInsteadOfOperationAddress field to given value.

### HasUseAgencyAddressInsteadOfOperationAddress

`func (o *CreateCustomerCommand) HasUseAgencyAddressInsteadOfOperationAddress() bool`

HasUseAgencyAddressInsteadOfOperationAddress returns a boolean if a field has been set.

### GetOperationAddress

`func (o *CreateCustomerCommand) GetOperationAddress() OperationAddressEditionDto`

GetOperationAddress returns the OperationAddress field if non-nil, zero value otherwise.

### GetOperationAddressOk

`func (o *CreateCustomerCommand) GetOperationAddressOk() (*OperationAddressEditionDto, bool)`

GetOperationAddressOk returns a tuple with the OperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddress

`func (o *CreateCustomerCommand) SetOperationAddress(v OperationAddressEditionDto)`

SetOperationAddress sets OperationAddress field to given value.

### HasOperationAddress

`func (o *CreateCustomerCommand) HasOperationAddress() bool`

HasOperationAddress returns a boolean if a field has been set.

### GetOperationalConfiguration

`func (o *CreateCustomerCommand) GetOperationalConfiguration() CustomerOperationalConfigurationEditionDto`

GetOperationalConfiguration returns the OperationalConfiguration field if non-nil, zero value otherwise.

### GetOperationalConfigurationOk

`func (o *CreateCustomerCommand) GetOperationalConfigurationOk() (*CustomerOperationalConfigurationEditionDto, bool)`

GetOperationalConfigurationOk returns a tuple with the OperationalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalConfiguration

`func (o *CreateCustomerCommand) SetOperationalConfiguration(v CustomerOperationalConfigurationEditionDto)`

SetOperationalConfiguration sets OperationalConfiguration field to given value.

### HasOperationalConfiguration

`func (o *CreateCustomerCommand) HasOperationalConfiguration() bool`

HasOperationalConfiguration returns a boolean if a field has been set.

### GetBillingAddress

`func (o *CreateCustomerCommand) GetBillingAddress() CustomerBillingAddressEditionDto`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CreateCustomerCommand) GetBillingAddressOk() (*CustomerBillingAddressEditionDto, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CreateCustomerCommand) SetBillingAddress(v CustomerBillingAddressEditionDto)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CreateCustomerCommand) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *CreateCustomerCommand) GetBillingConfiguration() CustomerBillingConfigurationEditionDto`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *CreateCustomerCommand) GetBillingConfigurationOk() (*CustomerBillingConfigurationEditionDto, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *CreateCustomerCommand) SetBillingConfiguration(v CustomerBillingConfigurationEditionDto)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *CreateCustomerCommand) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.

### GetBillingDetails

`func (o *CreateCustomerCommand) GetBillingDetails() CustomerBillingDetailsEditionDto`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *CreateCustomerCommand) GetBillingDetailsOk() (*CustomerBillingDetailsEditionDto, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *CreateCustomerCommand) SetBillingDetails(v CustomerBillingDetailsEditionDto)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *CreateCustomerCommand) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


