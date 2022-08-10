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

// AvailablePrefixRequest struct for AvailablePrefixRequest
type AvailablePrefixRequest struct {
	PrefixLength int32 `json:"prefix_length"`
	Status AvailablePrefixRequestStatusEnum `json:"status"`
}

// NewAvailablePrefixRequest instantiates a new AvailablePrefixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePrefixRequest(prefixLength int32, status AvailablePrefixRequestStatusEnum) *AvailablePrefixRequest {
	this := AvailablePrefixRequest{}
	this.PrefixLength = prefixLength
	this.Status = status
	return &this
}

// NewAvailablePrefixRequestWithDefaults instantiates a new AvailablePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePrefixRequestWithDefaults() *AvailablePrefixRequest {
	this := AvailablePrefixRequest{}
	return &this
}

// GetPrefixLength returns the PrefixLength field value
func (o *AvailablePrefixRequest) GetPrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefixRequest) GetPrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *AvailablePrefixRequest) SetPrefixLength(v int32) {
	o.PrefixLength = v
}

// GetStatus returns the Status field value
func (o *AvailablePrefixRequest) GetStatus() AvailablePrefixRequestStatusEnum {
	if o == nil {
		var ret AvailablePrefixRequestStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefixRequest) GetStatusOk() (*AvailablePrefixRequestStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AvailablePrefixRequest) SetStatus(v AvailablePrefixRequestStatusEnum) {
	o.Status = v
}

func (o AvailablePrefixRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prefix_length"] = o.PrefixLength
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAvailablePrefixRequest struct {
	value *AvailablePrefixRequest
	isSet bool
}

func (v NullableAvailablePrefixRequest) Get() *AvailablePrefixRequest {
	return v.value
}

func (v *NullableAvailablePrefixRequest) Set(val *AvailablePrefixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePrefixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePrefixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePrefixRequest(val *AvailablePrefixRequest) *NullableAvailablePrefixRequest {
	return &NullableAvailablePrefixRequest{value: val, isSet: true}
}

func (v NullableAvailablePrefixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePrefixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

