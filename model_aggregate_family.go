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

// AggregateFamily struct for AggregateFamily
type AggregateFamily struct {
	Value *int32 `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewAggregateFamily instantiates a new AggregateFamily object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateFamily() *AggregateFamily {
	this := AggregateFamily{}
	return &this
}

// NewAggregateFamilyWithDefaults instantiates a new AggregateFamily object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateFamilyWithDefaults() *AggregateFamily {
	this := AggregateFamily{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AggregateFamily) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateFamily) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AggregateFamily) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AggregateFamily) SetValue(v int32) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AggregateFamily) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateFamily) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AggregateFamily) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AggregateFamily) SetLabel(v string) {
	o.Label = &v
}

func (o AggregateFamily) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableAggregateFamily struct {
	value *AggregateFamily
	isSet bool
}

func (v NullableAggregateFamily) Get() *AggregateFamily {
	return v.value
}

func (v *NullableAggregateFamily) Set(val *AggregateFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateFamily(val *AggregateFamily) *NullableAggregateFamily {
	return &NullableAggregateFamily{value: val, isSet: true}
}

func (v NullableAggregateFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


