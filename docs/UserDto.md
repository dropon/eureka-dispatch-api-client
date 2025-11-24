# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **int32** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserDto

`func NewUserDto() *UserDto`

NewUserDto instantiates a new UserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoWithDefaults

`func NewUserDtoWithDefaults() *UserDto`

NewUserDtoWithDefaults instantiates a new UserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserDto) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserDto) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserDto) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLogin

`func (o *UserDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserDto) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserDto) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *UserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocale

`func (o *UserDto) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserDto) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserDto) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserDto) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


