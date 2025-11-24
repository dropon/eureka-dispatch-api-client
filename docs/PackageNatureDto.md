# PackageNatureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageNatureId** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**PackageTypologyCode** | Pointer to **string** |  | [optional] 
**PackageFamilyCode** | Pointer to **string** |  | [optional] 
**PackageFamilyLabel** | Pointer to **string** |  | [optional] 
**ParentPackageNatureId** | Pointer to **int32** |  | [optional] 
**DefaultQuantity** | Pointer to **float64** |  | [optional] 
**DimensionUnit** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **float64** |  | [optional] 
**Width** | Pointer to **float64** |  | [optional] 
**Length** | Pointer to **float64** |  | [optional] 
**IsOnlyOnePackageReferenceMandatory** | Pointer to **bool** |  | [optional] 
**References** | Pointer to [**CappedCollectionDtoPackageNatureReferenceDto**](CappedCollectionDtoPackageNatureReferenceDto.md) |  | [optional] 
**BulkSizes** | Pointer to [**CappedCollectionDtoPackageNatureBulkSizeDto**](CappedCollectionDtoPackageNatureBulkSizeDto.md) |  | [optional] 

## Methods

### NewPackageNatureDto

`func NewPackageNatureDto() *PackageNatureDto`

NewPackageNatureDto instantiates a new PackageNatureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageNatureDtoWithDefaults

`func NewPackageNatureDtoWithDefaults() *PackageNatureDto`

NewPackageNatureDtoWithDefaults instantiates a new PackageNatureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageNatureId

`func (o *PackageNatureDto) GetPackageNatureId() int32`

GetPackageNatureId returns the PackageNatureId field if non-nil, zero value otherwise.

### GetPackageNatureIdOk

`func (o *PackageNatureDto) GetPackageNatureIdOk() (*int32, bool)`

GetPackageNatureIdOk returns a tuple with the PackageNatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageNatureId

`func (o *PackageNatureDto) SetPackageNatureId(v int32)`

SetPackageNatureId sets PackageNatureId field to given value.

### HasPackageNatureId

`func (o *PackageNatureDto) HasPackageNatureId() bool`

HasPackageNatureId returns a boolean if a field has been set.

### GetCode

`func (o *PackageNatureDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PackageNatureDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PackageNatureDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PackageNatureDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *PackageNatureDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PackageNatureDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PackageNatureDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PackageNatureDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PackageNatureDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PackageNatureDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PackageNatureDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PackageNatureDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPackageTypologyCode

`func (o *PackageNatureDto) GetPackageTypologyCode() string`

GetPackageTypologyCode returns the PackageTypologyCode field if non-nil, zero value otherwise.

### GetPackageTypologyCodeOk

`func (o *PackageNatureDto) GetPackageTypologyCodeOk() (*string, bool)`

GetPackageTypologyCodeOk returns a tuple with the PackageTypologyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypologyCode

`func (o *PackageNatureDto) SetPackageTypologyCode(v string)`

SetPackageTypologyCode sets PackageTypologyCode field to given value.

### HasPackageTypologyCode

`func (o *PackageNatureDto) HasPackageTypologyCode() bool`

HasPackageTypologyCode returns a boolean if a field has been set.

### GetPackageFamilyCode

`func (o *PackageNatureDto) GetPackageFamilyCode() string`

GetPackageFamilyCode returns the PackageFamilyCode field if non-nil, zero value otherwise.

### GetPackageFamilyCodeOk

`func (o *PackageNatureDto) GetPackageFamilyCodeOk() (*string, bool)`

GetPackageFamilyCodeOk returns a tuple with the PackageFamilyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFamilyCode

`func (o *PackageNatureDto) SetPackageFamilyCode(v string)`

SetPackageFamilyCode sets PackageFamilyCode field to given value.

### HasPackageFamilyCode

`func (o *PackageNatureDto) HasPackageFamilyCode() bool`

HasPackageFamilyCode returns a boolean if a field has been set.

### GetPackageFamilyLabel

`func (o *PackageNatureDto) GetPackageFamilyLabel() string`

GetPackageFamilyLabel returns the PackageFamilyLabel field if non-nil, zero value otherwise.

### GetPackageFamilyLabelOk

`func (o *PackageNatureDto) GetPackageFamilyLabelOk() (*string, bool)`

GetPackageFamilyLabelOk returns a tuple with the PackageFamilyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFamilyLabel

`func (o *PackageNatureDto) SetPackageFamilyLabel(v string)`

SetPackageFamilyLabel sets PackageFamilyLabel field to given value.

### HasPackageFamilyLabel

`func (o *PackageNatureDto) HasPackageFamilyLabel() bool`

HasPackageFamilyLabel returns a boolean if a field has been set.

### GetParentPackageNatureId

`func (o *PackageNatureDto) GetParentPackageNatureId() int32`

GetParentPackageNatureId returns the ParentPackageNatureId field if non-nil, zero value otherwise.

### GetParentPackageNatureIdOk

`func (o *PackageNatureDto) GetParentPackageNatureIdOk() (*int32, bool)`

GetParentPackageNatureIdOk returns a tuple with the ParentPackageNatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPackageNatureId

`func (o *PackageNatureDto) SetParentPackageNatureId(v int32)`

SetParentPackageNatureId sets ParentPackageNatureId field to given value.

### HasParentPackageNatureId

`func (o *PackageNatureDto) HasParentPackageNatureId() bool`

HasParentPackageNatureId returns a boolean if a field has been set.

### GetDefaultQuantity

`func (o *PackageNatureDto) GetDefaultQuantity() float64`

GetDefaultQuantity returns the DefaultQuantity field if non-nil, zero value otherwise.

### GetDefaultQuantityOk

`func (o *PackageNatureDto) GetDefaultQuantityOk() (*float64, bool)`

GetDefaultQuantityOk returns a tuple with the DefaultQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuantity

`func (o *PackageNatureDto) SetDefaultQuantity(v float64)`

SetDefaultQuantity sets DefaultQuantity field to given value.

### HasDefaultQuantity

`func (o *PackageNatureDto) HasDefaultQuantity() bool`

HasDefaultQuantity returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *PackageNatureDto) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *PackageNatureDto) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *PackageNatureDto) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *PackageNatureDto) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetHeight

`func (o *PackageNatureDto) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *PackageNatureDto) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *PackageNatureDto) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *PackageNatureDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *PackageNatureDto) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PackageNatureDto) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PackageNatureDto) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PackageNatureDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetLength

`func (o *PackageNatureDto) GetLength() float64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PackageNatureDto) GetLengthOk() (*float64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PackageNatureDto) SetLength(v float64)`

SetLength sets Length field to given value.

### HasLength

`func (o *PackageNatureDto) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetIsOnlyOnePackageReferenceMandatory

`func (o *PackageNatureDto) GetIsOnlyOnePackageReferenceMandatory() bool`

GetIsOnlyOnePackageReferenceMandatory returns the IsOnlyOnePackageReferenceMandatory field if non-nil, zero value otherwise.

### GetIsOnlyOnePackageReferenceMandatoryOk

`func (o *PackageNatureDto) GetIsOnlyOnePackageReferenceMandatoryOk() (*bool, bool)`

GetIsOnlyOnePackageReferenceMandatoryOk returns a tuple with the IsOnlyOnePackageReferenceMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnlyOnePackageReferenceMandatory

`func (o *PackageNatureDto) SetIsOnlyOnePackageReferenceMandatory(v bool)`

SetIsOnlyOnePackageReferenceMandatory sets IsOnlyOnePackageReferenceMandatory field to given value.

### HasIsOnlyOnePackageReferenceMandatory

`func (o *PackageNatureDto) HasIsOnlyOnePackageReferenceMandatory() bool`

HasIsOnlyOnePackageReferenceMandatory returns a boolean if a field has been set.

### GetReferences

`func (o *PackageNatureDto) GetReferences() CappedCollectionDtoPackageNatureReferenceDto`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PackageNatureDto) GetReferencesOk() (*CappedCollectionDtoPackageNatureReferenceDto, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PackageNatureDto) SetReferences(v CappedCollectionDtoPackageNatureReferenceDto)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PackageNatureDto) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetBulkSizes

`func (o *PackageNatureDto) GetBulkSizes() CappedCollectionDtoPackageNatureBulkSizeDto`

GetBulkSizes returns the BulkSizes field if non-nil, zero value otherwise.

### GetBulkSizesOk

`func (o *PackageNatureDto) GetBulkSizesOk() (*CappedCollectionDtoPackageNatureBulkSizeDto, bool)`

GetBulkSizesOk returns a tuple with the BulkSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSizes

`func (o *PackageNatureDto) SetBulkSizes(v CappedCollectionDtoPackageNatureBulkSizeDto)`

SetBulkSizes sets BulkSizes field to given value.

### HasBulkSizes

`func (o *PackageNatureDto) HasBulkSizes() bool`

HasBulkSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


