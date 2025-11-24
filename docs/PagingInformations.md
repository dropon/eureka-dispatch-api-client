# PagingInformations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrevLink** | Pointer to **string** |  | [optional] 
**NextLink** | Pointer to **string** |  | [optional] 
**PageCount** | Pointer to **int32** |  | [optional] 
**TotalItemCount** | Pointer to **int32** |  | [optional] 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**HasPreviousPage** | Pointer to **bool** |  | [optional] 
**HasNextPage** | Pointer to **bool** |  | [optional] 
**IsFirstPage** | Pointer to **bool** |  | [optional] 
**IsLastPage** | Pointer to **bool** |  | [optional] 
**FirstItemOnPage** | Pointer to **int32** |  | [optional] 
**LastItemOnPage** | Pointer to **int32** |  | [optional] 

## Methods

### NewPagingInformations

`func NewPagingInformations() *PagingInformations`

NewPagingInformations instantiates a new PagingInformations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingInformationsWithDefaults

`func NewPagingInformationsWithDefaults() *PagingInformations`

NewPagingInformationsWithDefaults instantiates a new PagingInformations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevLink

`func (o *PagingInformations) GetPrevLink() string`

GetPrevLink returns the PrevLink field if non-nil, zero value otherwise.

### GetPrevLinkOk

`func (o *PagingInformations) GetPrevLinkOk() (*string, bool)`

GetPrevLinkOk returns a tuple with the PrevLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevLink

`func (o *PagingInformations) SetPrevLink(v string)`

SetPrevLink sets PrevLink field to given value.

### HasPrevLink

`func (o *PagingInformations) HasPrevLink() bool`

HasPrevLink returns a boolean if a field has been set.

### GetNextLink

`func (o *PagingInformations) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *PagingInformations) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *PagingInformations) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *PagingInformations) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.

### GetPageCount

`func (o *PagingInformations) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *PagingInformations) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *PagingInformations) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *PagingInformations) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetTotalItemCount

`func (o *PagingInformations) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *PagingInformations) GetTotalItemCountOk() (*int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemCount

`func (o *PagingInformations) SetTotalItemCount(v int32)`

SetTotalItemCount sets TotalItemCount field to given value.

### HasTotalItemCount

`func (o *PagingInformations) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.

### GetPageNumber

`func (o *PagingInformations) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *PagingInformations) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *PagingInformations) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *PagingInformations) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *PagingInformations) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PagingInformations) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PagingInformations) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PagingInformations) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetHasPreviousPage

`func (o *PagingInformations) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *PagingInformations) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *PagingInformations) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.

### HasHasPreviousPage

`func (o *PagingInformations) HasHasPreviousPage() bool`

HasHasPreviousPage returns a boolean if a field has been set.

### GetHasNextPage

`func (o *PagingInformations) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PagingInformations) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PagingInformations) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *PagingInformations) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetIsFirstPage

`func (o *PagingInformations) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *PagingInformations) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *PagingInformations) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *PagingInformations) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetIsLastPage

`func (o *PagingInformations) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *PagingInformations) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *PagingInformations) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *PagingInformations) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetFirstItemOnPage

`func (o *PagingInformations) GetFirstItemOnPage() int32`

GetFirstItemOnPage returns the FirstItemOnPage field if non-nil, zero value otherwise.

### GetFirstItemOnPageOk

`func (o *PagingInformations) GetFirstItemOnPageOk() (*int32, bool)`

GetFirstItemOnPageOk returns a tuple with the FirstItemOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstItemOnPage

`func (o *PagingInformations) SetFirstItemOnPage(v int32)`

SetFirstItemOnPage sets FirstItemOnPage field to given value.

### HasFirstItemOnPage

`func (o *PagingInformations) HasFirstItemOnPage() bool`

HasFirstItemOnPage returns a boolean if a field has been set.

### GetLastItemOnPage

`func (o *PagingInformations) GetLastItemOnPage() int32`

GetLastItemOnPage returns the LastItemOnPage field if non-nil, zero value otherwise.

### GetLastItemOnPageOk

`func (o *PagingInformations) GetLastItemOnPageOk() (*int32, bool)`

GetLastItemOnPageOk returns a tuple with the LastItemOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastItemOnPage

`func (o *PagingInformations) SetLastItemOnPage(v int32)`

SetLastItemOnPage sets LastItemOnPage field to given value.

### HasLastItemOnPage

`func (o *PagingInformations) HasLastItemOnPage() bool`

HasLastItemOnPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


