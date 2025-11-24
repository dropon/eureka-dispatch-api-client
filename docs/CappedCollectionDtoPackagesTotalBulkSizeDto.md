# CappedCollectionDtoPackagesTotalBulkSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PackagesTotalBulkSizeDto**](PackagesTotalBulkSizeDto.md) | List of items included in the capped collection | [optional] 
**HasAdditionalResult** | Pointer to **bool** | Indicates if the included data was partially returned because capped collection size is reached.  In this case, additional result can be retrieved from another route using pagination. | [optional] 

## Methods

### NewCappedCollectionDtoPackagesTotalBulkSizeDto

`func NewCappedCollectionDtoPackagesTotalBulkSizeDto() *CappedCollectionDtoPackagesTotalBulkSizeDto`

NewCappedCollectionDtoPackagesTotalBulkSizeDto instantiates a new CappedCollectionDtoPackagesTotalBulkSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCappedCollectionDtoPackagesTotalBulkSizeDtoWithDefaults

`func NewCappedCollectionDtoPackagesTotalBulkSizeDtoWithDefaults() *CappedCollectionDtoPackagesTotalBulkSizeDto`

NewCappedCollectionDtoPackagesTotalBulkSizeDtoWithDefaults instantiates a new CappedCollectionDtoPackagesTotalBulkSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) GetData() []PackagesTotalBulkSizeDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) GetDataOk() (*[]PackagesTotalBulkSizeDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) SetData(v []PackagesTotalBulkSizeDto)`

SetData sets Data field to given value.

### HasData

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasAdditionalResult

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) GetHasAdditionalResult() bool`

GetHasAdditionalResult returns the HasAdditionalResult field if non-nil, zero value otherwise.

### GetHasAdditionalResultOk

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) GetHasAdditionalResultOk() (*bool, bool)`

GetHasAdditionalResultOk returns a tuple with the HasAdditionalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalResult

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) SetHasAdditionalResult(v bool)`

SetHasAdditionalResult sets HasAdditionalResult field to given value.

### HasHasAdditionalResult

`func (o *CappedCollectionDtoPackagesTotalBulkSizeDto) HasHasAdditionalResult() bool`

HasHasAdditionalResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


