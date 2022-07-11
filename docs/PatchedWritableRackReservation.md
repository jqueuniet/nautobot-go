# PatchedWritableRackReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Rack** | Pointer to **string** |  | [optional] 
**Units** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**User** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputedFields** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedWritableRackReservation

`func NewPatchedWritableRackReservation() *PatchedWritableRackReservation`

NewPatchedWritableRackReservation instantiates a new PatchedWritableRackReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableRackReservationWithDefaults

`func NewPatchedWritableRackReservationWithDefaults() *PatchedWritableRackReservation`

NewPatchedWritableRackReservationWithDefaults instantiates a new PatchedWritableRackReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritableRackReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritableRackReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritableRackReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritableRackReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedWritableRackReservation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedWritableRackReservation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedWritableRackReservation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedWritableRackReservation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRack

`func (o *PatchedWritableRackReservation) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableRackReservation) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableRackReservation) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableRackReservation) HasRack() bool`

HasRack returns a boolean if a field has been set.

### GetUnits

`func (o *PatchedWritableRackReservation) GetUnits() map[string]interface{}`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PatchedWritableRackReservation) GetUnitsOk() (*map[string]interface{}, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PatchedWritableRackReservation) SetUnits(v map[string]interface{})`

SetUnits sets Units field to given value.

### HasUnits

`func (o *PatchedWritableRackReservation) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedWritableRackReservation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedWritableRackReservation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedWritableRackReservation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedWritableRackReservation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUser

`func (o *PatchedWritableRackReservation) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedWritableRackReservation) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedWritableRackReservation) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedWritableRackReservation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableRackReservation) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableRackReservation) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableRackReservation) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableRackReservation) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableRackReservation) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableRackReservation) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *PatchedWritableRackReservation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableRackReservation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableRackReservation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableRackReservation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableRackReservation) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableRackReservation) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableRackReservation) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableRackReservation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableRackReservation) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableRackReservation) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableRackReservation) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableRackReservation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComputedFields

`func (o *PatchedWritableRackReservation) GetComputedFields() map[string]interface{}`

GetComputedFields returns the ComputedFields field if non-nil, zero value otherwise.

### GetComputedFieldsOk

`func (o *PatchedWritableRackReservation) GetComputedFieldsOk() (*map[string]interface{}, bool)`

GetComputedFieldsOk returns a tuple with the ComputedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedFields

`func (o *PatchedWritableRackReservation) SetComputedFields(v map[string]interface{})`

SetComputedFields sets ComputedFields field to given value.

### HasComputedFields

`func (o *PatchedWritableRackReservation) HasComputedFields() bool`

HasComputedFields returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedWritableRackReservation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedWritableRackReservation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedWritableRackReservation) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedWritableRackReservation) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


