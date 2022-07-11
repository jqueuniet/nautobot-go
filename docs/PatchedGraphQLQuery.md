# PatchedGraphQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Url** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Display** | Pointer to **string** | Human friendly display value | [optional] [readonly] 

## Methods

### NewPatchedGraphQLQuery

`func NewPatchedGraphQLQuery() *PatchedGraphQLQuery`

NewPatchedGraphQLQuery instantiates a new PatchedGraphQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGraphQLQueryWithDefaults

`func NewPatchedGraphQLQueryWithDefaults() *PatchedGraphQLQuery`

NewPatchedGraphQLQueryWithDefaults instantiates a new PatchedGraphQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedGraphQLQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedGraphQLQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedGraphQLQuery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedGraphQLQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedGraphQLQuery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedGraphQLQuery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedGraphQLQuery) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedGraphQLQuery) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *PatchedGraphQLQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGraphQLQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGraphQLQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGraphQLQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedGraphQLQuery) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedGraphQLQuery) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedGraphQLQuery) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedGraphQLQuery) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetQuery

`func (o *PatchedGraphQLQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PatchedGraphQLQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PatchedGraphQLQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PatchedGraphQLQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetVariables

`func (o *PatchedGraphQLQuery) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PatchedGraphQLQuery) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PatchedGraphQLQuery) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PatchedGraphQLQuery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *PatchedGraphQLQuery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *PatchedGraphQLQuery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDisplay

`func (o *PatchedGraphQLQuery) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PatchedGraphQLQuery) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PatchedGraphQLQuery) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PatchedGraphQLQuery) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


