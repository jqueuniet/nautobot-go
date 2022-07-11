# DeviceWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**DeviceRole** | [**NestedDeviceRole**](NestedDeviceRole.md) |  | 
**Tenant** | Pointer to [**NullableAggregateTenant**](AggregateTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableDevicePlatform**](DevicePlatform.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | [**NestedSite**](NestedSite.md) |  | 
**Rack** | Pointer to [**NullableDeviceRack**](DeviceRack.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** | The lowest-numbered unit occupied by the device | [optional] 
**Face** | Pointer to [**DeviceFace**](DeviceFace.md) |  | [optional] 
**ParentDevice** | [**DeviceParentDevice**](DeviceParentDevice.md) |  | 
**Status** | [**DeviceStatus**](DeviceStatus.md) |  | 
**PrimaryIp** | [**DevicePrimaryIp**](DevicePrimaryIp.md) |  | 
**PrimaryIp4** | Pointer to [**NullableDevicePrimaryIp4**](DevicePrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableDevicePrimaryIp4**](DevicePrimaryIp4.md) |  | [optional] 
**SecretsGroup** | Pointer to [**NullableDeviceSecretsGroup**](DeviceSecretsGroup.md) |  | [optional] 
**Cluster** | Pointer to [**NullableDeviceCluster**](DeviceCluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableDeviceVirtualChassis**](DeviceVirtualChassis.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextSchema** | Pointer to [**NullableConfigContextSchema**](ConfigContextSchema.md) |  | [optional] 
**LocalContextData** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigContext** | **map[string]interface{}** |  | [readonly] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewDeviceWithConfigContext

`func NewDeviceWithConfigContext(id string, url string, deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite, parentDevice DeviceParentDevice, status DeviceStatus, primaryIp DevicePrimaryIp, configContext map[string]interface{}, created string, lastUpdated time.Time, display string, ) *DeviceWithConfigContext`

NewDeviceWithConfigContext instantiates a new DeviceWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithConfigContextWithDefaults

`func NewDeviceWithConfigContextWithDefaults() *DeviceWithConfigContext`

NewDeviceWithConfigContextWithDefaults instantiates a new DeviceWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceWithConfigContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceWithConfigContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceWithConfigContext) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *DeviceWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *DeviceWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceWithConfigContext) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceWithConfigContext) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *DeviceWithConfigContext) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceWithConfigContext) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceWithConfigContext) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetDeviceRole

`func (o *DeviceWithConfigContext) GetDeviceRole() NestedDeviceRole`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *DeviceWithConfigContext) GetDeviceRoleOk() (*NestedDeviceRole, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *DeviceWithConfigContext) SetDeviceRole(v NestedDeviceRole)`

SetDeviceRole sets DeviceRole field to given value.


### GetTenant

`func (o *DeviceWithConfigContext) GetTenant() AggregateTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DeviceWithConfigContext) GetTenantOk() (*AggregateTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DeviceWithConfigContext) SetTenant(v AggregateTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DeviceWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DeviceWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DeviceWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *DeviceWithConfigContext) GetPlatform() DevicePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceWithConfigContext) GetPlatformOk() (*DevicePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceWithConfigContext) SetPlatform(v DevicePlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *DeviceWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *DeviceWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *DeviceWithConfigContext) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DeviceWithConfigContext) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DeviceWithConfigContext) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DeviceWithConfigContext) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceWithConfigContext) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceWithConfigContext) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceWithConfigContext) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceWithConfigContext) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *DeviceWithConfigContext) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *DeviceWithConfigContext) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *DeviceWithConfigContext) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DeviceWithConfigContext) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DeviceWithConfigContext) SetSite(v NestedSite)`

SetSite sets Site field to given value.


### GetRack

`func (o *DeviceWithConfigContext) GetRack() DeviceRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DeviceWithConfigContext) GetRackOk() (*DeviceRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DeviceWithConfigContext) SetRack(v DeviceRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DeviceWithConfigContext) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DeviceWithConfigContext) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DeviceWithConfigContext) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *DeviceWithConfigContext) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DeviceWithConfigContext) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DeviceWithConfigContext) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DeviceWithConfigContext) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *DeviceWithConfigContext) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *DeviceWithConfigContext) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *DeviceWithConfigContext) GetFace() DeviceFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *DeviceWithConfigContext) GetFaceOk() (*DeviceFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *DeviceWithConfigContext) SetFace(v DeviceFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *DeviceWithConfigContext) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetParentDevice

`func (o *DeviceWithConfigContext) GetParentDevice() DeviceParentDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *DeviceWithConfigContext) GetParentDeviceOk() (*DeviceParentDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *DeviceWithConfigContext) SetParentDevice(v DeviceParentDevice)`

SetParentDevice sets ParentDevice field to given value.


### GetStatus

`func (o *DeviceWithConfigContext) GetStatus() DeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceWithConfigContext) GetStatusOk() (*DeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceWithConfigContext) SetStatus(v DeviceStatus)`

SetStatus sets Status field to given value.


### GetPrimaryIp

`func (o *DeviceWithConfigContext) GetPrimaryIp() DevicePrimaryIp`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *DeviceWithConfigContext) GetPrimaryIpOk() (*DevicePrimaryIp, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *DeviceWithConfigContext) SetPrimaryIp(v DevicePrimaryIp)`

SetPrimaryIp sets PrimaryIp field to given value.


### GetPrimaryIp4

`func (o *DeviceWithConfigContext) GetPrimaryIp4() DevicePrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp4Ok() (*DevicePrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *DeviceWithConfigContext) SetPrimaryIp4(v DevicePrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *DeviceWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *DeviceWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *DeviceWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *DeviceWithConfigContext) GetPrimaryIp6() DevicePrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp6Ok() (*DevicePrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *DeviceWithConfigContext) SetPrimaryIp6(v DevicePrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *DeviceWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *DeviceWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *DeviceWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetSecretsGroup

`func (o *DeviceWithConfigContext) GetSecretsGroup() DeviceSecretsGroup`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *DeviceWithConfigContext) GetSecretsGroupOk() (*DeviceSecretsGroup, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *DeviceWithConfigContext) SetSecretsGroup(v DeviceSecretsGroup)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *DeviceWithConfigContext) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *DeviceWithConfigContext) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *DeviceWithConfigContext) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCluster

`func (o *DeviceWithConfigContext) GetCluster() DeviceCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeviceWithConfigContext) GetClusterOk() (*DeviceCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeviceWithConfigContext) SetCluster(v DeviceCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeviceWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *DeviceWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeviceWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *DeviceWithConfigContext) GetVirtualChassis() DeviceVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *DeviceWithConfigContext) GetVirtualChassisOk() (*DeviceVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *DeviceWithConfigContext) SetVirtualChassis(v DeviceVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *DeviceWithConfigContext) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *DeviceWithConfigContext) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *DeviceWithConfigContext) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *DeviceWithConfigContext) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *DeviceWithConfigContext) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *DeviceWithConfigContext) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *DeviceWithConfigContext) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *DeviceWithConfigContext) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *DeviceWithConfigContext) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *DeviceWithConfigContext) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *DeviceWithConfigContext) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *DeviceWithConfigContext) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *DeviceWithConfigContext) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *DeviceWithConfigContext) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *DeviceWithConfigContext) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetComments

`func (o *DeviceWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeviceWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeviceWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeviceWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextSchema

`func (o *DeviceWithConfigContext) GetLocalContextSchema() ConfigContextSchema`

GetLocalContextSchema returns the LocalContextSchema field if non-nil, zero value otherwise.

### GetLocalContextSchemaOk

`func (o *DeviceWithConfigContext) GetLocalContextSchemaOk() (*ConfigContextSchema, bool)`

GetLocalContextSchemaOk returns a tuple with the LocalContextSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextSchema

`func (o *DeviceWithConfigContext) SetLocalContextSchema(v ConfigContextSchema)`

SetLocalContextSchema sets LocalContextSchema field to given value.

### HasLocalContextSchema

`func (o *DeviceWithConfigContext) HasLocalContextSchema() bool`

HasLocalContextSchema returns a boolean if a field has been set.

### SetLocalContextSchemaNil

`func (o *DeviceWithConfigContext) SetLocalContextSchemaNil(b bool)`

 SetLocalContextSchemaNil sets the value for LocalContextSchema to be an explicit nil

### UnsetLocalContextSchema
`func (o *DeviceWithConfigContext) UnsetLocalContextSchema()`

UnsetLocalContextSchema ensures that no value is present for LocalContextSchema, not even an explicit nil
### GetLocalContextData

`func (o *DeviceWithConfigContext) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *DeviceWithConfigContext) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *DeviceWithConfigContext) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *DeviceWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *DeviceWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *DeviceWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *DeviceWithConfigContext) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceWithConfigContext) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceWithConfigContext) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *DeviceWithConfigContext) GetConfigContext() map[string]interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *DeviceWithConfigContext) GetConfigContextOk() (*map[string]interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *DeviceWithConfigContext) SetConfigContext(v map[string]interface{})`

SetConfigContext sets ConfigContext field to given value.


### GetCreated

`func (o *DeviceWithConfigContext) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceWithConfigContext) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceWithConfigContext) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *DeviceWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplay

`func (o *DeviceWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


