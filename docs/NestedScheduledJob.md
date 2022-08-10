# NestedScheduledJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Interval** | [**IntervalEnum**](IntervalEnum.md) |  | 
**Crontab** | Pointer to **string** | Cronjob syntax string for custom scheduling | [optional] 

## Methods

### NewNestedScheduledJob

`func NewNestedScheduledJob(interval IntervalEnum, ) *NestedScheduledJob`

NewNestedScheduledJob instantiates a new NestedScheduledJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedScheduledJobWithDefaults

`func NewNestedScheduledJobWithDefaults() *NestedScheduledJob`

NewNestedScheduledJobWithDefaults instantiates a new NestedScheduledJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedScheduledJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedScheduledJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedScheduledJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NestedScheduledJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *NestedScheduledJob) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NestedScheduledJob) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NestedScheduledJob) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NestedScheduledJob) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetInterval

`func (o *NestedScheduledJob) GetInterval() IntervalEnum`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *NestedScheduledJob) GetIntervalOk() (*IntervalEnum, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *NestedScheduledJob) SetInterval(v IntervalEnum)`

SetInterval sets Interval field to given value.


### GetCrontab

`func (o *NestedScheduledJob) GetCrontab() string`

GetCrontab returns the Crontab field if non-nil, zero value otherwise.

### GetCrontabOk

`func (o *NestedScheduledJob) GetCrontabOk() (*string, bool)`

GetCrontabOk returns a tuple with the Crontab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrontab

`func (o *NestedScheduledJob) SetCrontab(v string)`

SetCrontab sets Crontab field to given value.

### HasCrontab

`func (o *NestedScheduledJob) HasCrontab() bool`

HasCrontab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


