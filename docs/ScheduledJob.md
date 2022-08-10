# ScheduledJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** | Short Description For This Task | 
**User** | [**JobResultUser**](JobResultUser.md) |  | 
**JobModel** | [**JobResultJobModel**](JobResultJobModel.md) |  | 
**Task** | **string** | The name of the Celery task that should be run. (Example: \&quot;proj.tasks.import_contacts\&quot;) | 
**Interval** | [**IntervalEnum**](IntervalEnum.md) |  | 
**Queue** | Pointer to **NullableString** | Queue defined in CELERY_TASK_QUEUES. Leave None for default queuing. | [optional] 
**JobClass** | **string** | Name of the fully qualified Nautobot Job class path | 
**LastRunAt** | **time.Time** | Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False. | [readonly] 
**TotalRunCount** | **int32** | Running count of how many times the schedule has triggered the task | [readonly] 
**DateChanged** | **time.Time** | Datetime that this scheduled job was last modified | [readonly] 
**Description** | Pointer to **string** | Detailed description about the details of this scheduled job | [optional] 
**ApprovedByUser** | [**JobResultUser**](JobResultUser.md) |  | 
**ApprovalRequired** | Pointer to **bool** |  | [optional] 
**ApprovedAt** | **time.Time** | Datetime that the schedule was approved | [readonly] 
**Crontab** | Pointer to **string** | Cronjob syntax string for custom scheduling | [optional] 

## Methods

### NewScheduledJob

`func NewScheduledJob(id string, url string, name string, user JobResultUser, jobModel JobResultJobModel, task string, interval IntervalEnum, jobClass string, lastRunAt time.Time, totalRunCount int32, dateChanged time.Time, approvedByUser JobResultUser, approvedAt time.Time, ) *ScheduledJob`

NewScheduledJob instantiates a new ScheduledJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledJobWithDefaults

`func NewScheduledJobWithDefaults() *ScheduledJob`

NewScheduledJobWithDefaults instantiates a new ScheduledJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduledJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledJob) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ScheduledJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScheduledJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScheduledJob) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *ScheduledJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledJob) SetName(v string)`

SetName sets Name field to given value.


### GetUser

`func (o *ScheduledJob) GetUser() JobResultUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ScheduledJob) GetUserOk() (*JobResultUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ScheduledJob) SetUser(v JobResultUser)`

SetUser sets User field to given value.


### GetJobModel

`func (o *ScheduledJob) GetJobModel() JobResultJobModel`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *ScheduledJob) GetJobModelOk() (*JobResultJobModel, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *ScheduledJob) SetJobModel(v JobResultJobModel)`

SetJobModel sets JobModel field to given value.


### GetTask

`func (o *ScheduledJob) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ScheduledJob) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ScheduledJob) SetTask(v string)`

SetTask sets Task field to given value.


### GetInterval

`func (o *ScheduledJob) GetInterval() IntervalEnum`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduledJob) GetIntervalOk() (*IntervalEnum, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduledJob) SetInterval(v IntervalEnum)`

SetInterval sets Interval field to given value.


### GetQueue

`func (o *ScheduledJob) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ScheduledJob) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ScheduledJob) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *ScheduledJob) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *ScheduledJob) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *ScheduledJob) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetJobClass

`func (o *ScheduledJob) GetJobClass() string`

GetJobClass returns the JobClass field if non-nil, zero value otherwise.

### GetJobClassOk

`func (o *ScheduledJob) GetJobClassOk() (*string, bool)`

GetJobClassOk returns a tuple with the JobClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobClass

`func (o *ScheduledJob) SetJobClass(v string)`

SetJobClass sets JobClass field to given value.


### GetLastRunAt

`func (o *ScheduledJob) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *ScheduledJob) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *ScheduledJob) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.


### GetTotalRunCount

`func (o *ScheduledJob) GetTotalRunCount() int32`

GetTotalRunCount returns the TotalRunCount field if non-nil, zero value otherwise.

### GetTotalRunCountOk

`func (o *ScheduledJob) GetTotalRunCountOk() (*int32, bool)`

GetTotalRunCountOk returns a tuple with the TotalRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRunCount

`func (o *ScheduledJob) SetTotalRunCount(v int32)`

SetTotalRunCount sets TotalRunCount field to given value.


### GetDateChanged

`func (o *ScheduledJob) GetDateChanged() time.Time`

GetDateChanged returns the DateChanged field if non-nil, zero value otherwise.

### GetDateChangedOk

`func (o *ScheduledJob) GetDateChangedOk() (*time.Time, bool)`

GetDateChangedOk returns a tuple with the DateChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateChanged

`func (o *ScheduledJob) SetDateChanged(v time.Time)`

SetDateChanged sets DateChanged field to given value.


### GetDescription

`func (o *ScheduledJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduledJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduledJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduledJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApprovedByUser

`func (o *ScheduledJob) GetApprovedByUser() JobResultUser`

GetApprovedByUser returns the ApprovedByUser field if non-nil, zero value otherwise.

### GetApprovedByUserOk

`func (o *ScheduledJob) GetApprovedByUserOk() (*JobResultUser, bool)`

GetApprovedByUserOk returns a tuple with the ApprovedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUser

`func (o *ScheduledJob) SetApprovedByUser(v JobResultUser)`

SetApprovedByUser sets ApprovedByUser field to given value.


### GetApprovalRequired

`func (o *ScheduledJob) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *ScheduledJob) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *ScheduledJob) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *ScheduledJob) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetApprovedAt

`func (o *ScheduledJob) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *ScheduledJob) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *ScheduledJob) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.


### GetCrontab

`func (o *ScheduledJob) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *ScheduledJob) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *ScheduledJob) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *ScheduledJob) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


