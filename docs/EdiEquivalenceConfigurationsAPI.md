# \EdiEquivalenceConfigurationsAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomerEquivalencesConfigurationByUid**](EdiEquivalenceConfigurationsAPI.md#GetCustomerEquivalencesConfigurationByUid) | **Get** /v3/edi-equivalences-configurations/customer-type-configurations/{uid} | Get a customer equivalences configuration by its unique identifier
[**GetCustomerEquivalencesConfigurations**](EdiEquivalenceConfigurationsAPI.md#GetCustomerEquivalencesConfigurations) | **Get** /v3/edi-equivalences-configurations/customer-type-configurations | Get customer equivalences configurations
[**GetPackageNatureEquivalencesConfigurationByUid**](EdiEquivalenceConfigurationsAPI.md#GetPackageNatureEquivalencesConfigurationByUid) | **Get** /v3/edi-equivalences-configurations/package-nature-type-configurations/{uid} | Get a package nature equivalences configuration by its unique identifier
[**GetPackageNatureEquivalencesConfigurations**](EdiEquivalenceConfigurationsAPI.md#GetPackageNatureEquivalencesConfigurations) | **Get** /v3/edi-equivalences-configurations/package-nature-type-configurations | Get package nature equivalences configurations
[**GetServiceEquivalencesConfigurationByUid**](EdiEquivalenceConfigurationsAPI.md#GetServiceEquivalencesConfigurationByUid) | **Get** /v3/edi-equivalences-configurations/service-type-configurations/{uid} | Get a service equivalences configuration by its unique identifier
[**GetServiceEquivalencesConfigurations**](EdiEquivalenceConfigurationsAPI.md#GetServiceEquivalencesConfigurations) | **Get** /v3/edi-equivalences-configurations/service-type-configurations | Get service equivalences configurations
[**GetSubstateEquivalencesConfigurationByUid**](EdiEquivalenceConfigurationsAPI.md#GetSubstateEquivalencesConfigurationByUid) | **Get** /v3/edi-equivalences-configurations/substate-type-configurations/{uid} | Get a substate equivalences configuration by its unique identifier
[**GetSubstateEquivalencesConfigurations**](EdiEquivalenceConfigurationsAPI.md#GetSubstateEquivalencesConfigurations) | **Get** /v3/edi-equivalences-configurations/substate-type-configurations | Get substate equivalences configurations



## GetCustomerEquivalencesConfigurationByUid

> IResourceListCustomerEquivalenceDto GetCustomerEquivalencesConfigurationByUid(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get a customer equivalences configuration by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the customer equivalences configuration
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurationByUid(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurationByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEquivalencesConfigurationByUid`: IResourceListCustomerEquivalenceDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurationByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the customer equivalences configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEquivalencesConfigurationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IResourceListCustomerEquivalenceDto**](IResourceListCustomerEquivalenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerEquivalencesConfigurations

> IPagedResourceListEdiEquivalenceConfigurationDto GetCustomerEquivalencesConfigurations(ctx).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get customer equivalences configurations



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
	uids := []string{"Inner_example"} // []string | Equivalence configuration unique identifiers (optional)
	codes := []string{"Inner_example"} // []string | Equivalence configuration codes (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurations(context.Background()).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerEquivalencesConfigurations`: IPagedResourceListEdiEquivalenceConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetCustomerEquivalencesConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerEquivalencesConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **[]string** | Equivalence configuration unique identifiers | 
 **codes** | **[]string** | Equivalence configuration codes | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListEdiEquivalenceConfigurationDto**](IPagedResourceListEdiEquivalenceConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageNatureEquivalencesConfigurationByUid

> IResourceListPackageNatureEquivalenceDto GetPackageNatureEquivalencesConfigurationByUid(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get a package nature equivalences configuration by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the package nature equivalences configuration
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurationByUid(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurationByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatureEquivalencesConfigurationByUid`: IResourceListPackageNatureEquivalenceDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurationByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the package nature equivalences configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNatureEquivalencesConfigurationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IResourceListPackageNatureEquivalenceDto**](IResourceListPackageNatureEquivalenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageNatureEquivalencesConfigurations

> IPagedResourceListEdiEquivalenceConfigurationDto GetPackageNatureEquivalencesConfigurations(ctx).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get package nature equivalences configurations



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
	uids := []string{"Inner_example"} // []string | Equivalence configuration unique identifiers (optional)
	codes := []string{"Inner_example"} // []string | Equivalence configuration codes (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurations(context.Background()).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageNatureEquivalencesConfigurations`: IPagedResourceListEdiEquivalenceConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetPackageNatureEquivalencesConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageNatureEquivalencesConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **[]string** | Equivalence configuration unique identifiers | 
 **codes** | **[]string** | Equivalence configuration codes | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListEdiEquivalenceConfigurationDto**](IPagedResourceListEdiEquivalenceConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceEquivalencesConfigurationByUid

> IResourceListServiceEquivalenceDto GetServiceEquivalencesConfigurationByUid(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get a service equivalences configuration by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the service equivalences configuration
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurationByUid(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurationByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceEquivalencesConfigurationByUid`: IResourceListServiceEquivalenceDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurationByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the service equivalences configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceEquivalencesConfigurationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IResourceListServiceEquivalenceDto**](IResourceListServiceEquivalenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceEquivalencesConfigurations

> IPagedResourceListEdiEquivalenceConfigurationDto GetServiceEquivalencesConfigurations(ctx).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get service equivalences configurations



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
	uids := []string{"Inner_example"} // []string | Equivalence configuration unique identifiers (optional)
	codes := []string{"Inner_example"} // []string | Equivalence configuration codes (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurations(context.Background()).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceEquivalencesConfigurations`: IPagedResourceListEdiEquivalenceConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetServiceEquivalencesConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceEquivalencesConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **[]string** | Equivalence configuration unique identifiers | 
 **codes** | **[]string** | Equivalence configuration codes | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListEdiEquivalenceConfigurationDto**](IPagedResourceListEdiEquivalenceConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubstateEquivalencesConfigurationByUid

> IResourceListSubstateEquivalenceDto GetSubstateEquivalencesConfigurationByUid(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get a substate equivalences configuration by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique identifier of the substate equivalences configuration
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurationByUid(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurationByUid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubstateEquivalencesConfigurationByUid`: IResourceListSubstateEquivalenceDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurationByUid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Unique identifier of the substate equivalences configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubstateEquivalencesConfigurationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IResourceListSubstateEquivalenceDto**](IResourceListSubstateEquivalenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubstateEquivalencesConfigurations

> IPagedResourceListEdiEquivalenceConfigurationDto GetSubstateEquivalencesConfigurations(ctx).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get substate equivalences configurations



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
	uids := []string{"Inner_example"} // []string | Equivalence configuration unique identifiers (optional)
	codes := []string{"Inner_example"} // []string | Equivalence configuration codes (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurations(context.Background()).Uids(uids).Codes(codes).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubstateEquivalencesConfigurations`: IPagedResourceListEdiEquivalenceConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `EdiEquivalenceConfigurationsAPI.GetSubstateEquivalencesConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubstateEquivalencesConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uids** | **[]string** | Equivalence configuration unique identifiers | 
 **codes** | **[]string** | Equivalence configuration codes | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListEdiEquivalenceConfigurationDto**](IPagedResourceListEdiEquivalenceConfigurationDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

