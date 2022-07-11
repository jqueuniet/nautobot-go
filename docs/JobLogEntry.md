# JobLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**AbsoluteUrl** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**JobResult** | **string** |  | 
**LogLevel** | Pointer to [**LogLevelEnum**](LogLevelEnum.md) |  | [optional] 
**LogObject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewJobLogEntry

`func NewJobLogEntry(id string, url string, jobResult string, ) *JobLogEntry`

NewJobLogEntry instantiates a new JobLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobLogEntryWithDefaults

`func NewJobLogEntryWithDefaults() *JobLogEntry`

NewJobLogEntryWithDefaults instantiates a new JobLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *JobLogEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobLogEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobLogEntry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAbsoluteUrl

`func (o *JobLogEntry) GetAbsoluteUrl() string`

GetAbsoluteUrl returns the AbsoluteUrl field if non-nil, zero value otherwise.

### GetAbsoluteUrlOk

`func (o *JobLogEntry) GetAbsoluteUrlOk() (*string, bool)`

GetAbsoluteUrlOk returns a tuple with the AbsoluteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteUrl

`func (o *JobLogEntry) SetAbsoluteUrl(v string)`

SetAbsoluteUrl sets AbsoluteUrl field to given value.

### HasAbsoluteUrl

`func (o *JobLogEntry) HasAbsoluteUrl() bool`

HasAbsoluteUrl returns a boolean if a field has been set.

### SetAbsoluteUrlNil

`func (o *JobLogEntry) SetAbsoluteUrlNil(b bool)`

 SetAbsoluteUrlNil sets the value for AbsoluteUrl to be an explicit nil

### UnsetAbsoluteUrl
`func (o *JobLogEntry) UnsetAbsoluteUrl()`

UnsetAbsoluteUrl ensures that no value is present for AbsoluteUrl, not even an explicit nil
### GetCreated

`func (o *JobLogEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JobLogEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JobLogEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JobLogEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGrouping

`func (o *JobLogEntry) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *JobLogEntry) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *JobLogEntry) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *JobLogEntry) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetJobResult

`func (o *JobLogEntry) GetJobResult() string`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *JobLogEntry) GetJobResultOk() (*string, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *JobLogEntry) SetJobResult(v string)`

SetJobResult sets JobResult field to given value.


### GetLogLevel

`func (o *JobLogEntry) GetLogLevel() LogLevelEnum`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *JobLogEntry) GetLogLevelOk() (*LogLevelEnum, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *JobLogEntry) SetLogLevel(v LogLevelEnum)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *JobLogEntry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogObject

`func (o *JobLogEntry) GetLogObject() string`

GetLogObject returns the LogObject field if non-nil, zero value otherwise.

### GetLogObjectOk

`func (o *JobLogEntry) GetLogObjectOk() (*string, bool)`

GetLogObjectOk returns a tuple with the LogObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogObject

`func (o *JobLogEntry) SetLogObject(v string)`

SetLogObject sets LogObject field to given value.

### HasLogObject

`func (o *JobLogEntry) HasLogObject() bool`

HasLogObject returns a boolean if a field has been set.

### SetLogObjectNil

`func (o *JobLogEntry) SetLogObjectNil(b bool)`

 SetLogObjectNil sets the value for LogObject to be an explicit nil

### UnsetLogObject
`func (o *JobLogEntry) UnsetLogObject()`

UnsetLogObject ensures that no value is present for LogObject, not even an explicit nil
### GetMessage

`func (o *JobLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JobLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JobLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JobLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


