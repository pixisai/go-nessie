# Operations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMeta** | [**CommitMeta3**](CommitMeta3.md) |  | 
**Operations** | [**[]Operation1**](Operation1.md) |  | 

## Methods

### NewOperations

`func NewOperations(commitMeta CommitMeta3, operations []Operation1, ) *Operations`

NewOperations instantiates a new Operations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationsWithDefaults

`func NewOperationsWithDefaults() *Operations`

NewOperationsWithDefaults instantiates a new Operations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMeta

`func (o *Operations) GetCommitMeta() CommitMeta3`

GetCommitMeta returns the CommitMeta field if non-nil, zero value otherwise.

### GetCommitMetaOk

`func (o *Operations) GetCommitMetaOk() (*CommitMeta3, bool)`

GetCommitMetaOk returns a tuple with the CommitMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMeta

`func (o *Operations) SetCommitMeta(v CommitMeta3)`

SetCommitMeta sets CommitMeta field to given value.


### GetOperations

`func (o *Operations) GetOperations() []Operation1`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Operations) GetOperationsOk() (*[]Operation1, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Operations) SetOperations(v []Operation1)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


