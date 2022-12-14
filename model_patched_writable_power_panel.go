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

// PatchedWritablePowerPanel Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritablePowerPanel struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Site *string `json:"site,omitempty"`
	RackGroup NullableString `json:"rack_group,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	PowerfeedCount *int32 `json:"powerfeed_count,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritablePowerPanel instantiates a new PatchedWritablePowerPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerPanel() *PatchedWritablePowerPanel {
	this := PatchedWritablePowerPanel{}
	return &this
}

// NewPatchedWritablePowerPanelWithDefaults instantiates a new PatchedWritablePowerPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerPanelWithDefaults() *PatchedWritablePowerPanel {
	this := PatchedWritablePowerPanel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritablePowerPanel) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritablePowerPanel) SetUrl(v string) {
	o.Url = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *PatchedWritablePowerPanel) SetSite(v string) {
	o.Site = &v
}

// GetRackGroup returns the RackGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerPanel) GetRackGroup() string {
	if o == nil || o.RackGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.RackGroup.Get()
}

// GetRackGroupOk returns a tuple with the RackGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerPanel) GetRackGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackGroup.Get(), o.RackGroup.IsSet()
}

// HasRackGroup returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasRackGroup() bool {
	if o != nil && o.RackGroup.IsSet() {
		return true
	}

	return false
}

// SetRackGroup gets a reference to the given NullableString and assigns it to the RackGroup field.
func (o *PatchedWritablePowerPanel) SetRackGroup(v string) {
	o.RackGroup.Set(&v)
}
// SetRackGroupNil sets the value for RackGroup to be an explicit nil
func (o *PatchedWritablePowerPanel) SetRackGroupNil() {
	o.RackGroup.Set(nil)
}

// UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
func (o *PatchedWritablePowerPanel) UnsetRackGroup() {
	o.RackGroup.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerPanel) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritablePowerPanel) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerPanel) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetPowerfeedCount returns the PowerfeedCount field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetPowerfeedCount() int32 {
	if o == nil || o.PowerfeedCount == nil {
		var ret int32
		return ret
	}
	return *o.PowerfeedCount
}

// GetPowerfeedCountOk returns a tuple with the PowerfeedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetPowerfeedCountOk() (*int32, bool) {
	if o == nil || o.PowerfeedCount == nil {
		return nil, false
	}
	return o.PowerfeedCount, true
}

// HasPowerfeedCount returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasPowerfeedCount() bool {
	if o != nil && o.PowerfeedCount != nil {
		return true
	}

	return false
}

// SetPowerfeedCount gets a reference to the given int32 and assigns it to the PowerfeedCount field.
func (o *PatchedWritablePowerPanel) SetPowerfeedCount(v int32) {
	o.PowerfeedCount = &v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritablePowerPanel) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritablePowerPanel) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerPanel) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritablePowerPanel) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritablePowerPanel) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritablePowerPanel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.RackGroup.IsSet() {
		toSerialize["rack_group"] = o.RackGroup.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.PowerfeedCount != nil {
		toSerialize["powerfeed_count"] = o.PowerfeedCount
	}
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritablePowerPanel struct {
	value *PatchedWritablePowerPanel
	isSet bool
}

func (v NullablePatchedWritablePowerPanel) Get() *PatchedWritablePowerPanel {
	return v.value
}

func (v *NullablePatchedWritablePowerPanel) Set(val *PatchedWritablePowerPanel) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPanel) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPanel(val *PatchedWritablePowerPanel) *NullablePatchedWritablePowerPanel {
	return &NullablePatchedWritablePowerPanel{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


