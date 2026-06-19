# AssignBatchTransportsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportUids** | **[]string** | Mandatory. List of transports unique identifiers to assign/unassign.    Limited to 50 transports per request. | 
**Assignment** | [**AssignTransportDto**](AssignTransportDto.md) |  | 

## Methods

### NewAssignBatchTransportsDto

`func NewAssignBatchTransportsDto(transportUids []string, assignment AssignTransportDto, ) *AssignBatchTransportsDto`

NewAssignBatchTransportsDto instantiates a new AssignBatchTransportsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignBatchTransportsDtoWithDefaults

`func NewAssignBatchTransportsDtoWithDefaults() *AssignBatchTransportsDto`

NewAssignBatchTransportsDtoWithDefaults instantiates a new AssignBatchTransportsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUids

`func (o *AssignBatchTransportsDto) GetTransportUids() []string`

GetTransportUids returns the TransportUids field if non-nil, zero value otherwise.

### GetTransportUidsOk

`func (o *AssignBatchTransportsDto) GetTransportUidsOk() (*[]string, bool)`

GetTransportUidsOk returns a tuple with the TransportUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUids

`func (o *AssignBatchTransportsDto) SetTransportUids(v []string)`

SetTransportUids sets TransportUids field to given value.


### GetAssignment

`func (o *AssignBatchTransportsDto) GetAssignment() AssignTransportDto`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *AssignBatchTransportsDto) GetAssignmentOk() (*AssignTransportDto, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *AssignBatchTransportsDto) SetAssignment(v AssignTransportDto)`

SetAssignment sets Assignment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


