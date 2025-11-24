# TransportEntryDryRunDtoSubServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 
**SellQuantity** | Pointer to **float64** |  | [optional] 
**SellPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**SellUnitPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsSellGlobalPricingEnabled** | Pointer to **bool** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsSellQuantityForced** | Pointer to **bool** |  | [optional] 
**SellVatCode** | Pointer to **string** |  | [optional] 
**IsSellVatCodeForced** | Pointer to **bool** |  | [optional] 
**BuyQuantity** | Pointer to **float64** |  | [optional] 
**BuyPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyUnitPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsBuyGlobalPricingEnabled** | Pointer to **bool** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsBuyQuantityForced** | Pointer to **bool** |  | [optional] 
**BuyVatCode** | Pointer to **string** |  | [optional] 
**IsBuyVatCodeForced** | Pointer to **bool** |  | [optional] 
**SellPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s sell price forced in reference currency, otherwise sell price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**SellUnitPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s sell unit price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s buy price forced in reference currency, otherwise buy price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyUnitPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s buy unit price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 

## Methods

### NewTransportEntryDryRunDtoSubServiceDto

`func NewTransportEntryDryRunDtoSubServiceDto() *TransportEntryDryRunDtoSubServiceDto`

NewTransportEntryDryRunDtoSubServiceDto instantiates a new TransportEntryDryRunDtoSubServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEntryDryRunDtoSubServiceDtoWithDefaults

`func NewTransportEntryDryRunDtoSubServiceDtoWithDefaults() *TransportEntryDryRunDtoSubServiceDto`

NewTransportEntryDryRunDtoSubServiceDtoWithDefaults instantiates a new TransportEntryDryRunDtoSubServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransportEntryDryRunDtoSubServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransportEntryDryRunDtoSubServiceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransportEntryDryRunDtoSubServiceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUnitCode

`func (o *TransportEntryDryRunDtoSubServiceDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *TransportEntryDryRunDtoSubServiceDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *TransportEntryDryRunDtoSubServiceDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetSellQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellQuantity() float64`

GetSellQuantity returns the SellQuantity field if non-nil, zero value otherwise.

### GetSellQuantityOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellQuantityOk() (*float64, bool)`

GetSellQuantityOk returns a tuple with the SellQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellQuantity(v float64)`

SetSellQuantity sets SellQuantity field to given value.

### HasSellQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellQuantity() bool`

HasSellQuantity returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetSellUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellUnitPrice() float64`

GetSellUnitPrice returns the SellUnitPrice field if non-nil, zero value otherwise.

### GetSellUnitPriceOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellUnitPriceOk() (*float64, bool)`

GetSellUnitPriceOk returns a tuple with the SellUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellUnitPrice(v float64)`

SetSellUnitPrice sets SellUnitPrice field to given value.

### HasSellUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellUnitPrice() bool`

HasSellUnitPrice returns a boolean if a field has been set.

### GetIsSellGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellGlobalPricingEnabled() bool`

GetIsSellGlobalPricingEnabled returns the IsSellGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsSellGlobalPricingEnabledOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellGlobalPricingEnabledOk() (*bool, bool)`

GetIsSellGlobalPricingEnabledOk returns a tuple with the IsSellGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsSellGlobalPricingEnabled(v bool)`

SetIsSellGlobalPricingEnabled sets IsSellGlobalPricingEnabled field to given value.

### HasIsSellGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsSellGlobalPricingEnabled() bool`

HasIsSellGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsSellQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellQuantityForced() bool`

GetIsSellQuantityForced returns the IsSellQuantityForced field if non-nil, zero value otherwise.

### GetIsSellQuantityForcedOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellQuantityForcedOk() (*bool, bool)`

GetIsSellQuantityForcedOk returns a tuple with the IsSellQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsSellQuantityForced(v bool)`

SetIsSellQuantityForced sets IsSellQuantityForced field to given value.

### HasIsSellQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsSellQuantityForced() bool`

HasIsSellQuantityForced returns a boolean if a field has been set.

### GetSellVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellVatCode() string`

GetSellVatCode returns the SellVatCode field if non-nil, zero value otherwise.

### GetSellVatCodeOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellVatCodeOk() (*string, bool)`

GetSellVatCodeOk returns a tuple with the SellVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellVatCode(v string)`

SetSellVatCode sets SellVatCode field to given value.

### HasSellVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellVatCode() bool`

HasSellVatCode returns a boolean if a field has been set.

### GetIsSellVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellVatCodeForced() bool`

GetIsSellVatCodeForced returns the IsSellVatCodeForced field if non-nil, zero value otherwise.

### GetIsSellVatCodeForcedOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsSellVatCodeForcedOk() (*bool, bool)`

GetIsSellVatCodeForcedOk returns a tuple with the IsSellVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsSellVatCodeForced(v bool)`

SetIsSellVatCodeForced sets IsSellVatCodeForced field to given value.

### HasIsSellVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsSellVatCodeForced() bool`

HasIsSellVatCodeForced returns a boolean if a field has been set.

### GetBuyQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyQuantity() float64`

GetBuyQuantity returns the BuyQuantity field if non-nil, zero value otherwise.

### GetBuyQuantityOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyQuantityOk() (*float64, bool)`

GetBuyQuantityOk returns a tuple with the BuyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyQuantity(v float64)`

SetBuyQuantity sets BuyQuantity field to given value.

### HasBuyQuantity

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyQuantity() bool`

HasBuyQuantity returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetBuyUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyUnitPrice() float64`

GetBuyUnitPrice returns the BuyUnitPrice field if non-nil, zero value otherwise.

### GetBuyUnitPriceOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyUnitPriceOk() (*float64, bool)`

GetBuyUnitPriceOk returns a tuple with the BuyUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyUnitPrice(v float64)`

SetBuyUnitPrice sets BuyUnitPrice field to given value.

### HasBuyUnitPrice

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyUnitPrice() bool`

HasBuyUnitPrice returns a boolean if a field has been set.

### GetIsBuyGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyGlobalPricingEnabled() bool`

GetIsBuyGlobalPricingEnabled returns the IsBuyGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsBuyGlobalPricingEnabledOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyGlobalPricingEnabledOk() (*bool, bool)`

GetIsBuyGlobalPricingEnabledOk returns a tuple with the IsBuyGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsBuyGlobalPricingEnabled(v bool)`

SetIsBuyGlobalPricingEnabled sets IsBuyGlobalPricingEnabled field to given value.

### HasIsBuyGlobalPricingEnabled

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsBuyGlobalPricingEnabled() bool`

HasIsBuyGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsBuyQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyQuantityForced() bool`

GetIsBuyQuantityForced returns the IsBuyQuantityForced field if non-nil, zero value otherwise.

### GetIsBuyQuantityForcedOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyQuantityForcedOk() (*bool, bool)`

GetIsBuyQuantityForcedOk returns a tuple with the IsBuyQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsBuyQuantityForced(v bool)`

SetIsBuyQuantityForced sets IsBuyQuantityForced field to given value.

### HasIsBuyQuantityForced

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsBuyQuantityForced() bool`

HasIsBuyQuantityForced returns a boolean if a field has been set.

### GetBuyVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyVatCode() string`

GetBuyVatCode returns the BuyVatCode field if non-nil, zero value otherwise.

### GetBuyVatCodeOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyVatCodeOk() (*string, bool)`

GetBuyVatCodeOk returns a tuple with the BuyVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyVatCode(v string)`

SetBuyVatCode sets BuyVatCode field to given value.

### HasBuyVatCode

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyVatCode() bool`

HasBuyVatCode returns a boolean if a field has been set.

### GetIsBuyVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyVatCodeForced() bool`

GetIsBuyVatCodeForced returns the IsBuyVatCodeForced field if non-nil, zero value otherwise.

### GetIsBuyVatCodeForcedOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetIsBuyVatCodeForcedOk() (*bool, bool)`

GetIsBuyVatCodeForcedOk returns a tuple with the IsBuyVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) SetIsBuyVatCodeForced(v bool)`

SetIsBuyVatCodeForced sets IsBuyVatCodeForced field to given value.

### HasIsBuyVatCodeForced

`func (o *TransportEntryDryRunDtoSubServiceDto) HasIsBuyVatCodeForced() bool`

HasIsBuyVatCodeForced returns a boolean if a field has been set.

### GetSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellPriceInReferenceCurrency() float64`

GetSellPriceInReferenceCurrency returns the SellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellPriceInReferenceCurrencyOk returns a tuple with the SellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellPriceInReferenceCurrency(v float64)`

SetSellPriceInReferenceCurrency sets SellPriceInReferenceCurrency field to given value.

### HasSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellPriceInReferenceCurrency() bool`

HasSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetSellUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellUnitPriceInReferenceCurrency() float64`

GetSellUnitPriceInReferenceCurrency returns the SellUnitPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellUnitPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetSellUnitPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellUnitPriceInReferenceCurrencyOk returns a tuple with the SellUnitPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) SetSellUnitPriceInReferenceCurrency(v float64)`

SetSellUnitPriceInReferenceCurrency sets SellUnitPriceInReferenceCurrency field to given value.

### HasSellUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) HasSellUnitPriceInReferenceCurrency() bool`

HasSellUnitPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyPriceInReferenceCurrency() float64`

GetBuyPriceInReferenceCurrency returns the BuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyPriceInReferenceCurrencyOk returns a tuple with the BuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyPriceInReferenceCurrency(v float64)`

SetBuyPriceInReferenceCurrency sets BuyPriceInReferenceCurrency field to given value.

### HasBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyPriceInReferenceCurrency() bool`

HasBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyUnitPriceInReferenceCurrency() float64`

GetBuyUnitPriceInReferenceCurrency returns the BuyUnitPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyUnitPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDtoSubServiceDto) GetBuyUnitPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyUnitPriceInReferenceCurrencyOk returns a tuple with the BuyUnitPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) SetBuyUnitPriceInReferenceCurrency(v float64)`

SetBuyUnitPriceInReferenceCurrency sets BuyUnitPriceInReferenceCurrency field to given value.

### HasBuyUnitPriceInReferenceCurrency

`func (o *TransportEntryDryRunDtoSubServiceDto) HasBuyUnitPriceInReferenceCurrency() bool`

HasBuyUnitPriceInReferenceCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


