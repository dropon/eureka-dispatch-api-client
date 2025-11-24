# TransportPricingDtoTransportSubServicePricingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**UnitCode** | Pointer to **string** |  | [optional] 
**SellQuantity** | Pointer to **float64** |  | [optional] 
**SellPrice** | Pointer to **float64** |  | [optional] 
**SellUnitPrice** | Pointer to **float64** |  | [optional] 
**IsSellGlobalPricingEnabled** | Pointer to **bool** |  | [optional] 
**IsSellQuantityForced** | Pointer to **bool** |  | [optional] 
**SellVatCode** | Pointer to **string** |  | [optional] 
**IsSellVatCodeForced** | Pointer to **bool** |  | [optional] 
**BuyQuantity** | Pointer to **float64** |  | [optional] 
**BuyPrice** | Pointer to **float64** |  | [optional] 
**BuyUnitPrice** | Pointer to **float64** |  | [optional] 
**IsBuyGlobalPricingEnabled** | Pointer to **bool** |  | [optional] 
**IsBuyQuantityForced** | Pointer to **bool** |  | [optional] 
**BuyVatCode** | Pointer to **string** |  | [optional] 
**IsBuyVatCodeForced** | Pointer to **bool** |  | [optional] 
**SellPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s sell price forced in reference currency, otherwise sell price converted to reference currency | [optional] 
**SellUnitPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s sell unit price converted to reference currency | [optional] 
**BuyPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s buy price forced in reference currency, otherwise buy price converted to reference currency | [optional] 
**BuyUnitPriceInReferenceCurrency** | Pointer to **float64** | Sub service&#39;s buy unit price converted to reference currency | [optional] 

## Methods

### NewTransportPricingDtoTransportSubServicePricingDto

`func NewTransportPricingDtoTransportSubServicePricingDto() *TransportPricingDtoTransportSubServicePricingDto`

NewTransportPricingDtoTransportSubServicePricingDto instantiates a new TransportPricingDtoTransportSubServicePricingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPricingDtoTransportSubServicePricingDtoWithDefaults

`func NewTransportPricingDtoTransportSubServicePricingDtoWithDefaults() *TransportPricingDtoTransportSubServicePricingDto`

NewTransportPricingDtoTransportSubServicePricingDtoWithDefaults instantiates a new TransportPricingDtoTransportSubServicePricingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUnitCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetSellQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellQuantity() float64`

GetSellQuantity returns the SellQuantity field if non-nil, zero value otherwise.

### GetSellQuantityOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellQuantityOk() (*float64, bool)`

GetSellQuantityOk returns a tuple with the SellQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellQuantity(v float64)`

SetSellQuantity sets SellQuantity field to given value.

### HasSellQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellQuantity() bool`

HasSellQuantity returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetSellUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellUnitPrice() float64`

GetSellUnitPrice returns the SellUnitPrice field if non-nil, zero value otherwise.

### GetSellUnitPriceOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellUnitPriceOk() (*float64, bool)`

GetSellUnitPriceOk returns a tuple with the SellUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellUnitPrice(v float64)`

SetSellUnitPrice sets SellUnitPrice field to given value.

### HasSellUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellUnitPrice() bool`

HasSellUnitPrice returns a boolean if a field has been set.

### GetIsSellGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellGlobalPricingEnabled() bool`

GetIsSellGlobalPricingEnabled returns the IsSellGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsSellGlobalPricingEnabledOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellGlobalPricingEnabledOk() (*bool, bool)`

GetIsSellGlobalPricingEnabledOk returns a tuple with the IsSellGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsSellGlobalPricingEnabled(v bool)`

SetIsSellGlobalPricingEnabled sets IsSellGlobalPricingEnabled field to given value.

### HasIsSellGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsSellGlobalPricingEnabled() bool`

HasIsSellGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsSellQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellQuantityForced() bool`

GetIsSellQuantityForced returns the IsSellQuantityForced field if non-nil, zero value otherwise.

### GetIsSellQuantityForcedOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellQuantityForcedOk() (*bool, bool)`

GetIsSellQuantityForcedOk returns a tuple with the IsSellQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsSellQuantityForced(v bool)`

SetIsSellQuantityForced sets IsSellQuantityForced field to given value.

### HasIsSellQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsSellQuantityForced() bool`

HasIsSellQuantityForced returns a boolean if a field has been set.

### GetSellVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellVatCode() string`

GetSellVatCode returns the SellVatCode field if non-nil, zero value otherwise.

### GetSellVatCodeOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellVatCodeOk() (*string, bool)`

GetSellVatCodeOk returns a tuple with the SellVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellVatCode(v string)`

SetSellVatCode sets SellVatCode field to given value.

### HasSellVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellVatCode() bool`

HasSellVatCode returns a boolean if a field has been set.

### GetIsSellVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellVatCodeForced() bool`

GetIsSellVatCodeForced returns the IsSellVatCodeForced field if non-nil, zero value otherwise.

### GetIsSellVatCodeForcedOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsSellVatCodeForcedOk() (*bool, bool)`

GetIsSellVatCodeForcedOk returns a tuple with the IsSellVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsSellVatCodeForced(v bool)`

SetIsSellVatCodeForced sets IsSellVatCodeForced field to given value.

### HasIsSellVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsSellVatCodeForced() bool`

HasIsSellVatCodeForced returns a boolean if a field has been set.

### GetBuyQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyQuantity() float64`

GetBuyQuantity returns the BuyQuantity field if non-nil, zero value otherwise.

### GetBuyQuantityOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyQuantityOk() (*float64, bool)`

GetBuyQuantityOk returns a tuple with the BuyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyQuantity(v float64)`

SetBuyQuantity sets BuyQuantity field to given value.

### HasBuyQuantity

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyQuantity() bool`

HasBuyQuantity returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetBuyUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyUnitPrice() float64`

GetBuyUnitPrice returns the BuyUnitPrice field if non-nil, zero value otherwise.

### GetBuyUnitPriceOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyUnitPriceOk() (*float64, bool)`

GetBuyUnitPriceOk returns a tuple with the BuyUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyUnitPrice(v float64)`

SetBuyUnitPrice sets BuyUnitPrice field to given value.

### HasBuyUnitPrice

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyUnitPrice() bool`

HasBuyUnitPrice returns a boolean if a field has been set.

### GetIsBuyGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyGlobalPricingEnabled() bool`

GetIsBuyGlobalPricingEnabled returns the IsBuyGlobalPricingEnabled field if non-nil, zero value otherwise.

### GetIsBuyGlobalPricingEnabledOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyGlobalPricingEnabledOk() (*bool, bool)`

GetIsBuyGlobalPricingEnabledOk returns a tuple with the IsBuyGlobalPricingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsBuyGlobalPricingEnabled(v bool)`

SetIsBuyGlobalPricingEnabled sets IsBuyGlobalPricingEnabled field to given value.

### HasIsBuyGlobalPricingEnabled

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsBuyGlobalPricingEnabled() bool`

HasIsBuyGlobalPricingEnabled returns a boolean if a field has been set.

### GetIsBuyQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyQuantityForced() bool`

GetIsBuyQuantityForced returns the IsBuyQuantityForced field if non-nil, zero value otherwise.

### GetIsBuyQuantityForcedOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyQuantityForcedOk() (*bool, bool)`

GetIsBuyQuantityForcedOk returns a tuple with the IsBuyQuantityForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsBuyQuantityForced(v bool)`

SetIsBuyQuantityForced sets IsBuyQuantityForced field to given value.

### HasIsBuyQuantityForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsBuyQuantityForced() bool`

HasIsBuyQuantityForced returns a boolean if a field has been set.

### GetBuyVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyVatCode() string`

GetBuyVatCode returns the BuyVatCode field if non-nil, zero value otherwise.

### GetBuyVatCodeOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyVatCodeOk() (*string, bool)`

GetBuyVatCodeOk returns a tuple with the BuyVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyVatCode(v string)`

SetBuyVatCode sets BuyVatCode field to given value.

### HasBuyVatCode

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyVatCode() bool`

HasBuyVatCode returns a boolean if a field has been set.

### GetIsBuyVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyVatCodeForced() bool`

GetIsBuyVatCodeForced returns the IsBuyVatCodeForced field if non-nil, zero value otherwise.

### GetIsBuyVatCodeForcedOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetIsBuyVatCodeForcedOk() (*bool, bool)`

GetIsBuyVatCodeForcedOk returns a tuple with the IsBuyVatCodeForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetIsBuyVatCodeForced(v bool)`

SetIsBuyVatCodeForced sets IsBuyVatCodeForced field to given value.

### HasIsBuyVatCodeForced

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasIsBuyVatCodeForced() bool`

HasIsBuyVatCodeForced returns a boolean if a field has been set.

### GetSellPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellPriceInReferenceCurrency() float64`

GetSellPriceInReferenceCurrency returns the SellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellPriceInReferenceCurrencyOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellPriceInReferenceCurrencyOk returns a tuple with the SellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellPriceInReferenceCurrency(v float64)`

SetSellPriceInReferenceCurrency sets SellPriceInReferenceCurrency field to given value.

### HasSellPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellPriceInReferenceCurrency() bool`

HasSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetSellUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellUnitPriceInReferenceCurrency() float64`

GetSellUnitPriceInReferenceCurrency returns the SellUnitPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellUnitPriceInReferenceCurrencyOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetSellUnitPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellUnitPriceInReferenceCurrencyOk returns a tuple with the SellUnitPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetSellUnitPriceInReferenceCurrency(v float64)`

SetSellUnitPriceInReferenceCurrency sets SellUnitPriceInReferenceCurrency field to given value.

### HasSellUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasSellUnitPriceInReferenceCurrency() bool`

HasSellUnitPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyPriceInReferenceCurrency() float64`

GetBuyPriceInReferenceCurrency returns the BuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyPriceInReferenceCurrencyOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyPriceInReferenceCurrencyOk returns a tuple with the BuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyPriceInReferenceCurrency(v float64)`

SetBuyPriceInReferenceCurrency sets BuyPriceInReferenceCurrency field to given value.

### HasBuyPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyPriceInReferenceCurrency() bool`

HasBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyUnitPriceInReferenceCurrency() float64`

GetBuyUnitPriceInReferenceCurrency returns the BuyUnitPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyUnitPriceInReferenceCurrencyOk

`func (o *TransportPricingDtoTransportSubServicePricingDto) GetBuyUnitPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyUnitPriceInReferenceCurrencyOk returns a tuple with the BuyUnitPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) SetBuyUnitPriceInReferenceCurrency(v float64)`

SetBuyUnitPriceInReferenceCurrency sets BuyUnitPriceInReferenceCurrency field to given value.

### HasBuyUnitPriceInReferenceCurrency

`func (o *TransportPricingDtoTransportSubServicePricingDto) HasBuyUnitPriceInReferenceCurrency() bool`

HasBuyUnitPriceInReferenceCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


