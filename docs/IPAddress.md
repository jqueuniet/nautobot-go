# IPAddress

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
**NatOutside** | [**[]NestedIPAddress**](NestedIPAddress.md) |  | [readonly] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**ComputedFields** | **map[string]interface{}** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewIPAddress

`func NewIPAddress(id string, url string, family AggregateFamily, address string, status IPAddressStatus, assignedObject map[string]interface{}, natOutside []NestedIPAddress, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string, ) *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddress) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *IPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFamily

`func (o *IPAddress) GetFamily() AggregateFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *IPAddress) GetFamilyOk() (*AggregateFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *IPAddress) SetFamily(v AggregateFamily)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *IPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVrf

`func (o *IPAddress) GetVrf() IPAddressVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IPAddress) GetVrfOk() (*IPAddressVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IPAddress) SetVrf(v IPAddressVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IPAddress) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *IPAddress) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *IPAddress) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *IPAddress) GetTenant() AggregateTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddress) GetTenantOk() (*AggregateTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddress) SetTenant(v AggregateTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddress) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddress) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddress) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *IPAddress) GetStatus() IPAddressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddress) GetStatusOk() (*IPAddressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddress) SetStatus(v IPAddressStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *IPAddress) GetRole() IPAddressRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddress) GetRoleOk() (*IPAddressRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddress) SetRole(v IPAddressRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddress) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *IPAddress) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *IPAddress) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *IPAddress) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *IPAddress) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *IPAddress) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *IPAddress) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *IPAddress) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *IPAddress) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *IPAddress) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *IPAddress) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *IPAddress) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *IPAddress) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *IPAddress) GetAssignedObject() map[string]interface{}`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *IPAddress) GetAssignedObjectOk() (*map[string]interface{}, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *IPAddress) SetAssignedObject(v map[string]interface{})`

SetAssignedObject sets AssignedObject field to given value.


### SetAssignedObjectNil

`func (o *IPAddress) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *IPAddress) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetNatInside

`func (o *IPAddress) GetNatInside() DevicePrimaryIp4`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *IPAddress) GetNatInsideOk() (*DevicePrimaryIp4, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *IPAddress) SetNatInside(v DevicePrimaryIp4)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *IPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *IPAddress) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *IPAddress) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetNatOutside

`func (o *IPAddress) GetNatOutside() []NestedIPAddress`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *IPAddress) GetNatOutsideOk() (*[]NestedIPAddress, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *IPAddress) SetNatOutside(v []NestedIPAddress)`

SetNatOutside sets NatOutside field to given value.


### GetDnsName

`func (o *IPAddress) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IPAddress) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IPAddress) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *IPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *IPAddress) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddress) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddress) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *IPAddress) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPAddress) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPAddress) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *IPAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetComputedFields

`func (o *IPAddress) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *IPAddress) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *IPAddress) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.


### GetDisplay

`func (o *IPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


