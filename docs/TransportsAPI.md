# \TransportsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplySubState**](TransportsAPI.md#ApplySubState) | **Post** /v3/transports/{uid}/substates | Apply a sub state to the transport
[**ApplyTransportCancelledStatus**](TransportsAPI.md#ApplyTransportCancelledStatus) | **Post** /v3/transports/{uid}/cancellation-process | Apply transport cancelled status
[**ApplyTransportDeliveredStatus**](TransportsAPI.md#ApplyTransportDeliveredStatus) | **Post** /v3/transports/{uid}/delivery-process | Apply transport delivered state
[**ApplyTransportPickedUpStatus**](TransportsAPI.md#ApplyTransportPickedUpStatus) | **Post** /v3/transports/{uid}/pickup-process | Apply transport pickup state
[**ApplyTransportTerminatedStatus**](TransportsAPI.md#ApplyTransportTerminatedStatus) | **Post** /v3/transports/{uid}/termination-process | Apply transport terminated state
[**CheckTransportAssignmentsConstraints**](TransportsAPI.md#CheckTransportAssignmentsConstraints) | **Patch** /v3/transports/{uid}/assignment/constraints-validation | Check transport assignment&#39;s constraints
[**DefragmentTransport**](TransportsAPI.md#DefragmentTransport) | **Post** /v3/transports/{uid}/defragmentation-process | Defragment transport
[**FragmentTransport**](TransportsAPI.md#FragmentTransport) | **Post** /v3/transports/{uid}/fragmentation-process | Fragment transport
[**GetTransportAirData**](TransportsAPI.md#GetTransportAirData) | **Get** /v3/transports/{uid}/air-data | Get transport air data
[**GetTransportBillAddress**](TransportsAPI.md#GetTransportBillAddress) | **Get** /v3/transports/{uid}/transport-bill-address | Get transport bill address
[**GetTransportByUniqueIdentifier**](TransportsAPI.md#GetTransportByUniqueIdentifier) | **Get** /v3/transports/{uid} | Get a transport by uid
[**GetTransportCommunicationConfiguration**](TransportsAPI.md#GetTransportCommunicationConfiguration) | **Get** /v3/transports/{uid}/communication-configuration | Get transport communication configuration
[**GetTransportCustomerCustomParameters**](TransportsAPI.md#GetTransportCustomerCustomParameters) | **Get** /v3/transports/{uid}/customer-custom-parameters | Get transport customer custom parameters
[**GetTransportDangerousGoods**](TransportsAPI.md#GetTransportDangerousGoods) | **Get** /v3/transports/{uid}/dangerous-goods | Get transport dangerous goods
[**GetTransportHistoryLines**](TransportsAPI.md#GetTransportHistoryLines) | **Get** /v3/transports/{uid}/history | Get transport history lines
[**GetTransportPackageLines**](TransportsAPI.md#GetTransportPackageLines) | **Get** /v3/transports/{uid}/package-lines | Get transport package lines
[**GetTransportPricing**](TransportsAPI.md#GetTransportPricing) | **Get** /v3/transports/{uid}/pricing | Get transport pricing
[**GetTransportRoute**](TransportsAPI.md#GetTransportRoute) | **Get** /v3/transports/{uid}/transport-route | Get transport route
[**GetTransportServiceCustomParameters**](TransportsAPI.md#GetTransportServiceCustomParameters) | **Get** /v3/transports/{uid}/service-custom-parameters | Get transport service custom parameters
[**GetTransportSubServices**](TransportsAPI.md#GetTransportSubServices) | **Get** /v3/transports/{uid}/subservices | Get transport sub services
[**GetTransportTotalBulkSizes**](TransportsAPI.md#GetTransportTotalBulkSizes) | **Get** /v3/transports/{uid}/total-bulk-sizes | Get transport total bulk sizes
[**GetTransports**](TransportsAPI.md#GetTransports) | **Get** /v3/transports | Get transports
[**TransportUpdateDryRun**](TransportsAPI.md#TransportUpdateDryRun) | **Patch** /v3/transports/{uid}/dry-run | Transport update dry run
[**UpdateTransport**](TransportsAPI.md#UpdateTransport) | **Patch** /v3/transports/{uid} | Update transport
[**UpdateTransportAssignment**](TransportsAPI.md#UpdateTransportAssignment) | **Patch** /v3/transports/{uid}/assignment | Update transport assignment
[**UpdateTransportAssignmentStatus**](TransportsAPI.md#UpdateTransportAssignmentStatus) | **Post** /v3/transports/{uid}/assignment/status | Update transport assignment status



## ApplySubState

> ApplySubState(ctx, uid).Command(command).Execute()

Apply a sub state to the transport

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	command := *openapiclient.NewApplyTransportSubStateCommand("00000000-0000-0000-0000-000000000000", "SubStateCode_example") // ApplyTransportSubStateCommand | Sub state application command

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.ApplySubState(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.ApplySubState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplySubStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyTransportSubStateCommand**](ApplyTransportSubStateCommand.md) | Sub state application command | 

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


## ApplyTransportCancelledStatus

> ApplyTransportCancelledStatus(ctx, uid).Command(command).Execute()

Apply transport cancelled status



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
	command := *openapiclient.NewApplyCancelledStatusCommand("00000000-0000-0000-0000-000000000000") // ApplyCancelledStatusCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.ApplyTransportCancelledStatus(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.ApplyTransportCancelledStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTransportCancelledStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyCancelledStatusCommand**](ApplyCancelledStatusCommand.md) |  | 

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


## ApplyTransportDeliveredStatus

> ApplyTransportDeliveredStatus(ctx, uid).Command(command).Execute()

Apply transport delivered state



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
	command := *openapiclient.NewApplyDeliveredStatusCommand("00000000-0000-0000-0000-000000000000") // ApplyDeliveredStatusCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.ApplyTransportDeliveredStatus(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.ApplyTransportDeliveredStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTransportDeliveredStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyDeliveredStatusCommand**](ApplyDeliveredStatusCommand.md) |  | 

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


## ApplyTransportPickedUpStatus

> ApplyTransportPickedUpStatus(ctx, uid).Command(command).Execute()

Apply transport pickup state



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
	command := *openapiclient.NewApplyPickedUpStatusCommand("00000000-0000-0000-0000-000000000000") // ApplyPickedUpStatusCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.ApplyTransportPickedUpStatus(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.ApplyTransportPickedUpStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTransportPickedUpStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyPickedUpStatusCommand**](ApplyPickedUpStatusCommand.md) |  | 

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


## ApplyTransportTerminatedStatus

> ApplyTransportTerminatedStatus(ctx, uid).Command(command).Execute()

Apply transport terminated state



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
	command := *openapiclient.NewApplyTerminatedStatusCommand("00000000-0000-0000-0000-000000000000") // ApplyTerminatedStatusCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.ApplyTransportTerminatedStatus(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.ApplyTransportTerminatedStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyTransportTerminatedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyTerminatedStatusCommand**](ApplyTerminatedStatusCommand.md) |  | 

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


## CheckTransportAssignmentsConstraints

> CheckTransportAssignmentConstraintsResultDto CheckTransportAssignmentsConstraints(ctx, uid).Command(command).Execute()

Check transport assignment's constraints

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
	command := *openapiclient.NewAssignTransportDto() // AssignTransportDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.CheckTransportAssignmentsConstraints(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.CheckTransportAssignmentsConstraints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTransportAssignmentsConstraints`: CheckTransportAssignmentConstraintsResultDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.CheckTransportAssignmentsConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTransportAssignmentsConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**AssignTransportDto**](AssignTransportDto.md) |  | 

### Return type

[**CheckTransportAssignmentConstraintsResultDto**](CheckTransportAssignmentConstraintsResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DefragmentTransport

> DefragmentTransport(ctx, uid).Command(command).Execute()

Defragment transport

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
	command := *openapiclient.NewDefragmentTransportCommand("00000000-0000-0000-0000-000000000000") // DefragmentTransportCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.DefragmentTransport(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.DefragmentTransport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDefragmentTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**DefragmentTransportCommand**](DefragmentTransportCommand.md) |  | 

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


## FragmentTransport

> FragmentTransport(ctx, uid).Command(command).Execute()

Fragment transport

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	command := *openapiclient.NewFragmentTransportCommand("00000000-0000-0000-0000-000000000000", *openapiclient.NewFragmentTransportCommandStepDto(), time.Now()) // FragmentTransportCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.FragmentTransport(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.FragmentTransport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFragmentTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**FragmentTransportCommand**](FragmentTransportCommand.md) |  | 

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


## GetTransportAirData

> TransportAirDataDto GetTransportAirData(ctx, uid).Fields(fields).Execute()

Get transport air data



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportAirData(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportAirData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportAirData`: TransportAirDataDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportAirData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportAirDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**TransportAirDataDto**](TransportAirDataDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportBillAddress

> TransportBillAddressDto GetTransportBillAddress(ctx, uid).Fields(fields).Execute()

Get transport bill address



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportBillAddress(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportBillAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportBillAddress`: TransportBillAddressDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportBillAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportBillAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**TransportBillAddressDto**](TransportBillAddressDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportByUniqueIdentifier

> TransportDto GetTransportByUniqueIdentifier(ctx, uid).Fields(fields).Execute()

Get a transport by uid



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportByUniqueIdentifier(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportByUniqueIdentifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportByUniqueIdentifier`: TransportDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportByUniqueIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportByUniqueIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**TransportDto**](TransportDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportCommunicationConfiguration

> TransportCommunicationConfigurationDto GetTransportCommunicationConfiguration(ctx, uid).Execute()

Get transport communication configuration



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportCommunicationConfiguration(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportCommunicationConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportCommunicationConfiguration`: TransportCommunicationConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportCommunicationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportCommunicationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransportCommunicationConfigurationDto**](TransportCommunicationConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportCustomerCustomParameters

> IPagedResourceListTransportCustomParameterDto GetTransportCustomerCustomParameters(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport customer custom parameters



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportCustomerCustomParameters(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportCustomerCustomParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportCustomerCustomParameters`: IPagedResourceListTransportCustomParameterDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportCustomerCustomParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportCustomerCustomParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportCustomParameterDto**](IPagedResourceListTransportCustomParameterDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportDangerousGoods

> IResourceListTransportDangerousGoodDto GetTransportDangerousGoods(ctx, uid).Execute()

Get transport dangerous goods

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportDangerousGoods(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportDangerousGoods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDangerousGoods`: IResourceListTransportDangerousGoodDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportDangerousGoods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDangerousGoodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IResourceListTransportDangerousGoodDto**](IResourceListTransportDangerousGoodDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportHistoryLines

> IPagedResourceListTransportHistoryLineDto GetTransportHistoryLines(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport history lines



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportHistoryLines(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportHistoryLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportHistoryLines`: IPagedResourceListTransportHistoryLineDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportHistoryLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportHistoryLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportHistoryLineDto**](IPagedResourceListTransportHistoryLineDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportPackageLines

> IPagedResourceListPackageLineDto GetTransportPackageLines(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport package lines



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportPackageLines(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportPackageLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportPackageLines`: IPagedResourceListPackageLineDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportPackageLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportPackageLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPackageLineDto**](IPagedResourceListPackageLineDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportPricing

> TransportPricingDto GetTransportPricing(ctx, uid).Execute()

Get transport pricing



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportPricing(context.Background(), uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportPricing`: TransportPricingDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransportPricingDto**](TransportPricingDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportRoute

> TransportRouteDto GetTransportRoute(ctx, uid).StartDate(startDate).EndDate(endDate).UseSmoothRoute(useSmoothRoute).Execute()

Get transport route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startDate := time.Now() // time.Time | Optional start date, restrict to positions acquired after this date (optional)
	endDate := time.Now() // time.Time | Optional end date, restrict to positions acquired before this date (optional)
	useSmoothRoute := true // bool | Optional value to specify whether or not the route should be smoothed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportRoute(context.Background(), uid).StartDate(startDate).EndDate(endDate).UseSmoothRoute(useSmoothRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportRoute`: TransportRouteDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Optional start date, restrict to positions acquired after this date | 
 **endDate** | **time.Time** | Optional end date, restrict to positions acquired before this date | 
 **useSmoothRoute** | **bool** | Optional value to specify whether or not the route should be smoothed | 

### Return type

[**TransportRouteDto**](TransportRouteDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportServiceCustomParameters

> IPagedResourceListTransportCustomParameterDto GetTransportServiceCustomParameters(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport service custom parameters



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportServiceCustomParameters(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportServiceCustomParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportServiceCustomParameters`: IPagedResourceListTransportCustomParameterDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportServiceCustomParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportServiceCustomParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportCustomParameterDto**](IPagedResourceListTransportCustomParameterDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportSubServices

> IPagedResourceListTransportSubServiceDto GetTransportSubServices(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport sub services



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportSubServices(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportSubServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportSubServices`: IPagedResourceListTransportSubServiceDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportSubServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportSubServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportSubServiceDto**](IPagedResourceListTransportSubServiceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportTotalBulkSizes

> IPagedResourceListPackagesTotalBulkSizeDto GetTransportTotalBulkSizes(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport total bulk sizes



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransportTotalBulkSizes(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransportTotalBulkSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportTotalBulkSizes`: IPagedResourceListPackagesTotalBulkSizeDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransportTotalBulkSizes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportTotalBulkSizesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPackagesTotalBulkSizeDto**](IPagedResourceListPackagesTotalBulkSizeDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransports

> IPagedResourceListTransportDto GetTransports(ctx).TransportIds(transportIds).TransportUids(transportUids).MissionNumbers(missionNumbers).MissionUids(missionUids).MissionTrackIds(missionTrackIds).DriverIds(driverIds).VehicleCodes(vehicleCodes).TrailerCodes(trailerCodes).VehicleOrTrailerCodes(vehicleOrTrailerCodes).QuotationNumbers(quotationNumbers).QuotationUids(quotationUids).QuotationTrackIds(quotationTrackIds).SourceQuotationNumbers(sourceQuotationNumbers).DateFrom(dateFrom).DateTo(dateTo).DateActions(dateActions).DateToDateActions(dateToDateActions).DateCombinationCriterion(dateCombinationCriterion).Signs(signs).References1(references1).References2(references2).References3(references3).PickupCityIds(pickupCityIds).PickupPostCodes(pickupPostCodes).PickupCityNames(pickupCityNames).PickupAddressNames(pickupAddressNames).PickupCountryCodes(pickupCountryCodes).PickupSectorCodes(pickupSectorCodes).DeliveryCityIds(deliveryCityIds).DeliveryPostCodes(deliveryPostCodes).DeliveryCityNames(deliveryCityNames).DeliveryAddressNames(deliveryAddressNames).DeliveryCountryCodes(deliveryCountryCodes).DeliverySectorCodes(deliverySectorCodes).CustomerCodes(customerCodes).OrdererCodes(ordererCodes).AgencyCodes(agencyCodes).PackageBarCodes(packageBarCodes).Statuses(statuses).FileCategoryCodes(fileCategoryCodes).ServiceCodes(serviceCodes).PackageReferences1(packageReferences1).PackageReferences2(packageReferences2).PackageReferences3(packageReferences3).PackageReferences4(packageReferences4).PackageReferences5(packageReferences5).PackageReferences6(packageReferences6).PackageReferences7(packageReferences7).PackageReferences8(packageReferences8).Pattern(pattern).PatternFields(patternFields).MissionTypeVisibility(missionTypeVisibility).QuotationStatusVisibility(quotationStatusVisibility).QuotationApprovalStatusVisibility(quotationApprovalStatusVisibility).FlagLastSubStateByFamilyCodes(flagLastSubStateByFamilyCodes).FragmentationVisibility(fragmentationVisibility).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transportIds := []int32{int32(123)} // []int32 | Transports identifiers (optional)
	transportUids := []string{"Inner_example"} // []string | Transports unique identifiers (optional)
	missionNumbers := []int32{int32(123)} // []int32 | Missions numbers (optional)
	missionUids := []string{"Inner_example"} // []string | Missions unique identifiers (optional)
	missionTrackIds := []string{"Inner_example"} // []string | Missions track identifiers (optional)
	driverIds := []int32{int32(123)} // []int32 | Drivers identifiers (optional)
	vehicleCodes := []string{"Inner_example"} // []string | Vehicles codes (optional)
	trailerCodes := []string{"Inner_example"} // []string | Trailer codes (optional)
	vehicleOrTrailerCodes := []string{"Inner_example"} // []string | Codes for vehicles or trailers (optional)
	quotationNumbers := []int32{int32(123)} // []int32 | Quotations numbers (optional)
	quotationUids := []string{"Inner_example"} // []string | Quotations unique identifiers (optional)
	quotationTrackIds := []string{"Inner_example"} // []string | Quotations track identifiers (optional)
	sourceQuotationNumbers := []int32{int32(123)} // []int32 | Source quotation numbers, for transports included in a mission created from a quotation. (optional)
	dateFrom := time.Now() // time.Time | The inclusive start date and time for the search. (optional)
	dateTo := time.Now() // time.Time | The exclusive end date and time for the search. (optional)
	dateActions := []string{"DateActions_example"} // []string | Specify on which transport dates the date search will be performed.  If left empty, the search will be performed on any of the TransportDate. (optional)
	dateToDateActions := []string{"DateToDateActions_example"} // []string | Specify a range of transport dates to perform the research.  The expected schema is : From DateToDateActions dates to DateActions dates.  If left empty, the search will be performed using DateActions (optional)
	dateCombinationCriterion := "dateCombinationCriterion_example" // string | The combination criteria for date search.  Have no impact if DateActions is unspecified or contains a single element.    Default value is 'Or'. (optional) (default to "Or")
	signs := []string{"Inner_example"} // []string |  (optional)
	references1 := []string{"Inner_example"} // []string |  (optional)
	references2 := []string{"Inner_example"} // []string |  (optional)
	references3 := []string{"Inner_example"} // []string |  (optional)
	pickupCityIds := []int32{int32(123)} // []int32 |  (optional)
	pickupPostCodes := []string{"Inner_example"} // []string |  (optional)
	pickupCityNames := []string{"Inner_example"} // []string |  (optional)
	pickupAddressNames := []string{"Inner_example"} // []string |  (optional)
	pickupCountryCodes := []string{"Inner_example"} // []string |  (optional)
	pickupSectorCodes := []string{"Inner_example"} // []string |  (optional)
	deliveryCityIds := []int32{int32(123)} // []int32 |  (optional)
	deliveryPostCodes := []string{"Inner_example"} // []string |  (optional)
	deliveryCityNames := []string{"Inner_example"} // []string |  (optional)
	deliveryAddressNames := []string{"Inner_example"} // []string |  (optional)
	deliveryCountryCodes := []string{"Inner_example"} // []string |  (optional)
	deliverySectorCodes := []string{"Inner_example"} // []string |  (optional)
	customerCodes := []string{"Inner_example"} // []string |  (optional)
	ordererCodes := []string{"Inner_example"} // []string |  (optional)
	agencyCodes := []string{"Inner_example"} // []string | A list of exact agency codes to match against the transports agency's code. (optional)
	packageBarCodes := []string{"Inner_example"} // []string |  (optional)
	statuses := []string{"Inner_example"} // []string |  (optional)
	fileCategoryCodes := []string{"Inner_example"} // []string |  (optional)
	serviceCodes := []string{"Inner_example"} // []string |  (optional)
	packageReferences1 := []string{"Inner_example"} // []string |  (optional)
	packageReferences2 := []string{"Inner_example"} // []string |  (optional)
	packageReferences3 := []string{"Inner_example"} // []string |  (optional)
	packageReferences4 := []string{"Inner_example"} // []string |  (optional)
	packageReferences5 := []string{"Inner_example"} // []string |  (optional)
	packageReferences6 := []string{"Inner_example"} // []string |  (optional)
	packageReferences7 := []string{"Inner_example"} // []string |  (optional)
	packageReferences8 := []string{"Inner_example"} // []string |  (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	missionTypeVisibility := "missionTypeVisibility_example" // string | Should the query return only missions, only quotations, or both.    Default value is 'MissionsOnly'. (optional) (default to "MissionsOnly")
	quotationStatusVisibility := "quotationStatusVisibility_example" // string | Should the query return active quotations, archived quotations, or both.    Default value is 'Active'. (optional) (default to "Active")
	quotationApprovalStatusVisibility := "quotationApprovalStatusVisibility_example" // string | Should the query return quotations with no approval required, quotations awaiting for approval, or both.    Default value is 'All'. (optional) (default to "All")
	flagLastSubStateByFamilyCodes := []string{"Inner_example"} // []string | Flag the last substate applied to a transport for each substate family code. (optional)
	fragmentationVisibility := "fragmentationVisibility_example" // string | Should the query return parent, children, or both in the context of fragmented transports.  Defaults to All (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.GetTransports(context.Background()).TransportIds(transportIds).TransportUids(transportUids).MissionNumbers(missionNumbers).MissionUids(missionUids).MissionTrackIds(missionTrackIds).DriverIds(driverIds).VehicleCodes(vehicleCodes).TrailerCodes(trailerCodes).VehicleOrTrailerCodes(vehicleOrTrailerCodes).QuotationNumbers(quotationNumbers).QuotationUids(quotationUids).QuotationTrackIds(quotationTrackIds).SourceQuotationNumbers(sourceQuotationNumbers).DateFrom(dateFrom).DateTo(dateTo).DateActions(dateActions).DateToDateActions(dateToDateActions).DateCombinationCriterion(dateCombinationCriterion).Signs(signs).References1(references1).References2(references2).References3(references3).PickupCityIds(pickupCityIds).PickupPostCodes(pickupPostCodes).PickupCityNames(pickupCityNames).PickupAddressNames(pickupAddressNames).PickupCountryCodes(pickupCountryCodes).PickupSectorCodes(pickupSectorCodes).DeliveryCityIds(deliveryCityIds).DeliveryPostCodes(deliveryPostCodes).DeliveryCityNames(deliveryCityNames).DeliveryAddressNames(deliveryAddressNames).DeliveryCountryCodes(deliveryCountryCodes).DeliverySectorCodes(deliverySectorCodes).CustomerCodes(customerCodes).OrdererCodes(ordererCodes).AgencyCodes(agencyCodes).PackageBarCodes(packageBarCodes).Statuses(statuses).FileCategoryCodes(fileCategoryCodes).ServiceCodes(serviceCodes).PackageReferences1(packageReferences1).PackageReferences2(packageReferences2).PackageReferences3(packageReferences3).PackageReferences4(packageReferences4).PackageReferences5(packageReferences5).PackageReferences6(packageReferences6).PackageReferences7(packageReferences7).PackageReferences8(packageReferences8).Pattern(pattern).PatternFields(patternFields).MissionTypeVisibility(missionTypeVisibility).QuotationStatusVisibility(quotationStatusVisibility).QuotationApprovalStatusVisibility(quotationApprovalStatusVisibility).FlagLastSubStateByFamilyCodes(flagLastSubStateByFamilyCodes).FragmentationVisibility(fragmentationVisibility).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.GetTransports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransports`: IPagedResourceListTransportDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.GetTransports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportIds** | **[]int32** | Transports identifiers | 
 **transportUids** | **[]string** | Transports unique identifiers | 
 **missionNumbers** | **[]int32** | Missions numbers | 
 **missionUids** | **[]string** | Missions unique identifiers | 
 **missionTrackIds** | **[]string** | Missions track identifiers | 
 **driverIds** | **[]int32** | Drivers identifiers | 
 **vehicleCodes** | **[]string** | Vehicles codes | 
 **trailerCodes** | **[]string** | Trailer codes | 
 **vehicleOrTrailerCodes** | **[]string** | Codes for vehicles or trailers | 
 **quotationNumbers** | **[]int32** | Quotations numbers | 
 **quotationUids** | **[]string** | Quotations unique identifiers | 
 **quotationTrackIds** | **[]string** | Quotations track identifiers | 
 **sourceQuotationNumbers** | **[]int32** | Source quotation numbers, for transports included in a mission created from a quotation. | 
 **dateFrom** | **time.Time** | The inclusive start date and time for the search. | 
 **dateTo** | **time.Time** | The exclusive end date and time for the search. | 
 **dateActions** | **[]string** | Specify on which transport dates the date search will be performed.  If left empty, the search will be performed on any of the TransportDate. | 
 **dateToDateActions** | **[]string** | Specify a range of transport dates to perform the research.  The expected schema is : From DateToDateActions dates to DateActions dates.  If left empty, the search will be performed using DateActions | 
 **dateCombinationCriterion** | **string** | The combination criteria for date search.  Have no impact if DateActions is unspecified or contains a single element.    Default value is &#39;Or&#39;. | [default to &quot;Or&quot;]
 **signs** | **[]string** |  | 
 **references1** | **[]string** |  | 
 **references2** | **[]string** |  | 
 **references3** | **[]string** |  | 
 **pickupCityIds** | **[]int32** |  | 
 **pickupPostCodes** | **[]string** |  | 
 **pickupCityNames** | **[]string** |  | 
 **pickupAddressNames** | **[]string** |  | 
 **pickupCountryCodes** | **[]string** |  | 
 **pickupSectorCodes** | **[]string** |  | 
 **deliveryCityIds** | **[]int32** |  | 
 **deliveryPostCodes** | **[]string** |  | 
 **deliveryCityNames** | **[]string** |  | 
 **deliveryAddressNames** | **[]string** |  | 
 **deliveryCountryCodes** | **[]string** |  | 
 **deliverySectorCodes** | **[]string** |  | 
 **customerCodes** | **[]string** |  | 
 **ordererCodes** | **[]string** |  | 
 **agencyCodes** | **[]string** | A list of exact agency codes to match against the transports agency&#39;s code. | 
 **packageBarCodes** | **[]string** |  | 
 **statuses** | **[]string** |  | 
 **fileCategoryCodes** | **[]string** |  | 
 **serviceCodes** | **[]string** |  | 
 **packageReferences1** | **[]string** |  | 
 **packageReferences2** | **[]string** |  | 
 **packageReferences3** | **[]string** |  | 
 **packageReferences4** | **[]string** |  | 
 **packageReferences5** | **[]string** |  | 
 **packageReferences6** | **[]string** |  | 
 **packageReferences7** | **[]string** |  | 
 **packageReferences8** | **[]string** |  | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **missionTypeVisibility** | **string** | Should the query return only missions, only quotations, or both.    Default value is &#39;MissionsOnly&#39;. | [default to &quot;MissionsOnly&quot;]
 **quotationStatusVisibility** | **string** | Should the query return active quotations, archived quotations, or both.    Default value is &#39;Active&#39;. | [default to &quot;Active&quot;]
 **quotationApprovalStatusVisibility** | **string** | Should the query return quotations with no approval required, quotations awaiting for approval, or both.    Default value is &#39;All&#39;. | [default to &quot;All&quot;]
 **flagLastSubStateByFamilyCodes** | **[]string** | Flag the last substate applied to a transport for each substate family code. | 
 **fragmentationVisibility** | **string** | Should the query return parent, children, or both in the context of fragmented transports.  Defaults to All | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportDto**](IPagedResourceListTransportDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportUpdateDryRun

> TransportEntryDryRunDto TransportUpdateDryRun(ctx, uid).Command(command).Execute()

Transport update dry run



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
	command := *openapiclient.NewUpdateTransportDto() // UpdateTransportDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportsAPI.TransportUpdateDryRun(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.TransportUpdateDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransportUpdateDryRun`: TransportEntryDryRunDto
	fmt.Fprintf(os.Stdout, "Response from `TransportsAPI.TransportUpdateDryRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportUpdateDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateTransportDto**](UpdateTransportDto.md) |  | 

### Return type

[**TransportEntryDryRunDto**](TransportEntryDryRunDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransport

> UpdateTransport(ctx, uid).Command(command).Execute()

Update transport



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport's unique identifier
	command := *openapiclient.NewUpdateTransportDto() // UpdateTransportDto | Update transport command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.UpdateTransport(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.UpdateTransport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateTransportDto**](UpdateTransportDto.md) | Update transport command. | 

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


## UpdateTransportAssignment

> UpdateTransportAssignment(ctx, uid).Command(command).Execute()

Update transport assignment



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
	command := *openapiclient.NewAssignTransportDto() // AssignTransportDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.UpdateTransportAssignment(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.UpdateTransportAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransportAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**AssignTransportDto**](AssignTransportDto.md) |  | 

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


## UpdateTransportAssignmentStatus

> UpdateTransportAssignmentStatus(ctx, uid).Command(command).Execute()

Update transport assignment status

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
	command := *openapiclient.NewApplyAssignmentStatusChangeCommand("00000000-0000-0000-0000-000000000000", "Status_example") // ApplyAssignmentStatusChangeCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportsAPI.UpdateTransportAssignmentStatus(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportsAPI.UpdateTransportAssignmentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransportAssignmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**ApplyAssignmentStatusChangeCommand**](ApplyAssignmentStatusChangeCommand.md) |  | 

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

