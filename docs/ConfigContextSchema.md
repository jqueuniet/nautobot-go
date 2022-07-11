# ConfigContextSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewConfigContextSchema

`func NewConfigContextSchema(id string, url string, name string, display string, ) *ConfigContextSchema`

NewConfigContextSchema instantiates a new ConfigContextSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextSchemaWithDefaults

`func NewConfigContextSchemaWithDefaults() *ConfigContextSchema`

NewConfigContextSchemaWithDefaults instantiates a new ConfigContextSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContextSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContextSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContextSchema) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConfigContextSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContextSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContextSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *ConfigContextSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContextSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContextSchema) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ConfigContextSchema) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ConfigContextSchema) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ConfigContextSchema) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ConfigContextSchema) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplay

`func (o *ConfigContextSchema) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContextSchema) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContextSchema) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


