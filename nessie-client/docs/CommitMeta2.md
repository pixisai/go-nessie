# CommitMeta2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**Committer** | Pointer to **string** |  | [optional] 
**Authors** | **[]string** |  | 
**AllSignedOffBy** | **[]string** |  | 
**Message** | **string** |  | 
**CommitTime** | Pointer to **time.Time** |  | [optional] 
**AuthorTime** | Pointer to **time.Time** |  | [optional] 
**AllProperties** | **map[string][]string** |  | 
**ParentCommitHashes** | **[]string** |  | 

## Methods

### NewCommitMeta2

`func NewCommitMeta2(authors []string, allSignedOffBy []string, message string, allProperties map[string][]string, parentCommitHashes []string, ) *CommitMeta2`

NewCommitMeta2 instantiates a new CommitMeta2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitMeta2WithDefaults

`func NewCommitMeta2WithDefaults() *CommitMeta2`

NewCommitMeta2WithDefaults instantiates a new CommitMeta2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CommitMeta2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CommitMeta2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CommitMeta2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CommitMeta2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetCommitter

`func (o *CommitMeta2) GetCommitter() string`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitMeta2) GetCommitterOk() (*string, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitMeta2) SetCommitter(v string)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *CommitMeta2) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetAuthors

`func (o *CommitMeta2) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CommitMeta2) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CommitMeta2) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.


### GetAllSignedOffBy

`func (o *CommitMeta2) GetAllSignedOffBy() []string`

GetAllSignedOffBy returns the AllSignedOffBy field if non-nil, zero value otherwise.

### GetAllSignedOffByOk

`func (o *CommitMeta2) GetAllSignedOffByOk() (*[]string, bool)`

GetAllSignedOffByOk returns a tuple with the AllSignedOffBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSignedOffBy

`func (o *CommitMeta2) SetAllSignedOffBy(v []string)`

SetAllSignedOffBy sets AllSignedOffBy field to given value.


### GetMessage

`func (o *CommitMeta2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitMeta2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitMeta2) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCommitTime

`func (o *CommitMeta2) GetCommitTime() time.Time`

GetCommitTime returns the CommitTime field if non-nil, zero value otherwise.

### GetCommitTimeOk

`func (o *CommitMeta2) GetCommitTimeOk() (*time.Time, bool)`

GetCommitTimeOk returns a tuple with the CommitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTime

`func (o *CommitMeta2) SetCommitTime(v time.Time)`

SetCommitTime sets CommitTime field to given value.

### HasCommitTime

`func (o *CommitMeta2) HasCommitTime() bool`

HasCommitTime returns a boolean if a field has been set.

### GetAuthorTime

`func (o *CommitMeta2) GetAuthorTime() time.Time`

GetAuthorTime returns the AuthorTime field if non-nil, zero value otherwise.

### GetAuthorTimeOk

`func (o *CommitMeta2) GetAuthorTimeOk() (*time.Time, bool)`

GetAuthorTimeOk returns a tuple with the AuthorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTime

`func (o *CommitMeta2) SetAuthorTime(v time.Time)`

SetAuthorTime sets AuthorTime field to given value.

### HasAuthorTime

`func (o *CommitMeta2) HasAuthorTime() bool`

HasAuthorTime returns a boolean if a field has been set.

### GetAllProperties

`func (o *CommitMeta2) GetAllProperties() map[string][]string`

GetAllProperties returns the AllProperties field if non-nil, zero value otherwise.

### GetAllPropertiesOk

`func (o *CommitMeta2) GetAllPropertiesOk() (*map[string][]string, bool)`

GetAllPropertiesOk returns a tuple with the AllProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllProperties

`func (o *CommitMeta2) SetAllProperties(v map[string][]string)`

SetAllProperties sets AllProperties field to given value.


### GetParentCommitHashes

`func (o *CommitMeta2) GetParentCommitHashes() []string`

GetParentCommitHashes returns the ParentCommitHashes field if non-nil, zero value otherwise.

### GetParentCommitHashesOk

`func (o *CommitMeta2) GetParentCommitHashesOk() (*[]string, bool)`

GetParentCommitHashesOk returns a tuple with the ParentCommitHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommitHashes

`func (o *CommitMeta2) SetParentCommitHashes(v []string)`

SetParentCommitHashes sets ParentCommitHashes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


