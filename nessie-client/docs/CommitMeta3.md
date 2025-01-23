# CommitMeta3

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

### NewCommitMeta3

`func NewCommitMeta3(authors []string, allSignedOffBy []string, message string, allProperties map[string][]string, parentCommitHashes []string, ) *CommitMeta3`

NewCommitMeta3 instantiates a new CommitMeta3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitMeta3WithDefaults

`func NewCommitMeta3WithDefaults() *CommitMeta3`

NewCommitMeta3WithDefaults instantiates a new CommitMeta3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CommitMeta3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CommitMeta3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CommitMeta3) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CommitMeta3) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetCommitter

`func (o *CommitMeta3) GetCommitter() string`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitMeta3) GetCommitterOk() (*string, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitMeta3) SetCommitter(v string)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *CommitMeta3) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetAuthors

`func (o *CommitMeta3) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CommitMeta3) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CommitMeta3) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.


### GetAllSignedOffBy

`func (o *CommitMeta3) GetAllSignedOffBy() []string`

GetAllSignedOffBy returns the AllSignedOffBy field if non-nil, zero value otherwise.

### GetAllSignedOffByOk

`func (o *CommitMeta3) GetAllSignedOffByOk() (*[]string, bool)`

GetAllSignedOffByOk returns a tuple with the AllSignedOffBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSignedOffBy

`func (o *CommitMeta3) SetAllSignedOffBy(v []string)`

SetAllSignedOffBy sets AllSignedOffBy field to given value.


### GetMessage

`func (o *CommitMeta3) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommitMeta3) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommitMeta3) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCommitTime

`func (o *CommitMeta3) GetCommitTime() time.Time`

GetCommitTime returns the CommitTime field if non-nil, zero value otherwise.

### GetCommitTimeOk

`func (o *CommitMeta3) GetCommitTimeOk() (*time.Time, bool)`

GetCommitTimeOk returns a tuple with the CommitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTime

`func (o *CommitMeta3) SetCommitTime(v time.Time)`

SetCommitTime sets CommitTime field to given value.

### HasCommitTime

`func (o *CommitMeta3) HasCommitTime() bool`

HasCommitTime returns a boolean if a field has been set.

### GetAuthorTime

`func (o *CommitMeta3) GetAuthorTime() time.Time`

GetAuthorTime returns the AuthorTime field if non-nil, zero value otherwise.

### GetAuthorTimeOk

`func (o *CommitMeta3) GetAuthorTimeOk() (*time.Time, bool)`

GetAuthorTimeOk returns a tuple with the AuthorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorTime

`func (o *CommitMeta3) SetAuthorTime(v time.Time)`

SetAuthorTime sets AuthorTime field to given value.

### HasAuthorTime

`func (o *CommitMeta3) HasAuthorTime() bool`

HasAuthorTime returns a boolean if a field has been set.

### GetAllProperties

`func (o *CommitMeta3) GetAllProperties() map[string][]string`

GetAllProperties returns the AllProperties field if non-nil, zero value otherwise.

### GetAllPropertiesOk

`func (o *CommitMeta3) GetAllPropertiesOk() (*map[string][]string, bool)`

GetAllPropertiesOk returns a tuple with the AllProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllProperties

`func (o *CommitMeta3) SetAllProperties(v map[string][]string)`

SetAllProperties sets AllProperties field to given value.


### GetParentCommitHashes

`func (o *CommitMeta3) GetParentCommitHashes() []string`

GetParentCommitHashes returns the ParentCommitHashes field if non-nil, zero value otherwise.

### GetParentCommitHashesOk

`func (o *CommitMeta3) GetParentCommitHashesOk() (*[]string, bool)`

GetParentCommitHashesOk returns a tuple with the ParentCommitHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommitHashes

`func (o *CommitMeta3) SetParentCommitHashes(v []string)`

SetParentCommitHashes sets ParentCommitHashes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


