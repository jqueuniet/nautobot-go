# NestedRackRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**RackCount** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewNestedRackRole

`func NewNestedRackRole(id string, url string, name string, rackCount int32, display string, ) *NestedRackRole`

NewNestedRackRole instantiates a new NestedRackRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRackRoleWithDefaults

`func NewNestedRackRoleWithDefaults() *NestedRackRole`

NewNestedRackRoleWithDefaults instantiates a new NestedRackRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedRackRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedRackRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedRackRole) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedRackRole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedRackRole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedRackRole) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedRackRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedRackRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedRackRole) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedRackRole) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedRackRole) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedRackRole) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedRackRole) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetRackCount

`func (o *NestedRackRole) GetRackCount() int32`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *NestedRackRole) GetRackCountOk() (*int32, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *NestedRackRole) SetRackCount(v int32)`

SetRackCount sets RackCount field to given value.


### GetDisplay

`func (o *NestedRackRole) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedRackRole) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedRackRole) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


