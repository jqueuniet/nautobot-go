# JobInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Commit** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**NestedScheduledJob**](NestedScheduledJob.md) |  | [optional] 

## Methods

### NewJobInput

`func NewJobInput() *JobInput`

NewJobInput instantiates a new JobInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInputWithDefaults

`func NewJobInputWithDefaults() *JobInput`

NewJobInputWithDefaults instantiates a new JobInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JobInput) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobInput) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobInput) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *JobInput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCommit

`func (o *JobInput) GetCommit() bool`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *JobInput) GetCommitOk() (*bool, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *JobInput) SetCommit(v bool)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *JobInput) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSchedule

`func (o *JobInput) GetSchedule() NestedScheduledJob`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobInput) GetScheduleOk() (*NestedScheduledJob, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobInput) SetSchedule(v NestedScheduledJob)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *JobInput) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


