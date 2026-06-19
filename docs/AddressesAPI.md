# \AddressesAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAddressDeletionConstraints**](AddressesAPI.md#CheckAddressDeletionConstraints) | **Get** /v3/addresses/{id}/deletion-constraints | Check address deletion constraints
[**CreateAddress**](AddressesAPI.md#CreateAddress) | **Post** /v3/addresses | Create address
[**DeleteAddress**](AddressesAPI.md#DeleteAddress) | **Delete** /v3/addresses/{id} | Delete address
[**GetAddressAdditionalContacts**](AddressesAPI.md#GetAddressAdditionalContacts) | **Get** /v3/addresses/{id}/additional-contacts | Get address additional contacts
[**GetAddressById**](AddressesAPI.md#GetAddressById) | **Get** /v3/addresses/{id} | Get address by id
[**GetAddressByIdAlternativeRoute**](AddressesAPI.md#GetAddressByIdAlternativeRoute) | **Get** /v3/addresses/by-address-id | Get address by id (alternative route)
[**GetAddressForbiddenVehicleTypes**](AddressesAPI.md#GetAddressForbiddenVehicleTypes) | **Get** /v3/addresses/{id}/forbidden-vehicle-types | Get address forbidden vehicle types
[**GetAddresses**](AddressesAPI.md#GetAddresses) | **Get** /v3/addresses | Get addresses
[**GetStreetsSuggestions**](AddressesAPI.md#GetStreetsSuggestions) | **Get** /v3/addresses/streets/suggestions | Get streets suggestions
[**UpdateAddress**](AddressesAPI.md#UpdateAddress) | **Patch** /v3/addresses/{id} | Update address



## CheckAddressDeletionConstraints

> EntityDeletionConstraintsDto CheckAddressDeletionConstraints(ctx, id).Execute()

Check address deletion constraints



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
	id := int32(56) // int32 | Address identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.CheckAddressDeletionConstraints(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.CheckAddressDeletionConstraints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAddressDeletionConstraints`: EntityDeletionConstraintsDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.CheckAddressDeletionConstraints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAddressDeletionConstraintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityDeletionConstraintsDto**](EntityDeletionConstraintsDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddress

> CreateAddressResultDto CreateAddress(ctx).Command(command).Execute()

Create address



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
	command := *openapiclient.NewCreateAddressCommand("Name_example") // CreateAddressCommand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.CreateAddress(context.Background()).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.CreateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddress`: CreateAddressResultDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.CreateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **command** | [**CreateAddressCommand**](CreateAddressCommand.md) |  | 

### Return type

[**CreateAddressResultDto**](CreateAddressResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, id).Execute()

Delete address



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
	id := int32(56) // int32 | Address identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AddressesAPI.DeleteAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.DeleteAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressAdditionalContacts

> IPagedResourceListAddressContactDto GetAddressAdditionalContacts(ctx, id).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get address additional contacts



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
	id := int32(56) // int32 | 
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressAdditionalContacts(context.Background(), id).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressAdditionalContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressAdditionalContacts`: IPagedResourceListAddressContactDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressAdditionalContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressAdditionalContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListAddressContactDto**](IPagedResourceListAddressContactDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressById

> AddressDto GetAddressById(ctx, id).Fields(fields).Execute()

Get address by id



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
	id := int32(56) // int32 | Identifier of the address
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressById(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressById`: AddressDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Identifier of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetAddressByIdAlternativeRoute

> AddressDto GetAddressByIdAlternativeRoute(ctx).AddressId(addressId).Fields(fields).Execute()

Get address by id (alternative route)



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
	resp, r, err := apiClient.AddressesAPI.GetAddressByIdAlternativeRoute(context.Background()).AddressId(addressId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressByIdAlternativeRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressByIdAlternativeRoute`: AddressDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressByIdAlternativeRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressByIdAlternativeRouteRequest struct via the builder pattern


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


## GetAddressForbiddenVehicleTypes

> IPagedResourceListAddressVehicleTypeDto GetAddressForbiddenVehicleTypes(ctx, id).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()

Get address forbidden vehicle types



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
	id := int32(56) // int32 | 
	startIndex := int32(56) // int32 | Pagination start index (offset). Default is 0. (optional)
	count := int32(56) // int32 | Pagination count (page size). Default is 10, maximum is 200. (optional)
	sort := []string{"Inner_example"} // []string | Pagination sorting fields separated with a comma (optional)
	desc := []string{"Inner_example"} // []string | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. (optional)
	fields := "fields_example" // string | Projection fields separated with a comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddressForbiddenVehicleTypes(context.Background(), id).StartIndex(startIndex).Count(count).Sort(sort).Desc(desc).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddressForbiddenVehicleTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddressForbiddenVehicleTypes`: IPagedResourceListAddressVehicleTypeDto
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddressForbiddenVehicleTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressForbiddenVehicleTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Pagination start index (offset). Default is 0. | 
 **count** | **int32** | Pagination count (page size). Default is 10, maximum is 200. | 
 **sort** | **[]string** | Pagination sorting fields separated with a comma | 
 **desc** | **[]string** | Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields. | 
 **fields** | **string** | Projection fields separated with a comma | 

### Return type

[**IPagedResourceListAddressVehicleTypeDto**](IPagedResourceListAddressVehicleTypeDto.md)

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


## UpdateAddress

> UpdateAddress(ctx, id).Command(command).Execute()

Update address



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
	id := int32(56) // int32 | Address identifier
	command := *openapiclient.NewUpdateAddressDto() // UpdateAddressDto | Update address command.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AddressesAPI.UpdateAddress(context.Background(), id).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.UpdateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Address identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **command** | [**UpdateAddressDto**](UpdateAddressDto.md) | Update address command. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/merge-patch+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

