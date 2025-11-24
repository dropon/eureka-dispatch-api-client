# AddressDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | Pointer to **int32** |  | [optional] 
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
**Geocoding** | Pointer to [**GeocodingDto**](GeocodingDto.md) |  | [optional] 
**City** | Pointer to [**CityDto**](CityDto.md) |  | [optional] 
**Emails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddressDto

`func NewAddressDto() *AddressDto`

NewAddressDto instantiates a new AddressDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDtoWithDefaults

`func NewAddressDtoWithDefaults() *AddressDto`

NewAddressDtoWithDefaults instantiates a new AddressDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *AddressDto) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *AddressDto) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *AddressDto) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *AddressDto) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetName

`func (o *AddressDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreetNumber

`func (o *AddressDto) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *AddressDto) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *AddressDto) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *AddressDto) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### GetStreetName

`func (o *AddressDto) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *AddressDto) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *AddressDto) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *AddressDto) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *AddressDto) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressDto) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressDto) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressDto) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *AddressDto) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressDto) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressDto) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressDto) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetContact

`func (o *AddressDto) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AddressDto) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AddressDto) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AddressDto) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhone

`func (o *AddressDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *AddressDto) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *AddressDto) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *AddressDto) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *AddressDto) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *AddressDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *AddressDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *AddressDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *AddressDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetFax

`func (o *AddressDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *AddressDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *AddressDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *AddressDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetAgencyCode

`func (o *AddressDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *AddressDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *AddressDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *AddressDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetAddressCategoryCode

`func (o *AddressDto) GetAddressCategoryCode() string`

GetAddressCategoryCode returns the AddressCategoryCode field if non-nil, zero value otherwise.

### GetAddressCategoryCodeOk

`func (o *AddressDto) GetAddressCategoryCodeOk() (*string, bool)`

GetAddressCategoryCodeOk returns a tuple with the AddressCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCategoryCode

`func (o *AddressDto) SetAddressCategoryCode(v string)`

SetAddressCategoryCode sets AddressCategoryCode field to given value.

### HasAddressCategoryCode

`func (o *AddressDto) HasAddressCategoryCode() bool`

HasAddressCategoryCode returns a boolean if a field has been set.

### GetExternalAddressCode

`func (o *AddressDto) GetExternalAddressCode() string`

GetExternalAddressCode returns the ExternalAddressCode field if non-nil, zero value otherwise.

### GetExternalAddressCodeOk

`func (o *AddressDto) GetExternalAddressCodeOk() (*string, bool)`

GetExternalAddressCodeOk returns a tuple with the ExternalAddressCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddressCode

`func (o *AddressDto) SetExternalAddressCode(v string)`

SetExternalAddressCode sets ExternalAddressCode field to given value.

### HasExternalAddressCode

`func (o *AddressDto) HasExternalAddressCode() bool`

HasExternalAddressCode returns a boolean if a field has been set.

### GetNotes

`func (o *AddressDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AddressDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AddressDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AddressDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAddressComplement1

`func (o *AddressDto) GetAddressComplement1() string`

GetAddressComplement1 returns the AddressComplement1 field if non-nil, zero value otherwise.

### GetAddressComplement1Ok

`func (o *AddressDto) GetAddressComplement1Ok() (*string, bool)`

GetAddressComplement1Ok returns a tuple with the AddressComplement1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement1

`func (o *AddressDto) SetAddressComplement1(v string)`

SetAddressComplement1 sets AddressComplement1 field to given value.

### HasAddressComplement1

`func (o *AddressDto) HasAddressComplement1() bool`

HasAddressComplement1 returns a boolean if a field has been set.

### GetAddressComplement2

`func (o *AddressDto) GetAddressComplement2() string`

GetAddressComplement2 returns the AddressComplement2 field if non-nil, zero value otherwise.

### GetAddressComplement2Ok

`func (o *AddressDto) GetAddressComplement2Ok() (*string, bool)`

GetAddressComplement2Ok returns a tuple with the AddressComplement2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement2

`func (o *AddressDto) SetAddressComplement2(v string)`

SetAddressComplement2 sets AddressComplement2 field to given value.

### HasAddressComplement2

`func (o *AddressDto) HasAddressComplement2() bool`

HasAddressComplement2 returns a boolean if a field has been set.

### GetAddressComplement3

`func (o *AddressDto) GetAddressComplement3() string`

GetAddressComplement3 returns the AddressComplement3 field if non-nil, zero value otherwise.

### GetAddressComplement3Ok

`func (o *AddressDto) GetAddressComplement3Ok() (*string, bool)`

GetAddressComplement3Ok returns a tuple with the AddressComplement3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement3

`func (o *AddressDto) SetAddressComplement3(v string)`

SetAddressComplement3 sets AddressComplement3 field to given value.

### HasAddressComplement3

`func (o *AddressDto) HasAddressComplement3() bool`

HasAddressComplement3 returns a boolean if a field has been set.

### GetAddressComplement4

`func (o *AddressDto) GetAddressComplement4() string`

GetAddressComplement4 returns the AddressComplement4 field if non-nil, zero value otherwise.

### GetAddressComplement4Ok

`func (o *AddressDto) GetAddressComplement4Ok() (*string, bool)`

GetAddressComplement4Ok returns a tuple with the AddressComplement4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressComplement4

`func (o *AddressDto) SetAddressComplement4(v string)`

SetAddressComplement4 sets AddressComplement4 field to given value.

### HasAddressComplement4

`func (o *AddressDto) HasAddressComplement4() bool`

HasAddressComplement4 returns a boolean if a field has been set.

### GetIsHeightRestrictedAccess

`func (o *AddressDto) GetIsHeightRestrictedAccess() bool`

GetIsHeightRestrictedAccess returns the IsHeightRestrictedAccess field if non-nil, zero value otherwise.

### GetIsHeightRestrictedAccessOk

`func (o *AddressDto) GetIsHeightRestrictedAccessOk() (*bool, bool)`

GetIsHeightRestrictedAccessOk returns a tuple with the IsHeightRestrictedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeightRestrictedAccess

`func (o *AddressDto) SetIsHeightRestrictedAccess(v bool)`

SetIsHeightRestrictedAccess sets IsHeightRestrictedAccess field to given value.

### HasIsHeightRestrictedAccess

`func (o *AddressDto) HasIsHeightRestrictedAccess() bool`

HasIsHeightRestrictedAccess returns a boolean if a field has been set.

### GetHeightLimitation

`func (o *AddressDto) GetHeightLimitation() int32`

GetHeightLimitation returns the HeightLimitation field if non-nil, zero value otherwise.

### GetHeightLimitationOk

`func (o *AddressDto) GetHeightLimitationOk() (*int32, bool)`

GetHeightLimitationOk returns a tuple with the HeightLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightLimitation

`func (o *AddressDto) SetHeightLimitation(v int32)`

SetHeightLimitation sets HeightLimitation field to given value.

### HasHeightLimitation

`func (o *AddressDto) HasHeightLimitation() bool`

HasHeightLimitation returns a boolean if a field has been set.

### GetIsHatchBackMandatory

`func (o *AddressDto) GetIsHatchBackMandatory() bool`

GetIsHatchBackMandatory returns the IsHatchBackMandatory field if non-nil, zero value otherwise.

### GetIsHatchBackMandatoryOk

`func (o *AddressDto) GetIsHatchBackMandatoryOk() (*bool, bool)`

GetIsHatchBackMandatoryOk returns a tuple with the IsHatchBackMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHatchBackMandatory

`func (o *AddressDto) SetIsHatchBackMandatory(v bool)`

SetIsHatchBackMandatory sets IsHatchBackMandatory field to given value.

### HasIsHatchBackMandatory

`func (o *AddressDto) HasIsHatchBackMandatory() bool`

HasIsHatchBackMandatory returns a boolean if a field has been set.

### GetIsWarehousemanMandatory

`func (o *AddressDto) GetIsWarehousemanMandatory() bool`

GetIsWarehousemanMandatory returns the IsWarehousemanMandatory field if non-nil, zero value otherwise.

### GetIsWarehousemanMandatoryOk

`func (o *AddressDto) GetIsWarehousemanMandatoryOk() (*bool, bool)`

GetIsWarehousemanMandatoryOk returns a tuple with the IsWarehousemanMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWarehousemanMandatory

`func (o *AddressDto) SetIsWarehousemanMandatory(v bool)`

SetIsWarehousemanMandatory sets IsWarehousemanMandatory field to given value.

### HasIsWarehousemanMandatory

`func (o *AddressDto) HasIsWarehousemanMandatory() bool`

HasIsWarehousemanMandatory returns a boolean if a field has been set.

### GetIsHandlingEquipmentMandatory

`func (o *AddressDto) GetIsHandlingEquipmentMandatory() bool`

GetIsHandlingEquipmentMandatory returns the IsHandlingEquipmentMandatory field if non-nil, zero value otherwise.

### GetIsHandlingEquipmentMandatoryOk

`func (o *AddressDto) GetIsHandlingEquipmentMandatoryOk() (*bool, bool)`

GetIsHandlingEquipmentMandatoryOk returns a tuple with the IsHandlingEquipmentMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHandlingEquipmentMandatory

`func (o *AddressDto) SetIsHandlingEquipmentMandatory(v bool)`

SetIsHandlingEquipmentMandatory sets IsHandlingEquipmentMandatory field to given value.

### HasIsHandlingEquipmentMandatory

`func (o *AddressDto) HasIsHandlingEquipmentMandatory() bool`

HasIsHandlingEquipmentMandatory returns a boolean if a field has been set.

### GetIsSecuritySuitcaseMandatory

`func (o *AddressDto) GetIsSecuritySuitcaseMandatory() bool`

GetIsSecuritySuitcaseMandatory returns the IsSecuritySuitcaseMandatory field if non-nil, zero value otherwise.

### GetIsSecuritySuitcaseMandatoryOk

`func (o *AddressDto) GetIsSecuritySuitcaseMandatoryOk() (*bool, bool)`

GetIsSecuritySuitcaseMandatoryOk returns a tuple with the IsSecuritySuitcaseMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecuritySuitcaseMandatory

`func (o *AddressDto) SetIsSecuritySuitcaseMandatory(v bool)`

SetIsSecuritySuitcaseMandatory sets IsSecuritySuitcaseMandatory field to given value.

### HasIsSecuritySuitcaseMandatory

`func (o *AddressDto) HasIsSecuritySuitcaseMandatory() bool`

HasIsSecuritySuitcaseMandatory returns a boolean if a field has been set.

### GetIsSideUnloading

`func (o *AddressDto) GetIsSideUnloading() bool`

GetIsSideUnloading returns the IsSideUnloading field if non-nil, zero value otherwise.

### GetIsSideUnloadingOk

`func (o *AddressDto) GetIsSideUnloadingOk() (*bool, bool)`

GetIsSideUnloadingOk returns a tuple with the IsSideUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSideUnloading

`func (o *AddressDto) SetIsSideUnloading(v bool)`

SetIsSideUnloading sets IsSideUnloading field to given value.

### HasIsSideUnloading

`func (o *AddressDto) HasIsSideUnloading() bool`

HasIsSideUnloading returns a boolean if a field has been set.

### GetIsCraneUnloading

`func (o *AddressDto) GetIsCraneUnloading() bool`

GetIsCraneUnloading returns the IsCraneUnloading field if non-nil, zero value otherwise.

### GetIsCraneUnloadingOk

`func (o *AddressDto) GetIsCraneUnloadingOk() (*bool, bool)`

GetIsCraneUnloadingOk returns a tuple with the IsCraneUnloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCraneUnloading

`func (o *AddressDto) SetIsCraneUnloading(v bool)`

SetIsCraneUnloading sets IsCraneUnloading field to given value.

### HasIsCraneUnloading

`func (o *AddressDto) HasIsCraneUnloading() bool`

HasIsCraneUnloading returns a boolean if a field has been set.

### GetGeocoding

`func (o *AddressDto) GetGeocoding() GeocodingDto`

GetGeocoding returns the Geocoding field if non-nil, zero value otherwise.

### GetGeocodingOk

`func (o *AddressDto) GetGeocodingOk() (*GeocodingDto, bool)`

GetGeocodingOk returns a tuple with the Geocoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocoding

`func (o *AddressDto) SetGeocoding(v GeocodingDto)`

SetGeocoding sets Geocoding field to given value.

### HasGeocoding

`func (o *AddressDto) HasGeocoding() bool`

HasGeocoding returns a boolean if a field has been set.

### GetCity

`func (o *AddressDto) GetCity() CityDto`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressDto) GetCityOk() (*CityDto, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressDto) SetCity(v CityDto)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetEmails

`func (o *AddressDto) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AddressDto) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AddressDto) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AddressDto) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


