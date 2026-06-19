# AssignmentConstraintDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsBlocking** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewAssignmentConstraintDto

`func NewAssignmentConstraintDto() *AssignmentConstraintDto`

NewAssignmentConstraintDto instantiates a new AssignmentConstraintDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentConstraintDtoWithDefaults

`func NewAssignmentConstraintDtoWithDefaults() *AssignmentConstraintDto`

NewAssignmentConstraintDtoWithDefaults instantiates a new AssignmentConstraintDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsBlocking

`func (o *AssignmentConstraintDto) GetIsBlocking() bool`

GetIsBlocking returns the IsBlocking field if non-nil, zero value otherwise.

### GetIsBlockingOk

`func (o *AssignmentConstraintDto) GetIsBlockingOk() (*bool, bool)`

GetIsBlockingOk returns a tuple with the IsBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocking

`func (o *AssignmentConstraintDto) SetIsBlocking(v bool)`

SetIsBlocking sets IsBlocking field to given value.

### HasIsBlocking

`func (o *AssignmentConstraintDto) HasIsBlocking() bool`

HasIsBlocking returns a boolean if a field has been set.

### GetMessage

`func (o *AssignmentConstraintDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssignmentConstraintDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssignmentConstraintDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AssignmentConstraintDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


