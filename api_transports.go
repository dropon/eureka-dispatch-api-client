/*
Dispatch Api 3.0

This API is dedicated to operators.    Get API url and your client credentials from [integrators website](https://integrators.eureka-technology.fr/).      You can download the Open Api specifications (swagger json document) in the link mentioned above.    Please use our [customer support portal](https://eureka-technology.atlassian.net/servicedesk/customer/portals) for questions or issues.  If you do not have a support account yet, please reach out to your contact at Eureka-Technology to create one.  If you don't have a dedicated contact, you can send an email to support@eureka-technology.com.    # Authentication      ## Access tokens    This Api uses OAuth2 authentication.      All the requests to the api, except authentication ones, must include a valid bearer token.    To obtain an access token, call the _/token_ endpoint, provide an Authorization header with a basic authentication type, constructed using your client_id and client_secret.    The body must be a form url encoded with the parameters corresponding to the authentication flow (see grant types section for more details).    The api returns a response containing the access token, its expiration in seconds, a refresh token, and other properties. Example :    ```json  {      \"access_token\": \"ey...\",      \"token_type\": \"bearer\",      \"expires_in\": 3599,      \"refresh_token\": \"...\",      \".issued\": \"(date)\",      \".expires\": \"(date)\"  }  ```  When the client id or client secret is invalid, a 400 Bad request is returned :  ```json  {      \"error\": \"invalid_client\",      \"error_description\": \"The client id or secret is incorrect.\"  }  ```  The returned access token can then be used in the requests to the api, it must be included in the Authorization header, example : `Authorization: Bearer ey...`.    The access token is valid 1 hour by default, this value can change depending on the carrier.      ### Example using postman      Here is an example of token request using postman (to be adapted according to the grant type).      The request header contains the basic authentication.    ![Authentication authorization header](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-authorization-header.png)    The request body is a form url encoded containing the grant type and the additional parameters (depending on the grant type).    ![Authentication body with password](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-body-with-password.png)    The response contains the access token and additional information.    ![Authentication response](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-response.png)    ## Refresh token      In _/token_ response, you receive an access token and a refresh token.    The refresh token allows to get a new access token without providing the original credentials.    When an access token has expired, you can exchange the refresh token for a new access token.    The refresh token is valid 24 hours by default, this value can change depending on the carrier.    To generate a new access token using this refresh token, use refresh_token grant type. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=refresh_token&refresh_token=your_refresh_token  ```      You can invalidate the refresh token by calling _/revoke_ endpoint. Example :    ```  POST /revoke    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    token=your_refresh_token  ```      ### Refresh token limitation    There is a limit in the number of active refresh tokens per user. It is 10000 by default, this value can change depending on the carrier.    If this limit is reached, you will receive the error _max_number_of_sessions_reached_, you will need to wait for refresh token expiration before being able to get a new access token for the user.    To avoid the limit, reuse the access token, do not ask for a new one for each request to the Api.      ## Grant types    ### Password flow      Use password grant to connect as a user, using its credentials : provide the parameters grant_type _password_, username and password. Example :    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=password&username=user_username&password=user_password  ```      When user's credentials are invalid, a 400 Bad request is returned :  ```json  {      \"error\": \"invalid_grant\",      \"error_description\": \"Unknown user or incorrect password.\"  }  ```      ### Client credentials      Use client credentials to connect as an application, outside of the context of a user : provide the parameters grant_type _client_credentials_. Example :    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=client_credentials  ```      ### SSO      To connect with an external SSO provider, SSO settings must be configured in Dispatch, and you will have a SSO configuration code for each provider.    SSO users must already be configured in Dispatch.      #### SSO Authorization code      This is the standard method to use an external OAUTH SSO provider configured in Dispatch. Use this method if your application does not have its own client id and client secret in the external SSO provider.      - First, your application will need to open a browser to send the user to the external SSO oauth provider.      To get the authorization url, use the operation _sso/oauth-authorization-url_ and provide the parameters : ssoCode, redirectUri and state, and provide bearer token created with a client credentials authentication.      - When the user connects, the oauth provider will redirect back to your application url and provide an authorization code and the state you provided      - Call token endpoint in the API with grant_type authorization_code_TheSsoConfigurationCode and parameters authorization_code and redirect_uri. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=authorization_code_ABC&authorization_code=UrlEncode(authorizationCode)&redirect_uri=UrlEncode(https://yourwebsite/sso/oauth/callback)  ```      - The API will validate the authorization code, and you will receive a token for this API for the user corresponding to the SSO user.      #### SSO Access token      This is the another method to use an external OAUTH SSO provider configured in Dispatch. You can use this method if your application has its own external SSO provider information (client id and client secret).    - From your application, authenticate to the external SSO provider using authorization code flow :       - Build the authorization url with your SSO credentials       - After user connection, you receive an authorization code        - Call the SSO provider to get an access token    - Call token endpoint in the API with grant_type access_token_TheSsoConfigurationCode and parameters access_token. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=access_token_ABC&access_token=UrlEncode(ssoAccessToken)  ```      - The API will validate the sso access token, and you will receive a token for this API for the user corresponding to the SSO user.      #### SAML      This is the standard method to use an external SSO provider configured in Dispatch and using SAML V2 protocol.    - First, your application will need to open a browser to send the user to the external SSO SAML provider.      To get the authorization url, use the operation _sso/saml-authorization-url_ and provide the parameters : ssoCode and relayState, and provide bearer token created with a client credentials authentication.      - When the user connects, the SAML provider will redirect back to your application with additional information in query parameters on in a form        - if information is received with a POST, read the request form and build a query string using form data        - if information is received with a GET (redirect method), get the query string      - Call token endpoint in the API with grant_type saml_TheSsoConfigurationCode and parameters saml_http_method and saml_content. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=saml_ABC&saml_http_method=GET&saml_content=UrlEncode(SAMLResponse=value1&RelayState=value2)  ```      - The API will validate the SAML content, and you will receive a token for this API for the user corresponding to the SSO user.      # General instructions      ## Projection      To select the fields returned by the api, provide 'fields' parameter in query string: list of fields separated with a comma.     When fields are not provided, calls will only return the first level fields, unless specified otherwise in a route description.      Example with simple fields : `GET route_to_my_resource/myResourceUid?fields=field1,field2,field3` returns :    ```json  {      \"field1\": \"My field 1\",      \"field2\": \"My field 2\",      \"field3\": \"My field 3\"  }  ```      Example with nested properties and collections : `GET route_to_my_resource/myResourceUid?fields=field1,nested1/nested2/nestedProperty3,myCollection1/field2` returns :    ```json  {      \"field1\": \"My field 1\",      \"nested1\": {          \"nested2\": {              \"nestedProperty3\": \"my nested property value 3\"          }      }      \"myCollection1\": [          {              \"field2\": \"my field 2\"          }      ]  }  ```      You can group properties on a nested object using parenthesis. Example to get the properties nested1.field1 and nested1.field2 : `fields=nested1(field1,field2)`.      Example to get the properties at the first level only, without nested properties or collections : `route_to_my_resource?fields=~`.      Example to get all the properties, included nested properties and collections : `route_to_my_resource?fields=*`.      ## Filtering      Access to paginated resources accept filtering with query parameters. Example : `route_to_my_paginated_resources?myFilter1=myValue1&myFilter2=myValue2`.      When a query parameter is a collection, the values must be delimited with a comma. To escape a comma in a value, double it.    Example to pass the words : \"hello,world\", \"value1\", \"value2\" to \"names\" parameter : `route_to_my_paginated_resources?names=hello,,world,value1,value2`.      When the description of a filter indicates that it accepts a pattern, you can search with _startsWith_, pattern.    In some cases, when it is specified in the description, the filter also supports _endsWith_ or _contains_ patterns.      Examples :   - Exact match : `myFilter=abc`    - Starts with abc : `myFilter=abc*`    - Ends with abc : `myFilter=*abc`    - Contains abc : `myFilter=*abc*`       For some operations, you can search on multiple fields, using pattern and pattern fields. The fields can be different depending on the operation, check the description of _patternFields_ parameter to get the available fields.    Example, to search resources having a code or a label starting with abc : `pattern=abc*&patternFields=code,label`.      ## Pagination      For paginated resources, provide query option's _startIndex_, _count_, _fields_, _sort_ for ascending ordering and _desc_ for descending ordering.    Example : `route_to_my_paginated_resources?startIndex=0&count=10&fields=field1,field2&sort=field2`    Example of response :    ```json  {      \"data\": [          {              \"field1\": \"My field 1 value 1\",              \"field2\": \"My field 2 value 1\"          },          {              \"field1\": \"My field 1 value 2\",              \"field2\": \"My field 2 value 2\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```  ## Errors      In case of errors, an object Problem Details is returned with its http status code.    It corresponds to the standard problem details, with an additional property _data_.    See more details about RFC 7807 problem details [here](https://tools.ietf.org/html/rfc7807).    Example :    ```json  {      \"type\": \"/problems/types/service-unavailable\",      \"title\": \"Service unavailable.\",      \"status\": 503,      \"instance\": \"/problems/instances/dispatch-module-not-available\"  }  ```    When a business error occurs, an error code and error message are included in an additional property.    Example :    ```json  {      \"data\": {          \"errors\": [              {                  \"code\": \"resource_not_found\",                  \"message\": \"Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",              }          ]      },      \"type\": \"/problems/types/not-found\",      \"title\": \"A business error occurred\",      \"status\": 404,      \"detail\": \"resource_not_found : Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",      \"instance\": \"/problems/instances/business-error\"  }  ```    ### Error codes    #### General    Error code | Error description    ----|----    missing_mandatory_parameter | A mandatory parameter is missing  must_be_strictly_greater_than_zero | The parameter must be greater than zero  resource_not_found | The resource does not exist  resource_already_exists | A resource with the same identifier already exists  user_missing_permission | The user is not allowed to perform the operation or to access the resource  datetime_out_of_range | The date time is out of the range supported by the system  guid_should_not_be_empty | The unique identifier GUID must not be empty  invalid_enum_value | The provided value is not a valid enumeration value   invalid_field_name | The field provided in the query parameters is not valid  field_not_supported | The field in the 'fields' or 'sort' parameters is not supported  user_entity_access_forbidden | The user is not allowed to access the resource  value_out_of_range | The provided value is out of the expected range  max_length_exceeded | The maximum length of the parameter is exceeded  end_date_must_be_greater_than_start_date | The end date must be greater than the start date  max_date_interval_reached | The maximum date interval is reached  user_has_no_eureka_maps_key | The user does not have access to Eureka Maps  max_user_application_preferences_reached | The maximum number of application preferences is reached for this user    #### Drivers    Error code | Error description    ----|----    driver_unavailability_reason_in_use | The driver unavailability reason is used  driver_unavailability_reason_disabled | The driver unavailability reason is disabled  driver_unavailability_overlapping | There is a driver unavailability in the same interval  contractor_agent_subcontractor_mandatory | The subcontractor is mandatory when the contractor agent is provided  contractor_agent_invalid_subcontractor | The provided contractor agent cannot be used with the provided subcontractor    #### Missions    Error code | Error description    ----|----    transport_assignment_constraints_violation | Transport assignment constraints not fulfilled  transport_unassignment_constraints_violation | Transport unassignment constraints not fulfilled  transport_fragmentation_constraints_violation | Transport fragmentation constraints not fulfilled  transport_defragmentation_constraints_violation | Transport defragmentation constraints not fulfilled  invalid_file_category | The provided file category is invalid  invalid_custom_parameter | The provided custom parameter is invalid  invalid_packing_nature | The package nature is invalid  invalid_unit_code | The unit code is invalid  invalid_dimension_unit | The dimension unit is invalid  transport_unexpected_state | The operation cannot be performed because the transport has an unexpected status  document_report_unexpected_state | The operation cannot be performed because the transport document report has an unexpected status  invalid_service_code | The provided service code is invalid  service_disabled | The provided service is disabled  invalid_customer_code | The provided customer code is invalid  customer_not_usable | The provided customer cannot be used  invalid_pricing_path | The provided pricing path is invalid  itinerary_computation_failed | Unable to compute itinerary for the transport    #### Addresses    Error code | Error description    ----|----    invalid_address_id | The provided address id is invalid  invalid_city_id | The provided city id is invalid  geocoding_failed | Unable to geocode the address        ## Resource update      You can use PATCH operations to partially update a resource.    The update uses merge patch semantics : you need to provide only the fields to updates.    See more details about JSON merge patch standard [here](https://tools.ietf.org/html/rfc7396).      Fields that are not present in the request will be preserved, and fields set to null will be cleared.      There is a special process for lists with nested objects :    - The objects that are not provided in a list are not updated.    - For all objects provided in a list : specify the merge action (add, update, remove) and the entity key when the action is update or remove.      Example, with comments indicating the performed operations :    ```json  {      \"property1\": \"new value 1\",         // property1 is updated       \"nested1\": {                        // nested1 is added if it doesn't exist, otherwise, it is updated          \"property2\": \"my value 2\",      // property2 is updated          \"nested2\": null,                // nested2 is deleted          \"nested3\": {                    // nested3 is added if it doesn't exist              \"property3\": \"my value 3\"   // property3 is updated          }      },      \"nestedCollection1\": [              // nestedCollection1 is updated          {              \"myIdentifier\": \"id1\",              \"mergeAction\": \"update\",              \"property4\": \"my value 1\"   // the property property4 of the item with myIdentifier equals to id1 is updated          },          {              \"myIdentifier\": \"id2\",              \"mergeAction\": \"remove\"     //  the item with myIdentifier equal to id2 is removed           },          {              \"mergeAction\": \"add\",       // an item is added to the collection with the provided property4.              \"property4\": \"my value 2\"          }      ]  }  ```      ## Date formats    Many dates in the system are saved without timezone information. These dates depend on the context (for example : pickup date correponds generally to the timezone of pickup address, but the offset is not kept in the system). In these cases, the timezone is not taken into account when provided in the input. For example, if you provide `2022-01-31T10:00:00+05:00`, the system will keep only `2022-01-31T10:00:00`, the date will be returned with an unspecified timezone.      The api returns the timezone when it is available in the system.      Examples of dates returned by the api :  * Date with unspecified time zone : `2022-01-31T10:00:00`  * UTC date : `2022-01-31T10:00:00+00:00` or `2022-01-31T10:00:00Z`  * Date with offset +5h : `2022-01-31T10:00:00+05:00`    ## Versioning    This API is versioned to handle changes in API contract.    The versioning is done with URI path. Example : _https://myapi/v1/myresource_.    The versioning applies on all the operations except _/token_ and _/revoke_ endpoints. Example: _https://myapi/token_.      A new version is released when a breaking change is made.    The previous versions remain maintained for a certain time. After a time, old versions may be deprecated (marked as obsolete) then removed.    Remark: breaking changes can be made on pre-release versions. These versions can also be removed from the API and replaced with a release version.      Examples of breaking changes :    - Change a property name or type in a request or response    - Add a required parameter in a request    - Remove a property in a response    - Remove an operation in the API      Examples of non-breaking changes :  - Add a value in a response  - Add an optional parameter in a request    - Add an operation in the API    - Change the format of any opaque identifier  - Change the format of a date time : add or remove timezone information, for example : change from `2022-01-31T10:00:00` to `2022-01-31T10:00:00Z` or to `2022-01-31T10:00:00+05:00`.    - Replacing a success status code with another success status code (i.e: replacing a 200 with a 201)  - Replacing a failure status code with another failure status code (i.e: replacing a 500 with a 400)    When a feature is added (without breaking changes), it is available for the last released API version.    In this case, a flag can be added to indicate that this feature is available. See api flags endpoint.      ## Rate Limiting    This API uses consumption quotas to prevent excessive calls and to maintain service availability.    Two types of rate limiting are applied :  - leaky buckets    - concurrency      You can view these limits in Eureka integrator's portal.    There are also additional limits for calls performed by one bearer token, configured by the carrier. These limits are not visible in the portal.      When a quota is exceeded, the API returns the error code 429 (Too Many Requests). A _Retry-After_ header is included and indicates (in seconds) how long to wait before making a new request.      ### Leaky bucket rate limiting      A leaky bucket rate limiting is applied to each API client, it depends on two parameters : the bucket size and the leak rate.      Each API client have its own bucket of a limited size.    When a request is received :    - If the bucket is not full, the bucket counter is increased.    - If the bucket is full (the bucket counter reaches the bucket size), the request is rejected. In this case, the API client must wait before sending a new request.      The bucket counter is decreased periodically to accept new requests.      You can find more details about leaky bucket algorithm [here](https://en.wikipedia.org/wiki/Leaky_bucket).      ### Concurrency rate limiting      The API accepts only a limited number of concurrent calls for each API client.    If the maximum number of concurrent calls is reached, the request is rejected. In this case, the API client must wait before sending a new request.        # Dispatch Api    ## Available features    This API includes various features :  * Retrieval of multiple resources needed for mission entry  * Mission and quotation operations : creation, update, cancellation  * Operations on transport documents (download, upload, delete)  * Operations on transport : sub state application, assignment, pickup, delivery, termination, cancellation  * Retrieval of transport data  * Transport document reports generation  * Fragmentation    This API does not include these features :  * In mission entry : dangerous goods, air transport, deposits (aka consignes), orders in sequence  * Disputes and misfunction sheets, call for tender  * D2D  * Operations (creation, update, delete) on resources and tools except on transports  * Retrieval of resources not needed for mission entry  * Billing  * Regular orders  * Rounds  * Mission scheduling  * Exports and imports  * Reports not related to one transport    ## Missions, quotations and transports      A transport is a system for conveying something (goods or people) from a place (pickup) to another (delivery). A transport is related to one customer and to one service and may contain packages.      A mission is an order containing one or many transports.    In general, a mission will contain only one transport. But in case of a mission with return, it will contain a second transport.    This API supports missions with one or two transports.      A quotation can be requested to indicate how much an order will cost.    When a quotation is validated, it is transformed into a new mission : a new mission is created with quotation information, and the quotation is archived.      Use :  * GET /transports to retrieve a collection of transports (use filtering and projection)    * POST /missions to create a mission  * POST /quotations to create a quotation        ## Mission entry    The following steps can be used to create a mission or a quotation.    ### Authenticate as a user    Connect as a user, with a login and password    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=password&username=my_dispatch_user_login&password=my_dispatch_user_password  ```    The response is a json containing the access token    ```json  {      \"access_token\": \"eyJh...\",      \"token_type\": \"bearer\",      \"expires_in\": 3599,      \"refresh_token\": \"11c...\",      \"as:clientId\": \"...\",      \".issued\": \"Mon, 08 Jul 2024 14:43:45 GMT\",      \".expires\": \"Mon, 08 Jul 2024 15:43:45 GMT\"  }  ```    For all the following calls, add a header `Authorization: Bearer my_access_token`.      ### Select a customer      Get the customers available for the connected user, use filtering to get only the fields you need, use pagination.      ```  GET /customers?startIndex=0&count=10&fields=uid,code,label  ```    Response    ```json  {      \"data\": [          {              \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",              \"code\": \"CUST01\",              \"label\": \"Sample customer 1\"          },          {              \"uid\": \"9acc8b97729c447e9b52bcf87d392c35\",              \"code\": \"CUST02\",              \"label\": \"Sample customer 2\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```  Get customer's information by its unique identifier or by its code. Examples :  ```  GET /customers/1b1548e7d1f046b0ba357539b139be7f?fields==uid,code,label  GET /customers/by-customer-code?customerCode=CUST01&fields==uid,code,label  ```    Response    ```json  {      \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",      \"code\": \"CUST01\",      \"label\": \"Sample customer 1\"  }  ```    ### Addresses    In address step, validate the selected address by getting the cities and countries available in the system, a valid city can be needed for the pricing. Example :      ```  GET /cities?postCode=34000&cityName=montpellier&countryCode=FR&fields=cityId,name,postCode,country/label  ```    Response    ```json  {      \"data\": [          {              \"cityId\": 17037,              \"name\": \"MONTPELLIER\",              \"postCode\": \"34000\",              \"country\": {                  \"label\": \"FRANCE\"              }          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 1,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 1      }  }  ```    ### Packing    In package entry, get package natures.      ```  GET /package-natures?fields=packageNatureId,code,label,packageFamilyLabel  ```    Response    ```json  {      \"data\": [          {              \"packageNatureId\": 1,              \"code\": \"CARTONA3\",              \"label\": \"Carton A3\",              \"packageFamilyLabel\": \"Colis\"          },          {              \"packageNatureId\": 2,              \"code\": \"Boite\",              \"label\": \"Boite\",              \"packageFamilyLabel\": \"Consommables\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```      ### Services    In service entry, get services.    ```  GET /services?startIndex=0&count=10&fields=uid,code,label,serviceFamily/label  ```    Response    ```json  {      \"data\": [          {              \"uid\": \"4da0ba3ef10c4147b73debd01d375000\",              \"code\": \"T1\",              \"label\": \"BREAK 500 KG\",              \"serviceFamily\": {                  \"label\": \"Poids Lourds\"              }          },          {              \"uid\": \"1cfb5b59d877406aa7b86908c9c12f63\",              \"code\": \"T1-MO\",              \"label\": \"MOTO 50KG\",              \"serviceFamily\": {                  \"label\": \"Prestations intra muros\"              }          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```    TO see the sub services included in a service, get its sub services.    ```  GET /services/4da0ba3ef10c4147b73debd01d375000/sub-services?fields=subServiceCode,subServiceLabel,defaultQuantity  ```    Response    ```json  {      \"data\": [          {              \"subServiceCode\": \"ADV\",              \"subServiceLabel\": \"Assurance ad valorem\",              \"defaultQuantity\": 1.00          },          {              \"subServiceCode\": \"KM\",              \"subServiceLabel\": \"Kilomètres\",              \"defaultQuantity\": 0.00          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```    For additional sub services, use sub-services route with available filtering and paging. Example:     ```  GET /sub-services?isEnabled=true&code=KM*&patternFields=code,label  ```    Response    ```json  {      \"data\": [          {              \"code\": \"KM\",              \"label\": \"Kilomètres\",              \"isEnabled\": true,              \"unitCode\": \"KM\",              \"comment\": null,              \"vatCode\": \"1\",              \"quantityMustBeSelected\": false,              \"subServiceQuantities\": null          },          {              \"code\": \"KM-NS\",              \"label\": \"Kilomètres sans surcharge carburant\",              \"isEnabled\": true,              \"unitCode\": \"KM\",              \"comment\": null,              \"vatCode\": \"2\",              \"quantityMustBeSelected\": false,              \"subServiceQuantities\": null          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```      ### Mission creation    Before creating a mission, you can perform a dry run by using `POST /missions/dry-run` route. It simulates a mission creation without creating it. Example:      ```json  {      \"transports\": [          {              \"customerCode\": \"ClientDemo1\",              \"agencyCode\": \"01\",              \"pickupStep\": {                  \"stepAddress\": {                      \"addressId\": 314312                  },                  \"stepDate\": {                      \"plannedStart\": {                          \"dateTime\": \"2024-07-02T00:00:00\",                          \"isTimeOfDayIgnored\": true                      }                  }              },              \"deliveryStep\": {                  \"stepAddress\": {                      \"AddressId\": 339923                  },                  \"stepDate\": {                      \"plannedStart\": {                          \"dateTime\": \"2024-07-03T00:00:00\",                          \"isTimeOfDayIgnored\": true                      }                  }              },              \"serviceCode\": \"T8\",              \"subServices\": [                  {                      \"code\": \"KM\"                  },                  {                      \"code\": \"ADV\",                      \"forcedSellQuantity\": 11                  }              ],              \"serviceCustomParameters\": [                  {                      \"code\": \"PAL\",                      \"Value\": \"2\"                  }              ]          }      ]  }  ```    This route returns information about the mission with its pricing, but without creating it.    To create a mission, use the same request, with `POST /missions` route.        ### Additional remarks on transport creation    * The provided entities (example: customer, service, agency) must be available for the connected user regarding its restrictions.    * The provided customer usable and not a prospect. If the orderer is mandatory for a customer, the orderer must be provided.    * The provided service must be enabled. If it has an agency restriction, it must match customer's agency.    * All provided identifiers (example : city id, agency code, VAT code, unit code) must exist  * Customer and service's custom parameters are checked. Example : the command is rejected if a custom parameter is mandatory, not provided and does not have a default value.    * Transport's references are checked. Examples : if mandatory and not provided, if must be selected from a list but does not match this list.    * Pickup and delivery dates are mandatory  * For package lines, provide either the nature id or a nature (free text). Default values are copied from the nature.    To change this behavior, enable _skipPackageNatureDefaultValuesApplication_ in command's api behavior.      * For transport's sub services: when they are not provided, the transport uses service's default sub services. Otherwise, the transport uses only the specified sub services : service's default sub services are not included and must be specified manually if they are needed.    * Sell pricing path can be forced.   The pricing path must be a customer pricing path (not a subcontractor's one).    When provided, the addresses are overwritten with the ones defined in the provided sell pricing path.      The provided service must match the one defined in the pricing path.    * The mission or quotation is created as unique if at least two transports are provided. To force this parameter, provide _forceUniqueOrder_ in command's api behavior.    * Itinerary computation is done if configured in global settings and if the distance is not provided. To ignore failures during this step, enable _ignoreItineraryComputationFailures_ in command's api behavior.    * Transport steps geocoding is attempted if itinerary computation is required. To try to geocode a step in other cases, enable _tryGeocodeTransportSteps_ in command's api behavior.     * Gas emission is computed if enabled in global settings and not forced.   * Transport bill address is unsed only when the customer is occasional.      ## Transport retrieval    Here is an example to search transports using projection, filtering and pagination.    In this example, we want the 5 top transports having a pickup post code 34000, for customer code CUST01 and with statuses C (assigned) and 0 (recorded), we only want the fields uid, missionNumber, serviceCode and status.    ```  GET /transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=0&count=5  ```  Response  ```json  {      \"data\": [          {              \"uid\": \"fcba3be8e09740499079f86dff98f5e6\",              \"missionNumber\": 1,              \"status\": \"C\",              \"serviceCode\": \"MO\"          },          {              \"uid\": \"10b9c9ba9b5240cba50fb834bf5cb714\",              \"missionNumber\": 7203,              \"status\": null,              \"serviceCode\": \"TNT\"          },          {              \"uid\": \"bfe364f5478a4f2890ea1cba6c535a99\",              \"missionNumber\": 13443,              \"status\": null,              \"serviceCode\": \"T1\"          },          {              \"uid\": \"87253a9702134a198be1748d02eec64d\",              \"missionNumber\": 13444,              \"status\": null,              \"serviceCode\": \"T1\"          },          {              \"uid\": \"5febb2485f474b978730c805ad6e7b79\",              \"missionNumber\": 13445,              \"status\": null,              \"serviceCode\": \"T1\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": \"https://mylicense.dispatchapi.dispatch-rts.com:443/v3/transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=5&count=5\",          \"pageCount\": 2,          \"totalItemCount\": 6,          \"pageNumber\": 1,          \"pageSize\": 5,          \"hasPreviousPage\": false,          \"hasNextPage\": true,          \"isFirstPage\": true,          \"isLastPage\": false,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 5      }  }  ```      ## Transport update    Here is an example to partially update a transport using PATCH operation.      ```  PATCH /transports/dfa98649e1fc4cf8a77c43351b531007  {      \"apiBehavior\": {          \"SkipPackageNatureDefaultValuesApplication\": false,          \"updateTransportOperationComment\": \"My update operation 01\"      },      \"reference2\": \"new-reference-2\",                // The reference2 will be updated      \"pickupStep\": {          \"stepDate\": {              \"plannedStart\": {                  \"dateTime\": \"2022-02-24 10:00:00\",  // the pickup planned start time will be updated                  \"isTimeOfDayIgnored\": false              }          }      },      \"packing\": {          \"lines\": [              {                  \"mergeAction\": \"update\",            // the quantity and references for package id 2268635 will be updated                  \"packageLineId\": 2268635,                  \"quantity\": 3.00,                  \"references\": [                      {                          \"mergeAction\": \"update\",    // the value of the reference with index 3 will be updated                          \"referenceIndex\": 3,                          \"value\": \"abc\"                      }                  ]              },              {                  \"mergeAction\": \"remove\",            // the package with id 2268636 will be removed                  \"packageLineId\": 2268636              },              {                  \"mergeAction\": \"add\",               // this package will be added                  \"nature\": \"Plis\",                  \"quantity\": 2.00,                  \"dimensionUnit\": \"cm\",                  \"height\": 1.00,                  \"width\": 10.00,                  \"length\": 15.00              }          ]      },      \"subServices\": [          {              \"mergeAction\": \"add\",                   // this service will be added              \"code\": \"SP01\",              \"forcedSellQuantity\": 1.2          }      ]  }  ```      ## Transport statuses    To follow up the progress of a transport, check transport status.    When a transport is created, it is recorded, then assigned to a driver, picked up from pickup address, delivered to delivery address, terminated and invoiced.    A transport has one of the following statuses:      Status | Description  ------ | -------  0 (zero), null or empty string | Recorded  A | Canceled  X | Unassigned  b | Pre assigned  B | Pre assigned and accepted by the driver  C | Assigned  E | Picked up  L | Delivered  T | Terminated  s | Dispute undealt  S | Dispute dealt  f | Pre billed  F | Invoiced      Remarks :  - Transport statuses are case sensitive. Make sure you use the correct case when filtering the transports based on their statuses.    - To filter recorded transports only, use the value 0 (zero). The api does not provide filtering with empty values.         ## Transport history statuses    Transport history lines returned by this api have one of the following status codes :      Status | Description  ------ | -------  0 (zero) | Recorded  M | Modified  C | Assigned  E | Picked up  L | Delivered        ## Dispatch desktop app user legacy permissions    A user operator has a set of permissions used in Dispatch desktop application.    The usage of these permissions may change in the future.      Permission Id | Description  ------ | -------  5 | Orders corrections  6 | Write scheduled transports  7 | View scheduled transports  8 | Print scheduled transports  10 | Orders consistency  11 | Statistics  12 | Margins  13 | Global settings  14 | Imports / Exports  15 | See Prices  16 | Invoices calculation  17 | Invoices editing  18 | Ad Hoc invoices  19 | Vouchers book  20 | Credit notes  21 | Invoices clearing  22 | View not issued bills  23 | Accounting export  24 | Payment  25 | Write units management  26 | View units  27 | Print units  28 | Invoices management  29 | Generate bills states  30 | Write customers  31 | View customers  32 | Print customers  33 | Write employees  34 | View employees  35 | Print employees  36 | Write subcontractors  37 | View subcontractors  38 | Print subcontractors  39 | Write salespersons  40 | View salespersons  41 | Print salespersons  42 | Write users  43 | View users  44 | Print users  45 | Write addresses  46 | View addresses  47 | Print addresses  48 | Write sectors  49 | View sectors  50 | Print sectors  51 | Write pricing paths  52 | View pricing paths  53 | Print pricing paths  54 | Write services  55 | View services  56 | Print services  57 | Write sub services  58 | View sub services  59 | Print sub services  60 | Write currencies  61 | View currencies  62 | Print currencies  63 | Write discount scale  64 | View discount scale  65 | Print discount scale  66 | Write payment methods  67 | View payment methods  68 | Print payment methods  69 | Write VAT  70 | View VAT  71 | Print VAT  72 | Write vehicles  73 | View vehicles  74 | Print vehicles  75 | Write cities  76 | View cities  77 | Print cities  78 | (Reserved)  79 | (Reserved)  80 | Companies  81 | (Reserved)  82 | (Reserved)  83 | Write orderer profiles  84 | View orderer profiles  92 | Customizable planning  93 | Tracking alerts  94 | Driver's message alerts  95 | Alerts on mobile dispute  96 | Alerts on unread / rejected missions  97 | Alerts on new web mission  98 | Alert on subcontractor's reminder  99 | Distribution entry  100 | Write references  101 | View references  102 | Print references  103 | Delegated orders tracking  104 | Alerts on mobile anomalies  105 | D2D Services equivalences  106 | D2D Services equivalences  107 | D2D amount change  108 | D2D City equivalence  109 | Customer / Subcontractor configuration  110 | D2D City equivalence  111 | Modify terminated orders  112 | Modify billed orders  113 | D2D Errors  114 | New D2D Mission / Modify D2D Mission  115 | Orders ending  116 | Orders cancellation  118 | Write user profiles  119 | View user profiles  120 | Geo device associations  121 | Cartography acces  122 | Create, Update, Delete  123 | Default model choice  124 | Preview  125 | Write agencies  126 | View agencies  127 | Print agencies  128 | Write orderers  129 | View orderers  130 | Print orderers  131 | Create, update, delete POI  132 | Create, update, delete pictogrammes  133 | Invoices clearing  134 | View tems  135 | Tools  136 | Time management administration  137 | View time management  138 | Subcontractors documents  139 | Write transport statuses  140 | Write transport substates  141 | View transport substates  142 | Write anomalies  143 | View anomalies  144 | Write disputes  145 | View disputes  146 | Write deposits  147 | View deposits  148 | Notepad  149 | Write unavailabilities  150 | View unavailabilities  151 | (Reserved)  152 | Consult a dysfunction sheet  153 | Modify a dysfunction sheet  154 | Print a dysfunction sheet  155 | Close misfunction sheet  156 | ReOpen a dysfunction sheet  157 | Write contractor agents  158 | View contractor agents  159 | Print contractor agents  160 | Kick users via instant messaging  161 | GPS devices configuration window  162 | Instant messaging  163 | Dispatch's alerts  164 | Notes  165 | Baracoda Tracking  166 | Set VAT in order  167 | Transport order manual price setting  168 | Write EDI parameters  169 | Global communications history  170 | Write regions  171 | View regions  172 | Write price modification rules  173 | View price modification rules  174 | Saving tabs configuration  175 | Allow users to see all agencies vehicles on Dispatch map  176 | Allow user to see recurrent orders from all agencies on map  177 | Desactivate free package type input  178 | Allow GPS Device privacy mode change  179 | Allow automatic working time declaration module execusion on current day  180 | Temperature alert  181 | Write package natures  182 | View package natures  183 | Print package natures  184 | Add package history state  185 | Write package reference values  186 | View package reference values  187 | Validate round sessions  188 | Override round session selection  189 | Override round session values  190 | Allow to secure transport  191 | allow batch modification of transport references  192 | Write subcontractor employees  193 | View subcontractor employees  194 | Print subcontractor employees  195 | Allow receipt entry  196 | Write appointments  197 | View appointments  198 | Access to planning appointment module  199 | Access to audit data  200 | Access to audit configuration  201 | Write password security policies  202 | View password security policies  203 | Transports audit  204 | Access to application orderers  205 | Write SSO configurations  206 | View SSO configurations  207 | Write SSO users  208 | View SSO users  209 | Application access management  210 | Display open link in browser button  211 | Data erasure and anonymization  212 | Write EDI equivalence configurations  213 | View EDI equivalence configurations  214 | Access to anonymized data

API version: v3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// TransportsAPIService TransportsAPI service
type TransportsAPIService service

type TransportsAPIApplySubStateRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyTransportSubStateCommand
}

// Sub state application command
func (r TransportsAPIApplySubStateRequest) Command(command ApplyTransportSubStateCommand) TransportsAPIApplySubStateRequest {
	r.command = &command
	return r
}

func (r TransportsAPIApplySubStateRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApplySubStateExecute(r)
}

/*
ApplySubState Apply a sub state to the transport

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIApplySubStateRequest
*/
func (a *TransportsAPIService) ApplySubState(ctx context.Context, uid string) TransportsAPIApplySubStateRequest {
	return TransportsAPIApplySubStateRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) ApplySubStateExecute(r TransportsAPIApplySubStateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.ApplySubState")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/substates"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIApplyTransportCancelledStatusRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyCancelledStatusCommand
}

func (r TransportsAPIApplyTransportCancelledStatusRequest) Command(command ApplyCancelledStatusCommand) TransportsAPIApplyTransportCancelledStatusRequest {
	r.command = &command
	return r
}

func (r TransportsAPIApplyTransportCancelledStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApplyTransportCancelledStatusExecute(r)
}

/*
ApplyTransportCancelledStatus Apply transport cancelled status

The following permission(s) are required to access this route:  Cancel orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIApplyTransportCancelledStatusRequest
*/
func (a *TransportsAPIService) ApplyTransportCancelledStatus(ctx context.Context, uid string) TransportsAPIApplyTransportCancelledStatusRequest {
	return TransportsAPIApplyTransportCancelledStatusRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) ApplyTransportCancelledStatusExecute(r TransportsAPIApplyTransportCancelledStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.ApplyTransportCancelledStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/cancellation-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIApplyTransportDeliveredStatusRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyDeliveredStatusCommand
}

func (r TransportsAPIApplyTransportDeliveredStatusRequest) Command(command ApplyDeliveredStatusCommand) TransportsAPIApplyTransportDeliveredStatusRequest {
	r.command = &command
	return r
}

func (r TransportsAPIApplyTransportDeliveredStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApplyTransportDeliveredStatusExecute(r)
}

/*
ApplyTransportDeliveredStatus Apply transport delivered state

The following permission(s) are required to access this route:  Update orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIApplyTransportDeliveredStatusRequest
*/
func (a *TransportsAPIService) ApplyTransportDeliveredStatus(ctx context.Context, uid string) TransportsAPIApplyTransportDeliveredStatusRequest {
	return TransportsAPIApplyTransportDeliveredStatusRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) ApplyTransportDeliveredStatusExecute(r TransportsAPIApplyTransportDeliveredStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.ApplyTransportDeliveredStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/delivery-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIApplyTransportPickedUpStatusRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyPickedUpStatusCommand
}

func (r TransportsAPIApplyTransportPickedUpStatusRequest) Command(command ApplyPickedUpStatusCommand) TransportsAPIApplyTransportPickedUpStatusRequest {
	r.command = &command
	return r
}

func (r TransportsAPIApplyTransportPickedUpStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApplyTransportPickedUpStatusExecute(r)
}

/*
ApplyTransportPickedUpStatus Apply transport pickup state

The following permission(s) are required to access this route:  Update orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIApplyTransportPickedUpStatusRequest
*/
func (a *TransportsAPIService) ApplyTransportPickedUpStatus(ctx context.Context, uid string) TransportsAPIApplyTransportPickedUpStatusRequest {
	return TransportsAPIApplyTransportPickedUpStatusRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) ApplyTransportPickedUpStatusExecute(r TransportsAPIApplyTransportPickedUpStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.ApplyTransportPickedUpStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/pickup-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIApplyTransportTerminatedStatusRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyTerminatedStatusCommand
}

func (r TransportsAPIApplyTransportTerminatedStatusRequest) Command(command ApplyTerminatedStatusCommand) TransportsAPIApplyTransportTerminatedStatusRequest {
	r.command = &command
	return r
}

func (r TransportsAPIApplyTransportTerminatedStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApplyTransportTerminatedStatusExecute(r)
}

/*
ApplyTransportTerminatedStatus Apply transport terminated state

The following permission(s) are required to access this route:  Terminate orders,Update orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIApplyTransportTerminatedStatusRequest
*/
func (a *TransportsAPIService) ApplyTransportTerminatedStatus(ctx context.Context, uid string) TransportsAPIApplyTransportTerminatedStatusRequest {
	return TransportsAPIApplyTransportTerminatedStatusRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) ApplyTransportTerminatedStatusExecute(r TransportsAPIApplyTransportTerminatedStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.ApplyTransportTerminatedStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/termination-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPICheckTransportAssignmentsConstraintsRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *AssignTransportDto
}

func (r TransportsAPICheckTransportAssignmentsConstraintsRequest) Command(command AssignTransportDto) TransportsAPICheckTransportAssignmentsConstraintsRequest {
	r.command = &command
	return r
}

func (r TransportsAPICheckTransportAssignmentsConstraintsRequest) Execute() (*CheckTransportAssignmentConstraintsResultDto, *http.Response, error) {
	return r.ApiService.CheckTransportAssignmentsConstraintsExecute(r)
}

/*
CheckTransportAssignmentsConstraints Check transport assignment's constraints

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPICheckTransportAssignmentsConstraintsRequest
*/
func (a *TransportsAPIService) CheckTransportAssignmentsConstraints(ctx context.Context, uid string) TransportsAPICheckTransportAssignmentsConstraintsRequest {
	return TransportsAPICheckTransportAssignmentsConstraintsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return CheckTransportAssignmentConstraintsResultDto
func (a *TransportsAPIService) CheckTransportAssignmentsConstraintsExecute(r TransportsAPICheckTransportAssignmentsConstraintsRequest) (*CheckTransportAssignmentConstraintsResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CheckTransportAssignmentConstraintsResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.CheckTransportAssignmentsConstraints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/assignment/constraints-validation"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return localVarReturnValue, nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIDefragmentTransportRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *DefragmentTransportCommand
}

func (r TransportsAPIDefragmentTransportRequest) Command(command DefragmentTransportCommand) TransportsAPIDefragmentTransportRequest {
	r.command = &command
	return r
}

func (r TransportsAPIDefragmentTransportRequest) Execute() (*http.Response, error) {
	return r.ApiService.DefragmentTransportExecute(r)
}

/*
DefragmentTransport Defragment transport

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIDefragmentTransportRequest
*/
func (a *TransportsAPIService) DefragmentTransport(ctx context.Context, uid string) TransportsAPIDefragmentTransportRequest {
	return TransportsAPIDefragmentTransportRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) DefragmentTransportExecute(r TransportsAPIDefragmentTransportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.DefragmentTransport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/defragmentation-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIFragmentTransportRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *FragmentTransportCommand
}

func (r TransportsAPIFragmentTransportRequest) Command(command FragmentTransportCommand) TransportsAPIFragmentTransportRequest {
	r.command = &command
	return r
}

func (r TransportsAPIFragmentTransportRequest) Execute() (*http.Response, error) {
	return r.ApiService.FragmentTransportExecute(r)
}

/*
FragmentTransport Fragment transport

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIFragmentTransportRequest
*/
func (a *TransportsAPIService) FragmentTransport(ctx context.Context, uid string) TransportsAPIFragmentTransportRequest {
	return TransportsAPIFragmentTransportRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) FragmentTransportExecute(r TransportsAPIFragmentTransportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.FragmentTransport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/fragmentation-process"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIGetTransportAirDataRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	fields     *string
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportAirDataRequest) Fields(fields string) TransportsAPIGetTransportAirDataRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportAirDataRequest) Execute() (*TransportAirDataDto, *http.Response, error) {
	return r.ApiService.GetTransportAirDataExecute(r)
}

/*
GetTransportAirData Get transport air data

Get transport air data by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportAirDataRequest
*/
func (a *TransportsAPIService) GetTransportAirData(ctx context.Context, uid string) TransportsAPIGetTransportAirDataRequest {
	return TransportsAPIGetTransportAirDataRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportAirDataDto
func (a *TransportsAPIService) GetTransportAirDataExecute(r TransportsAPIGetTransportAirDataRequest) (*TransportAirDataDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportAirDataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportAirData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/air-data"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportBillAddressRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	fields     *string
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportBillAddressRequest) Fields(fields string) TransportsAPIGetTransportBillAddressRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportBillAddressRequest) Execute() (*TransportBillAddressDto, *http.Response, error) {
	return r.ApiService.GetTransportBillAddressExecute(r)
}

/*
GetTransportBillAddress Get transport bill address

Get transport bill address by its transport unique identifier, when it is different from customer's bill address.
The transport bill address is only available when the customer is occasional.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportBillAddressRequest
*/
func (a *TransportsAPIService) GetTransportBillAddress(ctx context.Context, uid string) TransportsAPIGetTransportBillAddressRequest {
	return TransportsAPIGetTransportBillAddressRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportBillAddressDto
func (a *TransportsAPIService) GetTransportBillAddressExecute(r TransportsAPIGetTransportBillAddressRequest) (*TransportBillAddressDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportBillAddressDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportBillAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/transport-bill-address"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportByUniqueIdentifierRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	fields     *string
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportByUniqueIdentifierRequest) Fields(fields string) TransportsAPIGetTransportByUniqueIdentifierRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportByUniqueIdentifierRequest) Execute() (*TransportDto, *http.Response, error) {
	return r.ApiService.GetTransportByUniqueIdentifierExecute(r)
}

/*
GetTransportByUniqueIdentifier Get a transport by uid

Get a transport by its unique identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportByUniqueIdentifierRequest
*/
func (a *TransportsAPIService) GetTransportByUniqueIdentifier(ctx context.Context, uid string) TransportsAPIGetTransportByUniqueIdentifierRequest {
	return TransportsAPIGetTransportByUniqueIdentifierRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportDto
func (a *TransportsAPIService) GetTransportByUniqueIdentifierExecute(r TransportsAPIGetTransportByUniqueIdentifierRequest) (*TransportDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportByUniqueIdentifier")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportCommunicationConfigurationRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
}

func (r TransportsAPIGetTransportCommunicationConfigurationRequest) Execute() (*TransportCommunicationConfigurationDto, *http.Response, error) {
	return r.ApiService.GetTransportCommunicationConfigurationExecute(r)
}

/*
GetTransportCommunicationConfiguration Get transport communication configuration

Get transport communication configuration by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportCommunicationConfigurationRequest
*/
func (a *TransportsAPIService) GetTransportCommunicationConfiguration(ctx context.Context, uid string) TransportsAPIGetTransportCommunicationConfigurationRequest {
	return TransportsAPIGetTransportCommunicationConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportCommunicationConfigurationDto
func (a *TransportsAPIService) GetTransportCommunicationConfigurationExecute(r TransportsAPIGetTransportCommunicationConfigurationRequest) (*TransportCommunicationConfigurationDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportCommunicationConfigurationDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportCommunicationConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/communication-configuration"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportCustomerCustomParametersRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportCustomerCustomParametersRequest) StartIndex(startIndex int32) TransportsAPIGetTransportCustomerCustomParametersRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportCustomerCustomParametersRequest) Count(count int32) TransportsAPIGetTransportCustomerCustomParametersRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportCustomerCustomParametersRequest) Sort(sort []string) TransportsAPIGetTransportCustomerCustomParametersRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportCustomerCustomParametersRequest) Desc(desc []string) TransportsAPIGetTransportCustomerCustomParametersRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportCustomerCustomParametersRequest) Fields(fields string) TransportsAPIGetTransportCustomerCustomParametersRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportCustomerCustomParametersRequest) Execute() (*IPagedResourceListTransportCustomParameterDto, *http.Response, error) {
	return r.ApiService.GetTransportCustomerCustomParametersExecute(r)
}

/*
GetTransportCustomerCustomParameters Get transport customer custom parameters

Get a paginated list of transport's customer custom parameters by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportCustomerCustomParametersRequest
*/
func (a *TransportsAPIService) GetTransportCustomerCustomParameters(ctx context.Context, uid string) TransportsAPIGetTransportCustomerCustomParametersRequest {
	return TransportsAPIGetTransportCustomerCustomParametersRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListTransportCustomParameterDto
func (a *TransportsAPIService) GetTransportCustomerCustomParametersExecute(r TransportsAPIGetTransportCustomerCustomParametersRequest) (*IPagedResourceListTransportCustomParameterDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListTransportCustomParameterDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportCustomerCustomParameters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/customer-custom-parameters"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportDangerousGoodsRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
}

func (r TransportsAPIGetTransportDangerousGoodsRequest) Execute() (*IResourceListTransportDangerousGoodDto, *http.Response, error) {
	return r.ApiService.GetTransportDangerousGoodsExecute(r)
}

/*
GetTransportDangerousGoods Get transport dangerous goods

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportDangerousGoodsRequest
*/
func (a *TransportsAPIService) GetTransportDangerousGoods(ctx context.Context, uid string) TransportsAPIGetTransportDangerousGoodsRequest {
	return TransportsAPIGetTransportDangerousGoodsRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IResourceListTransportDangerousGoodDto
func (a *TransportsAPIService) GetTransportDangerousGoodsExecute(r TransportsAPIGetTransportDangerousGoodsRequest) (*IResourceListTransportDangerousGoodDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IResourceListTransportDangerousGoodDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportDangerousGoods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/dangerous-goods"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportHistoryLinesRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportHistoryLinesRequest) StartIndex(startIndex int32) TransportsAPIGetTransportHistoryLinesRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportHistoryLinesRequest) Count(count int32) TransportsAPIGetTransportHistoryLinesRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportHistoryLinesRequest) Sort(sort []string) TransportsAPIGetTransportHistoryLinesRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportHistoryLinesRequest) Desc(desc []string) TransportsAPIGetTransportHistoryLinesRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportHistoryLinesRequest) Fields(fields string) TransportsAPIGetTransportHistoryLinesRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportHistoryLinesRequest) Execute() (*IPagedResourceListTransportHistoryLineDto, *http.Response, error) {
	return r.ApiService.GetTransportHistoryLinesExecute(r)
}

/*
GetTransportHistoryLines Get transport history lines

Get a paginated list of transport's history lines by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportHistoryLinesRequest
*/
func (a *TransportsAPIService) GetTransportHistoryLines(ctx context.Context, uid string) TransportsAPIGetTransportHistoryLinesRequest {
	return TransportsAPIGetTransportHistoryLinesRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListTransportHistoryLineDto
func (a *TransportsAPIService) GetTransportHistoryLinesExecute(r TransportsAPIGetTransportHistoryLinesRequest) (*IPagedResourceListTransportHistoryLineDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListTransportHistoryLineDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportHistoryLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportPackageLinesRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportPackageLinesRequest) StartIndex(startIndex int32) TransportsAPIGetTransportPackageLinesRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportPackageLinesRequest) Count(count int32) TransportsAPIGetTransportPackageLinesRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportPackageLinesRequest) Sort(sort []string) TransportsAPIGetTransportPackageLinesRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportPackageLinesRequest) Desc(desc []string) TransportsAPIGetTransportPackageLinesRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportPackageLinesRequest) Fields(fields string) TransportsAPIGetTransportPackageLinesRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportPackageLinesRequest) Execute() (*IPagedResourceListPackageLineDto, *http.Response, error) {
	return r.ApiService.GetTransportPackageLinesExecute(r)
}

/*
GetTransportPackageLines Get transport package lines

Get a paginated list of transport's package lines by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportPackageLinesRequest
*/
func (a *TransportsAPIService) GetTransportPackageLines(ctx context.Context, uid string) TransportsAPIGetTransportPackageLinesRequest {
	return TransportsAPIGetTransportPackageLinesRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListPackageLineDto
func (a *TransportsAPIService) GetTransportPackageLinesExecute(r TransportsAPIGetTransportPackageLinesRequest) (*IPagedResourceListPackageLineDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListPackageLineDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportPackageLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/package-lines"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportPricingRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
}

func (r TransportsAPIGetTransportPricingRequest) Execute() (*TransportPricingDto, *http.Response, error) {
	return r.ApiService.GetTransportPricingExecute(r)
}

/*
GetTransportPricing Get transport pricing

The following permission(s) are required to access this route:  See prices.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportPricingRequest
*/
func (a *TransportsAPIService) GetTransportPricing(ctx context.Context, uid string) TransportsAPIGetTransportPricingRequest {
	return TransportsAPIGetTransportPricingRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportPricingDto
func (a *TransportsAPIService) GetTransportPricingExecute(r TransportsAPIGetTransportPricingRequest) (*TransportPricingDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportPricingDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportPricing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/pricing"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportRouteRequest struct {
	ctx            context.Context
	ApiService     *TransportsAPIService
	uid            string
	startDate      *Time
	endDate        *Time
	useSmoothRoute *bool
}

// Optional start date, restrict to positions acquired after this date
func (r TransportsAPIGetTransportRouteRequest) StartDate(startDate Time) TransportsAPIGetTransportRouteRequest {
	r.startDate = &startDate
	return r
}

// Optional end date, restrict to positions acquired before this date
func (r TransportsAPIGetTransportRouteRequest) EndDate(endDate Time) TransportsAPIGetTransportRouteRequest {
	r.endDate = &endDate
	return r
}

// Optional value to specify whether or not the route should be smoothed
func (r TransportsAPIGetTransportRouteRequest) UseSmoothRoute(useSmoothRoute bool) TransportsAPIGetTransportRouteRequest {
	r.useSmoothRoute = &useSmoothRoute
	return r
}

func (r TransportsAPIGetTransportRouteRequest) Execute() (*TransportRouteDto, *http.Response, error) {
	return r.ApiService.GetTransportRouteExecute(r)
}

/*
GetTransportRoute Get transport route

Get transport's positions between two dates.
The positions are retrieved on a regular interval by geolocation module.
When there is no position, the application does not try to download positions from geolocation server.
No position is returned if the transport is over or if there is no driver associated to it.
The positions that occurred before the beginning of the transport are not retrieved.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportRouteRequest
*/
func (a *TransportsAPIService) GetTransportRoute(ctx context.Context, uid string) TransportsAPIGetTransportRouteRequest {
	return TransportsAPIGetTransportRouteRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportRouteDto
func (a *TransportsAPIService) GetTransportRouteExecute(r TransportsAPIGetTransportRouteRequest) (*TransportRouteDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportRouteDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/transport-route"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "", "")
	}
	if r.useSmoothRoute != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "useSmoothRoute", r.useSmoothRoute, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportServiceCustomParametersRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportServiceCustomParametersRequest) StartIndex(startIndex int32) TransportsAPIGetTransportServiceCustomParametersRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportServiceCustomParametersRequest) Count(count int32) TransportsAPIGetTransportServiceCustomParametersRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportServiceCustomParametersRequest) Sort(sort []string) TransportsAPIGetTransportServiceCustomParametersRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportServiceCustomParametersRequest) Desc(desc []string) TransportsAPIGetTransportServiceCustomParametersRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportServiceCustomParametersRequest) Fields(fields string) TransportsAPIGetTransportServiceCustomParametersRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportServiceCustomParametersRequest) Execute() (*IPagedResourceListTransportCustomParameterDto, *http.Response, error) {
	return r.ApiService.GetTransportServiceCustomParametersExecute(r)
}

/*
GetTransportServiceCustomParameters Get transport service custom parameters

Get a paginated list of transport's service custom parameters by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportServiceCustomParametersRequest
*/
func (a *TransportsAPIService) GetTransportServiceCustomParameters(ctx context.Context, uid string) TransportsAPIGetTransportServiceCustomParametersRequest {
	return TransportsAPIGetTransportServiceCustomParametersRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListTransportCustomParameterDto
func (a *TransportsAPIService) GetTransportServiceCustomParametersExecute(r TransportsAPIGetTransportServiceCustomParametersRequest) (*IPagedResourceListTransportCustomParameterDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListTransportCustomParameterDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportServiceCustomParameters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/service-custom-parameters"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportSubServicesRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportSubServicesRequest) StartIndex(startIndex int32) TransportsAPIGetTransportSubServicesRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportSubServicesRequest) Count(count int32) TransportsAPIGetTransportSubServicesRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportSubServicesRequest) Sort(sort []string) TransportsAPIGetTransportSubServicesRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportSubServicesRequest) Desc(desc []string) TransportsAPIGetTransportSubServicesRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportSubServicesRequest) Fields(fields string) TransportsAPIGetTransportSubServicesRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportSubServicesRequest) Execute() (*IPagedResourceListTransportSubServiceDto, *http.Response, error) {
	return r.ApiService.GetTransportSubServicesExecute(r)
}

/*
GetTransportSubServices Get transport sub services

Get a paginated list of transport's sub services by transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportSubServicesRequest
*/
func (a *TransportsAPIService) GetTransportSubServices(ctx context.Context, uid string) TransportsAPIGetTransportSubServicesRequest {
	return TransportsAPIGetTransportSubServicesRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListTransportSubServiceDto
func (a *TransportsAPIService) GetTransportSubServicesExecute(r TransportsAPIGetTransportSubServicesRequest) (*IPagedResourceListTransportSubServiceDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListTransportSubServiceDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportSubServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/subservices"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportTotalBulkSizesRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	startIndex *int32
	count      *int32
	sort       *[]string
	desc       *[]string
	fields     *string
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportTotalBulkSizesRequest) StartIndex(startIndex int32) TransportsAPIGetTransportTotalBulkSizesRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportTotalBulkSizesRequest) Count(count int32) TransportsAPIGetTransportTotalBulkSizesRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportTotalBulkSizesRequest) Sort(sort []string) TransportsAPIGetTransportTotalBulkSizesRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportTotalBulkSizesRequest) Desc(desc []string) TransportsAPIGetTransportTotalBulkSizesRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportTotalBulkSizesRequest) Fields(fields string) TransportsAPIGetTransportTotalBulkSizesRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportTotalBulkSizesRequest) Execute() (*IPagedResourceListPackagesTotalBulkSizeDto, *http.Response, error) {
	return r.ApiService.GetTransportTotalBulkSizesExecute(r)
}

/*
GetTransportTotalBulkSizes Get transport total bulk sizes

Get a paginated list of transport's packing total bulk sizes by its transport unique identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIGetTransportTotalBulkSizesRequest
*/
func (a *TransportsAPIService) GetTransportTotalBulkSizes(ctx context.Context, uid string) TransportsAPIGetTransportTotalBulkSizesRequest {
	return TransportsAPIGetTransportTotalBulkSizesRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListPackagesTotalBulkSizeDto
func (a *TransportsAPIService) GetTransportTotalBulkSizesExecute(r TransportsAPIGetTransportTotalBulkSizesRequest) (*IPagedResourceListPackagesTotalBulkSizeDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListPackagesTotalBulkSizeDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransportTotalBulkSizes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/total-bulk-sizes"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIGetTransportsRequest struct {
	ctx                               context.Context
	ApiService                        *TransportsAPIService
	transportIds                      *[]int32
	transportUids                     *[]string
	missionNumbers                    *[]int32
	missionUids                       *[]string
	missionTrackIds                   *[]string
	driverIds                         *[]int32
	vehicleCodes                      *[]string
	trailerCodes                      *[]string
	vehicleOrTrailerCodes             *[]string
	quotationNumbers                  *[]int32
	quotationUids                     *[]string
	quotationTrackIds                 *[]string
	sourceQuotationNumbers            *[]int32
	dateFrom                          *Time
	dateTo                            *Time
	dateActions                       *[]string
	dateToDateActions                 *[]string
	dateCombinationCriterion          *string
	signs                             *[]string
	references1                       *[]string
	references2                       *[]string
	references3                       *[]string
	pickupCityIds                     *[]int32
	pickupPostCodes                   *[]string
	pickupCityNames                   *[]string
	pickupAddressNames                *[]string
	pickupCountryCodes                *[]string
	pickupSectorCodes                 *[]string
	deliveryCityIds                   *[]int32
	deliveryPostCodes                 *[]string
	deliveryCityNames                 *[]string
	deliveryAddressNames              *[]string
	deliveryCountryCodes              *[]string
	deliverySectorCodes               *[]string
	customerCodes                     *[]string
	ordererCodes                      *[]string
	agencyCodes                       *[]string
	packageBarCodes                   *[]string
	statuses                          *[]string
	fileCategoryCodes                 *[]string
	serviceCodes                      *[]string
	packageReferences1                *[]string
	packageReferences2                *[]string
	packageReferences3                *[]string
	packageReferences4                *[]string
	packageReferences5                *[]string
	packageReferences6                *[]string
	packageReferences7                *[]string
	packageReferences8                *[]string
	pattern                           *string
	patternFields                     *[]string
	missionTypeVisibility             *string
	quotationStatusVisibility         *string
	quotationApprovalStatusVisibility *string
	flagLastSubStateByFamilyCodes     *[]string
	fragmentationVisibility           *string
	startIndex                        *int32
	count                             *int32
	sort                              *[]string
	desc                              *[]string
	fields                            *string
}

// Transports identifiers
func (r TransportsAPIGetTransportsRequest) TransportIds(transportIds []int32) TransportsAPIGetTransportsRequest {
	r.transportIds = &transportIds
	return r
}

// Transports unique identifiers
func (r TransportsAPIGetTransportsRequest) TransportUids(transportUids []string) TransportsAPIGetTransportsRequest {
	r.transportUids = &transportUids
	return r
}

// Missions numbers
func (r TransportsAPIGetTransportsRequest) MissionNumbers(missionNumbers []int32) TransportsAPIGetTransportsRequest {
	r.missionNumbers = &missionNumbers
	return r
}

// Missions unique identifiers
func (r TransportsAPIGetTransportsRequest) MissionUids(missionUids []string) TransportsAPIGetTransportsRequest {
	r.missionUids = &missionUids
	return r
}

// Missions track identifiers
func (r TransportsAPIGetTransportsRequest) MissionTrackIds(missionTrackIds []string) TransportsAPIGetTransportsRequest {
	r.missionTrackIds = &missionTrackIds
	return r
}

// Drivers identifiers
func (r TransportsAPIGetTransportsRequest) DriverIds(driverIds []int32) TransportsAPIGetTransportsRequest {
	r.driverIds = &driverIds
	return r
}

// Vehicles codes
func (r TransportsAPIGetTransportsRequest) VehicleCodes(vehicleCodes []string) TransportsAPIGetTransportsRequest {
	r.vehicleCodes = &vehicleCodes
	return r
}

// Trailer codes
func (r TransportsAPIGetTransportsRequest) TrailerCodes(trailerCodes []string) TransportsAPIGetTransportsRequest {
	r.trailerCodes = &trailerCodes
	return r
}

// Codes for vehicles or trailers
func (r TransportsAPIGetTransportsRequest) VehicleOrTrailerCodes(vehicleOrTrailerCodes []string) TransportsAPIGetTransportsRequest {
	r.vehicleOrTrailerCodes = &vehicleOrTrailerCodes
	return r
}

// Quotations numbers
func (r TransportsAPIGetTransportsRequest) QuotationNumbers(quotationNumbers []int32) TransportsAPIGetTransportsRequest {
	r.quotationNumbers = &quotationNumbers
	return r
}

// Quotations unique identifiers
func (r TransportsAPIGetTransportsRequest) QuotationUids(quotationUids []string) TransportsAPIGetTransportsRequest {
	r.quotationUids = &quotationUids
	return r
}

// Quotations track identifiers
func (r TransportsAPIGetTransportsRequest) QuotationTrackIds(quotationTrackIds []string) TransportsAPIGetTransportsRequest {
	r.quotationTrackIds = &quotationTrackIds
	return r
}

// Source quotation numbers, for transports included in a mission created from a quotation.
func (r TransportsAPIGetTransportsRequest) SourceQuotationNumbers(sourceQuotationNumbers []int32) TransportsAPIGetTransportsRequest {
	r.sourceQuotationNumbers = &sourceQuotationNumbers
	return r
}

// The inclusive start date and time for the search.
func (r TransportsAPIGetTransportsRequest) DateFrom(dateFrom Time) TransportsAPIGetTransportsRequest {
	r.dateFrom = &dateFrom
	return r
}

// The exclusive end date and time for the search.
func (r TransportsAPIGetTransportsRequest) DateTo(dateTo Time) TransportsAPIGetTransportsRequest {
	r.dateTo = &dateTo
	return r
}

// Specify on which transport dates the date search will be performed.  If left empty, the search will be performed on any of the TransportDate.
func (r TransportsAPIGetTransportsRequest) DateActions(dateActions []string) TransportsAPIGetTransportsRequest {
	r.dateActions = &dateActions
	return r
}

// Specify a range of transport dates to perform the research.  The expected schema is : From DateToDateActions dates to DateActions dates.  If left empty, the search will be performed using DateActions
func (r TransportsAPIGetTransportsRequest) DateToDateActions(dateToDateActions []string) TransportsAPIGetTransportsRequest {
	r.dateToDateActions = &dateToDateActions
	return r
}

// The combination criteria for date search.  Have no impact if DateActions is unspecified or contains a single element.    Default value is &#39;Or&#39;.
func (r TransportsAPIGetTransportsRequest) DateCombinationCriterion(dateCombinationCriterion string) TransportsAPIGetTransportsRequest {
	r.dateCombinationCriterion = &dateCombinationCriterion
	return r
}

func (r TransportsAPIGetTransportsRequest) Signs(signs []string) TransportsAPIGetTransportsRequest {
	r.signs = &signs
	return r
}

func (r TransportsAPIGetTransportsRequest) References1(references1 []string) TransportsAPIGetTransportsRequest {
	r.references1 = &references1
	return r
}

func (r TransportsAPIGetTransportsRequest) References2(references2 []string) TransportsAPIGetTransportsRequest {
	r.references2 = &references2
	return r
}

func (r TransportsAPIGetTransportsRequest) References3(references3 []string) TransportsAPIGetTransportsRequest {
	r.references3 = &references3
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupCityIds(pickupCityIds []int32) TransportsAPIGetTransportsRequest {
	r.pickupCityIds = &pickupCityIds
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupPostCodes(pickupPostCodes []string) TransportsAPIGetTransportsRequest {
	r.pickupPostCodes = &pickupPostCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupCityNames(pickupCityNames []string) TransportsAPIGetTransportsRequest {
	r.pickupCityNames = &pickupCityNames
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupAddressNames(pickupAddressNames []string) TransportsAPIGetTransportsRequest {
	r.pickupAddressNames = &pickupAddressNames
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupCountryCodes(pickupCountryCodes []string) TransportsAPIGetTransportsRequest {
	r.pickupCountryCodes = &pickupCountryCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) PickupSectorCodes(pickupSectorCodes []string) TransportsAPIGetTransportsRequest {
	r.pickupSectorCodes = &pickupSectorCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliveryCityIds(deliveryCityIds []int32) TransportsAPIGetTransportsRequest {
	r.deliveryCityIds = &deliveryCityIds
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliveryPostCodes(deliveryPostCodes []string) TransportsAPIGetTransportsRequest {
	r.deliveryPostCodes = &deliveryPostCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliveryCityNames(deliveryCityNames []string) TransportsAPIGetTransportsRequest {
	r.deliveryCityNames = &deliveryCityNames
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliveryAddressNames(deliveryAddressNames []string) TransportsAPIGetTransportsRequest {
	r.deliveryAddressNames = &deliveryAddressNames
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliveryCountryCodes(deliveryCountryCodes []string) TransportsAPIGetTransportsRequest {
	r.deliveryCountryCodes = &deliveryCountryCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) DeliverySectorCodes(deliverySectorCodes []string) TransportsAPIGetTransportsRequest {
	r.deliverySectorCodes = &deliverySectorCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) CustomerCodes(customerCodes []string) TransportsAPIGetTransportsRequest {
	r.customerCodes = &customerCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) OrdererCodes(ordererCodes []string) TransportsAPIGetTransportsRequest {
	r.ordererCodes = &ordererCodes
	return r
}

// A list of exact agency codes to match against the transports agency&#39;s code.
func (r TransportsAPIGetTransportsRequest) AgencyCodes(agencyCodes []string) TransportsAPIGetTransportsRequest {
	r.agencyCodes = &agencyCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageBarCodes(packageBarCodes []string) TransportsAPIGetTransportsRequest {
	r.packageBarCodes = &packageBarCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) Statuses(statuses []string) TransportsAPIGetTransportsRequest {
	r.statuses = &statuses
	return r
}

func (r TransportsAPIGetTransportsRequest) FileCategoryCodes(fileCategoryCodes []string) TransportsAPIGetTransportsRequest {
	r.fileCategoryCodes = &fileCategoryCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) ServiceCodes(serviceCodes []string) TransportsAPIGetTransportsRequest {
	r.serviceCodes = &serviceCodes
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences1(packageReferences1 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences1 = &packageReferences1
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences2(packageReferences2 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences2 = &packageReferences2
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences3(packageReferences3 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences3 = &packageReferences3
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences4(packageReferences4 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences4 = &packageReferences4
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences5(packageReferences5 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences5 = &packageReferences5
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences6(packageReferences6 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences6 = &packageReferences6
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences7(packageReferences7 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences7 = &packageReferences7
	return r
}

func (r TransportsAPIGetTransportsRequest) PackageReferences8(packageReferences8 []string) TransportsAPIGetTransportsRequest {
	r.packageReferences8 = &packageReferences8
	return r
}

// A pattern to look for in fields specified  by PatternFields.
func (r TransportsAPIGetTransportsRequest) Pattern(pattern string) TransportsAPIGetTransportsRequest {
	r.pattern = &pattern
	return r
}

// Fields in which to search for Pattern
func (r TransportsAPIGetTransportsRequest) PatternFields(patternFields []string) TransportsAPIGetTransportsRequest {
	r.patternFields = &patternFields
	return r
}

// Should the query return only missions, only quotations, or both.    Default value is &#39;MissionsOnly&#39;.
func (r TransportsAPIGetTransportsRequest) MissionTypeVisibility(missionTypeVisibility string) TransportsAPIGetTransportsRequest {
	r.missionTypeVisibility = &missionTypeVisibility
	return r
}

// Should the query return active quotations, archived quotations, or both.    Default value is &#39;Active&#39;.
func (r TransportsAPIGetTransportsRequest) QuotationStatusVisibility(quotationStatusVisibility string) TransportsAPIGetTransportsRequest {
	r.quotationStatusVisibility = &quotationStatusVisibility
	return r
}

// Should the query return quotations with no approval required, quotations awaiting for approval, or both.    Default value is &#39;All&#39;.
func (r TransportsAPIGetTransportsRequest) QuotationApprovalStatusVisibility(quotationApprovalStatusVisibility string) TransportsAPIGetTransportsRequest {
	r.quotationApprovalStatusVisibility = &quotationApprovalStatusVisibility
	return r
}

// Flag the last substate applied to a transport for each substate family code.
func (r TransportsAPIGetTransportsRequest) FlagLastSubStateByFamilyCodes(flagLastSubStateByFamilyCodes []string) TransportsAPIGetTransportsRequest {
	r.flagLastSubStateByFamilyCodes = &flagLastSubStateByFamilyCodes
	return r
}

// Should the query return parent, children, or both in the context of fragmented transports.  Defaults to All
func (r TransportsAPIGetTransportsRequest) FragmentationVisibility(fragmentationVisibility string) TransportsAPIGetTransportsRequest {
	r.fragmentationVisibility = &fragmentationVisibility
	return r
}

// Pagination start index (offset). Default is 0.
func (r TransportsAPIGetTransportsRequest) StartIndex(startIndex int32) TransportsAPIGetTransportsRequest {
	r.startIndex = &startIndex
	return r
}

// Pagination count (page size). Default is 10, maximum is 200.
func (r TransportsAPIGetTransportsRequest) Count(count int32) TransportsAPIGetTransportsRequest {
	r.count = &count
	return r
}

// Pagination sorting fields separated with a comma
func (r TransportsAPIGetTransportsRequest) Sort(sort []string) TransportsAPIGetTransportsRequest {
	r.sort = &sort
	return r
}

// Pagination sorting fields with descending order separated with a comma. They must also be provided in sort fields.
func (r TransportsAPIGetTransportsRequest) Desc(desc []string) TransportsAPIGetTransportsRequest {
	r.desc = &desc
	return r
}

// Projection fields separated with a comma
func (r TransportsAPIGetTransportsRequest) Fields(fields string) TransportsAPIGetTransportsRequest {
	r.fields = &fields
	return r
}

func (r TransportsAPIGetTransportsRequest) Execute() (*IPagedResourceListTransportDto, *http.Response, error) {
	return r.ApiService.GetTransportsExecute(r)
}

/*
GetTransports Get transports

Get a paginated list of transports corresponding to the specified filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TransportsAPIGetTransportsRequest
*/
func (a *TransportsAPIService) GetTransports(ctx context.Context) TransportsAPIGetTransportsRequest {
	return TransportsAPIGetTransportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IPagedResourceListTransportDto
func (a *TransportsAPIService) GetTransportsExecute(r TransportsAPIGetTransportsRequest) (*IPagedResourceListTransportDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IPagedResourceListTransportDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.GetTransports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.transportIds != nil {
		t := *r.transportIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "transportIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "transportIds", t, "form", "multi")
		}
	}
	if r.transportUids != nil {
		t := *r.transportUids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "transportUids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "transportUids", t, "form", "multi")
		}
	}
	if r.missionNumbers != nil {
		t := *r.missionNumbers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "missionNumbers", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "missionNumbers", t, "form", "multi")
		}
	}
	if r.missionUids != nil {
		t := *r.missionUids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "missionUids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "missionUids", t, "form", "multi")
		}
	}
	if r.missionTrackIds != nil {
		t := *r.missionTrackIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "missionTrackIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "missionTrackIds", t, "form", "multi")
		}
	}
	if r.driverIds != nil {
		t := *r.driverIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "driverIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "driverIds", t, "form", "multi")
		}
	}
	if r.vehicleCodes != nil {
		t := *r.vehicleCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "vehicleCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "vehicleCodes", t, "form", "multi")
		}
	}
	if r.trailerCodes != nil {
		t := *r.trailerCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "trailerCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "trailerCodes", t, "form", "multi")
		}
	}
	if r.vehicleOrTrailerCodes != nil {
		t := *r.vehicleOrTrailerCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "vehicleOrTrailerCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "vehicleOrTrailerCodes", t, "form", "multi")
		}
	}
	if r.quotationNumbers != nil {
		t := *r.quotationNumbers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "quotationNumbers", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "quotationNumbers", t, "form", "multi")
		}
	}
	if r.quotationUids != nil {
		t := *r.quotationUids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "quotationUids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "quotationUids", t, "form", "multi")
		}
	}
	if r.quotationTrackIds != nil {
		t := *r.quotationTrackIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "quotationTrackIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "quotationTrackIds", t, "form", "multi")
		}
	}
	if r.sourceQuotationNumbers != nil {
		t := *r.sourceQuotationNumbers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sourceQuotationNumbers", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sourceQuotationNumbers", t, "form", "multi")
		}
	}
	if r.dateFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dateFrom", r.dateFrom, "", "")
	}
	if r.dateTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dateTo", r.dateTo, "", "")
	}
	if r.dateActions != nil {
		t := *r.dateActions
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "dateActions", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "dateActions", t, "form", "multi")
		}
	}
	if r.dateToDateActions != nil {
		t := *r.dateToDateActions
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "dateToDateActions", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "dateToDateActions", t, "form", "multi")
		}
	}
	if r.dateCombinationCriterion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dateCombinationCriterion", r.dateCombinationCriterion, "", "")
	} else {
		var defaultValue string = "Or"
		r.dateCombinationCriterion = &defaultValue
	}
	if r.signs != nil {
		t := *r.signs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "signs", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "signs", t, "form", "multi")
		}
	}
	if r.references1 != nil {
		t := *r.references1
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "references1", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "references1", t, "form", "multi")
		}
	}
	if r.references2 != nil {
		t := *r.references2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "references2", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "references2", t, "form", "multi")
		}
	}
	if r.references3 != nil {
		t := *r.references3
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "references3", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "references3", t, "form", "multi")
		}
	}
	if r.pickupCityIds != nil {
		t := *r.pickupCityIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCityIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCityIds", t, "form", "multi")
		}
	}
	if r.pickupPostCodes != nil {
		t := *r.pickupPostCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupPostCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupPostCodes", t, "form", "multi")
		}
	}
	if r.pickupCityNames != nil {
		t := *r.pickupCityNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCityNames", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCityNames", t, "form", "multi")
		}
	}
	if r.pickupAddressNames != nil {
		t := *r.pickupAddressNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupAddressNames", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupAddressNames", t, "form", "multi")
		}
	}
	if r.pickupCountryCodes != nil {
		t := *r.pickupCountryCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCountryCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupCountryCodes", t, "form", "multi")
		}
	}
	if r.pickupSectorCodes != nil {
		t := *r.pickupSectorCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pickupSectorCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pickupSectorCodes", t, "form", "multi")
		}
	}
	if r.deliveryCityIds != nil {
		t := *r.deliveryCityIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCityIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCityIds", t, "form", "multi")
		}
	}
	if r.deliveryPostCodes != nil {
		t := *r.deliveryPostCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryPostCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryPostCodes", t, "form", "multi")
		}
	}
	if r.deliveryCityNames != nil {
		t := *r.deliveryCityNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCityNames", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCityNames", t, "form", "multi")
		}
	}
	if r.deliveryAddressNames != nil {
		t := *r.deliveryAddressNames
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryAddressNames", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryAddressNames", t, "form", "multi")
		}
	}
	if r.deliveryCountryCodes != nil {
		t := *r.deliveryCountryCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCountryCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryCountryCodes", t, "form", "multi")
		}
	}
	if r.deliverySectorCodes != nil {
		t := *r.deliverySectorCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "deliverySectorCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "deliverySectorCodes", t, "form", "multi")
		}
	}
	if r.customerCodes != nil {
		t := *r.customerCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "customerCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "customerCodes", t, "form", "multi")
		}
	}
	if r.ordererCodes != nil {
		t := *r.ordererCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ordererCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ordererCodes", t, "form", "multi")
		}
	}
	if r.agencyCodes != nil {
		t := *r.agencyCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "agencyCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "agencyCodes", t, "form", "multi")
		}
	}
	if r.packageBarCodes != nil {
		t := *r.packageBarCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageBarCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageBarCodes", t, "form", "multi")
		}
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "form", "multi")
		}
	}
	if r.fileCategoryCodes != nil {
		t := *r.fileCategoryCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fileCategoryCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fileCategoryCodes", t, "form", "multi")
		}
	}
	if r.serviceCodes != nil {
		t := *r.serviceCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serviceCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serviceCodes", t, "form", "multi")
		}
	}
	if r.packageReferences1 != nil {
		t := *r.packageReferences1
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences1", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences1", t, "form", "multi")
		}
	}
	if r.packageReferences2 != nil {
		t := *r.packageReferences2
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences2", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences2", t, "form", "multi")
		}
	}
	if r.packageReferences3 != nil {
		t := *r.packageReferences3
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences3", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences3", t, "form", "multi")
		}
	}
	if r.packageReferences4 != nil {
		t := *r.packageReferences4
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences4", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences4", t, "form", "multi")
		}
	}
	if r.packageReferences5 != nil {
		t := *r.packageReferences5
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences5", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences5", t, "form", "multi")
		}
	}
	if r.packageReferences6 != nil {
		t := *r.packageReferences6
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences6", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences6", t, "form", "multi")
		}
	}
	if r.packageReferences7 != nil {
		t := *r.packageReferences7
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences7", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences7", t, "form", "multi")
		}
	}
	if r.packageReferences8 != nil {
		t := *r.packageReferences8
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences8", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "packageReferences8", t, "form", "multi")
		}
	}
	if r.pattern != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pattern", r.pattern, "", "")
	}
	if r.patternFields != nil {
		t := *r.patternFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "patternFields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "patternFields", t, "form", "multi")
		}
	}
	if r.missionTypeVisibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "missionTypeVisibility", r.missionTypeVisibility, "", "")
	} else {
		var defaultValue string = "MissionsOnly"
		r.missionTypeVisibility = &defaultValue
	}
	if r.quotationStatusVisibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quotationStatusVisibility", r.quotationStatusVisibility, "", "")
	} else {
		var defaultValue string = "Active"
		r.quotationStatusVisibility = &defaultValue
	}
	if r.quotationApprovalStatusVisibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quotationApprovalStatusVisibility", r.quotationApprovalStatusVisibility, "", "")
	} else {
		var defaultValue string = "All"
		r.quotationApprovalStatusVisibility = &defaultValue
	}
	if r.flagLastSubStateByFamilyCodes != nil {
		t := *r.flagLastSubStateByFamilyCodes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "flagLastSubStateByFamilyCodes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "flagLastSubStateByFamilyCodes", t, "form", "multi")
		}
	}
	if r.fragmentationVisibility != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fragmentationVisibility", r.fragmentationVisibility, "", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "csv")
	}
	if r.desc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "desc", r.desc, "form", "csv")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPITransportUpdateDryRunRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *UpdateTransportDto
}

func (r TransportsAPITransportUpdateDryRunRequest) Command(command UpdateTransportDto) TransportsAPITransportUpdateDryRunRequest {
	r.command = &command
	return r
}

func (r TransportsAPITransportUpdateDryRunRequest) Execute() (*TransportEntryDryRunDto, *http.Response, error) {
	return r.ApiService.TransportUpdateDryRunExecute(r)
}

/*
TransportUpdateDryRun Transport update dry run

Simulate a transport update : performs a dry run.
The following permission(s) are required to access this route:  Update orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPITransportUpdateDryRunRequest
*/
func (a *TransportsAPIService) TransportUpdateDryRun(ctx context.Context, uid string) TransportsAPITransportUpdateDryRunRequest {
	return TransportsAPITransportUpdateDryRunRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
//
//	@return TransportEntryDryRunDto
func (a *TransportsAPIService) TransportUpdateDryRunExecute(r TransportsAPITransportUpdateDryRunRequest) (*TransportEntryDryRunDto, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TransportEntryDryRunDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.TransportUpdateDryRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/dry-run"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return localVarReturnValue, nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransportsAPIUpdateTransportRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *UpdateTransportDto
}

// Update transport command.
func (r TransportsAPIUpdateTransportRequest) Command(command UpdateTransportDto) TransportsAPIUpdateTransportRequest {
	r.command = &command
	return r
}

func (r TransportsAPIUpdateTransportRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateTransportExecute(r)
}

/*
UpdateTransport Update transport

Updates a transport using merge patch semantics : you need to provide only the fields to updates.
Fields that are not present in the request will be preserved, and fields set to null will be cleared.
Special process for lists with nested objects :
The objects that are not provided in a list are not updated.
For all objects provided in a list : specify the merge action (add, update, remove) and the entity key when the action is update or remove.
If the change requires a modification on a related transport (for example, for a round trip), you can update the related mission or quotation instead.
The following permission(s) are required to access this route:  Update orders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid Transport's unique identifier
	@return TransportsAPIUpdateTransportRequest
*/
func (a *TransportsAPIService) UpdateTransport(ctx context.Context, uid string) TransportsAPIUpdateTransportRequest {
	return TransportsAPIUpdateTransportRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) UpdateTransportExecute(r TransportsAPIUpdateTransportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.UpdateTransport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIUpdateTransportAssignmentRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *AssignTransportDto
}

func (r TransportsAPIUpdateTransportAssignmentRequest) Command(command AssignTransportDto) TransportsAPIUpdateTransportAssignmentRequest {
	r.command = &command
	return r
}

func (r TransportsAPIUpdateTransportAssignmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateTransportAssignmentExecute(r)
}

/*
UpdateTransportAssignment Update transport assignment

Update transport's assignment data: driver, vehicle, trailer, handler, contractor agent.
When a driver is unset, the transport is un-assigned.
When the driver is updated, the transport is assigned to this driver.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIUpdateTransportAssignmentRequest
*/
func (a *TransportsAPIService) UpdateTransportAssignment(ctx context.Context, uid string) TransportsAPIUpdateTransportAssignmentRequest {
	return TransportsAPIUpdateTransportAssignmentRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) UpdateTransportAssignmentExecute(r TransportsAPIUpdateTransportAssignmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.UpdateTransportAssignment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/assignment"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/merge-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TransportsAPIUpdateTransportAssignmentStatusRequest struct {
	ctx        context.Context
	ApiService *TransportsAPIService
	uid        string
	command    *ApplyAssignmentStatusChangeCommand
}

func (r TransportsAPIUpdateTransportAssignmentStatusRequest) Command(command ApplyAssignmentStatusChangeCommand) TransportsAPIUpdateTransportAssignmentStatusRequest {
	r.command = &command
	return r
}

func (r TransportsAPIUpdateTransportAssignmentStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateTransportAssignmentStatusExecute(r)
}

/*
UpdateTransportAssignmentStatus Update transport assignment status

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uid
	@return TransportsAPIUpdateTransportAssignmentStatusRequest
*/
func (a *TransportsAPIService) UpdateTransportAssignmentStatus(ctx context.Context, uid string) TransportsAPIUpdateTransportAssignmentStatusRequest {
	return TransportsAPIUpdateTransportAssignmentStatusRequest{
		ApiService: a,
		ctx:        ctx,
		uid:        uid,
	}
}

// Execute executes the request
func (a *TransportsAPIService) UpdateTransportAssignmentStatusExecute(r TransportsAPIUpdateTransportAssignmentStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransportsAPIService.UpdateTransportAssignmentStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/transports/{uid}/assignment/status"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", url.PathEscape(parameterValueToString(r.uid, "uid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.command == nil {
		return nil, reportError("command is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.command
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
