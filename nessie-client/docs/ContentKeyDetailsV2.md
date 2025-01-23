# ContentKeyDetailsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | [optional] 
**MergeBehavior** | Pointer to **string** |  | [optional] 
**Conflict** | Pointer to [**PerContentKeyConflictDetails**](PerContentKeyConflictDetails.md) |  | [optional] 

## Methods

### NewContentKeyDetailsV2

`func NewContentKeyDetailsV2() *ContentKeyDetailsV2`

NewContentKeyDetailsV2 instantiates a new ContentKeyDetailsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKeyDetailsV2WithDefaults

`func NewContentKeyDetailsV2WithDefaults() *ContentKeyDetailsV2`

NewContentKeyDetailsV2WithDefaults instantiates a new ContentKeyDetailsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContentKeyDetailsV2) GetKey() GetMultipleContentsRequest1RequestedKeysInner`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContentKeyDetailsV2) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContentKeyDetailsV2) SetKey(v GetMultipleContentsRequest1RequestedKeysInner)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContentKeyDetailsV2) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMergeBehavior

`func (o *ContentKeyDetailsV2) GetMergeBehavior() string`

GetMergeBehavior returns the MergeBehavior field if non-nil, zero value otherwise.

### GetMergeBehaviorOk

`func (o *ContentKeyDetailsV2) GetMergeBehaviorOk() (*string, bool)`

GetMergeBehaviorOk returns a tuple with the MergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeBehavior

`func (o *ContentKeyDetailsV2) SetMergeBehavior(v string)`

SetMergeBehavior sets MergeBehavior field to given value.

### HasMergeBehavior

`func (o *ContentKeyDetailsV2) HasMergeBehavior() bool`

HasMergeBehavior returns a boolean if a field has been set.

### GetConflict

`func (o *ContentKeyDetailsV2) GetConflict() PerContentKeyConflictDetails`

GetConflict returns the Conflict field if non-nil, zero value otherwise.

### GetConflictOk

`func (o *ContentKeyDetailsV2) GetConflictOk() (*PerContentKeyConflictDetails, bool)`

GetConflictOk returns a tuple with the Conflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflict

`func (o *ContentKeyDetailsV2) SetConflict(v PerContentKeyConflictDetails)`

SetConflict sets Conflict field to given value.

### HasConflict

`func (o *ContentKeyDetailsV2) HasConflict() bool`

HasConflict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


