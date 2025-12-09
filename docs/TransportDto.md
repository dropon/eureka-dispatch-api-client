# TransportDto

## Properties

| Name                             | Type                                                                                                     | Description                                                                                                                                                      | Notes      |
| -------------------------------- | -------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Uid**                          | Pointer to **string**                                                                                    | Transport&#39;s internal unique identifier                                                                                                                       | [optional] |
| **TransportId**                  | Pointer to **int32**                                                                                     | Transport&#39;s internal unique identifier                                                                                                                       | [optional] |
| **Index**                        | Pointer to **int32**                                                                                     | The index of the transport in the mission or quotation                                                                                                           | [optional] |
| **MissionNumber**                | Pointer to **int32**                                                                                     | Mission&#39;s public identifier                                                                                                                                  | [optional] |
| **MissionUid**                   | Pointer to **string**                                                                                    | Mission&#39;s internal unique identifier                                                                                                                         | [optional] |
| **MissionTrackId**               | Pointer to **string**                                                                                    | Mission&#39;s tracking number                                                                                                                                    | [optional] |
| **QuotationNumber**              | Pointer to **int32**                                                                                     | Quotation&#39;s public identifier                                                                                                                                | [optional] |
| **QuotationUid**                 | Pointer to **string**                                                                                    | Quotation&#39;s internal unique identifier                                                                                                                       | [optional] |
| **QuotationTrackId**             | Pointer to **string**                                                                                    | Quotation&#39;s tracking number                                                                                                                                  | [optional] |
| **IsQuotationArchived**          | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **IsQuotationFinalized**         | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **IsQuotationSubjectToApproval** | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **IsMissionReadyToBill**         | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **SourceQuotationNumber**        | Pointer to **int32**                                                                                     | Source quotation number, for a transport included in a mission created from a quotation.                                                                         | [optional] |
| **RoundId**                      | Pointer to **int32**                                                                                     |                                                                                                                                                                  | [optional] |
| **CustomerCode**                 | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **AgencyCode**                   | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **OrdererCode**                  | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **OrdererName**                  | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **ContractorAgentId**            | Pointer to **int32**                                                                                     |                                                                                                                                                                  | [optional] |
| **ContractorAgentName**          | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **HandlerId**                    | Pointer to **int32**                                                                                     |                                                                                                                                                                  | [optional] |
| **VehicleCode**                  | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **TrailerCode**                  | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **Status**                       | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **SubStateCode**                 | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **CustomerCallDateTime**         | Pointer to **Time**                                                                                      | The date and time at which the customer has contacted the transport company to request the transport. By default, it corresponds to the mission creation date.   | [optional] |
| **DriverId**                     | Pointer to **int32**                                                                                     |                                                                                                                                                                  | [optional] |
| **Sign**                         | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **SecretCode**                   | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **Notes**                        | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **IsRoundTrip**                  | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **ServiceCode**                  | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **SubServices**                  | Pointer to [**CappedCollectionDtoTransportSubServiceDto**](CappedCollectionDtoTransportSubServiceDto.md) |                                                                                                                                                                  | [optional] |
| **PickupStep**                   | Pointer to [**TransportPickupStepDto**](TransportPickupStepDto.md)                                       |                                                                                                                                                                  | [optional] |
| **DeliveryStep**                 | Pointer to [**TransportDeliveryStepDto**](TransportDeliveryStepDto.md)                                   |                                                                                                                                                                  | [optional] |
| **Packing**                      | Pointer to [**TransportPackingDto**](TransportPackingDto.md)                                             |                                                                                                                                                                  | [optional] |
| **DistanceKm**                   | Pointer to **float64**                                                                                   | Distance of the transport in kilometers.                                                                                                                         | [optional] |
| **TotalDurationMinutes**         | Pointer to **int32**                                                                                     | Total duration of the transport in minutes                                                                                                                       | [optional] |
| **IsStrategic**                  | Pointer to **bool**                                                                                      | Indicates if the transport has a strategic importance                                                                                                            | [optional] |
| **IsShared**                     | Pointer to **bool**                                                                                      | Indicates if the transport is &#39;shared&#39; between agencies. If true, the transport will be visible on Dispatch planning via the &#39;supervising&#39; mode. | [optional] |
| **ResponsibleOperatorId**        | Pointer to **int32**                                                                                     | The id of a dispatch user that is responsible for following the execution of the transport.                                                                      | [optional] |
| **GasEmission**                  | Pointer to **float64**                                                                                   | Gas emission (CO2) in Kilograms.                                                                                                                                 | [optional] |
| **IsGasEmissionForced**          | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **SellCurrencyCode**             | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **SellPrice**                    | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **ForcedSellPrice**              | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **SellFuelSurchargePrice**       | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **BuyCurrencyCode**              | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **BuyPrice**                     | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **ForcedBuyPrice**               | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **BuyFuelSurchargePrice**        | Pointer to **float64**                                                                                   | The following permission(s) are required to access this property: See prices.                                                                                    | [optional] |
| **ReferenceCurrencyCode**        | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **SellPricingPathId**            | Pointer to **int32**                                                                                     |                                                                                                                                                                  | [optional] |
| **IsFragmented**                 | Pointer to **bool**                                                                                      | Indicates if this transport is fragmented: in this case, this transport is a parent transport and has child fragments.                                           | [optional] |
| **ParentTransportId**            | Pointer to **int32**                                                                                     | Identifier of the parent transport for a child fragment, when the mission or quotation is fragmented                                                             | [optional] |
| **MissionIsInSequence**          | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **MissionIsUniqueOrder**         | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **QuotationIsInSequence**        | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **QuotationIsUniqueOrder**       | Pointer to **bool**                                                                                      |                                                                                                                                                                  | [optional] |
| **Reference1**                   | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **Reference2**                   | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **Reference3**                   | Pointer to **string**                                                                                    |                                                                                                                                                                  | [optional] |
| **Comment**                      | Pointer to [**TransportCommentDto**](TransportCommentDto.md)                                             |                                                                                                                                                                  | [optional] |
| **CashOnDelivery**               | Pointer to [**TransportCashOnDeliveryDto**](TransportCashOnDeliveryDto.md)                               |                                                                                                                                                                  | [optional] |
| **Included**                     | Pointer to [**TransportIncludedDto**](TransportIncludedDto.md)                                           |                                                                                                                                                                  | [optional] |

## Methods

### NewTransportDto

`func NewTransportDto() *TransportDto`

NewTransportDto instantiates a new TransportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportDtoWithDefaults

`func NewTransportDtoWithDefaults() *TransportDto`

NewTransportDtoWithDefaults instantiates a new TransportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *TransportDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *TransportDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *TransportDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *TransportDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetTransportId

`func (o *TransportDto) GetTransportId() int32`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *TransportDto) GetTransportIdOk() (*int32, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *TransportDto) SetTransportId(v int32)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *TransportDto) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.

### GetIndex

`func (o *TransportDto) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TransportDto) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TransportDto) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *TransportDto) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMissionNumber

`func (o *TransportDto) GetMissionNumber() int32`

GetMissionNumber returns the MissionNumber field if non-nil, zero value otherwise.

### GetMissionNumberOk

`func (o *TransportDto) GetMissionNumberOk() (*int32, bool)`

GetMissionNumberOk returns a tuple with the MissionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionNumber

`func (o *TransportDto) SetMissionNumber(v int32)`

SetMissionNumber sets MissionNumber field to given value.

### HasMissionNumber

`func (o *TransportDto) HasMissionNumber() bool`

HasMissionNumber returns a boolean if a field has been set.

### GetMissionUid

`func (o *TransportDto) GetMissionUid() string`

GetMissionUid returns the MissionUid field if non-nil, zero value otherwise.

### GetMissionUidOk

`func (o *TransportDto) GetMissionUidOk() (*string, bool)`

GetMissionUidOk returns a tuple with the MissionUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionUid

`func (o *TransportDto) SetMissionUid(v string)`

SetMissionUid sets MissionUid field to given value.

### HasMissionUid

`func (o *TransportDto) HasMissionUid() bool`

HasMissionUid returns a boolean if a field has been set.

### GetMissionTrackId

`func (o *TransportDto) GetMissionTrackId() string`

GetMissionTrackId returns the MissionTrackId field if non-nil, zero value otherwise.

### GetMissionTrackIdOk

`func (o *TransportDto) GetMissionTrackIdOk() (*string, bool)`

GetMissionTrackIdOk returns a tuple with the MissionTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionTrackId

`func (o *TransportDto) SetMissionTrackId(v string)`

SetMissionTrackId sets MissionTrackId field to given value.

### HasMissionTrackId

`func (o *TransportDto) HasMissionTrackId() bool`

HasMissionTrackId returns a boolean if a field has been set.

### GetQuotationNumber

`func (o *TransportDto) GetQuotationNumber() int32`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *TransportDto) GetQuotationNumberOk() (*int32, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *TransportDto) SetQuotationNumber(v int32)`

SetQuotationNumber sets QuotationNumber field to given value.

### HasQuotationNumber

`func (o *TransportDto) HasQuotationNumber() bool`

HasQuotationNumber returns a boolean if a field has been set.

### GetQuotationUid

`func (o *TransportDto) GetQuotationUid() string`

GetQuotationUid returns the QuotationUid field if non-nil, zero value otherwise.

### GetQuotationUidOk

`func (o *TransportDto) GetQuotationUidOk() (*string, bool)`

GetQuotationUidOk returns a tuple with the QuotationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationUid

`func (o *TransportDto) SetQuotationUid(v string)`

SetQuotationUid sets QuotationUid field to given value.

### HasQuotationUid

`func (o *TransportDto) HasQuotationUid() bool`

HasQuotationUid returns a boolean if a field has been set.

### GetQuotationTrackId

`func (o *TransportDto) GetQuotationTrackId() string`

GetQuotationTrackId returns the QuotationTrackId field if non-nil, zero value otherwise.

### GetQuotationTrackIdOk

`func (o *TransportDto) GetQuotationTrackIdOk() (*string, bool)`

GetQuotationTrackIdOk returns a tuple with the QuotationTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationTrackId

`func (o *TransportDto) SetQuotationTrackId(v string)`

SetQuotationTrackId sets QuotationTrackId field to given value.

### HasQuotationTrackId

`func (o *TransportDto) HasQuotationTrackId() bool`

HasQuotationTrackId returns a boolean if a field has been set.

### GetIsQuotationArchived

`func (o *TransportDto) GetIsQuotationArchived() bool`

GetIsQuotationArchived returns the IsQuotationArchived field if non-nil, zero value otherwise.

### GetIsQuotationArchivedOk

`func (o *TransportDto) GetIsQuotationArchivedOk() (*bool, bool)`

GetIsQuotationArchivedOk returns a tuple with the IsQuotationArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuotationArchived

`func (o *TransportDto) SetIsQuotationArchived(v bool)`

SetIsQuotationArchived sets IsQuotationArchived field to given value.

### HasIsQuotationArchived

`func (o *TransportDto) HasIsQuotationArchived() bool`

HasIsQuotationArchived returns a boolean if a field has been set.

### GetIsQuotationFinalized

`func (o *TransportDto) GetIsQuotationFinalized() bool`

GetIsQuotationFinalized returns the IsQuotationFinalized field if non-nil, zero value otherwise.

### GetIsQuotationFinalizedOk

`func (o *TransportDto) GetIsQuotationFinalizedOk() (*bool, bool)`

GetIsQuotationFinalizedOk returns a tuple with the IsQuotationFinalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuotationFinalized

`func (o *TransportDto) SetIsQuotationFinalized(v bool)`

SetIsQuotationFinalized sets IsQuotationFinalized field to given value.

### HasIsQuotationFinalized

`func (o *TransportDto) HasIsQuotationFinalized() bool`

HasIsQuotationFinalized returns a boolean if a field has been set.

### GetIsQuotationSubjectToApproval

`func (o *TransportDto) GetIsQuotationSubjectToApproval() bool`

GetIsQuotationSubjectToApproval returns the IsQuotationSubjectToApproval field if non-nil, zero value otherwise.

### GetIsQuotationSubjectToApprovalOk

`func (o *TransportDto) GetIsQuotationSubjectToApprovalOk() (*bool, bool)`

GetIsQuotationSubjectToApprovalOk returns a tuple with the IsQuotationSubjectToApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuotationSubjectToApproval

`func (o *TransportDto) SetIsQuotationSubjectToApproval(v bool)`

SetIsQuotationSubjectToApproval sets IsQuotationSubjectToApproval field to given value.

### HasIsQuotationSubjectToApproval

`func (o *TransportDto) HasIsQuotationSubjectToApproval() bool`

HasIsQuotationSubjectToApproval returns a boolean if a field has been set.

### GetIsMissionReadyToBill

`func (o *TransportDto) GetIsMissionReadyToBill() bool`

GetIsMissionReadyToBill returns the IsMissionReadyToBill field if non-nil, zero value otherwise.

### GetIsMissionReadyToBillOk

`func (o *TransportDto) GetIsMissionReadyToBillOk() (*bool, bool)`

GetIsMissionReadyToBillOk returns a tuple with the IsMissionReadyToBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMissionReadyToBill

`func (o *TransportDto) SetIsMissionReadyToBill(v bool)`

SetIsMissionReadyToBill sets IsMissionReadyToBill field to given value.

### HasIsMissionReadyToBill

`func (o *TransportDto) HasIsMissionReadyToBill() bool`

HasIsMissionReadyToBill returns a boolean if a field has been set.

### GetSourceQuotationNumber

`func (o *TransportDto) GetSourceQuotationNumber() int32`

GetSourceQuotationNumber returns the SourceQuotationNumber field if non-nil, zero value otherwise.

### GetSourceQuotationNumberOk

`func (o *TransportDto) GetSourceQuotationNumberOk() (*int32, bool)`

GetSourceQuotationNumberOk returns a tuple with the SourceQuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuotationNumber

`func (o *TransportDto) SetSourceQuotationNumber(v int32)`

SetSourceQuotationNumber sets SourceQuotationNumber field to given value.

### HasSourceQuotationNumber

`func (o *TransportDto) HasSourceQuotationNumber() bool`

HasSourceQuotationNumber returns a boolean if a field has been set.

### GetRoundId

`func (o *TransportDto) GetRoundId() int32`

GetRoundId returns the RoundId field if non-nil, zero value otherwise.

### GetRoundIdOk

`func (o *TransportDto) GetRoundIdOk() (*int32, bool)`

GetRoundIdOk returns a tuple with the RoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundId

`func (o *TransportDto) SetRoundId(v int32)`

SetRoundId sets RoundId field to given value.

### HasRoundId

`func (o *TransportDto) HasRoundId() bool`

HasRoundId returns a boolean if a field has been set.

### GetCustomerCode

`func (o *TransportDto) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *TransportDto) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *TransportDto) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.

### HasCustomerCode

`func (o *TransportDto) HasCustomerCode() bool`

HasCustomerCode returns a boolean if a field has been set.

### GetAgencyCode

`func (o *TransportDto) GetAgencyCode() string`

GetAgencyCode returns the AgencyCode field if non-nil, zero value otherwise.

### GetAgencyCodeOk

`func (o *TransportDto) GetAgencyCodeOk() (*string, bool)`

GetAgencyCodeOk returns a tuple with the AgencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyCode

`func (o *TransportDto) SetAgencyCode(v string)`

SetAgencyCode sets AgencyCode field to given value.

### HasAgencyCode

`func (o *TransportDto) HasAgencyCode() bool`

HasAgencyCode returns a boolean if a field has been set.

### GetOrdererCode

`func (o *TransportDto) GetOrdererCode() string`

GetOrdererCode returns the OrdererCode field if non-nil, zero value otherwise.

### GetOrdererCodeOk

`func (o *TransportDto) GetOrdererCodeOk() (*string, bool)`

GetOrdererCodeOk returns a tuple with the OrdererCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererCode

`func (o *TransportDto) SetOrdererCode(v string)`

SetOrdererCode sets OrdererCode field to given value.

### HasOrdererCode

`func (o *TransportDto) HasOrdererCode() bool`

HasOrdererCode returns a boolean if a field has been set.

### GetOrdererName

`func (o *TransportDto) GetOrdererName() string`

GetOrdererName returns the OrdererName field if non-nil, zero value otherwise.

### GetOrdererNameOk

`func (o *TransportDto) GetOrdererNameOk() (*string, bool)`

GetOrdererNameOk returns a tuple with the OrdererName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdererName

`func (o *TransportDto) SetOrdererName(v string)`

SetOrdererName sets OrdererName field to given value.

### HasOrdererName

`func (o *TransportDto) HasOrdererName() bool`

HasOrdererName returns a boolean if a field has been set.

### GetContractorAgentId

`func (o *TransportDto) GetContractorAgentId() int32`

GetContractorAgentId returns the ContractorAgentId field if non-nil, zero value otherwise.

### GetContractorAgentIdOk

`func (o *TransportDto) GetContractorAgentIdOk() (*int32, bool)`

GetContractorAgentIdOk returns a tuple with the ContractorAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorAgentId

`func (o *TransportDto) SetContractorAgentId(v int32)`

SetContractorAgentId sets ContractorAgentId field to given value.

### HasContractorAgentId

`func (o *TransportDto) HasContractorAgentId() bool`

HasContractorAgentId returns a boolean if a field has been set.

### GetContractorAgentName

`func (o *TransportDto) GetContractorAgentName() string`

GetContractorAgentName returns the ContractorAgentName field if non-nil, zero value otherwise.

### GetContractorAgentNameOk

`func (o *TransportDto) GetContractorAgentNameOk() (*string, bool)`

GetContractorAgentNameOk returns a tuple with the ContractorAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorAgentName

`func (o *TransportDto) SetContractorAgentName(v string)`

SetContractorAgentName sets ContractorAgentName field to given value.

### HasContractorAgentName

`func (o *TransportDto) HasContractorAgentName() bool`

HasContractorAgentName returns a boolean if a field has been set.

### GetHandlerId

`func (o *TransportDto) GetHandlerId() int32`

GetHandlerId returns the HandlerId field if non-nil, zero value otherwise.

### GetHandlerIdOk

`func (o *TransportDto) GetHandlerIdOk() (*int32, bool)`

GetHandlerIdOk returns a tuple with the HandlerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerId

`func (o *TransportDto) SetHandlerId(v int32)`

SetHandlerId sets HandlerId field to given value.

### HasHandlerId

`func (o *TransportDto) HasHandlerId() bool`

HasHandlerId returns a boolean if a field has been set.

### GetVehicleCode

`func (o *TransportDto) GetVehicleCode() string`

GetVehicleCode returns the VehicleCode field if non-nil, zero value otherwise.

### GetVehicleCodeOk

`func (o *TransportDto) GetVehicleCodeOk() (*string, bool)`

GetVehicleCodeOk returns a tuple with the VehicleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleCode

`func (o *TransportDto) SetVehicleCode(v string)`

SetVehicleCode sets VehicleCode field to given value.

### HasVehicleCode

`func (o *TransportDto) HasVehicleCode() bool`

HasVehicleCode returns a boolean if a field has been set.

### GetTrailerCode

`func (o *TransportDto) GetTrailerCode() string`

GetTrailerCode returns the TrailerCode field if non-nil, zero value otherwise.

### GetTrailerCodeOk

`func (o *TransportDto) GetTrailerCodeOk() (*string, bool)`

GetTrailerCodeOk returns a tuple with the TrailerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerCode

`func (o *TransportDto) SetTrailerCode(v string)`

SetTrailerCode sets TrailerCode field to given value.

### HasTrailerCode

`func (o *TransportDto) HasTrailerCode() bool`

HasTrailerCode returns a boolean if a field has been set.

### GetStatus

`func (o *TransportDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransportDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransportDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransportDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStateCode

`func (o *TransportDto) GetSubStateCode() string`

GetSubStateCode returns the SubStateCode field if non-nil, zero value otherwise.

### GetSubStateCodeOk

`func (o *TransportDto) GetSubStateCodeOk() (*string, bool)`

GetSubStateCodeOk returns a tuple with the SubStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStateCode

`func (o *TransportDto) SetSubStateCode(v string)`

SetSubStateCode sets SubStateCode field to given value.

### HasSubStateCode

`func (o *TransportDto) HasSubStateCode() bool`

HasSubStateCode returns a boolean if a field has been set.

### GetCustomerCallDateTime

`func (o *TransportDto) GetCustomerCallDateTime() Time`

GetCustomerCallDateTime returns the CustomerCallDateTime field if non-nil, zero value otherwise.

### GetCustomerCallDateTimeOk

`func (o *TransportDto) GetCustomerCallDateTimeOk() (*Time, bool)`

GetCustomerCallDateTimeOk returns a tuple with the CustomerCallDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCallDateTime

`func (o *TransportDto) SetCustomerCallDateTime(v Time)`

SetCustomerCallDateTime sets CustomerCallDateTime field to given value.

### HasCustomerCallDateTime

`func (o *TransportDto) HasCustomerCallDateTime() bool`

HasCustomerCallDateTime returns a boolean if a field has been set.

### GetDriverId

`func (o *TransportDto) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *TransportDto) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *TransportDto) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *TransportDto) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetSign

`func (o *TransportDto) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *TransportDto) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *TransportDto) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *TransportDto) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetSecretCode

`func (o *TransportDto) GetSecretCode() string`

GetSecretCode returns the SecretCode field if non-nil, zero value otherwise.

### GetSecretCodeOk

`func (o *TransportDto) GetSecretCodeOk() (*string, bool)`

GetSecretCodeOk returns a tuple with the SecretCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretCode

`func (o *TransportDto) SetSecretCode(v string)`

SetSecretCode sets SecretCode field to given value.

### HasSecretCode

`func (o *TransportDto) HasSecretCode() bool`

HasSecretCode returns a boolean if a field has been set.

### GetNotes

`func (o *TransportDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransportDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransportDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransportDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsRoundTrip

`func (o *TransportDto) GetIsRoundTrip() bool`

GetIsRoundTrip returns the IsRoundTrip field if non-nil, zero value otherwise.

### GetIsRoundTripOk

`func (o *TransportDto) GetIsRoundTripOk() (*bool, bool)`

GetIsRoundTripOk returns a tuple with the IsRoundTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoundTrip

`func (o *TransportDto) SetIsRoundTrip(v bool)`

SetIsRoundTrip sets IsRoundTrip field to given value.

### HasIsRoundTrip

`func (o *TransportDto) HasIsRoundTrip() bool`

HasIsRoundTrip returns a boolean if a field has been set.

### GetServiceCode

`func (o *TransportDto) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *TransportDto) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *TransportDto) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *TransportDto) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetSubServices

`func (o *TransportDto) GetSubServices() CappedCollectionDtoTransportSubServiceDto`

GetSubServices returns the SubServices field if non-nil, zero value otherwise.

### GetSubServicesOk

`func (o *TransportDto) GetSubServicesOk() (*CappedCollectionDtoTransportSubServiceDto, bool)`

GetSubServicesOk returns a tuple with the SubServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubServices

`func (o *TransportDto) SetSubServices(v CappedCollectionDtoTransportSubServiceDto)`

SetSubServices sets SubServices field to given value.

### HasSubServices

`func (o *TransportDto) HasSubServices() bool`

HasSubServices returns a boolean if a field has been set.

### GetPickupStep

`func (o *TransportDto) GetPickupStep() TransportPickupStepDto`

GetPickupStep returns the PickupStep field if non-nil, zero value otherwise.

### GetPickupStepOk

`func (o *TransportDto) GetPickupStepOk() (*TransportPickupStepDto, bool)`

GetPickupStepOk returns a tuple with the PickupStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStep

`func (o *TransportDto) SetPickupStep(v TransportPickupStepDto)`

SetPickupStep sets PickupStep field to given value.

### HasPickupStep

`func (o *TransportDto) HasPickupStep() bool`

HasPickupStep returns a boolean if a field has been set.

### GetDeliveryStep

`func (o *TransportDto) GetDeliveryStep() TransportDeliveryStepDto`

GetDeliveryStep returns the DeliveryStep field if non-nil, zero value otherwise.

### GetDeliveryStepOk

`func (o *TransportDto) GetDeliveryStepOk() (*TransportDeliveryStepDto, bool)`

GetDeliveryStepOk returns a tuple with the DeliveryStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStep

`func (o *TransportDto) SetDeliveryStep(v TransportDeliveryStepDto)`

SetDeliveryStep sets DeliveryStep field to given value.

### HasDeliveryStep

`func (o *TransportDto) HasDeliveryStep() bool`

HasDeliveryStep returns a boolean if a field has been set.

### GetPacking

`func (o *TransportDto) GetPacking() TransportPackingDto`

GetPacking returns the Packing field if non-nil, zero value otherwise.

### GetPackingOk

`func (o *TransportDto) GetPackingOk() (*TransportPackingDto, bool)`

GetPackingOk returns a tuple with the Packing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacking

`func (o *TransportDto) SetPacking(v TransportPackingDto)`

SetPacking sets Packing field to given value.

### HasPacking

`func (o *TransportDto) HasPacking() bool`

HasPacking returns a boolean if a field has been set.

### GetDistanceKm

`func (o *TransportDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *TransportDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *TransportDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *TransportDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### GetTotalDurationMinutes

`func (o *TransportDto) GetTotalDurationMinutes() int32`

GetTotalDurationMinutes returns the TotalDurationMinutes field if non-nil, zero value otherwise.

### GetTotalDurationMinutesOk

`func (o *TransportDto) GetTotalDurationMinutesOk() (*int32, bool)`

GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMinutes

`func (o *TransportDto) SetTotalDurationMinutes(v int32)`

SetTotalDurationMinutes sets TotalDurationMinutes field to given value.

### HasTotalDurationMinutes

`func (o *TransportDto) HasTotalDurationMinutes() bool`

HasTotalDurationMinutes returns a boolean if a field has been set.

### GetIsStrategic

`func (o *TransportDto) GetIsStrategic() bool`

GetIsStrategic returns the IsStrategic field if non-nil, zero value otherwise.

### GetIsStrategicOk

`func (o *TransportDto) GetIsStrategicOk() (*bool, bool)`

GetIsStrategicOk returns a tuple with the IsStrategic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStrategic

`func (o *TransportDto) SetIsStrategic(v bool)`

SetIsStrategic sets IsStrategic field to given value.

### HasIsStrategic

`func (o *TransportDto) HasIsStrategic() bool`

HasIsStrategic returns a boolean if a field has been set.

### GetIsShared

`func (o *TransportDto) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *TransportDto) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *TransportDto) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *TransportDto) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetResponsibleOperatorId

`func (o *TransportDto) GetResponsibleOperatorId() int32`

GetResponsibleOperatorId returns the ResponsibleOperatorId field if non-nil, zero value otherwise.

### GetResponsibleOperatorIdOk

`func (o *TransportDto) GetResponsibleOperatorIdOk() (*int32, bool)`

GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleOperatorId

`func (o *TransportDto) SetResponsibleOperatorId(v int32)`

SetResponsibleOperatorId sets ResponsibleOperatorId field to given value.

### HasResponsibleOperatorId

`func (o *TransportDto) HasResponsibleOperatorId() bool`

HasResponsibleOperatorId returns a boolean if a field has been set.

### GetGasEmission

`func (o *TransportDto) GetGasEmission() float64`

GetGasEmission returns the GasEmission field if non-nil, zero value otherwise.

### GetGasEmissionOk

`func (o *TransportDto) GetGasEmissionOk() (*float64, bool)`

GetGasEmissionOk returns a tuple with the GasEmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasEmission

`func (o *TransportDto) SetGasEmission(v float64)`

SetGasEmission sets GasEmission field to given value.

### HasGasEmission

`func (o *TransportDto) HasGasEmission() bool`

HasGasEmission returns a boolean if a field has been set.

### GetIsGasEmissionForced

`func (o *TransportDto) GetIsGasEmissionForced() bool`

GetIsGasEmissionForced returns the IsGasEmissionForced field if non-nil, zero value otherwise.

### GetIsGasEmissionForcedOk

`func (o *TransportDto) GetIsGasEmissionForcedOk() (*bool, bool)`

GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGasEmissionForced

`func (o *TransportDto) SetIsGasEmissionForced(v bool)`

SetIsGasEmissionForced sets IsGasEmissionForced field to given value.

### HasIsGasEmissionForced

`func (o *TransportDto) HasIsGasEmissionForced() bool`

HasIsGasEmissionForced returns a boolean if a field has been set.

### GetSellCurrencyCode

`func (o *TransportDto) GetSellCurrencyCode() string`

GetSellCurrencyCode returns the SellCurrencyCode field if non-nil, zero value otherwise.

### GetSellCurrencyCodeOk

`func (o *TransportDto) GetSellCurrencyCodeOk() (*string, bool)`

GetSellCurrencyCodeOk returns a tuple with the SellCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellCurrencyCode

`func (o *TransportDto) SetSellCurrencyCode(v string)`

SetSellCurrencyCode sets SellCurrencyCode field to given value.

### HasSellCurrencyCode

`func (o *TransportDto) HasSellCurrencyCode() bool`

HasSellCurrencyCode returns a boolean if a field has been set.

### GetSellPrice

`func (o *TransportDto) GetSellPrice() float64`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *TransportDto) GetSellPriceOk() (*float64, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *TransportDto) SetSellPrice(v float64)`

SetSellPrice sets SellPrice field to given value.

### HasSellPrice

`func (o *TransportDto) HasSellPrice() bool`

HasSellPrice returns a boolean if a field has been set.

### GetForcedSellPrice

`func (o *TransportDto) GetForcedSellPrice() float64`

GetForcedSellPrice returns the ForcedSellPrice field if non-nil, zero value otherwise.

### GetForcedSellPriceOk

`func (o *TransportDto) GetForcedSellPriceOk() (*float64, bool)`

GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSellPrice

`func (o *TransportDto) SetForcedSellPrice(v float64)`

SetForcedSellPrice sets ForcedSellPrice field to given value.

### HasForcedSellPrice

`func (o *TransportDto) HasForcedSellPrice() bool`

HasForcedSellPrice returns a boolean if a field has been set.

### GetSellFuelSurchargePrice

`func (o *TransportDto) GetSellFuelSurchargePrice() float64`

GetSellFuelSurchargePrice returns the SellFuelSurchargePrice field if non-nil, zero value otherwise.

### GetSellFuelSurchargePriceOk

`func (o *TransportDto) GetSellFuelSurchargePriceOk() (*float64, bool)`

GetSellFuelSurchargePriceOk returns a tuple with the SellFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFuelSurchargePrice

`func (o *TransportDto) SetSellFuelSurchargePrice(v float64)`

SetSellFuelSurchargePrice sets SellFuelSurchargePrice field to given value.

### HasSellFuelSurchargePrice

`func (o *TransportDto) HasSellFuelSurchargePrice() bool`

HasSellFuelSurchargePrice returns a boolean if a field has been set.

### GetBuyCurrencyCode

`func (o *TransportDto) GetBuyCurrencyCode() string`

GetBuyCurrencyCode returns the BuyCurrencyCode field if non-nil, zero value otherwise.

### GetBuyCurrencyCodeOk

`func (o *TransportDto) GetBuyCurrencyCodeOk() (*string, bool)`

GetBuyCurrencyCodeOk returns a tuple with the BuyCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyCurrencyCode

`func (o *TransportDto) SetBuyCurrencyCode(v string)`

SetBuyCurrencyCode sets BuyCurrencyCode field to given value.

### HasBuyCurrencyCode

`func (o *TransportDto) HasBuyCurrencyCode() bool`

HasBuyCurrencyCode returns a boolean if a field has been set.

### GetBuyPrice

`func (o *TransportDto) GetBuyPrice() float64`

GetBuyPrice returns the BuyPrice field if non-nil, zero value otherwise.

### GetBuyPriceOk

`func (o *TransportDto) GetBuyPriceOk() (*float64, bool)`

GetBuyPriceOk returns a tuple with the BuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyPrice

`func (o *TransportDto) SetBuyPrice(v float64)`

SetBuyPrice sets BuyPrice field to given value.

### HasBuyPrice

`func (o *TransportDto) HasBuyPrice() bool`

HasBuyPrice returns a boolean if a field has been set.

### GetForcedBuyPrice

`func (o *TransportDto) GetForcedBuyPrice() float64`

GetForcedBuyPrice returns the ForcedBuyPrice field if non-nil, zero value otherwise.

### GetForcedBuyPriceOk

`func (o *TransportDto) GetForcedBuyPriceOk() (*float64, bool)`

GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedBuyPrice

`func (o *TransportDto) SetForcedBuyPrice(v float64)`

SetForcedBuyPrice sets ForcedBuyPrice field to given value.

### HasForcedBuyPrice

`func (o *TransportDto) HasForcedBuyPrice() bool`

HasForcedBuyPrice returns a boolean if a field has been set.

### GetBuyFuelSurchargePrice

`func (o *TransportDto) GetBuyFuelSurchargePrice() float64`

GetBuyFuelSurchargePrice returns the BuyFuelSurchargePrice field if non-nil, zero value otherwise.

### GetBuyFuelSurchargePriceOk

`func (o *TransportDto) GetBuyFuelSurchargePriceOk() (*float64, bool)`

GetBuyFuelSurchargePriceOk returns a tuple with the BuyFuelSurchargePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFuelSurchargePrice

`func (o *TransportDto) SetBuyFuelSurchargePrice(v float64)`

SetBuyFuelSurchargePrice sets BuyFuelSurchargePrice field to given value.

### HasBuyFuelSurchargePrice

`func (o *TransportDto) HasBuyFuelSurchargePrice() bool`

HasBuyFuelSurchargePrice returns a boolean if a field has been set.

### GetReferenceCurrencyCode

`func (o *TransportDto) GetReferenceCurrencyCode() string`

GetReferenceCurrencyCode returns the ReferenceCurrencyCode field if non-nil, zero value otherwise.

### GetReferenceCurrencyCodeOk

`func (o *TransportDto) GetReferenceCurrencyCodeOk() (*string, bool)`

GetReferenceCurrencyCodeOk returns a tuple with the ReferenceCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCurrencyCode

`func (o *TransportDto) SetReferenceCurrencyCode(v string)`

SetReferenceCurrencyCode sets ReferenceCurrencyCode field to given value.

### HasReferenceCurrencyCode

`func (o *TransportDto) HasReferenceCurrencyCode() bool`

HasReferenceCurrencyCode returns a boolean if a field has been set.

### GetSellPricingPathId

`func (o *TransportDto) GetSellPricingPathId() int32`

GetSellPricingPathId returns the SellPricingPathId field if non-nil, zero value otherwise.

### GetSellPricingPathIdOk

`func (o *TransportDto) GetSellPricingPathIdOk() (*int32, bool)`

GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPricingPathId

`func (o *TransportDto) SetSellPricingPathId(v int32)`

SetSellPricingPathId sets SellPricingPathId field to given value.

### HasSellPricingPathId

`func (o *TransportDto) HasSellPricingPathId() bool`

HasSellPricingPathId returns a boolean if a field has been set.

### GetIsFragmented

`func (o *TransportDto) GetIsFragmented() bool`

GetIsFragmented returns the IsFragmented field if non-nil, zero value otherwise.

### GetIsFragmentedOk

`func (o *TransportDto) GetIsFragmentedOk() (*bool, bool)`

GetIsFragmentedOk returns a tuple with the IsFragmented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFragmented

`func (o *TransportDto) SetIsFragmented(v bool)`

SetIsFragmented sets IsFragmented field to given value.

### HasIsFragmented

`func (o *TransportDto) HasIsFragmented() bool`

HasIsFragmented returns a boolean if a field has been set.

### GetParentTransportId

`func (o *TransportDto) GetParentTransportId() int32`

GetParentTransportId returns the ParentTransportId field if non-nil, zero value otherwise.

### GetParentTransportIdOk

`func (o *TransportDto) GetParentTransportIdOk() (*int32, bool)`

GetParentTransportIdOk returns a tuple with the ParentTransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransportId

`func (o *TransportDto) SetParentTransportId(v int32)`

SetParentTransportId sets ParentTransportId field to given value.

### HasParentTransportId

`func (o *TransportDto) HasParentTransportId() bool`

HasParentTransportId returns a boolean if a field has been set.

### GetMissionIsInSequence

`func (o *TransportDto) GetMissionIsInSequence() bool`

GetMissionIsInSequence returns the MissionIsInSequence field if non-nil, zero value otherwise.

### GetMissionIsInSequenceOk

`func (o *TransportDto) GetMissionIsInSequenceOk() (*bool, bool)`

GetMissionIsInSequenceOk returns a tuple with the MissionIsInSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionIsInSequence

`func (o *TransportDto) SetMissionIsInSequence(v bool)`

SetMissionIsInSequence sets MissionIsInSequence field to given value.

### HasMissionIsInSequence

`func (o *TransportDto) HasMissionIsInSequence() bool`

HasMissionIsInSequence returns a boolean if a field has been set.

### GetMissionIsUniqueOrder

`func (o *TransportDto) GetMissionIsUniqueOrder() bool`

GetMissionIsUniqueOrder returns the MissionIsUniqueOrder field if non-nil, zero value otherwise.

### GetMissionIsUniqueOrderOk

`func (o *TransportDto) GetMissionIsUniqueOrderOk() (*bool, bool)`

GetMissionIsUniqueOrderOk returns a tuple with the MissionIsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissionIsUniqueOrder

`func (o *TransportDto) SetMissionIsUniqueOrder(v bool)`

SetMissionIsUniqueOrder sets MissionIsUniqueOrder field to given value.

### HasMissionIsUniqueOrder

`func (o *TransportDto) HasMissionIsUniqueOrder() bool`

HasMissionIsUniqueOrder returns a boolean if a field has been set.

### GetQuotationIsInSequence

`func (o *TransportDto) GetQuotationIsInSequence() bool`

GetQuotationIsInSequence returns the QuotationIsInSequence field if non-nil, zero value otherwise.

### GetQuotationIsInSequenceOk

`func (o *TransportDto) GetQuotationIsInSequenceOk() (*bool, bool)`

GetQuotationIsInSequenceOk returns a tuple with the QuotationIsInSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationIsInSequence

`func (o *TransportDto) SetQuotationIsInSequence(v bool)`

SetQuotationIsInSequence sets QuotationIsInSequence field to given value.

### HasQuotationIsInSequence

`func (o *TransportDto) HasQuotationIsInSequence() bool`

HasQuotationIsInSequence returns a boolean if a field has been set.

### GetQuotationIsUniqueOrder

`func (o *TransportDto) GetQuotationIsUniqueOrder() bool`

GetQuotationIsUniqueOrder returns the QuotationIsUniqueOrder field if non-nil, zero value otherwise.

### GetQuotationIsUniqueOrderOk

`func (o *TransportDto) GetQuotationIsUniqueOrderOk() (*bool, bool)`

GetQuotationIsUniqueOrderOk returns a tuple with the QuotationIsUniqueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationIsUniqueOrder

`func (o *TransportDto) SetQuotationIsUniqueOrder(v bool)`

SetQuotationIsUniqueOrder sets QuotationIsUniqueOrder field to given value.

### HasQuotationIsUniqueOrder

`func (o *TransportDto) HasQuotationIsUniqueOrder() bool`

HasQuotationIsUniqueOrder returns a boolean if a field has been set.

### GetReference1

`func (o *TransportDto) GetReference1() string`

GetReference1 returns the Reference1 field if non-nil, zero value otherwise.

### GetReference1Ok

`func (o *TransportDto) GetReference1Ok() (*string, bool)`

GetReference1Ok returns a tuple with the Reference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference1

`func (o *TransportDto) SetReference1(v string)`

SetReference1 sets Reference1 field to given value.

### HasReference1

`func (o *TransportDto) HasReference1() bool`

HasReference1 returns a boolean if a field has been set.

### GetReference2

`func (o *TransportDto) GetReference2() string`

GetReference2 returns the Reference2 field if non-nil, zero value otherwise.

### GetReference2Ok

`func (o *TransportDto) GetReference2Ok() (*string, bool)`

GetReference2Ok returns a tuple with the Reference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference2

`func (o *TransportDto) SetReference2(v string)`

SetReference2 sets Reference2 field to given value.

### HasReference2

`func (o *TransportDto) HasReference2() bool`

HasReference2 returns a boolean if a field has been set.

### GetReference3

`func (o *TransportDto) GetReference3() string`

GetReference3 returns the Reference3 field if non-nil, zero value otherwise.

### GetReference3Ok

`func (o *TransportDto) GetReference3Ok() (*string, bool)`

GetReference3Ok returns a tuple with the Reference3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference3

`func (o *TransportDto) SetReference3(v string)`

SetReference3 sets Reference3 field to given value.

### HasReference3

`func (o *TransportDto) HasReference3() bool`

HasReference3 returns a boolean if a field has been set.

### GetComment

`func (o *TransportDto) GetComment() TransportCommentDto`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransportDto) GetCommentOk() (*TransportCommentDto, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransportDto) SetComment(v TransportCommentDto)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransportDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCashOnDelivery

`func (o *TransportDto) GetCashOnDelivery() TransportCashOnDeliveryDto`

GetCashOnDelivery returns the CashOnDelivery field if non-nil, zero value otherwise.

### GetCashOnDeliveryOk

`func (o *TransportDto) GetCashOnDeliveryOk() (*TransportCashOnDeliveryDto, bool)`

GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOnDelivery

`func (o *TransportDto) SetCashOnDelivery(v TransportCashOnDeliveryDto)`

SetCashOnDelivery sets CashOnDelivery field to given value.

### HasCashOnDelivery

`func (o *TransportDto) HasCashOnDelivery() bool`

HasCashOnDelivery returns a boolean if a field has been set.

### GetIncluded

`func (o *TransportDto) GetIncluded() TransportIncludedDto`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TransportDto) GetIncludedOk() (*TransportIncludedDto, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TransportDto) SetIncluded(v TransportIncludedDto)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TransportDto) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
