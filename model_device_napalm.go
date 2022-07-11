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

// DeviceNAPALM struct for DeviceNAPALM
type DeviceNAPALM struct {
	Method map[string]interface{} `json:"method"`
}

// NewDeviceNAPALM instantiates a new DeviceNAPALM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceNAPALM(method map[string]interface{}) *DeviceNAPALM {
	this := DeviceNAPALM{}
	this.Method = method
	return &this
}

// NewDeviceNAPALMWithDefaults instantiates a new DeviceNAPALM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceNAPALMWithDefaults() *DeviceNAPALM {
	this := DeviceNAPALM{}
	return &this
}

// GetMethod returns the Method field value
func (o *DeviceNAPALM) GetMethod() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *DeviceNAPALM) GetMethodOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method, true
}

// SetMethod sets field value
func (o *DeviceNAPALM) SetMethod(v map[string]interface{}) {
	o.Method = v
}

func (o DeviceNAPALM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceNAPALM struct {
	value *DeviceNAPALM
	isSet bool
}

func (v NullableDeviceNAPALM) Get() *DeviceNAPALM {
	return v.value
}

func (v *NullableDeviceNAPALM) Set(val *DeviceNAPALM) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceNAPALM) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceNAPALM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceNAPALM(val *DeviceNAPALM) *NullableDeviceNAPALM {
	return &NullableDeviceNAPALM{value: val, isSet: true}
}

func (v NullableDeviceNAPALM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceNAPALM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


