# TransportPricingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellCurrencyCode** | Pointer to **string** |  | [optional] 
**SellPrice** | Pointer to **float64** |  | [optional] 
**ForcedSellPrice** | Pointer to **float64** |  | [optional] 
**SellFuelSurchargePrice** | Pointer to **float64** |  | [optional] 
**BuyCurrencyCode** | Pointer to **string** |  | [optional] 
**BuyPrice** | Pointer to **float64** |  | [optional] 
**ForcedBuyPrice** | Pointer to **float64** |  | [optional] 
**BuyFuelSurchargePrice** | Pointer to **float64** |  | [optional] 
**ReferenceCurrencyCode** | Pointer to **string** |  | [optional] 
**SellPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s sell price converted to reference currency | [optional] 
**ForcedSellPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s forced sell price converted to reference currency | [optional] 
**SellFuelSurchargePriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s sell fuel surcharge price converted to reference currency | [optional] 
**BuyPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s buy price converted to reference currency | [optional] 
**ForcedBuyPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s forced buy price converted to reference currency | [optional] 
**BuyFuelSurchargePriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s buy fuel surcharge price converted to reference currency | [optional] 
**SubServices** | Pointer to [**CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto**](CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto.md) |  | [optional] 

## Methods

### NewTransportPricingDto

`func NewTransportPricingDto() *TransportPricingDto`

NewTransportPricingDto instantiates a new TransportPricingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPricingDtoWithDefaults

`func NewTransportPricingDtoWithDefaults() *TransportPricingDto`

NewTransportPricingDtoWithDefaults instantiates a new TransportPricingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellCurrencyCode

`func (o *TransportPricingDto) GetSellCurrencyCode() string`

GetSellCurrencyCode returns the SellCurrencyCode field if non-nil, zero value otherwise.

### GetSellCurrencyCodeOk

`func (o *TransportPricingDto) GetSellCurrencyCodeOk() (*string, bool)`

GetSellCurrencyCodeOk returns a tuple with the SellCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellCurrencyCode

`func (o *TransportPricingDto) SetSellCurrencyCode(v string)`

SetSellCurrencyCode sets SellCurrencyCode field to given value.

### HasSellCurrencyCode

`func (o *TransportPricingDto) HasSellCurrencyCode() bool`

HasSellCurrencyCode returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportPricingDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportPricingDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportPricingDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportPricingDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *TransportPricingDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *TransportPricingDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *TransportPricingDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *TransportPricingDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetSellFuelSurchargePrice

`func (o *TransportPricingDto) GetSellFuelSurchargePrice() float64`

GetSellFuelSurchargePrice returns the SellFuelSurchargePrice field if non-nil, zero value otherwise.

### GetSellFuelSurchargePriceOk

`func (o *TransportPricingDto) GetSellFuelSurchargePriceOk() (*float64, bool)`

GetSellFuelSurchargePriceOk returns a tuple with the SellFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFuelSurchargePrice

`func (o *TransportPricingDto) SetSellFuelSurchargePrice(v float64)`

SetSellFuelSurchargePrice sets SellFuelSurchargePrice field to given value.

### HasSellFuelSurchargePrice

`func (o *TransportPricingDto) HasSellFuelSurchargePrice() bool`

HasSellFuelSurchargePrice returns a boolean if a field has been set.

### GetBuyCurrencyCode

`func (o *TransportPricingDto) GetBuyCurrencyCode() string`

GetBuyCurrencyCode returns the BuyCurrencyCode field if non-nil, zero value otherwise.

### GetBuyCurrencyCodeOk

`func (o *TransportPricingDto) GetBuyCurrencyCodeOk() (*string, bool)`

GetBuyCurrencyCodeOk returns a tuple with the BuyCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyCurrencyCode

`func (o *TransportPricingDto) SetBuyCurrencyCode(v string)`

SetBuyCurrencyCode sets BuyCurrencyCode field to given value.

### HasBuyCurrencyCode

`func (o *TransportPricingDto) HasBuyCurrencyCode() bool`

HasBuyCurrencyCode returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportPricingDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportPricingDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportPricingDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportPricingDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *TransportPricingDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *TransportPricingDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *TransportPricingDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *TransportPricingDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetBuyFuelSurchargePrice

`func (o *TransportPricingDto) GetBuyFuelSurchargePrice() float64`

GetBuyFuelSurchargePrice returns the BuyFuelSurchargePrice field if non-nil, zero value otherwise.

### GetBuyFuelSurchargePriceOk

`func (o *TransportPricingDto) GetBuyFuelSurchargePriceOk() (*float64, bool)`

GetBuyFuelSurchargePriceOk returns a tuple with the BuyFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFuelSurchargePrice

`func (o *TransportPricingDto) SetBuyFuelSurchargePrice(v float64)`

SetBuyFuelSurchargePrice sets BuyFuelSurchargePrice field to given value.

### HasBuyFuelSurchargePrice

`func (o *TransportPricingDto) HasBuyFuelSurchargePrice() bool`

HasBuyFuelSurchargePrice returns a boolean if a field has been set.

### GetReferenceCurrencyCode

`func (o *TransportPricingDto) GetReferenceCurrencyCode() string`

GetReferenceCurrencyCode returns the ReferenceCurrencyCode field if non-nil, zero value otherwise.

### GetReferenceCurrencyCodeOk

`func (o *TransportPricingDto) GetReferenceCurrencyCodeOk() (*string, bool)`

GetReferenceCurrencyCodeOk returns a tuple with the ReferenceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCurrencyCode

`func (o *TransportPricingDto) SetReferenceCurrencyCode(v string)`

SetReferenceCurrencyCode sets ReferenceCurrencyCode field to given value.

### HasReferenceCurrencyCode

`func (o *TransportPricingDto) HasReferenceCurrencyCode() bool`

HasReferenceCurrencyCode returns a boolean if a field has been set.

### GetSellPriceInReferenceCurrency

`func (o *TransportPricingDto) GetSellPriceInReferenceCurrency() float64`

GetSellPriceInReferenceCurrency returns the SellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellPriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellPriceInReferenceCurrencyOk returns a tuple with the SellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPriceInReferenceCurrency

`func (o *TransportPricingDto) SetSellPriceInReferenceCurrency(v float64)`

SetSellPriceInReferenceCurrency sets SellPriceInReferenceCurrency field to given value.

### HasSellPriceInReferenceCurrency

`func (o *TransportPricingDto) HasSellPriceInReferenceCurrency() bool`

HasSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedSellPriceInReferenceCurrency

`func (o *TransportPricingDto) GetForcedSellPriceInReferenceCurrency() float64`

GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedSellPriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPriceInReferenceCurrency

`func (o *TransportPricingDto) SetForcedSellPriceInReferenceCurrency(v float64)`

SetForcedSellPriceInReferenceCurrency sets ForcedSellPriceInReferenceCurrency field to given value.

### HasForcedSellPriceInReferenceCurrency

`func (o *TransportPricingDto) HasForcedSellPriceInReferenceCurrency() bool`

HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) GetSellFuelSurchargePriceInReferenceCurrency() float64`

GetSellFuelSurchargePriceInReferenceCurrency returns the SellFuelSurchargePriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellFuelSurchargePriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetSellFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool)`

GetSellFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the SellFuelSurchargePriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) SetSellFuelSurchargePriceInReferenceCurrency(v float64)`

SetSellFuelSurchargePriceInReferenceCurrency sets SellFuelSurchargePriceInReferenceCurrency field to given value.

### HasSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) HasSellFuelSurchargePriceInReferenceCurrency() bool`

HasSellFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) GetBuyPriceInReferenceCurrency() float64`

GetBuyPriceInReferenceCurrency returns the BuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyPriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyPriceInReferenceCurrencyOk returns a tuple with the BuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) SetBuyPriceInReferenceCurrency(v float64)`

SetBuyPriceInReferenceCurrency sets BuyPriceInReferenceCurrency field to given value.

### HasBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) HasBuyPriceInReferenceCurrency() bool`

HasBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) GetForcedBuyPriceInReferenceCurrency() float64`

GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedBuyPriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) SetForcedBuyPriceInReferenceCurrency(v float64)`

SetForcedBuyPriceInReferenceCurrency sets ForcedBuyPriceInReferenceCurrency field to given value.

### HasForcedBuyPriceInReferenceCurrency

`func (o *TransportPricingDto) HasForcedBuyPriceInReferenceCurrency() bool`

HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) GetBuyFuelSurchargePriceInReferenceCurrency() float64`

GetBuyFuelSurchargePriceInReferenceCurrency returns the BuyFuelSurchargePriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyFuelSurchargePriceInReferenceCurrencyOk

`func (o *TransportPricingDto) GetBuyFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the BuyFuelSurchargePriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) SetBuyFuelSurchargePriceInReferenceCurrency(v float64)`

SetBuyFuelSurchargePriceInReferenceCurrency sets BuyFuelSurchargePriceInReferenceCurrency field to given value.

### HasBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportPricingDto) HasBuyFuelSurchargePriceInReferenceCurrency() bool`

HasBuyFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.

### GetSubServices

`func (o *TransportPricingDto) GetSubServices() CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *TransportPricingDto) GetSubServicesOk() (*CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *TransportPricingDto) SetSubServices(v CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *TransportPricingDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


