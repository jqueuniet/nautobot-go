# PatchedJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Source** | Pointer to **string** | Source of the Python code for this job - local, Git repository, or plugins | [optional] [readonly] 
**ModuleName** | Pointer to **string** | Dotted name of the Python module providing this job | [optional] [readonly] 
**JobClassName** | Pointer to **string** | Name of the Python class providing this job | [optional] [readonly] 
**Grouping** | Pointer to **string** | Human-readable grouping that this job belongs to | [optional] 
**GroupingOverride** | Pointer to **bool** | If set, the configured grouping will remain even if the underlying Job source code changes | [optional] 
**Name** | Pointer to **string** | Human-readable name of this job | [optional] 
**NameOverride** | Pointer to **bool** | If set, the configured name will remain even if the underlying Job source code changes | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Markdown formatting is supported | [optional] 
**DescriptionOverride** | Pointer to **bool** | If set, the configured description will remain even if the underlying Job source code changes | [optional] 
**Installed** | Pointer to **bool** | Whether the Python module and class providing this job are presently installed and loadable | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Whether this job can be executed by users | [optional] 
**HasSensitiveVariables** | Pointer to **bool** | Whether this job contains sensitive variables | [optional] 
**HasSensitiveVariablesOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**ApprovalRequired** | Pointer to **bool** | Whether the job requires approval from another user before running | [optional] 
**ApprovalRequiredOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**CommitDefault** | Pointer to **bool** | Whether the job defaults to committing changes when run, or defaults to a dry-run | [optional] 
**CommitDefaultOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**Hidden** | Pointer to **bool** | Whether the job defaults to not being shown in the UI | [optional] 
**HiddenOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**ReadOnly** | Pointer to **bool** | Whether the job is prevented from making lasting changes to the database | [optional] 
**ReadOnlyOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**SoftTimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will receive a &lt;code&gt;SoftTimeLimitExceeded&lt;/code&gt; exception.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**SoftTimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**TimeLimit** | Pointer to **float64** | Maximum runtime in seconds before the job will be forcibly terminated.&lt;br&gt;Set to 0 to use Nautobot system default | [optional] 
**TimeLimitOverride** | Pointer to **bool** | If set, the configured value will remain even if the underlying Job source code changes | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedJob

`func NewPatchedJob() *PatchedJob`

NewPatchedJob instantiates a new PatchedJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJobWithDefaults

`func NewPatchedJobWithDefaults() *PatchedJob`

NewPatchedJobWithDefaults instantiates a new PatchedJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedJob) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedJob) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSource

`func (o *PatchedJob) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchedJob) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchedJob) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchedJob) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetModuleName

`func (o *PatchedJob) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *PatchedJob) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *PatchedJob) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *PatchedJob) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetJobClassName

`func (o *PatchedJob) GetJobClassName() string`

GetJobClassName returns the JobClassName field if non-nil, zero value otherwise.

### GetJobClassNameOk

`func (o *PatchedJob) GetJobClassNameOk() (*string, bool)`

GetJobClassNameOk returns a tuple with the JobClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobClassName

`func (o *PatchedJob) SetJobClassName(v string)`

SetJobClassName sets JobClassName field to given value.

### HasJobClassName

`func (o *PatchedJob) HasJobClassName() bool`

HasJobClassName returns a boolean if a field has been set.

### GetGrouping

`func (o *PatchedJob) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *PatchedJob) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *PatchedJob) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *PatchedJob) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetGroupingOverride

`func (o *PatchedJob) GetGroupingOverride() bool`

GetGroupingOverride returns the GroupingOverride field if non-nil, zero value otherwise.

### GetGroupingOverrideOk

`func (o *PatchedJob) GetGroupingOverrideOk() (*bool, bool)`

GetGroupingOverrideOk returns a tuple with the GroupingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingOverride

`func (o *PatchedJob) SetGroupingOverride(v bool)`

SetGroupingOverride sets GroupingOverride field to given value.

### HasGroupingOverride

`func (o *PatchedJob) HasGroupingOverride() bool`

HasGroupingOverride returns a boolean if a field has been set.

### GetName

`func (o *PatchedJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameOverride

`func (o *PatchedJob) GetNameOverride() bool`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *PatchedJob) GetNameOverrideOk() (*bool, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *PatchedJob) SetNameOverride(v bool)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *PatchedJob) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedJob) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedJob) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedJob) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedJob) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionOverride

`func (o *PatchedJob) GetDescriptionOverride() bool`

GetDescriptionOverride returns the DescriptionOverride field if non-nil, zero value otherwise.

### GetDescriptionOverrideOk

`func (o *PatchedJob) GetDescriptionOverrideOk() (*bool, bool)`

GetDescriptionOverrideOk returns a tuple with the DescriptionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverride

`func (o *PatchedJob) SetDescriptionOverride(v bool)`

SetDescriptionOverride sets DescriptionOverride field to given value.

### HasDescriptionOverride

`func (o *PatchedJob) HasDescriptionOverride() bool`

HasDescriptionOverride returns a boolean if a field has been set.

### GetInstalled

`func (o *PatchedJob) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *PatchedJob) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *PatchedJob) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *PatchedJob) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasSensitiveVariables

`func (o *PatchedJob) GetHasSensitiveVariables() bool`

GetHasSensitiveVariables returns the HasSensitiveVariables field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOk

`func (o *PatchedJob) GetHasSensitiveVariablesOk() (*bool, bool)`

GetHasSensitiveVariablesOk returns a tuple with the HasSensitiveVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariables

`func (o *PatchedJob) SetHasSensitiveVariables(v bool)`

SetHasSensitiveVariables sets HasSensitiveVariables field to given value.

### HasHasSensitiveVariables

`func (o *PatchedJob) HasHasSensitiveVariables() bool`

HasHasSensitiveVariables returns a boolean if a field has been set.

### GetHasSensitiveVariablesOverride

`func (o *PatchedJob) GetHasSensitiveVariablesOverride() bool`

GetHasSensitiveVariablesOverride returns the HasSensitiveVariablesOverride field if non-nil, zero value otherwise.

### GetHasSensitiveVariablesOverrideOk

`func (o *PatchedJob) GetHasSensitiveVariablesOverrideOk() (*bool, bool)`

GetHasSensitiveVariablesOverrideOk returns a tuple with the HasSensitiveVariablesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSensitiveVariablesOverride

`func (o *PatchedJob) SetHasSensitiveVariablesOverride(v bool)`

SetHasSensitiveVariablesOverride sets HasSensitiveVariablesOverride field to given value.

### HasHasSensitiveVariablesOverride

`func (o *PatchedJob) HasHasSensitiveVariablesOverride() bool`

HasHasSensitiveVariablesOverride returns a boolean if a field has been set.

### GetApprovalRequired

`func (o *PatchedJob) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *PatchedJob) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *PatchedJob) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *PatchedJob) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetApprovalRequiredOverride

`func (o *PatchedJob) GetApprovalRequiredOverride() bool`

GetApprovalRequiredOverride returns the ApprovalRequiredOverride field if non-nil, zero value otherwise.

### GetApprovalRequiredOverrideOk

`func (o *PatchedJob) GetApprovalRequiredOverrideOk() (*bool, bool)`

GetApprovalRequiredOverrideOk returns a tuple with the ApprovalRequiredOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequiredOverride

`func (o *PatchedJob) SetApprovalRequiredOverride(v bool)`

SetApprovalRequiredOverride sets ApprovalRequiredOverride field to given value.

### HasApprovalRequiredOverride

`func (o *PatchedJob) HasApprovalRequiredOverride() bool`

HasApprovalRequiredOverride returns a boolean if a field has been set.

### GetCommitDefault

`func (o *PatchedJob) GetCommitDefault() bool`

GetCommitDefault returns the CommitDefault field if non-nil, zero value otherwise.

### GetCommitDefaultOk

`func (o *PatchedJob) GetCommitDefaultOk() (*bool, bool)`

GetCommitDefaultOk returns a tuple with the CommitDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDefault

`func (o *PatchedJob) SetCommitDefault(v bool)`

SetCommitDefault sets CommitDefault field to given value.

### HasCommitDefault

`func (o *PatchedJob) HasCommitDefault() bool`

HasCommitDefault returns a boolean if a field has been set.

### GetCommitDefaultOverride

`func (o *PatchedJob) GetCommitDefaultOverride() bool`

GetCommitDefaultOverride returns the CommitDefaultOverride field if non-nil, zero value otherwise.

### GetCommitDefaultOverrideOk

`func (o *PatchedJob) GetCommitDefaultOverrideOk() (*bool, bool)`

GetCommitDefaultOverrideOk returns a tuple with the CommitDefaultOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDefaultOverride

`func (o *PatchedJob) SetCommitDefaultOverride(v bool)`

SetCommitDefaultOverride sets CommitDefaultOverride field to given value.

### HasCommitDefaultOverride

`func (o *PatchedJob) HasCommitDefaultOverride() bool`

HasCommitDefaultOverride returns a boolean if a field has been set.

### GetHidden

`func (o *PatchedJob) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *PatchedJob) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *PatchedJob) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *PatchedJob) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHiddenOverride

`func (o *PatchedJob) GetHiddenOverride() bool`

GetHiddenOverride returns the HiddenOverride field if non-nil, zero value otherwise.

### GetHiddenOverrideOk

`func (o *PatchedJob) GetHiddenOverrideOk() (*bool, bool)`

GetHiddenOverrideOk returns a tuple with the HiddenOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenOverride

`func (o *PatchedJob) SetHiddenOverride(v bool)`

SetHiddenOverride sets HiddenOverride field to given value.

### HasHiddenOverride

`func (o *PatchedJob) HasHiddenOverride() bool`

HasHiddenOverride returns a boolean if a field has been set.

### GetReadOnly

`func (o *PatchedJob) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PatchedJob) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PatchedJob) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PatchedJob) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetReadOnlyOverride

`func (o *PatchedJob) GetReadOnlyOverride() bool`

GetReadOnlyOverride returns the ReadOnlyOverride field if non-nil, zero value otherwise.

### GetReadOnlyOverrideOk

`func (o *PatchedJob) GetReadOnlyOverrideOk() (*bool, bool)`

GetReadOnlyOverrideOk returns a tuple with the ReadOnlyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyOverride

`func (o *PatchedJob) SetReadOnlyOverride(v bool)`

SetReadOnlyOverride sets ReadOnlyOverride field to given value.

### HasReadOnlyOverride

`func (o *PatchedJob) HasReadOnlyOverride() bool`

HasReadOnlyOverride returns a boolean if a field has been set.

### GetSoftTimeLimit

`func (o *PatchedJob) GetSoftTimeLimit() float64`

GetSoftTimeLimit returns the SoftTimeLimit field if non-nil, zero value otherwise.

### GetSoftTimeLimitOk

`func (o *PatchedJob) GetSoftTimeLimitOk() (*float64, bool)`

GetSoftTimeLimitOk returns a tuple with the SoftTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimit

`func (o *PatchedJob) SetSoftTimeLimit(v float64)`

SetSoftTimeLimit sets SoftTimeLimit field to given value.

### HasSoftTimeLimit

`func (o *PatchedJob) HasSoftTimeLimit() bool`

HasSoftTimeLimit returns a boolean if a field has been set.

### GetSoftTimeLimitOverride

`func (o *PatchedJob) GetSoftTimeLimitOverride() bool`

GetSoftTimeLimitOverride returns the SoftTimeLimitOverride field if non-nil, zero value otherwise.

### GetSoftTimeLimitOverrideOk

`func (o *PatchedJob) GetSoftTimeLimitOverrideOk() (*bool, bool)`

GetSoftTimeLimitOverrideOk returns a tuple with the SoftTimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftTimeLimitOverride

`func (o *PatchedJob) SetSoftTimeLimitOverride(v bool)`

SetSoftTimeLimitOverride sets SoftTimeLimitOverride field to given value.

### HasSoftTimeLimitOverride

`func (o *PatchedJob) HasSoftTimeLimitOverride() bool`

HasSoftTimeLimitOverride returns a boolean if a field has been set.

### GetTimeLimit

`func (o *PatchedJob) GetTimeLimit() float64`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *PatchedJob) GetTimeLimitOk() (*float64, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *PatchedJob) SetTimeLimit(v float64)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *PatchedJob) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetTimeLimitOverride

`func (o *PatchedJob) GetTimeLimitOverride() bool`

GetTimeLimitOverride returns the TimeLimitOverride field if non-nil, zero value otherwise.

### GetTimeLimitOverrideOk

`func (o *PatchedJob) GetTimeLimitOverrideOk() (*bool, bool)`

GetTimeLimitOverrideOk returns a tuple with the TimeLimitOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitOverride

`func (o *PatchedJob) SetTimeLimitOverride(v bool)`

SetTimeLimitOverride sets TimeLimitOverride field to given value.

### HasTimeLimitOverride

`func (o *PatchedJob) HasTimeLimitOverride() bool`

HasTimeLimitOverride returns a boolean if a field has been set.

### GetTags

`func (o *PatchedJob) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedJob) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedJob) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedJob) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedJob) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedJob) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedJob) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedJob) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedJob) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedJob) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedJob) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedJob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedJob) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedJob) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedJob) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedJob) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedJob) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedJob) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedJob) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedJob) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


