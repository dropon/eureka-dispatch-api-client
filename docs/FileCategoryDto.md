# FileCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The unique business code for the file category. | [optional] 
**Label** | Pointer to **string** | The file category label | [optional] 
**IsVisibleForOrderers** | Pointer to **bool** | Defines whether the file category is visible for orderers | [optional] 
**IsVisibleForDrivers** | Pointer to **bool** | Defines whether the file category is visible for drivers. | [optional] 
**IsBillingReceipt** | Pointer to **bool** | Defines whether thi file category is a billing receipt. | [optional] 
**SubStateCode** | Pointer to **string** | The substate triggered by the file category | [optional] 

## Methods

### NewFileCategoryDto

`func NewFileCategoryDto() *FileCategoryDto`

NewFileCategoryDto instantiates a new FileCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCategoryDtoWithDefaults

`func NewFileCategoryDtoWithDefaults() *FileCategoryDto`

NewFileCategoryDtoWithDefaults instantiates a new FileCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FileCategoryDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FileCategoryDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FileCategoryDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FileCategoryDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *FileCategoryDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FileCategoryDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FileCategoryDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FileCategoryDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsVisibleForOrderers

`func (o *FileCategoryDto) GetIsVisibleForOrderers() bool`

GetIsVisibleForOrderers returns the IsVisibleForOrderers field if non-nil, zero value otherwise.

### GetIsVisibleForOrderersOk

`func (o *FileCategoryDto) GetIsVisibleForOrderersOk() (*bool, bool)`

GetIsVisibleForOrderersOk returns a tuple with the IsVisibleForOrderers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisibleForOrderers

`func (o *FileCategoryDto) SetIsVisibleForOrderers(v bool)`

SetIsVisibleForOrderers sets IsVisibleForOrderers field to given value.

### HasIsVisibleForOrderers

`func (o *FileCategoryDto) HasIsVisibleForOrderers() bool`

HasIsVisibleForOrderers returns a boolean if a field has been set.

### GetIsVisibleForDrivers

`func (o *FileCategoryDto) GetIsVisibleForDrivers() bool`

GetIsVisibleForDrivers returns the IsVisibleForDrivers field if non-nil, zero value otherwise.

### GetIsVisibleForDriversOk

`func (o *FileCategoryDto) GetIsVisibleForDriversOk() (*bool, bool)`

GetIsVisibleForDriversOk returns a tuple with the IsVisibleForDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisibleForDrivers

`func (o *FileCategoryDto) SetIsVisibleForDrivers(v bool)`

SetIsVisibleForDrivers sets IsVisibleForDrivers field to given value.

### HasIsVisibleForDrivers

`func (o *FileCategoryDto) HasIsVisibleForDrivers() bool`

HasIsVisibleForDrivers returns a boolean if a field has been set.

### GetIsBillingReceipt

`func (o *FileCategoryDto) GetIsBillingReceipt() bool`

GetIsBillingReceipt returns the IsBillingReceipt field if non-nil, zero value otherwise.

### GetIsBillingReceiptOk

`func (o *FileCategoryDto) GetIsBillingReceiptOk() (*bool, bool)`

GetIsBillingReceiptOk returns a tuple with the IsBillingReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingReceipt

`func (o *FileCategoryDto) SetIsBillingReceipt(v bool)`

SetIsBillingReceipt sets IsBillingReceipt field to given value.

### HasIsBillingReceipt

`func (o *FileCategoryDto) HasIsBillingReceipt() bool`

HasIsBillingReceipt returns a boolean if a field has been set.

### GetSubStateCode

`func (o *FileCategoryDto) GetSubStateCode() string`

GetSubStateCode returns the SubStateCode field if non-nil, zero value otherwise.

### GetSubStateCodeOk

`func (o *FileCategoryDto) GetSubStateCodeOk() (*string, bool)`

GetSubStateCodeOk returns a tuple with the SubStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStateCode

`func (o *FileCategoryDto) SetSubStateCode(v string)`

SetSubStateCode sets SubStateCode field to given value.

### HasSubStateCode

`func (o *FileCategoryDto) HasSubStateCode() bool`

HasSubStateCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


