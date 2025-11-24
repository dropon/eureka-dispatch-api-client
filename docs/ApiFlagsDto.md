# ApiFlagsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyVersion** | Pointer to **string** | Dispatch build version | [optional] 
**Flags** | Pointer to [**[]ApiFlagsDtoApiFlag**](ApiFlagsDtoApiFlag.md) | Collection of feature flags | [optional] 

## Methods

### NewApiFlagsDto

`func NewApiFlagsDto() *ApiFlagsDto`

NewApiFlagsDto instantiates a new ApiFlagsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFlagsDtoWithDefaults

`func NewApiFlagsDtoWithDefaults() *ApiFlagsDto`

NewApiFlagsDtoWithDefaults instantiates a new ApiFlagsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyVersion

`func (o *ApiFlagsDto) GetAssemblyVersion() string`

GetAssemblyVersion returns the AssemblyVersion field if non-nil, zero value otherwise.

### GetAssemblyVersionOk

`func (o *ApiFlagsDto) GetAssemblyVersionOk() (*string, bool)`

GetAssemblyVersionOk returns a tuple with the AssemblyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyVersion

`func (o *ApiFlagsDto) SetAssemblyVersion(v string)`

SetAssemblyVersion sets AssemblyVersion field to given value.

### HasAssemblyVersion

`func (o *ApiFlagsDto) HasAssemblyVersion() bool`

HasAssemblyVersion returns a boolean if a field has been set.

### GetFlags

`func (o *ApiFlagsDto) GetFlags() []ApiFlagsDtoApiFlag`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ApiFlagsDto) GetFlagsOk() (*[]ApiFlagsDtoApiFlag, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ApiFlagsDto) SetFlags(v []ApiFlagsDtoApiFlag)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ApiFlagsDto) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


