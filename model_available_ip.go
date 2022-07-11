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

// AvailableIP Representation of an IP address which does not exist in the database.
type AvailableIP struct {
	Family int32 `json:"family"`
	Address string `json:"address"`
	Vrf AvailableIPVrf `json:"vrf"`
}

// NewAvailableIP instantiates a new AvailableIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableIP(family int32, address string, vrf AvailableIPVrf) *AvailableIP {
	this := AvailableIP{}
	this.Family = family
	this.Address = address
	this.Vrf = vrf
	return &this
}

// NewAvailableIPWithDefaults instantiates a new AvailableIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableIPWithDefaults() *AvailableIP {
	this := AvailableIP{}
	return &this
}

// GetFamily returns the Family field value
func (o *AvailableIP) GetFamily() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *AvailableIP) GetFamilyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *AvailableIP) SetFamily(v int32) {
	o.Family = v
}

// GetAddress returns the Address field value
func (o *AvailableIP) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AvailableIP) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AvailableIP) SetAddress(v string) {
	o.Address = v
}

// GetVrf returns the Vrf field value
func (o *AvailableIP) GetVrf() AvailableIPVrf {
	if o == nil {
		var ret AvailableIPVrf
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *AvailableIP) GetVrfOk() (*AvailableIPVrf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *AvailableIP) SetVrf(v AvailableIPVrf) {
	o.Vrf = v
}

func (o AvailableIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableIP struct {
	value *AvailableIP
	isSet bool
}

func (v NullableAvailableIP) Get() *AvailableIP {
	return v.value
}

func (v *NullableAvailableIP) Set(val *AvailableIP) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableIP) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableIP(val *AvailableIP) *NullableAvailableIP {
	return &NullableAvailableIP{value: val, isSet: true}
}

func (v NullableAvailableIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


