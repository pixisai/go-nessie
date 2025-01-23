# PutV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AddedContentKey**](AddedContentKey.md) |  | 
**Content** | [**Content3**](Content3.md) |  | 
**ExpectedContent** | Pointer to [**Content5**](Content5.md) |  | [optional] 

## Methods

### NewPutV1

`func NewPutV1(key AddedContentKey, content Content3, ) *PutV1`

NewPutV1 instantiates a new PutV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutV1WithDefaults

`func NewPutV1WithDefaults() *PutV1`

NewPutV1WithDefaults instantiates a new PutV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PutV1) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PutV1) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PutV1) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.


### GetContent

`func (o *PutV1) GetContent() Content3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PutV1) GetContentOk() (*Content3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PutV1) SetContent(v Content3)`

SetContent sets Content field to given value.


### GetExpectedContent

`func (o *PutV1) GetExpectedContent() Content5`

GetExpectedContent returns the ExpectedContent field if non-nil, zero value otherwise.

### GetExpectedContentOk

`func (o *PutV1) GetExpectedContentOk() (*Content5, bool)`

GetExpectedContentOk returns a tuple with the ExpectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedContent

`func (o *PutV1) SetExpectedContent(v Content5)`

SetExpectedContent sets ExpectedContent field to given value.

### HasExpectedContent

`func (o *PutV1) HasExpectedContent() bool`

HasExpectedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


