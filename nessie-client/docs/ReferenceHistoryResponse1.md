# ReferenceHistoryResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | [**Reference3**](Reference3.md) |  | 
**Current** | [**ReferenceHistoryState1**](ReferenceHistoryState1.md) |  | 
**Previous** | [**[]ReferenceHistoryState2**](ReferenceHistoryState2.md) | Consistency status of the recorded recent HEADs of the reference, including re-assign operations. | 
**CommitLogConsistency** | **string** | Combined consistency status of the commit-log of the reference, if requested by the client. | 

## Methods

### NewReferenceHistoryResponse1

`func NewReferenceHistoryResponse1(reference Reference3, current ReferenceHistoryState1, previous []ReferenceHistoryState2, commitLogConsistency string, ) *ReferenceHistoryResponse1`

NewReferenceHistoryResponse1 instantiates a new ReferenceHistoryResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceHistoryResponse1WithDefaults

`func NewReferenceHistoryResponse1WithDefaults() *ReferenceHistoryResponse1`

NewReferenceHistoryResponse1WithDefaults instantiates a new ReferenceHistoryResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *ReferenceHistoryResponse1) GetReference() Reference3`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ReferenceHistoryResponse1) GetReferenceOk() (*Reference3, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ReferenceHistoryResponse1) SetReference(v Reference3)`

SetReference sets Reference field to given value.


### GetCurrent

`func (o *ReferenceHistoryResponse1) GetCurrent() ReferenceHistoryState1`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ReferenceHistoryResponse1) GetCurrentOk() (*ReferenceHistoryState1, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ReferenceHistoryResponse1) SetCurrent(v ReferenceHistoryState1)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *ReferenceHistoryResponse1) GetPrevious() []ReferenceHistoryState2`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ReferenceHistoryResponse1) GetPreviousOk() (*[]ReferenceHistoryState2, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ReferenceHistoryResponse1) SetPrevious(v []ReferenceHistoryState2)`

SetPrevious sets Previous field to given value.


### GetCommitLogConsistency

`func (o *ReferenceHistoryResponse1) GetCommitLogConsistency() string`

GetCommitLogConsistency returns the CommitLogConsistency field if non-nil, zero value otherwise.

### GetCommitLogConsistencyOk

`func (o *ReferenceHistoryResponse1) GetCommitLogConsistencyOk() (*string, bool)`

GetCommitLogConsistencyOk returns a tuple with the CommitLogConsistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitLogConsistency

`func (o *ReferenceHistoryResponse1) SetCommitLogConsistency(v string)`

SetCommitLogConsistency sets CommitLogConsistency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


