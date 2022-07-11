/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TenantGroup Extends ModelSerializer to render any CustomFields and their values associated with an object.
type TenantGroup struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	Parent NullableTenantGroupParent `json:"parent,omitempty"`
	Description *string `json:"description,omitempty"`
	TenantCount int32 `json:"tenant_count"`
	Depth int32 `json:"_depth"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewTenantGroup instantiates a new TenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantGroup(id string, url string, name string, tenantCount int32, depth int32, created string, lastUpdated time.Time, display string) *TenantGroup {
	this := TenantGroup{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.TenantCount = tenantCount
	this.Depth = depth
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Display = display
	return &this
}

// NewTenantGroupWithDefaults instantiates a new TenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantGroupWithDefaults() *TenantGroup {
	this := TenantGroup{}
	return &this
}

// GetId returns the Id field value
func (o *TenantGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantGroup) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *TenantGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TenantGroup) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *TenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *TenantGroup) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *TenantGroup) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *TenantGroup) SetSlug(v string) {
	o.Slug = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantGroup) GetParent() TenantGroupParent {
	if o == nil || o.Parent.Get() == nil {
		var ret TenantGroupParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantGroup) GetParentOk() (*TenantGroupParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *TenantGroup) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableTenantGroupParent and assigns it to the Parent field.
func (o *TenantGroup) SetParent(v TenantGroupParent) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *TenantGroup) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *TenantGroup) UnsetParent() {
	o.Parent.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TenantGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TenantGroup) SetDescription(v string) {
	o.Description = &v
}

// GetTenantCount returns the TenantCount field value
func (o *TenantGroup) GetTenantCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantCount
}

// GetTenantCountOk returns a tuple with the TenantCount field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetTenantCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantCount, true
}

// SetTenantCount sets field value
func (o *TenantGroup) SetTenantCount(v int32) {
	o.TenantCount = v
}

// GetDepth returns the Depth field value
func (o *TenantGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *TenantGroup) SetDepth(v int32) {
	o.Depth = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TenantGroup) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TenantGroup) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TenantGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *TenantGroup) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TenantGroup) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *TenantGroup) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *TenantGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetDisplay returns the Display field value
func (o *TenantGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *TenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *TenantGroup) SetDisplay(v string) {
	o.Display = v
}

func (o TenantGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tenant_count"] = o.TenantCount
	}
	if true {
		toSerialize["_depth"] = o.Depth
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableTenantGroup struct {
	value *TenantGroup
	isSet bool
}

func (v NullableTenantGroup) Get() *TenantGroup {
	return v.value
}

func (v *NullableTenantGroup) Set(val *TenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantGroup(val *TenantGroup) *NullableTenantGroup {
	return &NullableTenantGroup{value: val, isSet: true}
}

func (v NullableTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


