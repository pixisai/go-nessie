# GetAllReferencesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**References** | [**[]Reference3**](Reference3.md) |  | 

## Methods

### NewGetAllReferencesV2200Response

`func NewGetAllReferencesV2200Response(references []Reference3, ) *GetAllReferencesV2200Response`

NewGetAllReferencesV2200Response instantiates a new GetAllReferencesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllReferencesV2200ResponseWithDefaults

`func NewGetAllReferencesV2200ResponseWithDefaults() *GetAllReferencesV2200Response`

NewGetAllReferencesV2200ResponseWithDefaults instantiates a new GetAllReferencesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllReferencesV2200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllReferencesV2200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllReferencesV2200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetAllReferencesV2200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *GetAllReferencesV2200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetAllReferencesV2200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetAllReferencesV2200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetAllReferencesV2200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetReferences

`func (o *GetAllReferencesV2200Response) GetReferences() []Reference3`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetAllReferencesV2200Response) GetReferencesOk() (*[]Reference3, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetAllReferencesV2200Response) SetReferences(v []Reference3)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


