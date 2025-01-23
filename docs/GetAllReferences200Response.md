# GetAllReferences200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**References** | [**[]Reference1**](Reference1.md) |  | 

## Methods

### NewGetAllReferences200Response

`func NewGetAllReferences200Response(references []Reference1, ) *GetAllReferences200Response`

NewGetAllReferences200Response instantiates a new GetAllReferences200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllReferences200ResponseWithDefaults

`func NewGetAllReferences200ResponseWithDefaults() *GetAllReferences200Response`

NewGetAllReferences200ResponseWithDefaults instantiates a new GetAllReferences200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllReferences200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllReferences200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllReferences200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetAllReferences200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *GetAllReferences200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetAllReferences200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetAllReferences200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetAllReferences200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetReferences

`func (o *GetAllReferences200Response) GetReferences() []Reference1`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetAllReferences200Response) GetReferencesOk() (*[]Reference1, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetAllReferences200Response) SetReferences(v []Reference1)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


