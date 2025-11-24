# CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PricingPathUnitPricingSlotEquivalenceDto**](PricingPathUnitPricingSlotEquivalenceDto.md) | List of items included in the capped collection | [optional] 
**HasAdditionalResult** | Pointer to **bool** | Indicates if the included data was partially returned because capped collection size is reached.  In this case, additional result can be retrieved from another route using pagination. | [optional] 

## Methods

### NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto

`func NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto() *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto`

NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto instantiates a new CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDtoWithDefaults

`func NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDtoWithDefaults() *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto`

NewCappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDtoWithDefaults instantiates a new CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) GetData() []PricingPathUnitPricingSlotEquivalenceDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) GetDataOk() (*[]PricingPathUnitPricingSlotEquivalenceDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) SetData(v []PricingPathUnitPricingSlotEquivalenceDto)`

SetData sets Data field to given value.

### HasData

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) GetHasAdditionalResult() bool`

GetHasAdditionalResult returns the HasAdditionalResult field if non-nil, zero value otherwise.

### GetHasAdditionalResultOk

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) GetHasAdditionalResultOk() (*bool, bool)`

GetHasAdditionalResultOk returns a tuple with the HasAdditionalResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) SetHasAdditionalResult(v bool)`

SetHasAdditionalResult sets HasAdditionalResult field to given value.

### HasHasAdditionalResult

`func (o *CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto) HasHasAdditionalResult() bool`

HasHasAdditionalResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


