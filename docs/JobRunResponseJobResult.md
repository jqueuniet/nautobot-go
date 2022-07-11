# JobRunResponseJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Created** | **time.Time** |  | [readonly] 
**Completed** | Pointer to **NullableTime** |  | [optional] 
**User** | [**JobResultUser**](JobResultUser.md) |  | 
**Status** | [**NestedJobResultStatus**](NestedJobResultStatus.md) |  | 

## Methods

### NewJobRunResponseJobResult

`func NewJobRunResponseJobResult(id string, url string, name string, created time.Time, user JobResultUser, status NestedJobResultStatus, ) *JobRunResponseJobResult`

NewJobRunResponseJobResult instantiates a new JobRunResponseJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRunResponseJobResultWithDefaults

`func NewJobRunResponseJobResultWithDefaults() *JobRunResponseJobResult`

NewJobRunResponseJobResultWithDefaults instantiates a new JobRunResponseJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobRunResponseJobResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobRunResponseJobResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobRunResponseJobResult) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *JobRunResponseJobResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobRunResponseJobResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobRunResponseJobResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *JobRunResponseJobResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobRunResponseJobResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobRunResponseJobResult) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *JobRunResponseJobResult) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JobRunResponseJobResult) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JobRunResponseJobResult) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCompleted

`func (o *JobRunResponseJobResult) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *JobRunResponseJobResult) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *JobRunResponseJobResult) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *JobRunResponseJobResult) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *JobRunResponseJobResult) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *JobRunResponseJobResult) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetUser

`func (o *JobRunResponseJobResult) GetUser() JobResultUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JobRunResponseJobResult) GetUserOk() (*JobResultUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JobRunResponseJobResult) SetUser(v JobResultUser)`

SetUser sets User field to given value.


### GetStatus

`func (o *JobRunResponseJobResult) GetStatus() NestedJobResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobRunResponseJobResult) GetStatusOk() (*NestedJobResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobRunResponseJobResult) SetStatus(v NestedJobResultStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


