# \PackageNaturesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPackageNatureById**](PackageNaturesAPI.md#GetPackageNatureById) | **Get** /v3/package-natures/{packageNatureId} | Get package nature by id
[**GetPackageNatureReferenceValues**](PackageNaturesAPI.md#GetPackageNatureReferenceValues) | **Get** /v3/package-natures/{packageNatureId}/references/{referenceIndex}/values | Get package nature reference values
[**GetPackageNatureReferences**](PackageNaturesAPI.md#GetPackageNatureReferences) | **Get** /v3/package-natures/{packageNatureId}/references | Get package nature references
[**GetPackageNatures**](PackageNaturesAPI.md#GetPackageNatures) | **Get** /v3/package-natures | Get package natures



## GetPackageNatureById

> PackageNatureDto GetPackageNatureById(ctx, packageNatureId).Fields(fields).Execute()

Get package nature by id



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
	packageNatureId := int32(56) // int32 | Package nature's identifier.
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageNaturesAPI.GetPackageNatureById(context.Background(), packageNatureId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageNaturesAPI.GetPackageNatureById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatureById`: PackageNatureDto
	fmt.Fprintf(os.Stdout, "Response from `PackageNaturesAPI.GetPackageNatureById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageNatureId** | **int32** | Package nature&#39;s identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNatureByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**PackageNatureDto**](PackageNatureDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageNatureReferenceValues

> IPagedResourceListPackageNatureReferenceValueDto GetPackageNatureReferenceValues(ctx, packageNatureId, referenceIndex).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get package nature reference values



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
	packageNatureId := int32(56) // int32 | Package nature's identifier.
	referenceIndex := int32(56) // int32 | Index of the reference in the package nature.
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageNaturesAPI.GetPackageNatureReferenceValues(context.Background(), packageNatureId, referenceIndex).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageNaturesAPI.GetPackageNatureReferenceValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatureReferenceValues`: IPagedResourceListPackageNatureReferenceValueDto
	fmt.Fprintf(os.Stdout, "Response from `PackageNaturesAPI.GetPackageNatureReferenceValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageNatureId** | **int32** | Package nature&#39;s identifier. | 
**referenceIndex** | **int32** | Index of the reference in the package nature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNatureReferenceValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPackageNatureReferenceValueDto**](IPagedResourceListPackageNatureReferenceValueDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageNatureReferences

> IPagedResourceListPackageNatureReferenceDto GetPackageNatureReferences(ctx, packageNatureId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get package nature references



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
	packageNatureId := int32(56) // int32 | Package nature's identifier.
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageNaturesAPI.GetPackageNatureReferences(context.Background(), packageNatureId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageNaturesAPI.GetPackageNatureReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatureReferences`: IPagedResourceListPackageNatureReferenceDto
	fmt.Fprintf(os.Stdout, "Response from `PackageNaturesAPI.GetPackageNatureReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageNatureId** | **int32** | Package nature&#39;s identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNatureReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPackageNatureReferenceDto**](IPagedResourceListPackageNatureReferenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageNatures

> IPagedResourceListPackageNatureDto GetPackageNatures(ctx).Code(code).Label(label).IsEnabled(isEnabled).PackageFamilyCode(packageFamilyCode).CustomerUids(customerUids).ParentPackageNatureIds(parentPackageNatureIds).Labels(labels).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get package natures



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
	code := "code_example" // string | Package nature code pattern (optional)
	label := "label_example" // string | Package nature label pattern (optional)
	isEnabled := true // bool |  (optional)
	packageFamilyCode := "packageFamilyCode_example" // string | Package nature family code pattern (optional)
	customerUids := []string{"Inner_example"} // []string | Collection of customer unique identifiers (optional)
	parentPackageNatureIds := []int32{int32(123)} // []int32 | Collection of parent package nature identifiers (optional)
	labels := []string{"Inner_example"} // []string | Collection of package nature labels (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageNaturesAPI.GetPackageNatures(context.Background()).Code(code).Label(label).IsEnabled(isEnabled).PackageFamilyCode(packageFamilyCode).CustomerUids(customerUids).ParentPackageNatureIds(parentPackageNatureIds).Labels(labels).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageNaturesAPI.GetPackageNatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatures`: IPagedResourceListPackageNatureDto
	fmt.Fprintf(os.Stdout, "Response from `PackageNaturesAPI.GetPackageNatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Package nature code pattern | 
 **label** | **string** | Package nature label pattern | 
 **isEnabled** | **bool** |  | 
 **packageFamilyCode** | **string** | Package nature family code pattern | 
 **customerUids** | **[]string** | Collection of customer unique identifiers | 
 **parentPackageNatureIds** | **[]int32** | Collection of parent package nature identifiers | 
 **labels** | **[]string** | Collection of package nature labels | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPackageNatureDto**](IPagedResourceListPackageNatureDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

