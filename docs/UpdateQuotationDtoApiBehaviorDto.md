# UpdateQuotationDtoApiBehaviorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceUniqueOrder** | Pointer to **bool** | This option is used to force the option unique order in the mission.    By default, the unique order option is enabled when the mission contains many transports. | [optional] 
**SkipPackageNatureDefaultValuesApplication** | Pointer to **bool** | This option is used to avoid copying the default values from the package nature.  By default, values are copied from the package nature when they are not provided :   quantity, package code, dimensions (height, width, length), sizes, references. | [optional] 
**TryGeocodeTransportSteps** | Pointer to **bool** | Try to geocode transport steps if the step is not yet geocoded.    This is only done if the geolocation service is available.  By default, geocoding is attempted only when distance computation is enabled. | [optional] 
**IgnoreItineraryComputationFailures** | Pointer to **bool** | This option is used to ignore errors if itinerary computation failed. This step is done only for new transports.  By default, if itinerary computation fails, an error is raised. | [optional] 
**UpdateTransportOperationComment** | Pointer to **string** | Comment for the update operation, will be added in transport history if specified. | [optional] 

## Methods

### NewUpdateQuotationDtoApiBehaviorDto

`func NewUpdateQuotationDtoApiBehaviorDto() *UpdateQuotationDtoApiBehaviorDto`

NewUpdateQuotationDtoApiBehaviorDto instantiates a new UpdateQuotationDtoApiBehaviorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQuotationDtoApiBehaviorDtoWithDefaults

`func NewUpdateQuotationDtoApiBehaviorDtoWithDefaults() *UpdateQuotationDtoApiBehaviorDto`

NewUpdateQuotationDtoApiBehaviorDtoWithDefaults instantiates a new UpdateQuotationDtoApiBehaviorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceUniqueOrder

`func (o *UpdateQuotationDtoApiBehaviorDto) GetForceUniqueOrder() bool`

GetForceUniqueOrder returns the ForceUniqueOrder field if non-nil, zero value otherwise.

### GetForceUniqueOrderOk

`func (o *UpdateQuotationDtoApiBehaviorDto) GetForceUniqueOrderOk() (*bool, bool)`

GetForceUniqueOrderOk returns a tuple with the ForceUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUniqueOrder

`func (o *UpdateQuotationDtoApiBehaviorDto) SetForceUniqueOrder(v bool)`

SetForceUniqueOrder sets ForceUniqueOrder field to given value.

### HasForceUniqueOrder

`func (o *UpdateQuotationDtoApiBehaviorDto) HasForceUniqueOrder() bool`

HasForceUniqueOrder returns a boolean if a field has been set.

### GetSkipPackageNatureDefaultValuesApplication

`func (o *UpdateQuotationDtoApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplication() bool`

GetSkipPackageNatureDefaultValuesApplication returns the SkipPackageNatureDefaultValuesApplication field if non-nil, zero value otherwise.

### GetSkipPackageNatureDefaultValuesApplicationOk

`func (o *UpdateQuotationDtoApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplicationOk() (*bool, bool)`

GetSkipPackageNatureDefaultValuesApplicationOk returns a tuple with the SkipPackageNatureDefaultValuesApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPackageNatureDefaultValuesApplication

`func (o *UpdateQuotationDtoApiBehaviorDto) SetSkipPackageNatureDefaultValuesApplication(v bool)`

SetSkipPackageNatureDefaultValuesApplication sets SkipPackageNatureDefaultValuesApplication field to given value.

### HasSkipPackageNatureDefaultValuesApplication

`func (o *UpdateQuotationDtoApiBehaviorDto) HasSkipPackageNatureDefaultValuesApplication() bool`

HasSkipPackageNatureDefaultValuesApplication returns a boolean if a field has been set.

### GetTryGeocodeTransportSteps

`func (o *UpdateQuotationDtoApiBehaviorDto) GetTryGeocodeTransportSteps() bool`

GetTryGeocodeTransportSteps returns the TryGeocodeTransportSteps field if non-nil, zero value otherwise.

### GetTryGeocodeTransportStepsOk

`func (o *UpdateQuotationDtoApiBehaviorDto) GetTryGeocodeTransportStepsOk() (*bool, bool)`

GetTryGeocodeTransportStepsOk returns a tuple with the TryGeocodeTransportSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryGeocodeTransportSteps

`func (o *UpdateQuotationDtoApiBehaviorDto) SetTryGeocodeTransportSteps(v bool)`

SetTryGeocodeTransportSteps sets TryGeocodeTransportSteps field to given value.

### HasTryGeocodeTransportSteps

`func (o *UpdateQuotationDtoApiBehaviorDto) HasTryGeocodeTransportSteps() bool`

HasTryGeocodeTransportSteps returns a boolean if a field has been set.

### GetIgnoreItineraryComputationFailures

`func (o *UpdateQuotationDtoApiBehaviorDto) GetIgnoreItineraryComputationFailures() bool`

GetIgnoreItineraryComputationFailures returns the IgnoreItineraryComputationFailures field if non-nil, zero value otherwise.

### GetIgnoreItineraryComputationFailuresOk

`func (o *UpdateQuotationDtoApiBehaviorDto) GetIgnoreItineraryComputationFailuresOk() (*bool, bool)`

GetIgnoreItineraryComputationFailuresOk returns a tuple with the IgnoreItineraryComputationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreItineraryComputationFailures

`func (o *UpdateQuotationDtoApiBehaviorDto) SetIgnoreItineraryComputationFailures(v bool)`

SetIgnoreItineraryComputationFailures sets IgnoreItineraryComputationFailures field to given value.

### HasIgnoreItineraryComputationFailures

`func (o *UpdateQuotationDtoApiBehaviorDto) HasIgnoreItineraryComputationFailures() bool`

HasIgnoreItineraryComputationFailures returns a boolean if a field has been set.

### GetUpdateTransportOperationComment

`func (o *UpdateQuotationDtoApiBehaviorDto) GetUpdateTransportOperationComment() string`

GetUpdateTransportOperationComment returns the UpdateTransportOperationComment field if non-nil, zero value otherwise.

### GetUpdateTransportOperationCommentOk

`func (o *UpdateQuotationDtoApiBehaviorDto) GetUpdateTransportOperationCommentOk() (*string, bool)`

GetUpdateTransportOperationCommentOk returns a tuple with the UpdateTransportOperationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTransportOperationComment

`func (o *UpdateQuotationDtoApiBehaviorDto) SetUpdateTransportOperationComment(v string)`

SetUpdateTransportOperationComment sets UpdateTransportOperationComment field to given value.

### HasUpdateTransportOperationComment

`func (o *UpdateQuotationDtoApiBehaviorDto) HasUpdateTransportOperationComment() bool`

HasUpdateTransportOperationComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


