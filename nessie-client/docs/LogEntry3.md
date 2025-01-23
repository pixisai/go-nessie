# LogEntry3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMeta** | [**CommitMeta2**](CommitMeta2.md) |  | 
**AdditionalParents** | Pointer to **[]string** |  | [optional] 
**ParentCommitHash** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]Operation1**](Operation1.md) |  | [optional] 

## Methods

### NewLogEntry3

`func NewLogEntry3(commitMeta CommitMeta2, ) *LogEntry3`

NewLogEntry3 instantiates a new LogEntry3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntry3WithDefaults

`func NewLogEntry3WithDefaults() *LogEntry3`

NewLogEntry3WithDefaults instantiates a new LogEntry3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMeta

`func (o *LogEntry3) GetCommitMeta() CommitMeta2`

GetCommitMeta returns the CommitMeta field if non-nil, zero value otherwise.

### GetCommitMetaOk

`func (o *LogEntry3) GetCommitMetaOk() (*CommitMeta2, bool)`

GetCommitMetaOk returns a tuple with the CommitMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMeta

`func (o *LogEntry3) SetCommitMeta(v CommitMeta2)`

SetCommitMeta sets CommitMeta field to given value.


### GetAdditionalParents

`func (o *LogEntry3) GetAdditionalParents() []string`

GetAdditionalParents returns the AdditionalParents field if non-nil, zero value otherwise.

### GetAdditionalParentsOk

`func (o *LogEntry3) GetAdditionalParentsOk() (*[]string, bool)`

GetAdditionalParentsOk returns a tuple with the AdditionalParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParents

`func (o *LogEntry3) SetAdditionalParents(v []string)`

SetAdditionalParents sets AdditionalParents field to given value.

### HasAdditionalParents

`func (o *LogEntry3) HasAdditionalParents() bool`

HasAdditionalParents returns a boolean if a field has been set.

### GetParentCommitHash

`func (o *LogEntry3) GetParentCommitHash() string`

GetParentCommitHash returns the ParentCommitHash field if non-nil, zero value otherwise.

### GetParentCommitHashOk

`func (o *LogEntry3) GetParentCommitHashOk() (*string, bool)`

GetParentCommitHashOk returns a tuple with the ParentCommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommitHash

`func (o *LogEntry3) SetParentCommitHash(v string)`

SetParentCommitHash sets ParentCommitHash field to given value.

### HasParentCommitHash

`func (o *LogEntry3) HasParentCommitHash() bool`

HasParentCommitHash returns a boolean if a field has been set.

### GetOperations

`func (o *LogEntry3) GetOperations() []Operation1`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *LogEntry3) GetOperationsOk() (*[]Operation1, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *LogEntry3) SetOperations(v []Operation1)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *LogEntry3) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


