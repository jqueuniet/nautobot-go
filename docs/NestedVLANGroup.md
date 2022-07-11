# NestedVLANGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**VlanCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewNestedVLANGroup

`func NewNestedVLANGroup(id string, url string, name string, vlanCount int32, display string, ) *NestedVLANGroup`

NewNestedVLANGroup instantiates a new NestedVLANGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVLANGroupWithDefaults

`func NewNestedVLANGroupWithDefaults() *NestedVLANGroup`

NewNestedVLANGroupWithDefaults instantiates a new NestedVLANGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedVLANGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedVLANGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedVLANGroup) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedVLANGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedVLANGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedVLANGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedVLANGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVLANGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVLANGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedVLANGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedVLANGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedVLANGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedVLANGroup) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetVlanCount

`func (o *NestedVLANGroup) GetVlanCount() int32`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *NestedVLANGroup) GetVlanCountOk() (*int32, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *NestedVLANGroup) SetVlanCount(v int32)`

SetVlanCount sets VlanCount field to given value.


### GetDisplay

`func (o *NestedVLANGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedVLANGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedVLANGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


