# GetEntries200ResponseEntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Declares the type of a Nessie content object, which is currently one of ICEBERG_TABLE, DELTA_LAKE_TABLE, ICEBERG_VIEW, NAMESPACE or UDF, which are the discriminator mapping values of the &#39;Content&#39; type. | [optional] 
**Name** | [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | 

## Methods

### NewGetEntries200ResponseEntriesInner

`func NewGetEntries200ResponseEntriesInner(name GetMultipleContentsRequest1RequestedKeysInner, ) *GetEntries200ResponseEntriesInner`

NewGetEntries200ResponseEntriesInner instantiates a new GetEntries200ResponseEntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntries200ResponseEntriesInnerWithDefaults

`func NewGetEntries200ResponseEntriesInnerWithDefaults() *GetEntries200ResponseEntriesInner`

NewGetEntries200ResponseEntriesInnerWithDefaults instantiates a new GetEntries200ResponseEntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetEntries200ResponseEntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEntries200ResponseEntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEntries200ResponseEntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEntries200ResponseEntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetEntries200ResponseEntriesInner) GetName() GetMultipleContentsRequest1RequestedKeysInner`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEntries200ResponseEntriesInner) GetNameOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEntries200ResponseEntriesInner) SetName(v GetMultipleContentsRequest1RequestedKeysInner)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


