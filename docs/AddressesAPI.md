# \AddressesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressById**](AddressesAPI.md#GetAddressById) | **Get** /v3/addresses/by-address-id | Get an address by its identifier
[**GetAddresses**](AddressesAPI.md#GetAddresses) | **Get** /v3/addresses | Get addresses
[**GetStreetsSuggestions**](AddressesAPI.md#GetStreetsSuggestions) | **Get** /v3/addresses/streets/suggestions | Get streets suggestions



## GetAddressById

> AddressDto GetAddressById(ctx).AddressId(addressId).Fields(fields).Execute()

Get an address by its identifier

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
	addressId := int32(56) // int32 | Identifier of the address
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressById(context.Background()).AddressId(addressId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressById`: AddressDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressId** | **int32** | Identifier of the address | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**AddressDto**](AddressDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddresses

> IPagedResourceListAddressDto GetAddresses(ctx).AddressesIds(addressesIds).AddressName(addressName).StreetNumber(streetNumber).StreetName(streetName).PostCode(postCode).CityName(cityName).CountryCode(countryCode).ExternalAddressCode(externalAddressCode).AgencyCode(agencyCode).SectorCode(sectorCode).OperationZoneCode(operationZoneCode).OnlyFavoritesOfCustomerCode(onlyFavoritesOfCustomerCode).RestrictToAgencyCode(restrictToAgencyCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get addresses



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
	addressesIds := []int32{int32(123)} // []int32 | Get addresses having the specified ids (optional)
	addressName := "addressName_example" // string | Get addresses matching the specified address name pattern.  This field supports _startsWith_, _endsWith_ and _contains_ patterns. (optional)
	streetNumber := "streetNumber_example" // string | Get addresses matching the specified street number pattern (optional)
	streetName := "streetName_example" // string | Get addresses matching the specified street name pattern (optional)
	postCode := "postCode_example" // string | Get addresses matching the specified post code pattern (optional)
	cityName := "cityName_example" // string | Get addresses matching the specified city name pattern (optional)
	countryCode := "countryCode_example" // string | Get addresses matching the specified country code pattern (optional)
	externalAddressCode := "externalAddressCode_example" // string | Get addresses matching this specified external address code pattern (optional)
	agencyCode := "agencyCode_example" // string | Get addresses matching the specified agency code pattern (optional)
	sectorCode := "sectorCode_example" // string | Get addresses matching the specified sector code pattern (optional)
	operationZoneCode := "operationZoneCode_example" // string | Get addresses matching the specified operation zone code pattern (optional)
	onlyFavoritesOfCustomerCode := "onlyFavoritesOfCustomerCode_example" // string | If set, only the addresses that are favorites of the specified customer  will be retrieved. (optional)
	restrictToAgencyCode := "restrictToAgencyCode_example" // string | If this filter is specified, the results will be filtered to addresses that are either :  - Available from any agency.  - Available only from the specific agency (the restriction can be directly on the address, or in the city) (optional)
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddresses(context.Background()).AddressesIds(addressesIds).AddressName(addressName).StreetNumber(streetNumber).StreetName(streetName).PostCode(postCode).CityName(cityName).CountryCode(countryCode).ExternalAddressCode(externalAddressCode).AgencyCode(agencyCode).SectorCode(sectorCode).OperationZoneCode(operationZoneCode).OnlyFavoritesOfCustomerCode(onlyFavoritesOfCustomerCode).RestrictToAgencyCode(restrictToAgencyCode).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddresses`: IPagedResourceListAddressDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressesIds** | **[]int32** | Get addresses having the specified ids | 
 **addressName** | **string** | Get addresses matching the specified address name pattern.  This field supports _startsWith_, _endsWith_ and _contains_ patterns. | 
 **streetNumber** | **string** | Get addresses matching the specified street number pattern | 
 **streetName** | **string** | Get addresses matching the specified street name pattern | 
 **postCode** | **string** | Get addresses matching the specified post code pattern | 
 **cityName** | **string** | Get addresses matching the specified city name pattern | 
 **countryCode** | **string** | Get addresses matching the specified country code pattern | 
 **externalAddressCode** | **string** | Get addresses matching this specified external address code pattern | 
 **agencyCode** | **string** | Get addresses matching the specified agency code pattern | 
 **sectorCode** | **string** | Get addresses matching the specified sector code pattern | 
 **operationZoneCode** | **string** | Get addresses matching the specified operation zone code pattern | 
 **onlyFavoritesOfCustomerCode** | **string** | If set, only the addresses that are favorites of the specified customer  will be retrieved. | 
 **restrictToAgencyCode** | **string** | If this filter is specified, the results will be filtered to addresses that are either :  - Available from any agency.  - Available only from the specific agency (the restriction can be directly on the address, or in the city) | 
 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListAddressDto**](IPagedResourceListAddressDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreetsSuggestions

> IResourceListStreetSuggestionDto GetStreetsSuggestions(ctx).CountryCode(countryCode).CityName(cityName).PostCode(postCode).StreetName(streetName).Count(count).Execute()

Get streets suggestions



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
	countryCode := "countryCode_example" // string | Mandatory. The exact country code
	cityName := "cityName_example" // string | Mandatory. An exact or partial city name
	postCode := "postCode_example" // string | Mandatory. An exact or partial post code
	streetName := "streetName_example" // string | Mandatory. Search term. An exact or partial street name
	count := int32(56) // int32 | Default value is '100'. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetStreetsSuggestions(context.Background()).CountryCode(countryCode).CityName(cityName).PostCode(postCode).StreetName(streetName).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetStreetsSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreetsSuggestions`: IResourceListStreetSuggestionDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetStreetsSuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreetsSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | Mandatory. The exact country code | 
 **cityName** | **string** | Mandatory. An exact or partial city name | 
 **postCode** | **string** | Mandatory. An exact or partial post code | 
 **streetName** | **string** | Mandatory. Search term. An exact or partial street name | 
 **count** | **int32** | Default value is &#39;100&#39;. | [default to 100]

### Return type

[**IResourceListStreetSuggestionDto**](IResourceListStreetSuggestionDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

