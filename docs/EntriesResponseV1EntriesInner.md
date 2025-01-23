# EntriesResponseV1EntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Declares the type of a Nessie content object, which is currently one of ICEBERG_TABLE, DELTA_LAKE_TABLE, ICEBERG_VIEW, NAMESPACE or UDF, which are the discriminator mapping values of the &#39;Content&#39; type. | [optional] 
**Name** | [**AddedContentKey**](AddedContentKey.md) |  | 

## Methods

### NewEntriesResponseV1EntriesInner

`func NewEntriesResponseV1EntriesInner(name AddedContentKey, ) *EntriesResponseV1EntriesInner`

NewEntriesResponseV1EntriesInner instantiates a new EntriesResponseV1EntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesResponseV1EntriesInnerWithDefaults

`func NewEntriesResponseV1EntriesInnerWithDefaults() *EntriesResponseV1EntriesInner`

NewEntriesResponseV1EntriesInnerWithDefaults instantiates a new EntriesResponseV1EntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntriesResponseV1EntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntriesResponseV1EntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntriesResponseV1EntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntriesResponseV1EntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EntriesResponseV1EntriesInner) GetName() AddedContentKey`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntriesResponseV1EntriesInner) GetNameOk() (*AddedContentKey, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntriesResponseV1EntriesInner) SetName(v AddedContentKey)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


