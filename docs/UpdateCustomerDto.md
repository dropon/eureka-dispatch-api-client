# UpdateCustomerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgencyCode** | **string** | Mandatory. | 
**CustomerFamilyCode** | Pointer to **string** |  | [optional] 
**UseAgencyAddressInsteadOfOperationAddress** | Pointer to **bool** | Indicates whether the agency address should be used instead of the operation address.    If true, the provided operation address will be ignored and the agency address will be used instead. | [optional] 
**OperationAddress** | Pointer to [**OperationAddressEditionDto**](OperationAddressEditionDto.md) |  | [optional] 
**OperationalConfiguration** | Pointer to [**CustomerOperationalConfigurationEditionDto**](CustomerOperationalConfigurationEditionDto.md) |  | [optional] 
**BillingAddress** | Pointer to [**CustomerBillingAddressEditionDto**](CustomerBillingAddressEditionDto.md) |  | [optional] 
**BillingConfiguration** | Pointer to [**CustomerBillingConfigurationEditionDto**](CustomerBillingConfigurationEditionDto.md) |  | [optional] 
**BillingDetails** | Pointer to [**CustomerBillingDetailsEditionDto**](CustomerBillingDetailsEditionDto.md) |  | [optional] 

## Methods

### NewUpdateCustomerDto

`func NewUpdateCustomerDto(agencyCode string, ) *UpdateCustomerDto`

NewUpdateCustomerDto instantiates a new UpdateCustomerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerDtoWithDefaults

`func NewUpdateCustomerDtoWithDefaults() *UpdateCustomerDto`

NewUpdateCustomerDtoWithDefaults instantiates a new UpdateCustomerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgencyCode

`func (o *UpdateCustomerDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *UpdateCustomerDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *UpdateCustomerDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.


### GetCustomerFamilyCode

`func (o *UpdateCustomerDto) GetCustomerFamilyCode() string`

GetCustomerFamilyCode returns the CustomerFamilyCode field if non-nil, zero value otherwise.

### GetCustomerFamilyCodeOk

`func (o *UpdateCustomerDto) GetCustomerFamilyCodeOk() (*string, bool)`

GetCustomerFamilyCodeOk returns a tuple with the CustomerFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFamilyCode

`func (o *UpdateCustomerDto) SetCustomerFamilyCode(v string)`

SetCustomerFamilyCode sets CustomerFamilyCode field to given value.

### HasCustomerFamilyCode

`func (o *UpdateCustomerDto) HasCustomerFamilyCode() bool`

HasCustomerFamilyCode returns a boolean if a field has been set.

### GetUseAgencyAddressInsteadOfOperationAddress

`func (o *UpdateCustomerDto) GetUseAgencyAddressInsteadOfOperationAddress() bool`

GetUseAgencyAddressInsteadOfOperationAddress returns the UseAgencyAddressInsteadOfOperationAddress field if non-nil, zero value otherwise.

### GetUseAgencyAddressInsteadOfOperationAddressOk

`func (o *UpdateCustomerDto) GetUseAgencyAddressInsteadOfOperationAddressOk() (*bool, bool)`

GetUseAgencyAddressInsteadOfOperationAddressOk returns a tuple with the UseAgencyAddressInsteadOfOperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgencyAddressInsteadOfOperationAddress

`func (o *UpdateCustomerDto) SetUseAgencyAddressInsteadOfOperationAddress(v bool)`

SetUseAgencyAddressInsteadOfOperationAddress sets UseAgencyAddressInsteadOfOperationAddress field to given value.

### HasUseAgencyAddressInsteadOfOperationAddress

`func (o *UpdateCustomerDto) HasUseAgencyAddressInsteadOfOperationAddress() bool`

HasUseAgencyAddressInsteadOfOperationAddress returns a boolean if a field has been set.

### GetOperationAddress

`func (o *UpdateCustomerDto) GetOperationAddress() OperationAddressEditionDto`

GetOperationAddress returns the OperationAddress field if non-nil, zero value otherwise.

### GetOperationAddressOk

`func (o *UpdateCustomerDto) GetOperationAddressOk() (*OperationAddressEditionDto, bool)`

GetOperationAddressOk returns a tuple with the OperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddress

`func (o *UpdateCustomerDto) SetOperationAddress(v OperationAddressEditionDto)`

SetOperationAddress sets OperationAddress field to given value.

### HasOperationAddress

`func (o *UpdateCustomerDto) HasOperationAddress() bool`

HasOperationAddress returns a boolean if a field has been set.

### GetOperationalConfiguration

`func (o *UpdateCustomerDto) GetOperationalConfiguration() CustomerOperationalConfigurationEditionDto`

GetOperationalConfiguration returns the OperationalConfiguration field if non-nil, zero value otherwise.

### GetOperationalConfigurationOk

`func (o *UpdateCustomerDto) GetOperationalConfigurationOk() (*CustomerOperationalConfigurationEditionDto, bool)`

GetOperationalConfigurationOk returns a tuple with the OperationalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalConfiguration

`func (o *UpdateCustomerDto) SetOperationalConfiguration(v CustomerOperationalConfigurationEditionDto)`

SetOperationalConfiguration sets OperationalConfiguration field to given value.

### HasOperationalConfiguration

`func (o *UpdateCustomerDto) HasOperationalConfiguration() bool`

HasOperationalConfiguration returns a boolean if a field has been set.

### GetBillingAddress

`func (o *UpdateCustomerDto) GetBillingAddress() CustomerBillingAddressEditionDto`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *UpdateCustomerDto) GetBillingAddressOk() (*CustomerBillingAddressEditionDto, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *UpdateCustomerDto) SetBillingAddress(v CustomerBillingAddressEditionDto)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *UpdateCustomerDto) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *UpdateCustomerDto) GetBillingConfiguration() CustomerBillingConfigurationEditionDto`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *UpdateCustomerDto) GetBillingConfigurationOk() (*CustomerBillingConfigurationEditionDto, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *UpdateCustomerDto) SetBillingConfiguration(v CustomerBillingConfigurationEditionDto)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *UpdateCustomerDto) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.

### GetBillingDetails

`func (o *UpdateCustomerDto) GetBillingDetails() CustomerBillingDetailsEditionDto`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *UpdateCustomerDto) GetBillingDetailsOk() (*CustomerBillingDetailsEditionDto, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *UpdateCustomerDto) SetBillingDetails(v CustomerBillingDetailsEditionDto)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *UpdateCustomerDto) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


