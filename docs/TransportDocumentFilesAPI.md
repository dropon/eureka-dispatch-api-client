# \TransportDocumentFilesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTransportDocumentFile**](TransportDocumentFilesAPI.md#DeleteTransportDocumentFile) | **Delete** /v3/transports/{uid}/documents/files/{fileId} | Delete transport document file
[**DownloadTransportDocumentFile**](TransportDocumentFilesAPI.md#DownloadTransportDocumentFile) | **Get** /v3/transports/{uid}/documents/files/{fileId}/data | Download transport document file
[**GetTransportDocumentFileById**](TransportDocumentFilesAPI.md#GetTransportDocumentFileById) | **Get** /v3/transports/{uid}/documents/files/{fileId} | Get transport document file by id
[**GetTransportDocumentFiles**](TransportDocumentFilesAPI.md#GetTransportDocumentFiles) | **Get** /v3/transports/{uid}/documents/files | Get transport document files
[**UploadTransportDocumentFile**](TransportDocumentFilesAPI.md#UploadTransportDocumentFile) | **Post** /v3/transports/{uid}/documents/files | Upload transport document file



## DeleteTransportDocumentFile

> DeleteTransportDocumentFile(ctx, uid, fileId).Execute()

Delete transport document file



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
	fileId := "fileId_example" // string | File identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportDocumentFilesAPI.DeleteTransportDocumentFile(context.Background(), uid, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentFilesAPI.DeleteTransportDocumentFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 
**fileId** | **string** | File identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransportDocumentFileRequest struct via the builder pattern


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


## DownloadTransportDocumentFile

> *os.File DownloadTransportDocumentFile(ctx, uid, fileId).Execute()

Download transport document file



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
	fileId := "fileId_example" // string | File identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentFilesAPI.DownloadTransportDocumentFile(context.Background(), uid, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentFilesAPI.DownloadTransportDocumentFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadTransportDocumentFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentFilesAPI.DownloadTransportDocumentFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 
**fileId** | **string** | File identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTransportDocumentFileRequest struct via the builder pattern


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


## GetTransportDocumentFileById

> TransportDocumentFileDto GetTransportDocumentFileById(ctx, uid, fileId).Fields(fields).Execute()

Get transport document file by id



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
	fileId := "fileId_example" // string | File identifier
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentFilesAPI.GetTransportDocumentFileById(context.Background(), uid, fileId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentFilesAPI.GetTransportDocumentFileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDocumentFileById`: TransportDocumentFileDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentFilesAPI.GetTransportDocumentFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 
**fileId** | **string** | File identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDocumentFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**TransportDocumentFileDto**](TransportDocumentFileDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportDocumentFiles

> IPagedResourceListTransportDocumentFileDto GetTransportDocumentFiles(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get transport document files



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
	resp, r, err := apiClient.TransportDocumentFilesAPI.GetTransportDocumentFiles(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentFilesAPI.GetTransportDocumentFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportDocumentFiles`: IPagedResourceListTransportDocumentFileDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentFilesAPI.GetTransportDocumentFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportDocumentFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListTransportDocumentFileDto**](IPagedResourceListTransportDocumentFileDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTransportDocumentFile

> UploadTransportDocumentFileResultDto UploadTransportDocumentFile(ctx, uid).File(file).FileCategoryCode(fileCategoryCode).FileLabel(fileLabel).Execute()

Upload transport document file



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	fileCategoryCode := "fileCategoryCode_example" // string | 
	fileLabel := "fileLabel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportDocumentFilesAPI.UploadTransportDocumentFile(context.Background(), uid).File(file).FileCategoryCode(fileCategoryCode).FileLabel(fileLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentFilesAPI.UploadTransportDocumentFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadTransportDocumentFile`: UploadTransportDocumentFileResultDto
	fmt.Fprintf(os.Stdout, "Response from `TransportDocumentFilesAPI.UploadTransportDocumentFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Transport&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadTransportDocumentFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **fileCategoryCode** | **string** |  | 
 **fileLabel** | **string** |  | 

### Return type

[**UploadTransportDocumentFileResultDto**](UploadTransportDocumentFileResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

