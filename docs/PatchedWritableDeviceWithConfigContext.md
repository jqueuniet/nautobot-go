# PatchedWritableDeviceWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**DeviceRole** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Rack** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**Face** | Pointer to [**RackFace**](RackFace.md) |  | [optional] 
**ParentDevice** | Pointer to [**DeviceParentDevice**](DeviceParentDevice.md) |  | [optional] 
**Status** | Pointer to [**WritableDeviceWithConfigContextStatusEnum**](WritableDeviceWithConfigContextStatusEnum.md) |  | [optional] 
**PrimaryIp** | Pointer to [**DevicePrimaryIp**](DevicePrimaryIp.md) |  | [optional] 
**PrimaryIp4** | Pointer to **NullableString** |  | [optional] 
**PrimaryIp6** | Pointer to **NullableString** |  | [optional] 
**SecretsGroup** | Pointer to **NullableString** |  | [optional] 
**Cluster** | Pointer to **NullableString** |  | [optional] 
**VirtualChassis** | Pointer to **NullableString** |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextSchema** | Pointer to **NullableString** | Optional schema to validate the structure of the data | [optional] 
**LocalContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigContext** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableDeviceWithConfigContext

`func NewPatchedWritableDeviceWithConfigContext() *PatchedWritableDeviceWithConfigContext`

NewPatchedWritableDeviceWithConfigContext instantiates a new PatchedWritableDeviceWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeviceWithConfigContextWithDefaults

`func NewPatchedWritableDeviceWithConfigContextWithDefaults() *PatchedWritableDeviceWithConfigContext`

NewPatchedWritableDeviceWithConfigContextWithDefaults instantiates a new PatchedWritableDeviceWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableDeviceWithConfigContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableDeviceWithConfigContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableDeviceWithConfigContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableDeviceWithConfigContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableDeviceWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableDeviceWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableDeviceWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableDeviceWithConfigContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableDeviceWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeviceWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeviceWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeviceWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableDeviceWithConfigContext) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableDeviceWithConfigContext) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *PatchedWritableDeviceWithConfigContext) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableDeviceWithConfigContext) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableDeviceWithConfigContext) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableDeviceWithConfigContext) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceRole

`func (o *PatchedWritableDeviceWithConfigContext) GetDeviceRole() string`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *PatchedWritableDeviceWithConfigContext) GetDeviceRoleOk() (*string, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *PatchedWritableDeviceWithConfigContext) SetDeviceRole(v string)`

SetDeviceRole sets DeviceRole field to given value.

### HasDeviceRole

`func (o *PatchedWritableDeviceWithConfigContext) HasDeviceRole() bool`

HasDeviceRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableDeviceWithConfigContext) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableDeviceWithConfigContext) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableDeviceWithConfigContext) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableDeviceWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableDeviceWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableDeviceWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableDeviceWithConfigContext) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableDeviceWithConfigContext) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableDeviceWithConfigContext) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableDeviceWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableDeviceWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableDeviceWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *PatchedWritableDeviceWithConfigContext) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableDeviceWithConfigContext) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableDeviceWithConfigContext) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableDeviceWithConfigContext) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableDeviceWithConfigContext) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableDeviceWithConfigContext) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableDeviceWithConfigContext) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableDeviceWithConfigContext) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableDeviceWithConfigContext) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableDeviceWithConfigContext) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *PatchedWritableDeviceWithConfigContext) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableDeviceWithConfigContext) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableDeviceWithConfigContext) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableDeviceWithConfigContext) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetRack

`func (o *PatchedWritableDeviceWithConfigContext) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableDeviceWithConfigContext) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableDeviceWithConfigContext) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableDeviceWithConfigContext) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritableDeviceWithConfigContext) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritableDeviceWithConfigContext) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *PatchedWritableDeviceWithConfigContext) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedWritableDeviceWithConfigContext) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedWritableDeviceWithConfigContext) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedWritableDeviceWithConfigContext) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedWritableDeviceWithConfigContext) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedWritableDeviceWithConfigContext) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedWritableDeviceWithConfigContext) GetFace() RackFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedWritableDeviceWithConfigContext) GetFaceOk() (*RackFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedWritableDeviceWithConfigContext) SetFace(v RackFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedWritableDeviceWithConfigContext) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetParentDevice

`func (o *PatchedWritableDeviceWithConfigContext) GetParentDevice() DeviceParentDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *PatchedWritableDeviceWithConfigContext) GetParentDeviceOk() (*DeviceParentDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *PatchedWritableDeviceWithConfigContext) SetParentDevice(v DeviceParentDevice)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *PatchedWritableDeviceWithConfigContext) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableDeviceWithConfigContext) GetStatus() WritableDeviceWithConfigContextStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableDeviceWithConfigContext) GetStatusOk() (*WritableDeviceWithConfigContextStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableDeviceWithConfigContext) SetStatus(v WritableDeviceWithConfigContextStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableDeviceWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIp() DevicePrimaryIp`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIpOk() (*DevicePrimaryIp, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *PatchedWritableDeviceWithConfigContext) SetPrimaryIp(v DevicePrimaryIp)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *PatchedWritableDeviceWithConfigContext) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIp4() string`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIp4Ok() (*string, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContext) SetPrimaryIp4(v string)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableDeviceWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableDeviceWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIp6() string`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableDeviceWithConfigContext) GetPrimaryIp6Ok() (*string, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContext) SetPrimaryIp6(v string)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableDeviceWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableDeviceWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetSecretsGroup

`func (o *PatchedWritableDeviceWithConfigContext) GetSecretsGroup() string`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedWritableDeviceWithConfigContext) GetSecretsGroupOk() (*string, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedWritableDeviceWithConfigContext) SetSecretsGroup(v string)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedWritableDeviceWithConfigContext) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedWritableDeviceWithConfigContext) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedWritableDeviceWithConfigContext) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCluster

`func (o *PatchedWritableDeviceWithConfigContext) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableDeviceWithConfigContext) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableDeviceWithConfigContext) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableDeviceWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableDeviceWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableDeviceWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContext) GetVirtualChassis() string`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedWritableDeviceWithConfigContext) GetVirtualChassisOk() (*string, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContext) SetVirtualChassis(v string)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContext) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedWritableDeviceWithConfigContext) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedWritableDeviceWithConfigContext) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *PatchedWritableDeviceWithConfigContext) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedWritableDeviceWithConfigContext) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedWritableDeviceWithConfigContext) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedWritableDeviceWithConfigContext) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedWritableDeviceWithConfigContext) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedWritableDeviceWithConfigContext) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedWritableDeviceWithConfigContext) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedWritableDeviceWithConfigContext) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedWritableDeviceWithConfigContext) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedWritableDeviceWithConfigContext) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedWritableDeviceWithConfigContext) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedWritableDeviceWithConfigContext) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *PatchedWritableDeviceWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeviceWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeviceWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeviceWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextSchema

`func (o *PatchedWritableDeviceWithConfigContext) GetLocalContextSchema() string`

GetLocalContextSchema returns the LocalContextSchema field if non-nil, zero value otherwise.

### GetLocalContextSchemaOk

`func (o *PatchedWritableDeviceWithConfigContext) GetLocalContextSchemaOk() (*string, bool)`

GetLocalContextSchemaOk returns a tuple with the LocalContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextSchema

`func (o *PatchedWritableDeviceWithConfigContext) SetLocalContextSchema(v string)`

SetLocalContextSchema sets LocalContextSchema field to given value.

### HasLocalContextSchema

`func (o *PatchedWritableDeviceWithConfigContext) HasLocalContextSchema() bool`

HasLocalContextSchema returns a boolean if a field has been set.

### SetLocalContextSchemaNil

`func (o *PatchedWritableDeviceWithConfigContext) SetLocalContextSchemaNil(b bool)`

 SetLocalContextSchemaNil sets the value for LocalContextSchema to be an explicit nil

### UnsetLocalContextSchema
`func (o *PatchedWritableDeviceWithConfigContext) UnsetLocalContextSchema()`

UnsetLocalContextSchema ensures that no value is present for LocalContextSchema, not even an explicit nil
### GetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContext) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableDeviceWithConfigContext) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContext) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableDeviceWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableDeviceWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableDeviceWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *PatchedWritableDeviceWithConfigContext) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeviceWithConfigContext) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeviceWithConfigContext) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeviceWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableDeviceWithConfigContext) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableDeviceWithConfigContext) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableDeviceWithConfigContext) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableDeviceWithConfigContext) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeviceWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeviceWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeviceWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeviceWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *PatchedWritableDeviceWithConfigContext) GetConfigContext() map[string]interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *PatchedWritableDeviceWithConfigContext) GetConfigContextOk() (*map[string]interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *PatchedWritableDeviceWithConfigContext) SetConfigContext(v map[string]interface{})`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *PatchedWritableDeviceWithConfigContext) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableDeviceWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableDeviceWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableDeviceWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableDeviceWithConfigContext) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedWritableDeviceWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedWritableDeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedWritableDeviceWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedWritableDeviceWithConfigContext) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableDeviceWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableDeviceWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableDeviceWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableDeviceWithConfigContext) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


