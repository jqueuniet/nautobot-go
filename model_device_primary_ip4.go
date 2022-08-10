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

// DevicePrimaryIp4 struct for DevicePrimaryIp4
type DevicePrimaryIp4 struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Family int32 `json:"family"`
	Address string `json:"address"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewDevicePrimaryIp4 instantiates a new DevicePrimaryIp4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePrimaryIp4(id string, url string, family int32, address string, display string) *DevicePrimaryIp4 {
	this := DevicePrimaryIp4{}
	this.Id = id
	this.Url = url
	this.Family = family
	this.Address = address
	this.Display = display
	return &this
}

// NewDevicePrimaryIp4WithDefaults instantiates a new DevicePrimaryIp4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePrimaryIp4WithDefaults() *DevicePrimaryIp4 {
	this := DevicePrimaryIp4{}
	return &this
}

// GetId returns the Id field value
func (o *DevicePrimaryIp4) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DevicePrimaryIp4) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DevicePrimaryIp4) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DevicePrimaryIp4) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DevicePrimaryIp4) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DevicePrimaryIp4) SetUrl(v string) {
	o.Url = v
}

// GetFamily returns the Family field value
func (o *DevicePrimaryIp4) GetFamily() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *DevicePrimaryIp4) GetFamilyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *DevicePrimaryIp4) SetFamily(v int32) {
	o.Family = v
}

// GetAddress returns the Address field value
func (o *DevicePrimaryIp4) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DevicePrimaryIp4) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DevicePrimaryIp4) SetAddress(v string) {
	o.Address = v
}

// GetDisplay returns the Display field value
func (o *DevicePrimaryIp4) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DevicePrimaryIp4) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DevicePrimaryIp4) SetDisplay(v string) {
	o.Display = v
}

func (o DevicePrimaryIp4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableDevicePrimaryIp4 struct {
	value *DevicePrimaryIp4
	isSet bool
}

func (v NullableDevicePrimaryIp4) Get() *DevicePrimaryIp4 {
	return v.value
}

func (v *NullableDevicePrimaryIp4) Set(val *DevicePrimaryIp4) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePrimaryIp4) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePrimaryIp4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePrimaryIp4(val *DevicePrimaryIp4) *NullableDevicePrimaryIp4 {
	return &NullableDevicePrimaryIp4{value: val, isSet: true}
}

func (v NullableDevicePrimaryIp4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePrimaryIp4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


