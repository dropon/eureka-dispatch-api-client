# AssignTransportDtoApiBehaviorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForcePreAssignTransport** | Pointer to **bool** | Option to pre-assign the transport during the assignment process. | [optional] 
**ForceSendCommunications** | Pointer to **bool** | Option to force sending communications (email, fax, sms), when an assignment was already realized on the transport | [optional] 

## Methods

### NewAssignTransportDtoApiBehaviorDto

`func NewAssignTransportDtoApiBehaviorDto() *AssignTransportDtoApiBehaviorDto`

NewAssignTransportDtoApiBehaviorDto instantiates a new AssignTransportDtoApiBehaviorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignTransportDtoApiBehaviorDtoWithDefaults

`func NewAssignTransportDtoApiBehaviorDtoWithDefaults() *AssignTransportDtoApiBehaviorDto`

NewAssignTransportDtoApiBehaviorDtoWithDefaults instantiates a new AssignTransportDtoApiBehaviorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForcePreAssignTransport

`func (o *AssignTransportDtoApiBehaviorDto) GetForcePreAssignTransport() bool`

GetForcePreAssignTransport returns the ForcePreAssignTransport field if non-nil, zero value otherwise.

### GetForcePreAssignTransportOk

`func (o *AssignTransportDtoApiBehaviorDto) GetForcePreAssignTransportOk() (*bool, bool)`

GetForcePreAssignTransportOk returns a tuple with the ForcePreAssignTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePreAssignTransport

`func (o *AssignTransportDtoApiBehaviorDto) SetForcePreAssignTransport(v bool)`

SetForcePreAssignTransport sets ForcePreAssignTransport field to given value.

### HasForcePreAssignTransport

`func (o *AssignTransportDtoApiBehaviorDto) HasForcePreAssignTransport() bool`

HasForcePreAssignTransport returns a boolean if a field has been set.

### GetForceSendCommunications

`func (o *AssignTransportDtoApiBehaviorDto) GetForceSendCommunications() bool`

GetForceSendCommunications returns the ForceSendCommunications field if non-nil, zero value otherwise.

### GetForceSendCommunicationsOk

`func (o *AssignTransportDtoApiBehaviorDto) GetForceSendCommunicationsOk() (*bool, bool)`

GetForceSendCommunicationsOk returns a tuple with the ForceSendCommunications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSendCommunications

`func (o *AssignTransportDtoApiBehaviorDto) SetForceSendCommunications(v bool)`

SetForceSendCommunications sets ForceSendCommunications field to given value.

### HasForceSendCommunications

`func (o *AssignTransportDtoApiBehaviorDto) HasForceSendCommunications() bool`

HasForceSendCommunications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


