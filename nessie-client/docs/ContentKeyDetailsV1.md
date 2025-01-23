# ContentKeyDetailsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**GetMultipleContentsRequest1RequestedKeysInner**](GetMultipleContentsRequest1RequestedKeysInner.md) |  | [optional] 
**MergeBehavior** | Pointer to **string** |  | [optional] 

## Methods

### NewContentKeyDetailsV1

`func NewContentKeyDetailsV1() *ContentKeyDetailsV1`

NewContentKeyDetailsV1 instantiates a new ContentKeyDetailsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKeyDetailsV1WithDefaults

`func NewContentKeyDetailsV1WithDefaults() *ContentKeyDetailsV1`

NewContentKeyDetailsV1WithDefaults instantiates a new ContentKeyDetailsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContentKeyDetailsV1) GetKey() GetMultipleContentsRequest1RequestedKeysInner`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContentKeyDetailsV1) GetKeyOk() (*GetMultipleContentsRequest1RequestedKeysInner, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContentKeyDetailsV1) SetKey(v GetMultipleContentsRequest1RequestedKeysInner)`

SetKey sets Key field to given value.

### HasKey

`func (o *ContentKeyDetailsV1) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMergeBehavior

`func (o *ContentKeyDetailsV1) GetMergeBehavior() string`

GetMergeBehavior returns the MergeBehavior field if non-nil, zero value otherwise.

### GetMergeBehaviorOk

`func (o *ContentKeyDetailsV1) GetMergeBehaviorOk() (*string, bool)`

GetMergeBehaviorOk returns a tuple with the MergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeBehavior

`func (o *ContentKeyDetailsV1) SetMergeBehavior(v string)`

SetMergeBehavior sets MergeBehavior field to given value.

### HasMergeBehavior

`func (o *ContentKeyDetailsV1) HasMergeBehavior() bool`

HasMergeBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


