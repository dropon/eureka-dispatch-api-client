# UpdateTransportBaseDtoSubServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeAction** | Pointer to **string** | Merge action to perform : add, update, remove. | [optional] 
**Uid** | Pointer to **string** | This property is an entity identifier. | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ForcedSellQuantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**ForcedSellUnitPrice** | Pointer to **float64** | To use default sell prices, unset (set to null) forced sell unit price, forced sell price and forced sell price in reference currency (if applicable).    If the forced sell unit price is set, the forced sell price and the forced sell price in reference currency (if applicable) must be unset.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellPrice** | Pointer to **float64** | If the forced sell price is set, the forced sell unit price and the forced sell price in reference currency (if applicable) must be unset.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellPriceInReferenceCurrency** | Pointer to **float64** | Forced sell price can be specified in reference currency when it is different from customer&#39;s currency    and when this option is available in general settings.    If the forced sell price in reference currency is set, the forced sell unit price and the forced sell price must be unset.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellVatCode** | Pointer to **string** | The following permission(s) are required to access this property:  Set VAT in orders. | [optional] 
**ForcedBuyQuantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**ForcedBuyUnitPrice** | Pointer to **float64** | To use default buy prices, unset (set to null) forced buy unit price, forced buy price and forced buy price in reference currency (if applicable).    If the forced buy unit price is set, the forced buy price and the forced buy price in reference currency (if applicable) must be unset.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyPrice** | Pointer to **float64** | If the forced buy price is set, the forced buy unit price and the forced buy price in reference currency (if applicable) must be unset.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyPriceInReferenceCurrency** | Pointer to **float64** | Forced buy price can be specified in reference currency when it is different from customer&#39;s currency    and when this option is available in general settings.    If the forced buy price in reference currency is set, the forced buy unit price and the forced buy price must be unset.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyVatCode** | Pointer to **string** | The following permission(s) are required to access this property:  Set VAT in orders. | [optional] 

## Methods

### NewUpdateTransportBaseDtoSubServiceDto

`func NewUpdateTransportBaseDtoSubServiceDto() *UpdateTransportBaseDtoSubServiceDto`

NewUpdateTransportBaseDtoSubServiceDto instantiates a new UpdateTransportBaseDtoSubServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransportBaseDtoSubServiceDtoWithDefaults

`func NewUpdateTransportBaseDtoSubServiceDtoWithDefaults() *UpdateTransportBaseDtoSubServiceDto`

NewUpdateTransportBaseDtoSubServiceDtoWithDefaults instantiates a new UpdateTransportBaseDtoSubServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeAction

`func (o *UpdateTransportBaseDtoSubServiceDto) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *UpdateTransportBaseDtoSubServiceDto) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *UpdateTransportBaseDtoSubServiceDto) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetUid

`func (o *UpdateTransportBaseDtoSubServiceDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UpdateTransportBaseDtoSubServiceDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UpdateTransportBaseDtoSubServiceDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *UpdateTransportBaseDtoSubServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateTransportBaseDtoSubServiceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateTransportBaseDtoSubServiceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetForcedSellQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellQuantity() float64`

GetForcedSellQuantity returns the ForcedSellQuantity field if non-nil, zero value otherwise.

### GetForcedSellQuantityOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellQuantityOk() (*float64, bool)`

GetForcedSellQuantityOk returns a tuple with the ForcedSellQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedSellQuantity(v float64)`

SetForcedSellQuantity sets ForcedSellQuantity field to given value.

### HasForcedSellQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedSellQuantity() bool`

HasForcedSellQuantity returns a boolean if a field has been set.

### GetForcedSellUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellUnitPrice() float64`

GetForcedSellUnitPrice returns the ForcedSellUnitPrice field if non-nil, zero value otherwise.

### GetForcedSellUnitPriceOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellUnitPriceOk() (*float64, bool)`

GetForcedSellUnitPriceOk returns a tuple with the ForcedSellUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedSellUnitPrice(v float64)`

SetForcedSellUnitPrice sets ForcedSellUnitPrice field to given value.

### HasForcedSellUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedSellUnitPrice() bool`

HasForcedSellUnitPrice returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetForcedSellPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellPriceInReferenceCurrency() float64`

GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedSellPriceInReferenceCurrencyOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedSellPriceInReferenceCurrency(v float64)`

SetForcedSellPriceInReferenceCurrency sets ForcedSellPriceInReferenceCurrency field to given value.

### HasForcedSellPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedSellPriceInReferenceCurrency() bool`

HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedSellVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellVatCode() string`

GetForcedSellVatCode returns the ForcedSellVatCode field if non-nil, zero value otherwise.

### GetForcedSellVatCodeOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedSellVatCodeOk() (*string, bool)`

GetForcedSellVatCodeOk returns a tuple with the ForcedSellVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedSellVatCode(v string)`

SetForcedSellVatCode sets ForcedSellVatCode field to given value.

### HasForcedSellVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedSellVatCode() bool`

HasForcedSellVatCode returns a boolean if a field has been set.

### GetForcedBuyQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyQuantity() float64`

GetForcedBuyQuantity returns the ForcedBuyQuantity field if non-nil, zero value otherwise.

### GetForcedBuyQuantityOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyQuantityOk() (*float64, bool)`

GetForcedBuyQuantityOk returns a tuple with the ForcedBuyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedBuyQuantity(v float64)`

SetForcedBuyQuantity sets ForcedBuyQuantity field to given value.

### HasForcedBuyQuantity

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedBuyQuantity() bool`

HasForcedBuyQuantity returns a boolean if a field has been set.

### GetForcedBuyUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyUnitPrice() float64`

GetForcedBuyUnitPrice returns the ForcedBuyUnitPrice field if non-nil, zero value otherwise.

### GetForcedBuyUnitPriceOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyUnitPriceOk() (*float64, bool)`

GetForcedBuyUnitPriceOk returns a tuple with the ForcedBuyUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedBuyUnitPrice(v float64)`

SetForcedBuyUnitPrice sets ForcedBuyUnitPrice field to given value.

### HasForcedBuyUnitPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedBuyUnitPrice() bool`

HasForcedBuyUnitPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyPriceInReferenceCurrency() float64`

GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedBuyPriceInReferenceCurrencyOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedBuyPriceInReferenceCurrency(v float64)`

SetForcedBuyPriceInReferenceCurrency sets ForcedBuyPriceInReferenceCurrency field to given value.

### HasForcedBuyPriceInReferenceCurrency

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedBuyPriceInReferenceCurrency() bool`

HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedBuyVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyVatCode() string`

GetForcedBuyVatCode returns the ForcedBuyVatCode field if non-nil, zero value otherwise.

### GetForcedBuyVatCodeOk

`func (o *UpdateTransportBaseDtoSubServiceDto) GetForcedBuyVatCodeOk() (*string, bool)`

GetForcedBuyVatCodeOk returns a tuple with the ForcedBuyVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) SetForcedBuyVatCode(v string)`

SetForcedBuyVatCode sets ForcedBuyVatCode field to given value.

### HasForcedBuyVatCode

`func (o *UpdateTransportBaseDtoSubServiceDto) HasForcedBuyVatCode() bool`

HasForcedBuyVatCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


