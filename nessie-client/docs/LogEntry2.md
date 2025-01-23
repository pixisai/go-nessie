# LogEntry2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMeta** | [**CommitMeta1**](CommitMeta1.md) |  | 
**AdditionalParents** | Pointer to **[]string** |  | [optional] 
**ParentCommitHash** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]Operation2**](Operation2.md) |  | [optional] 

## Methods

### NewLogEntry2

`func NewLogEntry2(commitMeta CommitMeta1, ) *LogEntry2`

NewLogEntry2 instantiates a new LogEntry2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntry2WithDefaults

`func NewLogEntry2WithDefaults() *LogEntry2`

NewLogEntry2WithDefaults instantiates a new LogEntry2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMeta

`func (o *LogEntry2) GetCommitMeta() CommitMeta1`

GetCommitMeta returns the CommitMeta field if non-nil, zero value otherwise.

### GetCommitMetaOk

`func (o *LogEntry2) GetCommitMetaOk() (*CommitMeta1, bool)`

GetCommitMetaOk returns a tuple with the CommitMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMeta

`func (o *LogEntry2) SetCommitMeta(v CommitMeta1)`

SetCommitMeta sets CommitMeta field to given value.


### GetAdditionalParents

`func (o *LogEntry2) GetAdditionalParents() []string`

GetAdditionalParents returns the AdditionalParents field if non-nil, zero value otherwise.

### GetAdditionalParentsOk

`func (o *LogEntry2) GetAdditionalParentsOk() (*[]string, bool)`

GetAdditionalParentsOk returns a tuple with the AdditionalParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParents

`func (o *LogEntry2) SetAdditionalParents(v []string)`

SetAdditionalParents sets AdditionalParents field to given value.

### HasAdditionalParents

`func (o *LogEntry2) HasAdditionalParents() bool`

HasAdditionalParents returns a boolean if a field has been set.

### GetParentCommitHash

`func (o *LogEntry2) GetParentCommitHash() string`

GetParentCommitHash returns the ParentCommitHash field if non-nil, zero value otherwise.

### GetParentCommitHashOk

`func (o *LogEntry2) GetParentCommitHashOk() (*string, bool)`

GetParentCommitHashOk returns a tuple with the ParentCommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommitHash

`func (o *LogEntry2) SetParentCommitHash(v string)`

SetParentCommitHash sets ParentCommitHash field to given value.

### HasParentCommitHash

`func (o *LogEntry2) HasParentCommitHash() bool`

HasParentCommitHash returns a boolean if a field has been set.

### GetOperations

`func (o *LogEntry2) GetOperations() []Operation2`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *LogEntry2) GetOperationsOk() (*[]Operation2, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *LogEntry2) SetOperations(v []Operation2)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *LogEntry2) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


