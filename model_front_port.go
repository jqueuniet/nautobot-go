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

// FrontPort Extends ModelSerializer to render any CustomFields and their values associated with an object.
type FrontPort struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Device NestedDevice `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type FrontPortType `json:"type"`
	RearPort FrontPortRearPort `json:"rear_port"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	Description *string `json:"description,omitempty"`
	Cable CircuitTerminationCable `json:"cable"`
	CablePeer map[string]interface{} `json:"cable_peer"`
	CablePeerType NullableString `json:"cable_peer_type"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewFrontPort instantiates a new FrontPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPort(id string, url string, device NestedDevice, name string, type_ FrontPortType, rearPort FrontPortRearPort, cable CircuitTerminationCable, cablePeer map[string]interface{}, cablePeerType NullableString, display string) *FrontPort {
	this := FrontPort{}
	this.Id = id
	this.Url = url
	this.Device = device
	this.Name = name
	this.Type = type_
	this.RearPort = rearPort
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	this.Cable = cable
	this.CablePeer = cablePeer
	this.CablePeerType = cablePeerType
	this.Display = display
	return &this
}

// NewFrontPortWithDefaults instantiates a new FrontPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortWithDefaults() *FrontPort {
	this := FrontPort{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetId returns the Id field value
func (o *FrontPort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FrontPort) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *FrontPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FrontPort) SetUrl(v string) {
	o.Url = v
}

// GetDevice returns the Device field value
func (o *FrontPort) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *FrontPort) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetName returns the Name field value
func (o *FrontPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPort) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPort) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *FrontPort) GetType() FrontPortType {
	if o == nil {
		var ret FrontPortType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetTypeOk() (*FrontPortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPort) SetType(v FrontPortType) {
	o.Type = v
}

// GetRearPort returns the RearPort field value
func (o *FrontPort) GetRearPort() FrontPortRearPort {
	if o == nil {
		var ret FrontPortRearPort
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetRearPortOk() (*FrontPortRearPort, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *FrontPort) SetRearPort(v FrontPortRearPort) {
	o.RearPort = v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPort) GetRearPortPosition() int32 {
	if o == nil || o.RearPortPosition == nil {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || o.RearPortPosition == nil {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPort) HasRearPortPosition() bool {
	if o != nil && o.RearPortPosition != nil {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPort) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPort) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPort) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPort) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value
func (o *FrontPort) GetCable() CircuitTerminationCable {
	if o == nil {
		var ret CircuitTerminationCable
		return ret
	}

	return o.Cable
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cable, true
}

// SetCable sets field value
func (o *FrontPort) SetCable(v CircuitTerminationCable) {
	o.Cable = v
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *FrontPort) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// SetCablePeer sets field value
func (o *FrontPort) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FrontPort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *FrontPort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FrontPort) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FrontPort) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *FrontPort) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FrontPort) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FrontPort) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FrontPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *FrontPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FrontPort) SetDisplay(v string) {
	o.Display = v
}

func (o FrontPort) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["rear_port"] = o.RearPort
	}
	if o.RearPortPosition != nil {
		toSerialize["rear_port_position"] = o.RearPortPosition
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
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableFrontPort struct {
	value *FrontPort
	isSet bool
}

func (v NullableFrontPort) Get() *FrontPort {
	return v.value
}

func (v *NullableFrontPort) Set(val *FrontPort) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPort) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPort(val *FrontPort) *NullableFrontPort {
	return &NullableFrontPort{value: val, isSet: true}
}

func (v NullableFrontPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


