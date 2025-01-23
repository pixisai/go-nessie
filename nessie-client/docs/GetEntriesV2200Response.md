# GetEntriesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Entries** | [**[]GetEntriesV2200ResponseEntriesInner**](GetEntriesV2200ResponseEntriesInner.md) |  | 
**EffectiveReference** | Pointer to [**Reference3**](Reference3.md) |  | [optional] 

## Methods

### NewGetEntriesV2200Response

`func NewGetEntriesV2200Response(entries []GetEntriesV2200ResponseEntriesInner, ) *GetEntriesV2200Response`

NewGetEntriesV2200Response instantiates a new GetEntriesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntriesV2200ResponseWithDefaults

`func NewGetEntriesV2200ResponseWithDefaults() *GetEntriesV2200Response`

NewGetEntriesV2200ResponseWithDefaults instantiates a new GetEntriesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetEntriesV2200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetEntriesV2200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetEntriesV2200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetEntriesV2200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *GetEntriesV2200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetEntriesV2200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetEntriesV2200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetEntriesV2200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetEntries

`func (o *GetEntriesV2200Response) GetEntries() []GetEntriesV2200ResponseEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *GetEntriesV2200Response) GetEntriesOk() (*[]GetEntriesV2200ResponseEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *GetEntriesV2200Response) SetEntries(v []GetEntriesV2200ResponseEntriesInner)`

SetEntries sets Entries field to given value.


### GetEffectiveReference

`func (o *GetEntriesV2200Response) GetEffectiveReference() Reference3`

GetEffectiveReference returns the EffectiveReference field if non-nil, zero value otherwise.

### GetEffectiveReferenceOk

`func (o *GetEntriesV2200Response) GetEffectiveReferenceOk() (*Reference3, bool)`

GetEffectiveReferenceOk returns a tuple with the EffectiveReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveReference

`func (o *GetEntriesV2200Response) SetEffectiveReference(v Reference3)`

SetEffectiveReference sets EffectiveReference field to given value.

### HasEffectiveReference

`func (o *GetEntriesV2200Response) HasEffectiveReference() bool`

HasEffectiveReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


