# \ApiFlagsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiFlagByKey**](ApiFlagsAPI.md#GetApiFlagByKey) | **Get** /v3/apiflags/{key} | Get api flag by key
[**GetApiFlags**](ApiFlagsAPI.md#GetApiFlags) | **Get** /v3/apiflags | Get api flags



## GetApiFlagByKey

> ApiFlagsDtoApiFlag GetApiFlagByKey(ctx, key).Execute()

Get api flag by key



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
	key := "key_example" // string | Api flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiFlagsAPI.GetApiFlagByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiFlagsAPI.GetApiFlagByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiFlagByKey`: ApiFlagsDtoApiFlag
	fmt.Fprintf(os.Stdout, "Response from `ApiFlagsAPI.GetApiFlagByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Api flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiFlagByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiFlagsDtoApiFlag**](ApiFlagsDtoApiFlag.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiFlags

> ApiFlagsDto GetApiFlags(ctx).Execute()

Get api flags



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiFlagsAPI.GetApiFlags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiFlagsAPI.GetApiFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiFlags`: ApiFlagsDto
	fmt.Fprintf(os.Stdout, "Response from `ApiFlagsAPI.GetApiFlags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiFlagsRequest struct via the builder pattern


### Return type

[**ApiFlagsDto**](ApiFlagsDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

