# JobResultSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Interval** | [**IntervalEnum**](IntervalEnum.md) |  | 

## Methods

### NewJobResultSchedule

`func NewJobResultSchedule(interval IntervalEnum, ) *JobResultSchedule`

NewJobResultSchedule instantiates a new JobResultSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultScheduleWithDefaults

`func NewJobResultScheduleWithDefaults() *JobResultSchedule`

NewJobResultScheduleWithDefaults instantiates a new JobResultSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobResultSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResultSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResultSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JobResultSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *JobResultSchedule) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *JobResultSchedule) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *JobResultSchedule) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *JobResultSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetInterval

`func (o *JobResultSchedule) GetInterval() IntervalEnum`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *JobResultSchedule) GetIntervalOk() (*IntervalEnum, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *JobResultSchedule) SetInterval(v IntervalEnum)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


