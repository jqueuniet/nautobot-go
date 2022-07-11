# PowerPanelRackGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**RackCount** | **int32** |  | [readonly] 
**Depth** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewPowerPanelRackGroup

`func NewPowerPanelRackGroup(id string, url string, name string, rackCount int32, depth int32, display string, ) *PowerPanelRackGroup`

NewPowerPanelRackGroup instantiates a new PowerPanelRackGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPanelRackGroupWithDefaults

`func NewPowerPanelRackGroupWithDefaults() *PowerPanelRackGroup`

NewPowerPanelRackGroupWithDefaults instantiates a new PowerPanelRackGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPanelRackGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPanelRackGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPanelRackGroup) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *PowerPanelRackGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPanelRackGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPanelRackGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *PowerPanelRackGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPanelRackGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPanelRackGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PowerPanelRackGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PowerPanelRackGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PowerPanelRackGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PowerPanelRackGroup) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRackCount

`func (o *PowerPanelRackGroup) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *PowerPanelRackGroup) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *PowerPanelRackGroup) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.


### GetDepth

`func (o *PowerPanelRackGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *PowerPanelRackGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *PowerPanelRackGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetDisplay

`func (o *PowerPanelRackGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPanelRackGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPanelRackGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


