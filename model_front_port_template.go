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

// FrontPortTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type FrontPortTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	DeviceType NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type FrontPortType `json:"type"`
	RearPort NestedRearPortTemplate `json:"rear_port"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewFrontPortTemplate instantiates a new FrontPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortTemplate(id string, url string, deviceType NestedDeviceType, name string, type_ FrontPortType, rearPort NestedRearPortTemplate, display string) *FrontPortTemplate {
	this := FrontPortTemplate{}
	this.Id = id
	this.Url = url
	this.DeviceType = deviceType
	this.Name = name
	this.Type = type_
	this.RearPort = rearPort
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	this.Display = display
	return &this
}

// NewFrontPortTemplateWithDefaults instantiates a new FrontPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortTemplateWithDefaults() *FrontPortTemplate {
	this := FrontPortTemplate{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetId returns the Id field value
func (o *FrontPortTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FrontPortTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *FrontPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FrontPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDeviceType returns the DeviceType field value
func (o *FrontPortTemplate) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *FrontPortTemplate) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *FrontPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *FrontPortTemplate) GetType() FrontPortType {
	if o == nil {
		var ret FrontPortType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetTypeOk() (*FrontPortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPortTemplate) SetType(v FrontPortType) {
	o.Type = v
}

// GetRearPort returns the RearPort field value
func (o *FrontPortTemplate) GetRearPort() NestedRearPortTemplate {
	if o == nil {
		var ret NestedRearPortTemplate
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortOk() (*NestedRearPortTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *FrontPortTemplate) SetRearPort(v NestedRearPortTemplate) {
	o.RearPort = v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetRearPortPosition() int32 {
	if o == nil || o.RearPortPosition == nil {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || o.RearPortPosition == nil {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasRearPortPosition() bool {
	if o != nil && o.RearPortPosition != nil {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPortTemplate) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FrontPortTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *FrontPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FrontPortTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o FrontPortTemplate) MarshalJSON() ([]byte, error) {
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
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableFrontPortTemplate struct {
	value *FrontPortTemplate
	isSet bool
}

func (v NullableFrontPortTemplate) Get() *FrontPortTemplate {
	return v.value
}

func (v *NullableFrontPortTemplate) Set(val *FrontPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTemplate(val *FrontPortTemplate) *NullableFrontPortTemplate {
	return &NullableFrontPortTemplate{value: val, isSet: true}
}

func (v NullableFrontPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


