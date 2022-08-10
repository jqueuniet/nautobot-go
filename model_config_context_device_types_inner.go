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

// ConfigContextDeviceTypesInner struct for ConfigContextDeviceTypesInner
type ConfigContextDeviceTypesInner struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Manufacturer ConfigContextDeviceTypesInnerManufacturer `json:"manufacturer"`
	Model string `json:"model"`
	Slug string `json:"slug"`
	DeviceCount int32 `json:"device_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewConfigContextDeviceTypesInner instantiates a new ConfigContextDeviceTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigContextDeviceTypesInner(id string, url string, manufacturer ConfigContextDeviceTypesInnerManufacturer, model string, slug string, deviceCount int32, display string) *ConfigContextDeviceTypesInner {
	this := ConfigContextDeviceTypesInner{}
	this.Id = id
	this.Url = url
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	this.DeviceCount = deviceCount
	this.Display = display
	return &this
}

// NewConfigContextDeviceTypesInnerWithDefaults instantiates a new ConfigContextDeviceTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigContextDeviceTypesInnerWithDefaults() *ConfigContextDeviceTypesInner {
	this := ConfigContextDeviceTypesInner{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigContextDeviceTypesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigContextDeviceTypesInner) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConfigContextDeviceTypesInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConfigContextDeviceTypesInner) SetUrl(v string) {
	o.Url = v
}

// GetManufacturer returns the Manufacturer field value
func (o *ConfigContextDeviceTypesInner) GetManufacturer() ConfigContextDeviceTypesInnerManufacturer {
	if o == nil {
		var ret ConfigContextDeviceTypesInnerManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetManufacturerOk() (*ConfigContextDeviceTypesInnerManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *ConfigContextDeviceTypesInner) SetManufacturer(v ConfigContextDeviceTypesInnerManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *ConfigContextDeviceTypesInner) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ConfigContextDeviceTypesInner) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *ConfigContextDeviceTypesInner) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *ConfigContextDeviceTypesInner) SetSlug(v string) {
	o.Slug = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *ConfigContextDeviceTypesInner) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *ConfigContextDeviceTypesInner) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetDisplay returns the Display field value
func (o *ConfigContextDeviceTypesInner) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConfigContextDeviceTypesInner) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConfigContextDeviceTypesInner) SetDisplay(v string) {
	o.Display = v
}

func (o ConfigContextDeviceTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["device_count"] = o.DeviceCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableConfigContextDeviceTypesInner struct {
	value *ConfigContextDeviceTypesInner
	isSet bool
}

func (v NullableConfigContextDeviceTypesInner) Get() *ConfigContextDeviceTypesInner {
	return v.value
}

func (v *NullableConfigContextDeviceTypesInner) Set(val *ConfigContextDeviceTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigContextDeviceTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigContextDeviceTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigContextDeviceTypesInner(val *ConfigContextDeviceTypesInner) *NullableConfigContextDeviceTypesInner {
	return &NullableConfigContextDeviceTypesInner{value: val, isSet: true}
}

func (v NullableConfigContextDeviceTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigContextDeviceTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


