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

// InterfaceTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type InterfaceTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	DeviceType NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type InterfaceType `json:"type"`
	MgmtOnly *bool `json:"mgmt_only,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewInterfaceTemplate instantiates a new InterfaceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplate(id string, url string, deviceType NestedDeviceType, name string, type_ InterfaceType, display string) *InterfaceTemplate {
	this := InterfaceTemplate{}
	this.Id = id
	this.Url = url
	this.DeviceType = deviceType
	this.Name = name
	this.Type = type_
	this.Display = display
	return &this
}

// NewInterfaceTemplateWithDefaults instantiates a new InterfaceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplateWithDefaults() *InterfaceTemplate {
	this := InterfaceTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *InterfaceTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InterfaceTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDeviceType returns the DeviceType field value
func (o *InterfaceTemplate) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *InterfaceTemplate) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *InterfaceTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterfaceTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *InterfaceTemplate) GetType() InterfaceType {
	if o == nil {
		var ret InterfaceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetTypeOk() (*InterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InterfaceTemplate) SetType(v InterfaceType) {
	o.Type = v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetMgmtOnly() bool {
	if o == nil || o.MgmtOnly == nil {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || o.MgmtOnly == nil {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasMgmtOnly() bool {
	if o != nil && o.MgmtOnly != nil {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *InterfaceTemplate) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InterfaceTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *InterfaceTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InterfaceTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o InterfaceTemplate) MarshalJSON() ([]byte, error) {
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
	if o.MgmtOnly != nil {
		toSerialize["mgmt_only"] = o.MgmtOnly
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

type NullableInterfaceTemplate struct {
	value *InterfaceTemplate
	isSet bool
}

func (v NullableInterfaceTemplate) Get() *InterfaceTemplate {
	return v.value
}

func (v *NullableInterfaceTemplate) Set(val *InterfaceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplate(val *InterfaceTemplate) *NullableInterfaceTemplate {
	return &NullableInterfaceTemplate{value: val, isSet: true}
}

func (v NullableInterfaceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


