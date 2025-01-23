# DiffResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Diffs** | Pointer to [**[]DiffResponse1DiffsInner**](DiffResponse1DiffsInner.md) |  | [optional] 

## Methods

### NewDiffResponse1

`func NewDiffResponse1() *DiffResponse1`

NewDiffResponse1 instantiates a new DiffResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffResponse1WithDefaults

`func NewDiffResponse1WithDefaults() *DiffResponse1`

NewDiffResponse1WithDefaults instantiates a new DiffResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *DiffResponse1) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DiffResponse1) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DiffResponse1) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *DiffResponse1) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *DiffResponse1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DiffResponse1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DiffResponse1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DiffResponse1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDiffs

`func (o *DiffResponse1) GetDiffs() []DiffResponse1DiffsInner`

GetDiffs returns the Diffs field if non-nil, zero value otherwise.

### GetDiffsOk

`func (o *DiffResponse1) GetDiffsOk() (*[]DiffResponse1DiffsInner, bool)`

GetDiffsOk returns a tuple with the Diffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffs

`func (o *DiffResponse1) SetDiffs(v []DiffResponse1DiffsInner)`

SetDiffs sets Diffs field to given value.

### HasDiffs

`func (o *DiffResponse1) HasDiffs() bool`

HasDiffs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


