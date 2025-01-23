# PutContentOperationForAContentKey5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AddedContentKey**](AddedContentKey.md) |  | 
**Content** | [**Content10**](Content10.md) |  | 
**Metadata** | Pointer to [**[]ContentMetadata1**](ContentMetadata1.md) |  | [optional] 
**Documentation** | Pointer to [**ContentResponseV2Documentation**](ContentResponseV2Documentation.md) |  | [optional] 

## Methods

### NewPutContentOperationForAContentKey5

`func NewPutContentOperationForAContentKey5(key AddedContentKey, content Content10, ) *PutContentOperationForAContentKey5`

NewPutContentOperationForAContentKey5 instantiates a new PutContentOperationForAContentKey5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutContentOperationForAContentKey5WithDefaults

`func NewPutContentOperationForAContentKey5WithDefaults() *PutContentOperationForAContentKey5`

NewPutContentOperationForAContentKey5WithDefaults instantiates a new PutContentOperationForAContentKey5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PutContentOperationForAContentKey5) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PutContentOperationForAContentKey5) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PutContentOperationForAContentKey5) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.


### GetContent

`func (o *PutContentOperationForAContentKey5) GetContent() Content10`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PutContentOperationForAContentKey5) GetContentOk() (*Content10, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PutContentOperationForAContentKey5) SetContent(v Content10)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *PutContentOperationForAContentKey5) GetMetadata() []ContentMetadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutContentOperationForAContentKey5) GetMetadataOk() (*[]ContentMetadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutContentOperationForAContentKey5) SetMetadata(v []ContentMetadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PutContentOperationForAContentKey5) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocumentation

`func (o *PutContentOperationForAContentKey5) GetDocumentation() ContentResponseV2Documentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *PutContentOperationForAContentKey5) GetDocumentationOk() (*ContentResponseV2Documentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *PutContentOperationForAContentKey5) SetDocumentation(v ContentResponseV2Documentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *PutContentOperationForAContentKey5) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


