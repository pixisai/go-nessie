# OperationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **map[string]interface{}** |  | 
**Content** | [**Content10**](Content10.md) |  | 
**Metadata** | Pointer to [**[]ContentMetadata1**](ContentMetadata1.md) |  | [optional] 
**Documentation** | Pointer to [**PutContentOperationForAContentKeyDocumentation**](PutContentOperationForAContentKeyDocumentation.md) |  | [optional] 

## Methods

### NewOperationV2

`func NewOperationV2(key map[string]interface{}, content Content10, ) *OperationV2`

NewOperationV2 instantiates a new OperationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationV2WithDefaults

`func NewOperationV2WithDefaults() *OperationV2`

NewOperationV2WithDefaults instantiates a new OperationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OperationV2) GetKey() map[string]interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OperationV2) GetKeyOk() (*map[string]interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OperationV2) SetKey(v map[string]interface{})`

SetKey sets Key field to given value.


### GetContent

`func (o *OperationV2) GetContent() Content10`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OperationV2) GetContentOk() (*Content10, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OperationV2) SetContent(v Content10)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *OperationV2) GetMetadata() []ContentMetadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OperationV2) GetMetadataOk() (*[]ContentMetadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OperationV2) SetMetadata(v []ContentMetadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OperationV2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocumentation

`func (o *OperationV2) GetDocumentation() PutContentOperationForAContentKeyDocumentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *OperationV2) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *OperationV2) SetDocumentation(v PutContentOperationForAContentKeyDocumentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *OperationV2) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


