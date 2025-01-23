# EntriesResponseV2EntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Declares the type of a Nessie content object, which is currently one of ICEBERG_TABLE, DELTA_LAKE_TABLE, ICEBERG_VIEW, NAMESPACE or UDF, which are the discriminator mapping values of the &#39;Content&#39; type. | [optional] 
**Name** | [**AddedContentKey**](AddedContentKey.md) |  | 
**ContentId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**Content1**](Content1.md) |  | [optional] 

## Methods

### NewEntriesResponseV2EntriesInner

`func NewEntriesResponseV2EntriesInner(name AddedContentKey, ) *EntriesResponseV2EntriesInner`

NewEntriesResponseV2EntriesInner instantiates a new EntriesResponseV2EntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesResponseV2EntriesInnerWithDefaults

`func NewEntriesResponseV2EntriesInnerWithDefaults() *EntriesResponseV2EntriesInner`

NewEntriesResponseV2EntriesInnerWithDefaults instantiates a new EntriesResponseV2EntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntriesResponseV2EntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntriesResponseV2EntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntriesResponseV2EntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntriesResponseV2EntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EntriesResponseV2EntriesInner) GetName() AddedContentKey`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntriesResponseV2EntriesInner) GetNameOk() (*AddedContentKey, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntriesResponseV2EntriesInner) SetName(v AddedContentKey)`

SetName sets Name field to given value.


### GetContentId

`func (o *EntriesResponseV2EntriesInner) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *EntriesResponseV2EntriesInner) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *EntriesResponseV2EntriesInner) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *EntriesResponseV2EntriesInner) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetContent

`func (o *EntriesResponseV2EntriesInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EntriesResponseV2EntriesInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EntriesResponseV2EntriesInner) SetContent(v Content1)`

SetContent sets Content field to given value.

### HasContent

`func (o *EntriesResponseV2EntriesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


