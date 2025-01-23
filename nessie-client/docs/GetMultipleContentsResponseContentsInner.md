# GetMultipleContentsResponseContentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AddedContentKey**](AddedContentKey.md) |  | 
**Content** | [**Content2**](Content2.md) |  | 
**Documentation** | Pointer to [**ContentResponseV2Documentation**](ContentResponseV2Documentation.md) |  | [optional] 

## Methods

### NewGetMultipleContentsResponseContentsInner

`func NewGetMultipleContentsResponseContentsInner(key AddedContentKey, content Content2, ) *GetMultipleContentsResponseContentsInner`

NewGetMultipleContentsResponseContentsInner instantiates a new GetMultipleContentsResponseContentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultipleContentsResponseContentsInnerWithDefaults

`func NewGetMultipleContentsResponseContentsInnerWithDefaults() *GetMultipleContentsResponseContentsInner`

NewGetMultipleContentsResponseContentsInnerWithDefaults instantiates a new GetMultipleContentsResponseContentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetMultipleContentsResponseContentsInner) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMultipleContentsResponseContentsInner) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMultipleContentsResponseContentsInner) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.


### GetContent

`func (o *GetMultipleContentsResponseContentsInner) GetContent() Content2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetMultipleContentsResponseContentsInner) GetContentOk() (*Content2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetMultipleContentsResponseContentsInner) SetContent(v Content2)`

SetContent sets Content field to given value.


### GetDocumentation

`func (o *GetMultipleContentsResponseContentsInner) GetDocumentation() ContentResponseV2Documentation`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *GetMultipleContentsResponseContentsInner) GetDocumentationOk() (*ContentResponseV2Documentation, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *GetMultipleContentsResponseContentsInner) SetDocumentation(v ContentResponseV2Documentation)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *GetMultipleContentsResponseContentsInner) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


