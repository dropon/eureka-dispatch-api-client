# \QuotationsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelQuotation**](QuotationsAPI.md#CancelQuotation) | **Post** /v3/quotations/{uid}/cancellation-process | Cancel quotation
[**CreateMissionFromQuotation**](QuotationsAPI.md#CreateMissionFromQuotation) | **Post** /v3/quotations/{uid}/to-mission-process | Create mission from quotation
[**CreateQuotation**](QuotationsAPI.md#CreateQuotation) | **Post** /v3/quotations | Create a quotation
[**GetQuotationById**](QuotationsAPI.md#GetQuotationById) | **Get** /v3/quotations/{uid} | Get a quotation by its unique identifier
[**GetQuotationByQuotationNumber**](QuotationsAPI.md#GetQuotationByQuotationNumber) | **Get** /v3/quotations/by-quotation-number | Get a quotation by its quotation number
[**QuotationCreationDryRun**](QuotationsAPI.md#QuotationCreationDryRun) | **Post** /v3/quotations/dry-run | Quotation creation dry run
[**QuotationUpdateDryRun**](QuotationsAPI.md#QuotationUpdateDryRun) | **Patch** /v3/quotations/{uid}/dry-run | Quotation update dry run
[**UpdateQuotation**](QuotationsAPI.md#UpdateQuotation) | **Patch** /v3/quotations/{uid} | Update quotation



## CancelQuotation

> CancelQuotation(ctx, uid).Command(command).Execute()

Cancel quotation



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Quotation's unique identifier
	command := *openapiclient.NewCancelQuotationCommand("00000000-0000-0000-0000-000000000000") // CancelQuotationCommand | Cancel quotation command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QuotationsAPI.CancelQuotation(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.CancelQuotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Quotation&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**CancelQuotationCommand**](CancelQuotationCommand.md) | Cancel quotation command. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMissionFromQuotation

> CreateMissionFromQuotationResultDto CreateMissionFromQuotation(ctx, uid).Command(command).Execute()

Create mission from quotation



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Quotation's unique identifier
	command := *openapiclient.NewCreateMissionFromQuotationCommand("00000000-0000-0000-0000-000000000000") // CreateMissionFromQuotationCommand | Quotation to mission command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.CreateMissionFromQuotation(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.CreateMissionFromQuotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMissionFromQuotation`: CreateMissionFromQuotationResultDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.CreateMissionFromQuotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Quotation&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMissionFromQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**CreateMissionFromQuotationCommand**](CreateMissionFromQuotationCommand.md) | Quotation to mission command. | 

### Return type

[**CreateMissionFromQuotationResultDto**](CreateMissionFromQuotationResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuotation

> CreateQuotationResultDto CreateQuotation(ctx).Command(command).Execute()

Create a quotation

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
	command := *openapiclient.NewCreateQuotationCommand([]openapiclient.CreateTransportDto{*openapiclient.NewCreateTransportDto("CustomerCode_example", "AgencyCode_example", *openapiclient.NewCreateTransportDtoStepDto(), *openapiclient.NewCreateTransportDtoStepDto(), "ServiceCode_example")}) // CreateQuotationCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.CreateQuotation(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.CreateQuotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuotation`: CreateQuotationResultDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.CreateQuotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**CreateQuotationCommand**](CreateQuotationCommand.md) |  | 

### Return type

[**CreateQuotationResultDto**](CreateQuotationResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotationById

> QuotationDto GetQuotationById(ctx, uid).Fields(fields).Execute()

Get a quotation by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Quotation's unique identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.GetQuotationById(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.GetQuotationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotationById`: QuotationDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.GetQuotationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Quotation&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**QuotationDto**](QuotationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotationByQuotationNumber

> QuotationDto GetQuotationByQuotationNumber(ctx).QuotationNumber(quotationNumber).Fields(fields).Execute()

Get a quotation by its quotation number

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
	quotationNumber := int32(56) // int32 | Quotation number
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.GetQuotationByQuotationNumber(context.Background()).QuotationNumber(quotationNumber).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.GetQuotationByQuotationNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotationByQuotationNumber`: QuotationDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.GetQuotationByQuotationNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotationByQuotationNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotationNumber** | **int32** | Quotation number | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**QuotationDto**](QuotationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotationCreationDryRun

> QuotationEntryDryRunDto QuotationCreationDryRun(ctx).Command(command).Execute()

Quotation creation dry run



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
	command := *openapiclient.NewSimulateQuotationCreationDryRunCommand([]openapiclient.CreateTransportDto{*openapiclient.NewCreateTransportDto("CustomerCode_example", "AgencyCode_example", *openapiclient.NewCreateTransportDtoStepDto(), *openapiclient.NewCreateTransportDtoStepDto(), "ServiceCode_example")}) // SimulateQuotationCreationDryRunCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.QuotationCreationDryRun(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.QuotationCreationDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuotationCreationDryRun`: QuotationEntryDryRunDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.QuotationCreationDryRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuotationCreationDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**SimulateQuotationCreationDryRunCommand**](SimulateQuotationCreationDryRunCommand.md) |  | 

### Return type

[**QuotationEntryDryRunDto**](QuotationEntryDryRunDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotationUpdateDryRun

> QuotationEntryDryRunDto QuotationUpdateDryRun(ctx, uid).Command(command).Execute()

Quotation update dry run



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	command := *openapiclient.NewUpdateQuotationDto() // UpdateQuotationDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotationsAPI.QuotationUpdateDryRun(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.QuotationUpdateDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuotationUpdateDryRun`: QuotationEntryDryRunDto
	fmt.Fprintf(os.Stdout, "Response from `QuotationsAPI.QuotationUpdateDryRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotationUpdateDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateQuotationDto**](UpdateQuotationDto.md) |  | 

### Return type

[**QuotationEntryDryRunDto**](QuotationEntryDryRunDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuotation

> UpdateQuotation(ctx, uid).Command(command).Execute()

Update quotation



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Quotation's unique identifier
	command := *openapiclient.NewUpdateQuotationDto() // UpdateQuotationDto | Update quotation command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QuotationsAPI.UpdateQuotation(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotationsAPI.UpdateQuotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Quotation&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateQuotationDto**](UpdateQuotationDto.md) | Update quotation command. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

