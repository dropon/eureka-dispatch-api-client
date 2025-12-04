/*
Dispatch Api 3.0

This API is dedicated to operators.    Get API url and your client credentials from [integrators website](https://integrators.eureka-technology.fr/).      You can download the Open Api specifications (swagger json document) in the link mentioned above.    Please use our [customer support portal](https://eureka-technology.atlassian.net/servicedesk/customer/portals) for questions or issues.  If you do not have a support account yet, please reach out to your contact at Eureka-Technology to create one.  If you don't have a dedicated contact, you can send an email to support@eureka-technology.com.    # Authentication      ## Access tokens    This Api uses OAuth2 authentication.      All the requests to the api, except authentication ones, must include a valid bearer token.    To obtain an access token, call the _/token_ endpoint, provide an Authorization header with a basic authentication type, constructed using your client_id and client_secret.    The body must be a form url encoded with the parameters corresponding to the authentication flow (see grant types section for more details).    The api returns a response containing the access token, its expiration in seconds, a refresh token, and other properties. Example :    ```json  {      \"access_token\": \"ey...\",      \"token_type\": \"bearer\",      \"expires_in\": 3599,      \"refresh_token\": \"...\",      \".issued\": \"(date)\",      \".expires\": \"(date)\"  }  ```  When the client id or client secret is invalid, a 400 Bad request is returned :  ```json  {      \"error\": \"invalid_client\",      \"error_description\": \"The client id or secret is incorrect.\"  }  ```  The returned access token can then be used in the requests to the api, it must be included in the Authorization header, example : `Authorization: Bearer ey...`.    The access token is valid 1 hour by default, this value can change depending on the carrier.      ### Example using postman      Here is an example of token request using postman (to be adapted according to the grant type).      The request header contains the basic authentication.    ![Authentication authorization header](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-authorization-header.png)    The request body is a form url encoded containing the grant type and the additional parameters (depending on the grant type).    ![Authentication body with password](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-body-with-password.png)    The response contains the access token and additional information.    ![Authentication response](https://sterkfrcragrs.z28.web.core.windows.net/api/authentication/api-authentication-response.png)    ## Refresh token      In _/token_ response, you receive an access token and a refresh token.    The refresh token allows to get a new access token without providing the original credentials.    When an access token has expired, you can exchange the refresh token for a new access token.    The refresh token is valid 24 hours by default, this value can change depending on the carrier.    To generate a new access token using this refresh token, use refresh_token grant type. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=refresh_token&refresh_token=your_refresh_token  ```      You can invalidate the refresh token by calling _/revoke_ endpoint. Example :    ```  POST /revoke    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    token=your_refresh_token  ```      ### Refresh token limitation    There is a limit in the number of active refresh tokens per user. It is 10000 by default, this value can change depending on the carrier.    If this limit is reached, you will receive the error _max_number_of_sessions_reached_, you will need to wait for refresh token expiration before being able to get a new access token for the user.    To avoid the limit, reuse the access token, do not ask for a new one for each request to the Api.      ## Grant types    ### Password flow      Use password grant to connect as a user, using its credentials : provide the parameters grant_type _password_, username and password. Example :    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=password&username=user_username&password=user_password  ```      When user's credentials are invalid, a 400 Bad request is returned :  ```json  {      \"error\": \"invalid_grant\",      \"error_description\": \"Unknown user or incorrect password.\"  }  ```      ### Client credentials      Use client credentials to connect as an application, outside of the context of a user : provide the parameters grant_type _client_credentials_. Example :    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=client_credentials  ```      ### SSO      To connect with an external SSO provider, SSO settings must be configured in Dispatch, and you will have a SSO configuration code for each provider.    SSO users must already be configured in Dispatch.      #### SSO Authorization code      This is the standard method to use an external OAUTH SSO provider configured in Dispatch. Use this method if your application does not have its own client id and client secret in the external SSO provider.      - First, your application will need to open a browser to send the user to the external SSO oauth provider.      To get the authorization url, use the operation _sso/oauth-authorization-url_ and provide the parameters : ssoCode, redirectUri and state, and provide bearer token created with a client credentials authentication.      - When the user connects, the oauth provider will redirect back to your application url and provide an authorization code and the state you provided      - Call token endpoint in the API with grant_type authorization_code_TheSsoConfigurationCode and parameters authorization_code and redirect_uri. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=authorization_code_ABC&authorization_code=UrlEncode(authorizationCode)&redirect_uri=UrlEncode(https://yourwebsite/sso/oauth/callback)  ```      - The API will validate the authorization code, and you will receive a token for this API for the user corresponding to the SSO user.      #### SSO Access token      This is the another method to use an external OAUTH SSO provider configured in Dispatch. You can use this method if your application has its own external SSO provider information (client id and client secret).    - From your application, authenticate to the external SSO provider using authorization code flow :       - Build the authorization url with your SSO credentials       - After user connection, you receive an authorization code        - Call the SSO provider to get an access token    - Call token endpoint in the API with grant_type access_token_TheSsoConfigurationCode and parameters access_token. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=access_token_ABC&access_token=UrlEncode(ssoAccessToken)  ```      - The API will validate the sso access token, and you will receive a token for this API for the user corresponding to the SSO user.      #### SAML      This is the standard method to use an external SSO provider configured in Dispatch and using SAML V2 protocol.    - First, your application will need to open a browser to send the user to the external SSO SAML provider.      To get the authorization url, use the operation _sso/saml-authorization-url_ and provide the parameters : ssoCode and relayState, and provide bearer token created with a client credentials authentication.      - When the user connects, the SAML provider will redirect back to your application with additional information in query parameters on in a form        - if information is received with a POST, read the request form and build a query string using form data        - if information is received with a GET (redirect method), get the query string      - Call token endpoint in the API with grant_type saml_TheSsoConfigurationCode and parameters saml_http_method and saml_content. Example :      ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=saml_ABC&saml_http_method=GET&saml_content=UrlEncode(SAMLResponse=value1&RelayState=value2)  ```      - The API will validate the SAML content, and you will receive a token for this API for the user corresponding to the SSO user.      # General instructions      ## Projection      To select the fields returned by the api, provide 'fields' parameter in query string: list of fields separated with a comma.     When fields are not provided, calls will only return the first level fields, unless specified otherwise in a route description.      Example with simple fields : `GET route_to_my_resource/myResourceUid?fields=field1,field2,field3` returns :    ```json  {      \"field1\": \"My field 1\",      \"field2\": \"My field 2\",      \"field3\": \"My field 3\"  }  ```      Example with nested properties and collections : `GET route_to_my_resource/myResourceUid?fields=field1,nested1/nested2/nestedProperty3,myCollection1/field2` returns :    ```json  {      \"field1\": \"My field 1\",      \"nested1\": {          \"nested2\": {              \"nestedProperty3\": \"my nested property value 3\"          }      }      \"myCollection1\": [          {              \"field2\": \"my field 2\"          }      ]  }  ```      You can group properties on a nested object using parenthesis. Example to get the properties nested1.field1 and nested1.field2 : `fields=nested1(field1,field2)`.      Example to get the properties at the first level only, without nested properties or collections : `route_to_my_resource?fields=~`.      Example to get all the properties, included nested properties and collections : `route_to_my_resource?fields=*`.      ## Filtering      Access to paginated resources accept filtering with query parameters. Example : `route_to_my_paginated_resources?myFilter1=myValue1&myFilter2=myValue2`.      When a query parameter is a collection, the values must be delimited with a comma. To escape a comma in a value, double it.    Example to pass the words : \"hello,world\", \"value1\", \"value2\" to \"names\" parameter : `route_to_my_paginated_resources?names=hello,,world,value1,value2`.      When the description of a filter indicates that it accepts a pattern, you can search with _startsWith_, pattern.    In some cases, when it is specified in the description, the filter also supports _endsWith_ or _contains_ patterns.      Examples :   - Exact match : `myFilter=abc`    - Starts with abc : `myFilter=abc*`    - Ends with abc : `myFilter=*abc`    - Contains abc : `myFilter=*abc*`       For some operations, you can search on multiple fields, using pattern and pattern fields. The fields can be different depending on the operation, check the description of _patternFields_ parameter to get the available fields.    Example, to search resources having a code or a label starting with abc : `pattern=abc*&patternFields=code,label`.      ## Pagination      For paginated resources, provide query option's _startIndex_, _count_, _fields_, _sort_ for ascending ordering and _desc_ for descending ordering.    Example : `route_to_my_paginated_resources?startIndex=0&count=10&fields=field1,field2&sort=field2`    Example of response :    ```json  {      \"data\": [          {              \"field1\": \"My field 1 value 1\",              \"field2\": \"My field 2 value 1\"          },          {              \"field1\": \"My field 1 value 2\",              \"field2\": \"My field 2 value 2\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```  ## Errors      In case of errors, an object Problem Details is returned with its http status code.    It corresponds to the standard problem details, with an additional property _data_.    See more details about RFC 7807 problem details [here](https://tools.ietf.org/html/rfc7807).    Example :    ```json  {      \"type\": \"/problems/types/service-unavailable\",      \"title\": \"Service unavailable.\",      \"status\": 503,      \"instance\": \"/problems/instances/dispatch-module-not-available\"  }  ```    When a business error occurs, an error code and error message are included in an additional property.    Example :    ```json  {      \"data\": {          \"errors\": [              {                  \"code\": \"resource_not_found\",                  \"message\": \"Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",              }          ]      },      \"type\": \"/problems/types/not-found\",      \"title\": \"A business error occurred\",      \"status\": 404,      \"detail\": \"resource_not_found : Resource '26e32d5e-6412-408b-9a03-c5bf9ad4dbec' not found\",      \"instance\": \"/problems/instances/business-error\"  }  ```    ### Error codes    #### General    Error code | Error description    ----|----    missing_mandatory_parameter | A mandatory parameter is missing  must_be_strictly_greater_than_zero | The parameter must be greater than zero  resource_not_found | The resource does not exist  resource_already_exists | A resource with the same identifier already exists  user_missing_permission | The user is not allowed to perform the operation or to access the resource  datetime_out_of_range | The date time is out of the range supported by the system  guid_should_not_be_empty | The unique identifier GUID must not be empty  invalid_enum_value | The provided value is not a valid enumeration value   invalid_field_name | The field provided in the query parameters is not valid  field_not_supported | The field in the 'fields' or 'sort' parameters is not supported  user_entity_access_forbidden | The user is not allowed to access the resource  value_out_of_range | The provided value is out of the expected range  max_length_exceeded | The maximum length of the parameter is exceeded  end_date_must_be_greater_than_start_date | The end date must be greater than the start date  max_date_interval_reached | The maximum date interval is reached  user_has_no_eureka_maps_key | The user does not have access to Eureka Maps  max_user_application_preferences_reached | The maximum number of application preferences is reached for this user    #### Drivers    Error code | Error description    ----|----    driver_unavailability_reason_in_use | The driver unavailability reason is used  driver_unavailability_reason_disabled | The driver unavailability reason is disabled  driver_unavailability_overlapping | There is a driver unavailability in the same interval  contractor_agent_subcontractor_mandatory | The subcontractor is mandatory when the contractor agent is provided  contractor_agent_invalid_subcontractor | The provided contractor agent cannot be used with the provided subcontractor    #### Missions    Error code | Error description    ----|----    transport_assignment_constraints_violation | Transport assignment constraints not fulfilled  transport_unassignment_constraints_violation | Transport unassignment constraints not fulfilled  transport_fragmentation_constraints_violation | Transport fragmentation constraints not fulfilled  transport_defragmentation_constraints_violation | Transport defragmentation constraints not fulfilled  invalid_file_category | The provided file category is invalid  invalid_custom_parameter | The provided custom parameter is invalid  invalid_packing_nature | The package nature is invalid  invalid_unit_code | The unit code is invalid  invalid_dimension_unit | The dimension unit is invalid  transport_unexpected_state | The operation cannot be performed because the transport has an unexpected status  document_report_unexpected_state | The operation cannot be performed because the transport document report has an unexpected status  invalid_service_code | The provided service code is invalid  service_disabled | The provided service is disabled  invalid_customer_code | The provided customer code is invalid  customer_not_usable | The provided customer cannot be used  invalid_pricing_path | The provided pricing path is invalid  itinerary_computation_failed | Unable to compute itinerary for the transport    #### Addresses    Error code | Error description    ----|----    invalid_address_id | The provided address id is invalid  invalid_city_id | The provided city id is invalid  geocoding_failed | Unable to geocode the address        ## Resource update      You can use PATCH operations to partially update a resource.    The update uses merge patch semantics : you need to provide only the fields to updates.    See more details about JSON merge patch standard [here](https://tools.ietf.org/html/rfc7396).      Fields that are not present in the request will be preserved, and fields set to null will be cleared.      There is a special process for lists with nested objects :    - The objects that are not provided in a list are not updated.    - For all objects provided in a list : specify the merge action (add, update, remove) and the entity key when the action is update or remove.      Example, with comments indicating the performed operations :    ```json  {      \"property1\": \"new value 1\",         // property1 is updated       \"nested1\": {                        // nested1 is added if it doesn't exist, otherwise, it is updated          \"property2\": \"my value 2\",      // property2 is updated          \"nested2\": null,                // nested2 is deleted          \"nested3\": {                    // nested3 is added if it doesn't exist              \"property3\": \"my value 3\"   // property3 is updated          }      },      \"nestedCollection1\": [              // nestedCollection1 is updated          {              \"myIdentifier\": \"id1\",              \"mergeAction\": \"update\",              \"property4\": \"my value 1\"   // the property property4 of the item with myIdentifier equals to id1 is updated          },          {              \"myIdentifier\": \"id2\",              \"mergeAction\": \"remove\"     //  the item with myIdentifier equal to id2 is removed           },          {              \"mergeAction\": \"add\",       // an item is added to the collection with the provided property4.              \"property4\": \"my value 2\"          }      ]  }  ```      ## Date formats    Many dates in the system are saved without timezone information. These dates depend on the context (for example : pickup date correponds generally to the timezone of pickup address, but the offset is not kept in the system). In these cases, the timezone is not taken into account when provided in the input. For example, if you provide `2022-01-31T10:00:00+05:00`, the system will keep only `2022-01-31T10:00:00`, the date will be returned with an unspecified timezone.      The api returns the timezone when it is available in the system.      Examples of dates returned by the api :  * Date with unspecified time zone : `2022-01-31T10:00:00`  * UTC date : `2022-01-31T10:00:00+00:00` or `2022-01-31T10:00:00Z`  * Date with offset +5h : `2022-01-31T10:00:00+05:00`    ## Versioning    This API is versioned to handle changes in API contract.    The versioning is done with URI path. Example : _https://myapi/v1/myresource_.    The versioning applies on all the operations except _/token_ and _/revoke_ endpoints. Example: _https://myapi/token_.      A new version is released when a breaking change is made.    The previous versions remain maintained for a certain time. After a time, old versions may be deprecated (marked as obsolete) then removed.    Remark: breaking changes can be made on pre-release versions. These versions can also be removed from the API and replaced with a release version.      Examples of breaking changes :    - Change a property name or type in a request or response    - Add a required parameter in a request    - Remove a property in a response    - Remove an operation in the API      Examples of non-breaking changes :  - Add a value in a response  - Add an optional parameter in a request    - Add an operation in the API    - Change the format of any opaque identifier  - Change the format of a date time : add or remove timezone information, for example : change from `2022-01-31T10:00:00` to `2022-01-31T10:00:00Z` or to `2022-01-31T10:00:00+05:00`.    - Replacing a success status code with another success status code (i.e: replacing a 200 with a 201)  - Replacing a failure status code with another failure status code (i.e: replacing a 500 with a 400)    When a feature is added (without breaking changes), it is available for the last released API version.    In this case, a flag can be added to indicate that this feature is available. See api flags endpoint.      ## Rate Limiting    This API uses consumption quotas to prevent excessive calls and to maintain service availability.    Two types of rate limiting are applied :  - leaky buckets    - concurrency      You can view these limits in Eureka integrator's portal.    There are also additional limits for calls performed by one bearer token, configured by the carrier. These limits are not visible in the portal.      When a quota is exceeded, the API returns the error code 429 (Too Many Requests). A _Retry-After_ header is included and indicates (in seconds) how long to wait before making a new request.      ### Leaky bucket rate limiting      A leaky bucket rate limiting is applied to each API client, it depends on two parameters : the bucket size and the leak rate.      Each API client have its own bucket of a limited size.    When a request is received :    - If the bucket is not full, the bucket counter is increased.    - If the bucket is full (the bucket counter reaches the bucket size), the request is rejected. In this case, the API client must wait before sending a new request.      The bucket counter is decreased periodically to accept new requests.      You can find more details about leaky bucket algorithm [here](https://en.wikipedia.org/wiki/Leaky_bucket).      ### Concurrency rate limiting      The API accepts only a limited number of concurrent calls for each API client.    If the maximum number of concurrent calls is reached, the request is rejected. In this case, the API client must wait before sending a new request.        # Dispatch Api    ## Available features    This API includes various features :  * Retrieval of multiple resources needed for mission entry  * Mission and quotation operations : creation, update, cancellation  * Operations on transport documents (download, upload, delete)  * Operations on transport : sub state application, assignment, pickup, delivery, termination, cancellation  * Retrieval of transport data  * Transport document reports generation  * Fragmentation    This API does not include these features :  * In mission entry : dangerous goods, air transport, deposits (aka consignes), orders in sequence  * Disputes and misfunction sheets, call for tender  * D2D  * Operations (creation, update, delete) on resources and tools except on transports  * Retrieval of resources not needed for mission entry  * Billing  * Regular orders  * Rounds  * Mission scheduling  * Exports and imports  * Reports not related to one transport    ## Missions, quotations and transports      A transport is a system for conveying something (goods or people) from a place (pickup) to another (delivery). A transport is related to one customer and to one service and may contain packages.      A mission is an order containing one or many transports.    In general, a mission will contain only one transport. But in case of a mission with return, it will contain a second transport.    This API supports missions with one or two transports.      A quotation can be requested to indicate how much an order will cost.    When a quotation is validated, it is transformed into a new mission : a new mission is created with quotation information, and the quotation is archived.      Use :  * GET /transports to retrieve a collection of transports (use filtering and projection)    * POST /missions to create a mission  * POST /quotations to create a quotation        ## Mission entry    The following steps can be used to create a mission or a quotation.    ### Authenticate as a user    Connect as a user, with a login and password    ```  POST /token    Content-Type: application/x-www-form-urlencoded    Authorization: Basic Base64(client_id:client_secret)    grant_type=password&username=my_dispatch_user_login&password=my_dispatch_user_password  ```    The response is a json containing the access token    ```json  {      \"access_token\": \"eyJh...\",      \"token_type\": \"bearer\",      \"expires_in\": 3599,      \"refresh_token\": \"11c...\",      \"as:clientId\": \"...\",      \".issued\": \"Mon, 08 Jul 2024 14:43:45 GMT\",      \".expires\": \"Mon, 08 Jul 2024 15:43:45 GMT\"  }  ```    For all the following calls, add a header `Authorization: Bearer my_access_token`.      ### Select a customer      Get the customers available for the connected user, use filtering to get only the fields you need, use pagination.      ```  GET /customers?startIndex=0&count=10&fields=uid,code,label  ```    Response    ```json  {      \"data\": [          {              \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",              \"code\": \"CUST01\",              \"label\": \"Sample customer 1\"          },          {              \"uid\": \"9acc8b97729c447e9b52bcf87d392c35\",              \"code\": \"CUST02\",              \"label\": \"Sample customer 2\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```  Get customer's information by its unique identifier or by its code. Examples :  ```  GET /customers/1b1548e7d1f046b0ba357539b139be7f?fields==uid,code,label  GET /customers/by-customer-code?customerCode=CUST01&fields==uid,code,label  ```    Response    ```json  {      \"uid\": \"1b1548e7d1f046b0ba357539b139be7f\",      \"code\": \"CUST01\",      \"label\": \"Sample customer 1\"  }  ```    ### Addresses    In address step, validate the selected address by getting the cities and countries available in the system, a valid city can be needed for the pricing. Example :      ```  GET /cities?postCode=34000&cityName=montpellier&countryCode=FR&fields=cityId,name,postCode,country/label  ```    Response    ```json  {      \"data\": [          {              \"cityId\": 17037,              \"name\": \"MONTPELLIER\",              \"postCode\": \"34000\",              \"country\": {                  \"label\": \"FRANCE\"              }          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 1,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 1      }  }  ```    ### Packing    In package entry, get package natures.      ```  GET /package-natures?fields=packageNatureId,code,label,packageFamilyLabel  ```    Response    ```json  {      \"data\": [          {              \"packageNatureId\": 1,              \"code\": \"CARTONA3\",              \"label\": \"Carton A3\",              \"packageFamilyLabel\": \"Colis\"          },          {              \"packageNatureId\": 2,              \"code\": \"Boite\",              \"label\": \"Boite\",              \"packageFamilyLabel\": \"Consommables\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```      ### Services    In service entry, get services.    ```  GET /services?startIndex=0&count=10&fields=uid,code,label,serviceFamily/label  ```    Response    ```json  {      \"data\": [          {              \"uid\": \"4da0ba3ef10c4147b73debd01d375000\",              \"code\": \"T1\",              \"label\": \"BREAK 500 KG\",              \"serviceFamily\": {                  \"label\": \"Poids Lourds\"              }          },          {              \"uid\": \"1cfb5b59d877406aa7b86908c9c12f63\",              \"code\": \"T1-MO\",              \"label\": \"MOTO 50KG\",              \"serviceFamily\": {                  \"label\": \"Prestations intra muros\"              }          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```    TO see the sub services included in a service, get its sub services.    ```  GET /services/4da0ba3ef10c4147b73debd01d375000/sub-services?fields=subServiceCode,subServiceLabel,defaultQuantity  ```    Response    ```json  {      \"data\": [          {              \"subServiceCode\": \"ADV\",              \"subServiceLabel\": \"Assurance ad valorem\",              \"defaultQuantity\": 1.00          },          {              \"subServiceCode\": \"KM\",              \"subServiceLabel\": \"Kilomètres\",              \"defaultQuantity\": 0.00          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```    For additional sub services, use sub-services route with available filtering and paging. Example:     ```  GET /sub-services?isEnabled=true&code=KM*&patternFields=code,label  ```    Response    ```json  {      \"data\": [          {              \"code\": \"KM\",              \"label\": \"Kilomètres\",              \"isEnabled\": true,              \"unitCode\": \"KM\",              \"comment\": null,              \"vatCode\": \"1\",              \"quantityMustBeSelected\": false,              \"subServiceQuantities\": null          },          {              \"code\": \"KM-NS\",              \"label\": \"Kilomètres sans surcharge carburant\",              \"isEnabled\": true,              \"unitCode\": \"KM\",              \"comment\": null,              \"vatCode\": \"2\",              \"quantityMustBeSelected\": false,              \"subServiceQuantities\": null          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": null,          \"pageCount\": 1,          \"totalItemCount\": 2,          \"pageNumber\": 1,          \"pageSize\": 10,          \"hasPreviousPage\": false,          \"hasNextPage\": false,          \"isFirstPage\": true,          \"isLastPage\": true,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 2      }  }  ```      ### Mission creation    Before creating a mission, you can perform a dry run by using `POST /missions/dry-run` route. It simulates a mission creation without creating it. Example:      ```json  {      \"transports\": [          {              \"customerCode\": \"ClientDemo1\",              \"agencyCode\": \"01\",              \"pickupStep\": {                  \"stepAddress\": {                      \"addressId\": 314312                  },                  \"stepDate\": {                      \"plannedStart\": {                          \"dateTime\": \"2024-07-02T00:00:00\",                          \"isTimeOfDayIgnored\": true                      }                  }              },              \"deliveryStep\": {                  \"stepAddress\": {                      \"AddressId\": 339923                  },                  \"stepDate\": {                      \"plannedStart\": {                          \"dateTime\": \"2024-07-03T00:00:00\",                          \"isTimeOfDayIgnored\": true                      }                  }              },              \"serviceCode\": \"T8\",              \"subServices\": [                  {                      \"code\": \"KM\"                  },                  {                      \"code\": \"ADV\",                      \"forcedSellQuantity\": 11                  }              ],              \"serviceCustomParameters\": [                  {                      \"code\": \"PAL\",                      \"Value\": \"2\"                  }              ]          }      ]  }  ```    This route returns information about the mission with its pricing, but without creating it.    To create a mission, use the same request, with `POST /missions` route.        ### Additional remarks on transport creation    * The provided entities (example: customer, service, agency) must be available for the connected user regarding its restrictions.    * The provided customer usable and not a prospect. If the orderer is mandatory for a customer, the orderer must be provided.    * The provided service must be enabled. If it has an agency restriction, it must match customer's agency.    * All provided identifiers (example : city id, agency code, VAT code, unit code) must exist  * Customer and service's custom parameters are checked. Example : the command is rejected if a custom parameter is mandatory, not provided and does not have a default value.    * Transport's references are checked. Examples : if mandatory and not provided, if must be selected from a list but does not match this list.    * Pickup and delivery dates are mandatory  * For package lines, provide either the nature id or a nature (free text). Default values are copied from the nature.    To change this behavior, enable _skipPackageNatureDefaultValuesApplication_ in command's api behavior.      * For transport's sub services: when they are not provided, the transport uses service's default sub services. Otherwise, the transport uses only the specified sub services : service's default sub services are not included and must be specified manually if they are needed.    * Sell pricing path can be forced.   The pricing path must be a customer pricing path (not a subcontractor's one).    When provided, the addresses are overwritten with the ones defined in the provided sell pricing path.      The provided service must match the one defined in the pricing path.    * The mission or quotation is created as unique if at least two transports are provided. To force this parameter, provide _forceUniqueOrder_ in command's api behavior.    * Itinerary computation is done if configured in global settings and if the distance is not provided. To ignore failures during this step, enable _ignoreItineraryComputationFailures_ in command's api behavior.    * Transport steps geocoding is attempted if itinerary computation is required. To try to geocode a step in other cases, enable _tryGeocodeTransportSteps_ in command's api behavior.     * Gas emission is computed if enabled in global settings and not forced.   * Transport bill address is unsed only when the customer is occasional.      ## Transport retrieval    Here is an example to search transports using projection, filtering and pagination.    In this example, we want the 5 top transports having a pickup post code 34000, for customer code CUST01 and with statuses C (assigned) and 0 (recorded), we only want the fields uid, missionNumber, serviceCode and status.    ```  GET /transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=0&count=5  ```  Response  ```json  {      \"data\": [          {              \"uid\": \"fcba3be8e09740499079f86dff98f5e6\",              \"missionNumber\": 1,              \"status\": \"C\",              \"serviceCode\": \"MO\"          },          {              \"uid\": \"10b9c9ba9b5240cba50fb834bf5cb714\",              \"missionNumber\": 7203,              \"status\": null,              \"serviceCode\": \"TNT\"          },          {              \"uid\": \"bfe364f5478a4f2890ea1cba6c535a99\",              \"missionNumber\": 13443,              \"status\": null,              \"serviceCode\": \"T1\"          },          {              \"uid\": \"87253a9702134a198be1748d02eec64d\",              \"missionNumber\": 13444,              \"status\": null,              \"serviceCode\": \"T1\"          },          {              \"uid\": \"5febb2485f474b978730c805ad6e7b79\",              \"missionNumber\": 13445,              \"status\": null,              \"serviceCode\": \"T1\"          }      ],      \"paging\": {          \"prevLink\": null,          \"nextLink\": \"https://mylicense.dispatchapi.dispatch-rts.com:443/v3/transports?pickupPostCodes=34000&customerCodes=CUST01&statuses=C,0&fields=uid,missionNumber,serviceCode,status&startIndex=5&count=5\",          \"pageCount\": 2,          \"totalItemCount\": 6,          \"pageNumber\": 1,          \"pageSize\": 5,          \"hasPreviousPage\": false,          \"hasNextPage\": true,          \"isFirstPage\": true,          \"isLastPage\": false,          \"firstItemOnPage\": 1,          \"lastItemOnPage\": 5      }  }  ```      ## Transport update    Here is an example to partially update a transport using PATCH operation.      ```  PATCH /transports/dfa98649e1fc4cf8a77c43351b531007  {      \"apiBehavior\": {          \"SkipPackageNatureDefaultValuesApplication\": false,          \"updateTransportOperationComment\": \"My update operation 01\"      },      \"reference2\": \"new-reference-2\",                // The reference2 will be updated      \"pickupStep\": {          \"stepDate\": {              \"plannedStart\": {                  \"dateTime\": \"2022-02-24 10:00:00\",  // the pickup planned start time will be updated                  \"isTimeOfDayIgnored\": false              }          }      },      \"packing\": {          \"lines\": [              {                  \"mergeAction\": \"update\",            // the quantity and references for package id 2268635 will be updated                  \"packageLineId\": 2268635,                  \"quantity\": 3.00,                  \"references\": [                      {                          \"mergeAction\": \"update\",    // the value of the reference with index 3 will be updated                          \"referenceIndex\": 3,                          \"value\": \"abc\"                      }                  ]              },              {                  \"mergeAction\": \"remove\",            // the package with id 2268636 will be removed                  \"packageLineId\": 2268636              },              {                  \"mergeAction\": \"add\",               // this package will be added                  \"nature\": \"Plis\",                  \"quantity\": 2.00,                  \"dimensionUnit\": \"cm\",                  \"height\": 1.00,                  \"width\": 10.00,                  \"length\": 15.00              }          ]      },      \"subServices\": [          {              \"mergeAction\": \"add\",                   // this service will be added              \"code\": \"SP01\",              \"forcedSellQuantity\": 1.2          }      ]  }  ```      ## Transport statuses    To follow up the progress of a transport, check transport status.    When a transport is created, it is recorded, then assigned to a driver, picked up from pickup address, delivered to delivery address, terminated and invoiced.    A transport has one of the following statuses:      Status | Description  ------ | -------  0 (zero), null or empty string | Recorded  A | Canceled  X | Unassigned  b | Pre assigned  B | Pre assigned and accepted by the driver  C | Assigned  E | Picked up  L | Delivered  T | Terminated  s | Dispute undealt  S | Dispute dealt  f | Pre billed  F | Invoiced      Remarks :  - Transport statuses are case sensitive. Make sure you use the correct case when filtering the transports based on their statuses.    - To filter recorded transports only, use the value 0 (zero). The api does not provide filtering with empty values.         ## Transport history statuses    Transport history lines returned by this api have one of the following status codes :      Status | Description  ------ | -------  0 (zero) | Recorded  M | Modified  C | Assigned  E | Picked up  L | Delivered        ## Dispatch desktop app user legacy permissions    A user operator has a set of permissions used in Dispatch desktop application.    The usage of these permissions may change in the future.      Permission Id | Description  ------ | -------  5 | Orders corrections  6 | Write scheduled transports  7 | View scheduled transports  8 | Print scheduled transports  10 | Orders consistency  11 | Statistics  12 | Margins  13 | Global settings  14 | Imports / Exports  15 | See Prices  16 | Invoices calculation  17 | Invoices editing  18 | Ad Hoc invoices  19 | Vouchers book  20 | Credit notes  21 | Invoices clearing  22 | View not issued bills  23 | Accounting export  24 | Payment  25 | Write units management  26 | View units  27 | Print units  28 | Invoices management  29 | Generate bills states  30 | Write customers  31 | View customers  32 | Print customers  33 | Write employees  34 | View employees  35 | Print employees  36 | Write subcontractors  37 | View subcontractors  38 | Print subcontractors  39 | Write salespersons  40 | View salespersons  41 | Print salespersons  42 | Write users  43 | View users  44 | Print users  45 | Write addresses  46 | View addresses  47 | Print addresses  48 | Write sectors  49 | View sectors  50 | Print sectors  51 | Write pricing paths  52 | View pricing paths  53 | Print pricing paths  54 | Write services  55 | View services  56 | Print services  57 | Write sub services  58 | View sub services  59 | Print sub services  60 | Write currencies  61 | View currencies  62 | Print currencies  63 | Write discount scale  64 | View discount scale  65 | Print discount scale  66 | Write payment methods  67 | View payment methods  68 | Print payment methods  69 | Write VAT  70 | View VAT  71 | Print VAT  72 | Write vehicles  73 | View vehicles  74 | Print vehicles  75 | Write cities  76 | View cities  77 | Print cities  78 | (Reserved)  79 | (Reserved)  80 | Companies  81 | (Reserved)  82 | (Reserved)  83 | Write orderer profiles  84 | View orderer profiles  92 | Customizable planning  93 | Tracking alerts  94 | Driver's message alerts  95 | Alerts on mobile dispute  96 | Alerts on unread / rejected missions  97 | Alerts on new web mission  98 | Alert on subcontractor's reminder  99 | Distribution entry  100 | Write references  101 | View references  102 | Print references  103 | Delegated orders tracking  104 | Alerts on mobile anomalies  105 | D2D Services equivalences  106 | D2D Services equivalences  107 | D2D amount change  108 | D2D City equivalence  109 | Customer / Subcontractor configuration  110 | D2D City equivalence  111 | Modify terminated orders  112 | Modify billed orders  113 | D2D Errors  114 | New D2D Mission / Modify D2D Mission  115 | Orders ending  116 | Orders cancellation  118 | Write user profiles  119 | View user profiles  120 | Geo device associations  121 | Cartography acces  122 | Create, Update, Delete  123 | Default model choice  124 | Preview  125 | Write agencies  126 | View agencies  127 | Print agencies  128 | Write orderers  129 | View orderers  130 | Print orderers  131 | Create, update, delete POI  132 | Create, update, delete pictogrammes  133 | Invoices clearing  134 | View tems  135 | Tools  136 | Time management administration  137 | View time management  138 | Subcontractors documents  139 | Write transport statuses  140 | Write transport substates  141 | View transport substates  142 | Write anomalies  143 | View anomalies  144 | Write disputes  145 | View disputes  146 | Write deposits  147 | View deposits  148 | Notepad  149 | Write unavailabilities  150 | View unavailabilities  151 | (Reserved)  152 | Consult a dysfunction sheet  153 | Modify a dysfunction sheet  154 | Print a dysfunction sheet  155 | Close misfunction sheet  156 | ReOpen a dysfunction sheet  157 | Write contractor agents  158 | View contractor agents  159 | Print contractor agents  160 | Kick users via instant messaging  161 | GPS devices configuration window  162 | Instant messaging  163 | Dispatch's alerts  164 | Notes  165 | Baracoda Tracking  166 | Set VAT in order  167 | Transport order manual price setting  168 | Write EDI parameters  169 | Global communications history  170 | Write regions  171 | View regions  172 | Write price modification rules  173 | View price modification rules  174 | Saving tabs configuration  175 | Allow users to see all agencies vehicles on Dispatch map  176 | Allow user to see recurrent orders from all agencies on map  177 | Desactivate free package type input  178 | Allow GPS Device privacy mode change  179 | Allow automatic working time declaration module execusion on current day  180 | Temperature alert  181 | Write package natures  182 | View package natures  183 | Print package natures  184 | Add package history state  185 | Write package reference values  186 | View package reference values  187 | Validate round sessions  188 | Override round session selection  189 | Override round session values  190 | Allow to secure transport  191 | allow batch modification of transport references  192 | Write subcontractor employees  193 | View subcontractor employees  194 | Print subcontractor employees  195 | Allow receipt entry  196 | Write appointments  197 | View appointments  198 | Access to planning appointment module  199 | Access to audit data  200 | Access to audit configuration  201 | Write password security policies  202 | View password security policies  203 | Transports audit  204 | Access to application orderers  205 | Write SSO configurations  206 | View SSO configurations  207 | Write SSO users  208 | View SSO users  209 | Application access management  210 | Display open link in browser button  211 | Data erasure and anonymization  212 | Write EDI equivalence configurations  213 | View EDI equivalence configurations  214 | Access to anonymized data

API version: v3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportDto{}

// TransportDto struct for TransportDto
type TransportDto struct {
	// Transport's internal unique identifier
	Uid *string `json:"uid,omitempty"`
	// Transport's internal unique identifier
	TransportId *int32 `json:"transportId,omitempty"`
	// The index of the transport in the mission or quotation
	Index *int32 `json:"index,omitempty"`
	// Mission's public identifier
	MissionNumber *int32 `json:"missionNumber,omitempty"`
	// Mission's internal unique identifier
	MissionUid *string `json:"missionUid,omitempty"`
	// Mission's tracking number
	MissionTrackId *string `json:"missionTrackId,omitempty"`
	// Quotation's public identifier
	QuotationNumber *int32 `json:"quotationNumber,omitempty"`
	// Quotation's internal unique identifier
	QuotationUid *string `json:"quotationUid,omitempty"`
	// Quotation's tracking number
	QuotationTrackId             *string `json:"quotationTrackId,omitempty"`
	IsQuotationArchived          *bool   `json:"isQuotationArchived,omitempty"`
	IsQuotationFinalized         *bool   `json:"isQuotationFinalized,omitempty"`
	IsQuotationSubjectToApproval *bool   `json:"isQuotationSubjectToApproval,omitempty"`
	IsMissionReadyToBill         *bool   `json:"isMissionReadyToBill,omitempty"`
	// Source quotation number, for a transport included in a mission created from a quotation.
	SourceQuotationNumber *int32  `json:"sourceQuotationNumber,omitempty"`
	RoundId               *int32  `json:"roundId,omitempty"`
	CustomerCode          *string `json:"customerCode,omitempty"`
	AgencyCode            *string `json:"agencyCode,omitempty"`
	OrdererCode           *string `json:"ordererCode,omitempty"`
	OrdererName           *string `json:"ordererName,omitempty"`
	ContractorAgentId     *int32  `json:"contractorAgentId,omitempty"`
	ContractorAgentName   *string `json:"contractorAgentName,omitempty"`
	HandlerId             *int32  `json:"handlerId,omitempty"`
	VehicleCode           *string `json:"vehicleCode,omitempty"`
	TrailerCode           *string `json:"trailerCode,omitempty"`
	Status                *string `json:"status,omitempty"`
	SubStateCode          *string `json:"subStateCode,omitempty"`
	// The date and time at which the customer has contacted the transport company to request the transport.  By default, it corresponds to the mission creation date.
	DriverId     *int32                                     `json:"driverId,omitempty"`
	Sign         *string                                    `json:"sign,omitempty"`
	SecretCode   *string                                    `json:"secretCode,omitempty"`
	Notes        *string                                    `json:"notes,omitempty"`
	IsRoundTrip  *bool                                      `json:"isRoundTrip,omitempty"`
	ServiceCode  *string                                    `json:"serviceCode,omitempty"`
	SubServices  *CappedCollectionDtoTransportSubServiceDto `json:"subServices,omitempty"`
	PickupStep   *TransportPickupStepDto                    `json:"pickupStep,omitempty"`
	DeliveryStep *TransportDeliveryStepDto                  `json:"deliveryStep,omitempty"`
	Packing      *TransportPackingDto                       `json:"packing,omitempty"`
	// Distance of the transport in kilometers.
	DistanceKm *float64 `json:"distanceKm,omitempty"`
	// Total duration of the transport in minutes
	TotalDurationMinutes *int32 `json:"totalDurationMinutes,omitempty"`
	// Indicates if the transport has a strategic importance
	IsStrategic *bool `json:"isStrategic,omitempty"`
	// Indicates if the transport is 'shared' between agencies.    If true, the transport will be visible on Dispatch planning   via the 'supervising' mode.
	IsShared *bool `json:"isShared,omitempty"`
	// The id of a dispatch user that is responsible for following the execution of the transport.
	ResponsibleOperatorId *int32 `json:"responsibleOperatorId,omitempty"`
	// Gas emission (CO2) in Kilograms.
	GasEmission         *float64 `json:"gasEmission,omitempty"`
	IsGasEmissionForced *bool    `json:"isGasEmissionForced,omitempty"`
	SellCurrencyCode    *string  `json:"sellCurrencyCode,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	SellPrice *float64 `json:"sellPrice,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	ForcedSellPrice *float64 `json:"forcedSellPrice,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	SellFuelSurchargePrice *float64 `json:"sellFuelSurchargePrice,omitempty"`
	BuyCurrencyCode        *string  `json:"buyCurrencyCode,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	BuyPrice *float64 `json:"buyPrice,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	ForcedBuyPrice *float64 `json:"forcedBuyPrice,omitempty"`
	// The following permission(s) are required to access this property:  See prices.
	BuyFuelSurchargePrice *float64 `json:"buyFuelSurchargePrice,omitempty"`
	ReferenceCurrencyCode *string  `json:"referenceCurrencyCode,omitempty"`
	SellPricingPathId     *int32   `json:"sellPricingPathId,omitempty"`
	// Indicates if this transport is fragmented:  in this case, this transport is a parent transport and has child fragments.
	IsFragmented *bool `json:"isFragmented,omitempty"`
	// Identifier of the parent transport for a child fragment,  when the mission or quotation is fragmented
	ParentTransportId      *int32                      `json:"parentTransportId,omitempty"`
	MissionIsInSequence    *bool                       `json:"missionIsInSequence,omitempty"`
	MissionIsUniqueOrder   *bool                       `json:"missionIsUniqueOrder,omitempty"`
	QuotationIsInSequence  *bool                       `json:"quotationIsInSequence,omitempty"`
	QuotationIsUniqueOrder *bool                       `json:"quotationIsUniqueOrder,omitempty"`
	Reference1             *string                     `json:"reference1,omitempty"`
	Reference2             *string                     `json:"reference2,omitempty"`
	Reference3             *string                     `json:"reference3,omitempty"`
	Comment                *TransportCommentDto        `json:"comment,omitempty"`
	CashOnDelivery         *TransportCashOnDeliveryDto `json:"cashOnDelivery,omitempty"`
	Included               *TransportIncludedDto       `json:"included,omitempty"`
}

// NewTransportDto instantiates a new TransportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportDto() *TransportDto {
	this := TransportDto{}
	return &this
}

// NewTransportDtoWithDefaults instantiates a new TransportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportDtoWithDefaults() *TransportDto {
	this := TransportDto{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *TransportDto) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *TransportDto) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *TransportDto) SetUid(v string) {
	o.Uid = &v
}

// GetTransportId returns the TransportId field value if set, zero value otherwise.
func (o *TransportDto) GetTransportId() int32 {
	if o == nil || IsNil(o.TransportId) {
		var ret int32
		return ret
	}
	return *o.TransportId
}

// GetTransportIdOk returns a tuple with the TransportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetTransportIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TransportId) {
		return nil, false
	}
	return o.TransportId, true
}

// HasTransportId returns a boolean if a field has been set.
func (o *TransportDto) HasTransportId() bool {
	if o != nil && !IsNil(o.TransportId) {
		return true
	}

	return false
}

// SetTransportId gets a reference to the given int32 and assigns it to the TransportId field.
func (o *TransportDto) SetTransportId(v int32) {
	o.TransportId = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TransportDto) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TransportDto) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TransportDto) SetIndex(v int32) {
	o.Index = &v
}

// GetMissionNumber returns the MissionNumber field value if set, zero value otherwise.
func (o *TransportDto) GetMissionNumber() int32 {
	if o == nil || IsNil(o.MissionNumber) {
		var ret int32
		return ret
	}
	return *o.MissionNumber
}

// GetMissionNumberOk returns a tuple with the MissionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetMissionNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.MissionNumber) {
		return nil, false
	}
	return o.MissionNumber, true
}

// HasMissionNumber returns a boolean if a field has been set.
func (o *TransportDto) HasMissionNumber() bool {
	if o != nil && !IsNil(o.MissionNumber) {
		return true
	}

	return false
}

// SetMissionNumber gets a reference to the given int32 and assigns it to the MissionNumber field.
func (o *TransportDto) SetMissionNumber(v int32) {
	o.MissionNumber = &v
}

// GetMissionUid returns the MissionUid field value if set, zero value otherwise.
func (o *TransportDto) GetMissionUid() string {
	if o == nil || IsNil(o.MissionUid) {
		var ret string
		return ret
	}
	return *o.MissionUid
}

// GetMissionUidOk returns a tuple with the MissionUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetMissionUidOk() (*string, bool) {
	if o == nil || IsNil(o.MissionUid) {
		return nil, false
	}
	return o.MissionUid, true
}

// HasMissionUid returns a boolean if a field has been set.
func (o *TransportDto) HasMissionUid() bool {
	if o != nil && !IsNil(o.MissionUid) {
		return true
	}

	return false
}

// SetMissionUid gets a reference to the given string and assigns it to the MissionUid field.
func (o *TransportDto) SetMissionUid(v string) {
	o.MissionUid = &v
}

// GetMissionTrackId returns the MissionTrackId field value if set, zero value otherwise.
func (o *TransportDto) GetMissionTrackId() string {
	if o == nil || IsNil(o.MissionTrackId) {
		var ret string
		return ret
	}
	return *o.MissionTrackId
}

// GetMissionTrackIdOk returns a tuple with the MissionTrackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetMissionTrackIdOk() (*string, bool) {
	if o == nil || IsNil(o.MissionTrackId) {
		return nil, false
	}
	return o.MissionTrackId, true
}

// HasMissionTrackId returns a boolean if a field has been set.
func (o *TransportDto) HasMissionTrackId() bool {
	if o != nil && !IsNil(o.MissionTrackId) {
		return true
	}

	return false
}

// SetMissionTrackId gets a reference to the given string and assigns it to the MissionTrackId field.
func (o *TransportDto) SetMissionTrackId(v string) {
	o.MissionTrackId = &v
}

// GetQuotationNumber returns the QuotationNumber field value if set, zero value otherwise.
func (o *TransportDto) GetQuotationNumber() int32 {
	if o == nil || IsNil(o.QuotationNumber) {
		var ret int32
		return ret
	}
	return *o.QuotationNumber
}

// GetQuotationNumberOk returns a tuple with the QuotationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetQuotationNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotationNumber) {
		return nil, false
	}
	return o.QuotationNumber, true
}

// HasQuotationNumber returns a boolean if a field has been set.
func (o *TransportDto) HasQuotationNumber() bool {
	if o != nil && !IsNil(o.QuotationNumber) {
		return true
	}

	return false
}

// SetQuotationNumber gets a reference to the given int32 and assigns it to the QuotationNumber field.
func (o *TransportDto) SetQuotationNumber(v int32) {
	o.QuotationNumber = &v
}

// GetQuotationUid returns the QuotationUid field value if set, zero value otherwise.
func (o *TransportDto) GetQuotationUid() string {
	if o == nil || IsNil(o.QuotationUid) {
		var ret string
		return ret
	}
	return *o.QuotationUid
}

// GetQuotationUidOk returns a tuple with the QuotationUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetQuotationUidOk() (*string, bool) {
	if o == nil || IsNil(o.QuotationUid) {
		return nil, false
	}
	return o.QuotationUid, true
}

// HasQuotationUid returns a boolean if a field has been set.
func (o *TransportDto) HasQuotationUid() bool {
	if o != nil && !IsNil(o.QuotationUid) {
		return true
	}

	return false
}

// SetQuotationUid gets a reference to the given string and assigns it to the QuotationUid field.
func (o *TransportDto) SetQuotationUid(v string) {
	o.QuotationUid = &v
}

// GetQuotationTrackId returns the QuotationTrackId field value if set, zero value otherwise.
func (o *TransportDto) GetQuotationTrackId() string {
	if o == nil || IsNil(o.QuotationTrackId) {
		var ret string
		return ret
	}
	return *o.QuotationTrackId
}

// GetQuotationTrackIdOk returns a tuple with the QuotationTrackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetQuotationTrackIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuotationTrackId) {
		return nil, false
	}
	return o.QuotationTrackId, true
}

// HasQuotationTrackId returns a boolean if a field has been set.
func (o *TransportDto) HasQuotationTrackId() bool {
	if o != nil && !IsNil(o.QuotationTrackId) {
		return true
	}

	return false
}

// SetQuotationTrackId gets a reference to the given string and assigns it to the QuotationTrackId field.
func (o *TransportDto) SetQuotationTrackId(v string) {
	o.QuotationTrackId = &v
}

// GetIsQuotationArchived returns the IsQuotationArchived field value if set, zero value otherwise.
func (o *TransportDto) GetIsQuotationArchived() bool {
	if o == nil || IsNil(o.IsQuotationArchived) {
		var ret bool
		return ret
	}
	return *o.IsQuotationArchived
}

// GetIsQuotationArchivedOk returns a tuple with the IsQuotationArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsQuotationArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuotationArchived) {
		return nil, false
	}
	return o.IsQuotationArchived, true
}

// HasIsQuotationArchived returns a boolean if a field has been set.
func (o *TransportDto) HasIsQuotationArchived() bool {
	if o != nil && !IsNil(o.IsQuotationArchived) {
		return true
	}

	return false
}

// SetIsQuotationArchived gets a reference to the given bool and assigns it to the IsQuotationArchived field.
func (o *TransportDto) SetIsQuotationArchived(v bool) {
	o.IsQuotationArchived = &v
}

// GetIsQuotationFinalized returns the IsQuotationFinalized field value if set, zero value otherwise.
func (o *TransportDto) GetIsQuotationFinalized() bool {
	if o == nil || IsNil(o.IsQuotationFinalized) {
		var ret bool
		return ret
	}
	return *o.IsQuotationFinalized
}

// GetIsQuotationFinalizedOk returns a tuple with the IsQuotationFinalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsQuotationFinalizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuotationFinalized) {
		return nil, false
	}
	return o.IsQuotationFinalized, true
}

// HasIsQuotationFinalized returns a boolean if a field has been set.
func (o *TransportDto) HasIsQuotationFinalized() bool {
	if o != nil && !IsNil(o.IsQuotationFinalized) {
		return true
	}

	return false
}

// SetIsQuotationFinalized gets a reference to the given bool and assigns it to the IsQuotationFinalized field.
func (o *TransportDto) SetIsQuotationFinalized(v bool) {
	o.IsQuotationFinalized = &v
}

// GetIsQuotationSubjectToApproval returns the IsQuotationSubjectToApproval field value if set, zero value otherwise.
func (o *TransportDto) GetIsQuotationSubjectToApproval() bool {
	if o == nil || IsNil(o.IsQuotationSubjectToApproval) {
		var ret bool
		return ret
	}
	return *o.IsQuotationSubjectToApproval
}

// GetIsQuotationSubjectToApprovalOk returns a tuple with the IsQuotationSubjectToApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsQuotationSubjectToApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQuotationSubjectToApproval) {
		return nil, false
	}
	return o.IsQuotationSubjectToApproval, true
}

// HasIsQuotationSubjectToApproval returns a boolean if a field has been set.
func (o *TransportDto) HasIsQuotationSubjectToApproval() bool {
	if o != nil && !IsNil(o.IsQuotationSubjectToApproval) {
		return true
	}

	return false
}

// SetIsQuotationSubjectToApproval gets a reference to the given bool and assigns it to the IsQuotationSubjectToApproval field.
func (o *TransportDto) SetIsQuotationSubjectToApproval(v bool) {
	o.IsQuotationSubjectToApproval = &v
}

// GetIsMissionReadyToBill returns the IsMissionReadyToBill field value if set, zero value otherwise.
func (o *TransportDto) GetIsMissionReadyToBill() bool {
	if o == nil || IsNil(o.IsMissionReadyToBill) {
		var ret bool
		return ret
	}
	return *o.IsMissionReadyToBill
}

// GetIsMissionReadyToBillOk returns a tuple with the IsMissionReadyToBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsMissionReadyToBillOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMissionReadyToBill) {
		return nil, false
	}
	return o.IsMissionReadyToBill, true
}

// HasIsMissionReadyToBill returns a boolean if a field has been set.
func (o *TransportDto) HasIsMissionReadyToBill() bool {
	if o != nil && !IsNil(o.IsMissionReadyToBill) {
		return true
	}

	return false
}

// SetIsMissionReadyToBill gets a reference to the given bool and assigns it to the IsMissionReadyToBill field.
func (o *TransportDto) SetIsMissionReadyToBill(v bool) {
	o.IsMissionReadyToBill = &v
}

// GetSourceQuotationNumber returns the SourceQuotationNumber field value if set, zero value otherwise.
func (o *TransportDto) GetSourceQuotationNumber() int32 {
	if o == nil || IsNil(o.SourceQuotationNumber) {
		var ret int32
		return ret
	}
	return *o.SourceQuotationNumber
}

// GetSourceQuotationNumberOk returns a tuple with the SourceQuotationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSourceQuotationNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceQuotationNumber) {
		return nil, false
	}
	return o.SourceQuotationNumber, true
}

// HasSourceQuotationNumber returns a boolean if a field has been set.
func (o *TransportDto) HasSourceQuotationNumber() bool {
	if o != nil && !IsNil(o.SourceQuotationNumber) {
		return true
	}

	return false
}

// SetSourceQuotationNumber gets a reference to the given int32 and assigns it to the SourceQuotationNumber field.
func (o *TransportDto) SetSourceQuotationNumber(v int32) {
	o.SourceQuotationNumber = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *TransportDto) GetRoundId() int32 {
	if o == nil || IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetRoundIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *TransportDto) HasRoundId() bool {
	if o != nil && !IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *TransportDto) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetCustomerCode returns the CustomerCode field value if set, zero value otherwise.
func (o *TransportDto) GetCustomerCode() string {
	if o == nil || IsNil(o.CustomerCode) {
		var ret string
		return ret
	}
	return *o.CustomerCode
}

// GetCustomerCodeOk returns a tuple with the CustomerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetCustomerCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCode) {
		return nil, false
	}
	return o.CustomerCode, true
}

// HasCustomerCode returns a boolean if a field has been set.
func (o *TransportDto) HasCustomerCode() bool {
	if o != nil && !IsNil(o.CustomerCode) {
		return true
	}

	return false
}

// SetCustomerCode gets a reference to the given string and assigns it to the CustomerCode field.
func (o *TransportDto) SetCustomerCode(v string) {
	o.CustomerCode = &v
}

// GetAgencyCode returns the AgencyCode field value if set, zero value otherwise.
func (o *TransportDto) GetAgencyCode() string {
	if o == nil || IsNil(o.AgencyCode) {
		var ret string
		return ret
	}
	return *o.AgencyCode
}

// GetAgencyCodeOk returns a tuple with the AgencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetAgencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AgencyCode) {
		return nil, false
	}
	return o.AgencyCode, true
}

// HasAgencyCode returns a boolean if a field has been set.
func (o *TransportDto) HasAgencyCode() bool {
	if o != nil && !IsNil(o.AgencyCode) {
		return true
	}

	return false
}

// SetAgencyCode gets a reference to the given string and assigns it to the AgencyCode field.
func (o *TransportDto) SetAgencyCode(v string) {
	o.AgencyCode = &v
}

// GetOrdererCode returns the OrdererCode field value if set, zero value otherwise.
func (o *TransportDto) GetOrdererCode() string {
	if o == nil || IsNil(o.OrdererCode) {
		var ret string
		return ret
	}
	return *o.OrdererCode
}

// GetOrdererCodeOk returns a tuple with the OrdererCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetOrdererCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OrdererCode) {
		return nil, false
	}
	return o.OrdererCode, true
}

// HasOrdererCode returns a boolean if a field has been set.
func (o *TransportDto) HasOrdererCode() bool {
	if o != nil && !IsNil(o.OrdererCode) {
		return true
	}

	return false
}

// SetOrdererCode gets a reference to the given string and assigns it to the OrdererCode field.
func (o *TransportDto) SetOrdererCode(v string) {
	o.OrdererCode = &v
}

// GetOrdererName returns the OrdererName field value if set, zero value otherwise.
func (o *TransportDto) GetOrdererName() string {
	if o == nil || IsNil(o.OrdererName) {
		var ret string
		return ret
	}
	return *o.OrdererName
}

// GetOrdererNameOk returns a tuple with the OrdererName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetOrdererNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrdererName) {
		return nil, false
	}
	return o.OrdererName, true
}

// HasOrdererName returns a boolean if a field has been set.
func (o *TransportDto) HasOrdererName() bool {
	if o != nil && !IsNil(o.OrdererName) {
		return true
	}

	return false
}

// SetOrdererName gets a reference to the given string and assigns it to the OrdererName field.
func (o *TransportDto) SetOrdererName(v string) {
	o.OrdererName = &v
}

// GetContractorAgentId returns the ContractorAgentId field value if set, zero value otherwise.
func (o *TransportDto) GetContractorAgentId() int32 {
	if o == nil || IsNil(o.ContractorAgentId) {
		var ret int32
		return ret
	}
	return *o.ContractorAgentId
}

// GetContractorAgentIdOk returns a tuple with the ContractorAgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetContractorAgentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractorAgentId) {
		return nil, false
	}
	return o.ContractorAgentId, true
}

// HasContractorAgentId returns a boolean if a field has been set.
func (o *TransportDto) HasContractorAgentId() bool {
	if o != nil && !IsNil(o.ContractorAgentId) {
		return true
	}

	return false
}

// SetContractorAgentId gets a reference to the given int32 and assigns it to the ContractorAgentId field.
func (o *TransportDto) SetContractorAgentId(v int32) {
	o.ContractorAgentId = &v
}

// GetContractorAgentName returns the ContractorAgentName field value if set, zero value otherwise.
func (o *TransportDto) GetContractorAgentName() string {
	if o == nil || IsNil(o.ContractorAgentName) {
		var ret string
		return ret
	}
	return *o.ContractorAgentName
}

// GetContractorAgentNameOk returns a tuple with the ContractorAgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetContractorAgentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContractorAgentName) {
		return nil, false
	}
	return o.ContractorAgentName, true
}

// HasContractorAgentName returns a boolean if a field has been set.
func (o *TransportDto) HasContractorAgentName() bool {
	if o != nil && !IsNil(o.ContractorAgentName) {
		return true
	}

	return false
}

// SetContractorAgentName gets a reference to the given string and assigns it to the ContractorAgentName field.
func (o *TransportDto) SetContractorAgentName(v string) {
	o.ContractorAgentName = &v
}

// GetHandlerId returns the HandlerId field value if set, zero value otherwise.
func (o *TransportDto) GetHandlerId() int32 {
	if o == nil || IsNil(o.HandlerId) {
		var ret int32
		return ret
	}
	return *o.HandlerId
}

// GetHandlerIdOk returns a tuple with the HandlerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetHandlerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.HandlerId) {
		return nil, false
	}
	return o.HandlerId, true
}

// HasHandlerId returns a boolean if a field has been set.
func (o *TransportDto) HasHandlerId() bool {
	if o != nil && !IsNil(o.HandlerId) {
		return true
	}

	return false
}

// SetHandlerId gets a reference to the given int32 and assigns it to the HandlerId field.
func (o *TransportDto) SetHandlerId(v int32) {
	o.HandlerId = &v
}

// GetVehicleCode returns the VehicleCode field value if set, zero value otherwise.
func (o *TransportDto) GetVehicleCode() string {
	if o == nil || IsNil(o.VehicleCode) {
		var ret string
		return ret
	}
	return *o.VehicleCode
}

// GetVehicleCodeOk returns a tuple with the VehicleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetVehicleCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VehicleCode) {
		return nil, false
	}
	return o.VehicleCode, true
}

// HasVehicleCode returns a boolean if a field has been set.
func (o *TransportDto) HasVehicleCode() bool {
	if o != nil && !IsNil(o.VehicleCode) {
		return true
	}

	return false
}

// SetVehicleCode gets a reference to the given string and assigns it to the VehicleCode field.
func (o *TransportDto) SetVehicleCode(v string) {
	o.VehicleCode = &v
}

// GetTrailerCode returns the TrailerCode field value if set, zero value otherwise.
func (o *TransportDto) GetTrailerCode() string {
	if o == nil || IsNil(o.TrailerCode) {
		var ret string
		return ret
	}
	return *o.TrailerCode
}

// GetTrailerCodeOk returns a tuple with the TrailerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetTrailerCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TrailerCode) {
		return nil, false
	}
	return o.TrailerCode, true
}

// HasTrailerCode returns a boolean if a field has been set.
func (o *TransportDto) HasTrailerCode() bool {
	if o != nil && !IsNil(o.TrailerCode) {
		return true
	}

	return false
}

// SetTrailerCode gets a reference to the given string and assigns it to the TrailerCode field.
func (o *TransportDto) SetTrailerCode(v string) {
	o.TrailerCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransportDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransportDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransportDto) SetStatus(v string) {
	o.Status = &v
}

// GetSubStateCode returns the SubStateCode field value if set, zero value otherwise.
func (o *TransportDto) GetSubStateCode() string {
	if o == nil || IsNil(o.SubStateCode) {
		var ret string
		return ret
	}
	return *o.SubStateCode
}

// GetSubStateCodeOk returns a tuple with the SubStateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSubStateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SubStateCode) {
		return nil, false
	}
	return o.SubStateCode, true
}

// HasSubStateCode returns a boolean if a field has been set.
func (o *TransportDto) HasSubStateCode() bool {
	if o != nil && !IsNil(o.SubStateCode) {
		return true
	}

	return false
}

// SetSubStateCode gets a reference to the given string and assigns it to the SubStateCode field.
func (o *TransportDto) SetSubStateCode(v string) {
	o.SubStateCode = &v
}

// GetDriverId returns the DriverId field value if set, zero value otherwise.
func (o *TransportDto) GetDriverId() int32 {
	if o == nil || IsNil(o.DriverId) {
		var ret int32
		return ret
	}
	return *o.DriverId
}

// GetDriverIdOk returns a tuple with the DriverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetDriverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DriverId) {
		return nil, false
	}
	return o.DriverId, true
}

// HasDriverId returns a boolean if a field has been set.
func (o *TransportDto) HasDriverId() bool {
	if o != nil && !IsNil(o.DriverId) {
		return true
	}

	return false
}

// SetDriverId gets a reference to the given int32 and assigns it to the DriverId field.
func (o *TransportDto) SetDriverId(v int32) {
	o.DriverId = &v
}

// GetSign returns the Sign field value if set, zero value otherwise.
func (o *TransportDto) GetSign() string {
	if o == nil || IsNil(o.Sign) {
		var ret string
		return ret
	}
	return *o.Sign
}

// GetSignOk returns a tuple with the Sign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSignOk() (*string, bool) {
	if o == nil || IsNil(o.Sign) {
		return nil, false
	}
	return o.Sign, true
}

// HasSign returns a boolean if a field has been set.
func (o *TransportDto) HasSign() bool {
	if o != nil && !IsNil(o.Sign) {
		return true
	}

	return false
}

// SetSign gets a reference to the given string and assigns it to the Sign field.
func (o *TransportDto) SetSign(v string) {
	o.Sign = &v
}

// GetSecretCode returns the SecretCode field value if set, zero value otherwise.
func (o *TransportDto) GetSecretCode() string {
	if o == nil || IsNil(o.SecretCode) {
		var ret string
		return ret
	}
	return *o.SecretCode
}

// GetSecretCodeOk returns a tuple with the SecretCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSecretCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SecretCode) {
		return nil, false
	}
	return o.SecretCode, true
}

// HasSecretCode returns a boolean if a field has been set.
func (o *TransportDto) HasSecretCode() bool {
	if o != nil && !IsNil(o.SecretCode) {
		return true
	}

	return false
}

// SetSecretCode gets a reference to the given string and assigns it to the SecretCode field.
func (o *TransportDto) SetSecretCode(v string) {
	o.SecretCode = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TransportDto) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TransportDto) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *TransportDto) SetNotes(v string) {
	o.Notes = &v
}

// GetIsRoundTrip returns the IsRoundTrip field value if set, zero value otherwise.
func (o *TransportDto) GetIsRoundTrip() bool {
	if o == nil || IsNil(o.IsRoundTrip) {
		var ret bool
		return ret
	}
	return *o.IsRoundTrip
}

// GetIsRoundTripOk returns a tuple with the IsRoundTrip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsRoundTripOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoundTrip) {
		return nil, false
	}
	return o.IsRoundTrip, true
}

// HasIsRoundTrip returns a boolean if a field has been set.
func (o *TransportDto) HasIsRoundTrip() bool {
	if o != nil && !IsNil(o.IsRoundTrip) {
		return true
	}

	return false
}

// SetIsRoundTrip gets a reference to the given bool and assigns it to the IsRoundTrip field.
func (o *TransportDto) SetIsRoundTrip(v bool) {
	o.IsRoundTrip = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *TransportDto) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *TransportDto) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *TransportDto) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetSubServices returns the SubServices field value if set, zero value otherwise.
func (o *TransportDto) GetSubServices() CappedCollectionDtoTransportSubServiceDto {
	if o == nil || IsNil(o.SubServices) {
		var ret CappedCollectionDtoTransportSubServiceDto
		return ret
	}
	return *o.SubServices
}

// GetSubServicesOk returns a tuple with the SubServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSubServicesOk() (*CappedCollectionDtoTransportSubServiceDto, bool) {
	if o == nil || IsNil(o.SubServices) {
		return nil, false
	}
	return o.SubServices, true
}

// HasSubServices returns a boolean if a field has been set.
func (o *TransportDto) HasSubServices() bool {
	if o != nil && !IsNil(o.SubServices) {
		return true
	}

	return false
}

// SetSubServices gets a reference to the given CappedCollectionDtoTransportSubServiceDto and assigns it to the SubServices field.
func (o *TransportDto) SetSubServices(v CappedCollectionDtoTransportSubServiceDto) {
	o.SubServices = &v
}

// GetPickupStep returns the PickupStep field value if set, zero value otherwise.
func (o *TransportDto) GetPickupStep() TransportPickupStepDto {
	if o == nil || IsNil(o.PickupStep) {
		var ret TransportPickupStepDto
		return ret
	}
	return *o.PickupStep
}

// GetPickupStepOk returns a tuple with the PickupStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetPickupStepOk() (*TransportPickupStepDto, bool) {
	if o == nil || IsNil(o.PickupStep) {
		return nil, false
	}
	return o.PickupStep, true
}

// HasPickupStep returns a boolean if a field has been set.
func (o *TransportDto) HasPickupStep() bool {
	if o != nil && !IsNil(o.PickupStep) {
		return true
	}

	return false
}

// SetPickupStep gets a reference to the given TransportPickupStepDto and assigns it to the PickupStep field.
func (o *TransportDto) SetPickupStep(v TransportPickupStepDto) {
	o.PickupStep = &v
}

// GetDeliveryStep returns the DeliveryStep field value if set, zero value otherwise.
func (o *TransportDto) GetDeliveryStep() TransportDeliveryStepDto {
	if o == nil || IsNil(o.DeliveryStep) {
		var ret TransportDeliveryStepDto
		return ret
	}
	return *o.DeliveryStep
}

// GetDeliveryStepOk returns a tuple with the DeliveryStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetDeliveryStepOk() (*TransportDeliveryStepDto, bool) {
	if o == nil || IsNil(o.DeliveryStep) {
		return nil, false
	}
	return o.DeliveryStep, true
}

// HasDeliveryStep returns a boolean if a field has been set.
func (o *TransportDto) HasDeliveryStep() bool {
	if o != nil && !IsNil(o.DeliveryStep) {
		return true
	}

	return false
}

// SetDeliveryStep gets a reference to the given TransportDeliveryStepDto and assigns it to the DeliveryStep field.
func (o *TransportDto) SetDeliveryStep(v TransportDeliveryStepDto) {
	o.DeliveryStep = &v
}

// GetPacking returns the Packing field value if set, zero value otherwise.
func (o *TransportDto) GetPacking() TransportPackingDto {
	if o == nil || IsNil(o.Packing) {
		var ret TransportPackingDto
		return ret
	}
	return *o.Packing
}

// GetPackingOk returns a tuple with the Packing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetPackingOk() (*TransportPackingDto, bool) {
	if o == nil || IsNil(o.Packing) {
		return nil, false
	}
	return o.Packing, true
}

// HasPacking returns a boolean if a field has been set.
func (o *TransportDto) HasPacking() bool {
	if o != nil && !IsNil(o.Packing) {
		return true
	}

	return false
}

// SetPacking gets a reference to the given TransportPackingDto and assigns it to the Packing field.
func (o *TransportDto) SetPacking(v TransportPackingDto) {
	o.Packing = &v
}

// GetDistanceKm returns the DistanceKm field value if set, zero value otherwise.
func (o *TransportDto) GetDistanceKm() float64 {
	if o == nil || IsNil(o.DistanceKm) {
		var ret float64
		return ret
	}
	return *o.DistanceKm
}

// GetDistanceKmOk returns a tuple with the DistanceKm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetDistanceKmOk() (*float64, bool) {
	if o == nil || IsNil(o.DistanceKm) {
		return nil, false
	}
	return o.DistanceKm, true
}

// HasDistanceKm returns a boolean if a field has been set.
func (o *TransportDto) HasDistanceKm() bool {
	if o != nil && !IsNil(o.DistanceKm) {
		return true
	}

	return false
}

// SetDistanceKm gets a reference to the given float64 and assigns it to the DistanceKm field.
func (o *TransportDto) SetDistanceKm(v float64) {
	o.DistanceKm = &v
}

// GetTotalDurationMinutes returns the TotalDurationMinutes field value if set, zero value otherwise.
func (o *TransportDto) GetTotalDurationMinutes() int32 {
	if o == nil || IsNil(o.TotalDurationMinutes) {
		var ret int32
		return ret
	}
	return *o.TotalDurationMinutes
}

// GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetTotalDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalDurationMinutes) {
		return nil, false
	}
	return o.TotalDurationMinutes, true
}

// HasTotalDurationMinutes returns a boolean if a field has been set.
func (o *TransportDto) HasTotalDurationMinutes() bool {
	if o != nil && !IsNil(o.TotalDurationMinutes) {
		return true
	}

	return false
}

// SetTotalDurationMinutes gets a reference to the given int32 and assigns it to the TotalDurationMinutes field.
func (o *TransportDto) SetTotalDurationMinutes(v int32) {
	o.TotalDurationMinutes = &v
}

// GetIsStrategic returns the IsStrategic field value if set, zero value otherwise.
func (o *TransportDto) GetIsStrategic() bool {
	if o == nil || IsNil(o.IsStrategic) {
		var ret bool
		return ret
	}
	return *o.IsStrategic
}

// GetIsStrategicOk returns a tuple with the IsStrategic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsStrategicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStrategic) {
		return nil, false
	}
	return o.IsStrategic, true
}

// HasIsStrategic returns a boolean if a field has been set.
func (o *TransportDto) HasIsStrategic() bool {
	if o != nil && !IsNil(o.IsStrategic) {
		return true
	}

	return false
}

// SetIsStrategic gets a reference to the given bool and assigns it to the IsStrategic field.
func (o *TransportDto) SetIsStrategic(v bool) {
	o.IsStrategic = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *TransportDto) GetIsShared() bool {
	if o == nil || IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *TransportDto) HasIsShared() bool {
	if o != nil && !IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *TransportDto) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetResponsibleOperatorId returns the ResponsibleOperatorId field value if set, zero value otherwise.
func (o *TransportDto) GetResponsibleOperatorId() int32 {
	if o == nil || IsNil(o.ResponsibleOperatorId) {
		var ret int32
		return ret
	}
	return *o.ResponsibleOperatorId
}

// GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetResponsibleOperatorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponsibleOperatorId) {
		return nil, false
	}
	return o.ResponsibleOperatorId, true
}

// HasResponsibleOperatorId returns a boolean if a field has been set.
func (o *TransportDto) HasResponsibleOperatorId() bool {
	if o != nil && !IsNil(o.ResponsibleOperatorId) {
		return true
	}

	return false
}

// SetResponsibleOperatorId gets a reference to the given int32 and assigns it to the ResponsibleOperatorId field.
func (o *TransportDto) SetResponsibleOperatorId(v int32) {
	o.ResponsibleOperatorId = &v
}

// GetGasEmission returns the GasEmission field value if set, zero value otherwise.
func (o *TransportDto) GetGasEmission() float64 {
	if o == nil || IsNil(o.GasEmission) {
		var ret float64
		return ret
	}
	return *o.GasEmission
}

// GetGasEmissionOk returns a tuple with the GasEmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetGasEmissionOk() (*float64, bool) {
	if o == nil || IsNil(o.GasEmission) {
		return nil, false
	}
	return o.GasEmission, true
}

// HasGasEmission returns a boolean if a field has been set.
func (o *TransportDto) HasGasEmission() bool {
	if o != nil && !IsNil(o.GasEmission) {
		return true
	}

	return false
}

// SetGasEmission gets a reference to the given float64 and assigns it to the GasEmission field.
func (o *TransportDto) SetGasEmission(v float64) {
	o.GasEmission = &v
}

// GetIsGasEmissionForced returns the IsGasEmissionForced field value if set, zero value otherwise.
func (o *TransportDto) GetIsGasEmissionForced() bool {
	if o == nil || IsNil(o.IsGasEmissionForced) {
		var ret bool
		return ret
	}
	return *o.IsGasEmissionForced
}

// GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsGasEmissionForcedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGasEmissionForced) {
		return nil, false
	}
	return o.IsGasEmissionForced, true
}

// HasIsGasEmissionForced returns a boolean if a field has been set.
func (o *TransportDto) HasIsGasEmissionForced() bool {
	if o != nil && !IsNil(o.IsGasEmissionForced) {
		return true
	}

	return false
}

// SetIsGasEmissionForced gets a reference to the given bool and assigns it to the IsGasEmissionForced field.
func (o *TransportDto) SetIsGasEmissionForced(v bool) {
	o.IsGasEmissionForced = &v
}

// GetSellCurrencyCode returns the SellCurrencyCode field value if set, zero value otherwise.
func (o *TransportDto) GetSellCurrencyCode() string {
	if o == nil || IsNil(o.SellCurrencyCode) {
		var ret string
		return ret
	}
	return *o.SellCurrencyCode
}

// GetSellCurrencyCodeOk returns a tuple with the SellCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSellCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SellCurrencyCode) {
		return nil, false
	}
	return o.SellCurrencyCode, true
}

// HasSellCurrencyCode returns a boolean if a field has been set.
func (o *TransportDto) HasSellCurrencyCode() bool {
	if o != nil && !IsNil(o.SellCurrencyCode) {
		return true
	}

	return false
}

// SetSellCurrencyCode gets a reference to the given string and assigns it to the SellCurrencyCode field.
func (o *TransportDto) SetSellCurrencyCode(v string) {
	o.SellCurrencyCode = &v
}

// GetSellPrice returns the SellPrice field value if set, zero value otherwise.
func (o *TransportDto) GetSellPrice() float64 {
	if o == nil || IsNil(o.SellPrice) {
		var ret float64
		return ret
	}
	return *o.SellPrice
}

// GetSellPriceOk returns a tuple with the SellPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSellPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.SellPrice) {
		return nil, false
	}
	return o.SellPrice, true
}

// HasSellPrice returns a boolean if a field has been set.
func (o *TransportDto) HasSellPrice() bool {
	if o != nil && !IsNil(o.SellPrice) {
		return true
	}

	return false
}

// SetSellPrice gets a reference to the given float64 and assigns it to the SellPrice field.
func (o *TransportDto) SetSellPrice(v float64) {
	o.SellPrice = &v
}

// GetForcedSellPrice returns the ForcedSellPrice field value if set, zero value otherwise.
func (o *TransportDto) GetForcedSellPrice() float64 {
	if o == nil || IsNil(o.ForcedSellPrice) {
		var ret float64
		return ret
	}
	return *o.ForcedSellPrice
}

// GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetForcedSellPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedSellPrice) {
		return nil, false
	}
	return o.ForcedSellPrice, true
}

// HasForcedSellPrice returns a boolean if a field has been set.
func (o *TransportDto) HasForcedSellPrice() bool {
	if o != nil && !IsNil(o.ForcedSellPrice) {
		return true
	}

	return false
}

// SetForcedSellPrice gets a reference to the given float64 and assigns it to the ForcedSellPrice field.
func (o *TransportDto) SetForcedSellPrice(v float64) {
	o.ForcedSellPrice = &v
}

// GetSellFuelSurchargePrice returns the SellFuelSurchargePrice field value if set, zero value otherwise.
func (o *TransportDto) GetSellFuelSurchargePrice() float64 {
	if o == nil || IsNil(o.SellFuelSurchargePrice) {
		var ret float64
		return ret
	}
	return *o.SellFuelSurchargePrice
}

// GetSellFuelSurchargePriceOk returns a tuple with the SellFuelSurchargePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSellFuelSurchargePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.SellFuelSurchargePrice) {
		return nil, false
	}
	return o.SellFuelSurchargePrice, true
}

// HasSellFuelSurchargePrice returns a boolean if a field has been set.
func (o *TransportDto) HasSellFuelSurchargePrice() bool {
	if o != nil && !IsNil(o.SellFuelSurchargePrice) {
		return true
	}

	return false
}

// SetSellFuelSurchargePrice gets a reference to the given float64 and assigns it to the SellFuelSurchargePrice field.
func (o *TransportDto) SetSellFuelSurchargePrice(v float64) {
	o.SellFuelSurchargePrice = &v
}

// GetBuyCurrencyCode returns the BuyCurrencyCode field value if set, zero value otherwise.
func (o *TransportDto) GetBuyCurrencyCode() string {
	if o == nil || IsNil(o.BuyCurrencyCode) {
		var ret string
		return ret
	}
	return *o.BuyCurrencyCode
}

// GetBuyCurrencyCodeOk returns a tuple with the BuyCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetBuyCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BuyCurrencyCode) {
		return nil, false
	}
	return o.BuyCurrencyCode, true
}

// HasBuyCurrencyCode returns a boolean if a field has been set.
func (o *TransportDto) HasBuyCurrencyCode() bool {
	if o != nil && !IsNil(o.BuyCurrencyCode) {
		return true
	}

	return false
}

// SetBuyCurrencyCode gets a reference to the given string and assigns it to the BuyCurrencyCode field.
func (o *TransportDto) SetBuyCurrencyCode(v string) {
	o.BuyCurrencyCode = &v
}

// GetBuyPrice returns the BuyPrice field value if set, zero value otherwise.
func (o *TransportDto) GetBuyPrice() float64 {
	if o == nil || IsNil(o.BuyPrice) {
		var ret float64
		return ret
	}
	return *o.BuyPrice
}

// GetBuyPriceOk returns a tuple with the BuyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetBuyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyPrice) {
		return nil, false
	}
	return o.BuyPrice, true
}

// HasBuyPrice returns a boolean if a field has been set.
func (o *TransportDto) HasBuyPrice() bool {
	if o != nil && !IsNil(o.BuyPrice) {
		return true
	}

	return false
}

// SetBuyPrice gets a reference to the given float64 and assigns it to the BuyPrice field.
func (o *TransportDto) SetBuyPrice(v float64) {
	o.BuyPrice = &v
}

// GetForcedBuyPrice returns the ForcedBuyPrice field value if set, zero value otherwise.
func (o *TransportDto) GetForcedBuyPrice() float64 {
	if o == nil || IsNil(o.ForcedBuyPrice) {
		var ret float64
		return ret
	}
	return *o.ForcedBuyPrice
}

// GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetForcedBuyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedBuyPrice) {
		return nil, false
	}
	return o.ForcedBuyPrice, true
}

// HasForcedBuyPrice returns a boolean if a field has been set.
func (o *TransportDto) HasForcedBuyPrice() bool {
	if o != nil && !IsNil(o.ForcedBuyPrice) {
		return true
	}

	return false
}

// SetForcedBuyPrice gets a reference to the given float64 and assigns it to the ForcedBuyPrice field.
func (o *TransportDto) SetForcedBuyPrice(v float64) {
	o.ForcedBuyPrice = &v
}

// GetBuyFuelSurchargePrice returns the BuyFuelSurchargePrice field value if set, zero value otherwise.
func (o *TransportDto) GetBuyFuelSurchargePrice() float64 {
	if o == nil || IsNil(o.BuyFuelSurchargePrice) {
		var ret float64
		return ret
	}
	return *o.BuyFuelSurchargePrice
}

// GetBuyFuelSurchargePriceOk returns a tuple with the BuyFuelSurchargePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetBuyFuelSurchargePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyFuelSurchargePrice) {
		return nil, false
	}
	return o.BuyFuelSurchargePrice, true
}

// HasBuyFuelSurchargePrice returns a boolean if a field has been set.
func (o *TransportDto) HasBuyFuelSurchargePrice() bool {
	if o != nil && !IsNil(o.BuyFuelSurchargePrice) {
		return true
	}

	return false
}

// SetBuyFuelSurchargePrice gets a reference to the given float64 and assigns it to the BuyFuelSurchargePrice field.
func (o *TransportDto) SetBuyFuelSurchargePrice(v float64) {
	o.BuyFuelSurchargePrice = &v
}

// GetReferenceCurrencyCode returns the ReferenceCurrencyCode field value if set, zero value otherwise.
func (o *TransportDto) GetReferenceCurrencyCode() string {
	if o == nil || IsNil(o.ReferenceCurrencyCode) {
		var ret string
		return ret
	}
	return *o.ReferenceCurrencyCode
}

// GetReferenceCurrencyCodeOk returns a tuple with the ReferenceCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetReferenceCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceCurrencyCode) {
		return nil, false
	}
	return o.ReferenceCurrencyCode, true
}

// HasReferenceCurrencyCode returns a boolean if a field has been set.
func (o *TransportDto) HasReferenceCurrencyCode() bool {
	if o != nil && !IsNil(o.ReferenceCurrencyCode) {
		return true
	}

	return false
}

// SetReferenceCurrencyCode gets a reference to the given string and assigns it to the ReferenceCurrencyCode field.
func (o *TransportDto) SetReferenceCurrencyCode(v string) {
	o.ReferenceCurrencyCode = &v
}

// GetSellPricingPathId returns the SellPricingPathId field value if set, zero value otherwise.
func (o *TransportDto) GetSellPricingPathId() int32 {
	if o == nil || IsNil(o.SellPricingPathId) {
		var ret int32
		return ret
	}
	return *o.SellPricingPathId
}

// GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetSellPricingPathIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SellPricingPathId) {
		return nil, false
	}
	return o.SellPricingPathId, true
}

// HasSellPricingPathId returns a boolean if a field has been set.
func (o *TransportDto) HasSellPricingPathId() bool {
	if o != nil && !IsNil(o.SellPricingPathId) {
		return true
	}

	return false
}

// SetSellPricingPathId gets a reference to the given int32 and assigns it to the SellPricingPathId field.
func (o *TransportDto) SetSellPricingPathId(v int32) {
	o.SellPricingPathId = &v
}

// GetIsFragmented returns the IsFragmented field value if set, zero value otherwise.
func (o *TransportDto) GetIsFragmented() bool {
	if o == nil || IsNil(o.IsFragmented) {
		var ret bool
		return ret
	}
	return *o.IsFragmented
}

// GetIsFragmentedOk returns a tuple with the IsFragmented field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIsFragmentedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFragmented) {
		return nil, false
	}
	return o.IsFragmented, true
}

// HasIsFragmented returns a boolean if a field has been set.
func (o *TransportDto) HasIsFragmented() bool {
	if o != nil && !IsNil(o.IsFragmented) {
		return true
	}

	return false
}

// SetIsFragmented gets a reference to the given bool and assigns it to the IsFragmented field.
func (o *TransportDto) SetIsFragmented(v bool) {
	o.IsFragmented = &v
}

// GetParentTransportId returns the ParentTransportId field value if set, zero value otherwise.
func (o *TransportDto) GetParentTransportId() int32 {
	if o == nil || IsNil(o.ParentTransportId) {
		var ret int32
		return ret
	}
	return *o.ParentTransportId
}

// GetParentTransportIdOk returns a tuple with the ParentTransportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetParentTransportIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentTransportId) {
		return nil, false
	}
	return o.ParentTransportId, true
}

// HasParentTransportId returns a boolean if a field has been set.
func (o *TransportDto) HasParentTransportId() bool {
	if o != nil && !IsNil(o.ParentTransportId) {
		return true
	}

	return false
}

// SetParentTransportId gets a reference to the given int32 and assigns it to the ParentTransportId field.
func (o *TransportDto) SetParentTransportId(v int32) {
	o.ParentTransportId = &v
}

// GetMissionIsInSequence returns the MissionIsInSequence field value if set, zero value otherwise.
func (o *TransportDto) GetMissionIsInSequence() bool {
	if o == nil || IsNil(o.MissionIsInSequence) {
		var ret bool
		return ret
	}
	return *o.MissionIsInSequence
}

// GetMissionIsInSequenceOk returns a tuple with the MissionIsInSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetMissionIsInSequenceOk() (*bool, bool) {
	if o == nil || IsNil(o.MissionIsInSequence) {
		return nil, false
	}
	return o.MissionIsInSequence, true
}

// HasMissionIsInSequence returns a boolean if a field has been set.
func (o *TransportDto) HasMissionIsInSequence() bool {
	if o != nil && !IsNil(o.MissionIsInSequence) {
		return true
	}

	return false
}

// SetMissionIsInSequence gets a reference to the given bool and assigns it to the MissionIsInSequence field.
func (o *TransportDto) SetMissionIsInSequence(v bool) {
	o.MissionIsInSequence = &v
}

// GetMissionIsUniqueOrder returns the MissionIsUniqueOrder field value if set, zero value otherwise.
func (o *TransportDto) GetMissionIsUniqueOrder() bool {
	if o == nil || IsNil(o.MissionIsUniqueOrder) {
		var ret bool
		return ret
	}
	return *o.MissionIsUniqueOrder
}

// GetMissionIsUniqueOrderOk returns a tuple with the MissionIsUniqueOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetMissionIsUniqueOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.MissionIsUniqueOrder) {
		return nil, false
	}
	return o.MissionIsUniqueOrder, true
}

// HasMissionIsUniqueOrder returns a boolean if a field has been set.
func (o *TransportDto) HasMissionIsUniqueOrder() bool {
	if o != nil && !IsNil(o.MissionIsUniqueOrder) {
		return true
	}

	return false
}

// SetMissionIsUniqueOrder gets a reference to the given bool and assigns it to the MissionIsUniqueOrder field.
func (o *TransportDto) SetMissionIsUniqueOrder(v bool) {
	o.MissionIsUniqueOrder = &v
}

// GetQuotationIsInSequence returns the QuotationIsInSequence field value if set, zero value otherwise.
func (o *TransportDto) GetQuotationIsInSequence() bool {
	if o == nil || IsNil(o.QuotationIsInSequence) {
		var ret bool
		return ret
	}
	return *o.QuotationIsInSequence
}

// GetQuotationIsInSequenceOk returns a tuple with the QuotationIsInSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetQuotationIsInSequenceOk() (*bool, bool) {
	if o == nil || IsNil(o.QuotationIsInSequence) {
		return nil, false
	}
	return o.QuotationIsInSequence, true
}

// HasQuotationIsInSequence returns a boolean if a field has been set.
func (o *TransportDto) HasQuotationIsInSequence() bool {
	if o != nil && !IsNil(o.QuotationIsInSequence) {
		return true
	}

	return false
}

// SetQuotationIsInSequence gets a reference to the given bool and assigns it to the QuotationIsInSequence field.
func (o *TransportDto) SetQuotationIsInSequence(v bool) {
	o.QuotationIsInSequence = &v
}

// GetQuotationIsUniqueOrder returns the QuotationIsUniqueOrder field value if set, zero value otherwise.
func (o *TransportDto) GetQuotationIsUniqueOrder() bool {
	if o == nil || IsNil(o.QuotationIsUniqueOrder) {
		var ret bool
		return ret
	}
	return *o.QuotationIsUniqueOrder
}

// GetQuotationIsUniqueOrderOk returns a tuple with the QuotationIsUniqueOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetQuotationIsUniqueOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.QuotationIsUniqueOrder) {
		return nil, false
	}
	return o.QuotationIsUniqueOrder, true
}

// HasQuotationIsUniqueOrder returns a boolean if a field has been set.
func (o *TransportDto) HasQuotationIsUniqueOrder() bool {
	if o != nil && !IsNil(o.QuotationIsUniqueOrder) {
		return true
	}

	return false
}

// SetQuotationIsUniqueOrder gets a reference to the given bool and assigns it to the QuotationIsUniqueOrder field.
func (o *TransportDto) SetQuotationIsUniqueOrder(v bool) {
	o.QuotationIsUniqueOrder = &v
}

// GetReference1 returns the Reference1 field value if set, zero value otherwise.
func (o *TransportDto) GetReference1() string {
	if o == nil || IsNil(o.Reference1) {
		var ret string
		return ret
	}
	return *o.Reference1
}

// GetReference1Ok returns a tuple with the Reference1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetReference1Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference1) {
		return nil, false
	}
	return o.Reference1, true
}

// HasReference1 returns a boolean if a field has been set.
func (o *TransportDto) HasReference1() bool {
	if o != nil && !IsNil(o.Reference1) {
		return true
	}

	return false
}

// SetReference1 gets a reference to the given string and assigns it to the Reference1 field.
func (o *TransportDto) SetReference1(v string) {
	o.Reference1 = &v
}

// GetReference2 returns the Reference2 field value if set, zero value otherwise.
func (o *TransportDto) GetReference2() string {
	if o == nil || IsNil(o.Reference2) {
		var ret string
		return ret
	}
	return *o.Reference2
}

// GetReference2Ok returns a tuple with the Reference2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetReference2Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference2) {
		return nil, false
	}
	return o.Reference2, true
}

// HasReference2 returns a boolean if a field has been set.
func (o *TransportDto) HasReference2() bool {
	if o != nil && !IsNil(o.Reference2) {
		return true
	}

	return false
}

// SetReference2 gets a reference to the given string and assigns it to the Reference2 field.
func (o *TransportDto) SetReference2(v string) {
	o.Reference2 = &v
}

// GetReference3 returns the Reference3 field value if set, zero value otherwise.
func (o *TransportDto) GetReference3() string {
	if o == nil || IsNil(o.Reference3) {
		var ret string
		return ret
	}
	return *o.Reference3
}

// GetReference3Ok returns a tuple with the Reference3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetReference3Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference3) {
		return nil, false
	}
	return o.Reference3, true
}

// HasReference3 returns a boolean if a field has been set.
func (o *TransportDto) HasReference3() bool {
	if o != nil && !IsNil(o.Reference3) {
		return true
	}

	return false
}

// SetReference3 gets a reference to the given string and assigns it to the Reference3 field.
func (o *TransportDto) SetReference3(v string) {
	o.Reference3 = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TransportDto) GetComment() TransportCommentDto {
	if o == nil || IsNil(o.Comment) {
		var ret TransportCommentDto
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetCommentOk() (*TransportCommentDto, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TransportDto) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given TransportCommentDto and assigns it to the Comment field.
func (o *TransportDto) SetComment(v TransportCommentDto) {
	o.Comment = &v
}

// GetCashOnDelivery returns the CashOnDelivery field value if set, zero value otherwise.
func (o *TransportDto) GetCashOnDelivery() TransportCashOnDeliveryDto {
	if o == nil || IsNil(o.CashOnDelivery) {
		var ret TransportCashOnDeliveryDto
		return ret
	}
	return *o.CashOnDelivery
}

// GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetCashOnDeliveryOk() (*TransportCashOnDeliveryDto, bool) {
	if o == nil || IsNil(o.CashOnDelivery) {
		return nil, false
	}
	return o.CashOnDelivery, true
}

// HasCashOnDelivery returns a boolean if a field has been set.
func (o *TransportDto) HasCashOnDelivery() bool {
	if o != nil && !IsNil(o.CashOnDelivery) {
		return true
	}

	return false
}

// SetCashOnDelivery gets a reference to the given TransportCashOnDeliveryDto and assigns it to the CashOnDelivery field.
func (o *TransportDto) SetCashOnDelivery(v TransportCashOnDeliveryDto) {
	o.CashOnDelivery = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TransportDto) GetIncluded() TransportIncludedDto {
	if o == nil || IsNil(o.Included) {
		var ret TransportIncludedDto
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDto) GetIncludedOk() (*TransportIncludedDto, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TransportDto) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given TransportIncludedDto and assigns it to the Included field.
func (o *TransportDto) SetIncluded(v TransportIncludedDto) {
	o.Included = &v
}

func (o TransportDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.TransportId) {
		toSerialize["transportId"] = o.TransportId
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.MissionNumber) {
		toSerialize["missionNumber"] = o.MissionNumber
	}
	if !IsNil(o.MissionUid) {
		toSerialize["missionUid"] = o.MissionUid
	}
	if !IsNil(o.MissionTrackId) {
		toSerialize["missionTrackId"] = o.MissionTrackId
	}
	if !IsNil(o.QuotationNumber) {
		toSerialize["quotationNumber"] = o.QuotationNumber
	}
	if !IsNil(o.QuotationUid) {
		toSerialize["quotationUid"] = o.QuotationUid
	}
	if !IsNil(o.QuotationTrackId) {
		toSerialize["quotationTrackId"] = o.QuotationTrackId
	}
	if !IsNil(o.IsQuotationArchived) {
		toSerialize["isQuotationArchived"] = o.IsQuotationArchived
	}
	if !IsNil(o.IsQuotationFinalized) {
		toSerialize["isQuotationFinalized"] = o.IsQuotationFinalized
	}
	if !IsNil(o.IsQuotationSubjectToApproval) {
		toSerialize["isQuotationSubjectToApproval"] = o.IsQuotationSubjectToApproval
	}
	if !IsNil(o.IsMissionReadyToBill) {
		toSerialize["isMissionReadyToBill"] = o.IsMissionReadyToBill
	}
	if !IsNil(o.SourceQuotationNumber) {
		toSerialize["sourceQuotationNumber"] = o.SourceQuotationNumber
	}
	if !IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !IsNil(o.CustomerCode) {
		toSerialize["customerCode"] = o.CustomerCode
	}
	if !IsNil(o.AgencyCode) {
		toSerialize["agencyCode"] = o.AgencyCode
	}
	if !IsNil(o.OrdererCode) {
		toSerialize["ordererCode"] = o.OrdererCode
	}
	if !IsNil(o.OrdererName) {
		toSerialize["ordererName"] = o.OrdererName
	}
	if !IsNil(o.ContractorAgentId) {
		toSerialize["contractorAgentId"] = o.ContractorAgentId
	}
	if !IsNil(o.ContractorAgentName) {
		toSerialize["contractorAgentName"] = o.ContractorAgentName
	}
	if !IsNil(o.HandlerId) {
		toSerialize["handlerId"] = o.HandlerId
	}
	if !IsNil(o.VehicleCode) {
		toSerialize["vehicleCode"] = o.VehicleCode
	}
	if !IsNil(o.TrailerCode) {
		toSerialize["trailerCode"] = o.TrailerCode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubStateCode) {
		toSerialize["subStateCode"] = o.SubStateCode
	}
	if !IsNil(o.DriverId) {
		toSerialize["driverId"] = o.DriverId
	}
	if !IsNil(o.Sign) {
		toSerialize["sign"] = o.Sign
	}
	if !IsNil(o.SecretCode) {
		toSerialize["secretCode"] = o.SecretCode
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.IsRoundTrip) {
		toSerialize["isRoundTrip"] = o.IsRoundTrip
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["serviceCode"] = o.ServiceCode
	}
	if !IsNil(o.SubServices) {
		toSerialize["subServices"] = o.SubServices
	}
	if !IsNil(o.PickupStep) {
		toSerialize["pickupStep"] = o.PickupStep
	}
	if !IsNil(o.DeliveryStep) {
		toSerialize["deliveryStep"] = o.DeliveryStep
	}
	if !IsNil(o.Packing) {
		toSerialize["packing"] = o.Packing
	}
	if !IsNil(o.DistanceKm) {
		toSerialize["distanceKm"] = o.DistanceKm
	}
	if !IsNil(o.TotalDurationMinutes) {
		toSerialize["totalDurationMinutes"] = o.TotalDurationMinutes
	}
	if !IsNil(o.IsStrategic) {
		toSerialize["isStrategic"] = o.IsStrategic
	}
	if !IsNil(o.IsShared) {
		toSerialize["isShared"] = o.IsShared
	}
	if !IsNil(o.ResponsibleOperatorId) {
		toSerialize["responsibleOperatorId"] = o.ResponsibleOperatorId
	}
	if !IsNil(o.GasEmission) {
		toSerialize["gasEmission"] = o.GasEmission
	}
	if !IsNil(o.IsGasEmissionForced) {
		toSerialize["isGasEmissionForced"] = o.IsGasEmissionForced
	}
	if !IsNil(o.SellCurrencyCode) {
		toSerialize["sellCurrencyCode"] = o.SellCurrencyCode
	}
	if !IsNil(o.SellPrice) {
		toSerialize["sellPrice"] = o.SellPrice
	}
	if !IsNil(o.ForcedSellPrice) {
		toSerialize["forcedSellPrice"] = o.ForcedSellPrice
	}
	if !IsNil(o.SellFuelSurchargePrice) {
		toSerialize["sellFuelSurchargePrice"] = o.SellFuelSurchargePrice
	}
	if !IsNil(o.BuyCurrencyCode) {
		toSerialize["buyCurrencyCode"] = o.BuyCurrencyCode
	}
	if !IsNil(o.BuyPrice) {
		toSerialize["buyPrice"] = o.BuyPrice
	}
	if !IsNil(o.ForcedBuyPrice) {
		toSerialize["forcedBuyPrice"] = o.ForcedBuyPrice
	}
	if !IsNil(o.BuyFuelSurchargePrice) {
		toSerialize["buyFuelSurchargePrice"] = o.BuyFuelSurchargePrice
	}
	if !IsNil(o.ReferenceCurrencyCode) {
		toSerialize["referenceCurrencyCode"] = o.ReferenceCurrencyCode
	}
	if !IsNil(o.SellPricingPathId) {
		toSerialize["sellPricingPathId"] = o.SellPricingPathId
	}
	if !IsNil(o.IsFragmented) {
		toSerialize["isFragmented"] = o.IsFragmented
	}
	if !IsNil(o.ParentTransportId) {
		toSerialize["parentTransportId"] = o.ParentTransportId
	}
	if !IsNil(o.MissionIsInSequence) {
		toSerialize["missionIsInSequence"] = o.MissionIsInSequence
	}
	if !IsNil(o.MissionIsUniqueOrder) {
		toSerialize["missionIsUniqueOrder"] = o.MissionIsUniqueOrder
	}
	if !IsNil(o.QuotationIsInSequence) {
		toSerialize["quotationIsInSequence"] = o.QuotationIsInSequence
	}
	if !IsNil(o.QuotationIsUniqueOrder) {
		toSerialize["quotationIsUniqueOrder"] = o.QuotationIsUniqueOrder
	}
	if !IsNil(o.Reference1) {
		toSerialize["reference1"] = o.Reference1
	}
	if !IsNil(o.Reference2) {
		toSerialize["reference2"] = o.Reference2
	}
	if !IsNil(o.Reference3) {
		toSerialize["reference3"] = o.Reference3
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CashOnDelivery) {
		toSerialize["cashOnDelivery"] = o.CashOnDelivery
	}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableTransportDto struct {
	value *TransportDto
	isSet bool
}

func (v NullableTransportDto) Get() *TransportDto {
	return v.value
}

func (v *NullableTransportDto) Set(val *TransportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportDto(val *TransportDto) *NullableTransportDto {
	return &NullableTransportDto{value: val, isSet: true}
}

func (v NullableTransportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
