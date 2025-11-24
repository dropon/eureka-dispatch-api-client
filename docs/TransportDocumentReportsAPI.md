# \TransportDocumentReportsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadTransportDocumentReport**](TransportDocumentReportsAPI.md#DownloadTransportDocumentReport) | **Get** /v3/transport-document-reports/{uid}/data | Download a transport document report by its unique identifier
[**GenerateTransportDocumentReport**](TransportDocumentReportsAPI.md#GenerateTransportDocumentReport) | **Post** /v3/transport-document-reports | Generate transport document reports
[**GetTransportDocumentReportById**](TransportDocumentReportsAPI.md#GetTransportDocumentReportById) | **Get** /v3/transport-document-reports/{uid} | Get a transport document report by its unique identifier



## DownloadTransportDocumentReport

> *os.File DownloadTransportDocumentReport(ctx, uid).Execute()

Download a transport document report by its unique identifier

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the transport document report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentReportsAPI.DownloadTransportDocumentReport(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentReportsAPI.DownloadTransportDocumentReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadTransportDocumentReport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentReportsAPI.DownloadTransportDocumentReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the transport document report | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTransportDocumentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateTransportDocumentReport

> GenerateTransportDocumentReportResultDto GenerateTransportDocumentReport(ctx).Command(command).Execute()

Generate transport document reports

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	command := *openapiclient.NewGenerateTransportDocumentReportCommand([]string{"00000000-0000-0000-0000-000000000000"}, "ReportType_example") // GenerateTransportDocumentReportCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentReportsAPI.GenerateTransportDocumentReport(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentReportsAPI.GenerateTransportDocumentReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateTransportDocumentReport`: GenerateTransportDocumentReportResultDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentReportsAPI.GenerateTransportDocumentReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTransportDocumentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**GenerateTransportDocumentReportCommand**](GenerateTransportDocumentReportCommand.md) |  | 

### Return type

[**GenerateTransportDocumentReportResultDto**](GenerateTransportDocumentReportResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportDocumentReportById

> DocumentReportDto GetTransportDocumentReportById(ctx, uid).Execute()

Get a transport document report by its unique identifier

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the transport document report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentReportsAPI.GetTransportDocumentReportById(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentReportsAPI.GetTransportDocumentReportById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDocumentReportById`: DocumentReportDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentReportsAPI.GetTransportDocumentReportById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the transport document report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDocumentReportByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentReportDto**](DocumentReportDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

