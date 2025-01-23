# RepositoryConfigResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]GarbageCollectorConfigObject**](GarbageCollectorConfigObject.md) | The existing configuration objects for the requested types will be returned. Non-existing config objects will not be returned. | [optional] 

## Methods

### NewRepositoryConfigResponseV2

`func NewRepositoryConfigResponseV2() *RepositoryConfigResponseV2`

NewRepositoryConfigResponseV2 instantiates a new RepositoryConfigResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigResponseV2WithDefaults

`func NewRepositoryConfigResponseV2WithDefaults() *RepositoryConfigResponseV2`

NewRepositoryConfigResponseV2WithDefaults instantiates a new RepositoryConfigResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *RepositoryConfigResponseV2) GetConfigs() []GarbageCollectorConfigObject`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *RepositoryConfigResponseV2) GetConfigsOk() (*[]GarbageCollectorConfigObject, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *RepositoryConfigResponseV2) SetConfigs(v []GarbageCollectorConfigObject)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *RepositoryConfigResponseV2) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


