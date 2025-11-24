# CappedCollectionDtoPackageNatureReferenceValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PackageNatureReferenceValueDto**](PackageNatureReferenceValueDto.md) | List of items included in the capped collection | [optional] 
**HasAdditionalResult** | Pointer to **bool** | Indicates if the included data was partially returned because capped collection size is reached.  In this case, additional result can be retrieved from another route using pagination. | [optional] 

## Methods

### NewCappedCollectionDtoPackageNatureReferenceValueDto

`func NewCappedCollectionDtoPackageNatureReferenceValueDto() *CappedCollectionDtoPackageNatureReferenceValueDto`

NewCappedCollectionDtoPackageNatureReferenceValueDto instantiates a new CappedCollectionDtoPackageNatureReferenceValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCappedCollectionDtoPackageNatureReferenceValueDtoWithDefaults

`func NewCappedCollectionDtoPackageNatureReferenceValueDtoWithDefaults() *CappedCollectionDtoPackageNatureReferenceValueDto`

NewCappedCollectionDtoPackageNatureReferenceValueDtoWithDefaults instantiates a new CappedCollectionDtoPackageNatureReferenceValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) GetData() []PackageNatureReferenceValueDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) GetDataOk() (*[]PackageNatureReferenceValueDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) SetData(v []PackageNatureReferenceValueDto)`

SetData sets Data field to given value.

### HasData

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasAdditionalResult

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) GetHasAdditionalResult() bool`

GetHasAdditionalResult returns the HasAdditionalResult field if non-nil, zero value otherwise.

### GetHasAdditionalResultOk

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) GetHasAdditionalResultOk() (*bool, bool)`

GetHasAdditionalResultOk returns a tuple with the HasAdditionalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalResult

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) SetHasAdditionalResult(v bool)`

SetHasAdditionalResult sets HasAdditionalResult field to given value.

### HasHasAdditionalResult

`func (o *CappedCollectionDtoPackageNatureReferenceValueDto) HasHasAdditionalResult() bool`

HasHasAdditionalResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


