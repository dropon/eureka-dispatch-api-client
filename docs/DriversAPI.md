# \DriversAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDriverByCode**](DriversAPI.md#GetDriverByCode) | **Get** /v3/drivers/by-code | Get a driver by its code
[**GetDriverById**](DriversAPI.md#GetDriverById) | **Get** /v3/drivers/{driverId} | Get a driver by its identifier
[**GetDrivers**](DriversAPI.md#GetDrivers) | **Get** /v3/drivers | Get drivers



## GetDriverByCode

> DriverDto GetDriverByCode(ctx).Code(code).Fields(fields).Execute()

Get a driver by its code

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
	code := "code_example" // string | Driver's code
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversAPI.GetDriverByCode(context.Background()).Code(code).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversAPI.GetDriverByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriverByCode`: DriverDto
	fmt.Fprintf(os.Stdout, "Response from `DriversAPI.GetDriverByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Driver&#39;s code | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**DriverDto**](DriverDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDriverById

> DriverDto GetDriverById(ctx, driverId).Fields(fields).Execute()

Get a driver by its identifier

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
	driverId := int32(56) // int32 | Identifier of the driver
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversAPI.GetDriverById(context.Background(), driverId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversAPI.GetDriverById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDriverById`: DriverDto
	fmt.Fprintf(os.Stdout, "Response from `DriversAPI.GetDriverById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driverId** | **int32** | Identifier of the driver | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDriverByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**DriverDto**](DriverDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrivers

> IPagedResourceListDriverDto GetDrivers(ctx).DriverIds(driverIds).Code(code).Name(name).CountryCode(countryCode).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).AgencyCodes(agencyCodes).TeamCode(teamCode).TeamCodes(teamCodes).IsEnabled(isEnabled).DefaultVehicleCode(defaultVehicleCode).DefaultVehicleLicensePlate(defaultVehicleLicensePlate).DefaultVehicleCodes(defaultVehicleCodes).CurrentVehicleCode(currentVehicleCode).CurrentVehicleLicensePlate(currentVehicleLicensePlate).DriverTypes(driverTypes).DriverLicenseTypes(driverLicenseTypes).MinimumGlobalMark(minimumGlobalMark).MinimumSubcontractorEmployeesCount(minimumSubcontractorEmployeesCount).WorkingTimeServiceStarted(workingTimeServiceStarted).IsDispatchMobileEnabled(isDispatchMobileEnabled).Pattern(pattern).PatternFields(patternFields).DangerousGoodsCertificationClasses(dangerousGoodsCertificationClasses).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get drivers



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
	driverIds := []int32{int32(123)} // []int32 | List of driver Ids (optional)
	code := "code_example" // string | Driver code pattern (optional)
	name := "name_example" // string | Driver name pattern (optional)
	countryCode := "countryCode_example" // string | Driver country code (optional)
	postCode := "postCode_example" // string | Driver post code (optional)
	cityName := "cityName_example" // string | Driver city name (optional)
	agencyCode := "agencyCode_example" // string | Driver agency code (optional)
	agencyCodes := []string{"Inner_example"} // []string | Driver agency codes (optional)
	teamCode := "teamCode_example" // string | Driver team code (optional)
	teamCodes := []string{"Inner_example"} // []string | Driver team codes (optional)
	isEnabled := true // bool |  (optional)
	defaultVehicleCode := "defaultVehicleCode_example" // string | Default vehicle code (optional)
	defaultVehicleLicensePlate := "defaultVehicleLicensePlate_example" // string | Default vehicle license plate (optional)
	defaultVehicleCodes := []string{"Inner_example"} // []string | Default vehicle codes (optional)
	currentVehicleCode := "currentVehicleCode_example" // string | Current vehicle code (optional)
	currentVehicleLicensePlate := "currentVehicleLicensePlate_example" // string | Current vehicle license plate (optional)
	driverTypes := []string{"DriverTypes_example"} // []string | List of driver types (optional)
	driverLicenseTypes := []string{"Inner_example"} // []string | List of driver license type (optional)
	minimumGlobalMark := int32(56) // int32 | Driver included minimum global mark (optional)
	minimumSubcontractorEmployeesCount := int64(789) // int64 | Minimum number of employee the subcontractor have, if specified only subcontractor are returned (optional)
	workingTimeServiceStarted := true // bool |  (optional)
	isDispatchMobileEnabled := true // bool |  (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	dangerousGoodsCertificationClasses := []string{"DangerousGoodsCertificationClasses_example"} // []string | List of dangerous good certification class (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversAPI.GetDrivers(context.Background()).DriverIds(driverIds).Code(code).Name(name).CountryCode(countryCode).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).AgencyCodes(agencyCodes).TeamCode(teamCode).TeamCodes(teamCodes).IsEnabled(isEnabled).DefaultVehicleCode(defaultVehicleCode).DefaultVehicleLicensePlate(defaultVehicleLicensePlate).DefaultVehicleCodes(defaultVehicleCodes).CurrentVehicleCode(currentVehicleCode).CurrentVehicleLicensePlate(currentVehicleLicensePlate).DriverTypes(driverTypes).DriverLicenseTypes(driverLicenseTypes).MinimumGlobalMark(minimumGlobalMark).MinimumSubcontractorEmployeesCount(minimumSubcontractorEmployeesCount).WorkingTimeServiceStarted(workingTimeServiceStarted).IsDispatchMobileEnabled(isDispatchMobileEnabled).Pattern(pattern).PatternFields(patternFields).DangerousGoodsCertificationClasses(dangerousGoodsCertificationClasses).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversAPI.GetDrivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrivers`: IPagedResourceListDriverDto
	fmt.Fprintf(os.Stdout, "Response from `DriversAPI.GetDrivers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDriversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **driverIds** | **[]int32** | List of driver Ids | 
 **code** | **string** | Driver code pattern | 
 **name** | **string** | Driver name pattern | 
 **countryCode** | **string** | Driver country code | 
 **postCode** | **string** | Driver post code | 
 **cityName** | **string** | Driver city name | 
 **agencyCode** | **string** | Driver agency code | 
 **agencyCodes** | **[]string** | Driver agency codes | 
 **teamCode** | **string** | Driver team code | 
 **teamCodes** | **[]string** | Driver team codes | 
 **isEnabled** | **bool** |  | 
 **defaultVehicleCode** | **string** | Default vehicle code | 
 **defaultVehicleLicensePlate** | **string** | Default vehicle license plate | 
 **defaultVehicleCodes** | **[]string** | Default vehicle codes | 
 **currentVehicleCode** | **string** | Current vehicle code | 
 **currentVehicleLicensePlate** | **string** | Current vehicle license plate | 
 **driverTypes** | **[]string** | List of driver types | 
 **driverLicenseTypes** | **[]string** | List of driver license type | 
 **minimumGlobalMark** | **int32** | Driver included minimum global mark | 
 **minimumSubcontractorEmployeesCount** | **int64** | Minimum number of employee the subcontractor have, if specified only subcontractor are returned | 
 **workingTimeServiceStarted** | **bool** |  | 
 **isDispatchMobileEnabled** | **bool** |  | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **dangerousGoodsCertificationClasses** | **[]string** | List of dangerous good certification class | 
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

