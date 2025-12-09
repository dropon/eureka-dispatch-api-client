# DocumentReportDto

## Properties

| Name                          | Type                  | Description                        | Notes      |
| ----------------------------- | --------------------- | ---------------------------------- | ---------- |
| **Uid**                       | Pointer to **string** | Identifier of the document report. | [optional] |
| **ReportGenerationStartDate** | Pointer to **Time**   |                                    | [optional] |
| **Status**                    | Pointer to **string** |                                    | [optional] |
| **ErrorDescription**          | Pointer to **string** |                                    | [optional] |

## Methods

### NewDocumentReportDto

`func NewDocumentReportDto() *DocumentReportDto`

NewDocumentReportDto instantiates a new DocumentReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentReportDtoWithDefaults

`func NewDocumentReportDtoWithDefaults() *DocumentReportDto`

NewDocumentReportDtoWithDefaults instantiates a new DocumentReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *DocumentReportDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DocumentReportDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DocumentReportDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DocumentReportDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetReportGenerationStartDate

`func (o *DocumentReportDto) GetReportGenerationStartDate() Time`

GetReportGenerationStartDate returns the ReportGenerationStartDate field if non-nil, zero value otherwise.

### GetReportGenerationStartDateOk

`func (o *DocumentReportDto) GetReportGenerationStartDateOk() (*Time, bool)`

GetReportGenerationStartDateOk returns a tuple with the ReportGenerationStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGenerationStartDate

`func (o *DocumentReportDto) SetReportGenerationStartDate(v Time)`

SetReportGenerationStartDate sets ReportGenerationStartDate field to given value.

### HasReportGenerationStartDate

`func (o *DocumentReportDto) HasReportGenerationStartDate() bool`

HasReportGenerationStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentReportDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentReportDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentReportDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentReportDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorDescription

`func (o *DocumentReportDto) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *DocumentReportDto) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *DocumentReportDto) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *DocumentReportDto) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
