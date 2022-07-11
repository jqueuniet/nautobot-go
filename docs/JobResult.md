# JobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Completed** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**JobModel** | [**JobResultJobModel**](JobResultJobModel.md) |  | 
**ObjType** | **string** |  | [readonly] 
**Status** | [**JobResultStatus**](JobResultStatus.md) |  | 
**User** | [**JobResultUser**](JobResultUser.md) |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**JobId** | **string** |  | 
**Schedule** | [**JobResultSchedule**](JobResultSchedule.md) |  | 

## Methods

### NewJobResult

`func NewJobResult(id string, url string, created time.Time, name string, jobModel JobResultJobModel, objType string, status JobResultStatus, user JobResultUser, jobId string, schedule JobResultSchedule, ) *JobResult`

NewJobResult instantiates a new JobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultWithDefaults

`func NewJobResultWithDefaults() *JobResult`

NewJobResultWithDefaults instantiates a new JobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResult) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *JobResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreated

`func (o *JobResult) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JobResult) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JobResult) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCompleted

`func (o *JobResult) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JobResult) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JobResult) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JobResult) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *JobResult) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *JobResult) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetName

`func (o *JobResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResult) SetName(v string)`

SetName sets Name field to given value.


### GetJobModel

`func (o *JobResult) GetJobModel() JobResultJobModel`

GetJobModel returns the JobModel field if non-nil, zero value otherwise.

### GetJobModelOk

`func (o *JobResult) GetJobModelOk() (*JobResultJobModel, bool)`

GetJobModelOk returns a tuple with the JobModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModel

`func (o *JobResult) SetJobModel(v JobResultJobModel)`

SetJobModel sets JobModel field to given value.


### GetObjType

`func (o *JobResult) GetObjType() string`

GetObjType returns the ObjType field if non-nil, zero value otherwise.

### GetObjTypeOk

`func (o *JobResult) GetObjTypeOk() (*string, bool)`

GetObjTypeOk returns a tuple with the ObjType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjType

`func (o *JobResult) SetObjType(v string)`

SetObjType sets ObjType field to given value.


### GetStatus

`func (o *JobResult) GetStatus() JobResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobResult) GetStatusOk() (*JobResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobResult) SetStatus(v JobResultStatus)`

SetStatus sets Status field to given value.


### GetUser

`func (o *JobResult) GetUser() JobResultUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobResult) GetUserOk() (*JobResultUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobResult) SetUser(v JobResultUser)`

SetUser sets User field to given value.


### GetData

`func (o *JobResult) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobResult) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobResult) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *JobResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *JobResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *JobResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetJobId

`func (o *JobResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobResult) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetSchedule

`func (o *JobResult) GetSchedule() JobResultSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobResult) GetScheduleOk() (*JobResultSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobResult) SetSchedule(v JobResultSchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


