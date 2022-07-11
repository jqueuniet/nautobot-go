# ConfigContextRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**DeviceCount** | **int32** |  | [readonly] 
**VirtualmachineCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewConfigContextRolesInner

`func NewConfigContextRolesInner(id string, url string, name string, deviceCount int32, virtualmachineCount int32, display string, ) *ConfigContextRolesInner`

NewConfigContextRolesInner instantiates a new ConfigContextRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextRolesInnerWithDefaults

`func NewConfigContextRolesInnerWithDefaults() *ConfigContextRolesInner`

NewConfigContextRolesInnerWithDefaults instantiates a new ConfigContextRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContextRolesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContextRolesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContextRolesInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConfigContextRolesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContextRolesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContextRolesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *ConfigContextRolesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContextRolesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContextRolesInner) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ConfigContextRolesInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ConfigContextRolesInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ConfigContextRolesInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ConfigContextRolesInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDeviceCount

`func (o *ConfigContextRolesInner) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *ConfigContextRolesInner) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *ConfigContextRolesInner) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.


### GetVirtualmachineCount

`func (o *ConfigContextRolesInner) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *ConfigContextRolesInner) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *ConfigContextRolesInner) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.


### GetDisplay

`func (o *ConfigContextRolesInner) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContextRolesInner) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContextRolesInner) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


