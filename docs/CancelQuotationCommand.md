# CancelQuotationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotationUid** | **string** | Mandatory. Quotation unique identifier | 

## Methods

### NewCancelQuotationCommand

`func NewCancelQuotationCommand(quotationUid string, ) *CancelQuotationCommand`

NewCancelQuotationCommand instantiates a new CancelQuotationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelQuotationCommandWithDefaults

`func NewCancelQuotationCommandWithDefaults() *CancelQuotationCommand`

NewCancelQuotationCommandWithDefaults instantiates a new CancelQuotationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotationUid

`func (o *CancelQuotationCommand) GetQuotationUid() string`

GetQuotationUid returns the QuotationUid field if non-nil, zero value otherwise.

### GetQuotationUidOk

`func (o *CancelQuotationCommand) GetQuotationUidOk() (*string, bool)`

GetQuotationUidOk returns a tuple with the QuotationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationUid

`func (o *CancelQuotationCommand) SetQuotationUid(v string)`

SetQuotationUid sets QuotationUid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


