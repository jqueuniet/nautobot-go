# GraphQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | Pointer to **string** |  | [optional] 
**Query** | **string** |  | 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Display** | **string** | Human friendly display value | [readonly] 

## Methods

### NewGraphQLQuery

`func NewGraphQLQuery(id string, url string, name string, query string, display string, ) *GraphQLQuery`

NewGraphQLQuery instantiates a new GraphQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLQueryWithDefaults

`func NewGraphQLQueryWithDefaults() *GraphQLQuery`

NewGraphQLQueryWithDefaults instantiates a new GraphQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphQLQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphQLQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphQLQuery) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *GraphQLQuery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GraphQLQuery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GraphQLQuery) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *GraphQLQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphQLQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphQLQuery) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *GraphQLQuery) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GraphQLQuery) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GraphQLQuery) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GraphQLQuery) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetQuery

`func (o *GraphQLQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetVariables

`func (o *GraphQLQuery) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLQuery) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLQuery) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLQuery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *GraphQLQuery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *GraphQLQuery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDisplay

`func (o *GraphQLQuery) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *GraphQLQuery) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *GraphQLQuery) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


