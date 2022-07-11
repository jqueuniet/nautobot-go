# Manufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DevicetypeCount** | **int32** |  | [readonly] 
**InventoryitemCount** | **int32** |  | [readonly] 
**PlatformCount** | **int32** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewManufacturer

`func NewManufacturer(id string, url string, name string, devicetypeCount int32, inventoryitemCount int32, platformCount int32, created string, lastUpdated time.Time, display string, ) *Manufacturer`

NewManufacturer instantiates a new Manufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturerWithDefaults

`func NewManufacturerWithDefaults() *Manufacturer`

NewManufacturerWithDefaults instantiates a new Manufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manufacturer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manufacturer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manufacturer) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *Manufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Manufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Manufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *Manufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Manufacturer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Manufacturer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Manufacturer) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Manufacturer) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *Manufacturer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manufacturer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manufacturer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manufacturer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevicetypeCount

`func (o *Manufacturer) GetDevicetypeCount() int32`

GetDevicetypeCount returns the DevicetypeCount field if non-nil, zero value otherwise.

### GetDevicetypeCountOk

`func (o *Manufacturer) GetDevicetypeCountOk() (*int32, bool)`

GetDevicetypeCountOk returns a tuple with the DevicetypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetypeCount

`func (o *Manufacturer) SetDevicetypeCount(v int32)`

SetDevicetypeCount sets DevicetypeCount field to given value.


### GetInventoryitemCount

`func (o *Manufacturer) GetInventoryitemCount() int32`

GetInventoryitemCount returns the InventoryitemCount field if non-nil, zero value otherwise.

### GetInventoryitemCountOk

`func (o *Manufacturer) GetInventoryitemCountOk() (*int32, bool)`

GetInventoryitemCountOk returns a tuple with the InventoryitemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryitemCount

`func (o *Manufacturer) SetInventoryitemCount(v int32)`

SetInventoryitemCount sets InventoryitemCount field to given value.


### GetPlatformCount

`func (o *Manufacturer) GetPlatformCount() int32`

GetPlatformCount returns the PlatformCount field if non-nil, zero value otherwise.

### GetPlatformCountOk

`func (o *Manufacturer) GetPlatformCountOk() (*int32, bool)`

GetPlatformCountOk returns a tuple with the PlatformCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCount

`func (o *Manufacturer) SetPlatformCount(v int32)`

SetPlatformCount sets PlatformCount field to given value.


### GetCustomFields

`func (o *Manufacturer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Manufacturer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Manufacturer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Manufacturer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Manufacturer) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Manufacturer) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Manufacturer) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *Manufacturer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Manufacturer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Manufacturer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplay

`func (o *Manufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Manufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Manufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


