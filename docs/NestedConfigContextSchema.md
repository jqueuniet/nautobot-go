# NestedConfigContextSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewNestedConfigContextSchema

`func NewNestedConfigContextSchema(id string, url string, name string, display string, ) *NestedConfigContextSchema`

NewNestedConfigContextSchema instantiates a new NestedConfigContextSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedConfigContextSchemaWithDefaults

`func NewNestedConfigContextSchemaWithDefaults() *NestedConfigContextSchema`

NewNestedConfigContextSchemaWithDefaults instantiates a new NestedConfigContextSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedConfigContextSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedConfigContextSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedConfigContextSchema) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedConfigContextSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedConfigContextSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedConfigContextSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedConfigContextSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedConfigContextSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedConfigContextSchema) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedConfigContextSchema) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedConfigContextSchema) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedConfigContextSchema) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedConfigContextSchema) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedConfigContextSchema) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedConfigContextSchema) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedConfigContextSchema) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


