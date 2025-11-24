# UpdateTransportDtoApiBehaviorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipPackageNatureDefaultValuesApplication** | Pointer to **bool** | This option is used to avoid copying the default values from the package nature.  By default, values are copied from the package nature when they are not provided :   quantity, package code, dimensions (height, width, length), sizes, references. | [optional] 
**TryGeocodeTransportSteps** | Pointer to **bool** | Try to geocode transport steps if the step is not yet geocoded.    This is only done if the geolocation service is available.    By default, geocoding is not attempted. | [optional] 
**UpdateTransportOperationComment** | Pointer to **string** | Comment for the update operation, will be added in transport history if specified. | [optional] 

## Methods

### NewUpdateTransportDtoApiBehaviorDto

`func NewUpdateTransportDtoApiBehaviorDto() *UpdateTransportDtoApiBehaviorDto`

NewUpdateTransportDtoApiBehaviorDto instantiates a new UpdateTransportDtoApiBehaviorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransportDtoApiBehaviorDtoWithDefaults

`func NewUpdateTransportDtoApiBehaviorDtoWithDefaults() *UpdateTransportDtoApiBehaviorDto`

NewUpdateTransportDtoApiBehaviorDtoWithDefaults instantiates a new UpdateTransportDtoApiBehaviorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipPackageNatureDefaultValuesApplication

`func (o *UpdateTransportDtoApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplication() bool`

GetSkipPackageNatureDefaultValuesApplication returns the SkipPackageNatureDefaultValuesApplication field if non-nil, zero value otherwise.

### GetSkipPackageNatureDefaultValuesApplicationOk

`func (o *UpdateTransportDtoApiBehaviorDto) GetSkipPackageNatureDefaultValuesApplicationOk() (*bool, bool)`

GetSkipPackageNatureDefaultValuesApplicationOk returns a tuple with the SkipPackageNatureDefaultValuesApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPackageNatureDefaultValuesApplication

`func (o *UpdateTransportDtoApiBehaviorDto) SetSkipPackageNatureDefaultValuesApplication(v bool)`

SetSkipPackageNatureDefaultValuesApplication sets SkipPackageNatureDefaultValuesApplication field to given value.

### HasSkipPackageNatureDefaultValuesApplication

`func (o *UpdateTransportDtoApiBehaviorDto) HasSkipPackageNatureDefaultValuesApplication() bool`

HasSkipPackageNatureDefaultValuesApplication returns a boolean if a field has been set.

### GetTryGeocodeTransportSteps

`func (o *UpdateTransportDtoApiBehaviorDto) GetTryGeocodeTransportSteps() bool`

GetTryGeocodeTransportSteps returns the TryGeocodeTransportSteps field if non-nil, zero value otherwise.

### GetTryGeocodeTransportStepsOk

`func (o *UpdateTransportDtoApiBehaviorDto) GetTryGeocodeTransportStepsOk() (*bool, bool)`

GetTryGeocodeTransportStepsOk returns a tuple with the TryGeocodeTransportSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryGeocodeTransportSteps

`func (o *UpdateTransportDtoApiBehaviorDto) SetTryGeocodeTransportSteps(v bool)`

SetTryGeocodeTransportSteps sets TryGeocodeTransportSteps field to given value.

### HasTryGeocodeTransportSteps

`func (o *UpdateTransportDtoApiBehaviorDto) HasTryGeocodeTransportSteps() bool`

HasTryGeocodeTransportSteps returns a boolean if a field has been set.

### GetUpdateTransportOperationComment

`func (o *UpdateTransportDtoApiBehaviorDto) GetUpdateTransportOperationComment() string`

GetUpdateTransportOperationComment returns the UpdateTransportOperationComment field if non-nil, zero value otherwise.

### GetUpdateTransportOperationCommentOk

`func (o *UpdateTransportDtoApiBehaviorDto) GetUpdateTransportOperationCommentOk() (*string, bool)`

GetUpdateTransportOperationCommentOk returns a tuple with the UpdateTransportOperationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTransportOperationComment

`func (o *UpdateTransportDtoApiBehaviorDto) SetUpdateTransportOperationComment(v string)`

SetUpdateTransportOperationComment sets UpdateTransportOperationComment field to given value.

### HasUpdateTransportOperationComment

`func (o *UpdateTransportDtoApiBehaviorDto) HasUpdateTransportOperationComment() bool`

HasUpdateTransportOperationComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


