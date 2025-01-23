# GetMultipleContentsResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]GetMultipleContentsResponseV2ContentsInner**](GetMultipleContentsResponseV2ContentsInner.md) |  | 
**EffectiveReference** | Pointer to [**Reference2**](Reference2.md) |  | [optional] 

## Methods

### NewGetMultipleContentsResponseV2

`func NewGetMultipleContentsResponseV2(contents []GetMultipleContentsResponseV2ContentsInner, ) *GetMultipleContentsResponseV2`

NewGetMultipleContentsResponseV2 instantiates a new GetMultipleContentsResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultipleContentsResponseV2WithDefaults

`func NewGetMultipleContentsResponseV2WithDefaults() *GetMultipleContentsResponseV2`

NewGetMultipleContentsResponseV2WithDefaults instantiates a new GetMultipleContentsResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *GetMultipleContentsResponseV2) GetContents() []GetMultipleContentsResponseV2ContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *GetMultipleContentsResponseV2) GetContentsOk() (*[]GetMultipleContentsResponseV2ContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *GetMultipleContentsResponseV2) SetContents(v []GetMultipleContentsResponseV2ContentsInner)`

SetContents sets Contents field to given value.


### GetEffectiveReference

`func (o *GetMultipleContentsResponseV2) GetEffectiveReference() Reference2`

GetEffectiveReference returns the EffectiveReference field if non-nil, zero value otherwise.

### GetEffectiveReferenceOk

`func (o *GetMultipleContentsResponseV2) GetEffectiveReferenceOk() (*Reference2, bool)`

GetEffectiveReferenceOk returns a tuple with the EffectiveReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveReference

`func (o *GetMultipleContentsResponseV2) SetEffectiveReference(v Reference2)`

SetEffectiveReference sets EffectiveReference field to given value.

### HasEffectiveReference

`func (o *GetMultipleContentsResponseV2) HasEffectiveReference() bool`

HasEffectiveReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


