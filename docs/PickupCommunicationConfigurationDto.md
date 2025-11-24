# PickupCommunicationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendLoadingAlertOnAssignment** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**AdditionalEmail** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalFax** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalSMS** | Pointer to **string** | May contains multiple concatenated values | [optional] 

## Methods

### NewPickupCommunicationConfigurationDto

`func NewPickupCommunicationConfigurationDto() *PickupCommunicationConfigurationDto`

NewPickupCommunicationConfigurationDto instantiates a new PickupCommunicationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupCommunicationConfigurationDtoWithDefaults

`func NewPickupCommunicationConfigurationDtoWithDefaults() *PickupCommunicationConfigurationDto`

NewPickupCommunicationConfigurationDtoWithDefaults instantiates a new PickupCommunicationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendLoadingAlertOnAssignment

`func (o *PickupCommunicationConfigurationDto) GetSendLoadingAlertOnAssignment() CommunicationTypeConfigurationDto`

GetSendLoadingAlertOnAssignment returns the SendLoadingAlertOnAssignment field if non-nil, zero value otherwise.

### GetSendLoadingAlertOnAssignmentOk

`func (o *PickupCommunicationConfigurationDto) GetSendLoadingAlertOnAssignmentOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendLoadingAlertOnAssignmentOk returns a tuple with the SendLoadingAlertOnAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendLoadingAlertOnAssignment

`func (o *PickupCommunicationConfigurationDto) SetSendLoadingAlertOnAssignment(v CommunicationTypeConfigurationDto)`

SetSendLoadingAlertOnAssignment sets SendLoadingAlertOnAssignment field to given value.

### HasSendLoadingAlertOnAssignment

`func (o *PickupCommunicationConfigurationDto) HasSendLoadingAlertOnAssignment() bool`

HasSendLoadingAlertOnAssignment returns a boolean if a field has been set.

### GetAdditionalEmail

`func (o *PickupCommunicationConfigurationDto) GetAdditionalEmail() string`

GetAdditionalEmail returns the AdditionalEmail field if non-nil, zero value otherwise.

### GetAdditionalEmailOk

`func (o *PickupCommunicationConfigurationDto) GetAdditionalEmailOk() (*string, bool)`

GetAdditionalEmailOk returns a tuple with the AdditionalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmail

`func (o *PickupCommunicationConfigurationDto) SetAdditionalEmail(v string)`

SetAdditionalEmail sets AdditionalEmail field to given value.

### HasAdditionalEmail

`func (o *PickupCommunicationConfigurationDto) HasAdditionalEmail() bool`

HasAdditionalEmail returns a boolean if a field has been set.

### GetAdditionalFax

`func (o *PickupCommunicationConfigurationDto) GetAdditionalFax() string`

GetAdditionalFax returns the AdditionalFax field if non-nil, zero value otherwise.

### GetAdditionalFaxOk

`func (o *PickupCommunicationConfigurationDto) GetAdditionalFaxOk() (*string, bool)`

GetAdditionalFaxOk returns a tuple with the AdditionalFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFax

`func (o *PickupCommunicationConfigurationDto) SetAdditionalFax(v string)`

SetAdditionalFax sets AdditionalFax field to given value.

### HasAdditionalFax

`func (o *PickupCommunicationConfigurationDto) HasAdditionalFax() bool`

HasAdditionalFax returns a boolean if a field has been set.

### GetAdditionalSMS

`func (o *PickupCommunicationConfigurationDto) GetAdditionalSMS() string`

GetAdditionalSMS returns the AdditionalSMS field if non-nil, zero value otherwise.

### GetAdditionalSMSOk

`func (o *PickupCommunicationConfigurationDto) GetAdditionalSMSOk() (*string, bool)`

GetAdditionalSMSOk returns a tuple with the AdditionalSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSMS

`func (o *PickupCommunicationConfigurationDto) SetAdditionalSMS(v string)`

SetAdditionalSMS sets AdditionalSMS field to given value.

### HasAdditionalSMS

`func (o *PickupCommunicationConfigurationDto) HasAdditionalSMS() bool`

HasAdditionalSMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


