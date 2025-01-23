# GetEntries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Entries** | [**[]GetEntries200ResponseEntriesInner**](GetEntries200ResponseEntriesInner.md) |  | 

## Methods

### NewGetEntries200Response

`func NewGetEntries200Response(entries []GetEntries200ResponseEntriesInner, ) *GetEntries200Response`

NewGetEntries200Response instantiates a new GetEntries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntries200ResponseWithDefaults

`func NewGetEntries200ResponseWithDefaults() *GetEntries200Response`

NewGetEntries200ResponseWithDefaults instantiates a new GetEntries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetEntries200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetEntries200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetEntries200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetEntries200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *GetEntries200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetEntries200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetEntries200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetEntries200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetEntries

`func (o *GetEntries200Response) GetEntries() []GetEntries200ResponseEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GetEntries200Response) GetEntriesOk() (*[]GetEntries200ResponseEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GetEntries200Response) SetEntries(v []GetEntries200ResponseEntriesInner)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


