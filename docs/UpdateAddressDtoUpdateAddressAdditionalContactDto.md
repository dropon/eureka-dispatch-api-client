# UpdateAddressDtoUpdateAddressAdditionalContactDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeAction** | Pointer to **string** | Merge action to perform : add, update, remove. | [optional] 
**AddressContactId** | Pointer to **int32** | This property is an entity identifier. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AlternativePhone** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAddressDtoUpdateAddressAdditionalContactDto

`func NewUpdateAddressDtoUpdateAddressAdditionalContactDto() *UpdateAddressDtoUpdateAddressAdditionalContactDto`

NewUpdateAddressDtoUpdateAddressAdditionalContactDto instantiates a new UpdateAddressDtoUpdateAddressAdditionalContactDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAddressDtoUpdateAddressAdditionalContactDtoWithDefaults

`func NewUpdateAddressDtoUpdateAddressAdditionalContactDtoWithDefaults() *UpdateAddressDtoUpdateAddressAdditionalContactDto`

NewUpdateAddressDtoUpdateAddressAdditionalContactDtoWithDefaults instantiates a new UpdateAddressDtoUpdateAddressAdditionalContactDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeAction

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetAddressContactId

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetAddressContactId() int32`

GetAddressContactId returns the AddressContactId field if non-nil, zero value otherwise.

### GetAddressContactIdOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetAddressContactIdOk() (*int32, bool)`

GetAddressContactIdOk returns a tuple with the AddressContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressContactId

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetAddressContactId(v int32)`

SetAddressContactId sets AddressContactId field to given value.

### HasAddressContactId

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasAddressContactId() bool`

HasAddressContactId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAlternativePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetAlternativePhone() string`

GetAlternativePhone returns the AlternativePhone field if non-nil, zero value otherwise.

### GetAlternativePhoneOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetAlternativePhoneOk() (*string, bool)`

GetAlternativePhoneOk returns a tuple with the AlternativePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetAlternativePhone(v string)`

SetAlternativePhone sets AlternativePhone field to given value.

### HasAlternativePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasAlternativePhone() bool`

HasAlternativePhone returns a boolean if a field has been set.

### GetMobilePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetFax

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateAddressDtoUpdateAddressAdditionalContactDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


