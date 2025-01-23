# Transplant2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Commit message for this transplant request. | [optional] 
**HashesToTransplant** | **[]string** | Lists the hashes of commits that should be transplanted into the target branch. | 
**FromRefName** | **string** | The name of the reference that contains the &#39;source&#39; commits for the requested merge or transplant operation.  | 
**KeyMergeModes** | Pointer to [**[]MergeOperationKeyMergeModesInner**](MergeOperationKeyMergeModesInner.md) | Specific merge behaviour requests by content key.  The default is set by the &#x60;defaultKeyMergeMode&#x60; parameter.  | [optional] 
**DefaultKeyMergeMode** | Pointer to **string** | The default merge mode. If not set, &#x60;NORMAL&#x60; is assumed.  This settings applies to key thaWhen set to &#39;true&#39; instructs the server to validate the request but to avoid committing any changes.t are not explicitly mentioned in the &#x60;keyMergeModes&#x60; property.  | [optional] 
**DryRun** | Pointer to **bool** | When set to &#39;true&#39; instructs the server to validate the request but to avoid committing any changes.  | [optional] 
**FetchAdditionalInfo** | Pointer to **bool** | Whether to provide optional response data.  | [optional] 
**ReturnConflictAsResult** | Pointer to **bool** | When set to &#39;true&#39; instructs the server to produce normal (non-error) responses in case a conflict is detected and report conflict details in the response payload. | [optional] 

## Methods

### NewTransplant2

`func NewTransplant2(hashesToTransplant []string, fromRefName string, ) *Transplant2`

NewTransplant2 instantiates a new Transplant2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransplant2WithDefaults

`func NewTransplant2WithDefaults() *Transplant2`

NewTransplant2WithDefaults instantiates a new Transplant2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Transplant2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Transplant2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Transplant2) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Transplant2) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHashesToTransplant

`func (o *Transplant2) GetHashesToTransplant() []string`

GetHashesToTransplant returns the HashesToTransplant field if non-nil, zero value otherwise.

### GetHashesToTransplantOk

`func (o *Transplant2) GetHashesToTransplantOk() (*[]string, bool)`

GetHashesToTransplantOk returns a tuple with the HashesToTransplant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashesToTransplant

`func (o *Transplant2) SetHashesToTransplant(v []string)`

SetHashesToTransplant sets HashesToTransplant field to given value.


### GetFromRefName

`func (o *Transplant2) GetFromRefName() string`

GetFromRefName returns the FromRefName field if non-nil, zero value otherwise.

### GetFromRefNameOk

`func (o *Transplant2) GetFromRefNameOk() (*string, bool)`

GetFromRefNameOk returns a tuple with the FromRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromRefName

`func (o *Transplant2) SetFromRefName(v string)`

SetFromRefName sets FromRefName field to given value.


### GetKeyMergeModes

`func (o *Transplant2) GetKeyMergeModes() []MergeOperationKeyMergeModesInner`

GetKeyMergeModes returns the KeyMergeModes field if non-nil, zero value otherwise.

### GetKeyMergeModesOk

`func (o *Transplant2) GetKeyMergeModesOk() (*[]MergeOperationKeyMergeModesInner, bool)`

GetKeyMergeModesOk returns a tuple with the KeyMergeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMergeModes

`func (o *Transplant2) SetKeyMergeModes(v []MergeOperationKeyMergeModesInner)`

SetKeyMergeModes sets KeyMergeModes field to given value.

### HasKeyMergeModes

`func (o *Transplant2) HasKeyMergeModes() bool`

HasKeyMergeModes returns a boolean if a field has been set.

### GetDefaultKeyMergeMode

`func (o *Transplant2) GetDefaultKeyMergeMode() string`

GetDefaultKeyMergeMode returns the DefaultKeyMergeMode field if non-nil, zero value otherwise.

### GetDefaultKeyMergeModeOk

`func (o *Transplant2) GetDefaultKeyMergeModeOk() (*string, bool)`

GetDefaultKeyMergeModeOk returns a tuple with the DefaultKeyMergeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyMergeMode

`func (o *Transplant2) SetDefaultKeyMergeMode(v string)`

SetDefaultKeyMergeMode sets DefaultKeyMergeMode field to given value.

### HasDefaultKeyMergeMode

`func (o *Transplant2) HasDefaultKeyMergeMode() bool`

HasDefaultKeyMergeMode returns a boolean if a field has been set.

### GetDryRun

`func (o *Transplant2) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *Transplant2) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *Transplant2) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *Transplant2) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetFetchAdditionalInfo

`func (o *Transplant2) GetFetchAdditionalInfo() bool`

GetFetchAdditionalInfo returns the FetchAdditionalInfo field if non-nil, zero value otherwise.

### GetFetchAdditionalInfoOk

`func (o *Transplant2) GetFetchAdditionalInfoOk() (*bool, bool)`

GetFetchAdditionalInfoOk returns a tuple with the FetchAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchAdditionalInfo

`func (o *Transplant2) SetFetchAdditionalInfo(v bool)`

SetFetchAdditionalInfo sets FetchAdditionalInfo field to given value.

### HasFetchAdditionalInfo

`func (o *Transplant2) HasFetchAdditionalInfo() bool`

HasFetchAdditionalInfo returns a boolean if a field has been set.

### GetReturnConflictAsResult

`func (o *Transplant2) GetReturnConflictAsResult() bool`

GetReturnConflictAsResult returns the ReturnConflictAsResult field if non-nil, zero value otherwise.

### GetReturnConflictAsResultOk

`func (o *Transplant2) GetReturnConflictAsResultOk() (*bool, bool)`

GetReturnConflictAsResultOk returns a tuple with the ReturnConflictAsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnConflictAsResult

`func (o *Transplant2) SetReturnConflictAsResult(v bool)`

SetReturnConflictAsResult sets ReturnConflictAsResult field to given value.

### HasReturnConflictAsResult

`func (o *Transplant2) HasReturnConflictAsResult() bool`

HasReturnConflictAsResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


