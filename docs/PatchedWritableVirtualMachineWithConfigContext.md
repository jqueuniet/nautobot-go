# PatchedWritableVirtualMachineWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**WritableVirtualMachineWithConfigContextStatusEnum**](WritableVirtualMachineWithConfigContextStatusEnum.md) |  | [optional] 
**Site** | Pointer to [**PatchedWritableVirtualMachineWithConfigContextSite**](PatchedWritableVirtualMachineWithConfigContextSite.md) |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**PrimaryIp** | Pointer to [**DevicePrimaryIp**](DevicePrimaryIp.md) |  | [optional] 
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
**ConfigContext** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableVirtualMachineWithConfigContext

`func NewPatchedWritableVirtualMachineWithConfigContext() *PatchedWritableVirtualMachineWithConfigContext`

NewPatchedWritableVirtualMachineWithConfigContext instantiates a new PatchedWritableVirtualMachineWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVirtualMachineWithConfigContextWithDefaults

`func NewPatchedWritableVirtualMachineWithConfigContextWithDefaults() *PatchedWritableVirtualMachineWithConfigContext`

NewPatchedWritableVirtualMachineWithConfigContextWithDefaults instantiates a new PatchedWritableVirtualMachineWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetStatus() WritableVirtualMachineWithConfigContextStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetStatusOk() (*WritableVirtualMachineWithConfigContextStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetStatus(v WritableVirtualMachineWithConfigContextStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetSite() PatchedWritableVirtualMachineWithConfigContextSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetSiteOk() (*PatchedWritableVirtualMachineWithConfigContextSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetSite(v PatchedWritableVirtualMachineWithConfigContextSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetCluster

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIp() DevicePrimaryIp`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIpOk() (*DevicePrimaryIp, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPrimaryIp(v DevicePrimaryIp)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIp4() string`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIp4Ok() (*string, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPrimaryIp4(v string)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIp6() string`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetPrimaryIp6Ok() (*string, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPrimaryIp6(v string)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetComments

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetLocalContextSchema

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLocalContextSchema() string`

GetLocalContextSchema returns the LocalContextSchema field if non-nil, zero value otherwise.

### GetLocalContextSchemaOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLocalContextSchemaOk() (*string, bool)`

GetLocalContextSchemaOk returns a tuple with the LocalContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextSchema

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetLocalContextSchema(v string)`

SetLocalContextSchema sets LocalContextSchema field to given value.

### HasLocalContextSchema

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasLocalContextSchema() bool`

HasLocalContextSchema returns a boolean if a field has been set.

### SetLocalContextSchemaNil

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetLocalContextSchemaNil(b bool)`

 SetLocalContextSchemaNil sets the value for LocalContextSchema to be an explicit nil

### UnsetLocalContextSchema
`func (o *PatchedWritableVirtualMachineWithConfigContext) UnsetLocalContextSchema()`

UnsetLocalContextSchema ensures that no value is present for LocalContextSchema, not even an explicit nil
### GetTags

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetConfigContext() map[string]interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetConfigContextOk() (*map[string]interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetConfigContext(v map[string]interface{})`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableVirtualMachineWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableVirtualMachineWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableVirtualMachineWithConfigContext) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


