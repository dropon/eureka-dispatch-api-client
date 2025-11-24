# TransportEntryDryRunDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerCode** | Pointer to **string** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**OrdererCode** | Pointer to **string** |  | [optional] 
**OrdererName** | Pointer to **string** |  | [optional] 
**Sign** | Pointer to **string** |  | [optional] 
**SecretCode** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**IsRoundTrip** | Pointer to **bool** |  | [optional] 
**CustomerCallDateTime** | Pointer to **time.Time** |  | [optional] 
**PickupStep** | Pointer to [**TransportEntryDryRunDtoStepDto**](TransportEntryDryRunDtoStepDto.md) |  | [optional] 
**DeliveryStep** | Pointer to [**TransportEntryDryRunDtoStepDto**](TransportEntryDryRunDtoStepDto.md) |  | [optional] 
**ServiceCode** | Pointer to **string** |  | [optional] 
**SubServices** | Pointer to [**[]TransportEntryDryRunDtoSubServiceDto**](TransportEntryDryRunDtoSubServiceDto.md) |  | [optional] 
**Packing** | Pointer to [**TransportEntryDryRunDtoPackingDto**](TransportEntryDryRunDtoPackingDto.md) |  | [optional] 
**CustomerCustomParameters** | Pointer to [**[]TransportEntryDryRunDtoCustomParameterDto**](TransportEntryDryRunDtoCustomParameterDto.md) |  | [optional] 
**ServiceCustomParameters** | Pointer to [**[]TransportEntryDryRunDtoCustomParameterDto**](TransportEntryDryRunDtoCustomParameterDto.md) |  | [optional] 
**Reference1** | Pointer to **string** |  | [optional] 
**Reference2** | Pointer to **string** |  | [optional] 
**Reference3** | Pointer to **string** |  | [optional] 
**DistanceKm** | Pointer to **float64** |  | [optional] 
**TotalDurationMinutes** | Pointer to **int32** |  | [optional] 
**IsStrategic** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**ResponsibleOperatorId** | Pointer to **int32** |  | [optional] 
**GasEmission** | Pointer to **float64** |  | [optional] 
**IsGasEmissionForced** | Pointer to **bool** |  | [optional] 
**SellCurrencyCode** | Pointer to **string** |  | [optional] 
**SellPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedSellPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**SellFuelSurchargePrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyCurrencyCode** | Pointer to **string** |  | [optional] 
**BuyPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedBuyPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyFuelSurchargePrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices. | [optional] 
**ReferenceCurrencyCode** | Pointer to **string** |  | [optional] 
**SellPricingPathId** | Pointer to **int32** |  | [optional] 
**SellPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s sell price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedSellPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s forced sell price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**SellFuelSurchargePriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s sell fuel surcharge price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s buy price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**ForcedBuyPriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s forced buy price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**BuyFuelSurchargePriceInReferenceCurrency** | Pointer to **float64** | Transport&#39;s buy fuel surcharge price converted to reference currency    The following permission(s) are required to access this property:  See prices. | [optional] 
**Comment** | Pointer to [**TransportEntryDryRunDtoCommentDto**](TransportEntryDryRunDtoCommentDto.md) |  | [optional] 
**TransportBillAddress** | Pointer to [**TransportEntryDryRunDtoTransportBillAddressDto**](TransportEntryDryRunDtoTransportBillAddressDto.md) |  | [optional] 
**CashOnDelivery** | Pointer to [**TransportEntryDryRunDtoCashOnDeliveryDto**](TransportEntryDryRunDtoCashOnDeliveryDto.md) |  | [optional] 
**Communication** | Pointer to [**TransportCommunicationConfigurationDto**](TransportCommunicationConfigurationDto.md) |  | [optional] 
**DangerousGoods** | Pointer to [**[]TransportEntryDryRunDtoDangerousGoodDto**](TransportEntryDryRunDtoDangerousGoodDto.md) |  | [optional] 

## Methods

### NewTransportEntryDryRunDto

`func NewTransportEntryDryRunDto() *TransportEntryDryRunDto`

NewTransportEntryDryRunDto instantiates a new TransportEntryDryRunDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEntryDryRunDtoWithDefaults

`func NewTransportEntryDryRunDtoWithDefaults() *TransportEntryDryRunDto`

NewTransportEntryDryRunDtoWithDefaults instantiates a new TransportEntryDryRunDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerCode

`func (o *TransportEntryDryRunDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *TransportEntryDryRunDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *TransportEntryDryRunDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *TransportEntryDryRunDto) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *TransportEntryDryRunDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *TransportEntryDryRunDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *TransportEntryDryRunDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *TransportEntryDryRunDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetOrdererCode

`func (o *TransportEntryDryRunDto) GetOrdererCode() string`

GetOrdererCode returns the OrdererCode field if non-nil, zero value otherwise.

### GetOrdererCodeOk

`func (o *TransportEntryDryRunDto) GetOrdererCodeOk() (*string, bool)`

GetOrdererCodeOk returns a tuple with the OrdererCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererCode

`func (o *TransportEntryDryRunDto) SetOrdererCode(v string)`

SetOrdererCode sets OrdererCode field to given value.

### HasOrdererCode

`func (o *TransportEntryDryRunDto) HasOrdererCode() bool`

HasOrdererCode returns a boolean if a field has been set.

### GetOrdererName

`func (o *TransportEntryDryRunDto) GetOrdererName() string`

GetOrdererName returns the OrdererName field if non-nil, zero value otherwise.

### GetOrdererNameOk

`func (o *TransportEntryDryRunDto) GetOrdererNameOk() (*string, bool)`

GetOrdererNameOk returns a tuple with the OrdererName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererName

`func (o *TransportEntryDryRunDto) SetOrdererName(v string)`

SetOrdererName sets OrdererName field to given value.

### HasOrdererName

`func (o *TransportEntryDryRunDto) HasOrdererName() bool`

HasOrdererName returns a boolean if a field has been set.

### GetSign

`func (o *TransportEntryDryRunDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *TransportEntryDryRunDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *TransportEntryDryRunDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *TransportEntryDryRunDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetSecretCode

`func (o *TransportEntryDryRunDto) GetSecretCode() string`

GetSecretCode returns the SecretCode field if non-nil, zero value otherwise.

### GetSecretCodeOk

`func (o *TransportEntryDryRunDto) GetSecretCodeOk() (*string, bool)`

GetSecretCodeOk returns a tuple with the SecretCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCode

`func (o *TransportEntryDryRunDto) SetSecretCode(v string)`

SetSecretCode sets SecretCode field to given value.

### HasSecretCode

`func (o *TransportEntryDryRunDto) HasSecretCode() bool`

HasSecretCode returns a boolean if a field has been set.

### GetNotes

`func (o *TransportEntryDryRunDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransportEntryDryRunDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransportEntryDryRunDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransportEntryDryRunDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsRoundTrip

`func (o *TransportEntryDryRunDto) GetIsRoundTrip() bool`

GetIsRoundTrip returns the IsRoundTrip field if non-nil, zero value otherwise.

### GetIsRoundTripOk

`func (o *TransportEntryDryRunDto) GetIsRoundTripOk() (*bool, bool)`

GetIsRoundTripOk returns a tuple with the IsRoundTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundTrip

`func (o *TransportEntryDryRunDto) SetIsRoundTrip(v bool)`

SetIsRoundTrip sets IsRoundTrip field to given value.

### HasIsRoundTrip

`func (o *TransportEntryDryRunDto) HasIsRoundTrip() bool`

HasIsRoundTrip returns a boolean if a field has been set.

### GetCustomerCallDateTime

`func (o *TransportEntryDryRunDto) GetCustomerCallDateTime() time.Time`

GetCustomerCallDateTime returns the CustomerCallDateTime field if non-nil, zero value otherwise.

### GetCustomerCallDateTimeOk

`func (o *TransportEntryDryRunDto) GetCustomerCallDateTimeOk() (*time.Time, bool)`

GetCustomerCallDateTimeOk returns a tuple with the CustomerCallDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCallDateTime

`func (o *TransportEntryDryRunDto) SetCustomerCallDateTime(v time.Time)`

SetCustomerCallDateTime sets CustomerCallDateTime field to given value.

### HasCustomerCallDateTime

`func (o *TransportEntryDryRunDto) HasCustomerCallDateTime() bool`

HasCustomerCallDateTime returns a boolean if a field has been set.

### GetPickupStep

`func (o *TransportEntryDryRunDto) GetPickupStep() TransportEntryDryRunDtoStepDto`

GetPickupStep returns the PickupStep field if non-nil, zero value otherwise.

### GetPickupStepOk

`func (o *TransportEntryDryRunDto) GetPickupStepOk() (*TransportEntryDryRunDtoStepDto, bool)`

GetPickupStepOk returns a tuple with the PickupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStep

`func (o *TransportEntryDryRunDto) SetPickupStep(v TransportEntryDryRunDtoStepDto)`

SetPickupStep sets PickupStep field to given value.

### HasPickupStep

`func (o *TransportEntryDryRunDto) HasPickupStep() bool`

HasPickupStep returns a boolean if a field has been set.

### GetDeliveryStep

`func (o *TransportEntryDryRunDto) GetDeliveryStep() TransportEntryDryRunDtoStepDto`

GetDeliveryStep returns the DeliveryStep field if non-nil, zero value otherwise.

### GetDeliveryStepOk

`func (o *TransportEntryDryRunDto) GetDeliveryStepOk() (*TransportEntryDryRunDtoStepDto, bool)`

GetDeliveryStepOk returns a tuple with the DeliveryStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStep

`func (o *TransportEntryDryRunDto) SetDeliveryStep(v TransportEntryDryRunDtoStepDto)`

SetDeliveryStep sets DeliveryStep field to given value.

### HasDeliveryStep

`func (o *TransportEntryDryRunDto) HasDeliveryStep() bool`

HasDeliveryStep returns a boolean if a field has been set.

### GetServiceCode

`func (o *TransportEntryDryRunDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *TransportEntryDryRunDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *TransportEntryDryRunDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *TransportEntryDryRunDto) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetSubServices

`func (o *TransportEntryDryRunDto) GetSubServices() []TransportEntryDryRunDtoSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *TransportEntryDryRunDto) GetSubServicesOk() (*[]TransportEntryDryRunDtoSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *TransportEntryDryRunDto) SetSubServices(v []TransportEntryDryRunDtoSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *TransportEntryDryRunDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetPacking

`func (o *TransportEntryDryRunDto) GetPacking() TransportEntryDryRunDtoPackingDto`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *TransportEntryDryRunDto) GetPackingOk() (*TransportEntryDryRunDtoPackingDto, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *TransportEntryDryRunDto) SetPacking(v TransportEntryDryRunDtoPackingDto)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *TransportEntryDryRunDto) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetCustomerCustomParameters

`func (o *TransportEntryDryRunDto) GetCustomerCustomParameters() []TransportEntryDryRunDtoCustomParameterDto`

GetCustomerCustomParameters returns the CustomerCustomParameters field if non-nil, zero value otherwise.

### GetCustomerCustomParametersOk

`func (o *TransportEntryDryRunDto) GetCustomerCustomParametersOk() (*[]TransportEntryDryRunDtoCustomParameterDto, bool)`

GetCustomerCustomParametersOk returns a tuple with the CustomerCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomParameters

`func (o *TransportEntryDryRunDto) SetCustomerCustomParameters(v []TransportEntryDryRunDtoCustomParameterDto)`

SetCustomerCustomParameters sets CustomerCustomParameters field to given value.

### HasCustomerCustomParameters

`func (o *TransportEntryDryRunDto) HasCustomerCustomParameters() bool`

HasCustomerCustomParameters returns a boolean if a field has been set.

### GetServiceCustomParameters

`func (o *TransportEntryDryRunDto) GetServiceCustomParameters() []TransportEntryDryRunDtoCustomParameterDto`

GetServiceCustomParameters returns the ServiceCustomParameters field if non-nil, zero value otherwise.

### GetServiceCustomParametersOk

`func (o *TransportEntryDryRunDto) GetServiceCustomParametersOk() (*[]TransportEntryDryRunDtoCustomParameterDto, bool)`

GetServiceCustomParametersOk returns a tuple with the ServiceCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCustomParameters

`func (o *TransportEntryDryRunDto) SetServiceCustomParameters(v []TransportEntryDryRunDtoCustomParameterDto)`

SetServiceCustomParameters sets ServiceCustomParameters field to given value.

### HasServiceCustomParameters

`func (o *TransportEntryDryRunDto) HasServiceCustomParameters() bool`

HasServiceCustomParameters returns a boolean if a field has been set.

### GetReference1

`func (o *TransportEntryDryRunDto) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *TransportEntryDryRunDto) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *TransportEntryDryRunDto) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *TransportEntryDryRunDto) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *TransportEntryDryRunDto) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *TransportEntryDryRunDto) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *TransportEntryDryRunDto) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *TransportEntryDryRunDto) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *TransportEntryDryRunDto) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *TransportEntryDryRunDto) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *TransportEntryDryRunDto) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *TransportEntryDryRunDto) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetDistanceKm

`func (o *TransportEntryDryRunDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *TransportEntryDryRunDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *TransportEntryDryRunDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *TransportEntryDryRunDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetTotalDurationMinutes

`func (o *TransportEntryDryRunDto) GetTotalDurationMinutes() int32`

GetTotalDurationMinutes returns the TotalDurationMinutes field if non-nil, zero value otherwise.

### GetTotalDurationMinutesOk

`func (o *TransportEntryDryRunDto) GetTotalDurationMinutesOk() (*int32, bool)`

GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMinutes

`func (o *TransportEntryDryRunDto) SetTotalDurationMinutes(v int32)`

SetTotalDurationMinutes sets TotalDurationMinutes field to given value.

### HasTotalDurationMinutes

`func (o *TransportEntryDryRunDto) HasTotalDurationMinutes() bool`

HasTotalDurationMinutes returns a boolean if a field has been set.

### GetIsStrategic

`func (o *TransportEntryDryRunDto) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *TransportEntryDryRunDto) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *TransportEntryDryRunDto) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *TransportEntryDryRunDto) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetIsShared

`func (o *TransportEntryDryRunDto) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *TransportEntryDryRunDto) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *TransportEntryDryRunDto) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *TransportEntryDryRunDto) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetResponsibleOperatorId

`func (o *TransportEntryDryRunDto) GetResponsibleOperatorId() int32`

GetResponsibleOperatorId returns the ResponsibleOperatorId field if non-nil, zero value otherwise.

### GetResponsibleOperatorIdOk

`func (o *TransportEntryDryRunDto) GetResponsibleOperatorIdOk() (*int32, bool)`

GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleOperatorId

`func (o *TransportEntryDryRunDto) SetResponsibleOperatorId(v int32)`

SetResponsibleOperatorId sets ResponsibleOperatorId field to given value.

### HasResponsibleOperatorId

`func (o *TransportEntryDryRunDto) HasResponsibleOperatorId() bool`

HasResponsibleOperatorId returns a boolean if a field has been set.

### GetGasEmission

`func (o *TransportEntryDryRunDto) GetGasEmission() float64`

GetGasEmission returns the GasEmission field if non-nil, zero value otherwise.

### GetGasEmissionOk

`func (o *TransportEntryDryRunDto) GetGasEmissionOk() (*float64, bool)`

GetGasEmissionOk returns a tuple with the GasEmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEmission

`func (o *TransportEntryDryRunDto) SetGasEmission(v float64)`

SetGasEmission sets GasEmission field to given value.

### HasGasEmission

`func (o *TransportEntryDryRunDto) HasGasEmission() bool`

HasGasEmission returns a boolean if a field has been set.

### GetIsGasEmissionForced

`func (o *TransportEntryDryRunDto) GetIsGasEmissionForced() bool`

GetIsGasEmissionForced returns the IsGasEmissionForced field if non-nil, zero value otherwise.

### GetIsGasEmissionForcedOk

`func (o *TransportEntryDryRunDto) GetIsGasEmissionForcedOk() (*bool, bool)`

GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGasEmissionForced

`func (o *TransportEntryDryRunDto) SetIsGasEmissionForced(v bool)`

SetIsGasEmissionForced sets IsGasEmissionForced field to given value.

### HasIsGasEmissionForced

`func (o *TransportEntryDryRunDto) HasIsGasEmissionForced() bool`

HasIsGasEmissionForced returns a boolean if a field has been set.

### GetSellCurrencyCode

`func (o *TransportEntryDryRunDto) GetSellCurrencyCode() string`

GetSellCurrencyCode returns the SellCurrencyCode field if non-nil, zero value otherwise.

### GetSellCurrencyCodeOk

`func (o *TransportEntryDryRunDto) GetSellCurrencyCodeOk() (*string, bool)`

GetSellCurrencyCodeOk returns a tuple with the SellCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellCurrencyCode

`func (o *TransportEntryDryRunDto) SetSellCurrencyCode(v string)`

SetSellCurrencyCode sets SellCurrencyCode field to given value.

### HasSellCurrencyCode

`func (o *TransportEntryDryRunDto) HasSellCurrencyCode() bool`

HasSellCurrencyCode returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportEntryDryRunDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportEntryDryRunDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportEntryDryRunDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportEntryDryRunDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *TransportEntryDryRunDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *TransportEntryDryRunDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *TransportEntryDryRunDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *TransportEntryDryRunDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetSellFuelSurchargePrice

`func (o *TransportEntryDryRunDto) GetSellFuelSurchargePrice() float64`

GetSellFuelSurchargePrice returns the SellFuelSurchargePrice field if non-nil, zero value otherwise.

### GetSellFuelSurchargePriceOk

`func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceOk() (*float64, bool)`

GetSellFuelSurchargePriceOk returns a tuple with the SellFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFuelSurchargePrice

`func (o *TransportEntryDryRunDto) SetSellFuelSurchargePrice(v float64)`

SetSellFuelSurchargePrice sets SellFuelSurchargePrice field to given value.

### HasSellFuelSurchargePrice

`func (o *TransportEntryDryRunDto) HasSellFuelSurchargePrice() bool`

HasSellFuelSurchargePrice returns a boolean if a field has been set.

### GetBuyCurrencyCode

`func (o *TransportEntryDryRunDto) GetBuyCurrencyCode() string`

GetBuyCurrencyCode returns the BuyCurrencyCode field if non-nil, zero value otherwise.

### GetBuyCurrencyCodeOk

`func (o *TransportEntryDryRunDto) GetBuyCurrencyCodeOk() (*string, bool)`

GetBuyCurrencyCodeOk returns a tuple with the BuyCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyCurrencyCode

`func (o *TransportEntryDryRunDto) SetBuyCurrencyCode(v string)`

SetBuyCurrencyCode sets BuyCurrencyCode field to given value.

### HasBuyCurrencyCode

`func (o *TransportEntryDryRunDto) HasBuyCurrencyCode() bool`

HasBuyCurrencyCode returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportEntryDryRunDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportEntryDryRunDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportEntryDryRunDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportEntryDryRunDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *TransportEntryDryRunDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *TransportEntryDryRunDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *TransportEntryDryRunDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *TransportEntryDryRunDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetBuyFuelSurchargePrice

`func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePrice() float64`

GetBuyFuelSurchargePrice returns the BuyFuelSurchargePrice field if non-nil, zero value otherwise.

### GetBuyFuelSurchargePriceOk

`func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceOk() (*float64, bool)`

GetBuyFuelSurchargePriceOk returns a tuple with the BuyFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFuelSurchargePrice

`func (o *TransportEntryDryRunDto) SetBuyFuelSurchargePrice(v float64)`

SetBuyFuelSurchargePrice sets BuyFuelSurchargePrice field to given value.

### HasBuyFuelSurchargePrice

`func (o *TransportEntryDryRunDto) HasBuyFuelSurchargePrice() bool`

HasBuyFuelSurchargePrice returns a boolean if a field has been set.

### GetReferenceCurrencyCode

`func (o *TransportEntryDryRunDto) GetReferenceCurrencyCode() string`

GetReferenceCurrencyCode returns the ReferenceCurrencyCode field if non-nil, zero value otherwise.

### GetReferenceCurrencyCodeOk

`func (o *TransportEntryDryRunDto) GetReferenceCurrencyCodeOk() (*string, bool)`

GetReferenceCurrencyCodeOk returns a tuple with the ReferenceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCurrencyCode

`func (o *TransportEntryDryRunDto) SetReferenceCurrencyCode(v string)`

SetReferenceCurrencyCode sets ReferenceCurrencyCode field to given value.

### HasReferenceCurrencyCode

`func (o *TransportEntryDryRunDto) HasReferenceCurrencyCode() bool`

HasReferenceCurrencyCode returns a boolean if a field has been set.

### GetSellPricingPathId

`func (o *TransportEntryDryRunDto) GetSellPricingPathId() int32`

GetSellPricingPathId returns the SellPricingPathId field if non-nil, zero value otherwise.

### GetSellPricingPathIdOk

`func (o *TransportEntryDryRunDto) GetSellPricingPathIdOk() (*int32, bool)`

GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPricingPathId

`func (o *TransportEntryDryRunDto) SetSellPricingPathId(v int32)`

SetSellPricingPathId sets SellPricingPathId field to given value.

### HasSellPricingPathId

`func (o *TransportEntryDryRunDto) HasSellPricingPathId() bool`

HasSellPricingPathId returns a boolean if a field has been set.

### GetSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetSellPriceInReferenceCurrency() float64`

GetSellPriceInReferenceCurrency returns the SellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetSellPriceInReferenceCurrencyOk returns a tuple with the SellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetSellPriceInReferenceCurrency(v float64)`

SetSellPriceInReferenceCurrency sets SellPriceInReferenceCurrency field to given value.

### HasSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasSellPriceInReferenceCurrency() bool`

HasSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetForcedSellPriceInReferenceCurrency() float64`

GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedSellPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetForcedSellPriceInReferenceCurrency(v float64)`

SetForcedSellPriceInReferenceCurrency sets ForcedSellPriceInReferenceCurrency field to given value.

### HasForcedSellPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasForcedSellPriceInReferenceCurrency() bool`

HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.

### GetSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceInReferenceCurrency() float64`

GetSellFuelSurchargePriceInReferenceCurrency returns the SellFuelSurchargePriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetSellFuelSurchargePriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool)`

GetSellFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the SellFuelSurchargePriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetSellFuelSurchargePriceInReferenceCurrency(v float64)`

SetSellFuelSurchargePriceInReferenceCurrency sets SellFuelSurchargePriceInReferenceCurrency field to given value.

### HasSellFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasSellFuelSurchargePriceInReferenceCurrency() bool`

HasSellFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetBuyPriceInReferenceCurrency() float64`

GetBuyPriceInReferenceCurrency returns the BuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyPriceInReferenceCurrencyOk returns a tuple with the BuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetBuyPriceInReferenceCurrency(v float64)`

SetBuyPriceInReferenceCurrency sets BuyPriceInReferenceCurrency field to given value.

### HasBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasBuyPriceInReferenceCurrency() bool`

HasBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetForcedBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetForcedBuyPriceInReferenceCurrency() float64`

GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetForcedBuyPriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool)`

GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetForcedBuyPriceInReferenceCurrency(v float64)`

SetForcedBuyPriceInReferenceCurrency sets ForcedBuyPriceInReferenceCurrency field to given value.

### HasForcedBuyPriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasForcedBuyPriceInReferenceCurrency() bool`

HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.

### GetBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceInReferenceCurrency() float64`

GetBuyFuelSurchargePriceInReferenceCurrency returns the BuyFuelSurchargePriceInReferenceCurrency field if non-nil, zero value otherwise.

### GetBuyFuelSurchargePriceInReferenceCurrencyOk

`func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool)`

GetBuyFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the BuyFuelSurchargePriceInReferenceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) SetBuyFuelSurchargePriceInReferenceCurrency(v float64)`

SetBuyFuelSurchargePriceInReferenceCurrency sets BuyFuelSurchargePriceInReferenceCurrency field to given value.

### HasBuyFuelSurchargePriceInReferenceCurrency

`func (o *TransportEntryDryRunDto) HasBuyFuelSurchargePriceInReferenceCurrency() bool`

HasBuyFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.

### GetComment

`func (o *TransportEntryDryRunDto) GetComment() TransportEntryDryRunDtoCommentDto`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransportEntryDryRunDto) GetCommentOk() (*TransportEntryDryRunDtoCommentDto, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransportEntryDryRunDto) SetComment(v TransportEntryDryRunDtoCommentDto)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransportEntryDryRunDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTransportBillAddress

`func (o *TransportEntryDryRunDto) GetTransportBillAddress() TransportEntryDryRunDtoTransportBillAddressDto`

GetTransportBillAddress returns the TransportBillAddress field if non-nil, zero value otherwise.

### GetTransportBillAddressOk

`func (o *TransportEntryDryRunDto) GetTransportBillAddressOk() (*TransportEntryDryRunDtoTransportBillAddressDto, bool)`

GetTransportBillAddressOk returns a tuple with the TransportBillAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportBillAddress

`func (o *TransportEntryDryRunDto) SetTransportBillAddress(v TransportEntryDryRunDtoTransportBillAddressDto)`

SetTransportBillAddress sets TransportBillAddress field to given value.

### HasTransportBillAddress

`func (o *TransportEntryDryRunDto) HasTransportBillAddress() bool`

HasTransportBillAddress returns a boolean if a field has been set.

### GetCashOnDelivery

`func (o *TransportEntryDryRunDto) GetCashOnDelivery() TransportEntryDryRunDtoCashOnDeliveryDto`

GetCashOnDelivery returns the CashOnDelivery field if non-nil, zero value otherwise.

### GetCashOnDeliveryOk

`func (o *TransportEntryDryRunDto) GetCashOnDeliveryOk() (*TransportEntryDryRunDtoCashOnDeliveryDto, bool)`

GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDelivery

`func (o *TransportEntryDryRunDto) SetCashOnDelivery(v TransportEntryDryRunDtoCashOnDeliveryDto)`

SetCashOnDelivery sets CashOnDelivery field to given value.

### HasCashOnDelivery

`func (o *TransportEntryDryRunDto) HasCashOnDelivery() bool`

HasCashOnDelivery returns a boolean if a field has been set.

### GetCommunication

`func (o *TransportEntryDryRunDto) GetCommunication() TransportCommunicationConfigurationDto`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *TransportEntryDryRunDto) GetCommunicationOk() (*TransportCommunicationConfigurationDto, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *TransportEntryDryRunDto) SetCommunication(v TransportCommunicationConfigurationDto)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *TransportEntryDryRunDto) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.

### GetDangerousGoods

`func (o *TransportEntryDryRunDto) GetDangerousGoods() []TransportEntryDryRunDtoDangerousGoodDto`

GetDangerousGoods returns the DangerousGoods field if non-nil, zero value otherwise.

### GetDangerousGoodsOk

`func (o *TransportEntryDryRunDto) GetDangerousGoodsOk() (*[]TransportEntryDryRunDtoDangerousGoodDto, bool)`

GetDangerousGoodsOk returns a tuple with the DangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoods

`func (o *TransportEntryDryRunDto) SetDangerousGoods(v []TransportEntryDryRunDtoDangerousGoodDto)`

SetDangerousGoods sets DangerousGoods field to given value.

### HasDangerousGoods

`func (o *TransportEntryDryRunDto) HasDangerousGoods() bool`

HasDangerousGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


