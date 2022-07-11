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

// NestedPowerPort Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedPowerPort struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Device DeviceParentDevice `json:"device"`
	Name string `json:"name"`
	Cable NullableString `json:"cable,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedPowerPort instantiates a new NestedPowerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPowerPort(id string, url string, device DeviceParentDevice, name string, display string) *NestedPowerPort {
	this := NestedPowerPort{}
	this.Id = id
	this.Url = url
	this.Device = device
	this.Name = name
	this.Display = display
	return &this
}

// NewNestedPowerPortWithDefaults instantiates a new NestedPowerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPowerPortWithDefaults() *NestedPowerPort {
	this := NestedPowerPort{}
	return &this
}

// GetId returns the Id field value
func (o *NestedPowerPort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedPowerPort) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedPowerPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedPowerPort) SetUrl(v string) {
	o.Url = v
}

// GetDevice returns the Device field value
func (o *NestedPowerPort) GetDevice() DeviceParentDevice {
	if o == nil {
		var ret DeviceParentDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPort) GetDeviceOk() (*DeviceParentDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *NestedPowerPort) SetDevice(v DeviceParentDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *NestedPowerPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedPowerPort) SetName(v string) {
	o.Name = v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedPowerPort) GetCable() string {
	if o == nil || o.Cable.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedPowerPort) GetCableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *NestedPowerPort) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableString and assigns it to the Cable field.
func (o *NestedPowerPort) SetCable(v string) {
	o.Cable.Set(&v)
}
// SetCableNil sets the value for Cable to be an explicit nil
func (o *NestedPowerPort) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *NestedPowerPort) UnsetCable() {
	o.Cable.Unset()
}

// GetDisplay returns the Display field value
func (o *NestedPowerPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedPowerPort) SetDisplay(v string) {
	o.Display = v
}

func (o NestedPowerPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["device"] = o.Device
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedPowerPort struct {
	value *NestedPowerPort
	isSet bool
}

func (v NullableNestedPowerPort) Get() *NestedPowerPort {
	return v.value
}

func (v *NullableNestedPowerPort) Set(val *NestedPowerPort) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPowerPort) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPowerPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPowerPort(val *NestedPowerPort) *NullableNestedPowerPort {
	return &NullableNestedPowerPort{value: val, isSet: true}
}

func (v NullableNestedPowerPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPowerPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


