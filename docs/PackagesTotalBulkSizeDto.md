# PackagesTotalBulkSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitCode** | Pointer to **string** | The package bulk size unit (ie : KG, M3, ...) | [optional] 
**IsAutoComputed** | Pointer to **bool** | Indicates if the the value is ignored and calculated from individual packages. | [optional] 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewPackagesTotalBulkSizeDto

`func NewPackagesTotalBulkSizeDto() *PackagesTotalBulkSizeDto`

NewPackagesTotalBulkSizeDto instantiates a new PackagesTotalBulkSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesTotalBulkSizeDtoWithDefaults

`func NewPackagesTotalBulkSizeDtoWithDefaults() *PackagesTotalBulkSizeDto`

NewPackagesTotalBulkSizeDtoWithDefaults instantiates a new PackagesTotalBulkSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitCode

`func (o *PackagesTotalBulkSizeDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *PackagesTotalBulkSizeDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *PackagesTotalBulkSizeDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *PackagesTotalBulkSizeDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetIsAutoComputed

`func (o *PackagesTotalBulkSizeDto) GetIsAutoComputed() bool`

GetIsAutoComputed returns the IsAutoComputed field if non-nil, zero value otherwise.

### GetIsAutoComputedOk

`func (o *PackagesTotalBulkSizeDto) GetIsAutoComputedOk() (*bool, bool)`

GetIsAutoComputedOk returns a tuple with the IsAutoComputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoComputed

`func (o *PackagesTotalBulkSizeDto) SetIsAutoComputed(v bool)`

SetIsAutoComputed sets IsAutoComputed field to given value.

### HasIsAutoComputed

`func (o *PackagesTotalBulkSizeDto) HasIsAutoComputed() bool`

HasIsAutoComputed returns a boolean if a field has been set.

### GetValue

`func (o *PackagesTotalBulkSizeDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PackagesTotalBulkSizeDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PackagesTotalBulkSizeDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *PackagesTotalBulkSizeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


