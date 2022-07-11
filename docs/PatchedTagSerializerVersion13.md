# PatchedTagSerializerVersion13

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaggedItems** | Pointer to **int32** |  | [optional] [readonly] 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedTagSerializerVersion13

`func NewPatchedTagSerializerVersion13() *PatchedTagSerializerVersion13`

NewPatchedTagSerializerVersion13 instantiates a new PatchedTagSerializerVersion13 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTagSerializerVersion13WithDefaults

`func NewPatchedTagSerializerVersion13WithDefaults() *PatchedTagSerializerVersion13`

NewPatchedTagSerializerVersion13WithDefaults instantiates a new PatchedTagSerializerVersion13 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTagSerializerVersion13) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTagSerializerVersion13) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTagSerializerVersion13) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTagSerializerVersion13) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedTagSerializerVersion13) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedTagSerializerVersion13) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedTagSerializerVersion13) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedTagSerializerVersion13) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedTagSerializerVersion13) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTagSerializerVersion13) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTagSerializerVersion13) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTagSerializerVersion13) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedTagSerializerVersion13) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedTagSerializerVersion13) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedTagSerializerVersion13) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedTagSerializerVersion13) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetColor

`func (o *PatchedTagSerializerVersion13) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedTagSerializerVersion13) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedTagSerializerVersion13) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedTagSerializerVersion13) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTagSerializerVersion13) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTagSerializerVersion13) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTagSerializerVersion13) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTagSerializerVersion13) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaggedItems

`func (o *PatchedTagSerializerVersion13) GetTaggedItems() int32`

GetTaggedItems returns the TaggedItems field if non-nil, zero value otherwise.

### GetTaggedItemsOk

`func (o *PatchedTagSerializerVersion13) GetTaggedItemsOk() (*int32, bool)`

GetTaggedItemsOk returns a tuple with the TaggedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedItems

`func (o *PatchedTagSerializerVersion13) SetTaggedItems(v int32)`

SetTaggedItems sets TaggedItems field to given value.

### HasTaggedItems

`func (o *PatchedTagSerializerVersion13) HasTaggedItems() bool`

HasTaggedItems returns a boolean if a field has been set.

### GetContentTypes

`func (o *PatchedTagSerializerVersion13) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedTagSerializerVersion13) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedTagSerializerVersion13) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedTagSerializerVersion13) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedTagSerializerVersion13) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedTagSerializerVersion13) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedTagSerializerVersion13) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedTagSerializerVersion13) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedTagSerializerVersion13) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedTagSerializerVersion13) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedTagSerializerVersion13) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedTagSerializerVersion13) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedTagSerializerVersion13) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedTagSerializerVersion13) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedTagSerializerVersion13) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedTagSerializerVersion13) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedTagSerializerVersion13) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedTagSerializerVersion13) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedTagSerializerVersion13) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedTagSerializerVersion13) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


