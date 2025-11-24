# TransportSubServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 
**SellQuantity** | Pointer to **float64** |  | [optional] 
**SellUnitPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**SellPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedSellPriceInReferenceCurrency** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsSellGlobalPricingEnabled** | Pointer to **bool** | Sell global pricing is disabled when sell price or sell unit price is forced.    The following permission(s) are required to access this property:  See prices. | [optional] 
**IsSellQuantityForced** | Pointer to **bool** |  | [optional] 
**SellVatCode** | Pointer to **string** |  | [optional] 
**IsSellVatCodeForced** | Pointer to **bool** |  | [optional] 
**BuyQuantity** | Pointer to **float64** |  | [optional] 
**BuyUnitPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedBuyPriceInReferenceCurrency** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**IsBuyGlobalPricingEnabled** | Pointer to **bool** | Buy global pricing is disabled when buy price or buy unit price is forced.    The following permission(s) are required to access this property:  See prices. | [optional] 
**IsBuyQuantityForced** | Pointer to **bool** |  | [optional] 
**BuyVatCode** | Pointer to **string** |  | [optional] 
**IsBuyVatCodeForced** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransportSubServiceDto

`func NewTransportSubServiceDto() *TransportSubServiceDto`

NewTransportSubServiceDto instantiates a new TransportSubServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportSubServiceDtoWithDefaults

`func NewTransportSubServiceDtoWithDefaults() *TransportSubServiceDto`

NewTransportSubServiceDtoWithDefaults instantiates a new TransportSubServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *TransportSubServiceDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TransportSubServiceDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TransportSubServiceDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TransportSubServiceDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *TransportSubServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransportSubServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransportSubServiceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransportSubServiceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *TransportSubServiceDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TransportSubServiceDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TransportSubServiceDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TransportSubServiceDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUnitCode

`func (o *TransportSubServiceDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *TransportSubServiceDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *TransportSubServiceDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *TransportSubServiceDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetSellQuantity

`func (o *TransportSubServiceDto) GetSellQuantity() float64`

GetSellQuantity returns the SellQuantity field if non-nil, zero value otherwise.

### GetSellQuantityOk

`func (o *TransportSubServiceDto) GetSellQuantityOk() (*float64, bool)`

GetSellQuantityOk returns a tuple with the SellQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellQuantity

`func (o *TransportSubServiceDto) SetSellQuantity(v float64)`

SetSellQuantity sets SellQuantity field to given value.

### HasSellQuantity

`func (o *TransportSubServiceDto) HasSellQuantity() bool`

HasSellQuantity returns a boolean if a field has been set.

### GetSellUnitPrice

`func (o *TransportSubServiceDto) GetSellUnitPrice() float64`

GetSellUnitPrice returns the SellUnitPrice field if non-nil, zero value otherwise.

### GetSellUnitPriceOk

`func (o *TransportSubServiceDto) GetSellUnitPriceOk() (*float64, bool)`

GetSellUnitPriceOk returns a tuple with the SellUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellUnitPrice

`func (o *TransportSubServiceDto) SetSellUnitPrice(v float64)`

SetSellUnitPrice sets SellUnitPrice field to given value.

### HasSellUnitPrice

`func (o *TransportSubServiceDto) HasSellUnitPrice() bool`

HasSellUnitPrice returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportSubServiceDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportSubServiceDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportSubServiceDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportSubServiceDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetForcedSellPriceInReferenceCurrency

`func (o *TransportSubServiceDto) GetForcedSellPriceInReferenceCurrency() float64`

GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedSellPriceInReferenceCurrencyOk

`func (o *TransportSubServiceDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPriceInReferenceCurrency

`func (o *TransportSubServiceDto) SetForcedSellPriceInReferenceCurrency(v float64)`

SetForcedSellPriceInReferenceCurrency sets ForcedSellPriceInReferenceCurrency field to given value.

### HasForcedSellPriceInReferenceCurrency

`func (o *TransportSubServiceDto) HasForcedSellPriceInReferenceCurrency() bool`

HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetIsSellGlobalPricingEnabled

`func (o *TransportSubServiceDto) GetIsSellGlobalPricingEnabled() bool`

GetIsSellGlobalPricingEnabled returns the IsSellGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsSellGlobalPricingEnabledOk

`func (o *TransportSubServiceDto) GetIsSellGlobalPricingEnabledOk() (*bool, bool)`

GetIsSellGlobalPricingEnabledOk returns a tuple with the IsSellGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellGlobalPricingEnabled

`func (o *TransportSubServiceDto) SetIsSellGlobalPricingEnabled(v bool)`

SetIsSellGlobalPricingEnabled sets IsSellGlobalPricingEnabled field to given value.

### HasIsSellGlobalPricingEnabled

`func (o *TransportSubServiceDto) HasIsSellGlobalPricingEnabled() bool`

HasIsSellGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsSellQuantityForced

`func (o *TransportSubServiceDto) GetIsSellQuantityForced() bool`

GetIsSellQuantityForced returns the IsSellQuantityForced field if non-nil, zero value otherwise.

### GetIsSellQuantityForcedOk

`func (o *TransportSubServiceDto) GetIsSellQuantityForcedOk() (*bool, bool)`

GetIsSellQuantityForcedOk returns a tuple with the IsSellQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellQuantityForced

`func (o *TransportSubServiceDto) SetIsSellQuantityForced(v bool)`

SetIsSellQuantityForced sets IsSellQuantityForced field to given value.

### HasIsSellQuantityForced

`func (o *TransportSubServiceDto) HasIsSellQuantityForced() bool`

HasIsSellQuantityForced returns a boolean if a field has been set.

### GetSellVatCode

`func (o *TransportSubServiceDto) GetSellVatCode() string`

GetSellVatCode returns the SellVatCode field if non-nil, zero value otherwise.

### GetSellVatCodeOk

`func (o *TransportSubServiceDto) GetSellVatCodeOk() (*string, bool)`

GetSellVatCodeOk returns a tuple with the SellVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellVatCode

`func (o *TransportSubServiceDto) SetSellVatCode(v string)`

SetSellVatCode sets SellVatCode field to given value.

### HasSellVatCode

`func (o *TransportSubServiceDto) HasSellVatCode() bool`

HasSellVatCode returns a boolean if a field has been set.

### GetIsSellVatCodeForced

`func (o *TransportSubServiceDto) GetIsSellVatCodeForced() bool`

GetIsSellVatCodeForced returns the IsSellVatCodeForced field if non-nil, zero value otherwise.

### GetIsSellVatCodeForcedOk

`func (o *TransportSubServiceDto) GetIsSellVatCodeForcedOk() (*bool, bool)`

GetIsSellVatCodeForcedOk returns a tuple with the IsSellVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellVatCodeForced

`func (o *TransportSubServiceDto) SetIsSellVatCodeForced(v bool)`

SetIsSellVatCodeForced sets IsSellVatCodeForced field to given value.

### HasIsSellVatCodeForced

`func (o *TransportSubServiceDto) HasIsSellVatCodeForced() bool`

HasIsSellVatCodeForced returns a boolean if a field has been set.

### GetBuyQuantity

`func (o *TransportSubServiceDto) GetBuyQuantity() float64`

GetBuyQuantity returns the BuyQuantity field if non-nil, zero value otherwise.

### GetBuyQuantityOk

`func (o *TransportSubServiceDto) GetBuyQuantityOk() (*float64, bool)`

GetBuyQuantityOk returns a tuple with the BuyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyQuantity

`func (o *TransportSubServiceDto) SetBuyQuantity(v float64)`

SetBuyQuantity sets BuyQuantity field to given value.

### HasBuyQuantity

`func (o *TransportSubServiceDto) HasBuyQuantity() bool`

HasBuyQuantity returns a boolean if a field has been set.

### GetBuyUnitPrice

`func (o *TransportSubServiceDto) GetBuyUnitPrice() float64`

GetBuyUnitPrice returns the BuyUnitPrice field if non-nil, zero value otherwise.

### GetBuyUnitPriceOk

`func (o *TransportSubServiceDto) GetBuyUnitPriceOk() (*float64, bool)`

GetBuyUnitPriceOk returns a tuple with the BuyUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUnitPrice

`func (o *TransportSubServiceDto) SetBuyUnitPrice(v float64)`

SetBuyUnitPrice sets BuyUnitPrice field to given value.

### HasBuyUnitPrice

`func (o *TransportSubServiceDto) HasBuyUnitPrice() bool`

HasBuyUnitPrice returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportSubServiceDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportSubServiceDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportSubServiceDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportSubServiceDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPriceInReferenceCurrency

`func (o *TransportSubServiceDto) GetForcedBuyPriceInReferenceCurrency() float64`

GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedBuyPriceInReferenceCurrencyOk

`func (o *TransportSubServiceDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPriceInReferenceCurrency

`func (o *TransportSubServiceDto) SetForcedBuyPriceInReferenceCurrency(v float64)`

SetForcedBuyPriceInReferenceCurrency sets ForcedBuyPriceInReferenceCurrency field to given value.

### HasForcedBuyPriceInReferenceCurrency

`func (o *TransportSubServiceDto) HasForcedBuyPriceInReferenceCurrency() bool`

HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetIsBuyGlobalPricingEnabled

`func (o *TransportSubServiceDto) GetIsBuyGlobalPricingEnabled() bool`

GetIsBuyGlobalPricingEnabled returns the IsBuyGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsBuyGlobalPricingEnabledOk

`func (o *TransportSubServiceDto) GetIsBuyGlobalPricingEnabledOk() (*bool, bool)`

GetIsBuyGlobalPricingEnabledOk returns a tuple with the IsBuyGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyGlobalPricingEnabled

`func (o *TransportSubServiceDto) SetIsBuyGlobalPricingEnabled(v bool)`

SetIsBuyGlobalPricingEnabled sets IsBuyGlobalPricingEnabled field to given value.

### HasIsBuyGlobalPricingEnabled

`func (o *TransportSubServiceDto) HasIsBuyGlobalPricingEnabled() bool`

HasIsBuyGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsBuyQuantityForced

`func (o *TransportSubServiceDto) GetIsBuyQuantityForced() bool`

GetIsBuyQuantityForced returns the IsBuyQuantityForced field if non-nil, zero value otherwise.

### GetIsBuyQuantityForcedOk

`func (o *TransportSubServiceDto) GetIsBuyQuantityForcedOk() (*bool, bool)`

GetIsBuyQuantityForcedOk returns a tuple with the IsBuyQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyQuantityForced

`func (o *TransportSubServiceDto) SetIsBuyQuantityForced(v bool)`

SetIsBuyQuantityForced sets IsBuyQuantityForced field to given value.

### HasIsBuyQuantityForced

`func (o *TransportSubServiceDto) HasIsBuyQuantityForced() bool`

HasIsBuyQuantityForced returns a boolean if a field has been set.

### GetBuyVatCode

`func (o *TransportSubServiceDto) GetBuyVatCode() string`

GetBuyVatCode returns the BuyVatCode field if non-nil, zero value otherwise.

### GetBuyVatCodeOk

`func (o *TransportSubServiceDto) GetBuyVatCodeOk() (*string, bool)`

GetBuyVatCodeOk returns a tuple with the BuyVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyVatCode

`func (o *TransportSubServiceDto) SetBuyVatCode(v string)`

SetBuyVatCode sets BuyVatCode field to given value.

### HasBuyVatCode

`func (o *TransportSubServiceDto) HasBuyVatCode() bool`

HasBuyVatCode returns a boolean if a field has been set.

### GetIsBuyVatCodeForced

`func (o *TransportSubServiceDto) GetIsBuyVatCodeForced() bool`

GetIsBuyVatCodeForced returns the IsBuyVatCodeForced field if non-nil, zero value otherwise.

### GetIsBuyVatCodeForcedOk

`func (o *TransportSubServiceDto) GetIsBuyVatCodeForcedOk() (*bool, bool)`

GetIsBuyVatCodeForcedOk returns a tuple with the IsBuyVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyVatCodeForced

`func (o *TransportSubServiceDto) SetIsBuyVatCodeForced(v bool)`

SetIsBuyVatCodeForced sets IsBuyVatCodeForced field to given value.

### HasIsBuyVatCodeForced

`func (o *TransportSubServiceDto) HasIsBuyVatCodeForced() bool`

HasIsBuyVatCodeForced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


