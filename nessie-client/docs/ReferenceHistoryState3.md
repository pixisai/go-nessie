# ReferenceHistoryState3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** | Nessie commit ID. | [optional] 
**CommitConsistency** | Pointer to **string** | Consistency status of the commit. | [optional] 
**Meta** | Pointer to [**CommitMeta5**](CommitMeta5.md) |  | [optional] 

## Methods

### NewReferenceHistoryState3

`func NewReferenceHistoryState3() *ReferenceHistoryState3`

NewReferenceHistoryState3 instantiates a new ReferenceHistoryState3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceHistoryState3WithDefaults

`func NewReferenceHistoryState3WithDefaults() *ReferenceHistoryState3`

NewReferenceHistoryState3WithDefaults instantiates a new ReferenceHistoryState3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ReferenceHistoryState3) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ReferenceHistoryState3) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ReferenceHistoryState3) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ReferenceHistoryState3) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitConsistency

`func (o *ReferenceHistoryState3) GetCommitConsistency() string`

GetCommitConsistency returns the CommitConsistency field if non-nil, zero value otherwise.

### GetCommitConsistencyOk

`func (o *ReferenceHistoryState3) GetCommitConsistencyOk() (*string, bool)`

GetCommitConsistencyOk returns a tuple with the CommitConsistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitConsistency

`func (o *ReferenceHistoryState3) SetCommitConsistency(v string)`

SetCommitConsistency sets CommitConsistency field to given value.

### HasCommitConsistency

`func (o *ReferenceHistoryState3) HasCommitConsistency() bool`

HasCommitConsistency returns a boolean if a field has been set.

### GetMeta

`func (o *ReferenceHistoryState3) GetMeta() CommitMeta5`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReferenceHistoryState3) GetMetaOk() (*CommitMeta5, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReferenceHistoryState3) SetMeta(v CommitMeta5)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReferenceHistoryState3) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


