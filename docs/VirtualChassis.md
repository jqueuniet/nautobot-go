# VirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Master** | Pointer to [**NullableDeviceBayInstalledDevice**](DeviceBayInstalledDevice.md) |  | [optional] 
**Tags** | Pointer to [**[]TagSerializerField**](TagSerializerField.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**MemberCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewVirtualChassis

`func NewVirtualChassis(id string, url string, name string, memberCount int32, display string, ) *VirtualChassis`

NewVirtualChassis instantiates a new VirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisWithDefaults

`func NewVirtualChassisWithDefaults() *VirtualChassis`

NewVirtualChassisWithDefaults instantiates a new VirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualChassis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualChassis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualChassis) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *VirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *VirtualChassis) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VirtualChassis) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VirtualChassis) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *VirtualChassis) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaster

`func (o *VirtualChassis) GetMaster() DeviceBayInstalledDevice`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VirtualChassis) GetMasterOk() (*DeviceBayInstalledDevice, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VirtualChassis) SetMaster(v DeviceBayInstalledDevice)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *VirtualChassis) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *VirtualChassis) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *VirtualChassis) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetTags

`func (o *VirtualChassis) GetTags() []TagSerializerField`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualChassis) GetTagsOk() (*[]TagSerializerField, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualChassis) SetTags(v []TagSerializerField)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualChassis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualChassis) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualChassis) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualChassis) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualChassis) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMemberCount

`func (o *VirtualChassis) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *VirtualChassis) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *VirtualChassis) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetDisplay

`func (o *VirtualChassis) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualChassis) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualChassis) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


