/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GraphqlCreate200Response struct for GraphqlCreate200Response
type GraphqlCreate200Response struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewGraphqlCreate200Response instantiates a new GraphqlCreate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphqlCreate200Response() *GraphqlCreate200Response {
	this := GraphqlCreate200Response{}
	return &this
}

// NewGraphqlCreate200ResponseWithDefaults instantiates a new GraphqlCreate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphqlCreate200ResponseWithDefaults() *GraphqlCreate200Response {
	this := GraphqlCreate200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GraphqlCreate200Response) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphqlCreate200Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GraphqlCreate200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *GraphqlCreate200Response) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o GraphqlCreate200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGraphqlCreate200Response struct {
	value *GraphqlCreate200Response
	isSet bool
}

func (v NullableGraphqlCreate200Response) Get() *GraphqlCreate200Response {
	return v.value
}

func (v *NullableGraphqlCreate200Response) Set(val *GraphqlCreate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphqlCreate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphqlCreate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphqlCreate200Response(val *GraphqlCreate200Response) *NullableGraphqlCreate200Response {
	return &NullableGraphqlCreate200Response{value: val, isSet: true}
}

func (v NullableGraphqlCreate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphqlCreate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


