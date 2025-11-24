# CreateTransportDocumentLinkCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportUid** | **string** | Mandatory. Transport unique identifier | 
**Link** | **string** | Mandatory. Document link, it must be a valid url. | 
**Label** | Pointer to **string** | Document link label | [optional] 
**FileCategoryCode** | Pointer to **string** | File category code | [optional] 

## Methods

### NewCreateTransportDocumentLinkCommand

`func NewCreateTransportDocumentLinkCommand(transportUid string, link string, ) *CreateTransportDocumentLinkCommand`

NewCreateTransportDocumentLinkCommand instantiates a new CreateTransportDocumentLinkCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportDocumentLinkCommandWithDefaults

`func NewCreateTransportDocumentLinkCommandWithDefaults() *CreateTransportDocumentLinkCommand`

NewCreateTransportDocumentLinkCommandWithDefaults instantiates a new CreateTransportDocumentLinkCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *CreateTransportDocumentLinkCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *CreateTransportDocumentLinkCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *CreateTransportDocumentLinkCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.


### GetLink

`func (o *CreateTransportDocumentLinkCommand) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CreateTransportDocumentLinkCommand) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CreateTransportDocumentLinkCommand) SetLink(v string)`

SetLink sets Link field to given value.


### GetLabel

`func (o *CreateTransportDocumentLinkCommand) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateTransportDocumentLinkCommand) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateTransportDocumentLinkCommand) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateTransportDocumentLinkCommand) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFileCategoryCode

`func (o *CreateTransportDocumentLinkCommand) GetFileCategoryCode() string`

GetFileCategoryCode returns the FileCategoryCode field if non-nil, zero value otherwise.

### GetFileCategoryCodeOk

`func (o *CreateTransportDocumentLinkCommand) GetFileCategoryCodeOk() (*string, bool)`

GetFileCategoryCodeOk returns a tuple with the FileCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCategoryCode

`func (o *CreateTransportDocumentLinkCommand) SetFileCategoryCode(v string)`

SetFileCategoryCode sets FileCategoryCode field to given value.

### HasFileCategoryCode

`func (o *CreateTransportDocumentLinkCommand) HasFileCategoryCode() bool`

HasFileCategoryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


