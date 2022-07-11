# PatchedWritableAggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Family** | Pointer to [**NullableFamilyEnum**](FamilyEnum.md) |  | [optional] [readonly] 
**Prefix** | Pointer to **string** |  | [optional] 
**Rir** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**DateAdded** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableAggregate

`func NewPatchedWritableAggregate() *PatchedWritableAggregate`

NewPatchedWritableAggregate instantiates a new PatchedWritableAggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableAggregateWithDefaults

`func NewPatchedWritableAggregateWithDefaults() *PatchedWritableAggregate`

NewPatchedWritableAggregateWithDefaults instantiates a new PatchedWritableAggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableAggregate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableAggregate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableAggregate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableAggregate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableAggregate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableAggregate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableAggregate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableAggregate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFamily

`func (o *PatchedWritableAggregate) GetFamily() FamilyEnum`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *PatchedWritableAggregate) GetFamilyOk() (*FamilyEnum, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *PatchedWritableAggregate) SetFamily(v FamilyEnum)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *PatchedWritableAggregate) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### SetFamilyNil

`func (o *PatchedWritableAggregate) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *PatchedWritableAggregate) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetPrefix

`func (o *PatchedWritableAggregate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedWritableAggregate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedWritableAggregate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedWritableAggregate) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRir

`func (o *PatchedWritableAggregate) GetRir() string`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *PatchedWritableAggregate) GetRirOk() (*string, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *PatchedWritableAggregate) SetRir(v string)`

SetRir sets Rir field to given value.

### HasRir

`func (o *PatchedWritableAggregate) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableAggregate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableAggregate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableAggregate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableAggregate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableAggregate) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableAggregate) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDateAdded

`func (o *PatchedWritableAggregate) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *PatchedWritableAggregate) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *PatchedWritableAggregate) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *PatchedWritableAggregate) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### SetDateAddedNil

`func (o *PatchedWritableAggregate) SetDateAddedNil(b bool)`

 SetDateAddedNil sets the value for DateAdded to be an explicit nil

### UnsetDateAdded
`func (o *PatchedWritableAggregate) UnsetDateAdded()`

UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
### GetDescription

`func (o *PatchedWritableAggregate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableAggregate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableAggregate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableAggregate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableAggregate) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableAggregate) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableAggregate) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableAggregate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableAggregate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableAggregate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableAggregate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableAggregate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableAggregate) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableAggregate) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableAggregate) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableAggregate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedWritableAggregate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedWritableAggregate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedWritableAggregate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedWritableAggregate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableAggregate) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableAggregate) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableAggregate) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableAggregate) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableAggregate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableAggregate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableAggregate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableAggregate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


