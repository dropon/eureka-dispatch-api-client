# VehicleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**HasHatchback** | Pointer to **bool** |  | [optional] 
**HasDangerousGoodsCase** | Pointer to **bool** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**TypeCode** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**LicensePlate** | Pointer to **string** |  | [optional] 
**TrailerUid** | Pointer to **string** |  | [optional] 
**TrailerCode** | Pointer to **string** |  | [optional] 
**TrailerLabel** | Pointer to **string** |  | [optional] 
**TrailerLicensePlate** | Pointer to **string** |  | [optional] 
**IsSubcontractorsSpecific** | Pointer to **bool** | Indicate if the vehicle is specific to one or many subcontractors | [optional] 

## Methods

### NewVehicleDto

`func NewVehicleDto() *VehicleDto`

NewVehicleDto instantiates a new VehicleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleDtoWithDefaults

`func NewVehicleDtoWithDefaults() *VehicleDto`

NewVehicleDtoWithDefaults instantiates a new VehicleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *VehicleDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *VehicleDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *VehicleDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *VehicleDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *VehicleDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VehicleDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VehicleDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VehicleDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *VehicleDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VehicleDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VehicleDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VehicleDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsEnabled

`func (o *VehicleDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *VehicleDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *VehicleDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *VehicleDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetHasHatchback

`func (o *VehicleDto) GetHasHatchback() bool`

GetHasHatchback returns the HasHatchback field if non-nil, zero value otherwise.

### GetHasHatchbackOk

`func (o *VehicleDto) GetHasHatchbackOk() (*bool, bool)`

GetHasHatchbackOk returns a tuple with the HasHatchback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHatchback

`func (o *VehicleDto) SetHasHatchback(v bool)`

SetHasHatchback sets HasHatchback field to given value.

### HasHasHatchback

`func (o *VehicleDto) HasHasHatchback() bool`

HasHasHatchback returns a boolean if a field has been set.

### GetHasDangerousGoodsCase

`func (o *VehicleDto) GetHasDangerousGoodsCase() bool`

GetHasDangerousGoodsCase returns the HasDangerousGoodsCase field if non-nil, zero value otherwise.

### GetHasDangerousGoodsCaseOk

`func (o *VehicleDto) GetHasDangerousGoodsCaseOk() (*bool, bool)`

GetHasDangerousGoodsCaseOk returns a tuple with the HasDangerousGoodsCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDangerousGoodsCase

`func (o *VehicleDto) SetHasDangerousGoodsCase(v bool)`

SetHasDangerousGoodsCase sets HasDangerousGoodsCase field to given value.

### HasHasDangerousGoodsCase

`func (o *VehicleDto) HasHasDangerousGoodsCase() bool`

HasHasDangerousGoodsCase returns a boolean if a field has been set.

### GetAgencyCode

`func (o *VehicleDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *VehicleDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *VehicleDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *VehicleDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetTypeCode

`func (o *VehicleDto) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *VehicleDto) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *VehicleDto) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *VehicleDto) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetBrand

`func (o *VehicleDto) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *VehicleDto) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *VehicleDto) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *VehicleDto) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetLicensePlate

`func (o *VehicleDto) GetLicensePlate() string`

GetLicensePlate returns the LicensePlate field if non-nil, zero value otherwise.

### GetLicensePlateOk

`func (o *VehicleDto) GetLicensePlateOk() (*string, bool)`

GetLicensePlateOk returns a tuple with the LicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlate

`func (o *VehicleDto) SetLicensePlate(v string)`

SetLicensePlate sets LicensePlate field to given value.

### HasLicensePlate

`func (o *VehicleDto) HasLicensePlate() bool`

HasLicensePlate returns a boolean if a field has been set.

### GetTrailerUid

`func (o *VehicleDto) GetTrailerUid() string`

GetTrailerUid returns the TrailerUid field if non-nil, zero value otherwise.

### GetTrailerUidOk

`func (o *VehicleDto) GetTrailerUidOk() (*string, bool)`

GetTrailerUidOk returns a tuple with the TrailerUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerUid

`func (o *VehicleDto) SetTrailerUid(v string)`

SetTrailerUid sets TrailerUid field to given value.

### HasTrailerUid

`func (o *VehicleDto) HasTrailerUid() bool`

HasTrailerUid returns a boolean if a field has been set.

### GetTrailerCode

`func (o *VehicleDto) GetTrailerCode() string`

GetTrailerCode returns the TrailerCode field if non-nil, zero value otherwise.

### GetTrailerCodeOk

`func (o *VehicleDto) GetTrailerCodeOk() (*string, bool)`

GetTrailerCodeOk returns a tuple with the TrailerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerCode

`func (o *VehicleDto) SetTrailerCode(v string)`

SetTrailerCode sets TrailerCode field to given value.

### HasTrailerCode

`func (o *VehicleDto) HasTrailerCode() bool`

HasTrailerCode returns a boolean if a field has been set.

### GetTrailerLabel

`func (o *VehicleDto) GetTrailerLabel() string`

GetTrailerLabel returns the TrailerLabel field if non-nil, zero value otherwise.

### GetTrailerLabelOk

`func (o *VehicleDto) GetTrailerLabelOk() (*string, bool)`

GetTrailerLabelOk returns a tuple with the TrailerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLabel

`func (o *VehicleDto) SetTrailerLabel(v string)`

SetTrailerLabel sets TrailerLabel field to given value.

### HasTrailerLabel

`func (o *VehicleDto) HasTrailerLabel() bool`

HasTrailerLabel returns a boolean if a field has been set.

### GetTrailerLicensePlate

`func (o *VehicleDto) GetTrailerLicensePlate() string`

GetTrailerLicensePlate returns the TrailerLicensePlate field if non-nil, zero value otherwise.

### GetTrailerLicensePlateOk

`func (o *VehicleDto) GetTrailerLicensePlateOk() (*string, bool)`

GetTrailerLicensePlateOk returns a tuple with the TrailerLicensePlate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLicensePlate

`func (o *VehicleDto) SetTrailerLicensePlate(v string)`

SetTrailerLicensePlate sets TrailerLicensePlate field to given value.

### HasTrailerLicensePlate

`func (o *VehicleDto) HasTrailerLicensePlate() bool`

HasTrailerLicensePlate returns a boolean if a field has been set.

### GetIsSubcontractorsSpecific

`func (o *VehicleDto) GetIsSubcontractorsSpecific() bool`

GetIsSubcontractorsSpecific returns the IsSubcontractorsSpecific field if non-nil, zero value otherwise.

### GetIsSubcontractorsSpecificOk

`func (o *VehicleDto) GetIsSubcontractorsSpecificOk() (*bool, bool)`

GetIsSubcontractorsSpecificOk returns a tuple with the IsSubcontractorsSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubcontractorsSpecific

`func (o *VehicleDto) SetIsSubcontractorsSpecific(v bool)`

SetIsSubcontractorsSpecific sets IsSubcontractorsSpecific field to given value.

### HasIsSubcontractorsSpecific

`func (o *VehicleDto) HasIsSubcontractorsSpecific() bool`

HasIsSubcontractorsSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


