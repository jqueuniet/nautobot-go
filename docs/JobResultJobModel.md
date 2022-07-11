# JobResultJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Source** | **string** | Source of the Python code for this job - local, Git repository, or plugins | [readonly] 
**ModuleName** | **string** | Dotted name of the Python module providing this job | [readonly] 
**JobClassName** | **string** | Name of the Python class providing this job | [readonly] 
**Grouping** | **string** | Human-readable grouping that this job belongs to | 
**Name** | **string** | Human-readable name of this job | 
**Slug** | Pointer to **string** |  | [optional] 

## Methods

### NewJobResultJobModel

`func NewJobResultJobModel(id string, url string, source string, moduleName string, jobClassName string, grouping string, name string, ) *JobResultJobModel`

NewJobResultJobModel instantiates a new JobResultJobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobResultJobModelWithDefaults

`func NewJobResultJobModelWithDefaults() *JobResultJobModel`

NewJobResultJobModelWithDefaults instantiates a new JobResultJobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobResultJobModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobResultJobModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobResultJobModel) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *JobResultJobModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JobResultJobModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JobResultJobModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *JobResultJobModel) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JobResultJobModel) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JobResultJobModel) SetSource(v string)`

SetSource sets Source field to given value.


### GetModuleName

`func (o *JobResultJobModel) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *JobResultJobModel) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *JobResultJobModel) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.


### GetJobClassName

`func (o *JobResultJobModel) GetJobClassName() string`

GetJobClassName returns the JobClassName field if non-nil, zero value otherwise.

### GetJobClassNameOk

`func (o *JobResultJobModel) GetJobClassNameOk() (*string, bool)`

GetJobClassNameOk returns a tuple with the JobClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobClassName

`func (o *JobResultJobModel) SetJobClassName(v string)`

SetJobClassName sets JobClassName field to given value.


### GetGrouping

`func (o *JobResultJobModel) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *JobResultJobModel) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *JobResultJobModel) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.


### GetName

`func (o *JobResultJobModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobResultJobModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobResultJobModel) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *JobResultJobModel) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *JobResultJobModel) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *JobResultJobModel) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *JobResultJobModel) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


