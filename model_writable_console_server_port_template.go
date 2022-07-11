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

// WritableConsoleServerPortTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type WritableConsoleServerPortTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	DeviceType string `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PatchedWritableConsolePortTemplateType `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableConsoleServerPortTemplate instantiates a new WritableConsoleServerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableConsoleServerPortTemplate(id string, url string, deviceType string, name string, display string) *WritableConsoleServerPortTemplate {
	this := WritableConsoleServerPortTemplate{}
	this.Id = id
	this.Url = url
	this.DeviceType = deviceType
	this.Name = name
	this.Display = display
	return &this
}

// NewWritableConsoleServerPortTemplateWithDefaults instantiates a new WritableConsoleServerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableConsoleServerPortTemplateWithDefaults() *WritableConsoleServerPortTemplate {
	this := WritableConsoleServerPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *WritableConsoleServerPortTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableConsoleServerPortTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableConsoleServerPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableConsoleServerPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDeviceType returns the DeviceType field value
func (o *WritableConsoleServerPortTemplate) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *WritableConsoleServerPortTemplate) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *WritableConsoleServerPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableConsoleServerPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableConsoleServerPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplate) GetType() PatchedWritableConsolePortTemplateType {
	if o == nil || o.Type == nil {
		var ret PatchedWritableConsolePortTemplateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetTypeOk() (*PatchedWritableConsolePortTemplateType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritableConsolePortTemplateType and assigns it to the Type field.
func (o *WritableConsoleServerPortTemplate) SetType(v PatchedWritableConsolePortTemplateType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableConsoleServerPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableConsoleServerPortTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *WritableConsoleServerPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableConsoleServerPortTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o WritableConsoleServerPortTemplate) MarshalJSON() ([]byte, error) {
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

type NullableWritableConsoleServerPortTemplate struct {
	value *WritableConsoleServerPortTemplate
	isSet bool
}

func (v NullableWritableConsoleServerPortTemplate) Get() *WritableConsoleServerPortTemplate {
	return v.value
}

func (v *NullableWritableConsoleServerPortTemplate) Set(val *WritableConsoleServerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableConsoleServerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableConsoleServerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableConsoleServerPortTemplate(val *WritableConsoleServerPortTemplate) *NullableWritableConsoleServerPortTemplate {
	return &NullableWritableConsoleServerPortTemplate{value: val, isSet: true}
}

func (v NullableWritableConsoleServerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableConsoleServerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


