# \FileCategoriesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileCategories**](FileCategoriesAPI.md#GetFileCategories) | **Get** /v3/file-categories | Get file categories



## GetFileCategories

> IPagedResourceListFileCategoryDto GetFileCategories(ctx).Code(code).Label(label).Codes(codes).Labels(labels).IsVisibleForOrderers(isVisibleForOrderers).IsVisibleForDrivers(isVisibleForDrivers).IsBillingReceipt(isBillingReceipt).ServiceUids(serviceUids).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get file categories



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
	code := "code_example" // string | Search by file category code (optional)
	label := "label_example" // string | Search by file category label (optional)
	codes := []string{"Inner_example"} // []string | Search by file category codes (optional)
	labels := []string{"Inner_example"} // []string | Search by file category labels (optional)
	isVisibleForOrderers := true // bool | Search by file category orderer visibility (optional)
	isVisibleForDrivers := true // bool | Search by file category driver visibility (optional)
	isBillingReceipt := true // bool | Search by file category being billing receipt (optional)
	serviceUids := []string{"Inner_example"} // []string | Search by file category service uids (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileCategoriesAPI.GetFileCategories(context.Background()).Code(code).Label(label).Codes(codes).Labels(labels).IsVisibleForOrderers(isVisibleForOrderers).IsVisibleForDrivers(isVisibleForDrivers).IsBillingReceipt(isBillingReceipt).ServiceUids(serviceUids).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileCategoriesAPI.GetFileCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileCategories`: IPagedResourceListFileCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `FileCategoriesAPI.GetFileCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Search by file category code | 
 **label** | **string** | Search by file category label | 
 **codes** | **[]string** | Search by file category codes | 
 **labels** | **[]string** | Search by file category labels | 
 **isVisibleForOrderers** | **bool** | Search by file category orderer visibility | 
 **isVisibleForDrivers** | **bool** | Search by file category driver visibility | 
 **isBillingReceipt** | **bool** | Search by file category being billing receipt | 
 **serviceUids** | **[]string** | Search by file category service uids | 
 **pattern** | **string** | A pattern to look for in fields specified by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListFileCategoryDto**](IPagedResourceListFileCategoryDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

