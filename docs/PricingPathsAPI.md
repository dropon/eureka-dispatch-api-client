# \PricingPathsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPricingPathAdditionalUnits**](PricingPathsAPI.md#GetPricingPathAdditionalUnits) | **Get** /v3/pricing-paths/{pricingPathId}/sub-service-additional-units | Get pricing path additional units
[**GetPricingPathById**](PricingPathsAPI.md#GetPricingPathById) | **Get** /v3/pricing-paths/{pricingPathId} | Get pricing path by id
[**GetPricingPathSubServiceUnitPricingSlots**](PricingPathsAPI.md#GetPricingPathSubServiceUnitPricingSlots) | **Get** /v3/pricing-paths/{pricingPathId}/sub-service-unit-pricing-slots | Get pricing path sub service unit pricing slots
[**GetPricingPathSubServices**](PricingPathsAPI.md#GetPricingPathSubServices) | **Get** /v3/pricing-paths/{pricingPathId}/sub-services | Get pricing path sub services
[**GetPricingPathUnitPricingSlotEquivalences**](PricingPathsAPI.md#GetPricingPathUnitPricingSlotEquivalences) | **Get** /v3/pricing-paths/{pricingPathId}/unit-pricing-slot-equivalences | Get pricing path unit pricing slot equivalences
[**GetPricingPaths**](PricingPathsAPI.md#GetPricingPaths) | **Get** /v3/pricing-paths | Get pricing paths



## GetPricingPathAdditionalUnits

> IPagedResourceListPricingPathSubServiceAdditionalUnitDto GetPricingPathAdditionalUnits(ctx, pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get pricing path additional units



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
	pricingPathId := int32(56) // int32 | Identifier of the pricing path
	subServiceCode := "subServiceCode_example" // string | Sub service code (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPathAdditionalUnits(context.Background(), pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPathAdditionalUnits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPathAdditionalUnits`: IPagedResourceListPricingPathSubServiceAdditionalUnitDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPathAdditionalUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingPathId** | **int32** | Identifier of the pricing path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathAdditionalUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subServiceCode** | **string** | Sub service code | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPricingPathSubServiceAdditionalUnitDto**](IPagedResourceListPricingPathSubServiceAdditionalUnitDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPathById

> PricingPathDto GetPricingPathById(ctx, pricingPathId).Fields(fields).Execute()

Get pricing path by id

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
	pricingPathId := int32(56) // int32 | 
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPathById(context.Background(), pricingPathId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPathById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPathById`: PricingPathDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPathById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingPathId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**PricingPathDto**](PricingPathDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPathSubServiceUnitPricingSlots

> IPagedResourceListPricingPathSubServiceUnitPricingSlotDto GetPricingPathSubServiceUnitPricingSlots(ctx, pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get pricing path sub service unit pricing slots



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
	pricingPathId := int32(56) // int32 | Identifier of the pricing path
	subServiceCode := "subServiceCode_example" // string | Sub service code (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPathSubServiceUnitPricingSlots(context.Background(), pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPathSubServiceUnitPricingSlots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPathSubServiceUnitPricingSlots`: IPagedResourceListPricingPathSubServiceUnitPricingSlotDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPathSubServiceUnitPricingSlots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingPathId** | **int32** | Identifier of the pricing path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathSubServiceUnitPricingSlotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subServiceCode** | **string** | Sub service code | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPricingPathSubServiceUnitPricingSlotDto**](IPagedResourceListPricingPathSubServiceUnitPricingSlotDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPathSubServices

> IPagedResourceListPricingPathSubServiceDto GetPricingPathSubServices(ctx, pricingPathId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get pricing path sub services



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
	pricingPathId := int32(56) // int32 | Identifier of the pricing path
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPathSubServices(context.Background(), pricingPathId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPathSubServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPathSubServices`: IPagedResourceListPricingPathSubServiceDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPathSubServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingPathId** | **int32** | Identifier of the pricing path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathSubServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPricingPathSubServiceDto**](IPagedResourceListPricingPathSubServiceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPathUnitPricingSlotEquivalences

> IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto GetPricingPathUnitPricingSlotEquivalences(ctx, pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get pricing path unit pricing slot equivalences



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
	pricingPathId := int32(56) // int32 | Identifier of the pricing path
	subServiceCode := "subServiceCode_example" // string | Sub service code (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPathUnitPricingSlotEquivalences(context.Background(), pricingPathId).SubServiceCode(subServiceCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPathUnitPricingSlotEquivalences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPathUnitPricingSlotEquivalences`: IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPathUnitPricingSlotEquivalences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingPathId** | **int32** | Identifier of the pricing path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathUnitPricingSlotEquivalencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subServiceCode** | **string** | Sub service code | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto**](IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPaths

> IPagedResourceListPricingPathDto GetPricingPaths(ctx).PriceCatalogCode(priceCatalogCode).ServiceCode(serviceCode).CustomerCode(customerCode).SubcontractorCode(subcontractorCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get pricing paths



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
	priceCatalogCode := "priceCatalogCode_example" // string | Prices catalog code pattern (optional)
	serviceCode := "serviceCode_example" // string | Service code pattern (optional)
	customerCode := "customerCode_example" // string | Customer code pattern (optional)
	subcontractorCode := "subcontractorCode_example" // string | Sub contractor code pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingPathsAPI.GetPricingPaths(context.Background()).PriceCatalogCode(priceCatalogCode).ServiceCode(serviceCode).CustomerCode(customerCode).SubcontractorCode(subcontractorCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPathsAPI.GetPricingPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPaths`: IPagedResourceListPricingPathDto
	fmt.Fprintf(os.Stdout, "Response from `PricingPathsAPI.GetPricingPaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceCatalogCode** | **string** | Prices catalog code pattern | 
 **serviceCode** | **string** | Service code pattern | 
 **customerCode** | **string** | Customer code pattern | 
 **subcontractorCode** | **string** | Sub contractor code pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListPricingPathDto**](IPagedResourceListPricingPathDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

