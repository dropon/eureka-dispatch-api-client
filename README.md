# Go API client for openapi

This API is dedicated to operators.  
Get API url and your client credentials from [integrators website](https://integrators.eureka-technology.fr/).  

You can download the Open Api specifications (swagger json document) in the link mentioned above.

Please use our [customer support portal](https://eureka-technology.atlassian.net/servicedesk/customer/portals) for questions or issues.
If you do not have a support account yet, please reach out to your contact at Eureka-Technology to create one.
If you don't have a dedicated contact, you can send an email to support@eureka-technology.com.

# Authentication  

## Access tokens

This Api uses OAuth2 authentication.  

All the requests to the api, except authentication ones, must include a valid bearer token.  
To obtain an access token, call the _/token_ endpoint, provide an Authorization header with a basic authentication type, constructed using your client_id and client_secret.  
The body must be a form url encoded with the parameters corresponding to the authentication flow (see grant types section for more details).  
The api returns a response containing the access token, its expiration in seconds, a refresh token, and other properties. Example :  
```json
{
    \"access_token\": \"ey...\",
    \"token_type\": \"bearer\",
    \"expires_in\": 3599,
    \"refresh_token\": \"...\",
    \".issued\": \"(date)\",
    \".expires\": \"(date)\"
}
```
When the client id or client secret is invalid, a 400 Bad request is returned :
```json
{
    \"error\": \"invalid_client\",
    \"error_description\": \"The client id or secret is incorrect.\"
}
```
The returned access token can then be used in the requests to the api, it must be included in the Authorization header, example : `Authorization: Bearer ey...`.  
The access token is valid 1 hour by default, this value can change depending on the carrier.  

### Example using postman  

Here is an example of token request using postman (to be adapted according to the grant type).  

The request header contains the basic authentication.  
![Authentication authorization header](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-authorization-header.png)

The request body is a form url encoded containing the grant type and the additional parameters (depending on the grant type).  
![Authentication body with password](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-body-with-password.png)

The response contains the access token and additional information.  
![Authentication response](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-response.png)

## Refresh token  

In _/token_ response, you receive an access token and a refresh token.  
The refresh token allows to get a new access token without providing the original credentials.  
When an access token has expired, you can exchange the refresh token for a new access token.  
The refresh token is valid 24 hours by default, this value can change depending on the carrier.  
To generate a new access token using this refresh token, use refresh_token grant type. Example :  

```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=refresh_token&refresh_token=your_refresh_token
```  

You can invalidate the refresh token by calling _/revoke_ endpoint. Example :  
```
POST /revoke  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
token=your_refresh_token
```  

### Refresh token limitation

There is a limit in the number of active refresh tokens per user. It is 10000 by default, this value can change depending on the carrier.  
If this limit is reached, you will receive the error _max_number_of_sessions_reached_, you will need to wait for refresh token expiration before being able to get a new access token for the user.  
To avoid the limit, reuse the access token, do not ask for a new one for each request to the Api.  

## Grant types

### Password flow  

Use password grant to connect as a user, using its credentials : provide the parameters grant_type _password_, username and password. Example :  
```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=password&username=user_username&password=user_password
```  

When user's credentials are invalid, a 400 Bad request is returned :
```json
{
    \"error\": \"invalid_grant\",
    \"error_description\": \"Unknown user or incorrect password.\"
}
```  

### Client credentials  

Use client credentials to connect as an application, outside of the context of a user : provide the parameters grant_type _client_credentials_. Example :  
```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=client_credentials
```  

### SSO  

To connect with an external SSO provider, SSO settings must be configured in Dispatch, and you will have a SSO configuration code for each provider.  
SSO users must already be configured in Dispatch.  

#### SSO Authorization code  

This is the standard method to use an external OAUTH SSO provider configured in Dispatch. Use this method if your application does not have its own client id and client secret in the external SSO provider.  

- First, your application will need to open a browser to send the user to the external SSO oauth provider.  

To get the authorization url, use the operation _sso/oauth-authorization-url_ and provide the parameters : ssoCode, redirectUri and state, and provide bearer token created with a client credentials authentication.  

- When the user connects, the oauth provider will redirect back to your application url and provide an authorization code and the state you provided  

- Call token endpoint in the API with grant_type authorization_code_TheSsoConfigurationCode and parameters authorization_code and redirect_uri. Example :  

```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=authorization_code_ABC&authorization_code=UrlEncode(authorizationCode)&redirect_uri=UrlEncode(https://yourwebsite/sso/oauth/callback)
```  

- The API will validate the authorization code, and you will receive a token for this API for the user corresponding to the SSO user.  

#### SSO Access token  

This is the another method to use an external OAUTH SSO provider configured in Dispatch. You can use this method if your application has its own external SSO provider information (client id and client secret).  
- From your application, authenticate to the external SSO provider using authorization code flow : 
    - Build the authorization url with your SSO credentials 
    - After user connection, you receive an authorization code  
    - Call the SSO provider to get an access token  
- Call token endpoint in the API with grant_type access_token_TheSsoConfigurationCode and parameters access_token. Example :  

```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=access_token_ABC&access_token=UrlEncode(ssoAccessToken)
```  

- The API will validate the sso access token, and you will receive a token for this API for the user corresponding to the SSO user.  

#### SAML  

This is the standard method to use an external SSO provider configured in Dispatch and using SAML V2 protocol.

- First, your application will need to open a browser to send the user to the external SSO SAML provider.  

To get the authorization url, use the operation _sso/saml-authorization-url_ and provide the parameters : ssoCode and relayState, and provide bearer token created with a client credentials authentication.  

- When the user connects, the SAML provider will redirect back to your application with additional information in query parameters on in a form  
    - if information is received with a POST, read the request form and build a query string using form data  
    - if information is received with a GET (redirect method), get the query string  

- Call token endpoint in the API with grant_type saml_TheSsoConfigurationCode and parameters saml_http_method and saml_content. Example :  

```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=saml_ABC&saml_http_method=GET&saml_content=UrlEncode(SAMLResponse=value1&RelayState=value2)
```  

- The API will validate the SAML content, and you will receive a token for this API for the user corresponding to the SSO user.  

# General instructions  

## Projection  

To select the fields returned by the api, provide 'fields' parameter in query string: list of fields separated with a comma.   
When fields are not provided, calls will only return the first level fields, unless specified otherwise in a route description.  

Example with simple fields : `GET route_to_my_resource/myResourceUid?fields=field1,field2,field3` returns :  
```json
{
    \"field1\": \"My field 1\",
    \"field2\": \"My field 2\",
    \"field3\": \"My field 3\"
}
```  

Example with nested properties and collections : `GET route_to_my_resource/myResourceUid?fields=field1,nested1/nested2/nestedProperty3,myCollection1/field2` returns :  
```json
{
    \"field1\": \"My field 1\",
    \"nested1\": {
        \"nested2\": {
            \"nestedProperty3\": \"my nested property value 3\"
        }
    }
    \"myCollection1\": [
        {
            \"field2\": \"my field 2\"
        }
    ]
}
```  

You can group properties on a nested object using parenthesis. Example to get the properties nested1.field1 and nested1.field2 : `fields=nested1(field1,field2)`.  

Example to get the properties at the first level only, without nested properties or collections : `route_to_my_resource?fields=~`.  

Example to get all the properties, included nested properties and collections : `route_to_my_resource?fields=*`.  

## Filtering  

Access to paginated resources accept filtering with query parameters. Example : `route_to_my_paginated_resources?myFilter1=myValue1&myFilter2=myValue2`.  

When a query parameter is a collection, the values must be delimited with a comma. To escape a comma in a value, double it.  
Example to pass the words : \"hello,world\", \"value1\", \"value2\" to \"names\" parameter : `route_to_my_paginated_resources?names=hello,,world,value1,value2`.  

When the description of a filter indicates that it accepts a pattern, you can search with _startsWith_, pattern.  
In some cases, when it is specified in the description, the filter also supports _endsWith_ or _contains_ patterns.  

Examples : 
- Exact match : `myFilter=abc`  
- Starts with abc : `myFilter=abc*`  
- Ends with abc : `myFilter=*abc`  
- Contains abc : `myFilter=*abc*`   

For some operations, you can search on multiple fields, using pattern and pattern fields. The fields can be different depending on the operation, check the description of _patternFields_ parameter to get the available fields.  
Example, to search resources having a code or a label starting with abc : `pattern=abc*&patternFields=code,label`.  

## Pagination  

For paginated resources, provide query option's _startIndex_, _count_, _fields_, _sort_ for ascending ordering and _desc_ for descending ordering.  
Example : `route_to_my_paginated_resources?startIndex=0&count=10&fields=field1,field2&sort=field2`  
Example of response :  
```json
{
    \"data\": [
        {
            \"field1\": \"My field 1 value 1\",
            \"field2\": \"My field 2 value 1\"
        },
        {
            \"field1\": \"My field 1 value 2\",
            \"field2\": \"My field 2 value 2\"
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```
## Errors  

In case of errors, an object Problem Details is returned with its http status code.  
It corresponds to the standard problem details, with an additional property _data_.  
See more details about RFC 7807 problem details [here](https://tools.ietf.org/html/rfc7807).  
Example :  
```json
{
    \"type\": \"/problems/types/service-unavailable\",
    \"title\": \"Service unavailable.\",
    \"status\": 503,
    \"instance\": \"/problems/instances/dispatch-module-not-available\"
}
```

When a business error occurs, an error code and error message are included in an additional property.  
Example :  
```json
{
    \"data\": {
        \"errors\": [
            {
                \"code\": \"resource_not_found\",
                \"message\": \"Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",
            }
        ]
    },
    \"type\": \"/problems/types/not-found\",
    \"title\": \"A business error occurred\",
    \"status\": 404,
    \"detail\": \"resource_not_found : Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",
    \"instance\": \"/problems/instances/business-error\"
}
```

### Error codes  
#### General  
Error code | Error description  
----|----  
missing_mandatory_parameter | A mandatory parameter is missing
must_be_strictly_greater_than_zero | The parameter must be greater than zero
resource_not_found | The resource does not exist
resource_already_exists | A resource with the same identifier already exists
user_missing_permission | The user is not allowed to perform the operation or to access the resource
datetime_out_of_range | The date time is out of the range supported by the system
guid_should_not_be_empty | The unique identifier GUID must not be empty
invalid_enum_value | The provided value is not a valid enumeration value 
invalid_field_name | The field provided in the query parameters is not valid
field_not_supported | The field in the 'fields' or 'sort' parameters is not supported
user_entity_access_forbidden | The user is not allowed to access the resource
value_out_of_range | The provided value is out of the expected range
max_length_exceeded | The maximum length of the parameter is exceeded
end_date_must_be_greater_than_start_date | The end date must be greater than the start date
max_date_interval_reached | The maximum date interval is reached
user_has_no_eureka_maps_key | The user does not have access to Eureka Maps
max_user_application_preferences_reached | The maximum number of application preferences is reached for this user

#### Drivers  
Error code | Error description  
----|----  
driver_unavailability_reason_in_use | The driver unavailability reason is used
driver_unavailability_reason_disabled | The driver unavailability reason is disabled
driver_unavailability_overlapping | There is a driver unavailability in the same interval
contractor_agent_subcontractor_mandatory | The subcontractor is mandatory when the contractor agent is provided
contractor_agent_invalid_subcontractor | The provided contractor agent cannot be used with the provided subcontractor

#### Missions  
Error code | Error description  
----|----  
transport_assignment_constraints_violation | Transport assignment constraints not fulfilled
transport_unassignment_constraints_violation | Transport unassignment constraints not fulfilled
transport_fragmentation_constraints_violation | Transport fragmentation constraints not fulfilled
transport_defragmentation_constraints_violation | Transport defragmentation constraints not fulfilled
invalid_file_category | The provided file category is invalid
invalid_custom_parameter | The provided custom parameter is invalid
invalid_packing_nature | The package nature is invalid
invalid_unit_code | The unit code is invalid
invalid_dimension_unit | The dimension unit is invalid
transport_unexpected_state | The operation cannot be performed because the transport has an unexpected status
document_report_unexpected_state | The operation cannot be performed because the transport document report has an unexpected status
invalid_service_code | The provided service code is invalid
service_disabled | The provided service is disabled
invalid_customer_code | The provided customer code is invalid
customer_not_usable | The provided customer cannot be used
invalid_pricing_path | The provided pricing path is invalid
itinerary_computation_failed | Unable to compute itinerary for the transport

#### Addresses  
Error code | Error description  
----|----  
invalid_address_id | The provided address id is invalid
invalid_city_id | The provided city id is invalid
geocoding_failed | Unable to geocode the address



## Resource update  

You can use PATCH operations to partially update a resource.  
The update uses merge patch semantics : you need to provide only the fields to updates.  
See more details about JSON merge patch standard [here](https://tools.ietf.org/html/rfc7396).  

Fields that are not present in the request will be preserved, and fields set to null will be cleared.  

There is a special process for lists with nested objects :  
- The objects that are not provided in a list are not updated.  
- For all objects provided in a list : specify the merge action (add, update, remove) and the entity key when the action is update or remove.  

Example, with comments indicating the performed operations :

```json
{
    \"property1\": \"new value 1\",         // property1 is updated 
    \"nested1\": {                        // nested1 is added if it doesn't exist, otherwise, it is updated
        \"property2\": \"my value 2\",      // property2 is updated
        \"nested2\": null,                // nested2 is deleted
        \"nested3\": {                    // nested3 is added if it doesn't exist
            \"property3\": \"my value 3\"   // property3 is updated
        }
    },
    \"nestedCollection1\": [              // nestedCollection1 is updated
        {
            \"myIdentifier\": \"id1\",
            \"mergeAction\": \"update\",
            \"property4\": \"my value 1\"   // the property property4 of the item with myIdentifier equals to id1 is updated
        },
        {
            \"myIdentifier\": \"id2\",
            \"mergeAction\": \"remove\"     //  the item with myIdentifier equal to id2 is removed 
        },
        {
            \"mergeAction\": \"add\",       // an item is added to the collection with the provided property4.
            \"property4\": \"my value 2\"
        }
    ]
}
```  

## Date formats

Many dates in the system are saved without timezone information. These dates depend on the context (for example : pickup date correponds generally to the timezone of pickup address, but the offset is not kept in the system). In these cases, the timezone is not taken into account when provided in the input. For example, if you provide `2022-01-31T10:00:00+05:00`, the system will keep only `2022-01-31T10:00:00`, the date will be returned with an unspecified timezone.  

The api returns the timezone when it is available in the system.  

Examples of dates returned by the api :
* Date with unspecified time zone : `2022-01-31T10:00:00`
* UTC date : `2022-01-31T10:00:00+00:00` or `2022-01-31T10:00:00Z`
* Date with offset +5h : `2022-01-31T10:00:00+05:00`

## Versioning

This API is versioned to handle changes in API contract.  
The versioning is done with URI path. Example : _https://myapi/v1/myresource_.  
The versioning applies on all the operations except _/token_ and _/revoke_ endpoints. Example: _https://myapi/token_.  

A new version is released when a breaking change is made.  
The previous versions remain maintained for a certain time. After a time, old versions may be deprecated (marked as obsolete) then removed.  
Remark: breaking changes can be made on pre-release versions. These versions can also be removed from the API and replaced with a release version.  

Examples of breaking changes :  
- Change a property name or type in a request or response  
- Add a required parameter in a request  
- Remove a property in a response  
- Remove an operation in the API  

Examples of non-breaking changes :
- Add a value in a response
- Add an optional parameter in a request  
- Add an operation in the API  
- Change the format of any opaque identifier
- Change the format of a date time : add or remove timezone information, for example : change from `2022-01-31T10:00:00` to `2022-01-31T10:00:00Z` or to `2022-01-31T10:00:00+05:00`.  
- Replacing a success status code with another success status code (i.e: replacing a 200 with a 201)
- Replacing a failure status code with another failure status code (i.e: replacing a 500 with a 400)

When a feature is added (without breaking changes), it is available for the last released API version.  
In this case, a flag can be added to indicate that this feature is available. See api flags endpoint.  

## Rate Limiting

This API uses consumption quotas to prevent excessive calls and to maintain service availability.  
Two types of rate limiting are applied :
- leaky buckets  
- concurrency  

You can view these limits in Eureka integrator's portal.  
There are also additional limits for calls performed by one bearer token, configured by the carrier. These limits are not visible in the portal.  

When a quota is exceeded, the API returns the error code 429 (Too Many Requests). A _Retry-After_ header is included and indicates (in seconds) how long to wait before making a new request.  

### Leaky bucket rate limiting  

A leaky bucket rate limiting is applied to each API client, it depends on two parameters : the bucket size and the leak rate.  

Each API client have its own bucket of a limited size.  
When a request is received :  
- If the bucket is not full, the bucket counter is increased.  
- If the bucket is full (the bucket counter reaches the bucket size), the request is rejected. In this case, the API client must wait before sending a new request.  

The bucket counter is decreased periodically to accept new requests.  

You can find more details about leaky bucket algorithm [here](https://en.wikipedia.org/wiki/Leaky_bucket).  

### Concurrency rate limiting  

The API accepts only a limited number of concurrent calls for each API client.  
If the maximum number of concurrent calls is reached, the request is rejected. In this case, the API client must wait before sending a new request.  
  
# Dispatch Api

## Available features

This API includes various features :
* Retrieval of multiple resources needed for mission entry
* Mission and quotation operations : creation, update, cancellation
* Operations on transport documents (download, upload, delete)
* Operations on transport : sub state application, assignment, pickup, delivery, termination, cancellation
* Retrieval of transport data
* Transport document reports generation
* Fragmentation

This API does not include these features :
* In mission entry : dangerous goods, air transport, deposits (aka consignes), orders in sequence
* Disputes and misfunction sheets, call for tender
* D2D
* Operations (creation, update, delete) on resources and tools except on transports
* Retrieval of resources not needed for mission entry
* Billing
* Regular orders
* Rounds
* Mission scheduling
* Exports and imports
* Reports not related to one transport

## Missions, quotations and transports  

A transport is a system for conveying something (goods or people) from a place (pickup) to another (delivery). A transport is related to one customer and to one service and may contain packages.  

A mission is an order containing one or many transports.  
In general, a mission will contain only one transport. But in case of a mission with return, it will contain a second transport.  
This API supports missions with one or two transports.  

A quotation can be requested to indicate how much an order will cost.  
When a quotation is validated, it is transformed into a new mission : a new mission is created with quotation information, and the quotation is archived.  

Use :
* GET /transports to retrieve a collection of transports (use filtering and projection)  
* POST /missions to create a mission
* POST /quotations to create a quotation



## Mission entry

The following steps can be used to create a mission or a quotation.

### Authenticate as a user

Connect as a user, with a login and password  
```
POST /token  
Content-Type: application/x-www-form-urlencoded  
Authorization: Basic Base64(client_id:client_secret)  
grant_type=password&username=my_dispatch_user_login&password=my_dispatch_user_password
```  
The response is a json containing the access token

```json
{
    \"access_token\": \"eyJh...\",
    \"token_type\": \"bearer\",
    \"expires_in\": 3599,
    \"refresh_token\": \"11c...\",
    \"as:clientId\": \"...\",
    \".issued\": \"Mon, 08 Jul 2024 14:43:45 GMT\",
    \".expires\": \"Mon, 08 Jul 2024 15:43:45 GMT\"
}
```

For all the following calls, add a header `Authorization: Bearer my_access_token`.  

### Select a customer  

Get the customers available for the connected user, use filtering to get only the fields you need, use pagination.  

```
GET /customers?startIndex=0&count=10&fields=uid,code,label
```

Response

```json
{
    \"data\": [
        {
            \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",
            \"code\": \"CUST01\",
            \"label\": \"Sample customer 1\"
        },
        {
            \"uid\": \"9acc8b97729c447e9b52bcf87d392c35\",
            \"code\": \"CUST02\",
            \"label\": \"Sample customer 2\"
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```
Get customer's information by its unique identifier or by its code. Examples :
```
GET /customers/1b1548e7d1f046b0ba357539b139be7f?fields==uid,code,label
GET /customers/by-customer-code?customerCode=CUST01&fields==uid,code,label
```

Response

```json
{
    \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",
    \"code\": \"CUST01\",
    \"label\": \"Sample customer 1\"
}
```

### Addresses

In address step, validate the selected address by getting the cities and countries available in the system, a valid city can be needed for the pricing. Example :  

```
GET /cities?postCode=34000&cityName=montpellier&countryCode=FR&fields=cityId,name,postCode,country/label
```

Response

```json
{
    \"data\": [
        {
            \"cityId\": 17037,
            \"name\": \"MONTPELLIER\",
            \"postCode\": \"34000\",
            \"country\": {
                \"label\": \"FRANCE\"
            }
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 1,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 1
    }
}
```

### Packing

In package entry, get package natures.  

```
GET /package-natures?fields=packageNatureId,code,label,packageFamilyLabel
```

Response

```json
{
    \"data\": [
        {
            \"packageNatureId\": 1,
            \"code\": \"CARTONA3\",
            \"label\": \"Carton A3\",
            \"packageFamilyLabel\": \"Colis\"
        },
        {
            \"packageNatureId\": 2,
            \"code\": \"Boite\",
            \"label\": \"Boite\",
            \"packageFamilyLabel\": \"Consommables\"
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```


### Services

In service entry, get services.

```
GET /services?startIndex=0&count=10&fields=uid,code,label,serviceFamily/label
```

Response

```json
{
    \"data\": [
        {
            \"uid\": \"4da0ba3ef10c4147b73debd01d375000\",
            \"code\": \"T1\",
            \"label\": \"BREAK 500 KG\",
            \"serviceFamily\": {
                \"label\": \"Poids Lourds\"
            }
        },
        {
            \"uid\": \"1cfb5b59d877406aa7b86908c9c12f63\",
            \"code\": \"T1-MO\",
            \"label\": \"MOTO 50KG\",
            \"serviceFamily\": {
                \"label\": \"Prestations intra muros\"
            }
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```

TO see the sub services included in a service, get its sub services.

```
GET /services/4da0ba3ef10c4147b73debd01d375000/sub-services?fields=subServiceCode,subServiceLabel,defaultQuantity
```

Response

```json
{
    \"data\": [
        {
            \"subServiceCode\": \"ADV\",
            \"subServiceLabel\": \"Assurance ad valorem\",
            \"defaultQuantity\": 1.00
        },
        {
            \"subServiceCode\": \"KM\",
            \"subServiceLabel\": \"Kilomètres\",
            \"defaultQuantity\": 0.00
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```

For additional sub services, use sub-services route with available filtering and paging. Example: 

```
GET /sub-services?isEnabled=true&code=KM*&patternFields=code,label
```

Response

```json
{
    \"data\": [
        {
            \"code\": \"KM\",
            \"label\": \"Kilomètres\",
            \"isEnabled\": true,
            \"unitCode\": \"KM\",
            \"comment\": null,
            \"vatCode\": \"1\",
            \"quantityMustBeSelected\": false,
            \"subServiceQuantities\": null
        },
        {
            \"code\": \"KM-NS\",
            \"label\": \"Kilomètres sans surcharge carburant\",
            \"isEnabled\": true,
            \"unitCode\": \"KM\",
            \"comment\": null,
            \"vatCode\": \"2\",
            \"quantityMustBeSelected\": false,
            \"subServiceQuantities\": null
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": null,
        \"pageCount\": 1,
        \"totalItemCount\": 2,
        \"pageNumber\": 1,
        \"pageSize\": 10,
        \"hasPreviousPage\": false,
        \"hasNextPage\": false,
        \"isFirstPage\": true,
        \"isLastPage\": true,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 2
    }
}
```


### Mission creation

Before creating a mission, you can perform a dry run by using `POST /missions/dry-run` route. It simulates a mission creation without creating it. Example:  

```json
{
    \"transports\": [
        {
            \"customerCode\": \"ClientDemo1\",
            \"agencyCode\": \"01\",
            \"pickupStep\": {
                \"stepAddress\": {
                    \"addressId\": 314312
                },
                \"stepDate\": {
                    \"plannedStart\": {
                        \"dateTime\": \"2024-07-02T00:00:00\",
                        \"isTimeOfDayIgnored\": true
                    }
                }
            },
            \"deliveryStep\": {
                \"stepAddress\": {
                    \"AddressId\": 339923
                },
                \"stepDate\": {
                    \"plannedStart\": {
                        \"dateTime\": \"2024-07-03T00:00:00\",
                        \"isTimeOfDayIgnored\": true
                    }
                }
            },
            \"serviceCode\": \"T8\",
            \"subServices\": [
                {
                    \"code\": \"KM\"
                },
                {
                    \"code\": \"ADV\",
                    \"forcedSellQuantity\": 11
                }
            ],
            \"serviceCustomParameters\": [
                {
                    \"code\": \"PAL\",
                    \"Value\": \"2\"
                }
            ]
        }
    ]
}
```

This route returns information about the mission with its pricing, but without creating it.  
To create a mission, use the same request, with `POST /missions` route.  


### Additional remarks on transport creation

* The provided entities (example: customer, service, agency) must be available for the connected user regarding its restrictions.  
* The provided customer usable and not a prospect. If the orderer is mandatory for a customer, the orderer must be provided.  
* The provided service must be enabled. If it has an agency restriction, it must match customer's agency.  
* All provided identifiers (example : city id, agency code, VAT code, unit code) must exist
* Customer and service's custom parameters are checked. Example : the command is rejected if a custom parameter is mandatory, not provided and does not have a default value.  
* Transport's references are checked. Examples : if mandatory and not provided, if must be selected from a list but does not match this list.  
* Pickup and delivery dates are mandatory
* For package lines, provide either the nature id or a nature (free text). Default values are copied from the nature.  
To change this behavior, enable _skipPackageNatureDefaultValuesApplication_ in command's api behavior.    
* For transport's sub services: when they are not provided, the transport uses service's default sub services. Otherwise, the transport uses only the specified sub services : service's default sub services are not included and must be specified manually if they are needed.  
* Sell pricing path can be forced. 
The pricing path must be a customer pricing path (not a subcontractor's one).  
When provided, the addresses are overwritten with the ones defined in the provided sell pricing path.    
The provided service must match the one defined in the pricing path.  
* The mission or quotation is created as unique if at least two transports are provided. To force this parameter, provide _forceUniqueOrder_ in command's api behavior.  
* Itinerary computation is done if configured in global settings and if the distance is not provided. To ignore failures during this step, enable _ignoreItineraryComputationFailures_ in command's api behavior.  
* Transport steps geocoding is attempted if itinerary computation is required. To try to geocode a step in other cases, enable _tryGeocodeTransportSteps_ in command's api behavior.   
* Gas emission is computed if enabled in global settings and not forced. 
* Transport bill address is unsed only when the customer is occasional.


## Transport retrieval

Here is an example to search transports using projection, filtering and pagination.  
In this example, we want the 5 top transports having a pickup post code 34000, for customer code CUST01 and with statuses C (assigned) and 0 (recorded), we only want the fields uid, missionNumber, serviceCode and status.  
```
GET /transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=0&count=5
```
Response
```json
{
    \"data\": [
        {
            \"uid\": \"fcba3be8e09740499079f86dff98f5e6\",
            \"missionNumber\": 1,
            \"status\": \"C\",
            \"serviceCode\": \"MO\"
        },
        {
            \"uid\": \"10b9c9ba9b5240cba50fb834bf5cb714\",
            \"missionNumber\": 7203,
            \"status\": null,
            \"serviceCode\": \"TNT\"
        },
        {
            \"uid\": \"bfe364f5478a4f2890ea1cba6c535a99\",
            \"missionNumber\": 13443,
            \"status\": null,
            \"serviceCode\": \"T1\"
        },
        {
            \"uid\": \"87253a9702134a198be1748d02eec64d\",
            \"missionNumber\": 13444,
            \"status\": null,
            \"serviceCode\": \"T1\"
        },
        {
            \"uid\": \"5febb2485f474b978730c805ad6e7b79\",
            \"missionNumber\": 13445,
            \"status\": null,
            \"serviceCode\": \"T1\"
        }
    ],
    \"paging\": {
        \"prevLink\": null,
        \"nextLink\": \"https://mylicense.dispatchapi.dispatch-rts.com:443/v3/transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=5&count=5\",
        \"pageCount\": 2,
        \"totalItemCount\": 6,
        \"pageNumber\": 1,
        \"pageSize\": 5,
        \"hasPreviousPage\": false,
        \"hasNextPage\": true,
        \"isFirstPage\": true,
        \"isLastPage\": false,
        \"firstItemOnPage\": 1,
        \"lastItemOnPage\": 5
    }
}
```


## Transport update

Here is an example to partially update a transport using PATCH operation.  

```
PATCH /transports/dfa98649e1fc4cf8a77c43351b531007
{
    \"apiBehavior\": {
        \"SkipPackageNatureDefaultValuesApplication\": false,
        \"updateTransportOperationComment\": \"My update operation 01\"
    },
    \"reference2\": \"new-reference-2\",                // The reference2 will be updated
    \"pickupStep\": {
        \"stepDate\": {
            \"plannedStart\": {
                \"dateTime\": \"2022-02-24 10:00:00\",  // the pickup planned start time will be updated
                \"isTimeOfDayIgnored\": false
            }
        }
    },
    \"packing\": {
        \"lines\": [
            {
                \"mergeAction\": \"update\",            // the quantity and references for package id 2268635 will be updated
                \"packageLineId\": 2268635,
                \"quantity\": 3.00,
                \"references\": [
                    {
                        \"mergeAction\": \"update\",    // the value of the reference with index 3 will be updated
                        \"referenceIndex\": 3,
                        \"value\": \"abc\"
                    }
                ]
            },
            {
                \"mergeAction\": \"remove\",            // the package with id 2268636 will be removed
                \"packageLineId\": 2268636
            },
            {
                \"mergeAction\": \"add\",               // this package will be added
                \"nature\": \"Plis\",
                \"quantity\": 2.00,
                \"dimensionUnit\": \"cm\",
                \"height\": 1.00,
                \"width\": 10.00,
                \"length\": 15.00
            }
        ]
    },
    \"subServices\": [
        {
            \"mergeAction\": \"add\",                   // this service will be added
            \"code\": \"SP01\",
            \"forcedSellQuantity\": 1.2
        }
    ]
}
```  

## Transport statuses

To follow up the progress of a transport, check transport status.  
When a transport is created, it is recorded, then assigned to a driver, picked up from pickup address, delivered to delivery address, terminated and invoiced.  
A transport has one of the following statuses:  

Status | Description
------ | -------
0 (zero), null or empty string | Recorded
A | Canceled
X | Unassigned
b | Pre assigned
B | Pre assigned and accepted by the driver
C | Assigned
E | Picked up
L | Delivered
T | Terminated
s | Dispute undealt
S | Dispute dealt
f | Pre billed
F | Invoiced  

Remarks :
- Transport statuses are case sensitive. Make sure you use the correct case when filtering the transports based on their statuses.  
- To filter recorded transports only, use the value 0 (zero). The api does not provide filtering with empty values. 



## Transport history statuses

Transport history lines returned by this api have one of the following status codes :  

Status | Description
------ | -------
0 (zero) | Recorded
M | Modified
C | Assigned
E | Picked up
L | Delivered



## Dispatch desktop app user legacy permissions

A user operator has a set of permissions used in Dispatch desktop application.  
The usage of these permissions may change in the future.  

Permission Id | Description
------ | -------
5 | Orders corrections
6 | Write scheduled transports
7 | View scheduled transports
8 | Print scheduled transports
10 | Orders consistency
11 | Statistics
12 | Margins
13 | Global settings
14 | Imports / Exports
15 | See Prices
16 | Invoices calculation
17 | Invoices editing
18 | Ad Hoc invoices
19 | Vouchers book
20 | Credit notes
21 | Invoices clearing
22 | View not issued bills
23 | Accounting export
24 | Payment
25 | Write units management
26 | View units
27 | Print units
28 | Invoices management
29 | Generate bills states
30 | Write customers
31 | View customers
32 | Print customers
33 | Write employees
34 | View employees
35 | Print employees
36 | Write subcontractors
37 | View subcontractors
38 | Print subcontractors
39 | Write salespersons
40 | View salespersons
41 | Print salespersons
42 | Write users
43 | View users
44 | Print users
45 | Write addresses
46 | View addresses
47 | Print addresses
48 | Write sectors
49 | View sectors
50 | Print sectors
51 | Write pricing paths
52 | View pricing paths
53 | Print pricing paths
54 | Write services
55 | View services
56 | Print services
57 | Write sub services
58 | View sub services
59 | Print sub services
60 | Write currencies
61 | View currencies
62 | Print currencies
63 | Write discount scale
64 | View discount scale
65 | Print discount scale
66 | Write payment methods
67 | View payment methods
68 | Print payment methods
69 | Write VAT
70 | View VAT
71 | Print VAT
72 | Write vehicles
73 | View vehicles
74 | Print vehicles
75 | Write cities
76 | View cities
77 | Print cities
78 | (Reserved)
79 | (Reserved)
80 | Companies
81 | (Reserved)
82 | (Reserved)
83 | Write orderer profiles
84 | View orderer profiles
92 | Customizable planning
93 | Tracking alerts
94 | Driver's message alerts
95 | Alerts on mobile dispute
96 | Alerts on unread / rejected missions
97 | Alerts on new web mission
98 | Alert on subcontractor's reminder
99 | Distribution entry
100 | Write references
101 | View references
102 | Print references
103 | Delegated orders tracking
104 | Alerts on mobile anomalies
105 | D2D Services equivalences
106 | D2D Services equivalences
107 | D2D amount change
108 | D2D City equivalence
109 | Customer / Subcontractor configuration
110 | D2D City equivalence
111 | Modify terminated orders
112 | Modify billed orders
113 | D2D Errors
114 | New D2D Mission / Modify D2D Mission
115 | Orders ending
116 | Orders cancellation
118 | Write user profiles
119 | View user profiles
120 | Geo device associations
121 | Cartography acces
122 | Create, Update, Delete
123 | Default model choice
124 | Preview
125 | Write agencies
126 | View agencies
127 | Print agencies
128 | Write orderers
129 | View orderers
130 | Print orderers
131 | Create, update, delete POI
132 | Create, update, delete pictogrammes
133 | Invoices clearing
134 | View tems
135 | Tools
136 | Time management administration
137 | View time management
138 | Subcontractors documents
139 | Write transport statuses
140 | Write transport substates
141 | View transport substates
142 | Write anomalies
143 | View anomalies
144 | Write disputes
145 | View disputes
146 | Write deposits
147 | View deposits
148 | Notepad
149 | Write unavailabilities
150 | View unavailabilities
151 | (Reserved)
152 | Consult a dysfunction sheet
153 | Modify a dysfunction sheet
154 | Print a dysfunction sheet
155 | Close misfunction sheet
156 | ReOpen a dysfunction sheet
157 | Write contractor agents
158 | View contractor agents
159 | Print contractor agents
160 | Kick users via instant messaging
161 | GPS devices configuration window
162 | Instant messaging
163 | Dispatch's alerts
164 | Notes
165 | Baracoda Tracking
166 | Set VAT in order
167 | Transport order manual price setting
168 | Write EDI parameters
169 | Global communications history
170 | Write regions
171 | View regions
172 | Write price modification rules
173 | View price modification rules
174 | Saving tabs configuration
175 | Allow users to see all agencies vehicles on Dispatch map
176 | Allow user to see recurrent orders from all agencies on map
177 | Desactivate free package type input
178 | Allow GPS Device privacy mode change
179 | Allow automatic working time declaration module execusion on current day
180 | Temperature alert
181 | Write package natures
182 | View package natures
183 | Print package natures
184 | Add package history state
185 | Write package reference values
186 | View package reference values
187 | Validate round sessions
188 | Override round session selection
189 | Override round session values
190 | Allow to secure transport
191 | allow batch modification of transport references
192 | Write subcontractor employees
193 | View subcontractor employees
194 | Print subcontractor employees
195 | Allow receipt entry
196 | Write appointments
197 | View appointments
198 | Access to planning appointment module
199 | Access to audit data
200 | Access to audit configuration
201 | Write password security policies
202 | View password security policies
203 | Transports audit
204 | Access to application orderers
205 | Write SSO configurations
206 | View SSO configurations
207 | Write SSO users
208 | View SSO users
209 | Application access management
210 | Display open link in browser button
211 | Data erasure and anonymization
212 | Write EDI equivalence configurations
213 | View EDI equivalence configurations
214 | Access to anonymized data


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v3.0
- Package version: 1.0.0
- Generator version: 7.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.eureka-technology.fr](https://www.eureka-technology.fr)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://mylicense.dispatchapi.dispatch-rts.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressesAPI* | [**GetAddressById**](docs/AddressesAPI.md#getaddressbyid) | **Get** /v3/addresses/by-address-id | Get an address by its identifier
*AddressesAPI* | [**GetAddresses**](docs/AddressesAPI.md#getaddresses) | **Get** /v3/addresses | Get addresses
*AddressesAPI* | [**GetStreetsSuggestions**](docs/AddressesAPI.md#getstreetssuggestions) | **Get** /v3/addresses/streets/suggestions | Get streets suggestions
*AgenciesAPI* | [**GetAgencies**](docs/AgenciesAPI.md#getagencies) | **Get** /v3/agencies | Get agencies
*ApiFlagsAPI* | [**GetApiFlagByKey**](docs/ApiFlagsAPI.md#getapiflagbykey) | **Get** /v3/apiflags/{key} | Get api flag by key
*ApiFlagsAPI* | [**GetApiFlags**](docs/ApiFlagsAPI.md#getapiflags) | **Get** /v3/apiflags | Get api flags
*CitiesAPI* | [**GetBestMatch**](docs/CitiesAPI.md#getbestmatch) | **Get** /v3/cities/best-match | Get best match
*CitiesAPI* | [**GetCities**](docs/CitiesAPI.md#getcities) | **Get** /v3/cities | Get cities
*CitiesAPI* | [**GetCityById**](docs/CitiesAPI.md#getcitybyid) | **Get** /v3/cities/by-city-id | Get a city by its identifier
*CompaniesAPI* | [**GetCompanies**](docs/CompaniesAPI.md#getcompanies) | **Get** /v3/companies | Get companies
*ContractorAgentsAPI* | [**GetContractorAgents**](docs/ContractorAgentsAPI.md#getcontractoragents) | **Get** /v3/contractor-agents | Get contractor agents
*CountriesAPI* | [**GetCountries**](docs/CountriesAPI.md#getcountries) | **Get** /v3/countries | Get countries
*CountriesAPI* | [**GetCountryByCode**](docs/CountriesAPI.md#getcountrybycode) | **Get** /v3/countries/by-country-code | Get country by code
*CustomersAPI* | [**GetCustomParameterValues**](docs/CustomersAPI.md#getcustomparametervalues) | **Get** /v3/customers/{uid}/custom-parameters/{parameterId}/values | Get custom parameter values
*CustomersAPI* | [**GetCustomParameters**](docs/CustomersAPI.md#getcustomparameters) | **Get** /v3/customers/{uid}/custom-parameters | Get custom parameters
*CustomersAPI* | [**GetCustomerByCode**](docs/CustomersAPI.md#getcustomerbycode) | **Get** /v3/customers/by-customer-code | Get a customer by its code
*CustomersAPI* | [**GetCustomerById**](docs/CustomersAPI.md#getcustomerbyid) | **Get** /v3/customers/{uid} | Get a customer by its unique identifier
*CustomersAPI* | [**GetCustomerReferences**](docs/CustomersAPI.md#getcustomerreferences) | **Get** /v3/customers/{uid}/references | Get customer references
*CustomersAPI* | [**GetCustomers**](docs/CustomersAPI.md#getcustomers) | **Get** /v3/customers | Get customers
*DriversAPI* | [**GetDriverByCode**](docs/DriversAPI.md#getdriverbycode) | **Get** /v3/drivers/by-code | Get a driver by its code
*DriversAPI* | [**GetDriverById**](docs/DriversAPI.md#getdriverbyid) | **Get** /v3/drivers/{driverId} | Get a driver by its identifier
*DriversAPI* | [**GetDrivers**](docs/DriversAPI.md#getdrivers) | **Get** /v3/drivers | Get drivers
*FileCategoriesAPI* | [**GetFileCategories**](docs/FileCategoriesAPI.md#getfilecategories) | **Get** /v3/file-categories | Get file categories
*MeAPI* | [**GetMyAccountInformation**](docs/MeAPI.md#getmyaccountinformation) | **Get** /v3/me | Get my account information
*MissionsAPI* | [**CancelMission**](docs/MissionsAPI.md#cancelmission) | **Post** /v3/missions/{uid}/cancellation-process | Cancel mission
*MissionsAPI* | [**CreateMission**](docs/MissionsAPI.md#createmission) | **Post** /v3/missions | Create a mission
*MissionsAPI* | [**GetMissionById**](docs/MissionsAPI.md#getmissionbyid) | **Get** /v3/missions/{uid} | Get a mission by its unique identifier
*MissionsAPI* | [**GetMissionByMissionNumber**](docs/MissionsAPI.md#getmissionbymissionnumber) | **Get** /v3/missions/by-mission-number | Get a mission by its mission number
*MissionsAPI* | [**MissionCreationDryRun**](docs/MissionsAPI.md#missioncreationdryrun) | **Post** /v3/missions/dry-run | Mission creation dry run
*MissionsAPI* | [**MissionUpdateDryRun**](docs/MissionsAPI.md#missionupdatedryrun) | **Patch** /v3/missions/{uid}/dry-run | Mission update dry run
*MissionsAPI* | [**UpdateMission**](docs/MissionsAPI.md#updatemission) | **Patch** /v3/missions/{uid} | Update mission
*OrderersAPI* | [**GetOrderers**](docs/OrderersAPI.md#getorderers) | **Get** /v3/orderers | Get orderers
*PackageNaturesAPI* | [**GetPackageNatureById**](docs/PackageNaturesAPI.md#getpackagenaturebyid) | **Get** /v3/package-natures/{packageNatureId} | Get package nature by id
*PackageNaturesAPI* | [**GetPackageNatureReferenceValues**](docs/PackageNaturesAPI.md#getpackagenaturereferencevalues) | **Get** /v3/package-natures/{packageNatureId}/references/{referenceIndex}/values | Get package nature reference values
*PackageNaturesAPI* | [**GetPackageNatureReferences**](docs/PackageNaturesAPI.md#getpackagenaturereferences) | **Get** /v3/package-natures/{packageNatureId}/references | Get package nature references
*PackageNaturesAPI* | [**GetPackageNatures**](docs/PackageNaturesAPI.md#getpackagenatures) | **Get** /v3/package-natures | Get package natures
*PricingPathsAPI* | [**GetPricingPathAdditionalUnits**](docs/PricingPathsAPI.md#getpricingpathadditionalunits) | **Get** /v3/pricing-paths/{pricingPathId}/sub-service-additional-units | Get pricing path additional units
*PricingPathsAPI* | [**GetPricingPathById**](docs/PricingPathsAPI.md#getpricingpathbyid) | **Get** /v3/pricing-paths/{pricingPathId} | Get pricing path by id
*PricingPathsAPI* | [**GetPricingPathSubServiceUnitPricingSlots**](docs/PricingPathsAPI.md#getpricingpathsubserviceunitpricingslots) | **Get** /v3/pricing-paths/{pricingPathId}/sub-service-unit-pricing-slots | Get pricing path sub service unit pricing slots
*PricingPathsAPI* | [**GetPricingPathSubServices**](docs/PricingPathsAPI.md#getpricingpathsubservices) | **Get** /v3/pricing-paths/{pricingPathId}/sub-services | Get pricing path sub services
*PricingPathsAPI* | [**GetPricingPathUnitPricingSlotEquivalences**](docs/PricingPathsAPI.md#getpricingpathunitpricingslotequivalences) | **Get** /v3/pricing-paths/{pricingPathId}/unit-pricing-slot-equivalences | Get pricing path unit pricing slot equivalences
*PricingPathsAPI* | [**GetPricingPaths**](docs/PricingPathsAPI.md#getpricingpaths) | **Get** /v3/pricing-paths | Get pricing paths
*QuotationsAPI* | [**CancelQuotation**](docs/QuotationsAPI.md#cancelquotation) | **Post** /v3/quotations/{uid}/cancellation-process | Cancel quotation
*QuotationsAPI* | [**CreateMissionFromQuotation**](docs/QuotationsAPI.md#createmissionfromquotation) | **Post** /v3/quotations/{uid}/to-mission-process | Create mission from quotation
*QuotationsAPI* | [**CreateQuotation**](docs/QuotationsAPI.md#createquotation) | **Post** /v3/quotations | Create a quotation
*QuotationsAPI* | [**GetQuotationById**](docs/QuotationsAPI.md#getquotationbyid) | **Get** /v3/quotations/{uid} | Get a quotation by its unique identifier
*QuotationsAPI* | [**GetQuotationByQuotationNumber**](docs/QuotationsAPI.md#getquotationbyquotationnumber) | **Get** /v3/quotations/by-quotation-number | Get a quotation by its quotation number
*QuotationsAPI* | [**QuotationCreationDryRun**](docs/QuotationsAPI.md#quotationcreationdryrun) | **Post** /v3/quotations/dry-run | Quotation creation dry run
*QuotationsAPI* | [**QuotationUpdateDryRun**](docs/QuotationsAPI.md#quotationupdatedryrun) | **Patch** /v3/quotations/{uid}/dry-run | Quotation update dry run
*QuotationsAPI* | [**UpdateQuotation**](docs/QuotationsAPI.md#updatequotation) | **Patch** /v3/quotations/{uid} | Update quotation
*SectorsAPI* | [**GetSectors**](docs/SectorsAPI.md#getsectors) | **Get** /v3/sectors | Get sectors
*ServicesAPI* | [**GetCustomParameterValues**](docs/ServicesAPI.md#getcustomparametervalues) | **Get** /v3/services/{uid}/custom-parameters/{parameterId}/values | Get custom parameter values
*ServicesAPI* | [**GetCustomParameters**](docs/ServicesAPI.md#getcustomparameters) | **Get** /v3/services/{uid}/custom-parameters | Get custom parameters
*ServicesAPI* | [**GetServices**](docs/ServicesAPI.md#getservices) | **Get** /v3/services | Get services
*ServicesAPI* | [**GetSubServices**](docs/ServicesAPI.md#getsubservices) | **Get** /v3/services/{uid}/sub-services | Get sub services
*SubServicesAPI* | [**GetQuantities**](docs/SubServicesAPI.md#getquantities) | **Get** /v3/sub-services/quantities | Get quantities
*SubServicesAPI* | [**GetSubServices**](docs/SubServicesAPI.md#getsubservices) | **Get** /v3/sub-services | Get sub services
*SubstatesAPI* | [**GetSubstateByCode**](docs/SubstatesAPI.md#getsubstatebycode) | **Get** /v3/substates/by-code | Get a substate by its code
*SubstatesAPI* | [**GetSubstates**](docs/SubstatesAPI.md#getsubstates) | **Get** /v3/substates | Get substates
*TeamsAPI* | [**GetTeams**](docs/TeamsAPI.md#getteams) | **Get** /v3/teams | Get teams
*TransportDocumentFilesAPI* | [**DeleteTransportDocumentFile**](docs/TransportDocumentFilesAPI.md#deletetransportdocumentfile) | **Delete** /v3/transports/{uid}/documents/files/{fileId} | Delete transport document file
*TransportDocumentFilesAPI* | [**DownloadTransportDocumentFile**](docs/TransportDocumentFilesAPI.md#downloadtransportdocumentfile) | **Get** /v3/transports/{uid}/documents/files/{fileId}/data | Download transport document file
*TransportDocumentFilesAPI* | [**GetTransportDocumentFileById**](docs/TransportDocumentFilesAPI.md#gettransportdocumentfilebyid) | **Get** /v3/transports/{uid}/documents/files/{fileId} | Get transport document file by id
*TransportDocumentFilesAPI* | [**GetTransportDocumentFiles**](docs/TransportDocumentFilesAPI.md#gettransportdocumentfiles) | **Get** /v3/transports/{uid}/documents/files | Get transport document files
*TransportDocumentFilesAPI* | [**UploadTransportDocumentFile**](docs/TransportDocumentFilesAPI.md#uploadtransportdocumentfile) | **Post** /v3/transports/{uid}/documents/files | Upload transport document file
*TransportDocumentLinksAPI* | [**CreateTransportDocumentLink**](docs/TransportDocumentLinksAPI.md#createtransportdocumentlink) | **Post** /v3/transports/{uid}/documents/links | Create a transport document link
*TransportDocumentLinksAPI* | [**DeleteTransportDocumentLink**](docs/TransportDocumentLinksAPI.md#deletetransportdocumentlink) | **Delete** /v3/transports/{uid}/documents/links/{linkId} | Delete transport document link
*TransportDocumentLinksAPI* | [**GetTransportDocumentLinkById**](docs/TransportDocumentLinksAPI.md#gettransportdocumentlinkbyid) | **Get** /v3/transports/{uid}/documents/links/{linkId} | Get transport document link by id
*TransportDocumentLinksAPI* | [**GetTransportDocumentLinks**](docs/TransportDocumentLinksAPI.md#gettransportdocumentlinks) | **Get** /v3/transports/{uid}/documents/links | Get transport document links
*TransportDocumentReportsAPI* | [**DownloadTransportDocumentReport**](docs/TransportDocumentReportsAPI.md#downloadtransportdocumentreport) | **Get** /v3/transport-document-reports/{uid}/data | Download a transport document report by its unique identifier
*TransportDocumentReportsAPI* | [**GenerateTransportDocumentReport**](docs/TransportDocumentReportsAPI.md#generatetransportdocumentreport) | **Post** /v3/transport-document-reports | Generate transport document reports
*TransportDocumentReportsAPI* | [**GetTransportDocumentReportById**](docs/TransportDocumentReportsAPI.md#gettransportdocumentreportbyid) | **Get** /v3/transport-document-reports/{uid} | Get a transport document report by its unique identifier
*TransportsAPI* | [**ApplySubState**](docs/TransportsAPI.md#applysubstate) | **Post** /v3/transports/{uid}/substates | Apply a sub state to the transport
*TransportsAPI* | [**ApplyTransportCancelledStatus**](docs/TransportsAPI.md#applytransportcancelledstatus) | **Post** /v3/transports/{uid}/cancellation-process | Apply transport cancelled status
*TransportsAPI* | [**ApplyTransportDeliveredStatus**](docs/TransportsAPI.md#applytransportdeliveredstatus) | **Post** /v3/transports/{uid}/delivery-process | Apply transport delivered state
*TransportsAPI* | [**ApplyTransportPickedUpStatus**](docs/TransportsAPI.md#applytransportpickedupstatus) | **Post** /v3/transports/{uid}/pickup-process | Apply transport pickup state
*TransportsAPI* | [**ApplyTransportTerminatedStatus**](docs/TransportsAPI.md#applytransportterminatedstatus) | **Post** /v3/transports/{uid}/termination-process | Apply transport terminated state
*TransportsAPI* | [**CheckTransportAssignmentsConstraints**](docs/TransportsAPI.md#checktransportassignmentsconstraints) | **Patch** /v3/transports/{uid}/assignment/constraints-validation | Check transport assignment&#39;s constraints
*TransportsAPI* | [**DefragmentTransport**](docs/TransportsAPI.md#defragmenttransport) | **Post** /v3/transports/{uid}/defragmentation-process | Defragment transport
*TransportsAPI* | [**FragmentTransport**](docs/TransportsAPI.md#fragmenttransport) | **Post** /v3/transports/{uid}/fragmentation-process | Fragment transport
*TransportsAPI* | [**GetTransportAirData**](docs/TransportsAPI.md#gettransportairdata) | **Get** /v3/transports/{uid}/air-data | Get transport air data
*TransportsAPI* | [**GetTransportBillAddress**](docs/TransportsAPI.md#gettransportbilladdress) | **Get** /v3/transports/{uid}/transport-bill-address | Get transport bill address
*TransportsAPI* | [**GetTransportByUniqueIdentifier**](docs/TransportsAPI.md#gettransportbyuniqueidentifier) | **Get** /v3/transports/{uid} | Get a transport by uid
*TransportsAPI* | [**GetTransportCommunicationConfiguration**](docs/TransportsAPI.md#gettransportcommunicationconfiguration) | **Get** /v3/transports/{uid}/communication-configuration | Get transport communication configuration
*TransportsAPI* | [**GetTransportCustomerCustomParameters**](docs/TransportsAPI.md#gettransportcustomercustomparameters) | **Get** /v3/transports/{uid}/customer-custom-parameters | Get transport customer custom parameters
*TransportsAPI* | [**GetTransportDangerousGoods**](docs/TransportsAPI.md#gettransportdangerousgoods) | **Get** /v3/transports/{uid}/dangerous-goods | Get transport dangerous goods
*TransportsAPI* | [**GetTransportHistoryLines**](docs/TransportsAPI.md#gettransporthistorylines) | **Get** /v3/transports/{uid}/history | Get transport history lines
*TransportsAPI* | [**GetTransportPackageLines**](docs/TransportsAPI.md#gettransportpackagelines) | **Get** /v3/transports/{uid}/package-lines | Get transport package lines
*TransportsAPI* | [**GetTransportPricing**](docs/TransportsAPI.md#gettransportpricing) | **Get** /v3/transports/{uid}/pricing | Get transport pricing
*TransportsAPI* | [**GetTransportRoute**](docs/TransportsAPI.md#gettransportroute) | **Get** /v3/transports/{uid}/transport-route | Get transport route
*TransportsAPI* | [**GetTransportServiceCustomParameters**](docs/TransportsAPI.md#gettransportservicecustomparameters) | **Get** /v3/transports/{uid}/service-custom-parameters | Get transport service custom parameters
*TransportsAPI* | [**GetTransportSubServices**](docs/TransportsAPI.md#gettransportsubservices) | **Get** /v3/transports/{uid}/subservices | Get transport sub services
*TransportsAPI* | [**GetTransportTotalBulkSizes**](docs/TransportsAPI.md#gettransporttotalbulksizes) | **Get** /v3/transports/{uid}/total-bulk-sizes | Get transport total bulk sizes
*TransportsAPI* | [**GetTransports**](docs/TransportsAPI.md#gettransports) | **Get** /v3/transports | Get transports
*TransportsAPI* | [**TransportUpdateDryRun**](docs/TransportsAPI.md#transportupdatedryrun) | **Patch** /v3/transports/{uid}/dry-run | Transport update dry run
*TransportsAPI* | [**UpdateTransport**](docs/TransportsAPI.md#updatetransport) | **Patch** /v3/transports/{uid} | Update transport
*TransportsAPI* | [**UpdateTransportAssignment**](docs/TransportsAPI.md#updatetransportassignment) | **Patch** /v3/transports/{uid}/assignment | Update transport assignment
*TransportsAPI* | [**UpdateTransportAssignmentStatus**](docs/TransportsAPI.md#updatetransportassignmentstatus) | **Post** /v3/transports/{uid}/assignment/status | Update transport assignment status
*UnitsAPI* | [**GetUnits**](docs/UnitsAPI.md#getunits) | **Get** /v3/units | Get units
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /v3/users | Get users
*VehicleTypesAPI* | [**GetVehicles**](docs/VehicleTypesAPI.md#getvehicles) | **Get** /v3/vehicle-types | Get vehicle types
*VehiclesAPI* | [**GetVehicles**](docs/VehiclesAPI.md#getvehicles) | **Get** /v3/vehicles | Get vehicles
*VehiclesAPI* | [**GetVehiclesSubcontractors**](docs/VehiclesAPI.md#getvehiclessubcontractors) | **Get** /v3/vehicles/{uid}/subcontractors | Get the subcontractors linked to the vehicle


## Documentation For Models

 - [AddressDto](docs/AddressDto.md)
 - [AgencyDto](docs/AgencyDto.md)
 - [ApiFlagsDto](docs/ApiFlagsDto.md)
 - [ApiFlagsDtoApiFlag](docs/ApiFlagsDtoApiFlag.md)
 - [ApplyAssignmentStatusChangeCommand](docs/ApplyAssignmentStatusChangeCommand.md)
 - [ApplyCancelledStatusCommand](docs/ApplyCancelledStatusCommand.md)
 - [ApplyDeliveredStatusCommand](docs/ApplyDeliveredStatusCommand.md)
 - [ApplyPickedUpStatusCommand](docs/ApplyPickedUpStatusCommand.md)
 - [ApplyTerminatedStatusCommand](docs/ApplyTerminatedStatusCommand.md)
 - [ApplyTransportSubStateCommand](docs/ApplyTransportSubStateCommand.md)
 - [AssignTransportDto](docs/AssignTransportDto.md)
 - [AssignTransportDtoApiBehaviorDto](docs/AssignTransportDtoApiBehaviorDto.md)
 - [CancelMissionCommand](docs/CancelMissionCommand.md)
 - [CancelQuotationCommand](docs/CancelQuotationCommand.md)
 - [CappedCollectionDtoCustomParameterValueDto](docs/CappedCollectionDtoCustomParameterValueDto.md)
 - [CappedCollectionDtoPackageLineDto](docs/CappedCollectionDtoPackageLineDto.md)
 - [CappedCollectionDtoPackageLineReferenceValueDto](docs/CappedCollectionDtoPackageLineReferenceValueDto.md)
 - [CappedCollectionDtoPackageLineSizeDto](docs/CappedCollectionDtoPackageLineSizeDto.md)
 - [CappedCollectionDtoPackageNatureBulkSizeDto](docs/CappedCollectionDtoPackageNatureBulkSizeDto.md)
 - [CappedCollectionDtoPackageNatureReferenceDto](docs/CappedCollectionDtoPackageNatureReferenceDto.md)
 - [CappedCollectionDtoPackageNatureReferenceValueDto](docs/CappedCollectionDtoPackageNatureReferenceValueDto.md)
 - [CappedCollectionDtoPackagesTotalBulkSizeDto](docs/CappedCollectionDtoPackagesTotalBulkSizeDto.md)
 - [CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto](docs/CappedCollectionDtoPricingPathSubServiceAdditionalUnitDto.md)
 - [CappedCollectionDtoPricingPathSubServiceDto](docs/CappedCollectionDtoPricingPathSubServiceDto.md)
 - [CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto](docs/CappedCollectionDtoPricingPathSubServiceUnitPricingSlotDto.md)
 - [CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto](docs/CappedCollectionDtoPricingPathUnitPricingSlotEquivalenceDto.md)
 - [CappedCollectionDtoSubServiceQuantityDto](docs/CappedCollectionDtoSubServiceQuantityDto.md)
 - [CappedCollectionDtoTransportDto](docs/CappedCollectionDtoTransportDto.md)
 - [CappedCollectionDtoTransportHistoryAnomalyDto](docs/CappedCollectionDtoTransportHistoryAnomalyDto.md)
 - [CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto](docs/CappedCollectionDtoTransportPricingDtoTransportSubServicePricingDto.md)
 - [CappedCollectionDtoTransportSubServiceDto](docs/CappedCollectionDtoTransportSubServiceDto.md)
 - [CheckTransportAssignmentConstraintsResultDto](docs/CheckTransportAssignmentConstraintsResultDto.md)
 - [CheckTransportAssignmentConstraintsResultDtoAssignmentConstraintDto](docs/CheckTransportAssignmentConstraintsResultDtoAssignmentConstraintDto.md)
 - [CityDto](docs/CityDto.md)
 - [CommunicationTypeConfigurationDto](docs/CommunicationTypeConfigurationDto.md)
 - [CompanyDto](docs/CompanyDto.md)
 - [ContractorAgentDto](docs/ContractorAgentDto.md)
 - [CountryDto](docs/CountryDto.md)
 - [CreateMissionCommand](docs/CreateMissionCommand.md)
 - [CreateMissionCommandApiBehaviorDto](docs/CreateMissionCommandApiBehaviorDto.md)
 - [CreateMissionFromQuotationCommand](docs/CreateMissionFromQuotationCommand.md)
 - [CreateMissionFromQuotationResultDto](docs/CreateMissionFromQuotationResultDto.md)
 - [CreateMissionResultDto](docs/CreateMissionResultDto.md)
 - [CreateQuotationCommand](docs/CreateQuotationCommand.md)
 - [CreateQuotationCommandApiBehaviorDto](docs/CreateQuotationCommandApiBehaviorDto.md)
 - [CreateQuotationResultDto](docs/CreateQuotationResultDto.md)
 - [CreateTransportDocumentLinkCommand](docs/CreateTransportDocumentLinkCommand.md)
 - [CreateTransportDocumentLinkResultDto](docs/CreateTransportDocumentLinkResultDto.md)
 - [CreateTransportDto](docs/CreateTransportDto.md)
 - [CreateTransportDtoCashOnDeliveryDto](docs/CreateTransportDtoCashOnDeliveryDto.md)
 - [CreateTransportDtoCommentDto](docs/CreateTransportDtoCommentDto.md)
 - [CreateTransportDtoCustomParameterDto](docs/CreateTransportDtoCustomParameterDto.md)
 - [CreateTransportDtoDangerousGoodDto](docs/CreateTransportDtoDangerousGoodDto.md)
 - [CreateTransportDtoPackingDto](docs/CreateTransportDtoPackingDto.md)
 - [CreateTransportDtoPackingDtoPackageLineDto](docs/CreateTransportDtoPackingDtoPackageLineDto.md)
 - [CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto](docs/CreateTransportDtoPackingDtoPackageLineDtoPackageReferenceDto.md)
 - [CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto](docs/CreateTransportDtoPackingDtoPackageLineDtoPackageSizeDto.md)
 - [CreateTransportDtoPackingDtoTotalBulkSizeDto](docs/CreateTransportDtoPackingDtoTotalBulkSizeDto.md)
 - [CreateTransportDtoStepDto](docs/CreateTransportDtoStepDto.md)
 - [CreateTransportDtoStepDtoStepAddressDto](docs/CreateTransportDtoStepDtoStepAddressDto.md)
 - [CreateTransportDtoStepDtoStepConstraintsDto](docs/CreateTransportDtoStepDtoStepConstraintsDto.md)
 - [CreateTransportDtoStepDtoStepContactDto](docs/CreateTransportDtoStepDtoStepContactDto.md)
 - [CreateTransportDtoStepDtoStepDateDto](docs/CreateTransportDtoStepDtoStepDateDto.md)
 - [CreateTransportDtoStepDtoStepDateDtoEndSlotDto](docs/CreateTransportDtoStepDtoStepDateDtoEndSlotDto.md)
 - [CreateTransportDtoStepDtoStepDateDtoStartSlotDto](docs/CreateTransportDtoStepDtoStepDateDtoStartSlotDto.md)
 - [CreateTransportDtoSubServiceDto](docs/CreateTransportDtoSubServiceDto.md)
 - [CreateTransportDtoTransportBillAddressDto](docs/CreateTransportDtoTransportBillAddressDto.md)
 - [CreateTransportResultDto](docs/CreateTransportResultDto.md)
 - [CustomParameterDto](docs/CustomParameterDto.md)
 - [CustomParameterValueDto](docs/CustomParameterValueDto.md)
 - [CustomerCommunicationConfigurationDto](docs/CustomerCommunicationConfigurationDto.md)
 - [CustomerDto](docs/CustomerDto.md)
 - [CustomerOperationalConfigurationDto](docs/CustomerOperationalConfigurationDto.md)
 - [CustomerReferenceDto](docs/CustomerReferenceDto.md)
 - [CustomerReferencesConfigurationDto](docs/CustomerReferencesConfigurationDto.md)
 - [DefragmentTransportCommand](docs/DefragmentTransportCommand.md)
 - [DeliveryCommunicationConfigurationDto](docs/DeliveryCommunicationConfigurationDto.md)
 - [DeliveryDataOnStatusAppliedDto](docs/DeliveryDataOnStatusAppliedDto.md)
 - [DocumentReportDto](docs/DocumentReportDto.md)
 - [DriverDto](docs/DriverDto.md)
 - [DriverIncludedDto](docs/DriverIncludedDto.md)
 - [FileCategoryDto](docs/FileCategoryDto.md)
 - [FragmentTransportCommand](docs/FragmentTransportCommand.md)
 - [FragmentTransportCommandStepDto](docs/FragmentTransportCommandStepDto.md)
 - [GenerateTransportDocumentReportCommand](docs/GenerateTransportDocumentReportCommand.md)
 - [GenerateTransportDocumentReportResultDto](docs/GenerateTransportDocumentReportResultDto.md)
 - [GeocodingDto](docs/GeocodingDto.md)
 - [IPagedResourceListAddressDto](docs/IPagedResourceListAddressDto.md)
 - [IPagedResourceListAgencyDto](docs/IPagedResourceListAgencyDto.md)
 - [IPagedResourceListCityDto](docs/IPagedResourceListCityDto.md)
 - [IPagedResourceListCompanyDto](docs/IPagedResourceListCompanyDto.md)
 - [IPagedResourceListContractorAgentDto](docs/IPagedResourceListContractorAgentDto.md)
 - [IPagedResourceListCountryDto](docs/IPagedResourceListCountryDto.md)
 - [IPagedResourceListCustomParameterDto](docs/IPagedResourceListCustomParameterDto.md)
 - [IPagedResourceListCustomParameterValueDto](docs/IPagedResourceListCustomParameterValueDto.md)
 - [IPagedResourceListCustomerDto](docs/IPagedResourceListCustomerDto.md)
 - [IPagedResourceListCustomerReferenceDto](docs/IPagedResourceListCustomerReferenceDto.md)
 - [IPagedResourceListDriverDto](docs/IPagedResourceListDriverDto.md)
 - [IPagedResourceListFileCategoryDto](docs/IPagedResourceListFileCategoryDto.md)
 - [IPagedResourceListOrdererDto](docs/IPagedResourceListOrdererDto.md)
 - [IPagedResourceListPackageLineDto](docs/IPagedResourceListPackageLineDto.md)
 - [IPagedResourceListPackageNatureDto](docs/IPagedResourceListPackageNatureDto.md)
 - [IPagedResourceListPackageNatureReferenceDto](docs/IPagedResourceListPackageNatureReferenceDto.md)
 - [IPagedResourceListPackageNatureReferenceValueDto](docs/IPagedResourceListPackageNatureReferenceValueDto.md)
 - [IPagedResourceListPackagesTotalBulkSizeDto](docs/IPagedResourceListPackagesTotalBulkSizeDto.md)
 - [IPagedResourceListPricingPathDto](docs/IPagedResourceListPricingPathDto.md)
 - [IPagedResourceListPricingPathSubServiceAdditionalUnitDto](docs/IPagedResourceListPricingPathSubServiceAdditionalUnitDto.md)
 - [IPagedResourceListPricingPathSubServiceDto](docs/IPagedResourceListPricingPathSubServiceDto.md)
 - [IPagedResourceListPricingPathSubServiceUnitPricingSlotDto](docs/IPagedResourceListPricingPathSubServiceUnitPricingSlotDto.md)
 - [IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto](docs/IPagedResourceListPricingPathUnitPricingSlotEquivalenceDto.md)
 - [IPagedResourceListSectorDto](docs/IPagedResourceListSectorDto.md)
 - [IPagedResourceListServiceDto](docs/IPagedResourceListServiceDto.md)
 - [IPagedResourceListServiceSubServiceDto](docs/IPagedResourceListServiceSubServiceDto.md)
 - [IPagedResourceListSubServiceDto](docs/IPagedResourceListSubServiceDto.md)
 - [IPagedResourceListSubServiceQuantityDto](docs/IPagedResourceListSubServiceQuantityDto.md)
 - [IPagedResourceListSubStateDto](docs/IPagedResourceListSubStateDto.md)
 - [IPagedResourceListTeamDto](docs/IPagedResourceListTeamDto.md)
 - [IPagedResourceListTransportCustomParameterDto](docs/IPagedResourceListTransportCustomParameterDto.md)
 - [IPagedResourceListTransportDocumentFileDto](docs/IPagedResourceListTransportDocumentFileDto.md)
 - [IPagedResourceListTransportDocumentLinkDto](docs/IPagedResourceListTransportDocumentLinkDto.md)
 - [IPagedResourceListTransportDto](docs/IPagedResourceListTransportDto.md)
 - [IPagedResourceListTransportHistoryLineDto](docs/IPagedResourceListTransportHistoryLineDto.md)
 - [IPagedResourceListTransportSubServiceDto](docs/IPagedResourceListTransportSubServiceDto.md)
 - [IPagedResourceListUnitDto](docs/IPagedResourceListUnitDto.md)
 - [IPagedResourceListUserDto](docs/IPagedResourceListUserDto.md)
 - [IPagedResourceListVehicleDto](docs/IPagedResourceListVehicleDto.md)
 - [IPagedResourceListVehicleTypeDto](docs/IPagedResourceListVehicleTypeDto.md)
 - [IResourceListStreetSuggestionDto](docs/IResourceListStreetSuggestionDto.md)
 - [IResourceListTransportDangerousGoodDto](docs/IResourceListTransportDangerousGoodDto.md)
 - [MergeableTransportDto](docs/MergeableTransportDto.md)
 - [MissionDto](docs/MissionDto.md)
 - [MissionEntryDryRunDto](docs/MissionEntryDryRunDto.md)
 - [MyAccountInfoDto](docs/MyAccountInfoDto.md)
 - [NonBlockingErrorDto](docs/NonBlockingErrorDto.md)
 - [OrdererDto](docs/OrdererDto.md)
 - [PackageLineDto](docs/PackageLineDto.md)
 - [PackageLineReferenceValueDto](docs/PackageLineReferenceValueDto.md)
 - [PackageLineSizeDto](docs/PackageLineSizeDto.md)
 - [PackageNatureBulkSizeDto](docs/PackageNatureBulkSizeDto.md)
 - [PackageNatureDto](docs/PackageNatureDto.md)
 - [PackageNatureReferenceDto](docs/PackageNatureReferenceDto.md)
 - [PackageNatureReferenceValueDto](docs/PackageNatureReferenceValueDto.md)
 - [PackagesTotalBulkSizeDto](docs/PackagesTotalBulkSizeDto.md)
 - [PagingInformations](docs/PagingInformations.md)
 - [PickupCommunicationConfigurationDto](docs/PickupCommunicationConfigurationDto.md)
 - [PickupDataOnStatusAppliedDto](docs/PickupDataOnStatusAppliedDto.md)
 - [PositionDto](docs/PositionDto.md)
 - [PricingPathDto](docs/PricingPathDto.md)
 - [PricingPathSubServiceAdditionalUnitDto](docs/PricingPathSubServiceAdditionalUnitDto.md)
 - [PricingPathSubServiceDto](docs/PricingPathSubServiceDto.md)
 - [PricingPathSubServiceUnitPricingSlotDto](docs/PricingPathSubServiceUnitPricingSlotDto.md)
 - [PricingPathUnitPricingSlotEquivalenceDto](docs/PricingPathUnitPricingSlotEquivalenceDto.md)
 - [PricingPointDto](docs/PricingPointDto.md)
 - [QuotationDto](docs/QuotationDto.md)
 - [QuotationEntryDryRunDto](docs/QuotationEntryDryRunDto.md)
 - [ReportPaperOptionsDto](docs/ReportPaperOptionsDto.md)
 - [SearchFlagsDto](docs/SearchFlagsDto.md)
 - [SearchFlagsDtoLastSubStateForFamilyCodeDto](docs/SearchFlagsDtoLastSubStateForFamilyCodeDto.md)
 - [SectorDto](docs/SectorDto.md)
 - [ServiceDto](docs/ServiceDto.md)
 - [ServiceFamilyDto](docs/ServiceFamilyDto.md)
 - [ServiceSubServiceDto](docs/ServiceSubServiceDto.md)
 - [ServiceTypeDto](docs/ServiceTypeDto.md)
 - [SimulateMissionCreationDryRunCommand](docs/SimulateMissionCreationDryRunCommand.md)
 - [SimulateMissionCreationDryRunCommandApiBehaviorDto](docs/SimulateMissionCreationDryRunCommandApiBehaviorDto.md)
 - [SimulateQuotationCreationDryRunCommand](docs/SimulateQuotationCreationDryRunCommand.md)
 - [SimulateQuotationCreationDryRunCommandApiBehaviorDto](docs/SimulateQuotationCreationDryRunCommandApiBehaviorDto.md)
 - [SlotDto](docs/SlotDto.md)
 - [StreetSuggestionDto](docs/StreetSuggestionDto.md)
 - [SubContractorCommunicationConfigurationDto](docs/SubContractorCommunicationConfigurationDto.md)
 - [SubServiceDto](docs/SubServiceDto.md)
 - [SubServiceQuantityDto](docs/SubServiceQuantityDto.md)
 - [SubStateDto](docs/SubStateDto.md)
 - [TeamDto](docs/TeamDto.md)
 - [TransportAirDataDto](docs/TransportAirDataDto.md)
 - [TransportAirDataDtoTransportAirWaybillDto](docs/TransportAirDataDtoTransportAirWaybillDto.md)
 - [TransportBillAddressDto](docs/TransportBillAddressDto.md)
 - [TransportCashOnDeliveryDto](docs/TransportCashOnDeliveryDto.md)
 - [TransportCommentDto](docs/TransportCommentDto.md)
 - [TransportCommunicationConfigurationDto](docs/TransportCommunicationConfigurationDto.md)
 - [TransportCustomParameterDto](docs/TransportCustomParameterDto.md)
 - [TransportDangerousGoodDto](docs/TransportDangerousGoodDto.md)
 - [TransportDeliveryDateTimeSlotDto](docs/TransportDeliveryDateTimeSlotDto.md)
 - [TransportDeliveryStepDto](docs/TransportDeliveryStepDto.md)
 - [TransportDocumentFileDto](docs/TransportDocumentFileDto.md)
 - [TransportDocumentLinkDto](docs/TransportDocumentLinkDto.md)
 - [TransportDto](docs/TransportDto.md)
 - [TransportEntryDryRunDto](docs/TransportEntryDryRunDto.md)
 - [TransportEntryDryRunDtoCashOnDeliveryDto](docs/TransportEntryDryRunDtoCashOnDeliveryDto.md)
 - [TransportEntryDryRunDtoCommentDto](docs/TransportEntryDryRunDtoCommentDto.md)
 - [TransportEntryDryRunDtoCustomParameterDto](docs/TransportEntryDryRunDtoCustomParameterDto.md)
 - [TransportEntryDryRunDtoDangerousGoodDto](docs/TransportEntryDryRunDtoDangerousGoodDto.md)
 - [TransportEntryDryRunDtoPackingDto](docs/TransportEntryDryRunDtoPackingDto.md)
 - [TransportEntryDryRunDtoPackingDtoPackageLineDto](docs/TransportEntryDryRunDtoPackingDtoPackageLineDto.md)
 - [TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto](docs/TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageReferenceDto.md)
 - [TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto](docs/TransportEntryDryRunDtoPackingDtoPackageLineDtoPackageSizeDto.md)
 - [TransportEntryDryRunDtoPackingDtoTotalBulkSizeDto](docs/TransportEntryDryRunDtoPackingDtoTotalBulkSizeDto.md)
 - [TransportEntryDryRunDtoStepDto](docs/TransportEntryDryRunDtoStepDto.md)
 - [TransportEntryDryRunDtoStepDtoStepAddressDto](docs/TransportEntryDryRunDtoStepDtoStepAddressDto.md)
 - [TransportEntryDryRunDtoStepDtoStepConstraintsDto](docs/TransportEntryDryRunDtoStepDtoStepConstraintsDto.md)
 - [TransportEntryDryRunDtoStepDtoStepContactDto](docs/TransportEntryDryRunDtoStepDtoStepContactDto.md)
 - [TransportEntryDryRunDtoStepDtoStepDateDto](docs/TransportEntryDryRunDtoStepDtoStepDateDto.md)
 - [TransportEntryDryRunDtoStepDtoStepDateDtoEndSlotDto](docs/TransportEntryDryRunDtoStepDtoStepDateDtoEndSlotDto.md)
 - [TransportEntryDryRunDtoStepDtoStepDateDtoStartSlotDto](docs/TransportEntryDryRunDtoStepDtoStepDateDtoStartSlotDto.md)
 - [TransportEntryDryRunDtoSubServiceDto](docs/TransportEntryDryRunDtoSubServiceDto.md)
 - [TransportEntryDryRunDtoTransportBillAddressDto](docs/TransportEntryDryRunDtoTransportBillAddressDto.md)
 - [TransportHistoryAnomalyDto](docs/TransportHistoryAnomalyDto.md)
 - [TransportHistoryLineDto](docs/TransportHistoryLineDto.md)
 - [TransportIncludedDto](docs/TransportIncludedDto.md)
 - [TransportPackingDto](docs/TransportPackingDto.md)
 - [TransportPickupDateTimeSlotDto](docs/TransportPickupDateTimeSlotDto.md)
 - [TransportPickupStepDto](docs/TransportPickupStepDto.md)
 - [TransportPricingDto](docs/TransportPricingDto.md)
 - [TransportPricingDtoTransportSubServicePricingDto](docs/TransportPricingDtoTransportSubServicePricingDto.md)
 - [TransportRouteDto](docs/TransportRouteDto.md)
 - [TransportRouteDtoTransportRoutePositionsDto](docs/TransportRouteDtoTransportRoutePositionsDto.md)
 - [TransportStepAddressDto](docs/TransportStepAddressDto.md)
 - [TransportStepConstraintsDto](docs/TransportStepConstraintsDto.md)
 - [TransportStepContactDto](docs/TransportStepContactDto.md)
 - [TransportSubServiceDto](docs/TransportSubServiceDto.md)
 - [UnitDto](docs/UnitDto.md)
 - [UpdateMissionDto](docs/UpdateMissionDto.md)
 - [UpdateMissionDtoApiBehaviorDto](docs/UpdateMissionDtoApiBehaviorDto.md)
 - [UpdateQuotationDto](docs/UpdateQuotationDto.md)
 - [UpdateQuotationDtoApiBehaviorDto](docs/UpdateQuotationDtoApiBehaviorDto.md)
 - [UpdateTransportBaseDtoCashOnDeliveryDto](docs/UpdateTransportBaseDtoCashOnDeliveryDto.md)
 - [UpdateTransportBaseDtoCommentDto](docs/UpdateTransportBaseDtoCommentDto.md)
 - [UpdateTransportBaseDtoCustomParameterDto](docs/UpdateTransportBaseDtoCustomParameterDto.md)
 - [UpdateTransportBaseDtoDangerousGoodDto](docs/UpdateTransportBaseDtoDangerousGoodDto.md)
 - [UpdateTransportBaseDtoPackingDto](docs/UpdateTransportBaseDtoPackingDto.md)
 - [UpdateTransportBaseDtoPackingDtoPackageLineDto](docs/UpdateTransportBaseDtoPackingDtoPackageLineDto.md)
 - [UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto](docs/UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageReferenceDto.md)
 - [UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto](docs/UpdateTransportBaseDtoPackingDtoPackageLineDtoPackageSizeDto.md)
 - [UpdateTransportBaseDtoPackingDtoTotalBulkSizeDto](docs/UpdateTransportBaseDtoPackingDtoTotalBulkSizeDto.md)
 - [UpdateTransportBaseDtoStepDto](docs/UpdateTransportBaseDtoStepDto.md)
 - [UpdateTransportBaseDtoStepDtoStepAddressDto](docs/UpdateTransportBaseDtoStepDtoStepAddressDto.md)
 - [UpdateTransportBaseDtoStepDtoStepConstraintsDto](docs/UpdateTransportBaseDtoStepDtoStepConstraintsDto.md)
 - [UpdateTransportBaseDtoStepDtoStepContactDto](docs/UpdateTransportBaseDtoStepDtoStepContactDto.md)
 - [UpdateTransportBaseDtoStepDtoStepDateDto](docs/UpdateTransportBaseDtoStepDtoStepDateDto.md)
 - [UpdateTransportBaseDtoStepDtoStepDateDtoEndSlotDto](docs/UpdateTransportBaseDtoStepDtoStepDateDtoEndSlotDto.md)
 - [UpdateTransportBaseDtoStepDtoStepDateDtoStartSlotDto](docs/UpdateTransportBaseDtoStepDtoStepDateDtoStartSlotDto.md)
 - [UpdateTransportBaseDtoSubServiceDto](docs/UpdateTransportBaseDtoSubServiceDto.md)
 - [UpdateTransportBaseDtoTransportBillAddressDto](docs/UpdateTransportBaseDtoTransportBillAddressDto.md)
 - [UpdateTransportDto](docs/UpdateTransportDto.md)
 - [UpdateTransportDtoApiBehaviorDto](docs/UpdateTransportDtoApiBehaviorDto.md)
 - [UploadTransportDocumentFileResultDto](docs/UploadTransportDocumentFileResultDto.md)
 - [UserDto](docs/UserDto.md)
 - [VehicleDto](docs/VehicleDto.md)
 - [VehicleTypeDto](docs/VehicleTypeDto.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



