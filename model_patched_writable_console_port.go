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

// PatchedWritableConsolePort Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritableConsolePort struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Device *string `json:"device,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PatchedWritableConsolePortType `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Cable *CircuitTerminationCable `json:"cable,omitempty"`
	CablePeer map[string]interface{} `json:"cable_peer,omitempty"`
	CablePeerType NullableString `json:"cable_peer_type,omitempty"`
	ConnectedEndpoint map[string]interface{} `json:"connected_endpoint,omitempty"`
	ConnectedEndpointType NullableString `json:"connected_endpoint_type,omitempty"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableConsolePort instantiates a new PatchedWritableConsolePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableConsolePort() *PatchedWritableConsolePort {
	this := PatchedWritableConsolePort{}
	return &this
}

// NewPatchedWritableConsolePortWithDefaults instantiates a new PatchedWritableConsolePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableConsolePortWithDefaults() *PatchedWritableConsolePort {
	this := PatchedWritableConsolePort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableConsolePort) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableConsolePort) SetUrl(v string) {
	o.Url = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *PatchedWritableConsolePort) SetDevice(v string) {
	o.Device = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableConsolePort) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableConsolePort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetType() PatchedWritableConsolePortType {
	if o == nil || o.Type == nil {
		var ret PatchedWritableConsolePortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetTypeOk() (*PatchedWritableConsolePortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableConsolePortType and assigns it to the Type field.
func (o *PatchedWritableConsolePort) SetType(v PatchedWritableConsolePortType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableConsolePort) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetCable() CircuitTerminationCable {
	if o == nil || o.Cable == nil {
		var ret CircuitTerminationCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given CircuitTerminationCable and assigns it to the Cable field.
func (o *PatchedWritableConsolePort) SetCable(v CircuitTerminationCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePort) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePort) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]interface{} and assigns it to the CablePeer field.
func (o *PatchedWritableConsolePort) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasCablePeerType() bool {
	if o != nil && o.CablePeerType.IsSet() {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given NullableString and assigns it to the CablePeerType field.
func (o *PatchedWritableConsolePort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}
// SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil
func (o *PatchedWritableConsolePort) SetCablePeerTypeNil() {
	o.CablePeerType.Set(nil)
}

// UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
func (o *PatchedWritableConsolePort) UnsetCablePeerType() {
	o.CablePeerType.Unset()
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePort) GetConnectedEndpoint() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePort) GetConnectedEndpointOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// HasConnectedEndpoint returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasConnectedEndpoint() bool {
	if o != nil && o.ConnectedEndpoint != nil {
		return true
	}

	return false
}

// SetConnectedEndpoint gets a reference to the given map[string]interface{} and assigns it to the ConnectedEndpoint field.
func (o *PatchedWritableConsolePort) SetConnectedEndpoint(v map[string]interface{}) {
	o.ConnectedEndpoint = v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePort) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePort) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// HasConnectedEndpointType returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasConnectedEndpointType() bool {
	if o != nil && o.ConnectedEndpointType.IsSet() {
		return true
	}

	return false
}

// SetConnectedEndpointType gets a reference to the given NullableString and assigns it to the ConnectedEndpointType field.
func (o *PatchedWritableConsolePort) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}
// SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil
func (o *PatchedWritableConsolePort) SetConnectedEndpointTypeNil() {
	o.ConnectedEndpointType.Set(nil)
}

// UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
func (o *PatchedWritableConsolePort) UnsetConnectedEndpointType() {
	o.ConnectedEndpointType.Unset()
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConsolePort) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConsolePort) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// HasConnectedEndpointReachable returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasConnectedEndpointReachable() bool {
	if o != nil && o.ConnectedEndpointReachable.IsSet() {
		return true
	}

	return false
}

// SetConnectedEndpointReachable gets a reference to the given NullableBool and assigns it to the ConnectedEndpointReachable field.
func (o *PatchedWritableConsolePort) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}
// SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil
func (o *PatchedWritableConsolePort) SetConnectedEndpointReachableNil() {
	o.ConnectedEndpointReachable.Set(nil)
}

// UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
func (o *PatchedWritableConsolePort) UnsetConnectedEndpointReachable() {
	o.ConnectedEndpointReachable.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableConsolePort) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableConsolePort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableConsolePort) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableConsolePort) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConsolePort) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableConsolePort) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableConsolePort) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableConsolePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Cable != nil {
		toSerialize["cable"] = o.Cable
	}
	if o.CablePeer != nil {
		toSerialize["cable_peer"] = o.CablePeer
	}
	if o.CablePeerType.IsSet() {
		toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	}
	if o.ConnectedEndpoint != nil {
		toSerialize["connected_endpoint"] = o.ConnectedEndpoint
	}
	if o.ConnectedEndpointType.IsSet() {
		toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType.Get()
	}
	if o.ConnectedEndpointReachable.IsSet() {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableConsolePort struct {
	value *PatchedWritableConsolePort
	isSet bool
}

func (v NullablePatchedWritableConsolePort) Get() *PatchedWritableConsolePort {
	return v.value
}

func (v *NullablePatchedWritableConsolePort) Set(val *PatchedWritableConsolePort) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsolePort) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsolePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsolePort(val *PatchedWritableConsolePort) *NullablePatchedWritableConsolePort {
	return &NullablePatchedWritableConsolePort{value: val, isSet: true}
}

func (v NullablePatchedWritableConsolePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsolePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


