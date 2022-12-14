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

// PowerOutletTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PowerOutletTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	DeviceType NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PowerOutletType `json:"type,omitempty"`
	PowerPort *NestedPowerPortTemplate `json:"power_port,omitempty"`
	FeedLeg *PowerOutletFeedLeg `json:"feed_leg,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewPowerOutletTemplate instantiates a new PowerOutletTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerOutletTemplate(id string, url string, deviceType NestedDeviceType, name string, display string) *PowerOutletTemplate {
	this := PowerOutletTemplate{}
	this.Id = id
	this.Url = url
	this.DeviceType = deviceType
	this.Name = name
	this.Display = display
	return &this
}

// NewPowerOutletTemplateWithDefaults instantiates a new PowerOutletTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerOutletTemplateWithDefaults() *PowerOutletTemplate {
	this := PowerOutletTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *PowerOutletTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerOutletTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerOutletTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerOutletTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDeviceType returns the DeviceType field value
func (o *PowerOutletTemplate) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *PowerOutletTemplate) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *PowerOutletTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerOutletTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PowerOutletTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetType() PowerOutletType {
	if o == nil || o.Type == nil {
		var ret PowerOutletType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetTypeOk() (*PowerOutletType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerOutletType and assigns it to the Type field.
func (o *PowerOutletTemplate) SetType(v PowerOutletType) {
	o.Type = &v
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetPowerPort() NestedPowerPortTemplate {
	if o == nil || o.PowerPort == nil {
		var ret NestedPowerPortTemplate
		return ret
	}
	return *o.PowerPort
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetPowerPortOk() (*NestedPowerPortTemplate, bool) {
	if o == nil || o.PowerPort == nil {
		return nil, false
	}
	return o.PowerPort, true
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasPowerPort() bool {
	if o != nil && o.PowerPort != nil {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NestedPowerPortTemplate and assigns it to the PowerPort field.
func (o *PowerOutletTemplate) SetPowerPort(v NestedPowerPortTemplate) {
	o.PowerPort = &v
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetFeedLeg() PowerOutletFeedLeg {
	if o == nil || o.FeedLeg == nil {
		var ret PowerOutletFeedLeg
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetFeedLegOk() (*PowerOutletFeedLeg, bool) {
	if o == nil || o.FeedLeg == nil {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasFeedLeg() bool {
	if o != nil && o.FeedLeg != nil {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given PowerOutletFeedLeg and assigns it to the FeedLeg field.
func (o *PowerOutletTemplate) SetFeedLeg(v PowerOutletFeedLeg) {
	o.FeedLeg = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerOutletTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerOutletTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerOutletTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerOutletTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *PowerOutletTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerOutletTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerOutletTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o PowerOutletTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["device_type"] = o.DeviceType
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
	if o.PowerPort != nil {
		toSerialize["power_port"] = o.PowerPort
	}
	if o.FeedLeg != nil {
		toSerialize["feed_leg"] = o.FeedLeg
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePowerOutletTemplate struct {
	value *PowerOutletTemplate
	isSet bool
}

func (v NullablePowerOutletTemplate) Get() *PowerOutletTemplate {
	return v.value
}

func (v *NullablePowerOutletTemplate) Set(val *PowerOutletTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletTemplate(val *PowerOutletTemplate) *NullablePowerOutletTemplate {
	return &NullablePowerOutletTemplate{value: val, isSet: true}
}

func (v NullablePowerOutletTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


