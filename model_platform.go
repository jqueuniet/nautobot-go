/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Platform Extends ModelSerializer to render any CustomFields and their values associated with an object.
type Platform struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	Manufacturer NullableInventoryItemManufacturer `json:"manufacturer,omitempty"`
	// The name of the NAPALM driver to use when interacting with devices
	NapalmDriver *string `json:"napalm_driver,omitempty"`
	// Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs map[string]interface{} `json:"napalm_args,omitempty"`
	Description *string `json:"description,omitempty"`
	DeviceCount int32 `json:"device_count"`
	VirtualmachineCount int32 `json:"virtualmachine_count"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewPlatform instantiates a new Platform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatform(id string, url string, name string, deviceCount int32, virtualmachineCount int32, created string, lastUpdated time.Time, display string) *Platform {
	this := Platform{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.DeviceCount = deviceCount
	this.VirtualmachineCount = virtualmachineCount
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Display = display
	return &this
}

// NewPlatformWithDefaults instantiates a new Platform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformWithDefaults() *Platform {
	this := Platform{}
	return &this
}

// GetId returns the Id field value
func (o *Platform) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Platform) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Platform) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Platform) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Platform) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Platform) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *Platform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Platform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Platform) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Platform) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Platform) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Platform) SetSlug(v string) {
	o.Slug = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Platform) GetManufacturer() InventoryItemManufacturer {
	if o == nil || o.Manufacturer.Get() == nil {
		var ret InventoryItemManufacturer
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Platform) GetManufacturerOk() (*InventoryItemManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *Platform) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableInventoryItemManufacturer and assigns it to the Manufacturer field.
func (o *Platform) SetManufacturer(v InventoryItemManufacturer) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *Platform) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *Platform) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetNapalmDriver returns the NapalmDriver field value if set, zero value otherwise.
func (o *Platform) GetNapalmDriver() string {
	if o == nil || o.NapalmDriver == nil {
		var ret string
		return ret
	}
	return *o.NapalmDriver
}

// GetNapalmDriverOk returns a tuple with the NapalmDriver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetNapalmDriverOk() (*string, bool) {
	if o == nil || o.NapalmDriver == nil {
		return nil, false
	}
	return o.NapalmDriver, true
}

// HasNapalmDriver returns a boolean if a field has been set.
func (o *Platform) HasNapalmDriver() bool {
	if o != nil && o.NapalmDriver != nil {
		return true
	}

	return false
}

// SetNapalmDriver gets a reference to the given string and assigns it to the NapalmDriver field.
func (o *Platform) SetNapalmDriver(v string) {
	o.NapalmDriver = &v
}

// GetNapalmArgs returns the NapalmArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Platform) GetNapalmArgs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NapalmArgs
}

// GetNapalmArgsOk returns a tuple with the NapalmArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Platform) GetNapalmArgsOk() (map[string]interface{}, bool) {
	if o == nil || o.NapalmArgs == nil {
		return nil, false
	}
	return o.NapalmArgs, true
}

// HasNapalmArgs returns a boolean if a field has been set.
func (o *Platform) HasNapalmArgs() bool {
	if o != nil && o.NapalmArgs != nil {
		return true
	}

	return false
}

// SetNapalmArgs gets a reference to the given map[string]interface{} and assigns it to the NapalmArgs field.
func (o *Platform) SetNapalmArgs(v map[string]interface{}) {
	o.NapalmArgs = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Platform) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Platform) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Platform) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceCount returns the DeviceCount field value
func (o *Platform) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Platform) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Platform) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Platform) GetVirtualmachineCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Platform) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Platform) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Platform) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Platform) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Platform) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Platform) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *Platform) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Platform) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Platform) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Platform) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Platform) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Platform) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetDisplay returns the Display field value
func (o *Platform) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Platform) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Platform) SetDisplay(v string) {
	o.Display = v
}

func (o Platform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.NapalmDriver != nil {
		toSerialize["napalm_driver"] = o.NapalmDriver
	}
	if o.NapalmArgs != nil {
		toSerialize["napalm_args"] = o.NapalmArgs
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["device_count"] = o.DeviceCount
	}
	if true {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePlatform struct {
	value *Platform
	isSet bool
}

func (v NullablePlatform) Get() *Platform {
	return v.value
}

func (v *NullablePlatform) Set(val *Platform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatform(val *Platform) *NullablePlatform {
	return &NullablePlatform{value: val, isSet: true}
}

func (v NullablePlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


