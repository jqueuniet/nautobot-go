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

// NestedDevice Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedDevice struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name NullableString `json:"name,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedDevice instantiates a new NestedDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDevice(id string, url string, display string) *NestedDevice {
	this := NestedDevice{}
	this.Id = id
	this.Url = url
	this.Display = display
	return &this
}

// NewNestedDeviceWithDefaults instantiates a new NestedDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceWithDefaults() *NestedDevice {
	this := NestedDevice{}
	return &this
}

// GetId returns the Id field value
func (o *NestedDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedDevice) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedDevice) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedDevice) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedDevice) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedDevice) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NestedDevice) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NestedDevice) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NestedDevice) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NestedDevice) UnsetName() {
	o.Name.Unset()
}

// GetDisplay returns the Display field value
func (o *NestedDevice) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedDevice) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedDevice) SetDisplay(v string) {
	o.Display = v
}

func (o NestedDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedDevice struct {
	value *NestedDevice
	isSet bool
}

func (v NullableNestedDevice) Get() *NestedDevice {
	return v.value
}

func (v *NullableNestedDevice) Set(val *NestedDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDevice(val *NestedDevice) *NullableNestedDevice {
	return &NullableNestedDevice{value: val, isSet: true}
}

func (v NullableNestedDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


