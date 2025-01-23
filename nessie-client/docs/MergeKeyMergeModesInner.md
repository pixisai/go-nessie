# MergeKeyMergeModesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**AddedContentKey**](AddedContentKey.md) |  | [optional] 
**MergeBehavior** | Pointer to **string** |  | [optional] 
**ExpectedTargetContent** | Pointer to [**Content2**](Content2.md) |  | [optional] 
**ResolvedContent** | Pointer to [**Content2**](Content2.md) |  | [optional] 
**ExpectedTargetDocumentation** | Pointer to [**ContentResponseV2Documentation**](ContentResponseV2Documentation.md) |  | [optional] 
**ResolvedDocumentation** | Pointer to [**ContentResponseV2Documentation**](ContentResponseV2Documentation.md) |  | [optional] 
**Metadata** | Pointer to [**[]ContentMetadata1**](ContentMetadata1.md) |  | [optional] 

## Methods

### NewMergeKeyMergeModesInner

`func NewMergeKeyMergeModesInner() *MergeKeyMergeModesInner`

NewMergeKeyMergeModesInner instantiates a new MergeKeyMergeModesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeKeyMergeModesInnerWithDefaults

`func NewMergeKeyMergeModesInnerWithDefaults() *MergeKeyMergeModesInner`

NewMergeKeyMergeModesInnerWithDefaults instantiates a new MergeKeyMergeModesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MergeKeyMergeModesInner) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MergeKeyMergeModesInner) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MergeKeyMergeModesInner) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *MergeKeyMergeModesInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMergeBehavior

`func (o *MergeKeyMergeModesInner) GetMergeBehavior() string`

GetMergeBehavior returns the MergeBehavior field if non-nil, zero value otherwise.

### GetMergeBehaviorOk

`func (o *MergeKeyMergeModesInner) GetMergeBehaviorOk() (*string, bool)`

GetMergeBehaviorOk returns a tuple with the MergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeBehavior

`func (o *MergeKeyMergeModesInner) SetMergeBehavior(v string)`

SetMergeBehavior sets MergeBehavior field to given value.

### HasMergeBehavior

`func (o *MergeKeyMergeModesInner) HasMergeBehavior() bool`

HasMergeBehavior returns a boolean if a field has been set.

### GetExpectedTargetContent

`func (o *MergeKeyMergeModesInner) GetExpectedTargetContent() Content2`

GetExpectedTargetContent returns the ExpectedTargetContent field if non-nil, zero value otherwise.

### GetExpectedTargetContentOk

`func (o *MergeKeyMergeModesInner) GetExpectedTargetContentOk() (*Content2, bool)`

GetExpectedTargetContentOk returns a tuple with the ExpectedTargetContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTargetContent

`func (o *MergeKeyMergeModesInner) SetExpectedTargetContent(v Content2)`

SetExpectedTargetContent sets ExpectedTargetContent field to given value.

### HasExpectedTargetContent

`func (o *MergeKeyMergeModesInner) HasExpectedTargetContent() bool`

HasExpectedTargetContent returns a boolean if a field has been set.

### GetResolvedContent

`func (o *MergeKeyMergeModesInner) GetResolvedContent() Content2`

GetResolvedContent returns the ResolvedContent field if non-nil, zero value otherwise.

### GetResolvedContentOk

`func (o *MergeKeyMergeModesInner) GetResolvedContentOk() (*Content2, bool)`

GetResolvedContentOk returns a tuple with the ResolvedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedContent

`func (o *MergeKeyMergeModesInner) SetResolvedContent(v Content2)`

SetResolvedContent sets ResolvedContent field to given value.

### HasResolvedContent

`func (o *MergeKeyMergeModesInner) HasResolvedContent() bool`

HasResolvedContent returns a boolean if a field has been set.

### GetExpectedTargetDocumentation

`func (o *MergeKeyMergeModesInner) GetExpectedTargetDocumentation() ContentResponseV2Documentation`

GetExpectedTargetDocumentation returns the ExpectedTargetDocumentation field if non-nil, zero value otherwise.

### GetExpectedTargetDocumentationOk

`func (o *MergeKeyMergeModesInner) GetExpectedTargetDocumentationOk() (*ContentResponseV2Documentation, bool)`

GetExpectedTargetDocumentationOk returns a tuple with the ExpectedTargetDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTargetDocumentation

`func (o *MergeKeyMergeModesInner) SetExpectedTargetDocumentation(v ContentResponseV2Documentation)`

SetExpectedTargetDocumentation sets ExpectedTargetDocumentation field to given value.

### HasExpectedTargetDocumentation

`func (o *MergeKeyMergeModesInner) HasExpectedTargetDocumentation() bool`

HasExpectedTargetDocumentation returns a boolean if a field has been set.

### GetResolvedDocumentation

`func (o *MergeKeyMergeModesInner) GetResolvedDocumentation() ContentResponseV2Documentation`

GetResolvedDocumentation returns the ResolvedDocumentation field if non-nil, zero value otherwise.

### GetResolvedDocumentationOk

`func (o *MergeKeyMergeModesInner) GetResolvedDocumentationOk() (*ContentResponseV2Documentation, bool)`

GetResolvedDocumentationOk returns a tuple with the ResolvedDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDocumentation

`func (o *MergeKeyMergeModesInner) SetResolvedDocumentation(v ContentResponseV2Documentation)`

SetResolvedDocumentation sets ResolvedDocumentation field to given value.

### HasResolvedDocumentation

`func (o *MergeKeyMergeModesInner) HasResolvedDocumentation() bool`

HasResolvedDocumentation returns a boolean if a field has been set.

### GetMetadata

`func (o *MergeKeyMergeModesInner) GetMetadata() []ContentMetadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MergeKeyMergeModesInner) GetMetadataOk() (*[]ContentMetadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MergeKeyMergeModesInner) SetMetadata(v []ContentMetadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MergeKeyMergeModesInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


