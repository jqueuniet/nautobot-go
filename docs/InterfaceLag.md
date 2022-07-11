# InterfaceLag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Device** | [**DeviceParentDevice**](DeviceParentDevice.md) |  | 
**Name** | **string** |  | 
**Cable** | Pointer to **NullableString** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewInterfaceLag

`func NewInterfaceLag(id string, url string, device DeviceParentDevice, name string, display string, ) *InterfaceLag`

NewInterfaceLag instantiates a new InterfaceLag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceLagWithDefaults

`func NewInterfaceLagWithDefaults() *InterfaceLag`

NewInterfaceLagWithDefaults instantiates a new InterfaceLag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceLag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceLag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceLag) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *InterfaceLag) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InterfaceLag) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InterfaceLag) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDevice

`func (o *InterfaceLag) GetDevice() DeviceParentDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InterfaceLag) GetDeviceOk() (*DeviceParentDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InterfaceLag) SetDevice(v DeviceParentDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *InterfaceLag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceLag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceLag) SetName(v string)`

SetName sets Name field to given value.


### GetCable

`func (o *InterfaceLag) GetCable() string`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *InterfaceLag) GetCableOk() (*string, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *InterfaceLag) SetCable(v string)`

SetCable sets Cable field to given value.

### HasCable

`func (o *InterfaceLag) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *InterfaceLag) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *InterfaceLag) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetDisplay

`func (o *InterfaceLag) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceLag) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceLag) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


