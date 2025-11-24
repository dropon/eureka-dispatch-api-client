# GenerateTransportDocumentReportCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportUids** | **[]string** | Mandatory. List of transport unique identifiers. | 
**ReportType** | **string** | Mandatory. Report document type to generate. | 
**PaperOptions** | Pointer to [**ReportPaperOptionsDto**](ReportPaperOptionsDto.md) |  | [optional] 

## Methods

### NewGenerateTransportDocumentReportCommand

`func NewGenerateTransportDocumentReportCommand(transportUids []string, reportType string, ) *GenerateTransportDocumentReportCommand`

NewGenerateTransportDocumentReportCommand instantiates a new GenerateTransportDocumentReportCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateTransportDocumentReportCommandWithDefaults

`func NewGenerateTransportDocumentReportCommandWithDefaults() *GenerateTransportDocumentReportCommand`

NewGenerateTransportDocumentReportCommandWithDefaults instantiates a new GenerateTransportDocumentReportCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUids

`func (o *GenerateTransportDocumentReportCommand) GetTransportUids() []string`

GetTransportUids returns the TransportUids field if non-nil, zero value otherwise.

### GetTransportUidsOk

`func (o *GenerateTransportDocumentReportCommand) GetTransportUidsOk() (*[]string, bool)`

GetTransportUidsOk returns a tuple with the TransportUids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUids

`func (o *GenerateTransportDocumentReportCommand) SetTransportUids(v []string)`

SetTransportUids sets TransportUids field to given value.


### GetReportType

`func (o *GenerateTransportDocumentReportCommand) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *GenerateTransportDocumentReportCommand) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *GenerateTransportDocumentReportCommand) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetPaperOptions

`func (o *GenerateTransportDocumentReportCommand) GetPaperOptions() ReportPaperOptionsDto`

GetPaperOptions returns the PaperOptions field if non-nil, zero value otherwise.

### GetPaperOptionsOk

`func (o *GenerateTransportDocumentReportCommand) GetPaperOptionsOk() (*ReportPaperOptionsDto, bool)`

GetPaperOptionsOk returns a tuple with the PaperOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperOptions

`func (o *GenerateTransportDocumentReportCommand) SetPaperOptions(v ReportPaperOptionsDto)`

SetPaperOptions sets PaperOptions field to given value.

### HasPaperOptions

`func (o *GenerateTransportDocumentReportCommand) HasPaperOptions() bool`

HasPaperOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


