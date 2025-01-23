# DiffResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Diffs** | Pointer to [**[]DiffResponse2DiffsInner**](DiffResponse2DiffsInner.md) |  | [optional] 
**EffectiveFromReference** | Pointer to [**Reference3**](Reference3.md) |  | [optional] 
**EffectiveToReference** | Pointer to [**Reference3**](Reference3.md) |  | [optional] 

## Methods

### NewDiffResponse2

`func NewDiffResponse2() *DiffResponse2`

NewDiffResponse2 instantiates a new DiffResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffResponse2WithDefaults

`func NewDiffResponse2WithDefaults() *DiffResponse2`

NewDiffResponse2WithDefaults instantiates a new DiffResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *DiffResponse2) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *DiffResponse2) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *DiffResponse2) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *DiffResponse2) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *DiffResponse2) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DiffResponse2) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DiffResponse2) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DiffResponse2) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDiffs

`func (o *DiffResponse2) GetDiffs() []DiffResponse2DiffsInner`

GetDiffs returns the Diffs field if non-nil, zero value otherwise.

### GetDiffsOk

`func (o *DiffResponse2) GetDiffsOk() (*[]DiffResponse2DiffsInner, bool)`

GetDiffsOk returns a tuple with the Diffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffs

`func (o *DiffResponse2) SetDiffs(v []DiffResponse2DiffsInner)`

SetDiffs sets Diffs field to given value.

### HasDiffs

`func (o *DiffResponse2) HasDiffs() bool`

HasDiffs returns a boolean if a field has been set.

### GetEffectiveFromReference

`func (o *DiffResponse2) GetEffectiveFromReference() Reference3`

GetEffectiveFromReference returns the EffectiveFromReference field if non-nil, zero value otherwise.

### GetEffectiveFromReferenceOk

`func (o *DiffResponse2) GetEffectiveFromReferenceOk() (*Reference3, bool)`

GetEffectiveFromReferenceOk returns a tuple with the EffectiveFromReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFromReference

`func (o *DiffResponse2) SetEffectiveFromReference(v Reference3)`

SetEffectiveFromReference sets EffectiveFromReference field to given value.

### HasEffectiveFromReference

`func (o *DiffResponse2) HasEffectiveFromReference() bool`

HasEffectiveFromReference returns a boolean if a field has been set.

### GetEffectiveToReference

`func (o *DiffResponse2) GetEffectiveToReference() Reference3`

GetEffectiveToReference returns the EffectiveToReference field if non-nil, zero value otherwise.

### GetEffectiveToReferenceOk

`func (o *DiffResponse2) GetEffectiveToReferenceOk() (*Reference3, bool)`

GetEffectiveToReferenceOk returns a tuple with the EffectiveToReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveToReference

`func (o *DiffResponse2) SetEffectiveToReference(v Reference3)`

SetEffectiveToReference sets EffectiveToReference field to given value.

### HasEffectiveToReference

`func (o *DiffResponse2) HasEffectiveToReference() bool`

HasEffectiveToReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


