# MergeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromRefName** | **string** |  | 
**FromHash** | **string** |  | 
**KeyMergeModes** | Pointer to [**[]MergeOperationKeyMergeModesInner**](MergeOperationKeyMergeModesInner.md) |  | [optional] 
**DefaultKeyMergeMode** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** |  | [optional] 
**FetchAdditionalInfo** | Pointer to **bool** |  | [optional] 
**ReturnConflictAsResult** | Pointer to **bool** |  | [optional] 

## Methods

### NewMergeOperation

`func NewMergeOperation(fromRefName string, fromHash string, ) *MergeOperation`

NewMergeOperation instantiates a new MergeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeOperationWithDefaults

`func NewMergeOperationWithDefaults() *MergeOperation`

NewMergeOperationWithDefaults instantiates a new MergeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromRefName

`func (o *MergeOperation) GetFromRefName() string`

GetFromRefName returns the FromRefName field if non-nil, zero value otherwise.

### GetFromRefNameOk

`func (o *MergeOperation) GetFromRefNameOk() (*string, bool)`

GetFromRefNameOk returns a tuple with the FromRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRefName

`func (o *MergeOperation) SetFromRefName(v string)`

SetFromRefName sets FromRefName field to given value.


### GetFromHash

`func (o *MergeOperation) GetFromHash() string`

GetFromHash returns the FromHash field if non-nil, zero value otherwise.

### GetFromHashOk

`func (o *MergeOperation) GetFromHashOk() (*string, bool)`

GetFromHashOk returns a tuple with the FromHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromHash

`func (o *MergeOperation) SetFromHash(v string)`

SetFromHash sets FromHash field to given value.


### GetKeyMergeModes

`func (o *MergeOperation) GetKeyMergeModes() []MergeOperationKeyMergeModesInner`

GetKeyMergeModes returns the KeyMergeModes field if non-nil, zero value otherwise.

### GetKeyMergeModesOk

`func (o *MergeOperation) GetKeyMergeModesOk() (*[]MergeOperationKeyMergeModesInner, bool)`

GetKeyMergeModesOk returns a tuple with the KeyMergeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMergeModes

`func (o *MergeOperation) SetKeyMergeModes(v []MergeOperationKeyMergeModesInner)`

SetKeyMergeModes sets KeyMergeModes field to given value.

### HasKeyMergeModes

`func (o *MergeOperation) HasKeyMergeModes() bool`

HasKeyMergeModes returns a boolean if a field has been set.

### GetDefaultKeyMergeMode

`func (o *MergeOperation) GetDefaultKeyMergeMode() string`

GetDefaultKeyMergeMode returns the DefaultKeyMergeMode field if non-nil, zero value otherwise.

### GetDefaultKeyMergeModeOk

`func (o *MergeOperation) GetDefaultKeyMergeModeOk() (*string, bool)`

GetDefaultKeyMergeModeOk returns a tuple with the DefaultKeyMergeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyMergeMode

`func (o *MergeOperation) SetDefaultKeyMergeMode(v string)`

SetDefaultKeyMergeMode sets DefaultKeyMergeMode field to given value.

### HasDefaultKeyMergeMode

`func (o *MergeOperation) HasDefaultKeyMergeMode() bool`

HasDefaultKeyMergeMode returns a boolean if a field has been set.

### GetDryRun

`func (o *MergeOperation) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MergeOperation) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MergeOperation) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MergeOperation) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFetchAdditionalInfo

`func (o *MergeOperation) GetFetchAdditionalInfo() bool`

GetFetchAdditionalInfo returns the FetchAdditionalInfo field if non-nil, zero value otherwise.

### GetFetchAdditionalInfoOk

`func (o *MergeOperation) GetFetchAdditionalInfoOk() (*bool, bool)`

GetFetchAdditionalInfoOk returns a tuple with the FetchAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchAdditionalInfo

`func (o *MergeOperation) SetFetchAdditionalInfo(v bool)`

SetFetchAdditionalInfo sets FetchAdditionalInfo field to given value.

### HasFetchAdditionalInfo

`func (o *MergeOperation) HasFetchAdditionalInfo() bool`

HasFetchAdditionalInfo returns a boolean if a field has been set.

### GetReturnConflictAsResult

`func (o *MergeOperation) GetReturnConflictAsResult() bool`

GetReturnConflictAsResult returns the ReturnConflictAsResult field if non-nil, zero value otherwise.

### GetReturnConflictAsResultOk

`func (o *MergeOperation) GetReturnConflictAsResultOk() (*bool, bool)`

GetReturnConflictAsResultOk returns a tuple with the ReturnConflictAsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnConflictAsResult

`func (o *MergeOperation) SetReturnConflictAsResult(v bool)`

SetReturnConflictAsResult sets ReturnConflictAsResult field to given value.

### HasReturnConflictAsResult

`func (o *MergeOperation) HasReturnConflictAsResult() bool`

HasReturnConflictAsResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


