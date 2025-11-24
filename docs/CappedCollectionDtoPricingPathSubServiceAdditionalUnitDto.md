# CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PricingPathSubServiceAdditionalUnitDto**](PricingPathSubServiceAdditionalUnitDto.md) | List of items included in the capped collection | [optional] 
**HasAdditionalResult** | Pointer to **bool** | Indicates if the included data was partially returned because capped collection size is reached.  In this case, additional result can be retrieved from another route using pagination. | [optional] 

## Methods

### NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDto

`func NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDto() *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto`

NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDto instantiates a new CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDtoWithDefaults

`func NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDtoWithDefaults() *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto`

NewCappedCollectionDtoPricingPathSubServiceAdditionalUnitDtoWithDefaults instantiates a new CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) GetData() []PricingPathSubServiceAdditionalUnitDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) GetDataOk() (*[]PricingPathSubServiceAdditionalUnitDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) SetData(v []PricingPathSubServiceAdditionalUnitDto)`

SetData sets Data field to given value.

### HasData

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) GetHasAdditionalResult() bool`

GetHasAdditionalResult returns the HasAdditionalResult field if non-nil, zero value otherwise.

### GetHasAdditionalResultOk

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) GetHasAdditionalResultOk() (*bool, bool)`

GetHasAdditionalResultOk returns a tuple with the HasAdditionalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) SetHasAdditionalResult(v bool)`

SetHasAdditionalResult sets HasAdditionalResult field to given value.

### HasHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto) HasHasAdditionalResult() bool`

HasHasAdditionalResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


