# \MeAPI

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyAccountInformation**](MeAPI.md#GetMyAccountInformation) | **Get** /v3/me | Get my account information
[**GetOauthAuthorizationUrl**](MeAPI.md#GetOauthAuthorizationUrl) | **Get** /v3/sso/oauth-authorization-url | Get oauth authorization url
[**GetSAMLAuthorizationUrl**](MeAPI.md#GetSAMLAuthorizationUrl) | **Get** /v3/sso/saml-authorization-url | Get SAML authorization url



## GetMyAccountInformation

> MyAccountInfoDto GetMyAccountInformation(ctx).Execute()

Get my account information

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyAccountInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyAccountInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAccountInformation`: MyAccountInfoDto
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyAccountInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAccountInformationRequest struct via the builder pattern


### Return type

[**MyAccountInfoDto**](MyAccountInfoDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthAuthorizationUrl

> GetSsoAuthorizationUrlResultDto GetOauthAuthorizationUrl(ctx).RedirectUri(redirectUri).SsoCode(ssoCode).State(state).Scope(scope).Execute()

Get oauth authorization url



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
	redirectUri := "redirectUri_example" // string | 
	ssoCode := "ssoCode_example" // string | 
	state := "state_example" // string | 
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetOauthAuthorizationUrl(context.Background()).RedirectUri(redirectUri).SsoCode(ssoCode).State(state).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetOauthAuthorizationUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthAuthorizationUrl`: GetSsoAuthorizationUrlResultDto
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetOauthAuthorizationUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthAuthorizationUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectUri** | **string** |  | 
 **ssoCode** | **string** |  | 
 **state** | **string** |  | 
 **scope** | **string** |  | 

### Return type

[**GetSsoAuthorizationUrlResultDto**](GetSsoAuthorizationUrlResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSAMLAuthorizationUrl

> GetSsoAuthorizationUrlResultDto GetSAMLAuthorizationUrl(ctx).SsoCode(ssoCode).RelayState(relayState).AssertionConsumerServiceUrl(assertionConsumerServiceUrl).Execute()

Get SAML authorization url



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
	ssoCode := "ssoCode_example" // string | 
	relayState := "relayState_example" // string |  (optional)
	assertionConsumerServiceUrl := "assertionConsumerServiceUrl_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetSAMLAuthorizationUrl(context.Background()).SsoCode(ssoCode).RelayState(relayState).AssertionConsumerServiceUrl(assertionConsumerServiceUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetSAMLAuthorizationUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSAMLAuthorizationUrl`: GetSsoAuthorizationUrlResultDto
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetSAMLAuthorizationUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSAMLAuthorizationUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoCode** | **string** |  | 
 **relayState** | **string** |  | 
 **assertionConsumerServiceUrl** | **string** |  | 

### Return type

[**GetSsoAuthorizationUrlResultDto**](GetSsoAuthorizationUrlResultDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

