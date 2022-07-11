# JobClassDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | [readonly] 
**Id** | **string** |  | [readonly] 
**Pk** | **NullableString** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Description** | **string** |  | [readonly] 
**TestMethods** | **[]string** |  | 
**Vars** | **map[string]interface{}** |  | [readonly] 
**Result** | Pointer to [**JobResult**](JobResult.md) |  | [optional] 

## Methods

### NewJobClassDetail

`func NewJobClassDetail(url string, id string, pk NullableString, name string, description string, testMethods []string, vars map[string]interface{}, ) *JobClassDetail`

NewJobClassDetail instantiates a new JobClassDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobClassDetailWithDefaults

`func NewJobClassDetailWithDefaults() *JobClassDetail`

NewJobClassDetailWithDefaults instantiates a new JobClassDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *JobClassDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobClassDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobClassDetail) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *JobClassDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobClassDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobClassDetail) SetId(v string)`

SetId sets Id field to given value.


### GetPk

`func (o *JobClassDetail) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *JobClassDetail) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *JobClassDetail) SetPk(v string)`

SetPk sets Pk field to given value.


### SetPkNil

`func (o *JobClassDetail) SetPkNil(b bool)`

 SetPkNil sets the value for Pk to be an explicit nil

### UnsetPk
`func (o *JobClassDetail) UnsetPk()`

UnsetPk ensures that no value is present for Pk, not even an explicit nil
### GetName

`func (o *JobClassDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobClassDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobClassDetail) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobClassDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobClassDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobClassDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTestMethods

`func (o *JobClassDetail) GetTestMethods() []string`

GetTestMethods returns the TestMethods field if non-nil, zero value otherwise.

### GetTestMethodsOk

`func (o *JobClassDetail) GetTestMethodsOk() (*[]string, bool)`

GetTestMethodsOk returns a tuple with the TestMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMethods

`func (o *JobClassDetail) SetTestMethods(v []string)`

SetTestMethods sets TestMethods field to given value.


### GetVars

`func (o *JobClassDetail) GetVars() map[string]interface{}`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *JobClassDetail) GetVarsOk() (*map[string]interface{}, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *JobClassDetail) SetVars(v map[string]interface{})`

SetVars sets Vars field to given value.


### GetResult

`func (o *JobClassDetail) GetResult() JobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobClassDetail) GetResultOk() (*JobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobClassDetail) SetResult(v JobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *JobClassDetail) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


