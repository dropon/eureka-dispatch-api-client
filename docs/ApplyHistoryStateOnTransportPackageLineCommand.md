# ApplyHistoryStateOnTransportPackageLineCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportUid** | Pointer to **string** | Transport unique identifier | [optional] 
**PackageLineId** | Pointer to **int32** | Package line identifier | [optional] 
**StatusCode** | Pointer to **string** | Status code | [optional] 
**StatusDate** | Pointer to **time.Time** | Status business date | [optional] 
**Comment** | Pointer to **string** | Status comment | [optional] 

## Methods

### NewApplyHistoryStateOnTransportPackageLineCommand

`func NewApplyHistoryStateOnTransportPackageLineCommand() *ApplyHistoryStateOnTransportPackageLineCommand`

NewApplyHistoryStateOnTransportPackageLineCommand instantiates a new ApplyHistoryStateOnTransportPackageLineCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyHistoryStateOnTransportPackageLineCommandWithDefaults

`func NewApplyHistoryStateOnTransportPackageLineCommandWithDefaults() *ApplyHistoryStateOnTransportPackageLineCommand`

NewApplyHistoryStateOnTransportPackageLineCommandWithDefaults instantiates a new ApplyHistoryStateOnTransportPackageLineCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportUid

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetTransportUid() string`

GetTransportUid returns the TransportUid field if non-nil, zero value otherwise.

### GetTransportUidOk

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetTransportUidOk() (*string, bool)`

GetTransportUidOk returns a tuple with the TransportUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportUid

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) SetTransportUid(v string)`

SetTransportUid sets TransportUid field to given value.

### HasTransportUid

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) HasTransportUid() bool`

HasTransportUid returns a boolean if a field has been set.

### GetPackageLineId

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetPackageLineId() int32`

GetPackageLineId returns the PackageLineId field if non-nil, zero value otherwise.

### GetPackageLineIdOk

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetPackageLineIdOk() (*int32, bool)`

GetPackageLineIdOk returns a tuple with the PackageLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineId

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) SetPackageLineId(v int32)`

SetPackageLineId sets PackageLineId field to given value.

### HasPackageLineId

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) HasPackageLineId() bool`

HasPackageLineId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusDate

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetComment

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApplyHistoryStateOnTransportPackageLineCommand) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


