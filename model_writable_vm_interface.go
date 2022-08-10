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

// WritableVMInterface Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableVMInterface struct {
	Id string `json:"id"`
	Url string `json:"url"`
	VirtualMachine string `json:"virtual_machine"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	MacAddress NullableString `json:"mac_address,omitempty"`
	Description *string `json:"description,omitempty"`
	Mode *PatchedWritableInterfaceMode `json:"mode,omitempty"`
	UntaggedVlan NullableString `json:"untagged_vlan,omitempty"`
	TaggedVlans []string `json:"tagged_vlans,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableVMInterface instantiates a new WritableVMInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVMInterface(id string, url string, virtualMachine string, name string, display string) *WritableVMInterface {
	this := WritableVMInterface{}
	this.Id = id
	this.Url = url
	this.VirtualMachine = virtualMachine
	this.Name = name
	this.Display = display
	return &this
}

// NewWritableVMInterfaceWithDefaults instantiates a new WritableVMInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVMInterfaceWithDefaults() *WritableVMInterface {
	this := WritableVMInterface{}
	return &this
}

// GetId returns the Id field value
func (o *WritableVMInterface) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableVMInterface) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableVMInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableVMInterface) SetUrl(v string) {
	o.Url = v
}

// GetVirtualMachine returns the VirtualMachine field value
func (o *WritableVMInterface) GetVirtualMachine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetVirtualMachineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualMachine, true
}

// SetVirtualMachine sets field value
func (o *WritableVMInterface) SetVirtualMachine(v string) {
	o.VirtualMachine = v
}

// GetName returns the Name field value
func (o *WritableVMInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVMInterface) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableVMInterface) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableVMInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableVMInterface) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterface) GetMtu() int32 {
	if o == nil || o.Mtu.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterface) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *WritableVMInterface) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *WritableVMInterface) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *WritableVMInterface) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *WritableVMInterface) UnsetMtu() {
	o.Mtu.Unset()
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterface) GetMacAddress() string {
	if o == nil || o.MacAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterface) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WritableVMInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *WritableVMInterface) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}
// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *WritableVMInterface) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *WritableVMInterface) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVMInterface) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVMInterface) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVMInterface) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WritableVMInterface) GetMode() PatchedWritableInterfaceMode {
	if o == nil || o.Mode == nil {
		var ret PatchedWritableInterfaceMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetModeOk() (*PatchedWritableInterfaceMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WritableVMInterface) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given PatchedWritableInterfaceMode and assigns it to the Mode field.
func (o *WritableVMInterface) SetMode(v PatchedWritableInterfaceMode) {
	o.Mode = &v
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVMInterface) GetUntaggedVlan() string {
	if o == nil || o.UntaggedVlan.Get() == nil {
		var ret string
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVMInterface) GetUntaggedVlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *WritableVMInterface) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableString and assigns it to the UntaggedVlan field.
func (o *WritableVMInterface) SetUntaggedVlan(v string) {
	o.UntaggedVlan.Set(&v)
}
// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *WritableVMInterface) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *WritableVMInterface) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *WritableVMInterface) GetTaggedVlans() []string {
	if o == nil || o.TaggedVlans == nil {
		var ret []string
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetTaggedVlansOk() ([]string, bool) {
	if o == nil || o.TaggedVlans == nil {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *WritableVMInterface) HasTaggedVlans() bool {
	if o != nil && o.TaggedVlans != nil {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []string and assigns it to the TaggedVlans field.
func (o *WritableVMInterface) SetTaggedVlans(v []string) {
	o.TaggedVlans = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVMInterface) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVMInterface) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritableVMInterface) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetDisplay returns the Display field value
func (o *WritableVMInterface) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableVMInterface) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableVMInterface) SetDisplay(v string) {
	o.Display = v
}

func (o WritableVMInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["virtual_machine"] = o.VirtualMachine
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.MacAddress.IsSet() {
		toSerialize["mac_address"] = o.MacAddress.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if o.TaggedVlans != nil {
		toSerialize["tagged_vlans"] = o.TaggedVlans
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritableVMInterface struct {
	value *WritableVMInterface
	isSet bool
}

func (v NullableWritableVMInterface) Get() *WritableVMInterface {
	return v.value
}

func (v *NullableWritableVMInterface) Set(val *WritableVMInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVMInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVMInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVMInterface(val *WritableVMInterface) *NullableWritableVMInterface {
	return &NullableWritableVMInterface{value: val, isSet: true}
}

func (v NullableWritableVMInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVMInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


