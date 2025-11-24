# CreateMissionCommandApiBehaviorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceUniqueOrder** | Pointer to **bool** | This option is used to force the option unique order in the mission.    By default, the unique order option is enabled when the mission contains many transports. | [optional] 
**SkipPackageNatureDefaultValuesApplication** | Pointer to **bool** | This option is used to avoid copying the default values from the package nature.  By default, values are copied from the package nature when they are not provided :   quantity, package code, dimensions (height, width, length), sizes, references. | [optional] 
**TryGeocodeTransportSteps** | Pointer to **bool** | Try to geocode transport steps if the step is not yet geocoded.    This is only done if the geolocation service is available.  By default, geocoding is attempted only when distance computation is enabled. | [optional] 
**IgnoreItineraryComputationFailures** | Pointer to **bool** | This option is used to ignore errors if itinerary computation failed.  By default, if itinerary computation fails, an error is raised. | [optional] 

## Methods

### NewCreateMissionCommandApiBehaviorDto

`func NewCreateMissionCommandApiBehaviorDto() *CreateMissionCommandApiBehaviorDto`

NewCreateMissionCommandApiBehaviorDto instantiates a new CreateMissionCommandApiBehaviorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMissionCommandApiBehaviorDtoWithDefaults

`func NewCreateMissionCommandApiBehaviorDtoWithDefaults() *CreateMissionCommandApiBehaviorDto`

NewCreateMissionCommandApiBehaviorDtoWithDefaults instantiates a new CreateMissionCommandApiBehaviorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceUniqueOrder

`func (o *CreateMissionCommandApiBehaviorDto) GetForceUniqueOrder() bool`

GetForceUniqueOrder returns the ForceUniqueOrder field if non-nil, zero value otherwise.

### GetForceUniqueOrderOk

`func (o *CreateMissionCommandApiBehaviorDto) GetForceUniqueOrderOk() (*bool, bool)`

GetForceUniqueOrderOk returns a tuple with the ForceUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUniqueOrder

`func (o *CreateMissionCommandApiBehaviorDto) SetForceUniqueOrder(v bool)`

SetForceUniqueOrder sets ForceUniqueOrder field to given value.

### HasForceUniqueOrder

`func (o *CreateMissionCommandApiBehaviorDto) HasForceUniqueOrder() bool`

HasForceUniqueOrder returns a boolean if a field has been set.

### GetSkipPackageNatureDefaultValuesApplication

`func (o *CreateMissionCommandApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplication() bool`

GetSkipPackageNatureDefaultValuesApplication returns the SkipPackageNatureDefaultValuesApplication field if non-nil, zero value otherwise.

### GetSkipPackageNatureDefaultValuesApplicationOk

`func (o *CreateMissionCommandApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplicationOk() (*bool, bool)`

GetSkipPackageNatureDefaultValuesApplicationOk returns a tuple with the SkipPackageNatureDefaultValuesApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPackageNatureDefaultValuesApplication

`func (o *CreateMissionCommandApiBehaviorDto) SetSkipPackageNatureDefaultValuesApplication(v bool)`

SetSkipPackageNatureDefaultValuesApplication sets SkipPackageNatureDefaultValuesApplication field to given value.

### HasSkipPackageNatureDefaultValuesApplication

`func (o *CreateMissionCommandApiBehaviorDto) HasSkipPackageNatureDefaultValuesApplication() bool`

HasSkipPackageNatureDefaultValuesApplication returns a boolean if a field has been set.

### GetTryGeocodeTransportSteps

`func (o *CreateMissionCommandApiBehaviorDto) GetTryGeocodeTransportSteps() bool`

GetTryGeocodeTransportSteps returns the TryGeocodeTransportSteps field if non-nil, zero value otherwise.

### GetTryGeocodeTransportStepsOk

`func (o *CreateMissionCommandApiBehaviorDto) GetTryGeocodeTransportStepsOk() (*bool, bool)`

GetTryGeocodeTransportStepsOk returns a tuple with the TryGeocodeTransportSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryGeocodeTransportSteps

`func (o *CreateMissionCommandApiBehaviorDto) SetTryGeocodeTransportSteps(v bool)`

SetTryGeocodeTransportSteps sets TryGeocodeTransportSteps field to given value.

### HasTryGeocodeTransportSteps

`func (o *CreateMissionCommandApiBehaviorDto) HasTryGeocodeTransportSteps() bool`

HasTryGeocodeTransportSteps returns a boolean if a field has been set.

### GetIgnoreItineraryComputationFailures

`func (o *CreateMissionCommandApiBehaviorDto) GetIgnoreItineraryComputationFailures() bool`

GetIgnoreItineraryComputationFailures returns the IgnoreItineraryComputationFailures field if non-nil, zero value otherwise.

### GetIgnoreItineraryComputationFailuresOk

`func (o *CreateMissionCommandApiBehaviorDto) GetIgnoreItineraryComputationFailuresOk() (*bool, bool)`

GetIgnoreItineraryComputationFailuresOk returns a tuple with the IgnoreItineraryComputationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreItineraryComputationFailures

`func (o *CreateMissionCommandApiBehaviorDto) SetIgnoreItineraryComputationFailures(v bool)`

SetIgnoreItineraryComputationFailures sets IgnoreItineraryComputationFailures field to given value.

### HasIgnoreItineraryComputationFailures

`func (o *CreateMissionCommandApiBehaviorDto) HasIgnoreItineraryComputationFailures() bool`

HasIgnoreItineraryComputationFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


