# TagSerializerVersion13

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaggedItems** | **int32** |  | [readonly] 
**ContentTypes** | **[]string** |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **string** |  | [readonly] 
**LastUpdated** | **time.Time** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewTagSerializerVersion13

`func NewTagSerializerVersion13(id string, url string, name string, slug string, taggedItems int32, contentTypes []string, created string, lastUpdated time.Time, display string, ) *TagSerializerVersion13`

NewTagSerializerVersion13 instantiates a new TagSerializerVersion13 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSerializerVersion13WithDefaults

`func NewTagSerializerVersion13WithDefaults() *TagSerializerVersion13`

NewTagSerializerVersion13WithDefaults instantiates a new TagSerializerVersion13 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagSerializerVersion13) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagSerializerVersion13) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagSerializerVersion13) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *TagSerializerVersion13) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TagSerializerVersion13) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TagSerializerVersion13) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *TagSerializerVersion13) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagSerializerVersion13) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagSerializerVersion13) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *TagSerializerVersion13) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TagSerializerVersion13) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TagSerializerVersion13) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *TagSerializerVersion13) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TagSerializerVersion13) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TagSerializerVersion13) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TagSerializerVersion13) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *TagSerializerVersion13) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagSerializerVersion13) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagSerializerVersion13) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagSerializerVersion13) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaggedItems

`func (o *TagSerializerVersion13) GetTaggedItems() int32`

GetTaggedItems returns the TaggedItems field if non-nil, zero value otherwise.

### GetTaggedItemsOk

`func (o *TagSerializerVersion13) GetTaggedItemsOk() (*int32, bool)`

GetTaggedItemsOk returns a tuple with the TaggedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedItems

`func (o *TagSerializerVersion13) SetTaggedItems(v int32)`

SetTaggedItems sets TaggedItems field to given value.


### GetContentTypes

`func (o *TagSerializerVersion13) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *TagSerializerVersion13) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *TagSerializerVersion13) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetCustomFields

`func (o *TagSerializerVersion13) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TagSerializerVersion13) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TagSerializerVersion13) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TagSerializerVersion13) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *TagSerializerVersion13) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TagSerializerVersion13) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TagSerializerVersion13) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *TagSerializerVersion13) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TagSerializerVersion13) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TagSerializerVersion13) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetDisplay

`func (o *TagSerializerVersion13) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TagSerializerVersion13) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TagSerializerVersion13) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


