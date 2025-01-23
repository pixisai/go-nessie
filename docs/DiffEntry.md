# DiffEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**AddedContentKey**](AddedContentKey.md) |  | [optional] 
**From** | Pointer to [**Content2**](Content2.md) |  | [optional] 
**To** | Pointer to [**Content2**](Content2.md) |  | [optional] 

## Methods

### NewDiffEntry

`func NewDiffEntry() *DiffEntry`

NewDiffEntry instantiates a new DiffEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffEntryWithDefaults

`func NewDiffEntryWithDefaults() *DiffEntry`

NewDiffEntryWithDefaults instantiates a new DiffEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DiffEntry) GetKey() AddedContentKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DiffEntry) GetKeyOk() (*AddedContentKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DiffEntry) SetKey(v AddedContentKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *DiffEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFrom

`func (o *DiffEntry) GetFrom() Content2`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DiffEntry) GetFromOk() (*Content2, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DiffEntry) SetFrom(v Content2)`

SetFrom sets From field to given value.

### HasFrom

`func (o *DiffEntry) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *DiffEntry) GetTo() Content2`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DiffEntry) GetToOk() (*Content2, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DiffEntry) SetTo(v Content2)`

SetTo sets To field to given value.

### HasTo

`func (o *DiffEntry) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


