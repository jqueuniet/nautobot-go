# PatchedConfigContextSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**DataSchema** | Pointer to **map[string]interface{}** | A JSON Schema document which is used to validate a config context object. | [optional] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedConfigContextSchema

`func NewPatchedConfigContextSchema() *PatchedConfigContextSchema`

NewPatchedConfigContextSchema instantiates a new PatchedConfigContextSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConfigContextSchemaWithDefaults

`func NewPatchedConfigContextSchemaWithDefaults() *PatchedConfigContextSchema`

NewPatchedConfigContextSchemaWithDefaults instantiates a new PatchedConfigContextSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedConfigContextSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedConfigContextSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedConfigContextSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedConfigContextSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedConfigContextSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedConfigContextSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedConfigContextSchema) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedConfigContextSchema) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedConfigContextSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConfigContextSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConfigContextSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConfigContextSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedConfigContextSchema) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedConfigContextSchema) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedConfigContextSchema) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedConfigContextSchema) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetOwnerContentType

`func (o *PatchedConfigContextSchema) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *PatchedConfigContextSchema) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *PatchedConfigContextSchema) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *PatchedConfigContextSchema) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *PatchedConfigContextSchema) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *PatchedConfigContextSchema) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetOwnerObjectId

`func (o *PatchedConfigContextSchema) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *PatchedConfigContextSchema) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *PatchedConfigContextSchema) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *PatchedConfigContextSchema) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *PatchedConfigContextSchema) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *PatchedConfigContextSchema) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetOwner

`func (o *PatchedConfigContextSchema) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedConfigContextSchema) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedConfigContextSchema) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedConfigContextSchema) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *PatchedConfigContextSchema) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *PatchedConfigContextSchema) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetDescription

`func (o *PatchedConfigContextSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedConfigContextSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedConfigContextSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedConfigContextSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDataSchema

`func (o *PatchedConfigContextSchema) GetDataSchema() map[string]interface{}`

GetDataSchema returns the DataSchema field if non-nil, zero value otherwise.

### GetDataSchemaOk

`func (o *PatchedConfigContextSchema) GetDataSchemaOk() (*map[string]interface{}, bool)`

GetDataSchemaOk returns a tuple with the DataSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSchema

`func (o *PatchedConfigContextSchema) SetDataSchema(v map[string]interface{})`

SetDataSchema sets DataSchema field to given value.

### HasDataSchema

`func (o *PatchedConfigContextSchema) HasDataSchema() bool`

HasDataSchema returns a boolean if a field has been set.

### GetCreated

`func (o *PatchedConfigContextSchema) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PatchedConfigContextSchema) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PatchedConfigContextSchema) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PatchedConfigContextSchema) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PatchedConfigContextSchema) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PatchedConfigContextSchema) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PatchedConfigContextSchema) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PatchedConfigContextSchema) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDisplay

`func (o *PatchedConfigContextSchema) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedConfigContextSchema) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedConfigContextSchema) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedConfigContextSchema) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


