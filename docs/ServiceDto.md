# ServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**BackgroundColor** | Pointer to **string** |  | [optional] 
**ForegroundColor** | Pointer to **string** |  | [optional] 
**WithReturn** | Pointer to **bool** |  | [optional] 
**EdiConfigurationId** | Pointer to **int32** |  | [optional] 
**ServiceFamily** | Pointer to [**ServiceFamilyDto**](ServiceFamilyDto.md) |  | [optional] 
**ServiceType** | Pointer to [**ServiceTypeDto**](ServiceTypeDto.md) |  | [optional] 

## Methods

### NewServiceDto

`func NewServiceDto() *ServiceDto`

NewServiceDto instantiates a new ServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDtoWithDefaults

`func NewServiceDtoWithDefaults() *ServiceDto`

NewServiceDtoWithDefaults instantiates a new ServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ServiceDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ServiceDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ServiceDto) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ServiceDto) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetCode

`func (o *ServiceDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetComment

`func (o *ServiceDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ServiceDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ServiceDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ServiceDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *ServiceDto) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *ServiceDto) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *ServiceDto) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *ServiceDto) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetForegroundColor

`func (o *ServiceDto) GetForegroundColor() string`

GetForegroundColor returns the ForegroundColor field if non-nil, zero value otherwise.

### GetForegroundColorOk

`func (o *ServiceDto) GetForegroundColorOk() (*string, bool)`

GetForegroundColorOk returns a tuple with the ForegroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundColor

`func (o *ServiceDto) SetForegroundColor(v string)`

SetForegroundColor sets ForegroundColor field to given value.

### HasForegroundColor

`func (o *ServiceDto) HasForegroundColor() bool`

HasForegroundColor returns a boolean if a field has been set.

### GetWithReturn

`func (o *ServiceDto) GetWithReturn() bool`

GetWithReturn returns the WithReturn field if non-nil, zero value otherwise.

### GetWithReturnOk

`func (o *ServiceDto) GetWithReturnOk() (*bool, bool)`

GetWithReturnOk returns a tuple with the WithReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithReturn

`func (o *ServiceDto) SetWithReturn(v bool)`

SetWithReturn sets WithReturn field to given value.

### HasWithReturn

`func (o *ServiceDto) HasWithReturn() bool`

HasWithReturn returns a boolean if a field has been set.

### GetEdiConfigurationId

`func (o *ServiceDto) GetEdiConfigurationId() int32`

GetEdiConfigurationId returns the EdiConfigurationId field if non-nil, zero value otherwise.

### GetEdiConfigurationIdOk

`func (o *ServiceDto) GetEdiConfigurationIdOk() (*int32, bool)`

GetEdiConfigurationIdOk returns a tuple with the EdiConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdiConfigurationId

`func (o *ServiceDto) SetEdiConfigurationId(v int32)`

SetEdiConfigurationId sets EdiConfigurationId field to given value.

### HasEdiConfigurationId

`func (o *ServiceDto) HasEdiConfigurationId() bool`

HasEdiConfigurationId returns a boolean if a field has been set.

### GetServiceFamily

`func (o *ServiceDto) GetServiceFamily() ServiceFamilyDto`

GetServiceFamily returns the ServiceFamily field if non-nil, zero value otherwise.

### GetServiceFamilyOk

`func (o *ServiceDto) GetServiceFamilyOk() (*ServiceFamilyDto, bool)`

GetServiceFamilyOk returns a tuple with the ServiceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFamily

`func (o *ServiceDto) SetServiceFamily(v ServiceFamilyDto)`

SetServiceFamily sets ServiceFamily field to given value.

### HasServiceFamily

`func (o *ServiceDto) HasServiceFamily() bool`

HasServiceFamily returns a boolean if a field has been set.

### GetServiceType

`func (o *ServiceDto) GetServiceType() ServiceTypeDto`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ServiceDto) GetServiceTypeOk() (*ServiceTypeDto, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ServiceDto) SetServiceType(v ServiceTypeDto)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ServiceDto) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


