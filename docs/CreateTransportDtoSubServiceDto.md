# CreateTransportDtoSubServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Mandatory. Sub service&#39;s unique code. | 
**ForcedSellQuantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**ForcedSellUnitPrice** | Pointer to **float64** | If sell unit price is forced, sell global pricing will be disabled.    Forced sell unit price and forced sell price must not be both specified.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellPrice** | Pointer to **float64** | If sell price is forced, sell global pricing will be disabled.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellPriceInReferenceCurrency** | Pointer to **float64** | Forced sell price can be specified in reference currency when it is different from customer&#39;s currency    and when this option is available in general settings.    Forced sell price and forced sell price in reference currency must not be both set.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedSellVatCode** | Pointer to **string** | The following permission(s) are required to access this property:  Set VAT in orders. | [optional] 
**ForcedBuyQuantity** | Pointer to **float64** | This number can have up to 2 decimals. | [optional] 
**ForcedBuyUnitPrice** | Pointer to **float64** | If buy unit price is forced, buy global pricing will be disabled.  Forced buy unit price and forced buy price must not be both specified.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyPrice** | Pointer to **float64** | If buy price is forced, buy global pricing will be disabled.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyPriceInReferenceCurrency** | Pointer to **float64** | Forced buy price can be specified in reference currency when it is different from customer&#39;s currency    and when this option is available in general settings.    Forced buy price and forced buy price in reference currency must not be both set.    This number can have up to 2 decimals.    The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyVatCode** | Pointer to **string** | The following permission(s) are required to access this property:  Set VAT in orders. | [optional] 

## Methods

### NewCreateTransportDtoSubServiceDto

`func NewCreateTransportDtoSubServiceDto(code string, ) *CreateTransportDtoSubServiceDto`

NewCreateTransportDtoSubServiceDto instantiates a new CreateTransportDtoSubServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoSubServiceDtoWithDefaults

`func NewCreateTransportDtoSubServiceDtoWithDefaults() *CreateTransportDtoSubServiceDto`

NewCreateTransportDtoSubServiceDtoWithDefaults instantiates a new CreateTransportDtoSubServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTransportDtoSubServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTransportDtoSubServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTransportDtoSubServiceDto) SetCode(v string)`

SetCode sets Code field to given value.


### GetForcedSellQuantity

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellQuantity() float64`

GetForcedSellQuantity returns the ForcedSellQuantity field if non-nil, zero value otherwise.

### GetForcedSellQuantityOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellQuantityOk() (*float64, bool)`

GetForcedSellQuantityOk returns a tuple with the ForcedSellQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellQuantity

`func (o *CreateTransportDtoSubServiceDto) SetForcedSellQuantity(v float64)`

SetForcedSellQuantity sets ForcedSellQuantity field to given value.

### HasForcedSellQuantity

`func (o *CreateTransportDtoSubServiceDto) HasForcedSellQuantity() bool`

HasForcedSellQuantity returns a boolean if a field has been set.

### GetForcedSellUnitPrice

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellUnitPrice() float64`

GetForcedSellUnitPrice returns the ForcedSellUnitPrice field if non-nil, zero value otherwise.

### GetForcedSellUnitPriceOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellUnitPriceOk() (*float64, bool)`

GetForcedSellUnitPriceOk returns a tuple with the ForcedSellUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellUnitPrice

`func (o *CreateTransportDtoSubServiceDto) SetForcedSellUnitPrice(v float64)`

SetForcedSellUnitPrice sets ForcedSellUnitPrice field to given value.

### HasForcedSellUnitPrice

`func (o *CreateTransportDtoSubServiceDto) HasForcedSellUnitPrice() bool`

HasForcedSellUnitPrice returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *CreateTransportDtoSubServiceDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *CreateTransportDtoSubServiceDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetForcedSellPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellPriceInReferenceCurrency() float64`

GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedSellPriceInReferenceCurrencyOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) SetForcedSellPriceInReferenceCurrency(v float64)`

SetForcedSellPriceInReferenceCurrency sets ForcedSellPriceInReferenceCurrency field to given value.

### HasForcedSellPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) HasForcedSellPriceInReferenceCurrency() bool`

HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedSellVatCode

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellVatCode() string`

GetForcedSellVatCode returns the ForcedSellVatCode field if non-nil, zero value otherwise.

### GetForcedSellVatCodeOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedSellVatCodeOk() (*string, bool)`

GetForcedSellVatCodeOk returns a tuple with the ForcedSellVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellVatCode

`func (o *CreateTransportDtoSubServiceDto) SetForcedSellVatCode(v string)`

SetForcedSellVatCode sets ForcedSellVatCode field to given value.

### HasForcedSellVatCode

`func (o *CreateTransportDtoSubServiceDto) HasForcedSellVatCode() bool`

HasForcedSellVatCode returns a boolean if a field has been set.

### GetForcedBuyQuantity

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyQuantity() float64`

GetForcedBuyQuantity returns the ForcedBuyQuantity field if non-nil, zero value otherwise.

### GetForcedBuyQuantityOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyQuantityOk() (*float64, bool)`

GetForcedBuyQuantityOk returns a tuple with the ForcedBuyQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyQuantity

`func (o *CreateTransportDtoSubServiceDto) SetForcedBuyQuantity(v float64)`

SetForcedBuyQuantity sets ForcedBuyQuantity field to given value.

### HasForcedBuyQuantity

`func (o *CreateTransportDtoSubServiceDto) HasForcedBuyQuantity() bool`

HasForcedBuyQuantity returns a boolean if a field has been set.

### GetForcedBuyUnitPrice

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyUnitPrice() float64`

GetForcedBuyUnitPrice returns the ForcedBuyUnitPrice field if non-nil, zero value otherwise.

### GetForcedBuyUnitPriceOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyUnitPriceOk() (*float64, bool)`

GetForcedBuyUnitPriceOk returns a tuple with the ForcedBuyUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyUnitPrice

`func (o *CreateTransportDtoSubServiceDto) SetForcedBuyUnitPrice(v float64)`

SetForcedBuyUnitPrice sets ForcedBuyUnitPrice field to given value.

### HasForcedBuyUnitPrice

`func (o *CreateTransportDtoSubServiceDto) HasForcedBuyUnitPrice() bool`

HasForcedBuyUnitPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *CreateTransportDtoSubServiceDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *CreateTransportDtoSubServiceDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyPriceInReferenceCurrency() float64`

GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedBuyPriceInReferenceCurrencyOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) SetForcedBuyPriceInReferenceCurrency(v float64)`

SetForcedBuyPriceInReferenceCurrency sets ForcedBuyPriceInReferenceCurrency field to given value.

### HasForcedBuyPriceInReferenceCurrency

`func (o *CreateTransportDtoSubServiceDto) HasForcedBuyPriceInReferenceCurrency() bool`

HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedBuyVatCode

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyVatCode() string`

GetForcedBuyVatCode returns the ForcedBuyVatCode field if non-nil, zero value otherwise.

### GetForcedBuyVatCodeOk

`func (o *CreateTransportDtoSubServiceDto) GetForcedBuyVatCodeOk() (*string, bool)`

GetForcedBuyVatCodeOk returns a tuple with the ForcedBuyVatCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyVatCode

`func (o *CreateTransportDtoSubServiceDto) SetForcedBuyVatCode(v string)`

SetForcedBuyVatCode sets ForcedBuyVatCode field to given value.

### HasForcedBuyVatCode

`func (o *CreateTransportDtoSubServiceDto) HasForcedBuyVatCode() bool`

HasForcedBuyVatCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


