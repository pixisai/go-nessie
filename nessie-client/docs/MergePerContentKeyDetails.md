# MergePerContentKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**AddedContentKey**](AddedContentKey.md) |  | [optional] 
**MergeBehavior** | Pointer to **string** |  | [optional] 
**Conflict** | Pointer to [**PerContentKeyConflictDetails**](PerContentKeyConflictDetails.md) |  | [optional] 

## Methods

### NewMergePerContentKeyDetails

`func NewMergePerContentKeyDetails() *MergePerContentKeyDetails`

NewMergePerContentKeyDetails instantiates a new MergePerContentKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergePerContentKeyDetailsWithDefaults

`func NewMergePerContentKeyDetailsWithDefaults() *MergePerContentKeyDetails`

NewMergePerContentKeyDetailsWithDefaults instantiates a new MergePerContentKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MergePerContentKeyDetails) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MergePerContentKeyDetails) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MergePerContentKeyDetails) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *MergePerContentKeyDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMergeBehavior

`func (o *MergePerContentKeyDetails) GetMergeBehavior() string`

GetMergeBehavior returns the MergeBehavior field if non-nil, zero value otherwise.

### GetMergeBehaviorOk

`func (o *MergePerContentKeyDetails) GetMergeBehaviorOk() (*string, bool)`

GetMergeBehaviorOk returns a tuple with the MergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeBehavior

`func (o *MergePerContentKeyDetails) SetMergeBehavior(v string)`

SetMergeBehavior sets MergeBehavior field to given value.

### HasMergeBehavior

`func (o *MergePerContentKeyDetails) HasMergeBehavior() bool`

HasMergeBehavior returns a boolean if a field has been set.

### GetConflict

`func (o *MergePerContentKeyDetails) GetConflict() PerContentKeyConflictDetails`

GetConflict returns the Conflict field if non-nil, zero value otherwise.

### GetConflictOk

`func (o *MergePerContentKeyDetails) GetConflictOk() (*PerContentKeyConflictDetails, bool)`

GetConflictOk returns a tuple with the Conflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflict

`func (o *MergePerContentKeyDetails) SetConflict(v PerContentKeyConflictDetails)`

SetConflict sets Conflict field to given value.

### HasConflict

`func (o *MergePerContentKeyDetails) HasConflict() bool`

HasConflict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


