# Reference2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Hash** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ReferenceMetadata2**](ReferenceMetadata2.md) |  | [optional] 

## Methods

### NewReference2

`func NewReference2(name string, ) *Reference2`

NewReference2 instantiates a new Reference2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReference2WithDefaults

`func NewReference2WithDefaults() *Reference2`

NewReference2WithDefaults instantiates a new Reference2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Reference2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reference2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reference2) SetName(v string)`

SetName sets Name field to given value.


### GetHash

`func (o *Reference2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Reference2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Reference2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Reference2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMetadata

`func (o *Reference2) GetMetadata() ReferenceMetadata2`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Reference2) GetMetadataOk() (*ReferenceMetadata2, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Reference2) SetMetadata(v ReferenceMetadata2)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Reference2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


