# CustomerOperationalConfigurationEditionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | Indicate whether the customer is enabled. If set to null, the value defaults to true. | [optional] 
**PriceCatalogCode** | Pointer to **string** |  | [optional] 
**ReferencesConfiguration** | Pointer to [**CustomerOperationalConfigurationEditionDtoCustomerReferencesConfigurationEditionDto**](CustomerOperationalConfigurationEditionDtoCustomerReferencesConfigurationEditionDto.md) |  | [optional] 

## Methods

### NewCustomerOperationalConfigurationEditionDto

`func NewCustomerOperationalConfigurationEditionDto() *CustomerOperationalConfigurationEditionDto`

NewCustomerOperationalConfigurationEditionDto instantiates a new CustomerOperationalConfigurationEditionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOperationalConfigurationEditionDtoWithDefaults

`func NewCustomerOperationalConfigurationEditionDtoWithDefaults() *CustomerOperationalConfigurationEditionDto`

NewCustomerOperationalConfigurationEditionDtoWithDefaults instantiates a new CustomerOperationalConfigurationEditionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *CustomerOperationalConfigurationEditionDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomerOperationalConfigurationEditionDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomerOperationalConfigurationEditionDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CustomerOperationalConfigurationEditionDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriceCatalogCode

`func (o *CustomerOperationalConfigurationEditionDto) GetPriceCatalogCode() string`

GetPriceCatalogCode returns the PriceCatalogCode field if non-nil, zero value otherwise.

### GetPriceCatalogCodeOk

`func (o *CustomerOperationalConfigurationEditionDto) GetPriceCatalogCodeOk() (*string, bool)`

GetPriceCatalogCodeOk returns a tuple with the PriceCatalogCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCatalogCode

`func (o *CustomerOperationalConfigurationEditionDto) SetPriceCatalogCode(v string)`

SetPriceCatalogCode sets PriceCatalogCode field to given value.

### HasPriceCatalogCode

`func (o *CustomerOperationalConfigurationEditionDto) HasPriceCatalogCode() bool`

HasPriceCatalogCode returns a boolean if a field has been set.

### GetReferencesConfiguration

`func (o *CustomerOperationalConfigurationEditionDto) GetReferencesConfiguration() CustomerOperationalConfigurationEditionDtoCustomerReferencesConfigurationEditionDto`

GetReferencesConfiguration returns the ReferencesConfiguration field if non-nil, zero value otherwise.

### GetReferencesConfigurationOk

`func (o *CustomerOperationalConfigurationEditionDto) GetReferencesConfigurationOk() (*CustomerOperationalConfigurationEditionDtoCustomerReferencesConfigurationEditionDto, bool)`

GetReferencesConfigurationOk returns a tuple with the ReferencesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencesConfiguration

`func (o *CustomerOperationalConfigurationEditionDto) SetReferencesConfiguration(v CustomerOperationalConfigurationEditionDtoCustomerReferencesConfigurationEditionDto)`

SetReferencesConfiguration sets ReferencesConfiguration field to given value.

### HasReferencesConfiguration

`func (o *CustomerOperationalConfigurationEditionDto) HasReferencesConfiguration() bool`

HasReferencesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


