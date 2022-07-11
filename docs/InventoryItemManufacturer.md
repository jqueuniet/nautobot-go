# InventoryItemManufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**DevicetypeCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewInventoryItemManufacturer

`func NewInventoryItemManufacturer(id string, url string, name string, devicetypeCount int32, display string, ) *InventoryItemManufacturer`

NewInventoryItemManufacturer instantiates a new InventoryItemManufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemManufacturerWithDefaults

`func NewInventoryItemManufacturerWithDefaults() *InventoryItemManufacturer`

NewInventoryItemManufacturerWithDefaults instantiates a new InventoryItemManufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryItemManufacturer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItemManufacturer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItemManufacturer) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *InventoryItemManufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InventoryItemManufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InventoryItemManufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *InventoryItemManufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItemManufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItemManufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *InventoryItemManufacturer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *InventoryItemManufacturer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *InventoryItemManufacturer) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *InventoryItemManufacturer) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDevicetypeCount

`func (o *InventoryItemManufacturer) GetDevicetypeCount() int32`

GetDevicetypeCount returns the DevicetypeCount field if non-nil, zero value otherwise.

### GetDevicetypeCountOk

`func (o *InventoryItemManufacturer) GetDevicetypeCountOk() (*int32, bool)`

GetDevicetypeCountOk returns a tuple with the DevicetypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetypeCount

`func (o *InventoryItemManufacturer) SetDevicetypeCount(v int32)`

SetDevicetypeCount sets DevicetypeCount field to given value.


### GetDisplay

`func (o *InventoryItemManufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InventoryItemManufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InventoryItemManufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


