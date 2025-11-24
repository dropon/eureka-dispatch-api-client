# SubContractorCommunicationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendOnCharterConfirmationOnAssignment** | Pointer to [**CommunicationTypeConfigurationDto**](CommunicationTypeConfigurationDto.md) |  | [optional] 
**AdditionalEmail** | Pointer to **string** | May contains multiple concatenated values | [optional] 
**AdditionalFax** | Pointer to **string** | May contains multiple concatenated values | [optional] 

## Methods

### NewSubContractorCommunicationConfigurationDto

`func NewSubContractorCommunicationConfigurationDto() *SubContractorCommunicationConfigurationDto`

NewSubContractorCommunicationConfigurationDto instantiates a new SubContractorCommunicationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubContractorCommunicationConfigurationDtoWithDefaults

`func NewSubContractorCommunicationConfigurationDtoWithDefaults() *SubContractorCommunicationConfigurationDto`

NewSubContractorCommunicationConfigurationDtoWithDefaults instantiates a new SubContractorCommunicationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendOnCharterConfirmationOnAssignment

`func (o *SubContractorCommunicationConfigurationDto) GetSendOnCharterConfirmationOnAssignment() CommunicationTypeConfigurationDto`

GetSendOnCharterConfirmationOnAssignment returns the SendOnCharterConfirmationOnAssignment field if non-nil, zero value otherwise.

### GetSendOnCharterConfirmationOnAssignmentOk

`func (o *SubContractorCommunicationConfigurationDto) GetSendOnCharterConfirmationOnAssignmentOk() (*CommunicationTypeConfigurationDto, bool)`

GetSendOnCharterConfirmationOnAssignmentOk returns a tuple with the SendOnCharterConfirmationOnAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnCharterConfirmationOnAssignment

`func (o *SubContractorCommunicationConfigurationDto) SetSendOnCharterConfirmationOnAssignment(v CommunicationTypeConfigurationDto)`

SetSendOnCharterConfirmationOnAssignment sets SendOnCharterConfirmationOnAssignment field to given value.

### HasSendOnCharterConfirmationOnAssignment

`func (o *SubContractorCommunicationConfigurationDto) HasSendOnCharterConfirmationOnAssignment() bool`

HasSendOnCharterConfirmationOnAssignment returns a boolean if a field has been set.

### GetAdditionalEmail

`func (o *SubContractorCommunicationConfigurationDto) GetAdditionalEmail() string`

GetAdditionalEmail returns the AdditionalEmail field if non-nil, zero value otherwise.

### GetAdditionalEmailOk

`func (o *SubContractorCommunicationConfigurationDto) GetAdditionalEmailOk() (*string, bool)`

GetAdditionalEmailOk returns a tuple with the AdditionalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmail

`func (o *SubContractorCommunicationConfigurationDto) SetAdditionalEmail(v string)`

SetAdditionalEmail sets AdditionalEmail field to given value.

### HasAdditionalEmail

`func (o *SubContractorCommunicationConfigurationDto) HasAdditionalEmail() bool`

HasAdditionalEmail returns a boolean if a field has been set.

### GetAdditionalFax

`func (o *SubContractorCommunicationConfigurationDto) GetAdditionalFax() string`

GetAdditionalFax returns the AdditionalFax field if non-nil, zero value otherwise.

### GetAdditionalFaxOk

`func (o *SubContractorCommunicationConfigurationDto) GetAdditionalFaxOk() (*string, bool)`

GetAdditionalFaxOk returns a tuple with the AdditionalFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFax

`func (o *SubContractorCommunicationConfigurationDto) SetAdditionalFax(v string)`

SetAdditionalFax sets AdditionalFax field to given value.

### HasAdditionalFax

`func (o *SubContractorCommunicationConfigurationDto) HasAdditionalFax() bool`

HasAdditionalFax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


