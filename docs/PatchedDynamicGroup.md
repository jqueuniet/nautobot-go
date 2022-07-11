# PatchedDynamicGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Dynamic Group name | [optional] 
**Slug** | Pointer to **string** | Unique slug | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **map[string]interface{}** | A JSON-encoded dictionary of filter parameters for group membership | [optional] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedDynamicGroup

`func NewPatchedDynamicGroup() *PatchedDynamicGroup`

NewPatchedDynamicGroup instantiates a new PatchedDynamicGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDynamicGroupWithDefaults

`func NewPatchedDynamicGroupWithDefaults() *PatchedDynamicGroup`

NewPatchedDynamicGroupWithDefaults instantiates a new PatchedDynamicGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDynamicGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDynamicGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDynamicGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDynamicGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedDynamicGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedDynamicGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedDynamicGroup) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedDynamicGroup) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedDynamicGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDynamicGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDynamicGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDynamicGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedDynamicGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedDynamicGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedDynamicGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedDynamicGroup) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedDynamicGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDynamicGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDynamicGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDynamicGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedDynamicGroup) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedDynamicGroup) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedDynamicGroup) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedDynamicGroup) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFilter

`func (o *PatchedDynamicGroup) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PatchedDynamicGroup) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PatchedDynamicGroup) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PatchedDynamicGroup) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedDynamicGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedDynamicGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedDynamicGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedDynamicGroup) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


