# EntityDeletionConstraintsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEntityFound** | Pointer to **bool** | Indicates whether the entity to be deleted was found | [optional] 
**IsDeletionAllowed** | Pointer to **bool** | Indicates whether deletion operation is permitted | [optional] 
**Message** | Pointer to **string** | Message indicating the reason why deletion is not allowed | [optional] 

## Methods

### NewEntityDeletionConstraintsDto

`func NewEntityDeletionConstraintsDto() *EntityDeletionConstraintsDto`

NewEntityDeletionConstraintsDto instantiates a new EntityDeletionConstraintsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityDeletionConstraintsDtoWithDefaults

`func NewEntityDeletionConstraintsDtoWithDefaults() *EntityDeletionConstraintsDto`

NewEntityDeletionConstraintsDtoWithDefaults instantiates a new EntityDeletionConstraintsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEntityFound

`func (o *EntityDeletionConstraintsDto) GetIsEntityFound() bool`

GetIsEntityFound returns the IsEntityFound field if non-nil, zero value otherwise.

### GetIsEntityFoundOk

`func (o *EntityDeletionConstraintsDto) GetIsEntityFoundOk() (*bool, bool)`

GetIsEntityFoundOk returns a tuple with the IsEntityFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntityFound

`func (o *EntityDeletionConstraintsDto) SetIsEntityFound(v bool)`

SetIsEntityFound sets IsEntityFound field to given value.

### HasIsEntityFound

`func (o *EntityDeletionConstraintsDto) HasIsEntityFound() bool`

HasIsEntityFound returns a boolean if a field has been set.

### GetIsDeletionAllowed

`func (o *EntityDeletionConstraintsDto) GetIsDeletionAllowed() bool`

GetIsDeletionAllowed returns the IsDeletionAllowed field if non-nil, zero value otherwise.

### GetIsDeletionAllowedOk

`func (o *EntityDeletionConstraintsDto) GetIsDeletionAllowedOk() (*bool, bool)`

GetIsDeletionAllowedOk returns a tuple with the IsDeletionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletionAllowed

`func (o *EntityDeletionConstraintsDto) SetIsDeletionAllowed(v bool)`

SetIsDeletionAllowed sets IsDeletionAllowed field to given value.

### HasIsDeletionAllowed

`func (o *EntityDeletionConstraintsDto) HasIsDeletionAllowed() bool`

HasIsDeletionAllowed returns a boolean if a field has been set.

### GetMessage

`func (o *EntityDeletionConstraintsDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EntityDeletionConstraintsDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EntityDeletionConstraintsDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EntityDeletionConstraintsDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


