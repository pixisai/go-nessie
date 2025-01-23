# GetContentV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content5**](Content5.md) |  | 
**EffectiveReference** | [**Reference3**](Reference3.md) |  | 
**Documentation** | Pointer to [**PutContentOperationForAContentKeyDocumentation**](PutContentOperationForAContentKeyDocumentation.md) |  | [optional] 

## Methods

### NewGetContentV2200Response

`func NewGetContentV2200Response(content Content5, effectiveReference Reference3, ) *GetContentV2200Response`

NewGetContentV2200Response instantiates a new GetContentV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContentV2200ResponseWithDefaults

`func NewGetContentV2200ResponseWithDefaults() *GetContentV2200Response`

NewGetContentV2200ResponseWithDefaults instantiates a new GetContentV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GetContentV2200Response) GetContent() Content5`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetContentV2200Response) GetContentOk() (*Content5, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetContentV2200Response) SetContent(v Content5)`

SetContent sets Content field to given value.


### GetEffectiveReference

`func (o *GetContentV2200Response) GetEffectiveReference() Reference3`

GetEffectiveReference returns the EffectiveReference field if non-nil, zero value otherwise.

### GetEffectiveReferenceOk

`func (o *GetContentV2200Response) GetEffectiveReferenceOk() (*Reference3, bool)`

GetEffectiveReferenceOk returns a tuple with the EffectiveReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveReference

`func (o *GetContentV2200Response) SetEffectiveReference(v Reference3)`

SetEffectiveReference sets EffectiveReference field to given value.


### GetDocumentation

`func (o *GetContentV2200Response) GetDocumentation() PutContentOperationForAContentKeyDocumentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *GetContentV2200Response) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *GetContentV2200Response) SetDocumentation(v PutContentOperationForAContentKeyDocumentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *GetContentV2200Response) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


