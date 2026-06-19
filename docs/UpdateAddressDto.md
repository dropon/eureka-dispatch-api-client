# UpdateAddressDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StreetNumber** | Pointer to **string** |  | [optional] 
**StreetName** | Pointer to **string** |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AlternativePhone** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
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
**CityId** | Pointer to **int32** |  | [optional] 
**Emails** | Pointer to **[]string** | The emails are replaced when provided in the update. When not provided, the existing emails are kept unchanged. | [optional] 
**AdditionalContacts** | Pointer to [**[]UpdateAddressDtoUpdateAddressAdditionalContactDto**](UpdateAddressDtoUpdateAddressAdditionalContactDto.md) |  | [optional] 
**AddressOpenings** | Pointer to [**[]UpdateAddressDtoUpdateAddressOpeningDto**](UpdateAddressDtoUpdateAddressOpeningDto.md) | Address openings are replaced when provided in the update. When not provided, the existing address openings are kept unchanged. | [optional] 
**ForbiddenVehicleTypeCodes** | Pointer to **[]string** | Forbidden vehicle type codes are replaced when provided in the update. When not provided, the existing forbidden vehicle types are kept unchanged. | [optional] 

## Methods

### NewUpdateAddressDto

`func NewUpdateAddressDto() *UpdateAddressDto`

NewUpdateAddressDto instantiates a new UpdateAddressDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAddressDtoWithDefaults

`func NewUpdateAddressDtoWithDefaults() *UpdateAddressDto`

NewUpdateAddressDtoWithDefaults instantiates a new UpdateAddressDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAddressDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAddressDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAddressDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAddressDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetNumber

`func (o *UpdateAddressDto) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *UpdateAddressDto) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *UpdateAddressDto) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *UpdateAddressDto) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *UpdateAddressDto) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *UpdateAddressDto) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *UpdateAddressDto) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *UpdateAddressDto) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *UpdateAddressDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *UpdateAddressDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *UpdateAddressDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *UpdateAddressDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *UpdateAddressDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *UpdateAddressDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *UpdateAddressDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *UpdateAddressDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetContact

`func (o *UpdateAddressDto) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *UpdateAddressDto) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *UpdateAddressDto) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *UpdateAddressDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateAddressDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateAddressDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateAddressDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateAddressDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *UpdateAddressDto) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *UpdateAddressDto) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *UpdateAddressDto) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *UpdateAddressDto) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *UpdateAddressDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *UpdateAddressDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *UpdateAddressDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *UpdateAddressDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetFax

`func (o *UpdateAddressDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *UpdateAddressDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *UpdateAddressDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *UpdateAddressDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetAgencyCode

`func (o *UpdateAddressDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *UpdateAddressDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *UpdateAddressDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *UpdateAddressDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetAddressCategoryCode

`func (o *UpdateAddressDto) GetAddressCategoryCode() string`

GetAddressCategoryCode returns the AddressCategoryCode field if non-nil, zero value otherwise.

### GetAddressCategoryCodeOk

`func (o *UpdateAddressDto) GetAddressCategoryCodeOk() (*string, bool)`

GetAddressCategoryCodeOk returns a tuple with the AddressCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCategoryCode

`func (o *UpdateAddressDto) SetAddressCategoryCode(v string)`

SetAddressCategoryCode sets AddressCategoryCode field to given value.

### HasAddressCategoryCode

`func (o *UpdateAddressDto) HasAddressCategoryCode() bool`

HasAddressCategoryCode returns a boolean if a field has been set.

### GetExternalAddressCode

`func (o *UpdateAddressDto) GetExternalAddressCode() string`

GetExternalAddressCode returns the ExternalAddressCode field if non-nil, zero value otherwise.

### GetExternalAddressCodeOk

`func (o *UpdateAddressDto) GetExternalAddressCodeOk() (*string, bool)`

GetExternalAddressCodeOk returns a tuple with the ExternalAddressCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddressCode

`func (o *UpdateAddressDto) SetExternalAddressCode(v string)`

SetExternalAddressCode sets ExternalAddressCode field to given value.

### HasExternalAddressCode

`func (o *UpdateAddressDto) HasExternalAddressCode() bool`

HasExternalAddressCode returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateAddressDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateAddressDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateAddressDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateAddressDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStopDurationInSeconds

`func (o *UpdateAddressDto) GetStopDurationInSeconds() int32`

GetStopDurationInSeconds returns the StopDurationInSeconds field if non-nil, zero value otherwise.

### GetStopDurationInSecondsOk

`func (o *UpdateAddressDto) GetStopDurationInSecondsOk() (*int32, bool)`

GetStopDurationInSecondsOk returns a tuple with the StopDurationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDurationInSeconds

`func (o *UpdateAddressDto) SetStopDurationInSeconds(v int32)`

SetStopDurationInSeconds sets StopDurationInSeconds field to given value.

### HasStopDurationInSeconds

`func (o *UpdateAddressDto) HasStopDurationInSeconds() bool`

HasStopDurationInSeconds returns a boolean if a field has been set.

### GetCircuitAddressSetId

`func (o *UpdateAddressDto) GetCircuitAddressSetId() int32`

GetCircuitAddressSetId returns the CircuitAddressSetId field if non-nil, zero value otherwise.

### GetCircuitAddressSetIdOk

`func (o *UpdateAddressDto) GetCircuitAddressSetIdOk() (*int32, bool)`

GetCircuitAddressSetIdOk returns a tuple with the CircuitAddressSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitAddressSetId

`func (o *UpdateAddressDto) SetCircuitAddressSetId(v int32)`

SetCircuitAddressSetId sets CircuitAddressSetId field to given value.

### HasCircuitAddressSetId

`func (o *UpdateAddressDto) HasCircuitAddressSetId() bool`

HasCircuitAddressSetId returns a boolean if a field has been set.

### GetAddressComplement1

`func (o *UpdateAddressDto) GetAddressComplement1() string`

GetAddressComplement1 returns the AddressComplement1 field if non-nil, zero value otherwise.

### GetAddressComplement1Ok

`func (o *UpdateAddressDto) GetAddressComplement1Ok() (*string, bool)`

GetAddressComplement1Ok returns a tuple with the AddressComplement1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement1

`func (o *UpdateAddressDto) SetAddressComplement1(v string)`

SetAddressComplement1 sets AddressComplement1 field to given value.

### HasAddressComplement1

`func (o *UpdateAddressDto) HasAddressComplement1() bool`

HasAddressComplement1 returns a boolean if a field has been set.

### GetAddressComplement2

`func (o *UpdateAddressDto) GetAddressComplement2() string`

GetAddressComplement2 returns the AddressComplement2 field if non-nil, zero value otherwise.

### GetAddressComplement2Ok

`func (o *UpdateAddressDto) GetAddressComplement2Ok() (*string, bool)`

GetAddressComplement2Ok returns a tuple with the AddressComplement2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement2

`func (o *UpdateAddressDto) SetAddressComplement2(v string)`

SetAddressComplement2 sets AddressComplement2 field to given value.

### HasAddressComplement2

`func (o *UpdateAddressDto) HasAddressComplement2() bool`

HasAddressComplement2 returns a boolean if a field has been set.

### GetAddressComplement3

`func (o *UpdateAddressDto) GetAddressComplement3() string`

GetAddressComplement3 returns the AddressComplement3 field if non-nil, zero value otherwise.

### GetAddressComplement3Ok

`func (o *UpdateAddressDto) GetAddressComplement3Ok() (*string, bool)`

GetAddressComplement3Ok returns a tuple with the AddressComplement3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement3

`func (o *UpdateAddressDto) SetAddressComplement3(v string)`

SetAddressComplement3 sets AddressComplement3 field to given value.

### HasAddressComplement3

`func (o *UpdateAddressDto) HasAddressComplement3() bool`

HasAddressComplement3 returns a boolean if a field has been set.

### GetAddressComplement4

`func (o *UpdateAddressDto) GetAddressComplement4() string`

GetAddressComplement4 returns the AddressComplement4 field if non-nil, zero value otherwise.

### GetAddressComplement4Ok

`func (o *UpdateAddressDto) GetAddressComplement4Ok() (*string, bool)`

GetAddressComplement4Ok returns a tuple with the AddressComplement4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement4

`func (o *UpdateAddressDto) SetAddressComplement4(v string)`

SetAddressComplement4 sets AddressComplement4 field to given value.

### HasAddressComplement4

`func (o *UpdateAddressDto) HasAddressComplement4() bool`

HasAddressComplement4 returns a boolean if a field has been set.

### GetIsHeightRestrictedAccess

`func (o *UpdateAddressDto) GetIsHeightRestrictedAccess() bool`

GetIsHeightRestrictedAccess returns the IsHeightRestrictedAccess field if non-nil, zero value otherwise.

### GetIsHeightRestrictedAccessOk

`func (o *UpdateAddressDto) GetIsHeightRestrictedAccessOk() (*bool, bool)`

GetIsHeightRestrictedAccessOk returns a tuple with the IsHeightRestrictedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeightRestrictedAccess

`func (o *UpdateAddressDto) SetIsHeightRestrictedAccess(v bool)`

SetIsHeightRestrictedAccess sets IsHeightRestrictedAccess field to given value.

### HasIsHeightRestrictedAccess

`func (o *UpdateAddressDto) HasIsHeightRestrictedAccess() bool`

HasIsHeightRestrictedAccess returns a boolean if a field has been set.

### GetHeightLimitation

`func (o *UpdateAddressDto) GetHeightLimitation() int32`

GetHeightLimitation returns the HeightLimitation field if non-nil, zero value otherwise.

### GetHeightLimitationOk

`func (o *UpdateAddressDto) GetHeightLimitationOk() (*int32, bool)`

GetHeightLimitationOk returns a tuple with the HeightLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightLimitation

`func (o *UpdateAddressDto) SetHeightLimitation(v int32)`

SetHeightLimitation sets HeightLimitation field to given value.

### HasHeightLimitation

`func (o *UpdateAddressDto) HasHeightLimitation() bool`

HasHeightLimitation returns a boolean if a field has been set.

### GetIsHatchBackMandatory

`func (o *UpdateAddressDto) GetIsHatchBackMandatory() bool`

GetIsHatchBackMandatory returns the IsHatchBackMandatory field if non-nil, zero value otherwise.

### GetIsHatchBackMandatoryOk

`func (o *UpdateAddressDto) GetIsHatchBackMandatoryOk() (*bool, bool)`

GetIsHatchBackMandatoryOk returns a tuple with the IsHatchBackMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHatchBackMandatory

`func (o *UpdateAddressDto) SetIsHatchBackMandatory(v bool)`

SetIsHatchBackMandatory sets IsHatchBackMandatory field to given value.

### HasIsHatchBackMandatory

`func (o *UpdateAddressDto) HasIsHatchBackMandatory() bool`

HasIsHatchBackMandatory returns a boolean if a field has been set.

### GetIsWarehousemanMandatory

`func (o *UpdateAddressDto) GetIsWarehousemanMandatory() bool`

GetIsWarehousemanMandatory returns the IsWarehousemanMandatory field if non-nil, zero value otherwise.

### GetIsWarehousemanMandatoryOk

`func (o *UpdateAddressDto) GetIsWarehousemanMandatoryOk() (*bool, bool)`

GetIsWarehousemanMandatoryOk returns a tuple with the IsWarehousemanMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWarehousemanMandatory

`func (o *UpdateAddressDto) SetIsWarehousemanMandatory(v bool)`

SetIsWarehousemanMandatory sets IsWarehousemanMandatory field to given value.

### HasIsWarehousemanMandatory

`func (o *UpdateAddressDto) HasIsWarehousemanMandatory() bool`

HasIsWarehousemanMandatory returns a boolean if a field has been set.

### GetIsHandlingEquipmentMandatory

`func (o *UpdateAddressDto) GetIsHandlingEquipmentMandatory() bool`

GetIsHandlingEquipmentMandatory returns the IsHandlingEquipmentMandatory field if non-nil, zero value otherwise.

### GetIsHandlingEquipmentMandatoryOk

`func (o *UpdateAddressDto) GetIsHandlingEquipmentMandatoryOk() (*bool, bool)`

GetIsHandlingEquipmentMandatoryOk returns a tuple with the IsHandlingEquipmentMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHandlingEquipmentMandatory

`func (o *UpdateAddressDto) SetIsHandlingEquipmentMandatory(v bool)`

SetIsHandlingEquipmentMandatory sets IsHandlingEquipmentMandatory field to given value.

### HasIsHandlingEquipmentMandatory

`func (o *UpdateAddressDto) HasIsHandlingEquipmentMandatory() bool`

HasIsHandlingEquipmentMandatory returns a boolean if a field has been set.

### GetIsSecuritySuitcaseMandatory

`func (o *UpdateAddressDto) GetIsSecuritySuitcaseMandatory() bool`

GetIsSecuritySuitcaseMandatory returns the IsSecuritySuitcaseMandatory field if non-nil, zero value otherwise.

### GetIsSecuritySuitcaseMandatoryOk

`func (o *UpdateAddressDto) GetIsSecuritySuitcaseMandatoryOk() (*bool, bool)`

GetIsSecuritySuitcaseMandatoryOk returns a tuple with the IsSecuritySuitcaseMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySuitcaseMandatory

`func (o *UpdateAddressDto) SetIsSecuritySuitcaseMandatory(v bool)`

SetIsSecuritySuitcaseMandatory sets IsSecuritySuitcaseMandatory field to given value.

### HasIsSecuritySuitcaseMandatory

`func (o *UpdateAddressDto) HasIsSecuritySuitcaseMandatory() bool`

HasIsSecuritySuitcaseMandatory returns a boolean if a field has been set.

### GetIsSideUnloading

`func (o *UpdateAddressDto) GetIsSideUnloading() bool`

GetIsSideUnloading returns the IsSideUnloading field if non-nil, zero value otherwise.

### GetIsSideUnloadingOk

`func (o *UpdateAddressDto) GetIsSideUnloadingOk() (*bool, bool)`

GetIsSideUnloadingOk returns a tuple with the IsSideUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSideUnloading

`func (o *UpdateAddressDto) SetIsSideUnloading(v bool)`

SetIsSideUnloading sets IsSideUnloading field to given value.

### HasIsSideUnloading

`func (o *UpdateAddressDto) HasIsSideUnloading() bool`

HasIsSideUnloading returns a boolean if a field has been set.

### GetIsCraneUnloading

`func (o *UpdateAddressDto) GetIsCraneUnloading() bool`

GetIsCraneUnloading returns the IsCraneUnloading field if non-nil, zero value otherwise.

### GetIsCraneUnloadingOk

`func (o *UpdateAddressDto) GetIsCraneUnloadingOk() (*bool, bool)`

GetIsCraneUnloadingOk returns a tuple with the IsCraneUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCraneUnloading

`func (o *UpdateAddressDto) SetIsCraneUnloading(v bool)`

SetIsCraneUnloading sets IsCraneUnloading field to given value.

### HasIsCraneUnloading

`func (o *UpdateAddressDto) HasIsCraneUnloading() bool`

HasIsCraneUnloading returns a boolean if a field has been set.

### GetGeocoding

`func (o *UpdateAddressDto) GetGeocoding() AddressGeocodingEditionDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *UpdateAddressDto) GetGeocodingOk() (*AddressGeocodingEditionDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *UpdateAddressDto) SetGeocoding(v AddressGeocodingEditionDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *UpdateAddressDto) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetCityId

`func (o *UpdateAddressDto) GetCityId() int32`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *UpdateAddressDto) GetCityIdOk() (*int32, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *UpdateAddressDto) SetCityId(v int32)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *UpdateAddressDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### GetEmails

`func (o *UpdateAddressDto) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UpdateAddressDto) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UpdateAddressDto) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UpdateAddressDto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetAdditionalContacts

`func (o *UpdateAddressDto) GetAdditionalContacts() []UpdateAddressDtoUpdateAddressAdditionalContactDto`

GetAdditionalContacts returns the AdditionalContacts field if non-nil, zero value otherwise.

### GetAdditionalContactsOk

`func (o *UpdateAddressDto) GetAdditionalContactsOk() (*[]UpdateAddressDtoUpdateAddressAdditionalContactDto, bool)`

GetAdditionalContactsOk returns a tuple with the AdditionalContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContacts

`func (o *UpdateAddressDto) SetAdditionalContacts(v []UpdateAddressDtoUpdateAddressAdditionalContactDto)`

SetAdditionalContacts sets AdditionalContacts field to given value.

### HasAdditionalContacts

`func (o *UpdateAddressDto) HasAdditionalContacts() bool`

HasAdditionalContacts returns a boolean if a field has been set.

### GetAddressOpenings

`func (o *UpdateAddressDto) GetAddressOpenings() []UpdateAddressDtoUpdateAddressOpeningDto`

GetAddressOpenings returns the AddressOpenings field if non-nil, zero value otherwise.

### GetAddressOpeningsOk

`func (o *UpdateAddressDto) GetAddressOpeningsOk() (*[]UpdateAddressDtoUpdateAddressOpeningDto, bool)`

GetAddressOpeningsOk returns a tuple with the AddressOpenings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOpenings

`func (o *UpdateAddressDto) SetAddressOpenings(v []UpdateAddressDtoUpdateAddressOpeningDto)`

SetAddressOpenings sets AddressOpenings field to given value.

### HasAddressOpenings

`func (o *UpdateAddressDto) HasAddressOpenings() bool`

HasAddressOpenings returns a boolean if a field has been set.

### GetForbiddenVehicleTypeCodes

`func (o *UpdateAddressDto) GetForbiddenVehicleTypeCodes() []string`

GetForbiddenVehicleTypeCodes returns the ForbiddenVehicleTypeCodes field if non-nil, zero value otherwise.

### GetForbiddenVehicleTypeCodesOk

`func (o *UpdateAddressDto) GetForbiddenVehicleTypeCodesOk() (*[]string, bool)`

GetForbiddenVehicleTypeCodesOk returns a tuple with the ForbiddenVehicleTypeCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenVehicleTypeCodes

`func (o *UpdateAddressDto) SetForbiddenVehicleTypeCodes(v []string)`

SetForbiddenVehicleTypeCodes sets ForbiddenVehicleTypeCodes field to given value.

### HasForbiddenVehicleTypeCodes

`func (o *UpdateAddressDto) HasForbiddenVehicleTypeCodes() bool`

HasForbiddenVehicleTypeCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


