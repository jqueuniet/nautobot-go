# DeviceBayInstalledDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewDeviceBayInstalledDevice

`func NewDeviceBayInstalledDevice(id string, url string, display string, ) *DeviceBayInstalledDevice`

NewDeviceBayInstalledDevice instantiates a new DeviceBayInstalledDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBayInstalledDeviceWithDefaults

`func NewDeviceBayInstalledDeviceWithDefaults() *DeviceBayInstalledDevice`

NewDeviceBayInstalledDeviceWithDefaults instantiates a new DeviceBayInstalledDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceBayInstalledDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceBayInstalledDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceBayInstalledDevice) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *DeviceBayInstalledDevice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceBayInstalledDevice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceBayInstalledDevice) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *DeviceBayInstalledDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBayInstalledDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBayInstalledDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceBayInstalledDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceBayInstalledDevice) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceBayInstalledDevice) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplay

`func (o *DeviceBayInstalledDevice) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceBayInstalledDevice) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceBayInstalledDevice) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


