# NestedRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** | Internal relationship name | 
**Slug** | Pointer to **string** |  | [optional] 
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewNestedRelationship

`func NewNestedRelationship(id string, url string, name string, display string, ) *NestedRelationship`

NewNestedRelationship instantiates a new NestedRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedRelationshipWithDefaults

`func NewNestedRelationshipWithDefaults() *NestedRelationship`

NewNestedRelationshipWithDefaults instantiates a new NestedRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedRelationship) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedRelationship) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedRelationship) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedRelationship) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *NestedRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedRelationship) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *NestedRelationship) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *NestedRelationship) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *NestedRelationship) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *NestedRelationship) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplay

`func (o *NestedRelationship) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedRelationship) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedRelationship) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


