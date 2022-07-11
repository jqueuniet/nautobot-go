# ConfigContextDeviceTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Manufacturer** | [**ConfigContextDeviceTypesInnerManufacturer**](ConfigContextDeviceTypesInnerManufacturer.md) |  | 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**DeviceCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewConfigContextDeviceTypesInner

`func NewConfigContextDeviceTypesInner(id string, url string, manufacturer ConfigContextDeviceTypesInnerManufacturer, model string, slug string, deviceCount int32, display string, ) *ConfigContextDeviceTypesInner`

NewConfigContextDeviceTypesInner instantiates a new ConfigContextDeviceTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextDeviceTypesInnerWithDefaults

`func NewConfigContextDeviceTypesInnerWithDefaults() *ConfigContextDeviceTypesInner`

NewConfigContextDeviceTypesInnerWithDefaults instantiates a new ConfigContextDeviceTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContextDeviceTypesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContextDeviceTypesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContextDeviceTypesInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConfigContextDeviceTypesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContextDeviceTypesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContextDeviceTypesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetManufacturer

`func (o *ConfigContextDeviceTypesInner) GetManufacturer() ConfigContextDeviceTypesInnerManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ConfigContextDeviceTypesInner) GetManufacturerOk() (*ConfigContextDeviceTypesInnerManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ConfigContextDeviceTypesInner) SetManufacturer(v ConfigContextDeviceTypesInnerManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *ConfigContextDeviceTypesInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConfigContextDeviceTypesInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConfigContextDeviceTypesInner) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *ConfigContextDeviceTypesInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ConfigContextDeviceTypesInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ConfigContextDeviceTypesInner) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDeviceCount

`func (o *ConfigContextDeviceTypesInner) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *ConfigContextDeviceTypesInner) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *ConfigContextDeviceTypesInner) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.


### GetDisplay

`func (o *ConfigContextDeviceTypesInner) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContextDeviceTypesInner) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContextDeviceTypesInner) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


