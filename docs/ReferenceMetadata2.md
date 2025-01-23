# ReferenceMetadata2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumCommitsAhead** | Pointer to **int32** |  | [optional] 
**NumCommitsBehind** | Pointer to **int32** |  | [optional] 
**CommitMetaOfHEAD** | Pointer to [**CommitMeta2**](CommitMeta2.md) |  | [optional] 
**CommonAncestorHash** | Pointer to **string** |  | [optional] 
**NumTotalCommits** | Pointer to **int64** |  | [optional] 

## Methods

### NewReferenceMetadata2

`func NewReferenceMetadata2() *ReferenceMetadata2`

NewReferenceMetadata2 instantiates a new ReferenceMetadata2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceMetadata2WithDefaults

`func NewReferenceMetadata2WithDefaults() *ReferenceMetadata2`

NewReferenceMetadata2WithDefaults instantiates a new ReferenceMetadata2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumCommitsAhead

`func (o *ReferenceMetadata2) GetNumCommitsAhead() int32`

GetNumCommitsAhead returns the NumCommitsAhead field if non-nil, zero value otherwise.

### GetNumCommitsAheadOk

`func (o *ReferenceMetadata2) GetNumCommitsAheadOk() (*int32, bool)`

GetNumCommitsAheadOk returns a tuple with the NumCommitsAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCommitsAhead

`func (o *ReferenceMetadata2) SetNumCommitsAhead(v int32)`

SetNumCommitsAhead sets NumCommitsAhead field to given value.

### HasNumCommitsAhead

`func (o *ReferenceMetadata2) HasNumCommitsAhead() bool`

HasNumCommitsAhead returns a boolean if a field has been set.

### GetNumCommitsBehind

`func (o *ReferenceMetadata2) GetNumCommitsBehind() int32`

GetNumCommitsBehind returns the NumCommitsBehind field if non-nil, zero value otherwise.

### GetNumCommitsBehindOk

`func (o *ReferenceMetadata2) GetNumCommitsBehindOk() (*int32, bool)`

GetNumCommitsBehindOk returns a tuple with the NumCommitsBehind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCommitsBehind

`func (o *ReferenceMetadata2) SetNumCommitsBehind(v int32)`

SetNumCommitsBehind sets NumCommitsBehind field to given value.

### HasNumCommitsBehind

`func (o *ReferenceMetadata2) HasNumCommitsBehind() bool`

HasNumCommitsBehind returns a boolean if a field has been set.

### GetCommitMetaOfHEAD

`func (o *ReferenceMetadata2) GetCommitMetaOfHEAD() CommitMeta2`

GetCommitMetaOfHEAD returns the CommitMetaOfHEAD field if non-nil, zero value otherwise.

### GetCommitMetaOfHEADOk

`func (o *ReferenceMetadata2) GetCommitMetaOfHEADOk() (*CommitMeta2, bool)`

GetCommitMetaOfHEADOk returns a tuple with the CommitMetaOfHEAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMetaOfHEAD

`func (o *ReferenceMetadata2) SetCommitMetaOfHEAD(v CommitMeta2)`

SetCommitMetaOfHEAD sets CommitMetaOfHEAD field to given value.

### HasCommitMetaOfHEAD

`func (o *ReferenceMetadata2) HasCommitMetaOfHEAD() bool`

HasCommitMetaOfHEAD returns a boolean if a field has been set.

### GetCommonAncestorHash

`func (o *ReferenceMetadata2) GetCommonAncestorHash() string`

GetCommonAncestorHash returns the CommonAncestorHash field if non-nil, zero value otherwise.

### GetCommonAncestorHashOk

`func (o *ReferenceMetadata2) GetCommonAncestorHashOk() (*string, bool)`

GetCommonAncestorHashOk returns a tuple with the CommonAncestorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonAncestorHash

`func (o *ReferenceMetadata2) SetCommonAncestorHash(v string)`

SetCommonAncestorHash sets CommonAncestorHash field to given value.

### HasCommonAncestorHash

`func (o *ReferenceMetadata2) HasCommonAncestorHash() bool`

HasCommonAncestorHash returns a boolean if a field has been set.

### GetNumTotalCommits

`func (o *ReferenceMetadata2) GetNumTotalCommits() int64`

GetNumTotalCommits returns the NumTotalCommits field if non-nil, zero value otherwise.

### GetNumTotalCommitsOk

`func (o *ReferenceMetadata2) GetNumTotalCommitsOk() (*int64, bool)`

GetNumTotalCommitsOk returns a tuple with the NumTotalCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalCommits

`func (o *ReferenceMetadata2) SetNumTotalCommits(v int64)`

SetNumTotalCommits sets NumTotalCommits field to given value.

### HasNumTotalCommits

`func (o *ReferenceMetadata2) HasNumTotalCommits() bool`

HasNumTotalCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


