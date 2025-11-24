# \CitiesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBestMatch**](CitiesAPI.md#GetBestMatch) | **Get** /v3/cities/best-match | Get best match
[**GetCities**](CitiesAPI.md#GetCities) | **Get** /v3/cities | Get cities
[**GetCityById**](CitiesAPI.md#GetCityById) | **Get** /v3/cities/by-city-id | Get a city by its identifier



## GetBestMatch

> CityDto GetBestMatch(ctx).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).CountryCode(countryCode).CountryFamilyCode(countryFamilyCode).MissionEntryAvailability(missionEntryAvailability).Fields(fields).Execute()

Get best match



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
	postCode := "postCode_example" // string | The post code to look for. (optional)
	cityName := "cityName_example" // string | The city name to look for. (optional)
	agencyCode := "agencyCode_example" // string | Optional : Restrict research to cities in the specified agency,  or that are not linked to a specific agency.  A city having the correct agency code will have priority over  other cities with no specific agency. (optional)
	countryCode := "countryCode_example" // string | Optional : Restrict research to cities in the specified country.  If no country code is specified, a city from the default Dispatch country code  will have priority over the others. (optional)
	countryFamilyCode := "countryFamilyCode_example" // string | Optional : Restrict to cities in the specified country family.  Used only if CountryCode is not provided. (optional)
	missionEntryAvailability := "missionEntryAvailability_example" // string | Restrict to cities that are available or not for mission entry.    Default value is 'Yes'. (optional) (default to "Yes")
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CitiesAPI.GetBestMatch(context.Background()).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).CountryCode(countryCode).CountryFamilyCode(countryFamilyCode).MissionEntryAvailability(missionEntryAvailability).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CitiesAPI.GetBestMatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBestMatch`: CityDto
	fmt.Fprintf(os.Stdout, "Response from `CitiesAPI.GetBestMatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBestMatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCode** | **string** | The post code to look for. | 
 **cityName** | **string** | The city name to look for. | 
 **agencyCode** | **string** | Optional : Restrict research to cities in the specified agency,  or that are not linked to a specific agency.  A city having the correct agency code will have priority over  other cities with no specific agency. | 
 **countryCode** | **string** | Optional : Restrict research to cities in the specified country.  If no country code is specified, a city from the default Dispatch country code  will have priority over the others. | 
 **countryFamilyCode** | **string** | Optional : Restrict to cities in the specified country family.  Used only if CountryCode is not provided. | 
 **missionEntryAvailability** | **string** | Restrict to cities that are available or not for mission entry.    Default value is &#39;Yes&#39;. | [default to &quot;Yes&quot;]
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**CityDto**](CityDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCities

> IPagedResourceListCityDto GetCities(ctx).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).CountryCode(countryCode).CountryFamilyCode(countryFamilyCode).MissionEntryAvailability(missionEntryAvailability).CityIds(cityIds).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get cities



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
	postCode := "postCode_example" // string | The postal code pattern to look for. (optional)
	cityName := "cityName_example" // string | The city name pattern to look for. (optional)
	agencyCode := "agencyCode_example" // string | Optional : Restrict research to cities in the specified agency,  or that are not linked to a specific agency. (optional)
	countryCode := "countryCode_example" // string | Optional : Restrict research to cities in the specified country. (optional)
	countryFamilyCode := "countryFamilyCode_example" // string | Optional : Restrict to cities in the specified country family.  Used only if CountryCode is not provided. (optional)
	missionEntryAvailability := "missionEntryAvailability_example" // string | Restrict to cities that are available or not for mission entry.    Default value is 'Yes'. (optional) (default to "Yes")
	cityIds := []int32{int32(123)} // []int32 | A list of city ids (optional)
	pattern := "pattern_example" // string | A pattern to look for in fields specified  by PatternFields. (optional)
	patternFields := []string{"PatternFields_example"} // []string | Fields in which to search for Pattern (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CitiesAPI.GetCities(context.Background()).PostCode(postCode).CityName(cityName).AgencyCode(agencyCode).CountryCode(countryCode).CountryFamilyCode(countryFamilyCode).MissionEntryAvailability(missionEntryAvailability).CityIds(cityIds).Pattern(pattern).PatternFields(patternFields).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CitiesAPI.GetCities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCities`: IPagedResourceListCityDto
	fmt.Fprintf(os.Stdout, "Response from `CitiesAPI.GetCities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCode** | **string** | The postal code pattern to look for. | 
 **cityName** | **string** | The city name pattern to look for. | 
 **agencyCode** | **string** | Optional : Restrict research to cities in the specified agency,  or that are not linked to a specific agency. | 
 **countryCode** | **string** | Optional : Restrict research to cities in the specified country. | 
 **countryFamilyCode** | **string** | Optional : Restrict to cities in the specified country family.  Used only if CountryCode is not provided. | 
 **missionEntryAvailability** | **string** | Restrict to cities that are available or not for mission entry.    Default value is &#39;Yes&#39;. | [default to &quot;Yes&quot;]
 **cityIds** | **[]int32** | A list of city ids | 
 **pattern** | **string** | A pattern to look for in fields specified  by PatternFields. | 
 **patternFields** | **[]string** | Fields in which to search for Pattern | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListCityDto**](IPagedResourceListCityDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCityById

> CityDto GetCityById(ctx).CityId(cityId).Fields(fields).Execute()

Get a city by its identifier

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
	cityId := int32(56) // int32 | 
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CitiesAPI.GetCityById(context.Background()).CityId(cityId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CitiesAPI.GetCityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCityById`: CityDto
	fmt.Fprintf(os.Stdout, "Response from `CitiesAPI.GetCityById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cityId** | **int32** |  | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**CityDto**](CityDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

