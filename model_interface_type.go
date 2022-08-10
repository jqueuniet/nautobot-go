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

// InterfaceType struct for InterfaceType
type InterfaceType struct {
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewInterfaceType instantiates a new InterfaceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceType() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// NewInterfaceTypeWithDefaults instantiates a new InterfaceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTypeWithDefaults() *InterfaceType {
	this := InterfaceType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceType) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceType) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfaceType) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceType) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceType) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceType) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceType) SetLabel(v string) {
	o.Label = &v
}

func (o InterfaceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceType struct {
	value *InterfaceType
	isSet bool
}

func (v NullableInterfaceType) Get() *InterfaceType {
	return v.value
}

func (v *NullableInterfaceType) Set(val *InterfaceType) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceType) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceType(val *InterfaceType) *NullableInterfaceType {
	return &NullableInterfaceType{value: val, isSet: true}
}

func (v NullableInterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


