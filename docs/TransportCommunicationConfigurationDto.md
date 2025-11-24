# TransportCommunicationConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CustomerCommunicationConfigurationDto**](CustomerCommunicationConfigurationDto.md) |  | [optional] 
**SubContractor** | Pointer to [**SubContractorCommunicationConfigurationDto**](SubContractorCommunicationConfigurationDto.md) |  | [optional] 
**Pickup** | Pointer to [**PickupCommunicationConfigurationDto**](PickupCommunicationConfigurationDto.md) |  | [optional] 
**Delivery** | Pointer to [**DeliveryCommunicationConfigurationDto**](DeliveryCommunicationConfigurationDto.md) |  | [optional] 

## Methods

### NewTransportCommunicationConfigurationDto

`func NewTransportCommunicationConfigurationDto() *TransportCommunicationConfigurationDto`

NewTransportCommunicationConfigurationDto instantiates a new TransportCommunicationConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportCommunicationConfigurationDtoWithDefaults

`func NewTransportCommunicationConfigurationDtoWithDefaults() *TransportCommunicationConfigurationDto`

NewTransportCommunicationConfigurationDtoWithDefaults instantiates a new TransportCommunicationConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *TransportCommunicationConfigurationDto) GetCustomer() CustomerCommunicationConfigurationDto`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *TransportCommunicationConfigurationDto) GetCustomerOk() (*CustomerCommunicationConfigurationDto, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *TransportCommunicationConfigurationDto) SetCustomer(v CustomerCommunicationConfigurationDto)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *TransportCommunicationConfigurationDto) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSubContractor

`func (o *TransportCommunicationConfigurationDto) GetSubContractor() SubContractorCommunicationConfigurationDto`

GetSubContractor returns the SubContractor field if non-nil, zero value otherwise.

### GetSubContractorOk

`func (o *TransportCommunicationConfigurationDto) GetSubContractorOk() (*SubContractorCommunicationConfigurationDto, bool)`

GetSubContractorOk returns a tuple with the SubContractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContractor

`func (o *TransportCommunicationConfigurationDto) SetSubContractor(v SubContractorCommunicationConfigurationDto)`

SetSubContractor sets SubContractor field to given value.

### HasSubContractor

`func (o *TransportCommunicationConfigurationDto) HasSubContractor() bool`

HasSubContractor returns a boolean if a field has been set.

### GetPickup

`func (o *TransportCommunicationConfigurationDto) GetPickup() PickupCommunicationConfigurationDto`

GetPickup returns the Pickup field if non-nil, zero value otherwise.

### GetPickupOk

`func (o *TransportCommunicationConfigurationDto) GetPickupOk() (*PickupCommunicationConfigurationDto, bool)`

GetPickupOk returns a tuple with the Pickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickup

`func (o *TransportCommunicationConfigurationDto) SetPickup(v PickupCommunicationConfigurationDto)`

SetPickup sets Pickup field to given value.

### HasPickup

`func (o *TransportCommunicationConfigurationDto) HasPickup() bool`

HasPickup returns a boolean if a field has been set.

### GetDelivery

`func (o *TransportCommunicationConfigurationDto) GetDelivery() DeliveryCommunicationConfigurationDto`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *TransportCommunicationConfigurationDto) GetDeliveryOk() (*DeliveryCommunicationConfigurationDto, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *TransportCommunicationConfigurationDto) SetDelivery(v DeliveryCommunicationConfigurationDto)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *TransportCommunicationConfigurationDto) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


