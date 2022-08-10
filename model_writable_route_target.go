/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WritableRouteTarget Extends ModelSerializer to render any CustomFields and their values associated with an object.
type WritableRouteTarget struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Route target value (formatted in accordance with RFC 4360)
	Name string `json:"name"`
	Tenant NullableString `json:"tenant,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableRouteTarget instantiates a new WritableRouteTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableRouteTarget(id string, url string, name string, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string) *WritableRouteTarget {
	this := WritableRouteTarget{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewWritableRouteTargetWithDefaults instantiates a new WritableRouteTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableRouteTargetWithDefaults() *WritableRouteTarget {
	this := WritableRouteTarget{}
	return &this
}

// GetId returns the Id field value
func (o *WritableRouteTarget) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableRouteTarget) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableRouteTarget) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableRouteTarget) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *WritableRouteTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableRouteTarget) SetName(v string) {
	o.Name = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableRouteTarget) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableRouteTarget) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableRouteTarget) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *WritableRouteTarget) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableRouteTarget) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableRouteTarget) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableRouteTarget) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableRouteTarget) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableRouteTarget) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableRouteTarget) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableRouteTarget) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritableRouteTarget) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableRouteTarget) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableRouteTarget) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableRouteTarget) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *WritableRouteTarget) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *WritableRouteTarget) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *WritableRouteTarget) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *WritableRouteTarget) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetComputedFields returns the ComputedFields field value
func (o *WritableRouteTarget) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *WritableRouteTarget) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *WritableRouteTarget) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableRouteTarget) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableRouteTarget) SetDisplay(v string) {
	o.Display = v
}

func (o WritableRouteTarget) MarshalJSON() ([]byte, error) {
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
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritableRouteTarget struct {
	value *WritableRouteTarget
	isSet bool
}

func (v NullableWritableRouteTarget) Get() *WritableRouteTarget {
	return v.value
}

func (v *NullableWritableRouteTarget) Set(val *WritableRouteTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableRouteTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableRouteTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableRouteTarget(val *WritableRouteTarget) *NullableWritableRouteTarget {
	return &NullableWritableRouteTarget{value: val, isSet: true}
}

func (v NullableWritableRouteTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableRouteTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


