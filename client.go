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
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// APIClient manages communication with the Dispatch Api 3.0 API vv3.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AddressesAPI *AddressesAPIService

	AgenciesAPI *AgenciesAPIService

	ApiFlagsAPI *ApiFlagsAPIService

	CitiesAPI *CitiesAPIService

	CompaniesAPI *CompaniesAPIService

	ContractorAgentsAPI *ContractorAgentsAPIService

	CountriesAPI *CountriesAPIService

	CustomersAPI *CustomersAPIService

	DriversAPI *DriversAPIService

	FileCategoriesAPI *FileCategoriesAPIService

	MeAPI *MeAPIService

	MissionsAPI *MissionsAPIService

	OrderersAPI *OrderersAPIService

	PackageNaturesAPI *PackageNaturesAPIService

	PricingPathsAPI *PricingPathsAPIService

	QuotationsAPI *QuotationsAPIService

	SectorsAPI *SectorsAPIService

	ServicesAPI *ServicesAPIService

	SubServicesAPI *SubServicesAPIService

	SubstatesAPI *SubstatesAPIService

	TeamsAPI *TeamsAPIService

	TransportDocumentFilesAPI *TransportDocumentFilesAPIService

	TransportDocumentLinksAPI *TransportDocumentLinksAPIService

	TransportDocumentReportsAPI *TransportDocumentReportsAPIService

	TransportsAPI *TransportsAPIService

	UnitsAPI *UnitsAPIService

	UsersAPI *UsersAPIService

	VehicleTypesAPI *VehicleTypesAPIService

	VehiclesAPI *VehiclesAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AddressesAPI = (*AddressesAPIService)(&c.common)
	c.AgenciesAPI = (*AgenciesAPIService)(&c.common)
	c.ApiFlagsAPI = (*ApiFlagsAPIService)(&c.common)
	c.CitiesAPI = (*CitiesAPIService)(&c.common)
	c.CompaniesAPI = (*CompaniesAPIService)(&c.common)
	c.ContractorAgentsAPI = (*ContractorAgentsAPIService)(&c.common)
	c.CountriesAPI = (*CountriesAPIService)(&c.common)
	c.CustomersAPI = (*CustomersAPIService)(&c.common)
	c.DriversAPI = (*DriversAPIService)(&c.common)
	c.FileCategoriesAPI = (*FileCategoriesAPIService)(&c.common)
	c.MeAPI = (*MeAPIService)(&c.common)
	c.MissionsAPI = (*MissionsAPIService)(&c.common)
	c.OrderersAPI = (*OrderersAPIService)(&c.common)
	c.PackageNaturesAPI = (*PackageNaturesAPIService)(&c.common)
	c.PricingPathsAPI = (*PricingPathsAPIService)(&c.common)
	c.QuotationsAPI = (*QuotationsAPIService)(&c.common)
	c.SectorsAPI = (*SectorsAPIService)(&c.common)
	c.ServicesAPI = (*ServicesAPIService)(&c.common)
	c.SubServicesAPI = (*SubServicesAPIService)(&c.common)
	c.SubstatesAPI = (*SubstatesAPIService)(&c.common)
	c.TeamsAPI = (*TeamsAPIService)(&c.common)
	c.TransportDocumentFilesAPI = (*TransportDocumentFilesAPIService)(&c.common)
	c.TransportDocumentLinksAPI = (*TransportDocumentLinksAPIService)(&c.common)
	c.TransportDocumentReportsAPI = (*TransportDocumentReportsAPIService)(&c.common)
	c.TransportsAPI = (*TransportsAPIService)(&c.common)
	c.UnitsAPI = (*UnitsAPIService)(&c.common)
	c.UsersAPI = (*UsersAPIService)(&c.common)
	c.VehicleTypesAPI = (*VehicleTypesAPIService)(&c.common)
	c.VehiclesAPI = (*VehiclesAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString(obj interface{}, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		if actualObj, ok := obj.(interface{ GetActualInstanceValue() interface{} }); ok {
			return fmt.Sprintf("%v", actualObj.GetActualInstanceValue())
		}

		return fmt.Sprintf("%v", obj)
	}
	param, ok := obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	v := reflect.ValueOf(obj)
	value := ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			indValue := reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			lenIndValue := indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				arrayValue := indValue.Index(i)
				keyPrefixForCollectionType := keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			indValue := reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
		break
	case map[string]string:
		valuesMap[keyPrefix] = value
		break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
	fileBytes    []byte
	fileName     string
	formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile,
) (localVarRequest *http.Request, err error) {
	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
					return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
