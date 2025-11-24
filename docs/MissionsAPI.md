# \MissionsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMission**](MissionsAPI.md#CancelMission) | **Post** /v3/missions/{uid}/cancellation-process | Cancel mission
[**CreateMission**](MissionsAPI.md#CreateMission) | **Post** /v3/missions | Create a mission
[**GetMissionById**](MissionsAPI.md#GetMissionById) | **Get** /v3/missions/{uid} | Get a mission by its unique identifier
[**GetMissionByMissionNumber**](MissionsAPI.md#GetMissionByMissionNumber) | **Get** /v3/missions/by-mission-number | Get a mission by its mission number
[**MissionCreationDryRun**](MissionsAPI.md#MissionCreationDryRun) | **Post** /v3/missions/dry-run | Mission creation dry run
[**MissionUpdateDryRun**](MissionsAPI.md#MissionUpdateDryRun) | **Patch** /v3/missions/{uid}/dry-run | Mission update dry run
[**UpdateMission**](MissionsAPI.md#UpdateMission) | **Patch** /v3/missions/{uid} | Update mission



## CancelMission

> CancelMission(ctx, uid).Command(command).Execute()

Cancel mission



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mission's unique identifier
	command := *openapiclient.NewCancelMissionCommand("00000000-0000-0000-0000-000000000000") // CancelMissionCommand | Cancel mission command

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MissionsAPI.CancelMission(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.CancelMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Mission&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**CancelMissionCommand**](CancelMissionCommand.md) | Cancel mission command | 

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


## CreateMission

> CreateMissionResultDto CreateMission(ctx).Command(command).Execute()

Create a mission

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
	command := *openapiclient.NewCreateMissionCommand([]openapiclient.CreateTransportDto{*openapiclient.NewCreateTransportDto("CustomerCode_example", "AgencyCode_example", *openapiclient.NewCreateTransportDtoStepDto(), *openapiclient.NewCreateTransportDtoStepDto(), "ServiceCode_example")}) // CreateMissionCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.CreateMission(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.CreateMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMission`: CreateMissionResultDto
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.CreateMission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**CreateMissionCommand**](CreateMissionCommand.md) |  | 

### Return type

[**CreateMissionResultDto**](CreateMissionResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMissionById

> MissionDto GetMissionById(ctx, uid).Fields(fields).Execute()

Get a mission by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mission's unique identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.GetMissionById(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.GetMissionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMissionById`: MissionDto
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.GetMissionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Mission&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMissionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**MissionDto**](MissionDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMissionByMissionNumber

> MissionDto GetMissionByMissionNumber(ctx).MissionNumber(missionNumber).Fields(fields).Execute()

Get a mission by its mission number

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
	missionNumber := int32(56) // int32 | Mission number
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.GetMissionByMissionNumber(context.Background()).MissionNumber(missionNumber).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.GetMissionByMissionNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMissionByMissionNumber`: MissionDto
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.GetMissionByMissionNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMissionByMissionNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **missionNumber** | **int32** | Mission number | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**MissionDto**](MissionDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MissionCreationDryRun

> MissionEntryDryRunDto MissionCreationDryRun(ctx).Command(command).Execute()

Mission creation dry run



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
	command := *openapiclient.NewSimulateMissionCreationDryRunCommand([]openapiclient.CreateTransportDto{*openapiclient.NewCreateTransportDto("CustomerCode_example", "AgencyCode_example", *openapiclient.NewCreateTransportDtoStepDto(), *openapiclient.NewCreateTransportDtoStepDto(), "ServiceCode_example")}) // SimulateMissionCreationDryRunCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.MissionCreationDryRun(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.MissionCreationDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MissionCreationDryRun`: MissionEntryDryRunDto
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.MissionCreationDryRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMissionCreationDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**SimulateMissionCreationDryRunCommand**](SimulateMissionCreationDryRunCommand.md) |  | 

### Return type

[**MissionEntryDryRunDto**](MissionEntryDryRunDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MissionUpdateDryRun

> MissionEntryDryRunDto MissionUpdateDryRun(ctx, uid).Command(command).Execute()

Mission update dry run



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
	command := *openapiclient.NewUpdateMissionDto() // UpdateMissionDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissionsAPI.MissionUpdateDryRun(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.MissionUpdateDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MissionUpdateDryRun`: MissionEntryDryRunDto
	fmt.Fprintf(os.Stdout, "Response from `MissionsAPI.MissionUpdateDryRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMissionUpdateDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateMissionDto**](UpdateMissionDto.md) |  | 

### Return type

[**MissionEntryDryRunDto**](MissionEntryDryRunDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMission

> UpdateMission(ctx, uid).Command(command).Execute()

Update mission



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mission's unique identifier
	command := *openapiclient.NewUpdateMissionDto() // UpdateMissionDto | Update mission command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MissionsAPI.UpdateMission(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissionsAPI.UpdateMission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Mission&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateMissionDto**](UpdateMissionDto.md) | Update mission command. | 

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

