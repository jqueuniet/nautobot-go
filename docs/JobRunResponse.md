# JobRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**JobResultSchedule**](JobResultSchedule.md) |  | 
**JobResult** | [**JobRunResponseJobResult**](JobRunResponseJobResult.md) |  | 

## Methods

### NewJobRunResponse

`func NewJobRunResponse(schedule JobResultSchedule, jobResult JobRunResponseJobResult, ) *JobRunResponse`

NewJobRunResponse instantiates a new JobRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRunResponseWithDefaults

`func NewJobRunResponseWithDefaults() *JobRunResponse`

NewJobRunResponseWithDefaults instantiates a new JobRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *JobRunResponse) GetSchedule() JobResultSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobRunResponse) GetScheduleOk() (*JobResultSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobRunResponse) SetSchedule(v JobResultSchedule)`

SetSchedule sets Schedule field to given value.


### GetJobResult

`func (o *JobRunResponse) GetJobResult() JobRunResponseJobResult`

GetJobResult returns the JobResult field if non-nil, zero value otherwise.

### GetJobResultOk

`func (o *JobRunResponse) GetJobResultOk() (*JobRunResponseJobResult, bool)`

GetJobResultOk returns a tuple with the JobResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobResult

`func (o *JobRunResponse) SetJobResult(v JobRunResponseJobResult)`

SetJobResult sets JobResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


