# UpdateTransportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBehavior** | Pointer to [**UpdateTransportDtoApiBehaviorDto**](UpdateTransportDtoApiBehaviorDto.md) |  | [optional] 
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

### NewUpdateTransportDto

`func NewUpdateTransportDto() *UpdateTransportDto`

NewUpdateTransportDto instantiates a new UpdateTransportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransportDtoWithDefaults

`func NewUpdateTransportDtoWithDefaults() *UpdateTransportDto`

NewUpdateTransportDtoWithDefaults instantiates a new UpdateTransportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBehavior

`func (o *UpdateTransportDto) GetApiBehavior() UpdateTransportDtoApiBehaviorDto`

GetApiBehavior returns the ApiBehavior field if non-nil, zero value otherwise.

### GetApiBehaviorOk

`func (o *UpdateTransportDto) GetApiBehaviorOk() (*UpdateTransportDtoApiBehaviorDto, bool)`

GetApiBehaviorOk returns a tuple with the ApiBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBehavior

`func (o *UpdateTransportDto) SetApiBehavior(v UpdateTransportDtoApiBehaviorDto)`

SetApiBehavior sets ApiBehavior field to given value.

### HasApiBehavior

`func (o *UpdateTransportDto) HasApiBehavior() bool`

HasApiBehavior returns a boolean if a field has been set.

### GetCustomerCode

`func (o *UpdateTransportDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *UpdateTransportDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *UpdateTransportDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *UpdateTransportDto) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *UpdateTransportDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *UpdateTransportDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *UpdateTransportDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *UpdateTransportDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetOrdererCode

`func (o *UpdateTransportDto) GetOrdererCode() string`

GetOrdererCode returns the OrdererCode field if non-nil, zero value otherwise.

### GetOrdererCodeOk

`func (o *UpdateTransportDto) GetOrdererCodeOk() (*string, bool)`

GetOrdererCodeOk returns a tuple with the OrdererCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererCode

`func (o *UpdateTransportDto) SetOrdererCode(v string)`

SetOrdererCode sets OrdererCode field to given value.

### HasOrdererCode

`func (o *UpdateTransportDto) HasOrdererCode() bool`

HasOrdererCode returns a boolean if a field has been set.

### GetOrdererName

`func (o *UpdateTransportDto) GetOrdererName() string`

GetOrdererName returns the OrdererName field if non-nil, zero value otherwise.

### GetOrdererNameOk

`func (o *UpdateTransportDto) GetOrdererNameOk() (*string, bool)`

GetOrdererNameOk returns a tuple with the OrdererName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererName

`func (o *UpdateTransportDto) SetOrdererName(v string)`

SetOrdererName sets OrdererName field to given value.

### HasOrdererName

`func (o *UpdateTransportDto) HasOrdererName() bool`

HasOrdererName returns a boolean if a field has been set.

### GetSign

`func (o *UpdateTransportDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *UpdateTransportDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *UpdateTransportDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *UpdateTransportDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetSecretCode

`func (o *UpdateTransportDto) GetSecretCode() string`

GetSecretCode returns the SecretCode field if non-nil, zero value otherwise.

### GetSecretCodeOk

`func (o *UpdateTransportDto) GetSecretCodeOk() (*string, bool)`

GetSecretCodeOk returns a tuple with the SecretCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCode

`func (o *UpdateTransportDto) SetSecretCode(v string)`

SetSecretCode sets SecretCode field to given value.

### HasSecretCode

`func (o *UpdateTransportDto) HasSecretCode() bool`

HasSecretCode returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateTransportDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateTransportDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateTransportDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateTransportDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsRoundTrip

`func (o *UpdateTransportDto) GetIsRoundTrip() bool`

GetIsRoundTrip returns the IsRoundTrip field if non-nil, zero value otherwise.

### GetIsRoundTripOk

`func (o *UpdateTransportDto) GetIsRoundTripOk() (*bool, bool)`

GetIsRoundTripOk returns a tuple with the IsRoundTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundTrip

`func (o *UpdateTransportDto) SetIsRoundTrip(v bool)`

SetIsRoundTrip sets IsRoundTrip field to given value.

### HasIsRoundTrip

`func (o *UpdateTransportDto) HasIsRoundTrip() bool`

HasIsRoundTrip returns a boolean if a field has been set.

### GetPickupStep

`func (o *UpdateTransportDto) GetPickupStep() UpdateTransportBaseDtoStepDto`

GetPickupStep returns the PickupStep field if non-nil, zero value otherwise.

### GetPickupStepOk

`func (o *UpdateTransportDto) GetPickupStepOk() (*UpdateTransportBaseDtoStepDto, bool)`

GetPickupStepOk returns a tuple with the PickupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStep

`func (o *UpdateTransportDto) SetPickupStep(v UpdateTransportBaseDtoStepDto)`

SetPickupStep sets PickupStep field to given value.

### HasPickupStep

`func (o *UpdateTransportDto) HasPickupStep() bool`

HasPickupStep returns a boolean if a field has been set.

### GetDeliveryStep

`func (o *UpdateTransportDto) GetDeliveryStep() UpdateTransportBaseDtoStepDto`

GetDeliveryStep returns the DeliveryStep field if non-nil, zero value otherwise.

### GetDeliveryStepOk

`func (o *UpdateTransportDto) GetDeliveryStepOk() (*UpdateTransportBaseDtoStepDto, bool)`

GetDeliveryStepOk returns a tuple with the DeliveryStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStep

`func (o *UpdateTransportDto) SetDeliveryStep(v UpdateTransportBaseDtoStepDto)`

SetDeliveryStep sets DeliveryStep field to given value.

### HasDeliveryStep

`func (o *UpdateTransportDto) HasDeliveryStep() bool`

HasDeliveryStep returns a boolean if a field has been set.

### GetServiceCode

`func (o *UpdateTransportDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *UpdateTransportDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *UpdateTransportDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *UpdateTransportDto) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetSubServices

`func (o *UpdateTransportDto) GetSubServices() []UpdateTransportBaseDtoSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *UpdateTransportDto) GetSubServicesOk() (*[]UpdateTransportBaseDtoSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *UpdateTransportDto) SetSubServices(v []UpdateTransportBaseDtoSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *UpdateTransportDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetPacking

`func (o *UpdateTransportDto) GetPacking() UpdateTransportBaseDtoPackingDto`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *UpdateTransportDto) GetPackingOk() (*UpdateTransportBaseDtoPackingDto, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *UpdateTransportDto) SetPacking(v UpdateTransportBaseDtoPackingDto)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *UpdateTransportDto) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetCustomerCustomParameters

`func (o *UpdateTransportDto) GetCustomerCustomParameters() []UpdateTransportBaseDtoCustomParameterDto`

GetCustomerCustomParameters returns the CustomerCustomParameters field if non-nil, zero value otherwise.

### GetCustomerCustomParametersOk

`func (o *UpdateTransportDto) GetCustomerCustomParametersOk() (*[]UpdateTransportBaseDtoCustomParameterDto, bool)`

GetCustomerCustomParametersOk returns a tuple with the CustomerCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomParameters

`func (o *UpdateTransportDto) SetCustomerCustomParameters(v []UpdateTransportBaseDtoCustomParameterDto)`

SetCustomerCustomParameters sets CustomerCustomParameters field to given value.

### HasCustomerCustomParameters

`func (o *UpdateTransportDto) HasCustomerCustomParameters() bool`

HasCustomerCustomParameters returns a boolean if a field has been set.

### GetServiceCustomParameters

`func (o *UpdateTransportDto) GetServiceCustomParameters() []UpdateTransportBaseDtoCustomParameterDto`

GetServiceCustomParameters returns the ServiceCustomParameters field if non-nil, zero value otherwise.

### GetServiceCustomParametersOk

`func (o *UpdateTransportDto) GetServiceCustomParametersOk() (*[]UpdateTransportBaseDtoCustomParameterDto, bool)`

GetServiceCustomParametersOk returns a tuple with the ServiceCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCustomParameters

`func (o *UpdateTransportDto) SetServiceCustomParameters(v []UpdateTransportBaseDtoCustomParameterDto)`

SetServiceCustomParameters sets ServiceCustomParameters field to given value.

### HasServiceCustomParameters

`func (o *UpdateTransportDto) HasServiceCustomParameters() bool`

HasServiceCustomParameters returns a boolean if a field has been set.

### GetDangerousGoods

`func (o *UpdateTransportDto) GetDangerousGoods() []UpdateTransportBaseDtoDangerousGoodDto`

GetDangerousGoods returns the DangerousGoods field if non-nil, zero value otherwise.

### GetDangerousGoodsOk

`func (o *UpdateTransportDto) GetDangerousGoodsOk() (*[]UpdateTransportBaseDtoDangerousGoodDto, bool)`

GetDangerousGoodsOk returns a tuple with the DangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoods

`func (o *UpdateTransportDto) SetDangerousGoods(v []UpdateTransportBaseDtoDangerousGoodDto)`

SetDangerousGoods sets DangerousGoods field to given value.

### HasDangerousGoods

`func (o *UpdateTransportDto) HasDangerousGoods() bool`

HasDangerousGoods returns a boolean if a field has been set.

### GetReference1

`func (o *UpdateTransportDto) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *UpdateTransportDto) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *UpdateTransportDto) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *UpdateTransportDto) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *UpdateTransportDto) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *UpdateTransportDto) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *UpdateTransportDto) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *UpdateTransportDto) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *UpdateTransportDto) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *UpdateTransportDto) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *UpdateTransportDto) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *UpdateTransportDto) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetDistanceKm

`func (o *UpdateTransportDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *UpdateTransportDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *UpdateTransportDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *UpdateTransportDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetTotalDurationMinutes

`func (o *UpdateTransportDto) GetTotalDurationMinutes() int32`

GetTotalDurationMinutes returns the TotalDurationMinutes field if non-nil, zero value otherwise.

### GetTotalDurationMinutesOk

`func (o *UpdateTransportDto) GetTotalDurationMinutesOk() (*int32, bool)`

GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMinutes

`func (o *UpdateTransportDto) SetTotalDurationMinutes(v int32)`

SetTotalDurationMinutes sets TotalDurationMinutes field to given value.

### HasTotalDurationMinutes

`func (o *UpdateTransportDto) HasTotalDurationMinutes() bool`

HasTotalDurationMinutes returns a boolean if a field has been set.

### GetIsStrategic

`func (o *UpdateTransportDto) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *UpdateTransportDto) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *UpdateTransportDto) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *UpdateTransportDto) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetIsShared

`func (o *UpdateTransportDto) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *UpdateTransportDto) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *UpdateTransportDto) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *UpdateTransportDto) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetResponsibleOperatorId

`func (o *UpdateTransportDto) GetResponsibleOperatorId() int32`

GetResponsibleOperatorId returns the ResponsibleOperatorId field if non-nil, zero value otherwise.

### GetResponsibleOperatorIdOk

`func (o *UpdateTransportDto) GetResponsibleOperatorIdOk() (*int32, bool)`

GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleOperatorId

`func (o *UpdateTransportDto) SetResponsibleOperatorId(v int32)`

SetResponsibleOperatorId sets ResponsibleOperatorId field to given value.

### HasResponsibleOperatorId

`func (o *UpdateTransportDto) HasResponsibleOperatorId() bool`

HasResponsibleOperatorId returns a boolean if a field has been set.

### GetGasEmission

`func (o *UpdateTransportDto) GetGasEmission() float64`

GetGasEmission returns the GasEmission field if non-nil, zero value otherwise.

### GetGasEmissionOk

`func (o *UpdateTransportDto) GetGasEmissionOk() (*float64, bool)`

GetGasEmissionOk returns a tuple with the GasEmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEmission

`func (o *UpdateTransportDto) SetGasEmission(v float64)`

SetGasEmission sets GasEmission field to given value.

### HasGasEmission

`func (o *UpdateTransportDto) HasGasEmission() bool`

HasGasEmission returns a boolean if a field has been set.

### GetIsGasEmissionForced

`func (o *UpdateTransportDto) GetIsGasEmissionForced() bool`

GetIsGasEmissionForced returns the IsGasEmissionForced field if non-nil, zero value otherwise.

### GetIsGasEmissionForcedOk

`func (o *UpdateTransportDto) GetIsGasEmissionForcedOk() (*bool, bool)`

GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGasEmissionForced

`func (o *UpdateTransportDto) SetIsGasEmissionForced(v bool)`

SetIsGasEmissionForced sets IsGasEmissionForced field to given value.

### HasIsGasEmissionForced

`func (o *UpdateTransportDto) HasIsGasEmissionForced() bool`

HasIsGasEmissionForced returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *UpdateTransportDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *UpdateTransportDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *UpdateTransportDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *UpdateTransportDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *UpdateTransportDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *UpdateTransportDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *UpdateTransportDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *UpdateTransportDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetSellPricingPathId

`func (o *UpdateTransportDto) GetSellPricingPathId() int32`

GetSellPricingPathId returns the SellPricingPathId field if non-nil, zero value otherwise.

### GetSellPricingPathIdOk

`func (o *UpdateTransportDto) GetSellPricingPathIdOk() (*int32, bool)`

GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPricingPathId

`func (o *UpdateTransportDto) SetSellPricingPathId(v int32)`

SetSellPricingPathId sets SellPricingPathId field to given value.

### HasSellPricingPathId

`func (o *UpdateTransportDto) HasSellPricingPathId() bool`

HasSellPricingPathId returns a boolean if a field has been set.

### GetComment

`func (o *UpdateTransportDto) GetComment() UpdateTransportBaseDtoCommentDto`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateTransportDto) GetCommentOk() (*UpdateTransportBaseDtoCommentDto, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateTransportDto) SetComment(v UpdateTransportBaseDtoCommentDto)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateTransportDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTransportBillAddress

`func (o *UpdateTransportDto) GetTransportBillAddress() UpdateTransportBaseDtoTransportBillAddressDto`

GetTransportBillAddress returns the TransportBillAddress field if non-nil, zero value otherwise.

### GetTransportBillAddressOk

`func (o *UpdateTransportDto) GetTransportBillAddressOk() (*UpdateTransportBaseDtoTransportBillAddressDto, bool)`

GetTransportBillAddressOk returns a tuple with the TransportBillAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportBillAddress

`func (o *UpdateTransportDto) SetTransportBillAddress(v UpdateTransportBaseDtoTransportBillAddressDto)`

SetTransportBillAddress sets TransportBillAddress field to given value.

### HasTransportBillAddress

`func (o *UpdateTransportDto) HasTransportBillAddress() bool`

HasTransportBillAddress returns a boolean if a field has been set.

### GetCashOnDelivery

`func (o *UpdateTransportDto) GetCashOnDelivery() UpdateTransportBaseDtoCashOnDeliveryDto`

GetCashOnDelivery returns the CashOnDelivery field if non-nil, zero value otherwise.

### GetCashOnDeliveryOk

`func (o *UpdateTransportDto) GetCashOnDeliveryOk() (*UpdateTransportBaseDtoCashOnDeliveryDto, bool)`

GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDelivery

`func (o *UpdateTransportDto) SetCashOnDelivery(v UpdateTransportBaseDtoCashOnDeliveryDto)`

SetCashOnDelivery sets CashOnDelivery field to given value.

### HasCashOnDelivery

`func (o *UpdateTransportDto) HasCashOnDelivery() bool`

HasCashOnDelivery returns a boolean if a field has been set.

### GetCommunication

`func (o *UpdateTransportDto) GetCommunication() TransportCommunicationConfigurationDto`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *UpdateTransportDto) GetCommunicationOk() (*TransportCommunicationConfigurationDto, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *UpdateTransportDto) SetCommunication(v TransportCommunicationConfigurationDto)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *UpdateTransportDto) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


