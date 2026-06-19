# CustomerBillingAddressDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StreetNumber** | Pointer to **string** |  | [optional] 
**StreetName** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**CityName** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AlternativePhone** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerBillingAddressDto

`func NewCustomerBillingAddressDto() *CustomerBillingAddressDto`

NewCustomerBillingAddressDto instantiates a new CustomerBillingAddressDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerBillingAddressDtoWithDefaults

`func NewCustomerBillingAddressDtoWithDefaults() *CustomerBillingAddressDto`

NewCustomerBillingAddressDtoWithDefaults instantiates a new CustomerBillingAddressDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerBillingAddressDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerBillingAddressDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerBillingAddressDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerBillingAddressDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetNumber

`func (o *CustomerBillingAddressDto) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *CustomerBillingAddressDto) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *CustomerBillingAddressDto) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *CustomerBillingAddressDto) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *CustomerBillingAddressDto) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *CustomerBillingAddressDto) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *CustomerBillingAddressDto) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *CustomerBillingAddressDto) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CustomerBillingAddressDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomerBillingAddressDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomerBillingAddressDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomerBillingAddressDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CustomerBillingAddressDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomerBillingAddressDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomerBillingAddressDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomerBillingAddressDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCityName

`func (o *CustomerBillingAddressDto) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *CustomerBillingAddressDto) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *CustomerBillingAddressDto) SetCityName(v string)`

SetCityName sets CityName field to given value.

### HasCityName

`func (o *CustomerBillingAddressDto) HasCityName() bool`

HasCityName returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerBillingAddressDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerBillingAddressDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerBillingAddressDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomerBillingAddressDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *CustomerBillingAddressDto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CustomerBillingAddressDto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CustomerBillingAddressDto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CustomerBillingAddressDto) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetContact

`func (o *CustomerBillingAddressDto) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CustomerBillingAddressDto) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CustomerBillingAddressDto) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CustomerBillingAddressDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetEmails

`func (o *CustomerBillingAddressDto) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CustomerBillingAddressDto) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CustomerBillingAddressDto) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CustomerBillingAddressDto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerBillingAddressDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerBillingAddressDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerBillingAddressDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerBillingAddressDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *CustomerBillingAddressDto) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *CustomerBillingAddressDto) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *CustomerBillingAddressDto) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *CustomerBillingAddressDto) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *CustomerBillingAddressDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *CustomerBillingAddressDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *CustomerBillingAddressDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *CustomerBillingAddressDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetFax

`func (o *CustomerBillingAddressDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CustomerBillingAddressDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CustomerBillingAddressDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CustomerBillingAddressDto) HasFax() bool`

HasFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


