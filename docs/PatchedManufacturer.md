# PatchedManufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevicetypeCount** | Pointer to **int32** |  | [optional] [readonly] 
**InventoryitemCount** | Pointer to **int32** |  | [optional] [readonly] 
**PlatformCount** | Pointer to **int32** |  | [optional] [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedManufacturer

`func NewPatchedManufacturer() *PatchedManufacturer`

NewPatchedManufacturer instantiates a new PatchedManufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedManufacturerWithDefaults

`func NewPatchedManufacturerWithDefaults() *PatchedManufacturer`

NewPatchedManufacturerWithDefaults instantiates a new PatchedManufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedManufacturer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedManufacturer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedManufacturer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedManufacturer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedManufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedManufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedManufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedManufacturer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedManufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedManufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedManufacturer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedManufacturer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedManufacturer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedManufacturer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedManufacturer) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedManufacturer) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedManufacturer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedManufacturer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedManufacturer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedManufacturer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevicetypeCount

`func (o *PatchedManufacturer) GetDevicetypeCount() int32`

GetDevicetypeCount returns the DevicetypeCount field if non-nil, zero value otherwise.

### GetDevicetypeCountOk

`func (o *PatchedManufacturer) GetDevicetypeCountOk() (*int32, bool)`

GetDevicetypeCountOk returns a tuple with the DevicetypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetypeCount

`func (o *PatchedManufacturer) SetDevicetypeCount(v int32)`

SetDevicetypeCount sets DevicetypeCount field to given value.

### HasDevicetypeCount

`func (o *PatchedManufacturer) HasDevicetypeCount() bool`

HasDevicetypeCount returns a boolean if a field has been set.

### GetInventoryitemCount

`func (o *PatchedManufacturer) GetInventoryitemCount() int32`

GetInventoryitemCount returns the InventoryitemCount field if non-nil, zero value otherwise.

### GetInventoryitemCountOk

`func (o *PatchedManufacturer) GetInventoryitemCountOk() (*int32, bool)`

GetInventoryitemCountOk returns a tuple with the InventoryitemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemCount

`func (o *PatchedManufacturer) SetInventoryitemCount(v int32)`

SetInventoryitemCount sets InventoryitemCount field to given value.

### HasInventoryitemCount

`func (o *PatchedManufacturer) HasInventoryitemCount() bool`

HasInventoryitemCount returns a boolean if a field has been set.

### GetPlatformCount

`func (o *PatchedManufacturer) GetPlatformCount() int32`

GetPlatformCount returns the PlatformCount field if non-nil, zero value otherwise.

### GetPlatformCountOk

`func (o *PatchedManufacturer) GetPlatformCountOk() (*int32, bool)`

GetPlatformCountOk returns a tuple with the PlatformCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCount

`func (o *PatchedManufacturer) SetPlatformCount(v int32)`

SetPlatformCount sets PlatformCount field to given value.

### HasPlatformCount

`func (o *PatchedManufacturer) HasPlatformCount() bool`

HasPlatformCount returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedManufacturer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedManufacturer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedManufacturer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedManufacturer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedManufacturer) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedManufacturer) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedManufacturer) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedManufacturer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedManufacturer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedManufacturer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedManufacturer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedManufacturer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedManufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedManufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedManufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedManufacturer) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


