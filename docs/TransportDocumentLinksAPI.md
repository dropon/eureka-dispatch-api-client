# \TransportDocumentLinksAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransportDocumentLink**](TransportDocumentLinksAPI.md#CreateTransportDocumentLink) | **Post** /v3/transports/{uid}/documents/links | Create a transport document link
[**DeleteTransportDocumentLink**](TransportDocumentLinksAPI.md#DeleteTransportDocumentLink) | **Delete** /v3/transports/{uid}/documents/links/{linkId} | Delete transport document link
[**GetTransportDocumentLinkById**](TransportDocumentLinksAPI.md#GetTransportDocumentLinkById) | **Get** /v3/transports/{uid}/documents/links/{linkId} | Get transport document link by id
[**GetTransportDocumentLinks**](TransportDocumentLinksAPI.md#GetTransportDocumentLinks) | **Get** /v3/transports/{uid}/documents/links | Get transport document links



## CreateTransportDocumentLink

> CreateTransportDocumentLinkResultDto CreateTransportDocumentLink(ctx, uid).Command(command).Execute()

Create a transport document link



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
	command := *openapiclient.NewCreateTransportDocumentLinkCommand("00000000-0000-0000-0000-000000000000", "Link_example") // CreateTransportDocumentLinkCommand | Transport link creation command

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentLinksAPI.CreateTransportDocumentLink(context.Background(), uid).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentLinksAPI.CreateTransportDocumentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransportDocumentLink`: CreateTransportDocumentLinkResultDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentLinksAPI.CreateTransportDocumentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransportDocumentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**CreateTransportDocumentLinkCommand**](CreateTransportDocumentLinkCommand.md) | Transport link creation command | 

### Return type

[**CreateTransportDocumentLinkResultDto**](CreateTransportDocumentLinkResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransportDocumentLink

> DeleteTransportDocumentLink(ctx, uid, linkId).Execute()

Delete transport document link



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
	linkId := "linkId_example" // string | Link identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportDocumentLinksAPI.DeleteTransportDocumentLink(context.Background(), uid, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentLinksAPI.DeleteTransportDocumentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 
**linkId** | **string** | Link identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransportDocumentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportDocumentLinkById

> TransportDocumentLinkDto GetTransportDocumentLinkById(ctx, uid, linkId).Fields(fields).Execute()

Get transport document link by id



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
	linkId := "linkId_example" // string | Link identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentLinksAPI.GetTransportDocumentLinkById(context.Background(), uid, linkId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentLinksAPI.GetTransportDocumentLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDocumentLinkById`: TransportDocumentLinkDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentLinksAPI.GetTransportDocumentLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 
**linkId** | **string** | Link identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDocumentLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**TransportDocumentLinkDto**](TransportDocumentLinkDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportDocumentLinks

> IPagedResourceListTransportDocumentLinkDto GetTransportDocumentLinks(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport document links



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
	resp, r, err := apiClient.TransportDocumentLinksAPI.GetTransportDocumentLinks(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentLinksAPI.GetTransportDocumentLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDocumentLinks`: IPagedResourceListTransportDocumentLinkDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentLinksAPI.GetTransportDocumentLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDocumentLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportDocumentLinkDto**](IPagedResourceListTransportDocumentLinkDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

