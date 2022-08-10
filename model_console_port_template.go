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

// ConsolePortTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type ConsolePortTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	DeviceType NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *ConsolePortType `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewConsolePortTemplate instantiates a new ConsolePortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePortTemplate(id string, url string, deviceType NestedDeviceType, name string, display string) *ConsolePortTemplate {
	this := ConsolePortTemplate{}
	this.Id = id
	this.Url = url
	this.DeviceType = deviceType
	this.Name = name
	this.Display = display
	return &this
}

// NewConsolePortTemplateWithDefaults instantiates a new ConsolePortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortTemplateWithDefaults() *ConsolePortTemplate {
	this := ConsolePortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ConsolePortTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsolePortTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConsolePortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConsolePortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDeviceType returns the DeviceType field value
func (o *ConsolePortTemplate) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *ConsolePortTemplate) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *ConsolePortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsolePortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePortTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePortTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsolePortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConsolePortTemplate) GetType() ConsolePortType {
	if o == nil || o.Type == nil {
		var ret ConsolePortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetTypeOk() (*ConsolePortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConsolePortTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ConsolePortType and assigns it to the Type field.
func (o *ConsolePortTemplate) SetType(v ConsolePortType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConsolePortTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConsolePortTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConsolePortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ConsolePortTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ConsolePortTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ConsolePortTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *ConsolePortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConsolePortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConsolePortTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o ConsolePortTemplate) MarshalJSON() ([]byte, error) {
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

type NullableConsolePortTemplate struct {
	value *ConsolePortTemplate
	isSet bool
}

func (v NullableConsolePortTemplate) Get() *ConsolePortTemplate {
	return v.value
}

func (v *NullableConsolePortTemplate) Set(val *ConsolePortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortTemplate(val *ConsolePortTemplate) *NullableConsolePortTemplate {
	return &NullableConsolePortTemplate{value: val, isSet: true}
}

func (v NullableConsolePortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


