# PatchedWritableIPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to [**NullableFamilyEnum**](FamilyEnum.md) |  | [optional] [readonly] 
**Address** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to **NullableString** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**WritableIPAddressStatusEnum**](WritableIPAddressStatusEnum.md) |  | [optional] 
**Role** | Pointer to [**PatchedWritableIPAddressRole**](PatchedWritableIPAddressRole.md) |  | [optional] 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableString** |  | [optional] 
**AssignedObject** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**NatInside** | Pointer to **NullableString** | The IP Addresses for which this address is the \&quot;outside\&quot; IP | [optional] 
**NatOutside** | Pointer to [**[]NestedIPAddress**](NestedIPAddress.md) |  | [optional] [readonly] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableIPAddress

`func NewPatchedWritableIPAddress() *PatchedWritableIPAddress`

NewPatchedWritableIPAddress instantiates a new PatchedWritableIPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIPAddressWithDefaults

`func NewPatchedWritableIPAddressWithDefaults() *PatchedWritableIPAddress`

NewPatchedWritableIPAddressWithDefaults instantiates a new PatchedWritableIPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableIPAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableIPAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableIPAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableIPAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableIPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableIPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableIPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableIPAddress) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFamily

`func (o *PatchedWritableIPAddress) GetFamily() FamilyEnum`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *PatchedWritableIPAddress) GetFamilyOk() (*FamilyEnum, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *PatchedWritableIPAddress) SetFamily(v FamilyEnum)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *PatchedWritableIPAddress) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *PatchedWritableIPAddress) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *PatchedWritableIPAddress) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetAddress

`func (o *PatchedWritableIPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedWritableIPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedWritableIPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedWritableIPAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedWritableIPAddress) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedWritableIPAddress) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedWritableIPAddress) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedWritableIPAddress) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedWritableIPAddress) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedWritableIPAddress) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *PatchedWritableIPAddress) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableIPAddress) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableIPAddress) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableIPAddress) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableIPAddress) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableIPAddress) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedWritableIPAddress) GetStatus() WritableIPAddressStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableIPAddress) GetStatusOk() (*WritableIPAddressStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableIPAddress) SetStatus(v WritableIPAddressStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableIPAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableIPAddress) GetRole() PatchedWritableIPAddressRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableIPAddress) GetRoleOk() (*PatchedWritableIPAddressRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableIPAddress) SetRole(v PatchedWritableIPAddressRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableIPAddress) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *PatchedWritableIPAddress) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedWritableIPAddress) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedWritableIPAddress) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedWritableIPAddress) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *PatchedWritableIPAddress) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *PatchedWritableIPAddress) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *PatchedWritableIPAddress) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedWritableIPAddress) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedWritableIPAddress) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedWritableIPAddress) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *PatchedWritableIPAddress) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *PatchedWritableIPAddress) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *PatchedWritableIPAddress) GetAssignedObject() map[string]interface{}`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *PatchedWritableIPAddress) GetAssignedObjectOk() (*map[string]interface{}, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *PatchedWritableIPAddress) SetAssignedObject(v map[string]interface{})`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *PatchedWritableIPAddress) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### SetAssignedObjectNil

`func (o *PatchedWritableIPAddress) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *PatchedWritableIPAddress) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetNatInside

`func (o *PatchedWritableIPAddress) GetNatInside() string`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *PatchedWritableIPAddress) GetNatInsideOk() (*string, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *PatchedWritableIPAddress) SetNatInside(v string)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *PatchedWritableIPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *PatchedWritableIPAddress) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *PatchedWritableIPAddress) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetNatOutside

`func (o *PatchedWritableIPAddress) GetNatOutside() []NestedIPAddress`

GetNatOutside returns the NatOutside field if non-nil, zero value otherwise.

### GetNatOutsideOk

`func (o *PatchedWritableIPAddress) GetNatOutsideOk() (*[]NestedIPAddress, bool)`

GetNatOutsideOk returns a tuple with the NatOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutside

`func (o *PatchedWritableIPAddress) SetNatOutside(v []NestedIPAddress)`

SetNatOutside sets NatOutside field to given value.

### HasNatOutside

`func (o *PatchedWritableIPAddress) HasNatOutside() bool`

HasNatOutside returns a boolean if a field has been set.

### GetDnsName

`func (o *PatchedWritableIPAddress) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *PatchedWritableIPAddress) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *PatchedWritableIPAddress) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *PatchedWritableIPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIPAddress) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIPAddress) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIPAddress) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIPAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIPAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIPAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIPAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIPAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableIPAddress) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableIPAddress) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableIPAddress) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableIPAddress) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedWritableIPAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedWritableIPAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedWritableIPAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedWritableIPAddress) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableIPAddress) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableIPAddress) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableIPAddress) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableIPAddress) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableIPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableIPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableIPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableIPAddress) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


