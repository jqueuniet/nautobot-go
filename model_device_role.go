/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// DeviceRole Extends ModelSerializer to render any CustomFields and their values associated with an object.
type DeviceRole struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	Color *string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole *bool `json:"vm_role,omitempty"`
	Description *string `json:"description,omitempty"`
	DeviceCount int32 `json:"device_count"`
	VirtualmachineCount int32 `json:"virtualmachine_count"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewDeviceRole instantiates a new DeviceRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceRole(id string, url string, name string, deviceCount int32, virtualmachineCount int32, created string, lastUpdated time.Time, display string) *DeviceRole {
	this := DeviceRole{}
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

// NewDeviceRoleWithDefaults instantiates a new DeviceRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceRoleWithDefaults() *DeviceRole {
	this := DeviceRole{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceRole) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DeviceRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceRole) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *DeviceRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *DeviceRole) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *DeviceRole) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *DeviceRole) SetSlug(v string) {
	o.Slug = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *DeviceRole) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *DeviceRole) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *DeviceRole) SetColor(v string) {
	o.Color = &v
}

// GetVmRole returns the VmRole field value if set, zero value otherwise.
func (o *DeviceRole) GetVmRole() bool {
	if o == nil || o.VmRole == nil {
		var ret bool
		return ret
	}
	return *o.VmRole
}

// GetVmRoleOk returns a tuple with the VmRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetVmRoleOk() (*bool, bool) {
	if o == nil || o.VmRole == nil {
		return nil, false
	}
	return o.VmRole, true
}

// HasVmRole returns a boolean if a field has been set.
func (o *DeviceRole) HasVmRole() bool {
	if o != nil && o.VmRole != nil {
		return true
	}

	return false
}

// SetVmRole gets a reference to the given bool and assigns it to the VmRole field.
func (o *DeviceRole) SetVmRole(v bool) {
	o.VmRole = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceRole) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceRole) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceRole) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceCount returns the DeviceCount field value
func (o *DeviceRole) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *DeviceRole) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *DeviceRole) GetVirtualmachineCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *DeviceRole) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceRole) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceRole) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceRole) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *DeviceRole) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DeviceRole) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *DeviceRole) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *DeviceRole) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetDisplay returns the Display field value
func (o *DeviceRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceRole) SetDisplay(v string) {
	o.Display = v
}

func (o DeviceRole) MarshalJSON() ([]byte, error) {
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
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.VmRole != nil {
		toSerialize["vm_role"] = o.VmRole
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

type NullableDeviceRole struct {
	value *DeviceRole
	isSet bool
}

func (v NullableDeviceRole) Get() *DeviceRole {
	return v.value
}

func (v *NullableDeviceRole) Set(val *DeviceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceRole(val *DeviceRole) *NullableDeviceRole {
	return &NullableDeviceRole{value: val, isSet: true}
}

func (v NullableDeviceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


