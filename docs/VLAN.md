# VLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Site** | Pointer to [**NullableCircuitTerminationSite**](CircuitTerminationSite.md) |  | [optional] 
**Group** | Pointer to [**NullableVLANGroup**](VLANGroup.md) |  | [optional] 
**Vid** | **int32** |  | 
**Name** | **string** |  | 
**Tenant** | Pointer to [**NullableAggregateTenant**](AggregateTenant.md) |  | [optional] 
**Status** | [**VLANStatus**](VLANStatus.md) |  | 
**Role** | Pointer to [**NullablePrefixRole**](PrefixRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**PrefixCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewVLAN

`func NewVLAN(id string, url string, vid int32, name string, status VLANStatus, created string, lastUpdated time.Time, prefixCount int32, display string, ) *VLAN`

NewVLAN instantiates a new VLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANWithDefaults

`func NewVLANWithDefaults() *VLAN`

NewVLANWithDefaults instantiates a new VLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VLAN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VLAN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VLAN) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *VLAN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VLAN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VLAN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSite

`func (o *VLAN) GetSite() CircuitTerminationSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VLAN) GetSiteOk() (*CircuitTerminationSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VLAN) SetSite(v CircuitTerminationSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *VLAN) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *VLAN) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *VLAN) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *VLAN) GetGroup() VLANGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VLAN) GetGroupOk() (*VLANGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VLAN) SetGroup(v VLANGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VLAN) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *VLAN) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *VLAN) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *VLAN) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *VLAN) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *VLAN) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *VLAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VLAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VLAN) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *VLAN) GetTenant() AggregateTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VLAN) GetTenantOk() (*AggregateTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VLAN) SetTenant(v AggregateTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VLAN) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VLAN) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VLAN) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *VLAN) GetStatus() VLANStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VLAN) GetStatusOk() (*VLANStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VLAN) SetStatus(v VLANStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *VLAN) GetRole() PrefixRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VLAN) GetRoleOk() (*PrefixRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VLAN) SetRole(v PrefixRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VLAN) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VLAN) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VLAN) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *VLAN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VLAN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VLAN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VLAN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *VLAN) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VLAN) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VLAN) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VLAN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VLAN) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VLAN) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VLAN) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VLAN) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *VLAN) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VLAN) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VLAN) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *VLAN) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VLAN) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VLAN) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetPrefixCount

`func (o *VLAN) GetPrefixCount() int32`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *VLAN) GetPrefixCountOk() (*int32, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *VLAN) SetPrefixCount(v int32)`

SetPrefixCount sets PrefixCount field to given value.


### GetDisplay

`func (o *VLAN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VLAN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VLAN) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


