# WritableVirtualMachineWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Status** | [**WritableVirtualMachineWithConfigContextStatusEnum**](WritableVirtualMachineWithConfigContextStatusEnum.md) |  | 
**Site** | [**PatchedWritableVirtualMachineWithConfigContextSite**](PatchedWritableVirtualMachineWithConfigContextSite.md) |  | 
**Cluster** | **string** |  | 
**Role** | Pointer to **NullableString** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**PrimaryIp** | [**DevicePrimaryIp**](DevicePrimaryIp.md) |  | 
**PrimaryIp4** | Pointer to **NullableString** |  | [optional] 
**PrimaryIp6** | Pointer to **NullableString** |  | [optional] 
**Vcpus** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**LocalContextSchema** | Pointer to **NullableString** | Optional schema to validate the structure of the data | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigContext** | **map[string]interface{}** |  | [readonly] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewWritableVirtualMachineWithConfigContext

`func NewWritableVirtualMachineWithConfigContext(id string, url string, name string, status WritableVirtualMachineWithConfigContextStatusEnum, site PatchedWritableVirtualMachineWithConfigContextSite, cluster string, primaryIp DevicePrimaryIp, configContext map[string]interface{}, created string, lastUpdated time.Time, display string, ) *WritableVirtualMachineWithConfigContext`

NewWritableVirtualMachineWithConfigContext instantiates a new WritableVirtualMachineWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVirtualMachineWithConfigContextWithDefaults

`func NewWritableVirtualMachineWithConfigContextWithDefaults() *WritableVirtualMachineWithConfigContext`

NewWritableVirtualMachineWithConfigContextWithDefaults instantiates a new WritableVirtualMachineWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableVirtualMachineWithConfigContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableVirtualMachineWithConfigContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableVirtualMachineWithConfigContext) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WritableVirtualMachineWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WritableVirtualMachineWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WritableVirtualMachineWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *WritableVirtualMachineWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVirtualMachineWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVirtualMachineWithConfigContext) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *WritableVirtualMachineWithConfigContext) GetStatus() WritableVirtualMachineWithConfigContextStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableVirtualMachineWithConfigContext) GetStatusOk() (*WritableVirtualMachineWithConfigContextStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableVirtualMachineWithConfigContext) SetStatus(v WritableVirtualMachineWithConfigContextStatusEnum)`

SetStatus sets Status field to given value.


### GetSite

`func (o *WritableVirtualMachineWithConfigContext) GetSite() PatchedWritableVirtualMachineWithConfigContextSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WritableVirtualMachineWithConfigContext) GetSiteOk() (*PatchedWritableVirtualMachineWithConfigContextSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WritableVirtualMachineWithConfigContext) SetSite(v PatchedWritableVirtualMachineWithConfigContextSite)`

SetSite sets Site field to given value.


### GetCluster

`func (o *WritableVirtualMachineWithConfigContext) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WritableVirtualMachineWithConfigContext) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WritableVirtualMachineWithConfigContext) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetRole

`func (o *WritableVirtualMachineWithConfigContext) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableVirtualMachineWithConfigContext) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableVirtualMachineWithConfigContext) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableVirtualMachineWithConfigContext) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritableVirtualMachineWithConfigContext) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritableVirtualMachineWithConfigContext) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *WritableVirtualMachineWithConfigContext) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableVirtualMachineWithConfigContext) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableVirtualMachineWithConfigContext) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableVirtualMachineWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableVirtualMachineWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableVirtualMachineWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *WritableVirtualMachineWithConfigContext) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WritableVirtualMachineWithConfigContext) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WritableVirtualMachineWithConfigContext) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *WritableVirtualMachineWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *WritableVirtualMachineWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *WritableVirtualMachineWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp() DevicePrimaryIp`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIpOk() (*DevicePrimaryIp, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp(v DevicePrimaryIp)`

SetPrimaryIp sets PrimaryIp field to given value.


### GetPrimaryIp4

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp4() string`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp4Ok() (*string, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp4(v string)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *WritableVirtualMachineWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *WritableVirtualMachineWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp6() string`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *WritableVirtualMachineWithConfigContext) GetPrimaryIp6Ok() (*string, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp6(v string)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *WritableVirtualMachineWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *WritableVirtualMachineWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *WritableVirtualMachineWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *WritableVirtualMachineWithConfigContext) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *WritableVirtualMachineWithConfigContext) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *WritableVirtualMachineWithConfigContext) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *WritableVirtualMachineWithConfigContext) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *WritableVirtualMachineWithConfigContext) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *WritableVirtualMachineWithConfigContext) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *WritableVirtualMachineWithConfigContext) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *WritableVirtualMachineWithConfigContext) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *WritableVirtualMachineWithConfigContext) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *WritableVirtualMachineWithConfigContext) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *WritableVirtualMachineWithConfigContext) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *WritableVirtualMachineWithConfigContext) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *WritableVirtualMachineWithConfigContext) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *WritableVirtualMachineWithConfigContext) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *WritableVirtualMachineWithConfigContext) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *WritableVirtualMachineWithConfigContext) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *WritableVirtualMachineWithConfigContext) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *WritableVirtualMachineWithConfigContext) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetComments

`func (o *WritableVirtualMachineWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableVirtualMachineWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableVirtualMachineWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableVirtualMachineWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *WritableVirtualMachineWithConfigContext) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *WritableVirtualMachineWithConfigContext) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *WritableVirtualMachineWithConfigContext) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *WritableVirtualMachineWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *WritableVirtualMachineWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *WritableVirtualMachineWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetLocalContextSchema

`func (o *WritableVirtualMachineWithConfigContext) GetLocalContextSchema() string`

GetLocalContextSchema returns the LocalContextSchema field if non-nil, zero value otherwise.

### GetLocalContextSchemaOk

`func (o *WritableVirtualMachineWithConfigContext) GetLocalContextSchemaOk() (*string, bool)`

GetLocalContextSchemaOk returns a tuple with the LocalContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextSchema

`func (o *WritableVirtualMachineWithConfigContext) SetLocalContextSchema(v string)`

SetLocalContextSchema sets LocalContextSchema field to given value.

### HasLocalContextSchema

`func (o *WritableVirtualMachineWithConfigContext) HasLocalContextSchema() bool`

HasLocalContextSchema returns a boolean if a field has been set.

### SetLocalContextSchemaNil

`func (o *WritableVirtualMachineWithConfigContext) SetLocalContextSchemaNil(b bool)`

 SetLocalContextSchemaNil sets the value for LocalContextSchema to be an explicit nil

### UnsetLocalContextSchema
`func (o *WritableVirtualMachineWithConfigContext) UnsetLocalContextSchema()`

UnsetLocalContextSchema ensures that no value is present for LocalContextSchema, not even an explicit nil
### GetTags

`func (o *WritableVirtualMachineWithConfigContext) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVirtualMachineWithConfigContext) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVirtualMachineWithConfigContext) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVirtualMachineWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVirtualMachineWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVirtualMachineWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVirtualMachineWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVirtualMachineWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *WritableVirtualMachineWithConfigContext) GetConfigContext() map[string]interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *WritableVirtualMachineWithConfigContext) GetConfigContextOk() (*map[string]interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *WritableVirtualMachineWithConfigContext) SetConfigContext(v map[string]interface{})`

SetConfigContext sets ConfigContext field to given value.


### GetCreated

`func (o *WritableVirtualMachineWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WritableVirtualMachineWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WritableVirtualMachineWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *WritableVirtualMachineWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WritableVirtualMachineWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WritableVirtualMachineWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplay

`func (o *WritableVirtualMachineWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WritableVirtualMachineWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WritableVirtualMachineWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


