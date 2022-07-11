# TenantGroupParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**TenantCount** | **int32** |  | [readonly] 
**Depth** | **int32** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewTenantGroupParent

`func NewTenantGroupParent(id string, url string, name string, tenantCount int32, depth int32, display string, ) *TenantGroupParent`

NewTenantGroupParent instantiates a new TenantGroupParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantGroupParentWithDefaults

`func NewTenantGroupParentWithDefaults() *TenantGroupParent`

NewTenantGroupParentWithDefaults instantiates a new TenantGroupParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantGroupParent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantGroupParent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantGroupParent) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *TenantGroupParent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TenantGroupParent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TenantGroupParent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *TenantGroupParent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantGroupParent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantGroupParent) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *TenantGroupParent) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TenantGroupParent) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TenantGroupParent) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TenantGroupParent) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTenantCount

`func (o *TenantGroupParent) GetTenantCount() int32`

GetTenantCount returns the TenantCount field if non-nil, zero value otherwise.

### GetTenantCountOk

`func (o *TenantGroupParent) GetTenantCountOk() (*int32, bool)`

GetTenantCountOk returns a tuple with the TenantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCount

`func (o *TenantGroupParent) SetTenantCount(v int32)`

SetTenantCount sets TenantCount field to given value.


### GetDepth

`func (o *TenantGroupParent) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *TenantGroupParent) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *TenantGroupParent) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetDisplay

`func (o *TenantGroupParent) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TenantGroupParent) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TenantGroupParent) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


