# NestedJob

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

### NewNestedJob

`func NewNestedJob(id string, url string, source string, moduleName string, jobClassName string, grouping string, name string, ) *NestedJob`

NewNestedJob instantiates a new NestedJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedJobWithDefaults

`func NewNestedJobWithDefaults() *NestedJob`

NewNestedJobWithDefaults instantiates a new NestedJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedJob) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedJob) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *NestedJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NestedJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NestedJob) SetSource(v string)`

SetSource sets Source field to given value.


### GetModuleName

`func (o *NestedJob) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *NestedJob) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *NestedJob) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.


### GetJobClassName

`func (o *NestedJob) GetJobClassName() string`

GetJobClassName returns the JobClassName field if non-nil, zero value otherwise.

### GetJobClassNameOk

`func (o *NestedJob) GetJobClassNameOk() (*string, bool)`

GetJobClassNameOk returns a tuple with the JobClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobClassName

`func (o *NestedJob) SetJobClassName(v string)`

SetJobClassName sets JobClassName field to given value.


### GetGrouping

`func (o *NestedJob) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *NestedJob) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *NestedJob) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.


### GetName

`func (o *NestedJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedJob) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedJob) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedJob) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedJob) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedJob) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


