# Operation2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AddedContentKey**](AddedContentKey.md) |  | 
**Content** | [**Content3**](Content3.md) |  | 
**ExpectedContent** | Pointer to [**Content5**](Content5.md) |  | [optional] 

## Methods

### NewOperation2

`func NewOperation2(key AddedContentKey, content Content3, ) *Operation2`

NewOperation2 instantiates a new Operation2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperation2WithDefaults

`func NewOperation2WithDefaults() *Operation2`

NewOperation2WithDefaults instantiates a new Operation2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Operation2) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Operation2) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Operation2) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.


### GetContent

`func (o *Operation2) GetContent() Content3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Operation2) GetContentOk() (*Content3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Operation2) SetContent(v Content3)`

SetContent sets Content field to given value.


### GetExpectedContent

`func (o *Operation2) GetExpectedContent() Content5`

GetExpectedContent returns the ExpectedContent field if non-nil, zero value otherwise.

### GetExpectedContentOk

`func (o *Operation2) GetExpectedContentOk() (*Content5, bool)`

GetExpectedContentOk returns a tuple with the ExpectedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedContent

`func (o *Operation2) SetExpectedContent(v Content5)`

SetExpectedContent sets ExpectedContent field to given value.

### HasExpectedContent

`func (o *Operation2) HasExpectedContent() bool`

HasExpectedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


