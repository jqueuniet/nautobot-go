# WritableCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Cid** | **string** |  | 
**Provider** | **string** |  | 
**Type** | **string** |  | 
**Status** | [**WritableCircuitStatusEnum**](WritableCircuitStatusEnum.md) |  | 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TerminationA** | [**CircuitTerminationA**](CircuitTerminationA.md) |  | 
**TerminationZ** | [**CircuitTerminationA**](CircuitTerminationA.md) |  | 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableCircuit

`func NewWritableCircuit(id string, url string, cid string, provider string, type_ string, status WritableCircuitStatusEnum, terminationA CircuitTerminationA, terminationZ CircuitTerminationA, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string, ) *WritableCircuit`

NewWritableCircuit instantiates a new WritableCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCircuitWithDefaults

`func NewWritableCircuitWithDefaults() *WritableCircuit`

NewWritableCircuitWithDefaults instantiates a new WritableCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableCircuit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableCircuit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableCircuit) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableCircuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableCircuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableCircuit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCid

`func (o *WritableCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *WritableCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *WritableCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *WritableCircuit) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WritableCircuit) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WritableCircuit) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *WritableCircuit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableCircuit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableCircuit) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *WritableCircuit) GetStatus() WritableCircuitStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableCircuit) GetStatusOk() (*WritableCircuitStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableCircuit) SetStatus(v WritableCircuitStatusEnum)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *WritableCircuit) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableCircuit) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableCircuit) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableCircuit) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableCircuit) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableCircuit) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *WritableCircuit) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *WritableCircuit) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *WritableCircuit) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *WritableCircuit) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *WritableCircuit) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *WritableCircuit) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetCommitRate

`func (o *WritableCircuit) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *WritableCircuit) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *WritableCircuit) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *WritableCircuit) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *WritableCircuit) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *WritableCircuit) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *WritableCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerminationA

`func (o *WritableCircuit) GetTerminationA() CircuitTerminationA`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *WritableCircuit) GetTerminationAOk() (*CircuitTerminationA, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *WritableCircuit) SetTerminationA(v CircuitTerminationA)`

SetTerminationA sets TerminationA field to given value.


### GetTerminationZ

`func (o *WritableCircuit) GetTerminationZ() CircuitTerminationA`

GetTerminationZ returns the TerminationZ field if non-nil, zero value otherwise.

### GetTerminationZOk

`func (o *WritableCircuit) GetTerminationZOk() (*CircuitTerminationA, bool)`

GetTerminationZOk returns a tuple with the TerminationZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationZ

`func (o *WritableCircuit) SetTerminationZ(v CircuitTerminationA)`

SetTerminationZ sets TerminationZ field to given value.


### GetComments

`func (o *WritableCircuit) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableCircuit) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableCircuit) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableCircuit) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableCircuit) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCircuit) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCircuit) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCircuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableCircuit) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableCircuit) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableCircuit) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableCircuit) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WritableCircuit) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableCircuit) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableCircuit) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *WritableCircuit) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableCircuit) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableCircuit) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetComputedFields

`func (o *WritableCircuit) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *WritableCircuit) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *WritableCircuit) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.


### GetDisplay

`func (o *WritableCircuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableCircuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableCircuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


