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

// ConsolePort Extends ModelSerializer to render any CustomFields and their values associated with an object.
type ConsolePort struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Device NestedDevice `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *ConsolePortType `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	Cable CircuitTerminationCable `json:"cable"`
	CablePeer map[string]interface{} `json:"cable_peer"`
	CablePeerType NullableString `json:"cable_peer_type"`
	ConnectedEndpoint map[string]interface{} `json:"connected_endpoint"`
	ConnectedEndpointType NullableString `json:"connected_endpoint_type"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewConsolePort instantiates a new ConsolePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePort(id string, url string, device NestedDevice, name string, cable CircuitTerminationCable, cablePeer map[string]interface{}, cablePeerType NullableString, connectedEndpoint map[string]interface{}, connectedEndpointType NullableString, connectedEndpointReachable NullableBool, computedFields map[string]interface{}, display string) *ConsolePort {
	this := ConsolePort{}
	this.Id = id
	this.Url = url
	this.Device = device
	this.Name = name
	this.Cable = cable
	this.CablePeer = cablePeer
	this.CablePeerType = cablePeerType
	this.ConnectedEndpoint = connectedEndpoint
	this.ConnectedEndpointType = connectedEndpointType
	this.ConnectedEndpointReachable = connectedEndpointReachable
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewConsolePortWithDefaults instantiates a new ConsolePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortWithDefaults() *ConsolePort {
	this := ConsolePort{}
	return &this
}

// GetId returns the Id field value
func (o *ConsolePort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsolePort) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConsolePort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConsolePort) SetUrl(v string) {
	o.Url = v
}

// GetDevice returns the Device field value
func (o *ConsolePort) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *ConsolePort) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *ConsolePort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsolePort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePort) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePort) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsolePort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsolePort) GetType() ConsolePortType {
	if o == nil || o.Type == nil {
		var ret ConsolePortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetTypeOk() (*ConsolePortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsolePort) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortType and assigns it to the Type field.
func (o *ConsolePort) SetType(v ConsolePortType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsolePort) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsolePort) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsolePort) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value
func (o *ConsolePort) GetCable() CircuitTerminationCable {
	if o == nil {
		var ret CircuitTerminationCable
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *ConsolePort) SetCable(v CircuitTerminationCable) {
	o.Cable = v
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ConsolePort) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// SetCablePeer sets field value
func (o *ConsolePort) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsolePort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *ConsolePort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetConnectedEndpoint returns the ConnectedEndpoint field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ConsolePort) GetConnectedEndpoint() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ConnectedEndpoint
}

// GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectedEndpoint == nil {
		return nil, false
	}
	return o.ConnectedEndpoint, true
}

// SetConnectedEndpoint sets field value
func (o *ConsolePort) SetConnectedEndpoint(v map[string]interface{}) {
	o.ConnectedEndpoint = v
}

// GetConnectedEndpointType returns the ConnectedEndpointType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConsolePort) GetConnectedEndpointType() string {
	if o == nil || o.ConnectedEndpointType.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConnectedEndpointType.Get()
}

// GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointType.Get(), o.ConnectedEndpointType.IsSet()
}

// SetConnectedEndpointType sets field value
func (o *ConsolePort) SetConnectedEndpointType(v string) {
	o.ConnectedEndpointType.Set(&v)
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ConsolePort) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsolePort) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// SetConnectedEndpointReachable sets field value
func (o *ConsolePort) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConsolePort) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConsolePort) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *ConsolePort) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ConsolePort) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ConsolePort) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ConsolePort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComputedFields returns the ComputedFields field value
func (o *ConsolePort) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *ConsolePort) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *ConsolePort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConsolePort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConsolePort) SetDisplay(v string) {
	o.Display = v
}

func (o ConsolePort) MarshalJSON() ([]byte, error) {
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
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["cable"] = o.Cable
	}
	if o.CablePeer != nil {
		toSerialize["cable_peer"] = o.CablePeer
	}
	if true {
		toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	}
	if o.ConnectedEndpoint != nil {
		toSerialize["connected_endpoint"] = o.ConnectedEndpoint
	}
	if true {
		toSerialize["connected_endpoint_type"] = o.ConnectedEndpointType.Get()
	}
	if true {
		toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableConsolePort struct {
	value *ConsolePort
	isSet bool
}

func (v NullableConsolePort) Get() *ConsolePort {
	return v.value
}

func (v *NullableConsolePort) Set(val *ConsolePort) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePort) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePort(val *ConsolePort) *NullableConsolePort {
	return &NullableConsolePort{value: val, isSet: true}
}

func (v NullableConsolePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


