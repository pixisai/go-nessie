# LogResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**LogEntries** | [**[]LogEntry2**](LogEntry2.md) |  | 

## Methods

### NewLogResponse2

`func NewLogResponse2(logEntries []LogEntry2, ) *LogResponse2`

NewLogResponse2 instantiates a new LogResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResponse2WithDefaults

`func NewLogResponse2WithDefaults() *LogResponse2`

NewLogResponse2WithDefaults instantiates a new LogResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *LogResponse2) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LogResponse2) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LogResponse2) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LogResponse2) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *LogResponse2) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogResponse2) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogResponse2) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogResponse2) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLogEntries

`func (o *LogResponse2) GetLogEntries() []LogEntry2`

GetLogEntries returns the LogEntries field if non-nil, zero value otherwise.

### GetLogEntriesOk

`func (o *LogResponse2) GetLogEntriesOk() (*[]LogEntry2, bool)`

GetLogEntriesOk returns a tuple with the LogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEntries

`func (o *LogResponse2) SetLogEntries(v []LogEntry2)`

SetLogEntries sets LogEntries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


