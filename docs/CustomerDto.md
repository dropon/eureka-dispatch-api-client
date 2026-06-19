# CustomerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | Customer&#39;s unique identifier | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**CreationDateTime** | Pointer to **time.Time** |  | [optional] 
**LastModificationDateTime** | Pointer to **time.Time** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**OperationAddressId** | Pointer to **int32** |  | [optional] 
**UseAgencyAddressInsteadOfOperationAddress** | Pointer to **bool** | If true, the agency address should be used instead of operation address. | [optional] 
**CustomerFamilyCode** | Pointer to **string** |  | [optional] 
**CustomerFamilyLabel** | Pointer to **string** |  | [optional] 
**EdiConfigurationId** | Pointer to **int32** |  | [optional] 
**OperationalConfiguration** | Pointer to [**CustomerOperationalConfigurationDto**](CustomerOperationalConfigurationDto.md) |  | [optional] 
**BillingAddress** | Pointer to [**CustomerBillingAddressDto**](CustomerBillingAddressDto.md) |  | [optional] 
**BillingConfiguration** | Pointer to [**CustomerBillingConfigurationDto**](CustomerBillingConfigurationDto.md) |  | [optional] 
**BillingDetails** | Pointer to [**CustomerBillingDetailsDto**](CustomerBillingDetailsDto.md) |  | [optional] 
**Included** | Pointer to [**CustomerIncludedDto**](CustomerIncludedDto.md) |  | [optional] 

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

### GetCreationDateTime

`func (o *CustomerDto) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *CustomerDto) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *CustomerDto) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *CustomerDto) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetLastModificationDateTime

`func (o *CustomerDto) GetLastModificationDateTime() time.Time`

GetLastModificationDateTime returns the LastModificationDateTime field if non-nil, zero value otherwise.

### GetLastModificationDateTimeOk

`func (o *CustomerDto) GetLastModificationDateTimeOk() (*time.Time, bool)`

GetLastModificationDateTimeOk returns a tuple with the LastModificationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationDateTime

`func (o *CustomerDto) SetLastModificationDateTime(v time.Time)`

SetLastModificationDateTime sets LastModificationDateTime field to given value.

### HasLastModificationDateTime

`func (o *CustomerDto) HasLastModificationDateTime() bool`

HasLastModificationDateTime returns a boolean if a field has been set.

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

### GetOperationAddressId

`func (o *CustomerDto) GetOperationAddressId() int32`

GetOperationAddressId returns the OperationAddressId field if non-nil, zero value otherwise.

### GetOperationAddressIdOk

`func (o *CustomerDto) GetOperationAddressIdOk() (*int32, bool)`

GetOperationAddressIdOk returns a tuple with the OperationAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAddressId

`func (o *CustomerDto) SetOperationAddressId(v int32)`

SetOperationAddressId sets OperationAddressId field to given value.

### HasOperationAddressId

`func (o *CustomerDto) HasOperationAddressId() bool`

HasOperationAddressId returns a boolean if a field has been set.

### GetUseAgencyAddressInsteadOfOperationAddress

`func (o *CustomerDto) GetUseAgencyAddressInsteadOfOperationAddress() bool`

GetUseAgencyAddressInsteadOfOperationAddress returns the UseAgencyAddressInsteadOfOperationAddress field if non-nil, zero value otherwise.

### GetUseAgencyAddressInsteadOfOperationAddressOk

`func (o *CustomerDto) GetUseAgencyAddressInsteadOfOperationAddressOk() (*bool, bool)`

GetUseAgencyAddressInsteadOfOperationAddressOk returns a tuple with the UseAgencyAddressInsteadOfOperationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgencyAddressInsteadOfOperationAddress

`func (o *CustomerDto) SetUseAgencyAddressInsteadOfOperationAddress(v bool)`

SetUseAgencyAddressInsteadOfOperationAddress sets UseAgencyAddressInsteadOfOperationAddress field to given value.

### HasUseAgencyAddressInsteadOfOperationAddress

`func (o *CustomerDto) HasUseAgencyAddressInsteadOfOperationAddress() bool`

HasUseAgencyAddressInsteadOfOperationAddress returns a boolean if a field has been set.

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

### GetBillingAddress

`func (o *CustomerDto) GetBillingAddress() CustomerBillingAddressDto`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CustomerDto) GetBillingAddressOk() (*CustomerBillingAddressDto, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CustomerDto) SetBillingAddress(v CustomerBillingAddressDto)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CustomerDto) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBillingConfiguration

`func (o *CustomerDto) GetBillingConfiguration() CustomerBillingConfigurationDto`

GetBillingConfiguration returns the BillingConfiguration field if non-nil, zero value otherwise.

### GetBillingConfigurationOk

`func (o *CustomerDto) GetBillingConfigurationOk() (*CustomerBillingConfigurationDto, bool)`

GetBillingConfigurationOk returns a tuple with the BillingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfiguration

`func (o *CustomerDto) SetBillingConfiguration(v CustomerBillingConfigurationDto)`

SetBillingConfiguration sets BillingConfiguration field to given value.

### HasBillingConfiguration

`func (o *CustomerDto) HasBillingConfiguration() bool`

HasBillingConfiguration returns a boolean if a field has been set.

### GetBillingDetails

`func (o *CustomerDto) GetBillingDetails() CustomerBillingDetailsDto`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *CustomerDto) GetBillingDetailsOk() (*CustomerBillingDetailsDto, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *CustomerDto) SetBillingDetails(v CustomerBillingDetailsDto)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *CustomerDto) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### GetIncluded

`func (o *CustomerDto) GetIncluded() CustomerIncludedDto`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *CustomerDto) GetIncludedOk() (*CustomerIncludedDto, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *CustomerDto) SetIncluded(v CustomerIncludedDto)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *CustomerDto) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


