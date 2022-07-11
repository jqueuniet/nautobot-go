# GraphQLAPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | GraphQL query | 
**Variables** | Pointer to **map[string]interface{}** | Variables in JSON Format | [optional] 

## Methods

### NewGraphQLAPI

`func NewGraphQLAPI(query string, ) *GraphQLAPI`

NewGraphQLAPI instantiates a new GraphQLAPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLAPIWithDefaults

`func NewGraphQLAPIWithDefaults() *GraphQLAPI`

NewGraphQLAPIWithDefaults instantiates a new GraphQLAPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *GraphQLAPI) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLAPI) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLAPI) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetVariables

`func (o *GraphQLAPI) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLAPI) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLAPI) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLAPI) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


