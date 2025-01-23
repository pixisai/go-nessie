# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **map[string]interface{}** |  | 
**Content** | [**Content6**](Content6.md) |  | 
**ExpectedContent** | Pointer to [**Content7**](Content7.md) |  | [optional] 
**Metadata** | Pointer to [**[]ContentMetadata1**](ContentMetadata1.md) |  | [optional] 
**Documentation** | Pointer to [**PutContentOperationForAContentKeyDocumentation**](PutContentOperationForAContentKeyDocumentation.md) |  | [optional] 

## Methods

### NewOperation

`func NewOperation(key map[string]interface{}, content Content6, ) *Operation`

NewOperation instantiates a new Operation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationWithDefaults

`func NewOperationWithDefaults() *Operation`

NewOperationWithDefaults instantiates a new Operation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Operation) GetKey() map[string]interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Operation) GetKeyOk() (*map[string]interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Operation) SetKey(v map[string]interface{})`

SetKey sets Key field to given value.


### GetContent

`func (o *Operation) GetContent() Content6`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Operation) GetContentOk() (*Content6, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Operation) SetContent(v Content6)`

SetContent sets Content field to given value.


### GetExpectedContent

`func (o *Operation) GetExpectedContent() Content7`

GetExpectedContent returns the ExpectedContent field if non-nil, zero value otherwise.

### GetExpectedContentOk

`func (o *Operation) GetExpectedContentOk() (*Content7, bool)`

GetExpectedContentOk returns a tuple with the ExpectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedContent

`func (o *Operation) SetExpectedContent(v Content7)`

SetExpectedContent sets ExpectedContent field to given value.

### HasExpectedContent

`func (o *Operation) HasExpectedContent() bool`

HasExpectedContent returns a boolean if a field has been set.

### GetMetadata

`func (o *Operation) GetMetadata() []ContentMetadata1`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Operation) GetMetadataOk() (*[]ContentMetadata1, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Operation) SetMetadata(v []ContentMetadata1)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Operation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDocumentation

`func (o *Operation) GetDocumentation() PutContentOperationForAContentKeyDocumentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Operation) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Operation) SetDocumentation(v PutContentOperationForAContentKeyDocumentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Operation) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


