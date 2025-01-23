# LogEntry1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMeta** | [**CommitMeta3**](CommitMeta3.md) |  | 
**AdditionalParents** | Pointer to **[]string** |  | [optional] 
**ParentCommitHash** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to [**[]Operation1**](Operation1.md) |  | [optional] 

## Methods

### NewLogEntry1

`func NewLogEntry1(commitMeta CommitMeta3, ) *LogEntry1`

NewLogEntry1 instantiates a new LogEntry1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntry1WithDefaults

`func NewLogEntry1WithDefaults() *LogEntry1`

NewLogEntry1WithDefaults instantiates a new LogEntry1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMeta

`func (o *LogEntry1) GetCommitMeta() CommitMeta3`

GetCommitMeta returns the CommitMeta field if non-nil, zero value otherwise.

### GetCommitMetaOk

`func (o *LogEntry1) GetCommitMetaOk() (*CommitMeta3, bool)`

GetCommitMetaOk returns a tuple with the CommitMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMeta

`func (o *LogEntry1) SetCommitMeta(v CommitMeta3)`

SetCommitMeta sets CommitMeta field to given value.


### GetAdditionalParents

`func (o *LogEntry1) GetAdditionalParents() []string`

GetAdditionalParents returns the AdditionalParents field if non-nil, zero value otherwise.

### GetAdditionalParentsOk

`func (o *LogEntry1) GetAdditionalParentsOk() (*[]string, bool)`

GetAdditionalParentsOk returns a tuple with the AdditionalParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParents

`func (o *LogEntry1) SetAdditionalParents(v []string)`

SetAdditionalParents sets AdditionalParents field to given value.

### HasAdditionalParents

`func (o *LogEntry1) HasAdditionalParents() bool`

HasAdditionalParents returns a boolean if a field has been set.

### GetParentCommitHash

`func (o *LogEntry1) GetParentCommitHash() string`

GetParentCommitHash returns the ParentCommitHash field if non-nil, zero value otherwise.

### GetParentCommitHashOk

`func (o *LogEntry1) GetParentCommitHashOk() (*string, bool)`

GetParentCommitHashOk returns a tuple with the ParentCommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommitHash

`func (o *LogEntry1) SetParentCommitHash(v string)`

SetParentCommitHash sets ParentCommitHash field to given value.

### HasParentCommitHash

`func (o *LogEntry1) HasParentCommitHash() bool`

HasParentCommitHash returns a boolean if a field has been set.

### GetOperations

`func (o *LogEntry1) GetOperations() []Operation1`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *LogEntry1) GetOperationsOk() (*[]Operation1, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *LogEntry1) SetOperations(v []Operation1)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *LogEntry1) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


