# \VehiclesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVehicles**](VehiclesAPI.md#GetVehicles) | **Get** /v3/vehicles | Get vehicles
[**GetVehiclesSubcontractors**](VehiclesAPI.md#GetVehiclesSubcontractors) | **Get** /v3/vehicles/{uid}/subcontractors | Get the subcontractors linked to the vehicle



## GetVehicles

> IPagedResourceListVehicleDto GetVehicles(ctx).Code(code).Codes(codes).TypeCode(typeCode).TypeCodes(typeCodes).LicensePlate(licensePlate).Brand(brand).Label(label).IsEnabled(isEnabled).AgencyCode(agencyCode).AgencyCodes(agencyCodes).HasDangerousGoodsCase(hasDangerousGoodsCase).HasHatchback(hasHatchback).IsSubcontractorsSpecific(isSubcontractorsSpecific).SubcontractorCode(subcontractorCode).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get vehicles



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
	code := "code_example" // string | Vehicle code pattern (optional)
	codes := []string{"Inner_example"} // []string | Vehicle codes (optional)
	typeCode := "typeCode_example" // string | Vehicle type code (optional)
	typeCodes := []string{"Inner_example"} // []string | Vehicle type codes (optional)
	licensePlate := "licensePlate_example" // string | Vehicle license plate (optional)
	brand := "brand_example" // string | Vehicle brand (optional)
	label := "label_example" // string | Vehicle label pattern (optional)
	isEnabled := true // bool |  (optional)
	agencyCode := "agencyCode_example" // string | Vehicle agency code (optional)
	agencyCodes := []string{"Inner_example"} // []string | Vehicle agency codes (optional)
	hasDangerousGoodsCase := true // bool | Defines whether the vehicle is habilitated to transport dangerous good (optional)
	hasHatchback := true // bool | Defines whether the vehicle has a hatchback (optional)
	isSubcontractorsSpecific := true // bool | Defines whether the vehicle is specific to subcontractors (optional)
	subcontractorCode := "subcontractorCode_example" // string | Code of a subcontractor specific to vehicles (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VehiclesAPI.GetVehicles(context.Background()).Code(code).Codes(codes).TypeCode(typeCode).TypeCodes(typeCodes).LicensePlate(licensePlate).Brand(brand).Label(label).IsEnabled(isEnabled).AgencyCode(agencyCode).AgencyCodes(agencyCodes).HasDangerousGoodsCase(hasDangerousGoodsCase).HasHatchback(hasHatchback).IsSubcontractorsSpecific(isSubcontractorsSpecific).SubcontractorCode(subcontractorCode).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VehiclesAPI.GetVehicles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehicles`: IPagedResourceListVehicleDto
	fmt.Fprintf(os.Stdout, "Response from `VehiclesAPI.GetVehicles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVehiclesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Vehicle code pattern | 
 **codes** | **[]string** | Vehicle codes | 
 **typeCode** | **string** | Vehicle type code | 
 **typeCodes** | **[]string** | Vehicle type codes | 
 **licensePlate** | **string** | Vehicle license plate | 
 **brand** | **string** | Vehicle brand | 
 **label** | **string** | Vehicle label pattern | 
 **isEnabled** | **bool** |  | 
 **agencyCode** | **string** | Vehicle agency code | 
 **agencyCodes** | **[]string** | Vehicle agency codes | 
 **hasDangerousGoodsCase** | **bool** | Defines whether the vehicle is habilitated to transport dangerous good | 
 **hasHatchback** | **bool** | Defines whether the vehicle has a hatchback | 
 **isSubcontractorsSpecific** | **bool** | Defines whether the vehicle is specific to subcontractors | 
 **subcontractorCode** | **string** | Code of a subcontractor specific to vehicles | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListVehicleDto**](IPagedResourceListVehicleDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVehiclesSubcontractors

> IPagedResourceListDriverDto GetVehiclesSubcontractors(ctx, uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get the subcontractors linked to the vehicle



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
	uid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VehiclesAPI.GetVehiclesSubcontractors(context.Background(), uid).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VehiclesAPI.GetVehiclesSubcontractors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVehiclesSubcontractors`: IPagedResourceListDriverDto
	fmt.Fprintf(os.Stdout, "Response from `VehiclesAPI.GetVehiclesSubcontractors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVehiclesSubcontractorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListDriverDto**](IPagedResourceListDriverDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

