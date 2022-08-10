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

// NestedDeviceType Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedDeviceType struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Manufacturer ConfigContextDeviceTypesInnerManufacturer `json:"manufacturer"`
	Model string `json:"model"`
	Slug string `json:"slug"`
	DeviceCount int32 `json:"device_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedDeviceType instantiates a new NestedDeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeviceType(id string, url string, manufacturer ConfigContextDeviceTypesInnerManufacturer, model string, slug string, deviceCount int32, display string) *NestedDeviceType {
	this := NestedDeviceType{}
	this.Id = id
	this.Url = url
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	this.DeviceCount = deviceCount
	this.Display = display
	return &this
}

// NewNestedDeviceTypeWithDefaults instantiates a new NestedDeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceTypeWithDefaults() *NestedDeviceType {
	this := NestedDeviceType{}
	return &this
}

// GetId returns the Id field value
func (o *NestedDeviceType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedDeviceType) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedDeviceType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedDeviceType) SetUrl(v string) {
	o.Url = v
}

// GetManufacturer returns the Manufacturer field value
func (o *NestedDeviceType) GetManufacturer() ConfigContextDeviceTypesInnerManufacturer {
	if o == nil {
		var ret ConfigContextDeviceTypesInnerManufacturer
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetManufacturerOk() (*ConfigContextDeviceTypesInnerManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *NestedDeviceType) SetManufacturer(v ConfigContextDeviceTypesInnerManufacturer) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *NestedDeviceType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *NestedDeviceType) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *NestedDeviceType) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedDeviceType) SetSlug(v string) {
	o.Slug = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *NestedDeviceType) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *NestedDeviceType) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetDisplay returns the Display field value
func (o *NestedDeviceType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedDeviceType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedDeviceType) SetDisplay(v string) {
	o.Display = v
}

func (o NestedDeviceType) MarshalJSON() ([]byte, error) {
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

type NullableNestedDeviceType struct {
	value *NestedDeviceType
	isSet bool
}

func (v NullableNestedDeviceType) Get() *NestedDeviceType {
	return v.value
}

func (v *NullableNestedDeviceType) Set(val *NestedDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeviceType(val *NestedDeviceType) *NullableNestedDeviceType {
	return &NullableNestedDeviceType{value: val, isSet: true}
}

func (v NullableNestedDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


