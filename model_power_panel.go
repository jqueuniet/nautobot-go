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

// PowerPanel Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PowerPanel struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Site NestedSite `json:"site"`
	RackGroup NullablePowerPanelRackGroup `json:"rack_group,omitempty"`
	Name string `json:"name"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	PowerfeedCount int32 `json:"powerfeed_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewPowerPanel instantiates a new PowerPanel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPanel(id string, url string, site NestedSite, name string, powerfeedCount int32, display string) *PowerPanel {
	this := PowerPanel{}
	this.Id = id
	this.Url = url
	this.Site = site
	this.Name = name
	this.PowerfeedCount = powerfeedCount
	this.Display = display
	return &this
}

// NewPowerPanelWithDefaults instantiates a new PowerPanel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPanelWithDefaults() *PowerPanel {
	this := PowerPanel{}
	return &this
}

// GetId returns the Id field value
func (o *PowerPanel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PowerPanel) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PowerPanel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PowerPanel) SetUrl(v string) {
	o.Url = v
}

// GetSite returns the Site field value
func (o *PowerPanel) GetSite() NestedSite {
	if o == nil {
		var ret NestedSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetSiteOk() (*NestedSite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *PowerPanel) SetSite(v NestedSite) {
	o.Site = v
}

// GetRackGroup returns the RackGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerPanel) GetRackGroup() PowerPanelRackGroup {
	if o == nil || o.RackGroup.Get() == nil {
		var ret PowerPanelRackGroup
		return ret
	}
	return *o.RackGroup.Get()
}

// GetRackGroupOk returns a tuple with the RackGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerPanel) GetRackGroupOk() (*PowerPanelRackGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RackGroup.Get(), o.RackGroup.IsSet()
}

// HasRackGroup returns a boolean if a field has been set.
func (o *PowerPanel) HasRackGroup() bool {
	if o != nil && o.RackGroup.IsSet() {
		return true
	}

	return false
}

// SetRackGroup gets a reference to the given NullablePowerPanelRackGroup and assigns it to the RackGroup field.
func (o *PowerPanel) SetRackGroup(v PowerPanelRackGroup) {
	o.RackGroup.Set(&v)
}
// SetRackGroupNil sets the value for RackGroup to be an explicit nil
func (o *PowerPanel) SetRackGroupNil() {
	o.RackGroup.Set(nil)
}

// UnsetRackGroup ensures that no value is present for RackGroup, not even an explicit nil
func (o *PowerPanel) UnsetRackGroup() {
	o.RackGroup.Unset()
}

// GetName returns the Name field value
func (o *PowerPanel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerPanel) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerPanel) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerPanel) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PowerPanel) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerPanel) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerPanel) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerPanel) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetPowerfeedCount returns the PowerfeedCount field value
func (o *PowerPanel) GetPowerfeedCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PowerfeedCount
}

// GetPowerfeedCountOk returns a tuple with the PowerfeedCount field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetPowerfeedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerfeedCount, true
}

// SetPowerfeedCount sets field value
func (o *PowerPanel) SetPowerfeedCount(v int32) {
	o.PowerfeedCount = v
}

// GetDisplay returns the Display field value
func (o *PowerPanel) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *PowerPanel) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *PowerPanel) SetDisplay(v string) {
	o.Display = v
}

func (o PowerPanel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if o.RackGroup.IsSet() {
		toSerialize["rack_group"] = o.RackGroup.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["powerfeed_count"] = o.PowerfeedCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePowerPanel struct {
	value *PowerPanel
	isSet bool
}

func (v NullablePowerPanel) Get() *PowerPanel {
	return v.value
}

func (v *NullablePowerPanel) Set(val *PowerPanel) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPanel) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPanel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPanel(val *PowerPanel) *NullablePowerPanel {
	return &NullablePowerPanel{value: val, isSet: true}
}

func (v NullablePowerPanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPanel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


