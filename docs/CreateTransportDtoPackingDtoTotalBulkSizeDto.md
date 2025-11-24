# CreateTransportDtoPackingDtoTotalBulkSizeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitCode** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float64** | The bulk size for the specified unit code. Ignored is \&quot;IsAutoComputed\&quot; is set to true.    This number can have up to 3 decimals. | [optional] 
**IsAutoComputed** | Pointer to **bool** | If true, the value will be ignored, and calculated from individual packages. | [optional] 

## Methods

### NewCreateTransportDtoPackingDtoTotalBulkSizeDto

`func NewCreateTransportDtoPackingDtoTotalBulkSizeDto() *CreateTransportDtoPackingDtoTotalBulkSizeDto`

NewCreateTransportDtoPackingDtoTotalBulkSizeDto instantiates a new CreateTransportDtoPackingDtoTotalBulkSizeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoPackingDtoTotalBulkSizeDtoWithDefaults

`func NewCreateTransportDtoPackingDtoTotalBulkSizeDtoWithDefaults() *CreateTransportDtoPackingDtoTotalBulkSizeDto`

NewCreateTransportDtoPackingDtoTotalBulkSizeDtoWithDefaults instantiates a new CreateTransportDtoPackingDtoTotalBulkSizeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitCode

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetValue

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsAutoComputed

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetIsAutoComputed() bool`

GetIsAutoComputed returns the IsAutoComputed field if non-nil, zero value otherwise.

### GetIsAutoComputedOk

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) GetIsAutoComputedOk() (*bool, bool)`

GetIsAutoComputedOk returns a tuple with the IsAutoComputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoComputed

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) SetIsAutoComputed(v bool)`

SetIsAutoComputed sets IsAutoComputed field to given value.

### HasIsAutoComputed

`func (o *CreateTransportDtoPackingDtoTotalBulkSizeDto) HasIsAutoComputed() bool`

HasIsAutoComputed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


