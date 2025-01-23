# CommitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetBranch** | [**Branch3**](Branch3.md) |  | 
**AddedContents** | Pointer to [**[]CommitResponseAddedContentsInner**](CommitResponseAddedContentsInner.md) |  | [optional] 

## Methods

### NewCommitResponse

`func NewCommitResponse(targetBranch Branch3, ) *CommitResponse`

NewCommitResponse instantiates a new CommitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitResponseWithDefaults

`func NewCommitResponseWithDefaults() *CommitResponse`

NewCommitResponseWithDefaults instantiates a new CommitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetBranch

`func (o *CommitResponse) GetTargetBranch() Branch3`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *CommitResponse) GetTargetBranchOk() (*Branch3, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *CommitResponse) SetTargetBranch(v Branch3)`

SetTargetBranch sets TargetBranch field to given value.


### GetAddedContents

`func (o *CommitResponse) GetAddedContents() []CommitResponseAddedContentsInner`

GetAddedContents returns the AddedContents field if non-nil, zero value otherwise.

### GetAddedContentsOk

`func (o *CommitResponse) GetAddedContentsOk() (*[]CommitResponseAddedContentsInner, bool)`

GetAddedContentsOk returns a tuple with the AddedContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedContents

`func (o *CommitResponse) SetAddedContents(v []CommitResponseAddedContentsInner)`

SetAddedContents sets AddedContents field to given value.

### HasAddedContents

`func (o *CommitResponse) HasAddedContents() bool`

HasAddedContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


