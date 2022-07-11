# NestedJobResult

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

### NewNestedJobResult

`func NewNestedJobResult(id string, url string, name string, created time.Time, user JobResultUser, status NestedJobResultStatus, ) *NestedJobResult`

NewNestedJobResult instantiates a new NestedJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedJobResultWithDefaults

`func NewNestedJobResultWithDefaults() *NestedJobResult`

NewNestedJobResultWithDefaults instantiates a new NestedJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedJobResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedJobResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedJobResult) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedJobResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedJobResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedJobResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedJobResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedJobResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedJobResult) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *NestedJobResult) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NestedJobResult) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NestedJobResult) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCompleted

`func (o *NestedJobResult) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *NestedJobResult) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *NestedJobResult) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *NestedJobResult) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *NestedJobResult) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *NestedJobResult) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetUser

`func (o *NestedJobResult) GetUser() JobResultUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NestedJobResult) GetUserOk() (*JobResultUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NestedJobResult) SetUser(v JobResultUser)`

SetUser sets User field to given value.


### GetStatus

`func (o *NestedJobResult) GetStatus() NestedJobResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NestedJobResult) GetStatusOk() (*NestedJobResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NestedJobResult) SetStatus(v NestedJobResultStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


