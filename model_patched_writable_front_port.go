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

// PatchedWritableFrontPort Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritableFrontPort struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Device *string `json:"device,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PortTypeChoices `json:"type,omitempty"`
	RearPort *string `json:"rear_port,omitempty"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	Description *string `json:"description,omitempty"`
	Cable *CircuitTerminationCable `json:"cable,omitempty"`
	CablePeer map[string]interface{} `json:"cable_peer,omitempty"`
	CablePeerType NullableString `json:"cable_peer_type,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableFrontPort instantiates a new PatchedWritableFrontPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableFrontPort() *PatchedWritableFrontPort {
	this := PatchedWritableFrontPort{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// NewPatchedWritableFrontPortWithDefaults instantiates a new PatchedWritableFrontPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableFrontPortWithDefaults() *PatchedWritableFrontPort {
	this := PatchedWritableFrontPort{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableFrontPort) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableFrontPort) SetUrl(v string) {
	o.Url = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *PatchedWritableFrontPort) SetDevice(v string) {
	o.Device = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableFrontPort) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableFrontPort) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetType() PortTypeChoices {
	if o == nil || o.Type == nil {
		var ret PortTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetTypeOk() (*PortTypeChoices, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PortTypeChoices and assigns it to the Type field.
func (o *PatchedWritableFrontPort) SetType(v PortTypeChoices) {
	o.Type = &v
}

// GetRearPort returns the RearPort field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetRearPort() string {
	if o == nil || o.RearPort == nil {
		var ret string
		return ret
	}
	return *o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetRearPortOk() (*string, bool) {
	if o == nil || o.RearPort == nil {
		return nil, false
	}
	return o.RearPort, true
}

// HasRearPort returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasRearPort() bool {
	if o != nil && o.RearPort != nil {
		return true
	}

	return false
}

// SetRearPort gets a reference to the given string and assigns it to the RearPort field.
func (o *PatchedWritableFrontPort) SetRearPort(v string) {
	o.RearPort = &v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetRearPortPosition() int32 {
	if o == nil || o.RearPortPosition == nil {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || o.RearPortPosition == nil {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasRearPortPosition() bool {
	if o != nil && o.RearPortPosition != nil {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *PatchedWritableFrontPort) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableFrontPort) SetDescription(v string) {
	o.Description = &v
}

// GetCable returns the Cable field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetCable() CircuitTerminationCable {
	if o == nil || o.Cable == nil {
		var ret CircuitTerminationCable
		return ret
	}
	return *o.Cable
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetCableOk() (*CircuitTerminationCable, bool) {
	if o == nil || o.Cable == nil {
		return nil, false
	}
	return o.Cable, true
}

// HasCable returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasCable() bool {
	if o != nil && o.Cable != nil {
		return true
	}

	return false
}

// SetCable gets a reference to the given CircuitTerminationCable and assigns it to the Cable field.
func (o *PatchedWritableFrontPort) SetCable(v CircuitTerminationCable) {
	o.Cable = &v
}

// GetCablePeer returns the CablePeer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableFrontPort) GetCablePeer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CablePeer
}

// GetCablePeerOk returns a tuple with the CablePeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableFrontPort) GetCablePeerOk() (map[string]interface{}, bool) {
	if o == nil || o.CablePeer == nil {
		return nil, false
	}
	return o.CablePeer, true
}

// HasCablePeer returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasCablePeer() bool {
	if o != nil && o.CablePeer != nil {
		return true
	}

	return false
}

// SetCablePeer gets a reference to the given map[string]interface{} and assigns it to the CablePeer field.
func (o *PatchedWritableFrontPort) SetCablePeer(v map[string]interface{}) {
	o.CablePeer = v
}

// GetCablePeerType returns the CablePeerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableFrontPort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableFrontPort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// HasCablePeerType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasCablePeerType() bool {
	if o != nil && o.CablePeerType.IsSet() {
		return true
	}

	return false
}

// SetCablePeerType gets a reference to the given NullableString and assigns it to the CablePeerType field.
func (o *PatchedWritableFrontPort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}
// SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil
func (o *PatchedWritableFrontPort) SetCablePeerTypeNil() {
	o.CablePeerType.Set(nil)
}

// UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
func (o *PatchedWritableFrontPort) UnsetCablePeerType() {
	o.CablePeerType.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableFrontPort) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableFrontPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableFrontPort) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableFrontPort) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPort) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableFrontPort) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableFrontPort) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableFrontPort) MarshalJSON() ([]byte, error) {
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
	if o.RearPort != nil {
		toSerialize["rear_port"] = o.RearPort
	}
	if o.RearPortPosition != nil {
		toSerialize["rear_port_position"] = o.RearPortPosition
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

type NullablePatchedWritableFrontPort struct {
	value *PatchedWritableFrontPort
	isSet bool
}

func (v NullablePatchedWritableFrontPort) Get() *PatchedWritableFrontPort {
	return v.value
}

func (v *NullablePatchedWritableFrontPort) Set(val *PatchedWritableFrontPort) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableFrontPort) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableFrontPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableFrontPort(val *PatchedWritableFrontPort) *NullablePatchedWritableFrontPort {
	return &NullablePatchedWritableFrontPort{value: val, isSet: true}
}

func (v NullablePatchedWritableFrontPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableFrontPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


