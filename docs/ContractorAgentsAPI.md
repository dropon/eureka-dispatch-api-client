# \ContractorAgentsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContractorAgents**](ContractorAgentsAPI.md#GetContractorAgents) | **Get** /v3/contractor-agents | Get contractor agents



## GetContractorAgents

> IPagedResourceListContractorAgentDto GetContractorAgents(ctx).ContractorAgentIds(contractorAgentIds).SubcontractorIds(subcontractorIds).Name(name).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get contractor agents



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
	contractorAgentIds := []int32{int32(123)} // []int32 | List of contractor agent Ids. (optional)
	subcontractorIds := []int32{int32(123)} // []int32 | List of subcontractor Ids (optional)
	name := "name_example" // string | Subcontractor name (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractorAgentsAPI.GetContractorAgents(context.Background()).ContractorAgentIds(contractorAgentIds).SubcontractorIds(subcontractorIds).Name(name).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractorAgentsAPI.GetContractorAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContractorAgents`: IPagedResourceListContractorAgentDto
	fmt.Fprintf(os.Stdout, "Response from `ContractorAgentsAPI.GetContractorAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContractorAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractorAgentIds** | **[]int32** | List of contractor agent Ids. | 
 **subcontractorIds** | **[]int32** | List of subcontractor Ids | 
 **name** | **string** | Subcontractor name | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListContractorAgentDto**](IPagedResourceListContractorAgentDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

