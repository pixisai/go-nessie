# LogResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**LogEntries** | [**[]LogEntry1**](LogEntry1.md) |  | 

## Methods

### NewLogResponse1

`func NewLogResponse1(logEntries []LogEntry1, ) *LogResponse1`

NewLogResponse1 instantiates a new LogResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogResponse1WithDefaults

`func NewLogResponse1WithDefaults() *LogResponse1`

NewLogResponse1WithDefaults instantiates a new LogResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *LogResponse1) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LogResponse1) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LogResponse1) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LogResponse1) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetToken

`func (o *LogResponse1) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogResponse1) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogResponse1) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogResponse1) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLogEntries

`func (o *LogResponse1) GetLogEntries() []LogEntry1`

GetLogEntries returns the LogEntries field if non-nil, zero value otherwise.

### GetLogEntriesOk

`func (o *LogResponse1) GetLogEntriesOk() (*[]LogEntry1, bool)`

GetLogEntriesOk returns a tuple with the LogEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEntries

`func (o *LogResponse1) SetLogEntries(v []LogEntry1)`

SetLogEntries sets LogEntries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


