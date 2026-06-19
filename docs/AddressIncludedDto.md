# AddressIncludedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressCategory** | Pointer to [**AddressCategoryDto**](AddressCategoryDto.md) |  | [optional] 
**Agency** | Pointer to [**AgencyDto**](AgencyDto.md) |  | [optional] 
**CircuitAddressSet** | Pointer to [**CircuitAddressSetDto**](CircuitAddressSetDto.md) |  | [optional] 

## Methods

### NewAddressIncludedDto

`func NewAddressIncludedDto() *AddressIncludedDto`

NewAddressIncludedDto instantiates a new AddressIncludedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressIncludedDtoWithDefaults

`func NewAddressIncludedDtoWithDefaults() *AddressIncludedDto`

NewAddressIncludedDtoWithDefaults instantiates a new AddressIncludedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressCategory

`func (o *AddressIncludedDto) GetAddressCategory() AddressCategoryDto`

GetAddressCategory returns the AddressCategory field if non-nil, zero value otherwise.

### GetAddressCategoryOk

`func (o *AddressIncludedDto) GetAddressCategoryOk() (*AddressCategoryDto, bool)`

GetAddressCategoryOk returns a tuple with the AddressCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCategory

`func (o *AddressIncludedDto) SetAddressCategory(v AddressCategoryDto)`

SetAddressCategory sets AddressCategory field to given value.

### HasAddressCategory

`func (o *AddressIncludedDto) HasAddressCategory() bool`

HasAddressCategory returns a boolean if a field has been set.

### GetAgency

`func (o *AddressIncludedDto) GetAgency() AgencyDto`

GetAgency returns the Agency field if non-nil, zero value otherwise.

### GetAgencyOk

`func (o *AddressIncludedDto) GetAgencyOk() (*AgencyDto, bool)`

GetAgencyOk returns a tuple with the Agency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgency

`func (o *AddressIncludedDto) SetAgency(v AgencyDto)`

SetAgency sets Agency field to given value.

### HasAgency

`func (o *AddressIncludedDto) HasAgency() bool`

HasAgency returns a boolean if a field has been set.

### GetCircuitAddressSet

`func (o *AddressIncludedDto) GetCircuitAddressSet() CircuitAddressSetDto`

GetCircuitAddressSet returns the CircuitAddressSet field if non-nil, zero value otherwise.

### GetCircuitAddressSetOk

`func (o *AddressIncludedDto) GetCircuitAddressSetOk() (*CircuitAddressSetDto, bool)`

GetCircuitAddressSetOk returns a tuple with the CircuitAddressSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitAddressSet

`func (o *AddressIncludedDto) SetCircuitAddressSet(v CircuitAddressSetDto)`

SetCircuitAddressSet sets CircuitAddressSet field to given value.

### HasCircuitAddressSet

`func (o *AddressIncludedDto) HasCircuitAddressSet() bool`

HasCircuitAddressSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


