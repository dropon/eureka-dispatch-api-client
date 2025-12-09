# CreateTransportDto

## Properties

| Name                         | Type                                                                                                     | Description                                                                                                                                                                                                                                                                                                                          | Notes      |
| ---------------------------- | -------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------- |
| **CustomerCode**             | **string**                                                                                               | Mandatory. Customer&#39;s unique code.                                                                                                                                                                                                                                                                                               |
| **AgencyCode**               | **string**                                                                                               | Mandatory. Agency unique code.                                                                                                                                                                                                                                                                                                       |
| **OrdererCode**              | Pointer to **string**                                                                                    | Orderer unique code.                                                                                                                                                                                                                                                                                                                 | [optional] |
| **OrdererName**              | Pointer to **string**                                                                                    | Orderer name, when orderer code is not specified.                                                                                                                                                                                                                                                                                    | [optional] |
| **Sign**                     | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **SecretCode**               | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **Notes**                    | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **IsRoundTrip**              | Pointer to **bool**                                                                                      |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **CustomerCallDateTime**     | Pointer to **Time**                                                                                      | The date and time at which the customer have contacted the transport company to request the transport. If not set, will default to the service current local date and time at server time zone.                                                                                                                                      | [optional] |
| **PickupStep**               | [**CreateTransportDtoStepDto**](CreateTransportDtoStepDto.md)                                            |                                                                                                                                                                                                                                                                                                                                      |
| **DeliveryStep**             | [**CreateTransportDtoStepDto**](CreateTransportDtoStepDto.md)                                            |                                                                                                                                                                                                                                                                                                                                      |
| **ServiceCode**              | **string**                                                                                               | Mandatory. Service&#39;s unique code.                                                                                                                                                                                                                                                                                                |
| **SubServices**              | Pointer to [**[]CreateTransportDtoSubServiceDto**](CreateTransportDtoSubServiceDto.md)                   | Transport&#39;s sub services. When this list is null or empty, the transport uses service&#39;s default sub services. When at least an element is specified in this list, the transport uses only the specified sub services: service&#39;s default sub services are not included and must be specified manually if they are needed. | [optional] |
| **Packing**                  | Pointer to [**CreateTransportDtoPackingDto**](CreateTransportDtoPackingDto.md)                           |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **CustomerCustomParameters** | Pointer to [**[]CreateTransportDtoCustomParameterDto**](CreateTransportDtoCustomParameterDto.md)         | Custom parameters related to the customer.                                                                                                                                                                                                                                                                                           | [optional] |
| **ServiceCustomParameters**  | Pointer to [**[]CreateTransportDtoCustomParameterDto**](CreateTransportDtoCustomParameterDto.md)         | Custom parameters related to the service.                                                                                                                                                                                                                                                                                            | [optional] |
| **Reference1**               | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **Reference2**               | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **Reference3**               | Pointer to **string**                                                                                    |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **DistanceKm**               | Pointer to **float64**                                                                                   | Distance of the transport in kilometers.                                                                                                                                                                                                                                                                                             | [optional] |
| **TotalDurationMinutes**     | Pointer to **int32**                                                                                     | Total duration of the transport in minutes                                                                                                                                                                                                                                                                                           | [optional] |
| **IsStrategic**              | Pointer to **bool**                                                                                      | Indicates if the transport has a strategic importance. When not specified, this information is filled from the customer.                                                                                                                                                                                                             | [optional] |
| **IsShared**                 | Pointer to **bool**                                                                                      | Indicates if the transport is &#39;shared&#39; between agencies. If true, the transport will be visible on Dispatch planning via the &#39;supervising&#39; mode.                                                                                                                                                                     | [optional] |
| **ResponsibleOperatorId**    | Pointer to **int32**                                                                                     | The identifier of a dispatch user that is responsible for following the execution of the transport.                                                                                                                                                                                                                                  | [optional] |
| **GasEmission**              | Pointer to **float64**                                                                                   | Gas emission (CO2) in Kilograms. Used only when gas emission is forced.                                                                                                                                                                                                                                                              | [optional] |
| **IsGasEmissionForced**      | Pointer to **bool**                                                                                      |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **ForcedSellPrice**          | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices,Set prices in orders.                                                                                                                                                                                                                                   | [optional] |
| **ForcedBuyPrice**           | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices,Set prices in orders.                                                                                                                                                                                                                                   | [optional] |
| **SellPricingPathId**        | Pointer to **int32**                                                                                     | Used to force a specific sell pricing path. The pricing path must be a customer pricing path (not a subcontractor&#39;s one). When provided, the addresses are overwritten with the ones defined in the provided sell pricing path. The provided service must match the one defined in the pricing path.                             | [optional] |
| **Comment**                  | Pointer to [**CreateTransportDtoCommentDto**](CreateTransportDtoCommentDto.md)                           |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **TransportBillAddress**     | Pointer to [**CreateTransportDtoTransportBillAddressDto**](CreateTransportDtoTransportBillAddressDto.md) |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **CashOnDelivery**           | Pointer to [**CreateTransportDtoCashOnDeliveryDto**](CreateTransportDtoCashOnDeliveryDto.md)             |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **Communication**            | Pointer to [**TransportCommunicationConfigurationDto**](TransportCommunicationConfigurationDto.md)       |                                                                                                                                                                                                                                                                                                                                      | [optional] |
| **DangerousGoods**           | Pointer to [**[]CreateTransportDtoDangerousGoodDto**](CreateTransportDtoDangerousGoodDto.md)             |                                                                                                                                                                                                                                                                                                                                      | [optional] |

## Methods

### NewCreateTransportDto

`func NewCreateTransportDto(customerCode string, agencyCode string, pickupStep CreateTransportDtoStepDto, deliveryStep CreateTransportDtoStepDto, serviceCode string, ) *CreateTransportDto`

NewCreateTransportDto instantiates a new CreateTransportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDtoWithDefaults

`func NewCreateTransportDtoWithDefaults() *CreateTransportDto`

NewCreateTransportDtoWithDefaults instantiates a new CreateTransportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerCode

`func (o *CreateTransportDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *CreateTransportDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *CreateTransportDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### GetAgencyCode

`func (o *CreateTransportDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *CreateTransportDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *CreateTransportDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### GetOrdererCode

`func (o *CreateTransportDto) GetOrdererCode() string`

GetOrdererCode returns the OrdererCode field if non-nil, zero value otherwise.

### GetOrdererCodeOk

`func (o *CreateTransportDto) GetOrdererCodeOk() (*string, bool)`

GetOrdererCodeOk returns a tuple with the OrdererCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererCode

`func (o *CreateTransportDto) SetOrdererCode(v string)`

SetOrdererCode sets OrdererCode field to given value.

### HasOrdererCode

`func (o *CreateTransportDto) HasOrdererCode() bool`

HasOrdererCode returns a boolean if a field has been set.

### GetOrdererName

`func (o *CreateTransportDto) GetOrdererName() string`

GetOrdererName returns the OrdererName field if non-nil, zero value otherwise.

### GetOrdererNameOk

`func (o *CreateTransportDto) GetOrdererNameOk() (*string, bool)`

GetOrdererNameOk returns a tuple with the OrdererName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererName

`func (o *CreateTransportDto) SetOrdererName(v string)`

SetOrdererName sets OrdererName field to given value.

### HasOrdererName

`func (o *CreateTransportDto) HasOrdererName() bool`

HasOrdererName returns a boolean if a field has been set.

### GetSign

`func (o *CreateTransportDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *CreateTransportDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *CreateTransportDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *CreateTransportDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetSecretCode

`func (o *CreateTransportDto) GetSecretCode() string`

GetSecretCode returns the SecretCode field if non-nil, zero value otherwise.

### GetSecretCodeOk

`func (o *CreateTransportDto) GetSecretCodeOk() (*string, bool)`

GetSecretCodeOk returns a tuple with the SecretCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCode

`func (o *CreateTransportDto) SetSecretCode(v string)`

SetSecretCode sets SecretCode field to given value.

### HasSecretCode

`func (o *CreateTransportDto) HasSecretCode() bool`

HasSecretCode returns a boolean if a field has been set.

### GetNotes

`func (o *CreateTransportDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateTransportDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateTransportDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateTransportDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsRoundTrip

`func (o *CreateTransportDto) GetIsRoundTrip() bool`

GetIsRoundTrip returns the IsRoundTrip field if non-nil, zero value otherwise.

### GetIsRoundTripOk

`func (o *CreateTransportDto) GetIsRoundTripOk() (*bool, bool)`

GetIsRoundTripOk returns a tuple with the IsRoundTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundTrip

`func (o *CreateTransportDto) SetIsRoundTrip(v bool)`

SetIsRoundTrip sets IsRoundTrip field to given value.

### HasIsRoundTrip

`func (o *CreateTransportDto) HasIsRoundTrip() bool`

HasIsRoundTrip returns a boolean if a field has been set.

### GetCustomerCallDateTime

`func (o *CreateTransportDto) GetCustomerCallDateTime() Time`

GetCustomerCallDateTime returns the CustomerCallDateTime field if non-nil, zero value otherwise.

### GetCustomerCallDateTimeOk

`func (o *CreateTransportDto) GetCustomerCallDateTimeOk() (*Time, bool)`

GetCustomerCallDateTimeOk returns a tuple with the CustomerCallDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCallDateTime

`func (o *CreateTransportDto) SetCustomerCallDateTime(v Time)`

SetCustomerCallDateTime sets CustomerCallDateTime field to given value.

### HasCustomerCallDateTime

`func (o *CreateTransportDto) HasCustomerCallDateTime() bool`

HasCustomerCallDateTime returns a boolean if a field has been set.

### GetPickupStep

`func (o *CreateTransportDto) GetPickupStep() CreateTransportDtoStepDto`

GetPickupStep returns the PickupStep field if non-nil, zero value otherwise.

### GetPickupStepOk

`func (o *CreateTransportDto) GetPickupStepOk() (*CreateTransportDtoStepDto, bool)`

GetPickupStepOk returns a tuple with the PickupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStep

`func (o *CreateTransportDto) SetPickupStep(v CreateTransportDtoStepDto)`

SetPickupStep sets PickupStep field to given value.

### GetDeliveryStep

`func (o *CreateTransportDto) GetDeliveryStep() CreateTransportDtoStepDto`

GetDeliveryStep returns the DeliveryStep field if non-nil, zero value otherwise.

### GetDeliveryStepOk

`func (o *CreateTransportDto) GetDeliveryStepOk() (*CreateTransportDtoStepDto, bool)`

GetDeliveryStepOk returns a tuple with the DeliveryStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStep

`func (o *CreateTransportDto) SetDeliveryStep(v CreateTransportDtoStepDto)`

SetDeliveryStep sets DeliveryStep field to given value.

### GetServiceCode

`func (o *CreateTransportDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *CreateTransportDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *CreateTransportDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### GetSubServices

`func (o *CreateTransportDto) GetSubServices() []CreateTransportDtoSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *CreateTransportDto) GetSubServicesOk() (*[]CreateTransportDtoSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *CreateTransportDto) SetSubServices(v []CreateTransportDtoSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *CreateTransportDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetPacking

`func (o *CreateTransportDto) GetPacking() CreateTransportDtoPackingDto`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *CreateTransportDto) GetPackingOk() (*CreateTransportDtoPackingDto, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *CreateTransportDto) SetPacking(v CreateTransportDtoPackingDto)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *CreateTransportDto) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetCustomerCustomParameters

`func (o *CreateTransportDto) GetCustomerCustomParameters() []CreateTransportDtoCustomParameterDto`

GetCustomerCustomParameters returns the CustomerCustomParameters field if non-nil, zero value otherwise.

### GetCustomerCustomParametersOk

`func (o *CreateTransportDto) GetCustomerCustomParametersOk() (*[]CreateTransportDtoCustomParameterDto, bool)`

GetCustomerCustomParametersOk returns a tuple with the CustomerCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCustomParameters

`func (o *CreateTransportDto) SetCustomerCustomParameters(v []CreateTransportDtoCustomParameterDto)`

SetCustomerCustomParameters sets CustomerCustomParameters field to given value.

### HasCustomerCustomParameters

`func (o *CreateTransportDto) HasCustomerCustomParameters() bool`

HasCustomerCustomParameters returns a boolean if a field has been set.

### GetServiceCustomParameters

`func (o *CreateTransportDto) GetServiceCustomParameters() []CreateTransportDtoCustomParameterDto`

GetServiceCustomParameters returns the ServiceCustomParameters field if non-nil, zero value otherwise.

### GetServiceCustomParametersOk

`func (o *CreateTransportDto) GetServiceCustomParametersOk() (*[]CreateTransportDtoCustomParameterDto, bool)`

GetServiceCustomParametersOk returns a tuple with the ServiceCustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCustomParameters

`func (o *CreateTransportDto) SetServiceCustomParameters(v []CreateTransportDtoCustomParameterDto)`

SetServiceCustomParameters sets ServiceCustomParameters field to given value.

### HasServiceCustomParameters

`func (o *CreateTransportDto) HasServiceCustomParameters() bool`

HasServiceCustomParameters returns a boolean if a field has been set.

### GetReference1

`func (o *CreateTransportDto) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *CreateTransportDto) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *CreateTransportDto) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *CreateTransportDto) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *CreateTransportDto) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *CreateTransportDto) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *CreateTransportDto) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *CreateTransportDto) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *CreateTransportDto) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *CreateTransportDto) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *CreateTransportDto) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *CreateTransportDto) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetDistanceKm

`func (o *CreateTransportDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *CreateTransportDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *CreateTransportDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *CreateTransportDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetTotalDurationMinutes

`func (o *CreateTransportDto) GetTotalDurationMinutes() int32`

GetTotalDurationMinutes returns the TotalDurationMinutes field if non-nil, zero value otherwise.

### GetTotalDurationMinutesOk

`func (o *CreateTransportDto) GetTotalDurationMinutesOk() (*int32, bool)`

GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMinutes

`func (o *CreateTransportDto) SetTotalDurationMinutes(v int32)`

SetTotalDurationMinutes sets TotalDurationMinutes field to given value.

### HasTotalDurationMinutes

`func (o *CreateTransportDto) HasTotalDurationMinutes() bool`

HasTotalDurationMinutes returns a boolean if a field has been set.

### GetIsStrategic

`func (o *CreateTransportDto) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *CreateTransportDto) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *CreateTransportDto) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *CreateTransportDto) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetIsShared

`func (o *CreateTransportDto) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *CreateTransportDto) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *CreateTransportDto) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *CreateTransportDto) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetResponsibleOperatorId

`func (o *CreateTransportDto) GetResponsibleOperatorId() int32`

GetResponsibleOperatorId returns the ResponsibleOperatorId field if non-nil, zero value otherwise.

### GetResponsibleOperatorIdOk

`func (o *CreateTransportDto) GetResponsibleOperatorIdOk() (*int32, bool)`

GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleOperatorId

`func (o *CreateTransportDto) SetResponsibleOperatorId(v int32)`

SetResponsibleOperatorId sets ResponsibleOperatorId field to given value.

### HasResponsibleOperatorId

`func (o *CreateTransportDto) HasResponsibleOperatorId() bool`

HasResponsibleOperatorId returns a boolean if a field has been set.

### GetGasEmission

`func (o *CreateTransportDto) GetGasEmission() float64`

GetGasEmission returns the GasEmission field if non-nil, zero value otherwise.

### GetGasEmissionOk

`func (o *CreateTransportDto) GetGasEmissionOk() (*float64, bool)`

GetGasEmissionOk returns a tuple with the GasEmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEmission

`func (o *CreateTransportDto) SetGasEmission(v float64)`

SetGasEmission sets GasEmission field to given value.

### HasGasEmission

`func (o *CreateTransportDto) HasGasEmission() bool`

HasGasEmission returns a boolean if a field has been set.

### GetIsGasEmissionForced

`func (o *CreateTransportDto) GetIsGasEmissionForced() bool`

GetIsGasEmissionForced returns the IsGasEmissionForced field if non-nil, zero value otherwise.

### GetIsGasEmissionForcedOk

`func (o *CreateTransportDto) GetIsGasEmissionForcedOk() (*bool, bool)`

GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGasEmissionForced

`func (o *CreateTransportDto) SetIsGasEmissionForced(v bool)`

SetIsGasEmissionForced sets IsGasEmissionForced field to given value.

### HasIsGasEmissionForced

`func (o *CreateTransportDto) HasIsGasEmissionForced() bool`

HasIsGasEmissionForced returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *CreateTransportDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *CreateTransportDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *CreateTransportDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *CreateTransportDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *CreateTransportDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *CreateTransportDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *CreateTransportDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *CreateTransportDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetSellPricingPathId

`func (o *CreateTransportDto) GetSellPricingPathId() int32`

GetSellPricingPathId returns the SellPricingPathId field if non-nil, zero value otherwise.

### GetSellPricingPathIdOk

`func (o *CreateTransportDto) GetSellPricingPathIdOk() (*int32, bool)`

GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPricingPathId

`func (o *CreateTransportDto) SetSellPricingPathId(v int32)`

SetSellPricingPathId sets SellPricingPathId field to given value.

### HasSellPricingPathId

`func (o *CreateTransportDto) HasSellPricingPathId() bool`

HasSellPricingPathId returns a boolean if a field has been set.

### GetComment

`func (o *CreateTransportDto) GetComment() CreateTransportDtoCommentDto`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateTransportDto) GetCommentOk() (*CreateTransportDtoCommentDto, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateTransportDto) SetComment(v CreateTransportDtoCommentDto)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateTransportDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTransportBillAddress

`func (o *CreateTransportDto) GetTransportBillAddress() CreateTransportDtoTransportBillAddressDto`

GetTransportBillAddress returns the TransportBillAddress field if non-nil, zero value otherwise.

### GetTransportBillAddressOk

`func (o *CreateTransportDto) GetTransportBillAddressOk() (*CreateTransportDtoTransportBillAddressDto, bool)`

GetTransportBillAddressOk returns a tuple with the TransportBillAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportBillAddress

`func (o *CreateTransportDto) SetTransportBillAddress(v CreateTransportDtoTransportBillAddressDto)`

SetTransportBillAddress sets TransportBillAddress field to given value.

### HasTransportBillAddress

`func (o *CreateTransportDto) HasTransportBillAddress() bool`

HasTransportBillAddress returns a boolean if a field has been set.

### GetCashOnDelivery

`func (o *CreateTransportDto) GetCashOnDelivery() CreateTransportDtoCashOnDeliveryDto`

GetCashOnDelivery returns the CashOnDelivery field if non-nil, zero value otherwise.

### GetCashOnDeliveryOk

`func (o *CreateTransportDto) GetCashOnDeliveryOk() (*CreateTransportDtoCashOnDeliveryDto, bool)`

GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDelivery

`func (o *CreateTransportDto) SetCashOnDelivery(v CreateTransportDtoCashOnDeliveryDto)`

SetCashOnDelivery sets CashOnDelivery field to given value.

### HasCashOnDelivery

`func (o *CreateTransportDto) HasCashOnDelivery() bool`

HasCashOnDelivery returns a boolean if a field has been set.

### GetCommunication

`func (o *CreateTransportDto) GetCommunication() TransportCommunicationConfigurationDto`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *CreateTransportDto) GetCommunicationOk() (*TransportCommunicationConfigurationDto, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *CreateTransportDto) SetCommunication(v TransportCommunicationConfigurationDto)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *CreateTransportDto) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.

### GetDangerousGoods

`func (o *CreateTransportDto) GetDangerousGoods() []CreateTransportDtoDangerousGoodDto`

GetDangerousGoods returns the DangerousGoods field if non-nil, zero value otherwise.

### GetDangerousGoodsOk

`func (o *CreateTransportDto) GetDangerousGoodsOk() (*[]CreateTransportDtoDangerousGoodDto, bool)`

GetDangerousGoodsOk returns a tuple with the DangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoods

`func (o *CreateTransportDto) SetDangerousGoods(v []CreateTransportDtoDangerousGoodDto)`

SetDangerousGoods sets DangerousGoods field to given value.

### HasDangerousGoods

`func (o *CreateTransportDto) HasDangerousGoods() bool`

HasDangerousGoods returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
