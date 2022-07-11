# JobClass

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
**Result** | Pointer to [**NestedJobResult**](NestedJobResult.md) |  | [optional] 

## Methods

### NewJobClass

`func NewJobClass(url string, id string, pk NullableString, name string, description string, testMethods []string, vars map[string]interface{}, ) *JobClass`

NewJobClass instantiates a new JobClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobClassWithDefaults

`func NewJobClassWithDefaults() *JobClass`

NewJobClassWithDefaults instantiates a new JobClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *JobClass) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobClass) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobClass) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *JobClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobClass) SetId(v string)`

SetId sets Id field to given value.


### GetPk

`func (o *JobClass) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *JobClass) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *JobClass) SetPk(v string)`

SetPk sets Pk field to given value.


### SetPkNil

`func (o *JobClass) SetPkNil(b bool)`

 SetPkNil sets the value for Pk to be an explicit nil

### UnsetPk
`func (o *JobClass) UnsetPk()`

UnsetPk ensures that no value is present for Pk, not even an explicit nil
### GetName

`func (o *JobClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobClass) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobClass) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTestMethods

`func (o *JobClass) GetTestMethods() []string`

GetTestMethods returns the TestMethods field if non-nil, zero value otherwise.

### GetTestMethodsOk

`func (o *JobClass) GetTestMethodsOk() (*[]string, bool)`

GetTestMethodsOk returns a tuple with the TestMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMethods

`func (o *JobClass) SetTestMethods(v []string)`

SetTestMethods sets TestMethods field to given value.


### GetVars

`func (o *JobClass) GetVars() map[string]interface{}`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *JobClass) GetVarsOk() (*map[string]interface{}, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *JobClass) SetVars(v map[string]interface{})`

SetVars sets Vars field to given value.


### GetResult

`func (o *JobClass) GetResult() NestedJobResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *JobClass) GetResultOk() (*NestedJobResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *JobClass) SetResult(v NestedJobResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *JobClass) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


