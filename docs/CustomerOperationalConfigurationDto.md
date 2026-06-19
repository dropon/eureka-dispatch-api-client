# CustomerOperationalConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**PriceCatalogCode** | Pointer to **string** |  | [optional] 
**ReferencesConfiguration** | Pointer to [**CustomerReferencesConfigurationDto**](CustomerReferencesConfigurationDto.md) |  | [optional] 

## Methods

### NewCustomerOperationalConfigurationDto

`func NewCustomerOperationalConfigurationDto() *CustomerOperationalConfigurationDto`

NewCustomerOperationalConfigurationDto instantiates a new CustomerOperationalConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOperationalConfigurationDtoWithDefaults

`func NewCustomerOperationalConfigurationDtoWithDefaults() *CustomerOperationalConfigurationDto`

NewCustomerOperationalConfigurationDtoWithDefaults instantiates a new CustomerOperationalConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *CustomerOperationalConfigurationDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomerOperationalConfigurationDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomerOperationalConfigurationDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CustomerOperationalConfigurationDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriceCatalogCode

`func (o *CustomerOperationalConfigurationDto) GetPriceCatalogCode() string`

GetPriceCatalogCode returns the PriceCatalogCode field if non-nil, zero value otherwise.

### GetPriceCatalogCodeOk

`func (o *CustomerOperationalConfigurationDto) GetPriceCatalogCodeOk() (*string, bool)`

GetPriceCatalogCodeOk returns a tuple with the PriceCatalogCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCatalogCode

`func (o *CustomerOperationalConfigurationDto) SetPriceCatalogCode(v string)`

SetPriceCatalogCode sets PriceCatalogCode field to given value.

### HasPriceCatalogCode

`func (o *CustomerOperationalConfigurationDto) HasPriceCatalogCode() bool`

HasPriceCatalogCode returns a boolean if a field has been set.

### GetReferencesConfiguration

`func (o *CustomerOperationalConfigurationDto) GetReferencesConfiguration() CustomerReferencesConfigurationDto`

GetReferencesConfiguration returns the ReferencesConfiguration field if non-nil, zero value otherwise.

### GetReferencesConfigurationOk

`func (o *CustomerOperationalConfigurationDto) GetReferencesConfigurationOk() (*CustomerReferencesConfigurationDto, bool)`

GetReferencesConfigurationOk returns a tuple with the ReferencesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencesConfiguration

`func (o *CustomerOperationalConfigurationDto) SetReferencesConfiguration(v CustomerReferencesConfigurationDto)`

SetReferencesConfiguration sets ReferencesConfiguration field to given value.

### HasReferencesConfiguration

`func (o *CustomerOperationalConfigurationDto) HasReferencesConfiguration() bool`

HasReferencesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


