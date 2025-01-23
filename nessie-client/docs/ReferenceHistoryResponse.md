# ReferenceHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | [**Reference2**](Reference2.md) |  | 
**Current** | [**ReferenceHistoryState3**](ReferenceHistoryState3.md) |  | 
**Previous** | [**[]ReferenceHistoryState4**](ReferenceHistoryState4.md) | Consistency status of the recorded recent HEADs of the reference, including re-assign operations. | 
**CommitLogConsistency** | **string** | Combined consistency status of the commit-log of the reference, if requested by the client. | 

## Methods

### NewReferenceHistoryResponse

`func NewReferenceHistoryResponse(reference Reference2, current ReferenceHistoryState3, previous []ReferenceHistoryState4, commitLogConsistency string, ) *ReferenceHistoryResponse`

NewReferenceHistoryResponse instantiates a new ReferenceHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceHistoryResponseWithDefaults

`func NewReferenceHistoryResponseWithDefaults() *ReferenceHistoryResponse`

NewReferenceHistoryResponseWithDefaults instantiates a new ReferenceHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *ReferenceHistoryResponse) GetReference() Reference2`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ReferenceHistoryResponse) GetReferenceOk() (*Reference2, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ReferenceHistoryResponse) SetReference(v Reference2)`

SetReference sets Reference field to given value.


### GetCurrent

`func (o *ReferenceHistoryResponse) GetCurrent() ReferenceHistoryState3`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ReferenceHistoryResponse) GetCurrentOk() (*ReferenceHistoryState3, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ReferenceHistoryResponse) SetCurrent(v ReferenceHistoryState3)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *ReferenceHistoryResponse) GetPrevious() []ReferenceHistoryState4`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ReferenceHistoryResponse) GetPreviousOk() (*[]ReferenceHistoryState4, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ReferenceHistoryResponse) SetPrevious(v []ReferenceHistoryState4)`

SetPrevious sets Previous field to given value.


### GetCommitLogConsistency

`func (o *ReferenceHistoryResponse) GetCommitLogConsistency() string`

GetCommitLogConsistency returns the CommitLogConsistency field if non-nil, zero value otherwise.

### GetCommitLogConsistencyOk

`func (o *ReferenceHistoryResponse) GetCommitLogConsistencyOk() (*string, bool)`

GetCommitLogConsistencyOk returns a tuple with the CommitLogConsistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitLogConsistency

`func (o *ReferenceHistoryResponse) SetCommitLogConsistency(v string)`

SetCommitLogConsistency sets CommitLogConsistency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


