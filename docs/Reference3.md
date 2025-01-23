# Reference3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Hash** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ReferenceMetadata3**](ReferenceMetadata3.md) |  | [optional] 

## Methods

### NewReference3

`func NewReference3(name string, ) *Reference3`

NewReference3 instantiates a new Reference3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReference3WithDefaults

`func NewReference3WithDefaults() *Reference3`

NewReference3WithDefaults instantiates a new Reference3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Reference3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reference3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reference3) SetName(v string)`

SetName sets Name field to given value.


### GetHash

`func (o *Reference3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Reference3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Reference3) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Reference3) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMetadata

`func (o *Reference3) GetMetadata() ReferenceMetadata3`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Reference3) GetMetadataOk() (*ReferenceMetadata3, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Reference3) SetMetadata(v ReferenceMetadata3)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Reference3) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


