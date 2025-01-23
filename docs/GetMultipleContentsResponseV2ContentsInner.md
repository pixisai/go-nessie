# GetMultipleContentsResponseV2ContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AddedContentKey**](AddedContentKey.md) |  | 
**Content** | [**Content1**](Content1.md) |  | 
**Documentation** | Pointer to [**ContentResponseV2Documentation**](ContentResponseV2Documentation.md) |  | [optional] 

## Methods

### NewGetMultipleContentsResponseV2ContentsInner

`func NewGetMultipleContentsResponseV2ContentsInner(key AddedContentKey, content Content1, ) *GetMultipleContentsResponseV2ContentsInner`

NewGetMultipleContentsResponseV2ContentsInner instantiates a new GetMultipleContentsResponseV2ContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultipleContentsResponseV2ContentsInnerWithDefaults

`func NewGetMultipleContentsResponseV2ContentsInnerWithDefaults() *GetMultipleContentsResponseV2ContentsInner`

NewGetMultipleContentsResponseV2ContentsInnerWithDefaults instantiates a new GetMultipleContentsResponseV2ContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetMultipleContentsResponseV2ContentsInner) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMultipleContentsResponseV2ContentsInner) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMultipleContentsResponseV2ContentsInner) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.


### GetContent

`func (o *GetMultipleContentsResponseV2ContentsInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetMultipleContentsResponseV2ContentsInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetMultipleContentsResponseV2ContentsInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetDocumentation

`func (o *GetMultipleContentsResponseV2ContentsInner) GetDocumentation() ContentResponseV2Documentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *GetMultipleContentsResponseV2ContentsInner) GetDocumentationOk() (*ContentResponseV2Documentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *GetMultipleContentsResponseV2ContentsInner) SetDocumentation(v ContentResponseV2Documentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *GetMultipleContentsResponseV2ContentsInner) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


