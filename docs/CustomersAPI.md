# \CustomersAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomParameterValues**](CustomersAPI.md#GetCustomParameterValues) | **Get** /v3/customers/{uid}/custom-parameters/{parameterId}/values | Get custom parameter values
[**GetCustomParameters**](CustomersAPI.md#GetCustomParameters) | **Get** /v3/customers/{uid}/custom-parameters | Get custom parameters
[**GetCustomerByCode**](CustomersAPI.md#GetCustomerByCode) | **Get** /v3/customers/by-customer-code | Get a customer by its code
[**GetCustomerById**](CustomersAPI.md#GetCustomerById) | **Get** /v3/customers/{uid} | Get a customer by its unique identifier
[**GetCustomerReferences**](CustomersAPI.md#GetCustomerReferences) | **Get** /v3/customers/{uid}/references | Get customer references
[**GetCustomers**](CustomersAPI.md#GetCustomers) | **Get** /v3/customers | Get customers



## GetCustomParameterValues

> IPagedResourceListCustomParameterValueDto GetCustomParameterValues(ctx, uid, parameterId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get custom parameter values



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer's unique identifier.
	parameterId := int32(56) // int32 | Identifier of the custom parameter
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomParameterValues(context.Background(), uid, parameterId).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomParameterValues`: IPagedResourceListCustomParameterValueDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Customer&#39;s unique identifier. | 
**parameterId** | **int32** | Identifier of the custom parameter | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomParameterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListCustomParameterValueDto**](IPagedResourceListCustomParameterValueDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomParameters

> IPagedResourceListCustomParameterDto GetCustomParameters(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get custom parameters



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer's unique identifier.
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomParameters(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomParameters`: IPagedResourceListCustomParameterDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Customer&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListCustomParameterDto**](IPagedResourceListCustomParameterDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerByCode

> CustomerDto GetCustomerByCode(ctx).CustomerCode(customerCode).Fields(fields).Execute()

Get a customer by its code

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
	customerCode := "customerCode_example" // string | Customer's code.
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomerByCode(context.Background()).CustomerCode(customerCode).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerByCode`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerCode** | **string** | Customer&#39;s code. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerById

> CustomerDto GetCustomerById(ctx, uid).Fields(fields).Execute()

Get a customer by its unique identifier

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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer's unique identifier.
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomerById(context.Background(), uid).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerById`: CustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Customer&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**CustomerDto**](CustomerDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerReferences

> IPagedResourceListCustomerReferenceDto GetCustomerReferences(ctx, uid).ReferenceLevel(referenceLevel).Value(value).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get customer references



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer's unique identifier.
	referenceLevel := int32(56) // int32 | Optional level of the references to get (optional)
	value := "value_example" // string | Pattern of the reference value (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomerReferences(context.Background(), uid).ReferenceLevel(referenceLevel).Value(value).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerReferences`: IPagedResourceListCustomerReferenceDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | Customer&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **referenceLevel** | **int32** | Optional level of the references to get | 
 **value** | **string** | Pattern of the reference value | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListCustomerReferenceDto**](IPagedResourceListCustomerReferenceDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomers

> IPagedResourceListCustomerDto GetCustomers(ctx).Code(code).Label(label).PostCode(postCode).CityName(cityName).CountryCode(countryCode).FamilyCode(familyCode).AgencyCode(agencyCode).AgencyCodes(agencyCodes).CustomerUids(customerUids).CustomerCodes(customerCodes).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get customers



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
	code := "code_example" // string | Customer code pattern (optional)
	label := "label_example" // string | Customer label pattern (optional)
	postCode := "postCode_example" // string | Customer operation address post code pattern (optional)
	cityName := "cityName_example" // string | Customer operation address city name pattern (optional)
	countryCode := "countryCode_example" // string | Customer operation address country code pattern (optional)
	familyCode := "familyCode_example" // string | Customer family code pattern (optional)
	agencyCode := "agencyCode_example" // string | Customer agency code pattern (optional)
	agencyCodes := []string{"Inner_example"} // []string | A list of exact agency codes (optional)
	customerUids := []string{"Inner_example"} // []string | Collection of customer unique identifiers (optional)
	customerCodes := []string{"Inner_example"} // []string | A list of exact customer codes (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomers(context.Background()).Code(code).Label(label).PostCode(postCode).CityName(cityName).CountryCode(countryCode).FamilyCode(familyCode).AgencyCode(agencyCode).AgencyCodes(agencyCodes).CustomerUids(customerUids).CustomerCodes(customerCodes).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomers`: IPagedResourceListCustomerDto
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Customer code pattern | 
 **label** | **string** | Customer label pattern | 
 **postCode** | **string** | Customer operation address post code pattern | 
 **cityName** | **string** | Customer operation address city name pattern | 
 **countryCode** | **string** | Customer operation address country code pattern | 
 **familyCode** | **string** | Customer family code pattern | 
 **agencyCode** | **string** | Customer agency code pattern | 
 **agencyCodes** | **[]string** | A list of exact agency codes | 
 **customerUids** | **[]string** | Collection of customer unique identifiers | 
 **customerCodes** | **[]string** | A list of exact customer codes | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListCustomerDto**](IPagedResourceListCustomerDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

