# PutContentOperationForAContentKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | 
**Content** | [**Content2**](Content2.md) |  | 
**ExpectedContent** | Pointer to [**Content3**](Content3.md) |  | [optional] 
**Metadata** | Pointer to [**[]ContentMetadata1**](ContentMetadata1.md) |  | [optional] 
**Documentation** | Pointer to [**PutContentOperationForAContentKeyDocumentation**](PutContentOperationForAContentKeyDocumentation.md) |  | [optional] 

## Methods

### NewPutContentOperationForAContentKey

`func NewPutContentOperationForAContentKey(key GetMultipleContentsRequest1RequestedKeysInner, content Content2, ) *PutContentOperationForAContentKey`

NewPutContentOperationForAContentKey instantiates a new PutContentOperationForAContentKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutContentOperationForAContentKeyWithDefaults

`func NewPutContentOperationForAContentKeyWithDefaults() *PutContentOperationForAContentKey`

NewPutContentOperationForAContentKeyWithDefaults instantiates a new PutContentOperationForAContentKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PutContentOperationForAContentKey) GetKey() GetMultipleContentsRequest1RequestedKeysInner`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PutContentOperationForAContentKey) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PutContentOperationForAContentKey) SetKey(v GetMultipleContentsRequest1RequestedKeysInner)`

SetKey sets Key field to given value.


### GetContent

`func (o *PutContentOperationForAContentKey) GetContent() Content2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PutContentOperationForAContentKey) GetContentOk() (*Content2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PutContentOperationForAContentKey) SetContent(v Content2)`

SetContent sets Content field to given value.


### GetExpectedContent

`func (o *PutContentOperationForAContentKey) GetExpectedContent() Content3`

GetExpectedContent returns the ExpectedContent field if non-nil, zero value otherwise.

### GetExpectedContentOk

`func (o *PutContentOperationForAContentKey) GetExpectedContentOk() (*Content3, bool)`

GetExpectedContentOk returns a tuple with the ExpectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedContent

`func (o *PutContentOperationForAContentKey) SetExpectedContent(v Content3)`

SetExpectedContent sets ExpectedContent field to given value.

### HasExpectedContent

`func (o *PutContentOperationForAContentKey) HasExpectedContent() bool`

HasExpectedContent returns a boolean if a field has been set.

### GetMetadata

`func (o *PutContentOperationForAContentKey) GetMetadata() []ContentMetadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutContentOperationForAContentKey) GetMetadataOk() (*[]ContentMetadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutContentOperationForAContentKey) SetMetadata(v []ContentMetadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PutContentOperationForAContentKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocumentation

`func (o *PutContentOperationForAContentKey) GetDocumentation() PutContentOperationForAContentKeyDocumentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *PutContentOperationForAContentKey) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *PutContentOperationForAContentKey) SetDocumentation(v PutContentOperationForAContentKeyDocumentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *PutContentOperationForAContentKey) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


