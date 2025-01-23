# Branch3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Metadata** | Pointer to [**ReferenceMetadata3**](ReferenceMetadata3.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewBranch3

`func NewBranch3(name string, ) *Branch3`

NewBranch3 instantiates a new Branch3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranch3WithDefaults

`func NewBranch3WithDefaults() *Branch3`

NewBranch3WithDefaults instantiates a new Branch3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Branch3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch3) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *Branch3) GetMetadata() ReferenceMetadata3`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Branch3) GetMetadataOk() (*ReferenceMetadata3, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Branch3) SetMetadata(v ReferenceMetadata3)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Branch3) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetHash

`func (o *Branch3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Branch3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Branch3) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Branch3) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


