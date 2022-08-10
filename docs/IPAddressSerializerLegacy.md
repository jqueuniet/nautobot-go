# IPAddressSerializerLegacy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Family** | [**AggregateFamily**](AggregateFamily.md) |  | 
**Address** | **string** |  | 
**Vrf** | Pointer to [**NullableIPAddressVrf**](IPAddressVrf.md) |  | [optional] 
**Tenant** | Pointer to [**NullableAggregateTenant**](AggregateTenant.md) |  | [optional] 
**Status** | [**IPAddressStatus**](IPAddressStatus.md) |  | 
**Role** | Pointer to [**IPAddressRole**](IPAddressRole.md) |  | [optional] 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableString** |  | [optional] 
**AssignedObject** | **map[string]interface{}** |  | [readonly] 
**NatInside** | Pointer to [**NullableDevicePrimaryIp4**](DevicePrimaryIp4.md) |  | [optional] 
**NatOutside** | [**DevicePrimaryIp**](DevicePrimaryIp.md) |  | 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewIPAddressSerializerLegacy

`func NewIPAddressSerializerLegacy(id string, url string, family AggregateFamily, address string, status IPAddressStatus, assignedObject map[string]interface{}, natOutside DevicePrimaryIp, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string, ) *IPAddressSerializerLegacy`

NewIPAddressSerializerLegacy instantiates a new IPAddressSerializerLegacy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressSerializerLegacyWithDefaults

`func NewIPAddressSerializerLegacyWithDefaults() *IPAddressSerializerLegacy`

NewIPAddressSerializerLegacyWithDefaults instantiates a new IPAddressSerializerLegacy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddressSerializerLegacy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddressSerializerLegacy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddressSerializerLegacy) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *IPAddressSerializerLegacy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddressSerializerLegacy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddressSerializerLegacy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFamily

`func (o *IPAddressSerializerLegacy) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *IPAddressSerializerLegacy) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *IPAddressSerializerLegacy) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *IPAddressSerializerLegacy) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddressSerializerLegacy) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddressSerializerLegacy) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVrf

`func (o *IPAddressSerializerLegacy) GetVrf() IPAddressVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IPAddressSerializerLegacy) GetVrfOk() (*IPAddressVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IPAddressSerializerLegacy) SetVrf(v IPAddressVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IPAddressSerializerLegacy) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *IPAddressSerializerLegacy) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *IPAddressSerializerLegacy) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *IPAddressSerializerLegacy) GetTenant() AggregateTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddressSerializerLegacy) GetTenantOk() (*AggregateTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddressSerializerLegacy) SetTenant(v AggregateTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddressSerializerLegacy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddressSerializerLegacy) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddressSerializerLegacy) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *IPAddressSerializerLegacy) GetStatus() IPAddressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddressSerializerLegacy) GetStatusOk() (*IPAddressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddressSerializerLegacy) SetStatus(v IPAddressStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *IPAddressSerializerLegacy) GetRole() IPAddressRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddressSerializerLegacy) GetRoleOk() (*IPAddressRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddressSerializerLegacy) SetRole(v IPAddressRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddressSerializerLegacy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *IPAddressSerializerLegacy) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *IPAddressSerializerLegacy) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *IPAddressSerializerLegacy) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *IPAddressSerializerLegacy) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *IPAddressSerializerLegacy) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *IPAddressSerializerLegacy) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *IPAddressSerializerLegacy) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *IPAddressSerializerLegacy) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *IPAddressSerializerLegacy) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *IPAddressSerializerLegacy) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *IPAddressSerializerLegacy) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *IPAddressSerializerLegacy) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *IPAddressSerializerLegacy) GetAssignedObject() map[string]interface{}`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *IPAddressSerializerLegacy) GetAssignedObjectOk() (*map[string]interface{}, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *IPAddressSerializerLegacy) SetAssignedObject(v map[string]interface{})`

SetAssignedObject sets AssignedObject field to given value.


### SetAssignedObjectNil

`func (o *IPAddressSerializerLegacy) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *IPAddressSerializerLegacy) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetNatInside

`func (o *IPAddressSerializerLegacy) GetNatInside() DevicePrimaryIp4`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *IPAddressSerializerLegacy) GetNatInsideOk() (*DevicePrimaryIp4, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *IPAddressSerializerLegacy) SetNatInside(v DevicePrimaryIp4)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *IPAddressSerializerLegacy) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *IPAddressSerializerLegacy) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *IPAddressSerializerLegacy) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetNatOutside

`func (o *IPAddressSerializerLegacy) GetNatOutside() DevicePrimaryIp`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *IPAddressSerializerLegacy) GetNatOutsideOk() (*DevicePrimaryIp, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *IPAddressSerializerLegacy) SetNatOutside(v DevicePrimaryIp)`

SetNatOutside sets NatOutside field to given value.


### GetDnsName

`func (o *IPAddressSerializerLegacy) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IPAddressSerializerLegacy) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IPAddressSerializerLegacy) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IPAddressSerializerLegacy) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *IPAddressSerializerLegacy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddressSerializerLegacy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddressSerializerLegacy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddressSerializerLegacy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *IPAddressSerializerLegacy) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddressSerializerLegacy) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddressSerializerLegacy) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddressSerializerLegacy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPAddressSerializerLegacy) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddressSerializerLegacy) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddressSerializerLegacy) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddressSerializerLegacy) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *IPAddressSerializerLegacy) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPAddressSerializerLegacy) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPAddressSerializerLegacy) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *IPAddressSerializerLegacy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPAddressSerializerLegacy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPAddressSerializerLegacy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetComputedFields

`func (o *IPAddressSerializerLegacy) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *IPAddressSerializerLegacy) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *IPAddressSerializerLegacy) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.


### GetDisplay

`func (o *IPAddressSerializerLegacy) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddressSerializerLegacy) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddressSerializerLegacy) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


