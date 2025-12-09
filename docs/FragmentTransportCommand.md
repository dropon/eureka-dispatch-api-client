# FragmentTransportCommand

## Properties

| Name             | Type                                                                      | Description                            | Notes |
| ---------------- | ------------------------------------------------------------------------- | -------------------------------------- | ----- |
| **TransportUid** | **string**                                                                | Mandatory. Transport unique identifier |
| **Step**         | [**FragmentTransportCommandStepDto**](FragmentTransportCommandStepDto.md) |                                        |
| **LinkDate**     | **Time**                                                                  | Mandatory. Link date                   |

## Methods

### NewFragmentTransportCommand

`func NewFragmentTransportCommand(transportUid string, step FragmentTransportCommandStepDto, linkDate Time, ) *FragmentTransportCommand`

NewFragmentTransportCommand instantiates a new FragmentTransportCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFragmentTransportCommandWithDefaults

`func NewFragmentTransportCommandWithDefaults() *FragmentTransportCommand`

NewFragmentTransportCommandWithDefaults instantiates a new FragmentTransportCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *FragmentTransportCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *FragmentTransportCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *FragmentTransportCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.

### GetStep

`func (o *FragmentTransportCommand) GetStep() FragmentTransportCommandStepDto`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *FragmentTransportCommand) GetStepOk() (*FragmentTransportCommandStepDto, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *FragmentTransportCommand) SetStep(v FragmentTransportCommandStepDto)`

SetStep sets Step field to given value.

### GetLinkDate

`func (o *FragmentTransportCommand) GetLinkDate() Time`

GetLinkDate returns the LinkDate field if non-nil, zero value otherwise.

### GetLinkDateOk

`func (o *FragmentTransportCommand) GetLinkDateOk() (*Time, bool)`

GetLinkDateOk returns a tuple with the LinkDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDate

`func (o *FragmentTransportCommand) SetLinkDate(v Time)`

SetLinkDate sets LinkDate field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
