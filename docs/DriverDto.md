# DriverDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverId** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DriverType** | Pointer to **string** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**TeamCode** | Pointer to **string** |  | [optional] 
**TeamLabel** | Pointer to **string** |  | [optional] 
**ProfileCode** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsDispatchMobileEnabled** | Pointer to **bool** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AlternativePhone** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**AddressId** | Pointer to **int32** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**CurrentVehicleCode** | Pointer to **string** |  | [optional] 
**DefaultVehicleCode** | Pointer to **string** |  | [optional] 
**DrivingLicenceType** | Pointer to **string** |  | [optional] 
**HasDangerousGoodsCertification** | Pointer to **bool** |  | [optional] 
**GlobalMark** | Pointer to **int32** |  | [optional] 
**Included** | Pointer to [**DriverIncludedDto**](DriverIncludedDto.md) |  | [optional] 
**SubcontractorCode** | Pointer to **string** |  | [optional] 
**RemoteAgencyCode** | Pointer to **string** |  | [optional] 
**RemoteLicenseCode** | Pointer to **string** |  | [optional] 
**SubcontractorEmployeesCount** | Pointer to **int64** |  | [optional] 
**ArrivalDateTime** | Pointer to **time.Time** |  | [optional] 
**DepartureDateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDriverDto

`func NewDriverDto() *DriverDto`

NewDriverDto instantiates a new DriverDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverDtoWithDefaults

`func NewDriverDtoWithDefaults() *DriverDto`

NewDriverDtoWithDefaults instantiates a new DriverDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverId

`func (o *DriverDto) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DriverDto) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DriverDto) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DriverDto) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetCode

`func (o *DriverDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DriverDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DriverDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DriverDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *DriverDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriverDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriverDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DriverDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDriverType

`func (o *DriverDto) GetDriverType() string`

GetDriverType returns the DriverType field if non-nil, zero value otherwise.

### GetDriverTypeOk

`func (o *DriverDto) GetDriverTypeOk() (*string, bool)`

GetDriverTypeOk returns a tuple with the DriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverType

`func (o *DriverDto) SetDriverType(v string)`

SetDriverType sets DriverType field to given value.

### HasDriverType

`func (o *DriverDto) HasDriverType() bool`

HasDriverType returns a boolean if a field has been set.

### GetAgencyCode

`func (o *DriverDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *DriverDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *DriverDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *DriverDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetTeamCode

`func (o *DriverDto) GetTeamCode() string`

GetTeamCode returns the TeamCode field if non-nil, zero value otherwise.

### GetTeamCodeOk

`func (o *DriverDto) GetTeamCodeOk() (*string, bool)`

GetTeamCodeOk returns a tuple with the TeamCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamCode

`func (o *DriverDto) SetTeamCode(v string)`

SetTeamCode sets TeamCode field to given value.

### HasTeamCode

`func (o *DriverDto) HasTeamCode() bool`

HasTeamCode returns a boolean if a field has been set.

### GetTeamLabel

`func (o *DriverDto) GetTeamLabel() string`

GetTeamLabel returns the TeamLabel field if non-nil, zero value otherwise.

### GetTeamLabelOk

`func (o *DriverDto) GetTeamLabelOk() (*string, bool)`

GetTeamLabelOk returns a tuple with the TeamLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamLabel

`func (o *DriverDto) SetTeamLabel(v string)`

SetTeamLabel sets TeamLabel field to given value.

### HasTeamLabel

`func (o *DriverDto) HasTeamLabel() bool`

HasTeamLabel returns a boolean if a field has been set.

### GetProfileCode

`func (o *DriverDto) GetProfileCode() string`

GetProfileCode returns the ProfileCode field if non-nil, zero value otherwise.

### GetProfileCodeOk

`func (o *DriverDto) GetProfileCodeOk() (*string, bool)`

GetProfileCodeOk returns a tuple with the ProfileCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileCode

`func (o *DriverDto) SetProfileCode(v string)`

SetProfileCode sets ProfileCode field to given value.

### HasProfileCode

`func (o *DriverDto) HasProfileCode() bool`

HasProfileCode returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DriverDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DriverDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DriverDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DriverDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsDispatchMobileEnabled

`func (o *DriverDto) GetIsDispatchMobileEnabled() bool`

GetIsDispatchMobileEnabled returns the IsDispatchMobileEnabled field if non-nil, zero value otherwise.

### GetIsDispatchMobileEnabledOk

`func (o *DriverDto) GetIsDispatchMobileEnabledOk() (*bool, bool)`

GetIsDispatchMobileEnabledOk returns a tuple with the IsDispatchMobileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDispatchMobileEnabled

`func (o *DriverDto) SetIsDispatchMobileEnabled(v bool)`

SetIsDispatchMobileEnabled sets IsDispatchMobileEnabled field to given value.

### HasIsDispatchMobileEnabled

`func (o *DriverDto) HasIsDispatchMobileEnabled() bool`

HasIsDispatchMobileEnabled returns a boolean if a field has been set.

### GetPhone

`func (o *DriverDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DriverDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DriverDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *DriverDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *DriverDto) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *DriverDto) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *DriverDto) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *DriverDto) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *DriverDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *DriverDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *DriverDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *DriverDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetEmail

`func (o *DriverDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DriverDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DriverDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DriverDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFax

`func (o *DriverDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *DriverDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *DriverDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *DriverDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetAddressId

`func (o *DriverDto) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *DriverDto) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *DriverDto) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *DriverDto) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetLocale

`func (o *DriverDto) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *DriverDto) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *DriverDto) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *DriverDto) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetCurrentVehicleCode

`func (o *DriverDto) GetCurrentVehicleCode() string`

GetCurrentVehicleCode returns the CurrentVehicleCode field if non-nil, zero value otherwise.

### GetCurrentVehicleCodeOk

`func (o *DriverDto) GetCurrentVehicleCodeOk() (*string, bool)`

GetCurrentVehicleCodeOk returns a tuple with the CurrentVehicleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVehicleCode

`func (o *DriverDto) SetCurrentVehicleCode(v string)`

SetCurrentVehicleCode sets CurrentVehicleCode field to given value.

### HasCurrentVehicleCode

`func (o *DriverDto) HasCurrentVehicleCode() bool`

HasCurrentVehicleCode returns a boolean if a field has been set.

### GetDefaultVehicleCode

`func (o *DriverDto) GetDefaultVehicleCode() string`

GetDefaultVehicleCode returns the DefaultVehicleCode field if non-nil, zero value otherwise.

### GetDefaultVehicleCodeOk

`func (o *DriverDto) GetDefaultVehicleCodeOk() (*string, bool)`

GetDefaultVehicleCodeOk returns a tuple with the DefaultVehicleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVehicleCode

`func (o *DriverDto) SetDefaultVehicleCode(v string)`

SetDefaultVehicleCode sets DefaultVehicleCode field to given value.

### HasDefaultVehicleCode

`func (o *DriverDto) HasDefaultVehicleCode() bool`

HasDefaultVehicleCode returns a boolean if a field has been set.

### GetDrivingLicenceType

`func (o *DriverDto) GetDrivingLicenceType() string`

GetDrivingLicenceType returns the DrivingLicenceType field if non-nil, zero value otherwise.

### GetDrivingLicenceTypeOk

`func (o *DriverDto) GetDrivingLicenceTypeOk() (*string, bool)`

GetDrivingLicenceTypeOk returns a tuple with the DrivingLicenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingLicenceType

`func (o *DriverDto) SetDrivingLicenceType(v string)`

SetDrivingLicenceType sets DrivingLicenceType field to given value.

### HasDrivingLicenceType

`func (o *DriverDto) HasDrivingLicenceType() bool`

HasDrivingLicenceType returns a boolean if a field has been set.

### GetHasDangerousGoodsCertification

`func (o *DriverDto) GetHasDangerousGoodsCertification() bool`

GetHasDangerousGoodsCertification returns the HasDangerousGoodsCertification field if non-nil, zero value otherwise.

### GetHasDangerousGoodsCertificationOk

`func (o *DriverDto) GetHasDangerousGoodsCertificationOk() (*bool, bool)`

GetHasDangerousGoodsCertificationOk returns a tuple with the HasDangerousGoodsCertification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDangerousGoodsCertification

`func (o *DriverDto) SetHasDangerousGoodsCertification(v bool)`

SetHasDangerousGoodsCertification sets HasDangerousGoodsCertification field to given value.

### HasHasDangerousGoodsCertification

`func (o *DriverDto) HasHasDangerousGoodsCertification() bool`

HasHasDangerousGoodsCertification returns a boolean if a field has been set.

### GetGlobalMark

`func (o *DriverDto) GetGlobalMark() int32`

GetGlobalMark returns the GlobalMark field if non-nil, zero value otherwise.

### GetGlobalMarkOk

`func (o *DriverDto) GetGlobalMarkOk() (*int32, bool)`

GetGlobalMarkOk returns a tuple with the GlobalMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalMark

`func (o *DriverDto) SetGlobalMark(v int32)`

SetGlobalMark sets GlobalMark field to given value.

### HasGlobalMark

`func (o *DriverDto) HasGlobalMark() bool`

HasGlobalMark returns a boolean if a field has been set.

### GetIncluded

`func (o *DriverDto) GetIncluded() DriverIncludedDto`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *DriverDto) GetIncludedOk() (*DriverIncludedDto, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *DriverDto) SetIncluded(v DriverIncludedDto)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *DriverDto) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetSubcontractorCode

`func (o *DriverDto) GetSubcontractorCode() string`

GetSubcontractorCode returns the SubcontractorCode field if non-nil, zero value otherwise.

### GetSubcontractorCodeOk

`func (o *DriverDto) GetSubcontractorCodeOk() (*string, bool)`

GetSubcontractorCodeOk returns a tuple with the SubcontractorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcontractorCode

`func (o *DriverDto) SetSubcontractorCode(v string)`

SetSubcontractorCode sets SubcontractorCode field to given value.

### HasSubcontractorCode

`func (o *DriverDto) HasSubcontractorCode() bool`

HasSubcontractorCode returns a boolean if a field has been set.

### GetRemoteAgencyCode

`func (o *DriverDto) GetRemoteAgencyCode() string`

GetRemoteAgencyCode returns the RemoteAgencyCode field if non-nil, zero value otherwise.

### GetRemoteAgencyCodeOk

`func (o *DriverDto) GetRemoteAgencyCodeOk() (*string, bool)`

GetRemoteAgencyCodeOk returns a tuple with the RemoteAgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAgencyCode

`func (o *DriverDto) SetRemoteAgencyCode(v string)`

SetRemoteAgencyCode sets RemoteAgencyCode field to given value.

### HasRemoteAgencyCode

`func (o *DriverDto) HasRemoteAgencyCode() bool`

HasRemoteAgencyCode returns a boolean if a field has been set.

### GetRemoteLicenseCode

`func (o *DriverDto) GetRemoteLicenseCode() string`

GetRemoteLicenseCode returns the RemoteLicenseCode field if non-nil, zero value otherwise.

### GetRemoteLicenseCodeOk

`func (o *DriverDto) GetRemoteLicenseCodeOk() (*string, bool)`

GetRemoteLicenseCodeOk returns a tuple with the RemoteLicenseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLicenseCode

`func (o *DriverDto) SetRemoteLicenseCode(v string)`

SetRemoteLicenseCode sets RemoteLicenseCode field to given value.

### HasRemoteLicenseCode

`func (o *DriverDto) HasRemoteLicenseCode() bool`

HasRemoteLicenseCode returns a boolean if a field has been set.

### GetSubcontractorEmployeesCount

`func (o *DriverDto) GetSubcontractorEmployeesCount() int64`

GetSubcontractorEmployeesCount returns the SubcontractorEmployeesCount field if non-nil, zero value otherwise.

### GetSubcontractorEmployeesCountOk

`func (o *DriverDto) GetSubcontractorEmployeesCountOk() (*int64, bool)`

GetSubcontractorEmployeesCountOk returns a tuple with the SubcontractorEmployeesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcontractorEmployeesCount

`func (o *DriverDto) SetSubcontractorEmployeesCount(v int64)`

SetSubcontractorEmployeesCount sets SubcontractorEmployeesCount field to given value.

### HasSubcontractorEmployeesCount

`func (o *DriverDto) HasSubcontractorEmployeesCount() bool`

HasSubcontractorEmployeesCount returns a boolean if a field has been set.

### GetArrivalDateTime

`func (o *DriverDto) GetArrivalDateTime() time.Time`

GetArrivalDateTime returns the ArrivalDateTime field if non-nil, zero value otherwise.

### GetArrivalDateTimeOk

`func (o *DriverDto) GetArrivalDateTimeOk() (*time.Time, bool)`

GetArrivalDateTimeOk returns a tuple with the ArrivalDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDateTime

`func (o *DriverDto) SetArrivalDateTime(v time.Time)`

SetArrivalDateTime sets ArrivalDateTime field to given value.

### HasArrivalDateTime

`func (o *DriverDto) HasArrivalDateTime() bool`

HasArrivalDateTime returns a boolean if a field has been set.

### GetDepartureDateTime

`func (o *DriverDto) GetDepartureDateTime() time.Time`

GetDepartureDateTime returns the DepartureDateTime field if non-nil, zero value otherwise.

### GetDepartureDateTimeOk

`func (o *DriverDto) GetDepartureDateTimeOk() (*time.Time, bool)`

GetDepartureDateTimeOk returns a tuple with the DepartureDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDateTime

`func (o *DriverDto) SetDepartureDateTime(v time.Time)`

SetDepartureDateTime sets DepartureDateTime field to given value.

### HasDepartureDateTime

`func (o *DriverDto) HasDepartureDateTime() bool`

HasDepartureDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


