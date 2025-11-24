# CustomerCommunicationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendOnMissionEntry** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnAssignment** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnPickup** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnDelivery** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnTermination** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnCheckPointDelay** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendOnAppointment** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**AdditionalEmail** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalFax** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalSMS** | Pointer to **string** | May contains multiple concatenated values | [optional] 

## Methods

### NewCustomerCommunicationConfigurationDto

`func NewCustomerCommunicationConfigurationDto() *CustomerCommunicationConfigurationDto`

NewCustomerCommunicationConfigurationDto instantiates a new CustomerCommunicationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCommunicationConfigurationDtoWithDefaults

`func NewCustomerCommunicationConfigurationDtoWithDefaults() *CustomerCommunicationConfigurationDto`

NewCustomerCommunicationConfigurationDtoWithDefaults instantiates a new CustomerCommunicationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendOnMissionEntry

`func (o *CustomerCommunicationConfigurationDto) GetSendOnMissionEntry() CommunicationTypeConfigurationDto`

GetSendOnMissionEntry returns the SendOnMissionEntry field if non-nil, zero value otherwise.

### GetSendOnMissionEntryOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnMissionEntryOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnMissionEntryOk returns a tuple with the SendOnMissionEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnMissionEntry

`func (o *CustomerCommunicationConfigurationDto) SetSendOnMissionEntry(v CommunicationTypeConfigurationDto)`

SetSendOnMissionEntry sets SendOnMissionEntry field to given value.

### HasSendOnMissionEntry

`func (o *CustomerCommunicationConfigurationDto) HasSendOnMissionEntry() bool`

HasSendOnMissionEntry returns a boolean if a field has been set.

### GetSendOnAssignment

`func (o *CustomerCommunicationConfigurationDto) GetSendOnAssignment() CommunicationTypeConfigurationDto`

GetSendOnAssignment returns the SendOnAssignment field if non-nil, zero value otherwise.

### GetSendOnAssignmentOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnAssignmentOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnAssignmentOk returns a tuple with the SendOnAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnAssignment

`func (o *CustomerCommunicationConfigurationDto) SetSendOnAssignment(v CommunicationTypeConfigurationDto)`

SetSendOnAssignment sets SendOnAssignment field to given value.

### HasSendOnAssignment

`func (o *CustomerCommunicationConfigurationDto) HasSendOnAssignment() bool`

HasSendOnAssignment returns a boolean if a field has been set.

### GetSendOnPickup

`func (o *CustomerCommunicationConfigurationDto) GetSendOnPickup() CommunicationTypeConfigurationDto`

GetSendOnPickup returns the SendOnPickup field if non-nil, zero value otherwise.

### GetSendOnPickupOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnPickupOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnPickupOk returns a tuple with the SendOnPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnPickup

`func (o *CustomerCommunicationConfigurationDto) SetSendOnPickup(v CommunicationTypeConfigurationDto)`

SetSendOnPickup sets SendOnPickup field to given value.

### HasSendOnPickup

`func (o *CustomerCommunicationConfigurationDto) HasSendOnPickup() bool`

HasSendOnPickup returns a boolean if a field has been set.

### GetSendOnDelivery

`func (o *CustomerCommunicationConfigurationDto) GetSendOnDelivery() CommunicationTypeConfigurationDto`

GetSendOnDelivery returns the SendOnDelivery field if non-nil, zero value otherwise.

### GetSendOnDeliveryOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnDeliveryOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnDeliveryOk returns a tuple with the SendOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnDelivery

`func (o *CustomerCommunicationConfigurationDto) SetSendOnDelivery(v CommunicationTypeConfigurationDto)`

SetSendOnDelivery sets SendOnDelivery field to given value.

### HasSendOnDelivery

`func (o *CustomerCommunicationConfigurationDto) HasSendOnDelivery() bool`

HasSendOnDelivery returns a boolean if a field has been set.

### GetSendOnTermination

`func (o *CustomerCommunicationConfigurationDto) GetSendOnTermination() CommunicationTypeConfigurationDto`

GetSendOnTermination returns the SendOnTermination field if non-nil, zero value otherwise.

### GetSendOnTerminationOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnTerminationOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnTerminationOk returns a tuple with the SendOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnTermination

`func (o *CustomerCommunicationConfigurationDto) SetSendOnTermination(v CommunicationTypeConfigurationDto)`

SetSendOnTermination sets SendOnTermination field to given value.

### HasSendOnTermination

`func (o *CustomerCommunicationConfigurationDto) HasSendOnTermination() bool`

HasSendOnTermination returns a boolean if a field has been set.

### GetSendOnCheckPointDelay

`func (o *CustomerCommunicationConfigurationDto) GetSendOnCheckPointDelay() CommunicationTypeConfigurationDto`

GetSendOnCheckPointDelay returns the SendOnCheckPointDelay field if non-nil, zero value otherwise.

### GetSendOnCheckPointDelayOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnCheckPointDelayOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnCheckPointDelayOk returns a tuple with the SendOnCheckPointDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnCheckPointDelay

`func (o *CustomerCommunicationConfigurationDto) SetSendOnCheckPointDelay(v CommunicationTypeConfigurationDto)`

SetSendOnCheckPointDelay sets SendOnCheckPointDelay field to given value.

### HasSendOnCheckPointDelay

`func (o *CustomerCommunicationConfigurationDto) HasSendOnCheckPointDelay() bool`

HasSendOnCheckPointDelay returns a boolean if a field has been set.

### GetSendOnAppointment

`func (o *CustomerCommunicationConfigurationDto) GetSendOnAppointment() CommunicationTypeConfigurationDto`

GetSendOnAppointment returns the SendOnAppointment field if non-nil, zero value otherwise.

### GetSendOnAppointmentOk

`func (o *CustomerCommunicationConfigurationDto) GetSendOnAppointmentOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnAppointmentOk returns a tuple with the SendOnAppointment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnAppointment

`func (o *CustomerCommunicationConfigurationDto) SetSendOnAppointment(v CommunicationTypeConfigurationDto)`

SetSendOnAppointment sets SendOnAppointment field to given value.

### HasSendOnAppointment

`func (o *CustomerCommunicationConfigurationDto) HasSendOnAppointment() bool`

HasSendOnAppointment returns a boolean if a field has been set.

### GetAdditionalEmail

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalEmail() string`

GetAdditionalEmail returns the AdditionalEmail field if non-nil, zero value otherwise.

### GetAdditionalEmailOk

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalEmailOk() (*string, bool)`

GetAdditionalEmailOk returns a tuple with the AdditionalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmail

`func (o *CustomerCommunicationConfigurationDto) SetAdditionalEmail(v string)`

SetAdditionalEmail sets AdditionalEmail field to given value.

### HasAdditionalEmail

`func (o *CustomerCommunicationConfigurationDto) HasAdditionalEmail() bool`

HasAdditionalEmail returns a boolean if a field has been set.

### GetAdditionalFax

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalFax() string`

GetAdditionalFax returns the AdditionalFax field if non-nil, zero value otherwise.

### GetAdditionalFaxOk

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalFaxOk() (*string, bool)`

GetAdditionalFaxOk returns a tuple with the AdditionalFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFax

`func (o *CustomerCommunicationConfigurationDto) SetAdditionalFax(v string)`

SetAdditionalFax sets AdditionalFax field to given value.

### HasAdditionalFax

`func (o *CustomerCommunicationConfigurationDto) HasAdditionalFax() bool`

HasAdditionalFax returns a boolean if a field has been set.

### GetAdditionalSMS

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalSMS() string`

GetAdditionalSMS returns the AdditionalSMS field if non-nil, zero value otherwise.

### GetAdditionalSMSOk

`func (o *CustomerCommunicationConfigurationDto) GetAdditionalSMSOk() (*string, bool)`

GetAdditionalSMSOk returns a tuple with the AdditionalSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSMS

`func (o *CustomerCommunicationConfigurationDto) SetAdditionalSMS(v string)`

SetAdditionalSMS sets AdditionalSMS field to given value.

### HasAdditionalSMS

`func (o *CustomerCommunicationConfigurationDto) HasAdditionalSMS() bool`

HasAdditionalSMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


