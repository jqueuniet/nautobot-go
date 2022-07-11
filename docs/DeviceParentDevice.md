# DeviceParentDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewDeviceParentDevice

`func NewDeviceParentDevice(id string, url string, display string, ) *DeviceParentDevice`

NewDeviceParentDevice instantiates a new DeviceParentDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceParentDeviceWithDefaults

`func NewDeviceParentDeviceWithDefaults() *DeviceParentDevice`

NewDeviceParentDeviceWithDefaults instantiates a new DeviceParentDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceParentDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceParentDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceParentDevice) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *DeviceParentDevice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceParentDevice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceParentDevice) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *DeviceParentDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceParentDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceParentDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceParentDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceParentDevice) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceParentDevice) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplay

`func (o *DeviceParentDevice) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceParentDevice) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceParentDevice) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


