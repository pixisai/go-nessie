# RepositoryConfigV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultCutoffPolicy** | Pointer to **string** | The default cutoff policy. Policies can be one of: - number of commits as an integer value - a duration (see java.time.Duration) - an ISO instant - &#39;NONE&#39;, means everything&#39;s considered as live | [optional] 
**PerRefCutoffPolicies** | Pointer to [**[]ReferencesCutoffPolicy**](ReferencesCutoffPolicy.md) |  | [optional] 
**NewFilesGracePeriod** | Pointer to **string** | Files that have been created after &#39;gc-start-time - new-files-grace-period&#39; are not being deleted. | [optional] 
**ExpectedFileCountPerContent** | Pointer to **int32** |  | [optional] 

## Methods

### NewRepositoryConfigV2

`func NewRepositoryConfigV2() *RepositoryConfigV2`

NewRepositoryConfigV2 instantiates a new RepositoryConfigV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigV2WithDefaults

`func NewRepositoryConfigV2WithDefaults() *RepositoryConfigV2`

NewRepositoryConfigV2WithDefaults instantiates a new RepositoryConfigV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultCutoffPolicy

`func (o *RepositoryConfigV2) GetDefaultCutoffPolicy() string`

GetDefaultCutoffPolicy returns the DefaultCutoffPolicy field if non-nil, zero value otherwise.

### GetDefaultCutoffPolicyOk

`func (o *RepositoryConfigV2) GetDefaultCutoffPolicyOk() (*string, bool)`

GetDefaultCutoffPolicyOk returns a tuple with the DefaultCutoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCutoffPolicy

`func (o *RepositoryConfigV2) SetDefaultCutoffPolicy(v string)`

SetDefaultCutoffPolicy sets DefaultCutoffPolicy field to given value.

### HasDefaultCutoffPolicy

`func (o *RepositoryConfigV2) HasDefaultCutoffPolicy() bool`

HasDefaultCutoffPolicy returns a boolean if a field has been set.

### GetPerRefCutoffPolicies

`func (o *RepositoryConfigV2) GetPerRefCutoffPolicies() []ReferencesCutoffPolicy`

GetPerRefCutoffPolicies returns the PerRefCutoffPolicies field if non-nil, zero value otherwise.

### GetPerRefCutoffPoliciesOk

`func (o *RepositoryConfigV2) GetPerRefCutoffPoliciesOk() (*[]ReferencesCutoffPolicy, bool)`

GetPerRefCutoffPoliciesOk returns a tuple with the PerRefCutoffPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerRefCutoffPolicies

`func (o *RepositoryConfigV2) SetPerRefCutoffPolicies(v []ReferencesCutoffPolicy)`

SetPerRefCutoffPolicies sets PerRefCutoffPolicies field to given value.

### HasPerRefCutoffPolicies

`func (o *RepositoryConfigV2) HasPerRefCutoffPolicies() bool`

HasPerRefCutoffPolicies returns a boolean if a field has been set.

### GetNewFilesGracePeriod

`func (o *RepositoryConfigV2) GetNewFilesGracePeriod() string`

GetNewFilesGracePeriod returns the NewFilesGracePeriod field if non-nil, zero value otherwise.

### GetNewFilesGracePeriodOk

`func (o *RepositoryConfigV2) GetNewFilesGracePeriodOk() (*string, bool)`

GetNewFilesGracePeriodOk returns a tuple with the NewFilesGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFilesGracePeriod

`func (o *RepositoryConfigV2) SetNewFilesGracePeriod(v string)`

SetNewFilesGracePeriod sets NewFilesGracePeriod field to given value.

### HasNewFilesGracePeriod

`func (o *RepositoryConfigV2) HasNewFilesGracePeriod() bool`

HasNewFilesGracePeriod returns a boolean if a field has been set.

### GetExpectedFileCountPerContent

`func (o *RepositoryConfigV2) GetExpectedFileCountPerContent() int32`

GetExpectedFileCountPerContent returns the ExpectedFileCountPerContent field if non-nil, zero value otherwise.

### GetExpectedFileCountPerContentOk

`func (o *RepositoryConfigV2) GetExpectedFileCountPerContentOk() (*int32, bool)`

GetExpectedFileCountPerContentOk returns a tuple with the ExpectedFileCountPerContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedFileCountPerContent

`func (o *RepositoryConfigV2) SetExpectedFileCountPerContent(v int32)`

SetExpectedFileCountPerContent sets ExpectedFileCountPerContent field to given value.

### HasExpectedFileCountPerContent

`func (o *RepositoryConfigV2) HasExpectedFileCountPerContent() bool`

HasExpectedFileCountPerContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


