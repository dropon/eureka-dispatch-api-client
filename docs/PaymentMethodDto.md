# PaymentMethodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentMethodDto

`func NewPaymentMethodDto() *PaymentMethodDto`

NewPaymentMethodDto instantiates a new PaymentMethodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDtoWithDefaults

`func NewPaymentMethodDtoWithDefaults() *PaymentMethodDto`

NewPaymentMethodDtoWithDefaults instantiates a new PaymentMethodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PaymentMethodDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentMethodDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentMethodDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentMethodDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentMethodDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethodDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethodDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethodDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


