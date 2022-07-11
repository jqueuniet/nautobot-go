# Prefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Family** | [**AggregateFamily**](AggregateFamily.md) |  | 
**Prefix** | **string** |  | 
**Site** | Pointer to [**NullableCircuitTerminationSite**](CircuitTerminationSite.md) |  | [optional] 
**Vrf** | Pointer to [**NullableIPAddressSerializerLegacyVrf**](IPAddressSerializerLegacyVrf.md) |  | [optional] 
**Tenant** | Pointer to [**NullableAggregateTenant**](AggregateTenant.md) |  | [optional] 
**Vlan** | Pointer to [**NullableInterfaceUntaggedVlan**](InterfaceUntaggedVlan.md) |  | [optional] 
**Status** | [**PrefixStatus**](PrefixStatus.md) |  | 
**Role** | Pointer to [**NullablePrefixRole**](PrefixRole.md) |  | [optional] 
**IsPool** | Pointer to **bool** | All IP addresses within this prefix are considered usable | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewPrefix

`func NewPrefix(id string, url string, family AggregateFamily, prefix string, status PrefixStatus, created string, lastUpdated time.Time, display string, ) *Prefix`

NewPrefix instantiates a new Prefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixWithDefaults

`func NewPrefixWithDefaults() *Prefix`

NewPrefixWithDefaults instantiates a new Prefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Prefix) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prefix) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prefix) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *Prefix) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Prefix) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Prefix) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFamily

`func (o *Prefix) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Prefix) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Prefix) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetPrefix

`func (o *Prefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Prefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Prefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetSite

`func (o *Prefix) GetSite() CircuitTerminationSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Prefix) GetSiteOk() (*CircuitTerminationSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Prefix) SetSite(v CircuitTerminationSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *Prefix) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *Prefix) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *Prefix) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetVrf

`func (o *Prefix) GetVrf() IPAddressSerializerLegacyVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *Prefix) GetVrfOk() (*IPAddressSerializerLegacyVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *Prefix) SetVrf(v IPAddressSerializerLegacyVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *Prefix) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *Prefix) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *Prefix) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *Prefix) GetTenant() AggregateTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Prefix) GetTenantOk() (*AggregateTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Prefix) SetTenant(v AggregateTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Prefix) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Prefix) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Prefix) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *Prefix) GetVlan() InterfaceUntaggedVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *Prefix) GetVlanOk() (*InterfaceUntaggedVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *Prefix) SetVlan(v InterfaceUntaggedVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *Prefix) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *Prefix) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *Prefix) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetStatus

`func (o *Prefix) GetStatus() PrefixStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Prefix) GetStatusOk() (*PrefixStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Prefix) SetStatus(v PrefixStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *Prefix) GetRole() PrefixRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Prefix) GetRoleOk() (*PrefixRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Prefix) SetRole(v PrefixRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Prefix) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Prefix) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Prefix) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetIsPool

`func (o *Prefix) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *Prefix) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *Prefix) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.

### HasIsPool

`func (o *Prefix) HasIsPool() bool`

HasIsPool returns a boolean if a field has been set.

### GetDescription

`func (o *Prefix) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Prefix) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Prefix) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Prefix) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Prefix) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Prefix) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Prefix) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Prefix) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Prefix) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Prefix) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Prefix) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Prefix) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Prefix) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Prefix) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Prefix) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *Prefix) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Prefix) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Prefix) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplay

`func (o *Prefix) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Prefix) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Prefix) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


