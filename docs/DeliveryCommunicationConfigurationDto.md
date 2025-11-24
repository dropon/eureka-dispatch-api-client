# DeliveryCommunicationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendDeliveryAlertOnAssignment** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**SendToDeliveryContactOnPickup** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**AdditionalEmail** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalFax** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalSMS** | Pointer to **string** | May contains multiple concatenated values | [optional] 

## Methods

### NewDeliveryCommunicationConfigurationDto

`func NewDeliveryCommunicationConfigurationDto() *DeliveryCommunicationConfigurationDto`

NewDeliveryCommunicationConfigurationDto instantiates a new DeliveryCommunicationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryCommunicationConfigurationDtoWithDefaults

`func NewDeliveryCommunicationConfigurationDtoWithDefaults() *DeliveryCommunicationConfigurationDto`

NewDeliveryCommunicationConfigurationDtoWithDefaults instantiates a new DeliveryCommunicationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendDeliveryAlertOnAssignment

`func (o *DeliveryCommunicationConfigurationDto) GetSendDeliveryAlertOnAssignment() CommunicationTypeConfigurationDto`

GetSendDeliveryAlertOnAssignment returns the SendDeliveryAlertOnAssignment field if non-nil, zero value otherwise.

### GetSendDeliveryAlertOnAssignmentOk

`func (o *DeliveryCommunicationConfigurationDto) GetSendDeliveryAlertOnAssignmentOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendDeliveryAlertOnAssignmentOk returns a tuple with the SendDeliveryAlertOnAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDeliveryAlertOnAssignment

`func (o *DeliveryCommunicationConfigurationDto) SetSendDeliveryAlertOnAssignment(v CommunicationTypeConfigurationDto)`

SetSendDeliveryAlertOnAssignment sets SendDeliveryAlertOnAssignment field to given value.

### HasSendDeliveryAlertOnAssignment

`func (o *DeliveryCommunicationConfigurationDto) HasSendDeliveryAlertOnAssignment() bool`

HasSendDeliveryAlertOnAssignment returns a boolean if a field has been set.

### GetSendToDeliveryContactOnPickup

`func (o *DeliveryCommunicationConfigurationDto) GetSendToDeliveryContactOnPickup() CommunicationTypeConfigurationDto`

GetSendToDeliveryContactOnPickup returns the SendToDeliveryContactOnPickup field if non-nil, zero value otherwise.

### GetSendToDeliveryContactOnPickupOk

`func (o *DeliveryCommunicationConfigurationDto) GetSendToDeliveryContactOnPickupOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendToDeliveryContactOnPickupOk returns a tuple with the SendToDeliveryContactOnPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToDeliveryContactOnPickup

`func (o *DeliveryCommunicationConfigurationDto) SetSendToDeliveryContactOnPickup(v CommunicationTypeConfigurationDto)`

SetSendToDeliveryContactOnPickup sets SendToDeliveryContactOnPickup field to given value.

### HasSendToDeliveryContactOnPickup

`func (o *DeliveryCommunicationConfigurationDto) HasSendToDeliveryContactOnPickup() bool`

HasSendToDeliveryContactOnPickup returns a boolean if a field has been set.

### GetAdditionalEmail

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalEmail() string`

GetAdditionalEmail returns the AdditionalEmail field if non-nil, zero value otherwise.

### GetAdditionalEmailOk

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalEmailOk() (*string, bool)`

GetAdditionalEmailOk returns a tuple with the AdditionalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmail

`func (o *DeliveryCommunicationConfigurationDto) SetAdditionalEmail(v string)`

SetAdditionalEmail sets AdditionalEmail field to given value.

### HasAdditionalEmail

`func (o *DeliveryCommunicationConfigurationDto) HasAdditionalEmail() bool`

HasAdditionalEmail returns a boolean if a field has been set.

### GetAdditionalFax

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalFax() string`

GetAdditionalFax returns the AdditionalFax field if non-nil, zero value otherwise.

### GetAdditionalFaxOk

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalFaxOk() (*string, bool)`

GetAdditionalFaxOk returns a tuple with the AdditionalFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFax

`func (o *DeliveryCommunicationConfigurationDto) SetAdditionalFax(v string)`

SetAdditionalFax sets AdditionalFax field to given value.

### HasAdditionalFax

`func (o *DeliveryCommunicationConfigurationDto) HasAdditionalFax() bool`

HasAdditionalFax returns a boolean if a field has been set.

### GetAdditionalSMS

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalSMS() string`

GetAdditionalSMS returns the AdditionalSMS field if non-nil, zero value otherwise.

### GetAdditionalSMSOk

`func (o *DeliveryCommunicationConfigurationDto) GetAdditionalSMSOk() (*string, bool)`

GetAdditionalSMSOk returns a tuple with the AdditionalSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSMS

`func (o *DeliveryCommunicationConfigurationDto) SetAdditionalSMS(v string)`

SetAdditionalSMS sets AdditionalSMS field to given value.

### HasAdditionalSMS

`func (o *DeliveryCommunicationConfigurationDto) HasAdditionalSMS() bool`

HasAdditionalSMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


