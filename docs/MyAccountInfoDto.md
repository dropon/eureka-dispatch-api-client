# MyAccountInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**HasAccessToEurekaMaps** | Pointer to **bool** |  | [optional] 
**HasAccessToGeolocation** | Pointer to **bool** |  | [optional] 
**HasAccessToOperatorRounds** | Pointer to **bool** |  | [optional] 
**HasAccessToWorkingTime** | Pointer to **bool** |  | [optional] 
**HasAccessToPRMReadOnly** | Pointer to **bool** |  | [optional] 
**HasAccessToPRMReadWrite** | Pointer to **bool** |  | [optional] 

## Methods

### NewMyAccountInfoDto

`func NewMyAccountInfoDto() *MyAccountInfoDto`

NewMyAccountInfoDto instantiates a new MyAccountInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyAccountInfoDtoWithDefaults

`func NewMyAccountInfoDtoWithDefaults() *MyAccountInfoDto`

NewMyAccountInfoDtoWithDefaults instantiates a new MyAccountInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *MyAccountInfoDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *MyAccountInfoDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *MyAccountInfoDto) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *MyAccountInfoDto) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *MyAccountInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MyAccountInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MyAccountInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MyAccountInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *MyAccountInfoDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MyAccountInfoDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MyAccountInfoDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MyAccountInfoDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLocale

`func (o *MyAccountInfoDto) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MyAccountInfoDto) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MyAccountInfoDto) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MyAccountInfoDto) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPhone

`func (o *MyAccountInfoDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MyAccountInfoDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MyAccountInfoDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MyAccountInfoDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetHasAccessToEurekaMaps

`func (o *MyAccountInfoDto) GetHasAccessToEurekaMaps() bool`

GetHasAccessToEurekaMaps returns the HasAccessToEurekaMaps field if non-nil, zero value otherwise.

### GetHasAccessToEurekaMapsOk

`func (o *MyAccountInfoDto) GetHasAccessToEurekaMapsOk() (*bool, bool)`

GetHasAccessToEurekaMapsOk returns a tuple with the HasAccessToEurekaMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToEurekaMaps

`func (o *MyAccountInfoDto) SetHasAccessToEurekaMaps(v bool)`

SetHasAccessToEurekaMaps sets HasAccessToEurekaMaps field to given value.

### HasHasAccessToEurekaMaps

`func (o *MyAccountInfoDto) HasHasAccessToEurekaMaps() bool`

HasHasAccessToEurekaMaps returns a boolean if a field has been set.

### GetHasAccessToGeolocation

`func (o *MyAccountInfoDto) GetHasAccessToGeolocation() bool`

GetHasAccessToGeolocation returns the HasAccessToGeolocation field if non-nil, zero value otherwise.

### GetHasAccessToGeolocationOk

`func (o *MyAccountInfoDto) GetHasAccessToGeolocationOk() (*bool, bool)`

GetHasAccessToGeolocationOk returns a tuple with the HasAccessToGeolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToGeolocation

`func (o *MyAccountInfoDto) SetHasAccessToGeolocation(v bool)`

SetHasAccessToGeolocation sets HasAccessToGeolocation field to given value.

### HasHasAccessToGeolocation

`func (o *MyAccountInfoDto) HasHasAccessToGeolocation() bool`

HasHasAccessToGeolocation returns a boolean if a field has been set.

### GetHasAccessToOperatorRounds

`func (o *MyAccountInfoDto) GetHasAccessToOperatorRounds() bool`

GetHasAccessToOperatorRounds returns the HasAccessToOperatorRounds field if non-nil, zero value otherwise.

### GetHasAccessToOperatorRoundsOk

`func (o *MyAccountInfoDto) GetHasAccessToOperatorRoundsOk() (*bool, bool)`

GetHasAccessToOperatorRoundsOk returns a tuple with the HasAccessToOperatorRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToOperatorRounds

`func (o *MyAccountInfoDto) SetHasAccessToOperatorRounds(v bool)`

SetHasAccessToOperatorRounds sets HasAccessToOperatorRounds field to given value.

### HasHasAccessToOperatorRounds

`func (o *MyAccountInfoDto) HasHasAccessToOperatorRounds() bool`

HasHasAccessToOperatorRounds returns a boolean if a field has been set.

### GetHasAccessToWorkingTime

`func (o *MyAccountInfoDto) GetHasAccessToWorkingTime() bool`

GetHasAccessToWorkingTime returns the HasAccessToWorkingTime field if non-nil, zero value otherwise.

### GetHasAccessToWorkingTimeOk

`func (o *MyAccountInfoDto) GetHasAccessToWorkingTimeOk() (*bool, bool)`

GetHasAccessToWorkingTimeOk returns a tuple with the HasAccessToWorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToWorkingTime

`func (o *MyAccountInfoDto) SetHasAccessToWorkingTime(v bool)`

SetHasAccessToWorkingTime sets HasAccessToWorkingTime field to given value.

### HasHasAccessToWorkingTime

`func (o *MyAccountInfoDto) HasHasAccessToWorkingTime() bool`

HasHasAccessToWorkingTime returns a boolean if a field has been set.

### GetHasAccessToPRMReadOnly

`func (o *MyAccountInfoDto) GetHasAccessToPRMReadOnly() bool`

GetHasAccessToPRMReadOnly returns the HasAccessToPRMReadOnly field if non-nil, zero value otherwise.

### GetHasAccessToPRMReadOnlyOk

`func (o *MyAccountInfoDto) GetHasAccessToPRMReadOnlyOk() (*bool, bool)`

GetHasAccessToPRMReadOnlyOk returns a tuple with the HasAccessToPRMReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToPRMReadOnly

`func (o *MyAccountInfoDto) SetHasAccessToPRMReadOnly(v bool)`

SetHasAccessToPRMReadOnly sets HasAccessToPRMReadOnly field to given value.

### HasHasAccessToPRMReadOnly

`func (o *MyAccountInfoDto) HasHasAccessToPRMReadOnly() bool`

HasHasAccessToPRMReadOnly returns a boolean if a field has been set.

### GetHasAccessToPRMReadWrite

`func (o *MyAccountInfoDto) GetHasAccessToPRMReadWrite() bool`

GetHasAccessToPRMReadWrite returns the HasAccessToPRMReadWrite field if non-nil, zero value otherwise.

### GetHasAccessToPRMReadWriteOk

`func (o *MyAccountInfoDto) GetHasAccessToPRMReadWriteOk() (*bool, bool)`

GetHasAccessToPRMReadWriteOk returns a tuple with the HasAccessToPRMReadWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToPRMReadWrite

`func (o *MyAccountInfoDto) SetHasAccessToPRMReadWrite(v bool)`

SetHasAccessToPRMReadWrite sets HasAccessToPRMReadWrite field to given value.

### HasHasAccessToPRMReadWrite

`func (o *MyAccountInfoDto) HasHasAccessToPRMReadWrite() bool`

HasHasAccessToPRMReadWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


