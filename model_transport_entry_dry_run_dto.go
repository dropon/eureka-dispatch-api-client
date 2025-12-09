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

// checks if the TransportEntryDryRunDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportEntryDryRunDto{}

// TransportEntryDryRunDto struct for TransportEntryDryRunDto
type TransportEntryDryRunDto struct {
	CustomerCode             *string                                     `json:"customerCode,omitempty"`
	AgencyCode               *string                                     `json:"agencyCode,omitempty"`
	OrdererCode              *string                                     `json:"ordererCode,omitempty"`
	OrdererName              *string                                     `json:"ordererName,omitempty"`
	Sign                     *string                                     `json:"sign,omitempty"`
	SecretCode               *string                                     `json:"secretCode,omitempty"`
	Notes                    *string                                     `json:"notes,omitempty"`
	IsRoundTrip              *bool                                       `json:"isRoundTrip,omitempty"`
	CustomerCallDateTime     *Time                                       `json:"customerCallDateTime,omitempty"`
	PickupStep               *TransportEntryDryRunDtoStepDto             `json:"pickupStep,omitempty"`
	DeliveryStep             *TransportEntryDryRunDtoStepDto             `json:"deliveryStep,omitempty"`
	ServiceCode              *string                                     `json:"serviceCode,omitempty"`
	SubServices              []TransportEntryDryRunDtoSubServiceDto      `json:"subServices,omitempty"`
	Packing                  *TransportEntryDryRunDtoPackingDto          `json:"packing,omitempty"`
	CustomerCustomParameters []TransportEntryDryRunDtoCustomParameterDto `json:"customerCustomParameters,omitempty"`
	ServiceCustomParameters  []TransportEntryDryRunDtoCustomParameterDto `json:"serviceCustomParameters,omitempty"`
	Reference1               *string                                     `json:"reference1,omitempty"`
	Reference2               *string                                     `json:"reference2,omitempty"`
	Reference3               *string                                     `json:"reference3,omitempty"`
	DistanceKm               *float64                                    `json:"distanceKm,omitempty"`
	TotalDurationMinutes     *int32                                      `json:"totalDurationMinutes,omitempty"`
	IsStrategic              *bool                                       `json:"isStrategic,omitempty"`
	IsShared                 *bool                                       `json:"isShared,omitempty"`
	ResponsibleOperatorId    *int32                                      `json:"responsibleOperatorId,omitempty"`
	GasEmission              *float64                                    `json:"gasEmission,omitempty"`
	IsGasEmissionForced      *bool                                       `json:"isGasEmissionForced,omitempty"`
	SellCurrencyCode         *string                                     `json:"sellCurrencyCode,omitempty"`
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
	// Transport's sell price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	SellPriceInReferenceCurrency *float64 `json:"sellPriceInReferenceCurrency,omitempty"`
	// Transport's forced sell price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	ForcedSellPriceInReferenceCurrency *float64 `json:"forcedSellPriceInReferenceCurrency,omitempty"`
	// Transport's sell fuel surcharge price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	SellFuelSurchargePriceInReferenceCurrency *float64 `json:"sellFuelSurchargePriceInReferenceCurrency,omitempty"`
	// Transport's buy price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	BuyPriceInReferenceCurrency *float64 `json:"buyPriceInReferenceCurrency,omitempty"`
	// Transport's forced buy price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	ForcedBuyPriceInReferenceCurrency *float64 `json:"forcedBuyPriceInReferenceCurrency,omitempty"`
	// Transport's buy fuel surcharge price converted to reference currency    The following permission(s) are required to access this property:  See prices.
	BuyFuelSurchargePriceInReferenceCurrency *float64                                        `json:"buyFuelSurchargePriceInReferenceCurrency,omitempty"`
	Comment                                  *TransportEntryDryRunDtoCommentDto              `json:"comment,omitempty"`
	TransportBillAddress                     *TransportEntryDryRunDtoTransportBillAddressDto `json:"transportBillAddress,omitempty"`
	CashOnDelivery                           *TransportEntryDryRunDtoCashOnDeliveryDto       `json:"cashOnDelivery,omitempty"`
	Communication                            *TransportCommunicationConfigurationDto         `json:"communication,omitempty"`
	DangerousGoods                           []TransportEntryDryRunDtoDangerousGoodDto       `json:"dangerousGoods,omitempty"`
}

// NewTransportEntryDryRunDto instantiates a new TransportEntryDryRunDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportEntryDryRunDto() *TransportEntryDryRunDto {
	this := TransportEntryDryRunDto{}
	return &this
}

// NewTransportEntryDryRunDtoWithDefaults instantiates a new TransportEntryDryRunDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportEntryDryRunDtoWithDefaults() *TransportEntryDryRunDto {
	this := TransportEntryDryRunDto{}
	return &this
}

// GetCustomerCode returns the CustomerCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetCustomerCode() string {
	if o == nil || IsNil(o.CustomerCode) {
		var ret string
		return ret
	}
	return *o.CustomerCode
}

// GetCustomerCodeOk returns a tuple with the CustomerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCustomerCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCode) {
		return nil, false
	}
	return o.CustomerCode, true
}

// HasCustomerCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasCustomerCode() bool {
	if o != nil && !IsNil(o.CustomerCode) {
		return true
	}

	return false
}

// SetCustomerCode gets a reference to the given string and assigns it to the CustomerCode field.
func (o *TransportEntryDryRunDto) SetCustomerCode(v string) {
	o.CustomerCode = &v
}

// GetAgencyCode returns the AgencyCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetAgencyCode() string {
	if o == nil || IsNil(o.AgencyCode) {
		var ret string
		return ret
	}
	return *o.AgencyCode
}

// GetAgencyCodeOk returns a tuple with the AgencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetAgencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AgencyCode) {
		return nil, false
	}
	return o.AgencyCode, true
}

// HasAgencyCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasAgencyCode() bool {
	if o != nil && !IsNil(o.AgencyCode) {
		return true
	}

	return false
}

// SetAgencyCode gets a reference to the given string and assigns it to the AgencyCode field.
func (o *TransportEntryDryRunDto) SetAgencyCode(v string) {
	o.AgencyCode = &v
}

// GetOrdererCode returns the OrdererCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetOrdererCode() string {
	if o == nil || IsNil(o.OrdererCode) {
		var ret string
		return ret
	}
	return *o.OrdererCode
}

// GetOrdererCodeOk returns a tuple with the OrdererCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetOrdererCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OrdererCode) {
		return nil, false
	}
	return o.OrdererCode, true
}

// HasOrdererCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasOrdererCode() bool {
	if o != nil && !IsNil(o.OrdererCode) {
		return true
	}

	return false
}

// SetOrdererCode gets a reference to the given string and assigns it to the OrdererCode field.
func (o *TransportEntryDryRunDto) SetOrdererCode(v string) {
	o.OrdererCode = &v
}

// GetOrdererName returns the OrdererName field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetOrdererName() string {
	if o == nil || IsNil(o.OrdererName) {
		var ret string
		return ret
	}
	return *o.OrdererName
}

// GetOrdererNameOk returns a tuple with the OrdererName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetOrdererNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrdererName) {
		return nil, false
	}
	return o.OrdererName, true
}

// HasOrdererName returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasOrdererName() bool {
	if o != nil && !IsNil(o.OrdererName) {
		return true
	}

	return false
}

// SetOrdererName gets a reference to the given string and assigns it to the OrdererName field.
func (o *TransportEntryDryRunDto) SetOrdererName(v string) {
	o.OrdererName = &v
}

// GetSign returns the Sign field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSign() string {
	if o == nil || IsNil(o.Sign) {
		var ret string
		return ret
	}
	return *o.Sign
}

// GetSignOk returns a tuple with the Sign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSignOk() (*string, bool) {
	if o == nil || IsNil(o.Sign) {
		return nil, false
	}
	return o.Sign, true
}

// HasSign returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSign() bool {
	if o != nil && !IsNil(o.Sign) {
		return true
	}

	return false
}

// SetSign gets a reference to the given string and assigns it to the Sign field.
func (o *TransportEntryDryRunDto) SetSign(v string) {
	o.Sign = &v
}

// GetSecretCode returns the SecretCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSecretCode() string {
	if o == nil || IsNil(o.SecretCode) {
		var ret string
		return ret
	}
	return *o.SecretCode
}

// GetSecretCodeOk returns a tuple with the SecretCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSecretCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SecretCode) {
		return nil, false
	}
	return o.SecretCode, true
}

// HasSecretCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSecretCode() bool {
	if o != nil && !IsNil(o.SecretCode) {
		return true
	}

	return false
}

// SetSecretCode gets a reference to the given string and assigns it to the SecretCode field.
func (o *TransportEntryDryRunDto) SetSecretCode(v string) {
	o.SecretCode = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *TransportEntryDryRunDto) SetNotes(v string) {
	o.Notes = &v
}

// GetIsRoundTrip returns the IsRoundTrip field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetIsRoundTrip() bool {
	if o == nil || IsNil(o.IsRoundTrip) {
		var ret bool
		return ret
	}
	return *o.IsRoundTrip
}

// GetIsRoundTripOk returns a tuple with the IsRoundTrip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetIsRoundTripOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoundTrip) {
		return nil, false
	}
	return o.IsRoundTrip, true
}

// HasIsRoundTrip returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasIsRoundTrip() bool {
	if o != nil && !IsNil(o.IsRoundTrip) {
		return true
	}

	return false
}

// SetIsRoundTrip gets a reference to the given bool and assigns it to the IsRoundTrip field.
func (o *TransportEntryDryRunDto) SetIsRoundTrip(v bool) {
	o.IsRoundTrip = &v
}

// GetCustomerCallDateTime returns the CustomerCallDateTime field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetCustomerCallDateTime() Time {
	if o == nil || IsNil(o.CustomerCallDateTime) {
		var ret Time
		return ret
	}
	return *o.CustomerCallDateTime
}

// GetCustomerCallDateTimeOk returns a tuple with the CustomerCallDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCustomerCallDateTimeOk() (*Time, bool) {
	if o == nil || IsNil(o.CustomerCallDateTime) {
		return nil, false
	}
	return o.CustomerCallDateTime, true
}

// HasCustomerCallDateTime returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasCustomerCallDateTime() bool {
	if o != nil && !IsNil(o.CustomerCallDateTime) {
		return true
	}

	return false
}

// SetCustomerCallDateTime gets a reference to the given Time and assigns it to the CustomerCallDateTime field.
func (o *TransportEntryDryRunDto) SetCustomerCallDateTime(v Time) {
	o.CustomerCallDateTime = &v
}

// GetPickupStep returns the PickupStep field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetPickupStep() TransportEntryDryRunDtoStepDto {
	if o == nil || IsNil(o.PickupStep) {
		var ret TransportEntryDryRunDtoStepDto
		return ret
	}
	return *o.PickupStep
}

// GetPickupStepOk returns a tuple with the PickupStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetPickupStepOk() (*TransportEntryDryRunDtoStepDto, bool) {
	if o == nil || IsNil(o.PickupStep) {
		return nil, false
	}
	return o.PickupStep, true
}

// HasPickupStep returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasPickupStep() bool {
	if o != nil && !IsNil(o.PickupStep) {
		return true
	}

	return false
}

// SetPickupStep gets a reference to the given TransportEntryDryRunDtoStepDto and assigns it to the PickupStep field.
func (o *TransportEntryDryRunDto) SetPickupStep(v TransportEntryDryRunDtoStepDto) {
	o.PickupStep = &v
}

// GetDeliveryStep returns the DeliveryStep field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetDeliveryStep() TransportEntryDryRunDtoStepDto {
	if o == nil || IsNil(o.DeliveryStep) {
		var ret TransportEntryDryRunDtoStepDto
		return ret
	}
	return *o.DeliveryStep
}

// GetDeliveryStepOk returns a tuple with the DeliveryStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetDeliveryStepOk() (*TransportEntryDryRunDtoStepDto, bool) {
	if o == nil || IsNil(o.DeliveryStep) {
		return nil, false
	}
	return o.DeliveryStep, true
}

// HasDeliveryStep returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasDeliveryStep() bool {
	if o != nil && !IsNil(o.DeliveryStep) {
		return true
	}

	return false
}

// SetDeliveryStep gets a reference to the given TransportEntryDryRunDtoStepDto and assigns it to the DeliveryStep field.
func (o *TransportEntryDryRunDto) SetDeliveryStep(v TransportEntryDryRunDtoStepDto) {
	o.DeliveryStep = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *TransportEntryDryRunDto) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetSubServices returns the SubServices field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSubServices() []TransportEntryDryRunDtoSubServiceDto {
	if o == nil || IsNil(o.SubServices) {
		var ret []TransportEntryDryRunDtoSubServiceDto
		return ret
	}
	return o.SubServices
}

// GetSubServicesOk returns a tuple with the SubServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSubServicesOk() ([]TransportEntryDryRunDtoSubServiceDto, bool) {
	if o == nil || IsNil(o.SubServices) {
		return nil, false
	}
	return o.SubServices, true
}

// HasSubServices returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSubServices() bool {
	if o != nil && !IsNil(o.SubServices) {
		return true
	}

	return false
}

// SetSubServices gets a reference to the given []TransportEntryDryRunDtoSubServiceDto and assigns it to the SubServices field.
func (o *TransportEntryDryRunDto) SetSubServices(v []TransportEntryDryRunDtoSubServiceDto) {
	o.SubServices = v
}

// GetPacking returns the Packing field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetPacking() TransportEntryDryRunDtoPackingDto {
	if o == nil || IsNil(o.Packing) {
		var ret TransportEntryDryRunDtoPackingDto
		return ret
	}
	return *o.Packing
}

// GetPackingOk returns a tuple with the Packing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetPackingOk() (*TransportEntryDryRunDtoPackingDto, bool) {
	if o == nil || IsNil(o.Packing) {
		return nil, false
	}
	return o.Packing, true
}

// HasPacking returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasPacking() bool {
	if o != nil && !IsNil(o.Packing) {
		return true
	}

	return false
}

// SetPacking gets a reference to the given TransportEntryDryRunDtoPackingDto and assigns it to the Packing field.
func (o *TransportEntryDryRunDto) SetPacking(v TransportEntryDryRunDtoPackingDto) {
	o.Packing = &v
}

// GetCustomerCustomParameters returns the CustomerCustomParameters field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetCustomerCustomParameters() []TransportEntryDryRunDtoCustomParameterDto {
	if o == nil || IsNil(o.CustomerCustomParameters) {
		var ret []TransportEntryDryRunDtoCustomParameterDto
		return ret
	}
	return o.CustomerCustomParameters
}

// GetCustomerCustomParametersOk returns a tuple with the CustomerCustomParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCustomerCustomParametersOk() ([]TransportEntryDryRunDtoCustomParameterDto, bool) {
	if o == nil || IsNil(o.CustomerCustomParameters) {
		return nil, false
	}
	return o.CustomerCustomParameters, true
}

// HasCustomerCustomParameters returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasCustomerCustomParameters() bool {
	if o != nil && !IsNil(o.CustomerCustomParameters) {
		return true
	}

	return false
}

// SetCustomerCustomParameters gets a reference to the given []TransportEntryDryRunDtoCustomParameterDto and assigns it to the CustomerCustomParameters field.
func (o *TransportEntryDryRunDto) SetCustomerCustomParameters(v []TransportEntryDryRunDtoCustomParameterDto) {
	o.CustomerCustomParameters = v
}

// GetServiceCustomParameters returns the ServiceCustomParameters field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetServiceCustomParameters() []TransportEntryDryRunDtoCustomParameterDto {
	if o == nil || IsNil(o.ServiceCustomParameters) {
		var ret []TransportEntryDryRunDtoCustomParameterDto
		return ret
	}
	return o.ServiceCustomParameters
}

// GetServiceCustomParametersOk returns a tuple with the ServiceCustomParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetServiceCustomParametersOk() ([]TransportEntryDryRunDtoCustomParameterDto, bool) {
	if o == nil || IsNil(o.ServiceCustomParameters) {
		return nil, false
	}
	return o.ServiceCustomParameters, true
}

// HasServiceCustomParameters returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasServiceCustomParameters() bool {
	if o != nil && !IsNil(o.ServiceCustomParameters) {
		return true
	}

	return false
}

// SetServiceCustomParameters gets a reference to the given []TransportEntryDryRunDtoCustomParameterDto and assigns it to the ServiceCustomParameters field.
func (o *TransportEntryDryRunDto) SetServiceCustomParameters(v []TransportEntryDryRunDtoCustomParameterDto) {
	o.ServiceCustomParameters = v
}

// GetReference1 returns the Reference1 field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetReference1() string {
	if o == nil || IsNil(o.Reference1) {
		var ret string
		return ret
	}
	return *o.Reference1
}

// GetReference1Ok returns a tuple with the Reference1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetReference1Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference1) {
		return nil, false
	}
	return o.Reference1, true
}

// HasReference1 returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasReference1() bool {
	if o != nil && !IsNil(o.Reference1) {
		return true
	}

	return false
}

// SetReference1 gets a reference to the given string and assigns it to the Reference1 field.
func (o *TransportEntryDryRunDto) SetReference1(v string) {
	o.Reference1 = &v
}

// GetReference2 returns the Reference2 field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetReference2() string {
	if o == nil || IsNil(o.Reference2) {
		var ret string
		return ret
	}
	return *o.Reference2
}

// GetReference2Ok returns a tuple with the Reference2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetReference2Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference2) {
		return nil, false
	}
	return o.Reference2, true
}

// HasReference2 returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasReference2() bool {
	if o != nil && !IsNil(o.Reference2) {
		return true
	}

	return false
}

// SetReference2 gets a reference to the given string and assigns it to the Reference2 field.
func (o *TransportEntryDryRunDto) SetReference2(v string) {
	o.Reference2 = &v
}

// GetReference3 returns the Reference3 field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetReference3() string {
	if o == nil || IsNil(o.Reference3) {
		var ret string
		return ret
	}
	return *o.Reference3
}

// GetReference3Ok returns a tuple with the Reference3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetReference3Ok() (*string, bool) {
	if o == nil || IsNil(o.Reference3) {
		return nil, false
	}
	return o.Reference3, true
}

// HasReference3 returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasReference3() bool {
	if o != nil && !IsNil(o.Reference3) {
		return true
	}

	return false
}

// SetReference3 gets a reference to the given string and assigns it to the Reference3 field.
func (o *TransportEntryDryRunDto) SetReference3(v string) {
	o.Reference3 = &v
}

// GetDistanceKm returns the DistanceKm field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetDistanceKm() float64 {
	if o == nil || IsNil(o.DistanceKm) {
		var ret float64
		return ret
	}
	return *o.DistanceKm
}

// GetDistanceKmOk returns a tuple with the DistanceKm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetDistanceKmOk() (*float64, bool) {
	if o == nil || IsNil(o.DistanceKm) {
		return nil, false
	}
	return o.DistanceKm, true
}

// HasDistanceKm returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasDistanceKm() bool {
	if o != nil && !IsNil(o.DistanceKm) {
		return true
	}

	return false
}

// SetDistanceKm gets a reference to the given float64 and assigns it to the DistanceKm field.
func (o *TransportEntryDryRunDto) SetDistanceKm(v float64) {
	o.DistanceKm = &v
}

// GetTotalDurationMinutes returns the TotalDurationMinutes field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetTotalDurationMinutes() int32 {
	if o == nil || IsNil(o.TotalDurationMinutes) {
		var ret int32
		return ret
	}
	return *o.TotalDurationMinutes
}

// GetTotalDurationMinutesOk returns a tuple with the TotalDurationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetTotalDurationMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalDurationMinutes) {
		return nil, false
	}
	return o.TotalDurationMinutes, true
}

// HasTotalDurationMinutes returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasTotalDurationMinutes() bool {
	if o != nil && !IsNil(o.TotalDurationMinutes) {
		return true
	}

	return false
}

// SetTotalDurationMinutes gets a reference to the given int32 and assigns it to the TotalDurationMinutes field.
func (o *TransportEntryDryRunDto) SetTotalDurationMinutes(v int32) {
	o.TotalDurationMinutes = &v
}

// GetIsStrategic returns the IsStrategic field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetIsStrategic() bool {
	if o == nil || IsNil(o.IsStrategic) {
		var ret bool
		return ret
	}
	return *o.IsStrategic
}

// GetIsStrategicOk returns a tuple with the IsStrategic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetIsStrategicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsStrategic) {
		return nil, false
	}
	return o.IsStrategic, true
}

// HasIsStrategic returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasIsStrategic() bool {
	if o != nil && !IsNil(o.IsStrategic) {
		return true
	}

	return false
}

// SetIsStrategic gets a reference to the given bool and assigns it to the IsStrategic field.
func (o *TransportEntryDryRunDto) SetIsStrategic(v bool) {
	o.IsStrategic = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetIsShared() bool {
	if o == nil || IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetIsSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasIsShared() bool {
	if o != nil && !IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *TransportEntryDryRunDto) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetResponsibleOperatorId returns the ResponsibleOperatorId field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetResponsibleOperatorId() int32 {
	if o == nil || IsNil(o.ResponsibleOperatorId) {
		var ret int32
		return ret
	}
	return *o.ResponsibleOperatorId
}

// GetResponsibleOperatorIdOk returns a tuple with the ResponsibleOperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetResponsibleOperatorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponsibleOperatorId) {
		return nil, false
	}
	return o.ResponsibleOperatorId, true
}

// HasResponsibleOperatorId returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasResponsibleOperatorId() bool {
	if o != nil && !IsNil(o.ResponsibleOperatorId) {
		return true
	}

	return false
}

// SetResponsibleOperatorId gets a reference to the given int32 and assigns it to the ResponsibleOperatorId field.
func (o *TransportEntryDryRunDto) SetResponsibleOperatorId(v int32) {
	o.ResponsibleOperatorId = &v
}

// GetGasEmission returns the GasEmission field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetGasEmission() float64 {
	if o == nil || IsNil(o.GasEmission) {
		var ret float64
		return ret
	}
	return *o.GasEmission
}

// GetGasEmissionOk returns a tuple with the GasEmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetGasEmissionOk() (*float64, bool) {
	if o == nil || IsNil(o.GasEmission) {
		return nil, false
	}
	return o.GasEmission, true
}

// HasGasEmission returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasGasEmission() bool {
	if o != nil && !IsNil(o.GasEmission) {
		return true
	}

	return false
}

// SetGasEmission gets a reference to the given float64 and assigns it to the GasEmission field.
func (o *TransportEntryDryRunDto) SetGasEmission(v float64) {
	o.GasEmission = &v
}

// GetIsGasEmissionForced returns the IsGasEmissionForced field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetIsGasEmissionForced() bool {
	if o == nil || IsNil(o.IsGasEmissionForced) {
		var ret bool
		return ret
	}
	return *o.IsGasEmissionForced
}

// GetIsGasEmissionForcedOk returns a tuple with the IsGasEmissionForced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetIsGasEmissionForcedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGasEmissionForced) {
		return nil, false
	}
	return o.IsGasEmissionForced, true
}

// HasIsGasEmissionForced returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasIsGasEmissionForced() bool {
	if o != nil && !IsNil(o.IsGasEmissionForced) {
		return true
	}

	return false
}

// SetIsGasEmissionForced gets a reference to the given bool and assigns it to the IsGasEmissionForced field.
func (o *TransportEntryDryRunDto) SetIsGasEmissionForced(v bool) {
	o.IsGasEmissionForced = &v
}

// GetSellCurrencyCode returns the SellCurrencyCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellCurrencyCode() string {
	if o == nil || IsNil(o.SellCurrencyCode) {
		var ret string
		return ret
	}
	return *o.SellCurrencyCode
}

// GetSellCurrencyCodeOk returns a tuple with the SellCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SellCurrencyCode) {
		return nil, false
	}
	return o.SellCurrencyCode, true
}

// HasSellCurrencyCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellCurrencyCode() bool {
	if o != nil && !IsNil(o.SellCurrencyCode) {
		return true
	}

	return false
}

// SetSellCurrencyCode gets a reference to the given string and assigns it to the SellCurrencyCode field.
func (o *TransportEntryDryRunDto) SetSellCurrencyCode(v string) {
	o.SellCurrencyCode = &v
}

// GetSellPrice returns the SellPrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellPrice() float64 {
	if o == nil || IsNil(o.SellPrice) {
		var ret float64
		return ret
	}
	return *o.SellPrice
}

// GetSellPriceOk returns a tuple with the SellPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.SellPrice) {
		return nil, false
	}
	return o.SellPrice, true
}

// HasSellPrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellPrice() bool {
	if o != nil && !IsNil(o.SellPrice) {
		return true
	}

	return false
}

// SetSellPrice gets a reference to the given float64 and assigns it to the SellPrice field.
func (o *TransportEntryDryRunDto) SetSellPrice(v float64) {
	o.SellPrice = &v
}

// GetForcedSellPrice returns the ForcedSellPrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetForcedSellPrice() float64 {
	if o == nil || IsNil(o.ForcedSellPrice) {
		var ret float64
		return ret
	}
	return *o.ForcedSellPrice
}

// GetForcedSellPriceOk returns a tuple with the ForcedSellPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetForcedSellPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedSellPrice) {
		return nil, false
	}
	return o.ForcedSellPrice, true
}

// HasForcedSellPrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasForcedSellPrice() bool {
	if o != nil && !IsNil(o.ForcedSellPrice) {
		return true
	}

	return false
}

// SetForcedSellPrice gets a reference to the given float64 and assigns it to the ForcedSellPrice field.
func (o *TransportEntryDryRunDto) SetForcedSellPrice(v float64) {
	o.ForcedSellPrice = &v
}

// GetSellFuelSurchargePrice returns the SellFuelSurchargePrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellFuelSurchargePrice() float64 {
	if o == nil || IsNil(o.SellFuelSurchargePrice) {
		var ret float64
		return ret
	}
	return *o.SellFuelSurchargePrice
}

// GetSellFuelSurchargePriceOk returns a tuple with the SellFuelSurchargePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.SellFuelSurchargePrice) {
		return nil, false
	}
	return o.SellFuelSurchargePrice, true
}

// HasSellFuelSurchargePrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellFuelSurchargePrice() bool {
	if o != nil && !IsNil(o.SellFuelSurchargePrice) {
		return true
	}

	return false
}

// SetSellFuelSurchargePrice gets a reference to the given float64 and assigns it to the SellFuelSurchargePrice field.
func (o *TransportEntryDryRunDto) SetSellFuelSurchargePrice(v float64) {
	o.SellFuelSurchargePrice = &v
}

// GetBuyCurrencyCode returns the BuyCurrencyCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetBuyCurrencyCode() string {
	if o == nil || IsNil(o.BuyCurrencyCode) {
		var ret string
		return ret
	}
	return *o.BuyCurrencyCode
}

// GetBuyCurrencyCodeOk returns a tuple with the BuyCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetBuyCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BuyCurrencyCode) {
		return nil, false
	}
	return o.BuyCurrencyCode, true
}

// HasBuyCurrencyCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasBuyCurrencyCode() bool {
	if o != nil && !IsNil(o.BuyCurrencyCode) {
		return true
	}

	return false
}

// SetBuyCurrencyCode gets a reference to the given string and assigns it to the BuyCurrencyCode field.
func (o *TransportEntryDryRunDto) SetBuyCurrencyCode(v string) {
	o.BuyCurrencyCode = &v
}

// GetBuyPrice returns the BuyPrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetBuyPrice() float64 {
	if o == nil || IsNil(o.BuyPrice) {
		var ret float64
		return ret
	}
	return *o.BuyPrice
}

// GetBuyPriceOk returns a tuple with the BuyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetBuyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyPrice) {
		return nil, false
	}
	return o.BuyPrice, true
}

// HasBuyPrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasBuyPrice() bool {
	if o != nil && !IsNil(o.BuyPrice) {
		return true
	}

	return false
}

// SetBuyPrice gets a reference to the given float64 and assigns it to the BuyPrice field.
func (o *TransportEntryDryRunDto) SetBuyPrice(v float64) {
	o.BuyPrice = &v
}

// GetForcedBuyPrice returns the ForcedBuyPrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetForcedBuyPrice() float64 {
	if o == nil || IsNil(o.ForcedBuyPrice) {
		var ret float64
		return ret
	}
	return *o.ForcedBuyPrice
}

// GetForcedBuyPriceOk returns a tuple with the ForcedBuyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetForcedBuyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedBuyPrice) {
		return nil, false
	}
	return o.ForcedBuyPrice, true
}

// HasForcedBuyPrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasForcedBuyPrice() bool {
	if o != nil && !IsNil(o.ForcedBuyPrice) {
		return true
	}

	return false
}

// SetForcedBuyPrice gets a reference to the given float64 and assigns it to the ForcedBuyPrice field.
func (o *TransportEntryDryRunDto) SetForcedBuyPrice(v float64) {
	o.ForcedBuyPrice = &v
}

// GetBuyFuelSurchargePrice returns the BuyFuelSurchargePrice field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePrice() float64 {
	if o == nil || IsNil(o.BuyFuelSurchargePrice) {
		var ret float64
		return ret
	}
	return *o.BuyFuelSurchargePrice
}

// GetBuyFuelSurchargePriceOk returns a tuple with the BuyFuelSurchargePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyFuelSurchargePrice) {
		return nil, false
	}
	return o.BuyFuelSurchargePrice, true
}

// HasBuyFuelSurchargePrice returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasBuyFuelSurchargePrice() bool {
	if o != nil && !IsNil(o.BuyFuelSurchargePrice) {
		return true
	}

	return false
}

// SetBuyFuelSurchargePrice gets a reference to the given float64 and assigns it to the BuyFuelSurchargePrice field.
func (o *TransportEntryDryRunDto) SetBuyFuelSurchargePrice(v float64) {
	o.BuyFuelSurchargePrice = &v
}

// GetReferenceCurrencyCode returns the ReferenceCurrencyCode field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetReferenceCurrencyCode() string {
	if o == nil || IsNil(o.ReferenceCurrencyCode) {
		var ret string
		return ret
	}
	return *o.ReferenceCurrencyCode
}

// GetReferenceCurrencyCodeOk returns a tuple with the ReferenceCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetReferenceCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceCurrencyCode) {
		return nil, false
	}
	return o.ReferenceCurrencyCode, true
}

// HasReferenceCurrencyCode returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasReferenceCurrencyCode() bool {
	if o != nil && !IsNil(o.ReferenceCurrencyCode) {
		return true
	}

	return false
}

// SetReferenceCurrencyCode gets a reference to the given string and assigns it to the ReferenceCurrencyCode field.
func (o *TransportEntryDryRunDto) SetReferenceCurrencyCode(v string) {
	o.ReferenceCurrencyCode = &v
}

// GetSellPricingPathId returns the SellPricingPathId field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellPricingPathId() int32 {
	if o == nil || IsNil(o.SellPricingPathId) {
		var ret int32
		return ret
	}
	return *o.SellPricingPathId
}

// GetSellPricingPathIdOk returns a tuple with the SellPricingPathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellPricingPathIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SellPricingPathId) {
		return nil, false
	}
	return o.SellPricingPathId, true
}

// HasSellPricingPathId returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellPricingPathId() bool {
	if o != nil && !IsNil(o.SellPricingPathId) {
		return true
	}

	return false
}

// SetSellPricingPathId gets a reference to the given int32 and assigns it to the SellPricingPathId field.
func (o *TransportEntryDryRunDto) SetSellPricingPathId(v int32) {
	o.SellPricingPathId = &v
}

// GetSellPriceInReferenceCurrency returns the SellPriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellPriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.SellPriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.SellPriceInReferenceCurrency
}

// GetSellPriceInReferenceCurrencyOk returns a tuple with the SellPriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellPriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.SellPriceInReferenceCurrency) {
		return nil, false
	}
	return o.SellPriceInReferenceCurrency, true
}

// HasSellPriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellPriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.SellPriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetSellPriceInReferenceCurrency gets a reference to the given float64 and assigns it to the SellPriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetSellPriceInReferenceCurrency(v float64) {
	o.SellPriceInReferenceCurrency = &v
}

// GetForcedSellPriceInReferenceCurrency returns the ForcedSellPriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetForcedSellPriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.ForcedSellPriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.ForcedSellPriceInReferenceCurrency
}

// GetForcedSellPriceInReferenceCurrencyOk returns a tuple with the ForcedSellPriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetForcedSellPriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedSellPriceInReferenceCurrency) {
		return nil, false
	}
	return o.ForcedSellPriceInReferenceCurrency, true
}

// HasForcedSellPriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasForcedSellPriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.ForcedSellPriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetForcedSellPriceInReferenceCurrency gets a reference to the given float64 and assigns it to the ForcedSellPriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetForcedSellPriceInReferenceCurrency(v float64) {
	o.ForcedSellPriceInReferenceCurrency = &v
}

// GetSellFuelSurchargePriceInReferenceCurrency returns the SellFuelSurchargePriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.SellFuelSurchargePriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.SellFuelSurchargePriceInReferenceCurrency
}

// GetSellFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the SellFuelSurchargePriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetSellFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.SellFuelSurchargePriceInReferenceCurrency) {
		return nil, false
	}
	return o.SellFuelSurchargePriceInReferenceCurrency, true
}

// HasSellFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasSellFuelSurchargePriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.SellFuelSurchargePriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetSellFuelSurchargePriceInReferenceCurrency gets a reference to the given float64 and assigns it to the SellFuelSurchargePriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetSellFuelSurchargePriceInReferenceCurrency(v float64) {
	o.SellFuelSurchargePriceInReferenceCurrency = &v
}

// GetBuyPriceInReferenceCurrency returns the BuyPriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetBuyPriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.BuyPriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.BuyPriceInReferenceCurrency
}

// GetBuyPriceInReferenceCurrencyOk returns a tuple with the BuyPriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetBuyPriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyPriceInReferenceCurrency) {
		return nil, false
	}
	return o.BuyPriceInReferenceCurrency, true
}

// HasBuyPriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasBuyPriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.BuyPriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetBuyPriceInReferenceCurrency gets a reference to the given float64 and assigns it to the BuyPriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetBuyPriceInReferenceCurrency(v float64) {
	o.BuyPriceInReferenceCurrency = &v
}

// GetForcedBuyPriceInReferenceCurrency returns the ForcedBuyPriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetForcedBuyPriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.ForcedBuyPriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.ForcedBuyPriceInReferenceCurrency
}

// GetForcedBuyPriceInReferenceCurrencyOk returns a tuple with the ForcedBuyPriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetForcedBuyPriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.ForcedBuyPriceInReferenceCurrency) {
		return nil, false
	}
	return o.ForcedBuyPriceInReferenceCurrency, true
}

// HasForcedBuyPriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasForcedBuyPriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.ForcedBuyPriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetForcedBuyPriceInReferenceCurrency gets a reference to the given float64 and assigns it to the ForcedBuyPriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetForcedBuyPriceInReferenceCurrency(v float64) {
	o.ForcedBuyPriceInReferenceCurrency = &v
}

// GetBuyFuelSurchargePriceInReferenceCurrency returns the BuyFuelSurchargePriceInReferenceCurrency field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceInReferenceCurrency() float64 {
	if o == nil || IsNil(o.BuyFuelSurchargePriceInReferenceCurrency) {
		var ret float64
		return ret
	}
	return *o.BuyFuelSurchargePriceInReferenceCurrency
}

// GetBuyFuelSurchargePriceInReferenceCurrencyOk returns a tuple with the BuyFuelSurchargePriceInReferenceCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetBuyFuelSurchargePriceInReferenceCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyFuelSurchargePriceInReferenceCurrency) {
		return nil, false
	}
	return o.BuyFuelSurchargePriceInReferenceCurrency, true
}

// HasBuyFuelSurchargePriceInReferenceCurrency returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasBuyFuelSurchargePriceInReferenceCurrency() bool {
	if o != nil && !IsNil(o.BuyFuelSurchargePriceInReferenceCurrency) {
		return true
	}

	return false
}

// SetBuyFuelSurchargePriceInReferenceCurrency gets a reference to the given float64 and assigns it to the BuyFuelSurchargePriceInReferenceCurrency field.
func (o *TransportEntryDryRunDto) SetBuyFuelSurchargePriceInReferenceCurrency(v float64) {
	o.BuyFuelSurchargePriceInReferenceCurrency = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetComment() TransportEntryDryRunDtoCommentDto {
	if o == nil || IsNil(o.Comment) {
		var ret TransportEntryDryRunDtoCommentDto
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCommentOk() (*TransportEntryDryRunDtoCommentDto, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given TransportEntryDryRunDtoCommentDto and assigns it to the Comment field.
func (o *TransportEntryDryRunDto) SetComment(v TransportEntryDryRunDtoCommentDto) {
	o.Comment = &v
}

// GetTransportBillAddress returns the TransportBillAddress field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetTransportBillAddress() TransportEntryDryRunDtoTransportBillAddressDto {
	if o == nil || IsNil(o.TransportBillAddress) {
		var ret TransportEntryDryRunDtoTransportBillAddressDto
		return ret
	}
	return *o.TransportBillAddress
}

// GetTransportBillAddressOk returns a tuple with the TransportBillAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetTransportBillAddressOk() (*TransportEntryDryRunDtoTransportBillAddressDto, bool) {
	if o == nil || IsNil(o.TransportBillAddress) {
		return nil, false
	}
	return o.TransportBillAddress, true
}

// HasTransportBillAddress returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasTransportBillAddress() bool {
	if o != nil && !IsNil(o.TransportBillAddress) {
		return true
	}

	return false
}

// SetTransportBillAddress gets a reference to the given TransportEntryDryRunDtoTransportBillAddressDto and assigns it to the TransportBillAddress field.
func (o *TransportEntryDryRunDto) SetTransportBillAddress(v TransportEntryDryRunDtoTransportBillAddressDto) {
	o.TransportBillAddress = &v
}

// GetCashOnDelivery returns the CashOnDelivery field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetCashOnDelivery() TransportEntryDryRunDtoCashOnDeliveryDto {
	if o == nil || IsNil(o.CashOnDelivery) {
		var ret TransportEntryDryRunDtoCashOnDeliveryDto
		return ret
	}
	return *o.CashOnDelivery
}

// GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCashOnDeliveryOk() (*TransportEntryDryRunDtoCashOnDeliveryDto, bool) {
	if o == nil || IsNil(o.CashOnDelivery) {
		return nil, false
	}
	return o.CashOnDelivery, true
}

// HasCashOnDelivery returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasCashOnDelivery() bool {
	if o != nil && !IsNil(o.CashOnDelivery) {
		return true
	}

	return false
}

// SetCashOnDelivery gets a reference to the given TransportEntryDryRunDtoCashOnDeliveryDto and assigns it to the CashOnDelivery field.
func (o *TransportEntryDryRunDto) SetCashOnDelivery(v TransportEntryDryRunDtoCashOnDeliveryDto) {
	o.CashOnDelivery = &v
}

// GetCommunication returns the Communication field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetCommunication() TransportCommunicationConfigurationDto {
	if o == nil || IsNil(o.Communication) {
		var ret TransportCommunicationConfigurationDto
		return ret
	}
	return *o.Communication
}

// GetCommunicationOk returns a tuple with the Communication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetCommunicationOk() (*TransportCommunicationConfigurationDto, bool) {
	if o == nil || IsNil(o.Communication) {
		return nil, false
	}
	return o.Communication, true
}

// HasCommunication returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasCommunication() bool {
	if o != nil && !IsNil(o.Communication) {
		return true
	}

	return false
}

// SetCommunication gets a reference to the given TransportCommunicationConfigurationDto and assigns it to the Communication field.
func (o *TransportEntryDryRunDto) SetCommunication(v TransportCommunicationConfigurationDto) {
	o.Communication = &v
}

// GetDangerousGoods returns the DangerousGoods field value if set, zero value otherwise.
func (o *TransportEntryDryRunDto) GetDangerousGoods() []TransportEntryDryRunDtoDangerousGoodDto {
	if o == nil || IsNil(o.DangerousGoods) {
		var ret []TransportEntryDryRunDtoDangerousGoodDto
		return ret
	}
	return o.DangerousGoods
}

// GetDangerousGoodsOk returns a tuple with the DangerousGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEntryDryRunDto) GetDangerousGoodsOk() ([]TransportEntryDryRunDtoDangerousGoodDto, bool) {
	if o == nil || IsNil(o.DangerousGoods) {
		return nil, false
	}
	return o.DangerousGoods, true
}

// HasDangerousGoods returns a boolean if a field has been set.
func (o *TransportEntryDryRunDto) HasDangerousGoods() bool {
	if o != nil && !IsNil(o.DangerousGoods) {
		return true
	}

	return false
}

// SetDangerousGoods gets a reference to the given []TransportEntryDryRunDtoDangerousGoodDto and assigns it to the DangerousGoods field.
func (o *TransportEntryDryRunDto) SetDangerousGoods(v []TransportEntryDryRunDtoDangerousGoodDto) {
	o.DangerousGoods = v
}

func (o TransportEntryDryRunDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportEntryDryRunDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.CustomerCallDateTime) {
		toSerialize["customerCallDateTime"] = o.CustomerCallDateTime
	}
	if !IsNil(o.PickupStep) {
		toSerialize["pickupStep"] = o.PickupStep
	}
	if !IsNil(o.DeliveryStep) {
		toSerialize["deliveryStep"] = o.DeliveryStep
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["serviceCode"] = o.ServiceCode
	}
	if !IsNil(o.SubServices) {
		toSerialize["subServices"] = o.SubServices
	}
	if !IsNil(o.Packing) {
		toSerialize["packing"] = o.Packing
	}
	if !IsNil(o.CustomerCustomParameters) {
		toSerialize["customerCustomParameters"] = o.CustomerCustomParameters
	}
	if !IsNil(o.ServiceCustomParameters) {
		toSerialize["serviceCustomParameters"] = o.ServiceCustomParameters
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
	if !IsNil(o.SellPriceInReferenceCurrency) {
		toSerialize["sellPriceInReferenceCurrency"] = o.SellPriceInReferenceCurrency
	}
	if !IsNil(o.ForcedSellPriceInReferenceCurrency) {
		toSerialize["forcedSellPriceInReferenceCurrency"] = o.ForcedSellPriceInReferenceCurrency
	}
	if !IsNil(o.SellFuelSurchargePriceInReferenceCurrency) {
		toSerialize["sellFuelSurchargePriceInReferenceCurrency"] = o.SellFuelSurchargePriceInReferenceCurrency
	}
	if !IsNil(o.BuyPriceInReferenceCurrency) {
		toSerialize["buyPriceInReferenceCurrency"] = o.BuyPriceInReferenceCurrency
	}
	if !IsNil(o.ForcedBuyPriceInReferenceCurrency) {
		toSerialize["forcedBuyPriceInReferenceCurrency"] = o.ForcedBuyPriceInReferenceCurrency
	}
	if !IsNil(o.BuyFuelSurchargePriceInReferenceCurrency) {
		toSerialize["buyFuelSurchargePriceInReferenceCurrency"] = o.BuyFuelSurchargePriceInReferenceCurrency
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.TransportBillAddress) {
		toSerialize["transportBillAddress"] = o.TransportBillAddress
	}
	if !IsNil(o.CashOnDelivery) {
		toSerialize["cashOnDelivery"] = o.CashOnDelivery
	}
	if !IsNil(o.Communication) {
		toSerialize["communication"] = o.Communication
	}
	if !IsNil(o.DangerousGoods) {
		toSerialize["dangerousGoods"] = o.DangerousGoods
	}
	return toSerialize, nil
}

type NullableTransportEntryDryRunDto struct {
	value *TransportEntryDryRunDto
	isSet bool
}

func (v NullableTransportEntryDryRunDto) Get() *TransportEntryDryRunDto {
	return v.value
}

func (v *NullableTransportEntryDryRunDto) Set(val *TransportEntryDryRunDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportEntryDryRunDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportEntryDryRunDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportEntryDryRunDto(val *TransportEntryDryRunDto) *NullableTransportEntryDryRunDto {
	return &NullableTransportEntryDryRunDto{value: val, isSet: true}
}

func (v NullableTransportEntryDryRunDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportEntryDryRunDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
