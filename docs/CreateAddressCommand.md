# CreateAddressCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Mandatory. | 
**StreetNumber** | Pointer to **string** |  | [optional] 
**StreetName** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AlternativePhone** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**CityId** | Pointer to **int32** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**AddressCategoryCode** | Pointer to **string** |  | [optional] 
**ExternalAddressCode** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**StopDurationInSeconds** | Pointer to **int32** |  | [optional] 
**CircuitAddressSetId** | Pointer to **int32** |  | [optional] 
**AddressComplement1** | Pointer to **string** |  | [optional] 
**AddressComplement2** | Pointer to **string** |  | [optional] 
**AddressComplement3** | Pointer to **string** |  | [optional] 
**AddressComplement4** | Pointer to **string** |  | [optional] 
**IsHeightRestrictedAccess** | Pointer to **bool** |  | [optional] 
**HeightLimitation** | Pointer to **int32** |  | [optional] 
**IsHatchBackMandatory** | Pointer to **bool** |  | [optional] 
**IsWarehousemanMandatory** | Pointer to **bool** |  | [optional] 
**IsHandlingEquipmentMandatory** | Pointer to **bool** |  | [optional] 
**IsSecuritySuitcaseMandatory** | Pointer to **bool** |  | [optional] 
**IsSideUnloading** | Pointer to **bool** |  | [optional] 
**IsCraneUnloading** | Pointer to **bool** |  | [optional] 
**Geocoding** | Pointer to [**AddressGeocodingEditionDto**](AddressGeocodingEditionDto.md) |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 
**AdditionalContacts** | Pointer to [**[]CreateAddressCommandCreateAddressAdditionalContactDto**](CreateAddressCommandCreateAddressAdditionalContactDto.md) |  | [optional] 
**AddressOpenings** | Pointer to [**[]CreateAddressCommandCreateAddressOpeningDto**](CreateAddressCommandCreateAddressOpeningDto.md) |  | [optional] 
**ForbiddenVehicleTypeCodes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateAddressCommand

`func NewCreateAddressCommand(name string, ) *CreateAddressCommand`

NewCreateAddressCommand instantiates a new CreateAddressCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAddressCommandWithDefaults

`func NewCreateAddressCommandWithDefaults() *CreateAddressCommand`

NewCreateAddressCommandWithDefaults instantiates a new CreateAddressCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAddressCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAddressCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAddressCommand) SetName(v string)`

SetName sets Name field to given value.


### GetStreetNumber

`func (o *CreateAddressCommand) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *CreateAddressCommand) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *CreateAddressCommand) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *CreateAddressCommand) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *CreateAddressCommand) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *CreateAddressCommand) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *CreateAddressCommand) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *CreateAddressCommand) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CreateAddressCommand) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CreateAddressCommand) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CreateAddressCommand) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CreateAddressCommand) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CreateAddressCommand) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CreateAddressCommand) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CreateAddressCommand) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CreateAddressCommand) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetContact

`func (o *CreateAddressCommand) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CreateAddressCommand) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CreateAddressCommand) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CreateAddressCommand) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhone

`func (o *CreateAddressCommand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateAddressCommand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateAddressCommand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateAddressCommand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *CreateAddressCommand) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *CreateAddressCommand) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *CreateAddressCommand) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *CreateAddressCommand) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *CreateAddressCommand) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *CreateAddressCommand) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *CreateAddressCommand) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *CreateAddressCommand) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetFax

`func (o *CreateAddressCommand) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CreateAddressCommand) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CreateAddressCommand) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CreateAddressCommand) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetCityId

`func (o *CreateAddressCommand) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *CreateAddressCommand) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *CreateAddressCommand) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *CreateAddressCommand) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetAgencyCode

`func (o *CreateAddressCommand) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CreateAddressCommand) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CreateAddressCommand) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *CreateAddressCommand) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetAddressCategoryCode

`func (o *CreateAddressCommand) GetAddressCategoryCode() string`

GetAddressCategoryCode returns the AddressCategoryCode field if non-nil, zero value otherwise.

### GetAddressCategoryCodeOk

`func (o *CreateAddressCommand) GetAddressCategoryCodeOk() (*string, bool)`

GetAddressCategoryCodeOk returns a tuple with the AddressCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCategoryCode

`func (o *CreateAddressCommand) SetAddressCategoryCode(v string)`

SetAddressCategoryCode sets AddressCategoryCode field to given value.

### HasAddressCategoryCode

`func (o *CreateAddressCommand) HasAddressCategoryCode() bool`

HasAddressCategoryCode returns a boolean if a field has been set.

### GetExternalAddressCode

`func (o *CreateAddressCommand) GetExternalAddressCode() string`

GetExternalAddressCode returns the ExternalAddressCode field if non-nil, zero value otherwise.

### GetExternalAddressCodeOk

`func (o *CreateAddressCommand) GetExternalAddressCodeOk() (*string, bool)`

GetExternalAddressCodeOk returns a tuple with the ExternalAddressCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddressCode

`func (o *CreateAddressCommand) SetExternalAddressCode(v string)`

SetExternalAddressCode sets ExternalAddressCode field to given value.

### HasExternalAddressCode

`func (o *CreateAddressCommand) HasExternalAddressCode() bool`

HasExternalAddressCode returns a boolean if a field has been set.

### GetNotes

`func (o *CreateAddressCommand) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateAddressCommand) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateAddressCommand) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateAddressCommand) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStopDurationInSeconds

`func (o *CreateAddressCommand) GetStopDurationInSeconds() int32`

GetStopDurationInSeconds returns the StopDurationInSeconds field if non-nil, zero value otherwise.

### GetStopDurationInSecondsOk

`func (o *CreateAddressCommand) GetStopDurationInSecondsOk() (*int32, bool)`

GetStopDurationInSecondsOk returns a tuple with the StopDurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDurationInSeconds

`func (o *CreateAddressCommand) SetStopDurationInSeconds(v int32)`

SetStopDurationInSeconds sets StopDurationInSeconds field to given value.

### HasStopDurationInSeconds

`func (o *CreateAddressCommand) HasStopDurationInSeconds() bool`

HasStopDurationInSeconds returns a boolean if a field has been set.

### GetCircuitAddressSetId

`func (o *CreateAddressCommand) GetCircuitAddressSetId() int32`

GetCircuitAddressSetId returns the CircuitAddressSetId field if non-nil, zero value otherwise.

### GetCircuitAddressSetIdOk

`func (o *CreateAddressCommand) GetCircuitAddressSetIdOk() (*int32, bool)`

GetCircuitAddressSetIdOk returns a tuple with the CircuitAddressSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitAddressSetId

`func (o *CreateAddressCommand) SetCircuitAddressSetId(v int32)`

SetCircuitAddressSetId sets CircuitAddressSetId field to given value.

### HasCircuitAddressSetId

`func (o *CreateAddressCommand) HasCircuitAddressSetId() bool`

HasCircuitAddressSetId returns a boolean if a field has been set.

### GetAddressComplement1

`func (o *CreateAddressCommand) GetAddressComplement1() string`

GetAddressComplement1 returns the AddressComplement1 field if non-nil, zero value otherwise.

### GetAddressComplement1Ok

`func (o *CreateAddressCommand) GetAddressComplement1Ok() (*string, bool)`

GetAddressComplement1Ok returns a tuple with the AddressComplement1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement1

`func (o *CreateAddressCommand) SetAddressComplement1(v string)`

SetAddressComplement1 sets AddressComplement1 field to given value.

### HasAddressComplement1

`func (o *CreateAddressCommand) HasAddressComplement1() bool`

HasAddressComplement1 returns a boolean if a field has been set.

### GetAddressComplement2

`func (o *CreateAddressCommand) GetAddressComplement2() string`

GetAddressComplement2 returns the AddressComplement2 field if non-nil, zero value otherwise.

### GetAddressComplement2Ok

`func (o *CreateAddressCommand) GetAddressComplement2Ok() (*string, bool)`

GetAddressComplement2Ok returns a tuple with the AddressComplement2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement2

`func (o *CreateAddressCommand) SetAddressComplement2(v string)`

SetAddressComplement2 sets AddressComplement2 field to given value.

### HasAddressComplement2

`func (o *CreateAddressCommand) HasAddressComplement2() bool`

HasAddressComplement2 returns a boolean if a field has been set.

### GetAddressComplement3

`func (o *CreateAddressCommand) GetAddressComplement3() string`

GetAddressComplement3 returns the AddressComplement3 field if non-nil, zero value otherwise.

### GetAddressComplement3Ok

`func (o *CreateAddressCommand) GetAddressComplement3Ok() (*string, bool)`

GetAddressComplement3Ok returns a tuple with the AddressComplement3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement3

`func (o *CreateAddressCommand) SetAddressComplement3(v string)`

SetAddressComplement3 sets AddressComplement3 field to given value.

### HasAddressComplement3

`func (o *CreateAddressCommand) HasAddressComplement3() bool`

HasAddressComplement3 returns a boolean if a field has been set.

### GetAddressComplement4

`func (o *CreateAddressCommand) GetAddressComplement4() string`

GetAddressComplement4 returns the AddressComplement4 field if non-nil, zero value otherwise.

### GetAddressComplement4Ok

`func (o *CreateAddressCommand) GetAddressComplement4Ok() (*string, bool)`

GetAddressComplement4Ok returns a tuple with the AddressComplement4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement4

`func (o *CreateAddressCommand) SetAddressComplement4(v string)`

SetAddressComplement4 sets AddressComplement4 field to given value.

### HasAddressComplement4

`func (o *CreateAddressCommand) HasAddressComplement4() bool`

HasAddressComplement4 returns a boolean if a field has been set.

### GetIsHeightRestrictedAccess

`func (o *CreateAddressCommand) GetIsHeightRestrictedAccess() bool`

GetIsHeightRestrictedAccess returns the IsHeightRestrictedAccess field if non-nil, zero value otherwise.

### GetIsHeightRestrictedAccessOk

`func (o *CreateAddressCommand) GetIsHeightRestrictedAccessOk() (*bool, bool)`

GetIsHeightRestrictedAccessOk returns a tuple with the IsHeightRestrictedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeightRestrictedAccess

`func (o *CreateAddressCommand) SetIsHeightRestrictedAccess(v bool)`

SetIsHeightRestrictedAccess sets IsHeightRestrictedAccess field to given value.

### HasIsHeightRestrictedAccess

`func (o *CreateAddressCommand) HasIsHeightRestrictedAccess() bool`

HasIsHeightRestrictedAccess returns a boolean if a field has been set.

### GetHeightLimitation

`func (o *CreateAddressCommand) GetHeightLimitation() int32`

GetHeightLimitation returns the HeightLimitation field if non-nil, zero value otherwise.

### GetHeightLimitationOk

`func (o *CreateAddressCommand) GetHeightLimitationOk() (*int32, bool)`

GetHeightLimitationOk returns a tuple with the HeightLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightLimitation

`func (o *CreateAddressCommand) SetHeightLimitation(v int32)`

SetHeightLimitation sets HeightLimitation field to given value.

### HasHeightLimitation

`func (o *CreateAddressCommand) HasHeightLimitation() bool`

HasHeightLimitation returns a boolean if a field has been set.

### GetIsHatchBackMandatory

`func (o *CreateAddressCommand) GetIsHatchBackMandatory() bool`

GetIsHatchBackMandatory returns the IsHatchBackMandatory field if non-nil, zero value otherwise.

### GetIsHatchBackMandatoryOk

`func (o *CreateAddressCommand) GetIsHatchBackMandatoryOk() (*bool, bool)`

GetIsHatchBackMandatoryOk returns a tuple with the IsHatchBackMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHatchBackMandatory

`func (o *CreateAddressCommand) SetIsHatchBackMandatory(v bool)`

SetIsHatchBackMandatory sets IsHatchBackMandatory field to given value.

### HasIsHatchBackMandatory

`func (o *CreateAddressCommand) HasIsHatchBackMandatory() bool`

HasIsHatchBackMandatory returns a boolean if a field has been set.

### GetIsWarehousemanMandatory

`func (o *CreateAddressCommand) GetIsWarehousemanMandatory() bool`

GetIsWarehousemanMandatory returns the IsWarehousemanMandatory field if non-nil, zero value otherwise.

### GetIsWarehousemanMandatoryOk

`func (o *CreateAddressCommand) GetIsWarehousemanMandatoryOk() (*bool, bool)`

GetIsWarehousemanMandatoryOk returns a tuple with the IsWarehousemanMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWarehousemanMandatory

`func (o *CreateAddressCommand) SetIsWarehousemanMandatory(v bool)`

SetIsWarehousemanMandatory sets IsWarehousemanMandatory field to given value.

### HasIsWarehousemanMandatory

`func (o *CreateAddressCommand) HasIsWarehousemanMandatory() bool`

HasIsWarehousemanMandatory returns a boolean if a field has been set.

### GetIsHandlingEquipmentMandatory

`func (o *CreateAddressCommand) GetIsHandlingEquipmentMandatory() bool`

GetIsHandlingEquipmentMandatory returns the IsHandlingEquipmentMandatory field if non-nil, zero value otherwise.

### GetIsHandlingEquipmentMandatoryOk

`func (o *CreateAddressCommand) GetIsHandlingEquipmentMandatoryOk() (*bool, bool)`

GetIsHandlingEquipmentMandatoryOk returns a tuple with the IsHandlingEquipmentMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHandlingEquipmentMandatory

`func (o *CreateAddressCommand) SetIsHandlingEquipmentMandatory(v bool)`

SetIsHandlingEquipmentMandatory sets IsHandlingEquipmentMandatory field to given value.

### HasIsHandlingEquipmentMandatory

`func (o *CreateAddressCommand) HasIsHandlingEquipmentMandatory() bool`

HasIsHandlingEquipmentMandatory returns a boolean if a field has been set.

### GetIsSecuritySuitcaseMandatory

`func (o *CreateAddressCommand) GetIsSecuritySuitcaseMandatory() bool`

GetIsSecuritySuitcaseMandatory returns the IsSecuritySuitcaseMandatory field if non-nil, zero value otherwise.

### GetIsSecuritySuitcaseMandatoryOk

`func (o *CreateAddressCommand) GetIsSecuritySuitcaseMandatoryOk() (*bool, bool)`

GetIsSecuritySuitcaseMandatoryOk returns a tuple with the IsSecuritySuitcaseMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySuitcaseMandatory

`func (o *CreateAddressCommand) SetIsSecuritySuitcaseMandatory(v bool)`

SetIsSecuritySuitcaseMandatory sets IsSecuritySuitcaseMandatory field to given value.

### HasIsSecuritySuitcaseMandatory

`func (o *CreateAddressCommand) HasIsSecuritySuitcaseMandatory() bool`

HasIsSecuritySuitcaseMandatory returns a boolean if a field has been set.

### GetIsSideUnloading

`func (o *CreateAddressCommand) GetIsSideUnloading() bool`

GetIsSideUnloading returns the IsSideUnloading field if non-nil, zero value otherwise.

### GetIsSideUnloadingOk

`func (o *CreateAddressCommand) GetIsSideUnloadingOk() (*bool, bool)`

GetIsSideUnloadingOk returns a tuple with the IsSideUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSideUnloading

`func (o *CreateAddressCommand) SetIsSideUnloading(v bool)`

SetIsSideUnloading sets IsSideUnloading field to given value.

### HasIsSideUnloading

`func (o *CreateAddressCommand) HasIsSideUnloading() bool`

HasIsSideUnloading returns a boolean if a field has been set.

### GetIsCraneUnloading

`func (o *CreateAddressCommand) GetIsCraneUnloading() bool`

GetIsCraneUnloading returns the IsCraneUnloading field if non-nil, zero value otherwise.

### GetIsCraneUnloadingOk

`func (o *CreateAddressCommand) GetIsCraneUnloadingOk() (*bool, bool)`

GetIsCraneUnloadingOk returns a tuple with the IsCraneUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCraneUnloading

`func (o *CreateAddressCommand) SetIsCraneUnloading(v bool)`

SetIsCraneUnloading sets IsCraneUnloading field to given value.

### HasIsCraneUnloading

`func (o *CreateAddressCommand) HasIsCraneUnloading() bool`

HasIsCraneUnloading returns a boolean if a field has been set.

### GetGeocoding

`func (o *CreateAddressCommand) GetGeocoding() AddressGeocodingEditionDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *CreateAddressCommand) GetGeocodingOk() (*AddressGeocodingEditionDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *CreateAddressCommand) SetGeocoding(v AddressGeocodingEditionDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *CreateAddressCommand) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetEmails

`func (o *CreateAddressCommand) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateAddressCommand) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateAddressCommand) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CreateAddressCommand) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *CreateAddressCommand) GetAdditionalContacts() []CreateAddressCommandCreateAddressAdditionalContactDto`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *CreateAddressCommand) GetAdditionalContactsOk() (*[]CreateAddressCommandCreateAddressAdditionalContactDto, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *CreateAddressCommand) SetAdditionalContacts(v []CreateAddressCommandCreateAddressAdditionalContactDto)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *CreateAddressCommand) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetAddressOpenings

`func (o *CreateAddressCommand) GetAddressOpenings() []CreateAddressCommandCreateAddressOpeningDto`

GetAddressOpenings returns the AddressOpenings field if non-nil, zero value otherwise.

### GetAddressOpeningsOk

`func (o *CreateAddressCommand) GetAddressOpeningsOk() (*[]CreateAddressCommandCreateAddressOpeningDto, bool)`

GetAddressOpeningsOk returns a tuple with the AddressOpenings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOpenings

`func (o *CreateAddressCommand) SetAddressOpenings(v []CreateAddressCommandCreateAddressOpeningDto)`

SetAddressOpenings sets AddressOpenings field to given value.

### HasAddressOpenings

`func (o *CreateAddressCommand) HasAddressOpenings() bool`

HasAddressOpenings returns a boolean if a field has been set.

### GetForbiddenVehicleTypeCodes

`func (o *CreateAddressCommand) GetForbiddenVehicleTypeCodes() []string`

GetForbiddenVehicleTypeCodes returns the ForbiddenVehicleTypeCodes field if non-nil, zero value otherwise.

### GetForbiddenVehicleTypeCodesOk

`func (o *CreateAddressCommand) GetForbiddenVehicleTypeCodesOk() (*[]string, bool)`

GetForbiddenVehicleTypeCodesOk returns a tuple with the ForbiddenVehicleTypeCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenVehicleTypeCodes

`func (o *CreateAddressCommand) SetForbiddenVehicleTypeCodes(v []string)`

SetForbiddenVehicleTypeCodes sets ForbiddenVehicleTypeCodes field to given value.

### HasForbiddenVehicleTypeCodes

`func (o *CreateAddressCommand) HasForbiddenVehicleTypeCodes() bool`

HasForbiddenVehicleTypeCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


