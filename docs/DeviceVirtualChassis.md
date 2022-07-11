# DeviceVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Url** | **string** |  | [readonly] 
**Master** | [**NestedDevice**](NestedDevice.md) |  | 
**MemberCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewDeviceVirtualChassis

`func NewDeviceVirtualChassis(id string, name string, url string, master NestedDevice, memberCount int32, display string, ) *DeviceVirtualChassis`

NewDeviceVirtualChassis instantiates a new DeviceVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceVirtualChassisWithDefaults

`func NewDeviceVirtualChassisWithDefaults() *DeviceVirtualChassis`

NewDeviceVirtualChassisWithDefaults instantiates a new DeviceVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceVirtualChassis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceVirtualChassis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceVirtualChassis) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *DeviceVirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceVirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceVirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMaster

`func (o *DeviceVirtualChassis) GetMaster() NestedDevice`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *DeviceVirtualChassis) GetMasterOk() (*NestedDevice, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *DeviceVirtualChassis) SetMaster(v NestedDevice)`

SetMaster sets Master field to given value.


### GetMemberCount

`func (o *DeviceVirtualChassis) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *DeviceVirtualChassis) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *DeviceVirtualChassis) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetDisplay

`func (o *DeviceVirtualChassis) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceVirtualChassis) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceVirtualChassis) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


