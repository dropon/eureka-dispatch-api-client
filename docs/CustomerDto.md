# CustomerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Customer&#39;s unique identifier | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**CustomerFamilyCode** | Pointer to **string** |  | [optional] 
**CustomerFamilyLabel** | Pointer to **string** |  | [optional] 
**EdiConfigurationId** | Pointer to **int32** |  | [optional] 
**OperationalConfiguration** | Pointer to [**CustomerOperationalConfigurationDto**](CustomerOperationalConfigurationDto.md) |  | [optional] 

## Methods

### NewCustomerDto

`func NewCustomerDto() *CustomerDto`

NewCustomerDto instantiates a new CustomerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDtoWithDefaults

`func NewCustomerDtoWithDefaults() *CustomerDto`

NewCustomerDtoWithDefaults instantiates a new CustomerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *CustomerDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CustomerDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CustomerDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CustomerDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *CustomerDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomerDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *CustomerDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomerDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomerDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomerDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAgencyCode

`func (o *CustomerDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CustomerDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CustomerDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *CustomerDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetCustomerFamilyCode

`func (o *CustomerDto) GetCustomerFamilyCode() string`

GetCustomerFamilyCode returns the CustomerFamilyCode field if non-nil, zero value otherwise.

### GetCustomerFamilyCodeOk

`func (o *CustomerDto) GetCustomerFamilyCodeOk() (*string, bool)`

GetCustomerFamilyCodeOk returns a tuple with the CustomerFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFamilyCode

`func (o *CustomerDto) SetCustomerFamilyCode(v string)`

SetCustomerFamilyCode sets CustomerFamilyCode field to given value.

### HasCustomerFamilyCode

`func (o *CustomerDto) HasCustomerFamilyCode() bool`

HasCustomerFamilyCode returns a boolean if a field has been set.

### GetCustomerFamilyLabel

`func (o *CustomerDto) GetCustomerFamilyLabel() string`

GetCustomerFamilyLabel returns the CustomerFamilyLabel field if non-nil, zero value otherwise.

### GetCustomerFamilyLabelOk

`func (o *CustomerDto) GetCustomerFamilyLabelOk() (*string, bool)`

GetCustomerFamilyLabelOk returns a tuple with the CustomerFamilyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFamilyLabel

`func (o *CustomerDto) SetCustomerFamilyLabel(v string)`

SetCustomerFamilyLabel sets CustomerFamilyLabel field to given value.

### HasCustomerFamilyLabel

`func (o *CustomerDto) HasCustomerFamilyLabel() bool`

HasCustomerFamilyLabel returns a boolean if a field has been set.

### GetEdiConfigurationId

`func (o *CustomerDto) GetEdiConfigurationId() int32`

GetEdiConfigurationId returns the EdiConfigurationId field if non-nil, zero value otherwise.

### GetEdiConfigurationIdOk

`func (o *CustomerDto) GetEdiConfigurationIdOk() (*int32, bool)`

GetEdiConfigurationIdOk returns a tuple with the EdiConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdiConfigurationId

`func (o *CustomerDto) SetEdiConfigurationId(v int32)`

SetEdiConfigurationId sets EdiConfigurationId field to given value.

### HasEdiConfigurationId

`func (o *CustomerDto) HasEdiConfigurationId() bool`

HasEdiConfigurationId returns a boolean if a field has been set.

### GetOperationalConfiguration

`func (o *CustomerDto) GetOperationalConfiguration() CustomerOperationalConfigurationDto`

GetOperationalConfiguration returns the OperationalConfiguration field if non-nil, zero value otherwise.

### GetOperationalConfigurationOk

`func (o *CustomerDto) GetOperationalConfigurationOk() (*CustomerOperationalConfigurationDto, bool)`

GetOperationalConfigurationOk returns a tuple with the OperationalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalConfiguration

`func (o *CustomerDto) SetOperationalConfiguration(v CustomerOperationalConfigurationDto)`

SetOperationalConfiguration sets OperationalConfiguration field to given value.

### HasOperationalConfiguration

`func (o *CustomerDto) HasOperationalConfiguration() bool`

HasOperationalConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


