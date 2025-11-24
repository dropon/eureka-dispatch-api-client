# \SubstatesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubstateByCode**](SubstatesAPI.md#GetSubstateByCode) | **Get** /v3/substates/by-code | Get a substate by its code
[**GetSubstates**](SubstatesAPI.md#GetSubstates) | **Get** /v3/substates | Get substates



## GetSubstateByCode

> SubStateDto GetSubstateByCode(ctx).Code(code).Fields(fields).Execute()

Get a substate by its code

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
	code := "code_example" // string | Substate's code
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatesAPI.GetSubstateByCode(context.Background()).Code(code).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatesAPI.GetSubstateByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubstateByCode`: SubStateDto
	fmt.Fprintf(os.Stdout, "Response from `SubstatesAPI.GetSubstateByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubstateByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Substate&#39;s code | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**SubStateDto**](SubStateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubstates

> IPagedResourceListSubStateDto GetSubstates(ctx).SubStateCodes(subStateCodes).SubStateFamilyCodes(subStateFamilyCodes).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get substates



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
	subStateCodes := []string{"Inner_example"} // []string | Substate unique codes (optional)
	subStateFamilyCodes := []string{"Inner_example"} // []string | Substate family unique codes (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatesAPI.GetSubstates(context.Background()).SubStateCodes(subStateCodes).SubStateFamilyCodes(subStateFamilyCodes).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatesAPI.GetSubstates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubstates`: IPagedResourceListSubStateDto
	fmt.Fprintf(os.Stdout, "Response from `SubstatesAPI.GetSubstates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubstatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subStateCodes** | **[]string** | Substate unique codes | 
 **subStateFamilyCodes** | **[]string** | Substate family unique codes | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListSubStateDto**](IPagedResourceListSubStateDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

