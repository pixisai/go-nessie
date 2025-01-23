# GetMultipleContentsResponse2ContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | 
**Content** | [**Content5**](Content5.md) |  | 
**Documentation** | Pointer to [**PutContentOperationForAContentKeyDocumentation**](PutContentOperationForAContentKeyDocumentation.md) |  | [optional] 

## Methods

### NewGetMultipleContentsResponse2ContentsInner

`func NewGetMultipleContentsResponse2ContentsInner(key GetMultipleContentsRequest1RequestedKeysInner, content Content5, ) *GetMultipleContentsResponse2ContentsInner`

NewGetMultipleContentsResponse2ContentsInner instantiates a new GetMultipleContentsResponse2ContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultipleContentsResponse2ContentsInnerWithDefaults

`func NewGetMultipleContentsResponse2ContentsInnerWithDefaults() *GetMultipleContentsResponse2ContentsInner`

NewGetMultipleContentsResponse2ContentsInnerWithDefaults instantiates a new GetMultipleContentsResponse2ContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetMultipleContentsResponse2ContentsInner) GetKey() GetMultipleContentsRequest1RequestedKeysInner`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMultipleContentsResponse2ContentsInner) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMultipleContentsResponse2ContentsInner) SetKey(v GetMultipleContentsRequest1RequestedKeysInner)`

SetKey sets Key field to given value.


### GetContent

`func (o *GetMultipleContentsResponse2ContentsInner) GetContent() Content5`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetMultipleContentsResponse2ContentsInner) GetContentOk() (*Content5, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetMultipleContentsResponse2ContentsInner) SetContent(v Content5)`

SetContent sets Content field to given value.


### GetDocumentation

`func (o *GetMultipleContentsResponse2ContentsInner) GetDocumentation() PutContentOperationForAContentKeyDocumentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *GetMultipleContentsResponse2ContentsInner) GetDocumentationOk() (*PutContentOperationForAContentKeyDocumentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *GetMultipleContentsResponse2ContentsInner) SetDocumentation(v PutContentOperationForAContentKeyDocumentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *GetMultipleContentsResponse2ContentsInner) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


