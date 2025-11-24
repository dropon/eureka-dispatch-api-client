# MergeableTransportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | This property is an entity identifier. | [optional] 
**MergeAction** | Pointer to **string** |  | [optional] 
**CustomerCode** | Pointer to **string** |  | [optional] 
**AgencyCode** | Pointer to **string** |  | [optional] 
**OrdererCode** | Pointer to **string** |  | [optional] 
**OrdererName** | Pointer to **string** |  | [optional] 
**Sign** | Pointer to **string** |  | [optional] 
**SecretCode** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**IsRoundTrip** | Pointer to **bool** |  | [optional] 
**PickupStep** | Pointer to [**UpdateTransportBaseDtoStepDto**](UpdateTransportBaseDtoStepDto.md) |  | [optional] 
**DeliveryStep** | Pointer to [**UpdateTransportBaseDtoStepDto**](UpdateTransportBaseDtoStepDto.md) |  | [optional] 
**ServiceCode** | Pointer to **string** |  | [optional] 
**SubServices** | Pointer to [**[]UpdateTransportBaseDtoSubServiceDto**](UpdateTransportBaseDtoSubServiceDto.md) | When the service is updated, this collection must be filled like for the creation :  Keep the collection empty to use service&#39;s default sub services,  or provide the full collection of sub services with merge action &#39;add&#39;. | [optional] 
**Packing** | Pointer to [**UpdateTransportBaseDtoPackingDto**](UpdateTransportBaseDtoPackingDto.md) |  | [optional] 
**CustomerCustomParameters** | Pointer to [**[]UpdateTransportBaseDtoCustomParameterDto**](UpdateTransportBaseDtoCustomParameterDto.md) |  | [optional] 
**ServiceCustomParameters** | Pointer to [**[]UpdateTransportBaseDtoCustomParameterDto**](UpdateTransportBaseDtoCustomParameterDto.md) |  | [optional] 
**DangerousGoods** | Pointer to [**[]UpdateTransportBaseDtoDangerousGoodDto**](UpdateTransportBaseDtoDangerousGoodDto.md) |  | [optional] 
**Reference1** | Pointer to **string** |  | [optional] 
**Reference2** | Pointer to **string** |  | [optional] 
**Reference3** | Pointer to **string** |  | [optional] 
**DistanceKm** | Pointer to **float64** |  | [optional] 
**TotalDurationMinutes** | Pointer to **int32** |  | [optional] 
**IsStrategic** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**ResponsibleOperatorId** | Pointer to **int32** |  | [optional] 
**GasEmission** | Pointer to **float64** | Gas emission is updated only if is gas emission forced is set | [optional] 
**IsGasEmissionForced** | Pointer to **bool** |  | [optional] 
**ForcedSellPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**ForcedBuyPrice** | Pointer to **float64** | The following permission(s) are required to access this property:  See prices,Set prices in orders. | [optional] 
**SellPricingPathId** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to [**UpdateTransportBaseDtoCommentDto**](UpdateTransportBaseDtoCommentDto.md) |  | [optional] 
**TransportBillAddress** | Pointer to [**UpdateTransportBaseDtoTransportBillAddressDto**](UpdateTransportBaseDtoTransportBillAddressDto.md) |  | [optional] 
**CashOnDelivery** | Pointer to [**UpdateTransportBaseDtoCashOnDeliveryDto**](UpdateTransportBaseDtoCashOnDeliveryDto.md) |  | [optional] 
**Communication** | Pointer to [**TransportCommunicationConfigurationDto**](TransportCommunicationConfigurationDto.md) |  | [optional] 

## Methods

### NewMergeableTransportDto

`func NewMergeableTransportDto() *MergeableTransportDto`

NewMergeableTransportDto instantiates a new MergeableTransportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeableTransportDtoWithDefaults

`func NewMergeableTransportDtoWithDefaults() *MergeableTransportDto`

NewMergeableTransportDtoWithDefaults instantiates a new MergeableTransportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *MergeableTransportDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *MergeableTransportDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *MergeableTransportDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *MergeableTransportDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetMergeAction

`func (o *MergeableTransportDto) GetMergeAction() string`

GetMergeAction returns the MergeAction field if non-nil, zero value otherwise.

### GetMergeActionOk

`func (o *MergeableTransportDto) GetMergeActionOk() (*string, bool)`

GetMergeActionOk returns a tuple with the MergeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeAction

`func (o *MergeableTransportDto) SetMergeAction(v string)`

SetMergeAction sets MergeAction field to given value.

### HasMergeAction

`func (o *MergeableTransportDto) HasMergeAction() bool`

HasMergeAction returns a boolean if a field has been set.

### GetCustomerCode

`func (o *MergeableTransportDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *MergeableTransportDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *MergeableTransportDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *MergeableTransportDto) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *MergeableTransportDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *MergeableTransportDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *MergeableTransportDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *MergeableTransportDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetOrdererCode

`func (o *MergeableTransportDto) GetOrdererCode() string`

GetOrdererCode returns the OrdererCode field if non-nil, zero value otherwise.

### GetOrdererCodeOk

`func (o *MergeableTransportDto) GetOrdererCodeOk() (*string, bool)`

GetOrdererCodeOk returns a tuple with the OrdererCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererCode

`func (o *MergeableTransportDto) SetOrdererCode(v string)`

SetOrdererCode sets OrdererCode field to given value.

### HasOrdererCode

`func (o *MergeableTransportDto) HasOrdererCode() bool`

HasOrdererCode returns a boolean if a field has been set.

### GetOrdererName

`func (o *MergeableTransportDto) GetOrdererName() string`

GetOrdererName returns the OrdererName field if non-nil, zero value otherwise.

### GetOrdererNameOk

`func (o *MergeableTransportDto) GetOrdererNameOk() (*string, bool)`

GetOrdererNameOk returns a tuple with the OrdererName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererName

`func (o *MergeableTransportDto) SetOrdererName(v string)`

SetOrdererName sets OrdererName field to given value.

### HasOrdererName

`func (o *MergeableTransportDto) HasOrdererName() bool`

HasOrdererName returns a boolean if a field has been set.

### GetSign

`func (o *MergeableTransportDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *MergeableTransportDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *MergeableTransportDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *MergeableTransportDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetSecretCode

`func (o *MergeableTransportDto) GetSecretCode() string`

GetSecretCode returns the SecretCode field if non-nil, zero value otherwise.

### GetSecretCodeOk

`func (o *MergeableTransportDto) GetSecretCodeOk() (*string, bool)`

GetSecretCodeOk returns a tuple with the SecretCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCode

`func (o *MergeableTransportDto) SetSecretCode(v string)`

SetSecretCode sets SecretCode field to given value.

### HasSecretCode

`func (o *MergeableTransportDto) HasSecretCode() bool`

HasSecretCode returns a boolean if a field has been set.

### GetNotes

`func (o *MergeableTransportDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MergeableTransportDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MergeableTransportDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MergeableTransportDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsRoundTrip

`func (o *MergeableTransportDto) GetIsRoundTrip() bool`

GetIsRoundTrip returns the IsRoundTrip field if non-nil, zero value otherwise.

### GetIsRoundTripOk

`func (o *MergeableTransportDto) GetIsRoundTripOk() (*bool, bool)`

GetIsRoundTripOk returns a tuple with the IsRoundTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundTrip

`func (o *MergeableTransportDto) SetIsRoundTrip(v bool)`

SetIsRoundTrip sets IsRoundTrip field to given value.

### HasIsRoundTrip

`func (o *MergeableTransportDto) HasIsRoundTrip() bool`

HasIsRoundTrip returns a boolean if a field has been set.

### GetPickupStep

`func (o *MergeableTransportDto) GetPickupStep() UpdateTransportBaseDtoStepDto`

GetPickupStep returns the PickupStep field if non-nil, zero value otherwise.

### GetPickupStepOk

`func (o *MergeableTransportDto) GetPickupStepOk() (*UpdateTransportBaseDtoStepDto, bool)`

GetPickupStepOk returns a tuple with the PickupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStep

`func (o *MergeableTransportDto) SetPickupStep(v UpdateTransportBaseDtoStepDto)`

SetPickupStep sets PickupStep field to given value.

### HasPickupStep

`func (o *MergeableTransportDto) HasPickupStep() bool`

HasPickupStep returns a boolean if a field has been set.

### GetDeliveryStep

`func (o *MergeableTransportDto) GetDeliveryStep() UpdateTransportBaseDtoStepDto`

GetDeliveryStep returns the DeliveryStep field if non-nil, zero value otherwise.

### GetDeliveryStepOk

`func (o *MergeableTransportDto) GetDeliveryStepOk() (*UpdateTransportBaseDtoStepDto, bool)`

GetDeliveryStepOk returns a tuple with the DeliveryStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStep

`func (o *MergeableTransportDto) SetDeliveryStep(v UpdateTransportBaseDtoStepDto)`

SetDeliveryStep sets DeliveryStep field to given value.

### HasDeliveryStep

`func (o *MergeableTransportDto) HasDeliveryStep() bool`

HasDeliveryStep returns a boolean if a field has been set.

### GetServiceCode

`func (o *MergeableTransportDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *MergeableTransportDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *MergeableTransportDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *MergeableTransportDto) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetSubServices

`func (o *MergeableTransportDto) GetSubServices() []UpdateTransportBaseDtoSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *MergeableTransportDto) GetSubServicesOk() (*[]UpdateTransportBaseDtoSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *MergeableTransportDto) SetSubServices(v []UpdateTransportBaseDtoSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *MergeableTransportDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetPacking

`func (o *MergeableTransportDto) GetPacking() UpdateTransportBaseDtoPackingDto`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *MergeableTransportDto) GetPackingOk() (*UpdateTransportBaseDtoPackingDto, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *MergeableTransportDto) SetPacking(v UpdateTransportBaseDtoPackingDto)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *MergeableTransportDto) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetCustomerCustomParameters

`func (o *MergeableTransportDto) GetCustomerCustomParameters() []UpdateTransportBaseDtoCustomParameterDto`

GetCustomerCustomParameters returns the CustomerCustomParameters field if non-nil, zero value otherwise.

### GetCustomerCustomParametersOk

`func (o *MergeableTransportDto) GetCustomerCustomParametersOk() (*[]UpdateTransportBaseDtoCustomParameterDto, bool)`

GetCustomerCustomParametersOk returns a tuple with the CustomerCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomParameters

`func (o *MergeableTransportDto) SetCustomerCustomParameters(v []UpdateTransportBaseDtoCustomParameterDto)`

SetCustomerCustomParameters sets CustomerCustomParameters field to given value.

### HasCustomerCustomParameters

`func (o *MergeableTransportDto) HasCustomerCustomParameters() bool`

HasCustomerCustomParameters returns a boolean if a field has been set.

### GetServiceCustomParameters

`func (o *MergeableTransportDto) GetServiceCustomParameters() []UpdateTransportBaseDtoCustomParameterDto`

GetServiceCustomParameters returns the ServiceCustomParameters field if non-nil, zero value otherwise.

### GetServiceCustomParametersOk

`func (o *MergeableTransportDto) GetServiceCustomParametersOk() (*[]UpdateTransportBaseDtoCustomParameterDto, bool)`

GetServiceCustomParametersOk returns a tuple with the ServiceCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCustomParameters

`func (o *MergeableTransportDto) SetServiceCustomParameters(v []UpdateTransportBaseDtoCustomParameterDto)`

SetServiceCustomParameters sets ServiceCustomParameters field to given value.

### HasServiceCustomParameters

`func (o *MergeableTransportDto) HasServiceCustomParameters() bool`

HasServiceCustomParameters returns a boolean if a field has been set.

### GetDangerousGoods

`func (o *MergeableTransportDto) GetDangerousGoods() []UpdateTransportBaseDtoDangerousGoodDto`

GetDangerousGoods returns the DangerousGoods field if non-nil, zero value otherwise.

### GetDangerousGoodsOk

`func (o *MergeableTransportDto) GetDangerousGoodsOk() (*[]UpdateTransportBaseDtoDangerousGoodDto, bool)`

GetDangerousGoodsOk returns a tuple with the DangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoods

`func (o *MergeableTransportDto) SetDangerousGoods(v []UpdateTransportBaseDtoDangerousGoodDto)`

SetDangerousGoods sets DangerousGoods field to given value.

### HasDangerousGoods

`func (o *MergeableTransportDto) HasDangerousGoods() bool`

HasDangerousGoods returns a boolean if a field has been set.

### GetReference1

`func (o *MergeableTransportDto) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *MergeableTransportDto) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *MergeableTransportDto) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *MergeableTransportDto) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *MergeableTransportDto) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *MergeableTransportDto) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *MergeableTransportDto) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *MergeableTransportDto) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *MergeableTransportDto) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *MergeableTransportDto) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *MergeableTransportDto) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *MergeableTransportDto) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetDistanceKm

`func (o *MergeableTransportDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *MergeableTransportDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *MergeableTransportDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *MergeableTransportDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetTotalDurationMinutes

`func (o *MergeableTransportDto) GetTotalDurationMinutes() int32`

GetTotalDurationMinutes returns the TotalDurationMinutes field if non-nil, zero value otherwise.

### GetTotalDurationMinutesOk

`func (o *MergeableTransportDto) GetTotalDurationMinutesOk() (*int32, bool)`

GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMinutes

`func (o *MergeableTransportDto) SetTotalDurationMinutes(v int32)`

SetTotalDurationMinutes sets TotalDurationMinutes field to given value.

### HasTotalDurationMinutes

`func (o *MergeableTransportDto) HasTotalDurationMinutes() bool`

HasTotalDurationMinutes returns a boolean if a field has been set.

### GetIsStrategic

`func (o *MergeableTransportDto) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *MergeableTransportDto) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *MergeableTransportDto) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *MergeableTransportDto) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetIsShared

`func (o *MergeableTransportDto) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MergeableTransportDto) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *MergeableTransportDto) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *MergeableTransportDto) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetResponsibleOperatorId

`func (o *MergeableTransportDto) GetResponsibleOperatorId() int32`

GetResponsibleOperatorId returns the ResponsibleOperatorId field if non-nil, zero value otherwise.

### GetResponsibleOperatorIdOk

`func (o *MergeableTransportDto) GetResponsibleOperatorIdOk() (*int32, bool)`

GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleOperatorId

`func (o *MergeableTransportDto) SetResponsibleOperatorId(v int32)`

SetResponsibleOperatorId sets ResponsibleOperatorId field to given value.

### HasResponsibleOperatorId

`func (o *MergeableTransportDto) HasResponsibleOperatorId() bool`

HasResponsibleOperatorId returns a boolean if a field has been set.

### GetGasEmission

`func (o *MergeableTransportDto) GetGasEmission() float64`

GetGasEmission returns the GasEmission field if non-nil, zero value otherwise.

### GetGasEmissionOk

`func (o *MergeableTransportDto) GetGasEmissionOk() (*float64, bool)`

GetGasEmissionOk returns a tuple with the GasEmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEmission

`func (o *MergeableTransportDto) SetGasEmission(v float64)`

SetGasEmission sets GasEmission field to given value.

### HasGasEmission

`func (o *MergeableTransportDto) HasGasEmission() bool`

HasGasEmission returns a boolean if a field has been set.

### GetIsGasEmissionForced

`func (o *MergeableTransportDto) GetIsGasEmissionForced() bool`

GetIsGasEmissionForced returns the IsGasEmissionForced field if non-nil, zero value otherwise.

### GetIsGasEmissionForcedOk

`func (o *MergeableTransportDto) GetIsGasEmissionForcedOk() (*bool, bool)`

GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGasEmissionForced

`func (o *MergeableTransportDto) SetIsGasEmissionForced(v bool)`

SetIsGasEmissionForced sets IsGasEmissionForced field to given value.

### HasIsGasEmissionForced

`func (o *MergeableTransportDto) HasIsGasEmissionForced() bool`

HasIsGasEmissionForced returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *MergeableTransportDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *MergeableTransportDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *MergeableTransportDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *MergeableTransportDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *MergeableTransportDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *MergeableTransportDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *MergeableTransportDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *MergeableTransportDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetSellPricingPathId

`func (o *MergeableTransportDto) GetSellPricingPathId() int32`

GetSellPricingPathId returns the SellPricingPathId field if non-nil, zero value otherwise.

### GetSellPricingPathIdOk

`func (o *MergeableTransportDto) GetSellPricingPathIdOk() (*int32, bool)`

GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPricingPathId

`func (o *MergeableTransportDto) SetSellPricingPathId(v int32)`

SetSellPricingPathId sets SellPricingPathId field to given value.

### HasSellPricingPathId

`func (o *MergeableTransportDto) HasSellPricingPathId() bool`

HasSellPricingPathId returns a boolean if a field has been set.

### GetComment

`func (o *MergeableTransportDto) GetComment() UpdateTransportBaseDtoCommentDto`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MergeableTransportDto) GetCommentOk() (*UpdateTransportBaseDtoCommentDto, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MergeableTransportDto) SetComment(v UpdateTransportBaseDtoCommentDto)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MergeableTransportDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTransportBillAddress

`func (o *MergeableTransportDto) GetTransportBillAddress() UpdateTransportBaseDtoTransportBillAddressDto`

GetTransportBillAddress returns the TransportBillAddress field if non-nil, zero value otherwise.

### GetTransportBillAddressOk

`func (o *MergeableTransportDto) GetTransportBillAddressOk() (*UpdateTransportBaseDtoTransportBillAddressDto, bool)`

GetTransportBillAddressOk returns a tuple with the TransportBillAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportBillAddress

`func (o *MergeableTransportDto) SetTransportBillAddress(v UpdateTransportBaseDtoTransportBillAddressDto)`

SetTransportBillAddress sets TransportBillAddress field to given value.

### HasTransportBillAddress

`func (o *MergeableTransportDto) HasTransportBillAddress() bool`

HasTransportBillAddress returns a boolean if a field has been set.

### GetCashOnDelivery

`func (o *MergeableTransportDto) GetCashOnDelivery() UpdateTransportBaseDtoCashOnDeliveryDto`

GetCashOnDelivery returns the CashOnDelivery field if non-nil, zero value otherwise.

### GetCashOnDeliveryOk

`func (o *MergeableTransportDto) GetCashOnDeliveryOk() (*UpdateTransportBaseDtoCashOnDeliveryDto, bool)`

GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDelivery

`func (o *MergeableTransportDto) SetCashOnDelivery(v UpdateTransportBaseDtoCashOnDeliveryDto)`

SetCashOnDelivery sets CashOnDelivery field to given value.

### HasCashOnDelivery

`func (o *MergeableTransportDto) HasCashOnDelivery() bool`

HasCashOnDelivery returns a boolean if a field has been set.

### GetCommunication

`func (o *MergeableTransportDto) GetCommunication() TransportCommunicationConfigurationDto`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *MergeableTransportDto) GetCommunicationOk() (*TransportCommunicationConfigurationDto, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *MergeableTransportDto) SetCommunication(v TransportCommunicationConfigurationDto)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *MergeableTransportDto) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


