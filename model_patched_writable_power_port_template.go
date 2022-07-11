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

// PatchedWritablePowerPortTemplate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritablePowerPortTemplate struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	DeviceType *string `json:"device_type,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Type *PatchedWritablePowerPortTemplateType `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw NullableInt32 `json:"allocated_draw,omitempty"`
	Description *string `json:"description,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritablePowerPortTemplate instantiates a new PatchedWritablePowerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerPortTemplate() *PatchedWritablePowerPortTemplate {
	this := PatchedWritablePowerPortTemplate{}
	return &this
}

// NewPatchedWritablePowerPortTemplateWithDefaults instantiates a new PatchedWritablePowerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerPortTemplateWithDefaults() *PatchedWritablePowerPortTemplate {
	this := PatchedWritablePowerPortTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritablePowerPortTemplate) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritablePowerPortTemplate) SetUrl(v string) {
	o.Url = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *PatchedWritablePowerPortTemplate) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerPortTemplate) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritablePowerPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetType() PatchedWritablePowerPortTemplateType {
	if o == nil || o.Type == nil {
		var ret PatchedWritablePowerPortTemplateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetTypeOk() (*PatchedWritablePowerPortTemplateType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PatchedWritablePowerPortTemplateType and assigns it to the Type field.
func (o *PatchedWritablePowerPortTemplate) SetType(v PatchedWritablePowerPortTemplateType) {
	o.Type = &v
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortTemplate) GetMaximumDraw() int32 {
	if o == nil || o.MaximumDraw.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortTemplate) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *PatchedWritablePowerPortTemplate) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}
// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *PatchedWritablePowerPortTemplate) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *PatchedWritablePowerPortTemplate) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPortTemplate) GetAllocatedDraw() int32 {
	if o == nil || o.AllocatedDraw.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPortTemplate) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *PatchedWritablePowerPortTemplate) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}
// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *PatchedWritablePowerPortTemplate) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *PatchedWritablePowerPortTemplate) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritablePowerPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerPortTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritablePowerPortTemplate) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritablePowerPortTemplate) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritablePowerPortTemplate) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritablePowerPortTemplate) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritablePowerPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
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
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
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

type NullablePatchedWritablePowerPortTemplate struct {
	value *PatchedWritablePowerPortTemplate
	isSet bool
}

func (v NullablePatchedWritablePowerPortTemplate) Get() *PatchedWritablePowerPortTemplate {
	return v.value
}

func (v *NullablePatchedWritablePowerPortTemplate) Set(val *PatchedWritablePowerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPortTemplate(val *PatchedWritablePowerPortTemplate) *NullablePatchedWritablePowerPortTemplate {
	return &NullablePatchedWritablePowerPortTemplate{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


