# GetEntriesV2200ResponseEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Declares the type of a Nessie content object, which is currently one of ICEBERG_TABLE, DELTA_LAKE_TABLE, ICEBERG_VIEW, NAMESPACE or UDF, which are the discriminator mapping values of the &#39;Content&#39; type. | [optional] 
**Name** | [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | 
**ContentId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**Content5**](Content5.md) |  | [optional] 

## Methods

### NewGetEntriesV2200ResponseEntriesInner

`func NewGetEntriesV2200ResponseEntriesInner(name GetMultipleContentsRequest1RequestedKeysInner, ) *GetEntriesV2200ResponseEntriesInner`

NewGetEntriesV2200ResponseEntriesInner instantiates a new GetEntriesV2200ResponseEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntriesV2200ResponseEntriesInnerWithDefaults

`func NewGetEntriesV2200ResponseEntriesInnerWithDefaults() *GetEntriesV2200ResponseEntriesInner`

NewGetEntriesV2200ResponseEntriesInnerWithDefaults instantiates a new GetEntriesV2200ResponseEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetEntriesV2200ResponseEntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEntriesV2200ResponseEntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEntriesV2200ResponseEntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEntriesV2200ResponseEntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetEntriesV2200ResponseEntriesInner) GetName() GetMultipleContentsRequest1RequestedKeysInner`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEntriesV2200ResponseEntriesInner) GetNameOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEntriesV2200ResponseEntriesInner) SetName(v GetMultipleContentsRequest1RequestedKeysInner)`

SetName sets Name field to given value.


### GetContentId

`func (o *GetEntriesV2200ResponseEntriesInner) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *GetEntriesV2200ResponseEntriesInner) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *GetEntriesV2200ResponseEntriesInner) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *GetEntriesV2200ResponseEntriesInner) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetContent

`func (o *GetEntriesV2200ResponseEntriesInner) GetContent() Content5`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetEntriesV2200ResponseEntriesInner) GetContentOk() (*Content5, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetEntriesV2200ResponseEntriesInner) SetContent(v Content5)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetEntriesV2200ResponseEntriesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


