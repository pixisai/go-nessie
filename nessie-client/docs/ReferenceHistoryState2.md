# ReferenceHistoryState2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | Pointer to **string** | Nessie commit ID. | [optional] 
**CommitConsistency** | Pointer to **string** | Consistency status of the commit. | [optional] 
**Meta** | Pointer to [**CommitMeta4**](CommitMeta4.md) |  | [optional] 

## Methods

### NewReferenceHistoryState2

`func NewReferenceHistoryState2() *ReferenceHistoryState2`

NewReferenceHistoryState2 instantiates a new ReferenceHistoryState2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceHistoryState2WithDefaults

`func NewReferenceHistoryState2WithDefaults() *ReferenceHistoryState2`

NewReferenceHistoryState2WithDefaults instantiates a new ReferenceHistoryState2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ReferenceHistoryState2) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ReferenceHistoryState2) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ReferenceHistoryState2) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ReferenceHistoryState2) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommitConsistency

`func (o *ReferenceHistoryState2) GetCommitConsistency() string`

GetCommitConsistency returns the CommitConsistency field if non-nil, zero value otherwise.

### GetCommitConsistencyOk

`func (o *ReferenceHistoryState2) GetCommitConsistencyOk() (*string, bool)`

GetCommitConsistencyOk returns a tuple with the CommitConsistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitConsistency

`func (o *ReferenceHistoryState2) SetCommitConsistency(v string)`

SetCommitConsistency sets CommitConsistency field to given value.

### HasCommitConsistency

`func (o *ReferenceHistoryState2) HasCommitConsistency() bool`

HasCommitConsistency returns a boolean if a field has been set.

### GetMeta

`func (o *ReferenceHistoryState2) GetMeta() CommitMeta4`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ReferenceHistoryState2) GetMetaOk() (*CommitMeta4, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ReferenceHistoryState2) SetMeta(v CommitMeta4)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ReferenceHistoryState2) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


