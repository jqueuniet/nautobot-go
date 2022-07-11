/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GraphQLAPI struct for GraphQLAPI
type GraphQLAPI struct {
	// GraphQL query
	Query string `json:"query"`
	// Variables in JSON Format
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// NewGraphQLAPI instantiates a new GraphQLAPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLAPI(query string) *GraphQLAPI {
	this := GraphQLAPI{}
	this.Query = query
	return &this
}

// NewGraphQLAPIWithDefaults instantiates a new GraphQLAPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLAPIWithDefaults() *GraphQLAPI {
	this := GraphQLAPI{}
	return &this
}

// GetQuery returns the Query field value
func (o *GraphQLAPI) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GraphQLAPI) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GraphQLAPI) SetQuery(v string) {
	o.Query = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *GraphQLAPI) GetVariables() map[string]interface{} {
	if o == nil || o.Variables == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphQLAPI) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *GraphQLAPI) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *GraphQLAPI) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

func (o GraphQLAPI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableGraphQLAPI struct {
	value *GraphQLAPI
	isSet bool
}

func (v NullableGraphQLAPI) Get() *GraphQLAPI {
	return v.value
}

func (v *NullableGraphQLAPI) Set(val *GraphQLAPI) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLAPI) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLAPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLAPI(val *GraphQLAPI) *NullableGraphQLAPI {
	return &NullableGraphQLAPI{value: val, isSet: true}
}

func (v NullableGraphQLAPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLAPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


