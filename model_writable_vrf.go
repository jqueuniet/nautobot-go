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

// WritableVRF Extends ModelSerializer to render any CustomFields and their values associated with an object.
type WritableVRF struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd NullableString `json:"rd,omitempty"`
	Tenant NullableString `json:"tenant,omitempty"`
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique *bool `json:"enforce_unique,omitempty"`
	Description *string `json:"description,omitempty"`
	ImportTargets []string `json:"import_targets,omitempty"`
	ExportTargets []string `json:"export_targets,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	IpaddressCount int32 `json:"ipaddress_count"`
	PrefixCount int32 `json:"prefix_count"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableVRF instantiates a new WritableVRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVRF(id string, url string, name string, created string, lastUpdated time.Time, ipaddressCount int32, prefixCount int32, computedFields map[string]interface{}, display string) *WritableVRF {
	this := WritableVRF{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.IpaddressCount = ipaddressCount
	this.PrefixCount = prefixCount
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewWritableVRFWithDefaults instantiates a new WritableVRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVRFWithDefaults() *WritableVRF {
	this := WritableVRF{}
	return &this
}

// GetId returns the Id field value
func (o *WritableVRF) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableVRF) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableVRF) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableVRF) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *WritableVRF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVRF) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVRF) GetRd() string {
	if o == nil || o.Rd.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *WritableVRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *WritableVRF) SetRd(v string) {
	o.Rd.Set(&v)
}
// SetRdNil sets the value for Rd to be an explicit nil
func (o *WritableVRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *WritableVRF) UnsetRd() {
	o.Rd.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVRF) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVRF) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableVRF) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *WritableVRF) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableVRF) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableVRF) UnsetTenant() {
	o.Tenant.Unset()
}

// GetEnforceUnique returns the EnforceUnique field value if set, zero value otherwise.
func (o *WritableVRF) GetEnforceUnique() bool {
	if o == nil || o.EnforceUnique == nil {
		var ret bool
		return ret
	}
	return *o.EnforceUnique
}

// GetEnforceUniqueOk returns a tuple with the EnforceUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetEnforceUniqueOk() (*bool, bool) {
	if o == nil || o.EnforceUnique == nil {
		return nil, false
	}
	return o.EnforceUnique, true
}

// HasEnforceUnique returns a boolean if a field has been set.
func (o *WritableVRF) HasEnforceUnique() bool {
	if o != nil && o.EnforceUnique != nil {
		return true
	}

	return false
}

// SetEnforceUnique gets a reference to the given bool and assigns it to the EnforceUnique field.
func (o *WritableVRF) SetEnforceUnique(v bool) {
	o.EnforceUnique = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableVRF) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableVRF) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableVRF) SetDescription(v string) {
	o.Description = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *WritableVRF) GetImportTargets() []string {
	if o == nil || o.ImportTargets == nil {
		var ret []string
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetImportTargetsOk() ([]string, bool) {
	if o == nil || o.ImportTargets == nil {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *WritableVRF) HasImportTargets() bool {
	if o != nil && o.ImportTargets != nil {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []string and assigns it to the ImportTargets field.
func (o *WritableVRF) SetImportTargets(v []string) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *WritableVRF) GetExportTargets() []string {
	if o == nil || o.ExportTargets == nil {
		var ret []string
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetExportTargetsOk() ([]string, bool) {
	if o == nil || o.ExportTargets == nil {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *WritableVRF) HasExportTargets() bool {
	if o != nil && o.ExportTargets != nil {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []string and assigns it to the ExportTargets field.
func (o *WritableVRF) SetExportTargets(v []string) {
	o.ExportTargets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVRF) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVRF) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritableVRF) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVRF) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVRF) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVRF) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *WritableVRF) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *WritableVRF) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *WritableVRF) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *WritableVRF) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetIpaddressCount returns the IpaddressCount field value
func (o *WritableVRF) GetIpaddressCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetIpaddressCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpaddressCount, true
}

// SetIpaddressCount sets field value
func (o *WritableVRF) SetIpaddressCount(v int32) {
	o.IpaddressCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *WritableVRF) GetPrefixCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetPrefixCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *WritableVRF) SetPrefixCount(v int32) {
	o.PrefixCount = v
}

// GetComputedFields returns the ComputedFields field value
func (o *WritableVRF) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *WritableVRF) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *WritableVRF) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableVRF) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableVRF) SetDisplay(v string) {
	o.Display = v
}

func (o WritableVRF) MarshalJSON() ([]byte, error) {
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
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.EnforceUnique != nil {
		toSerialize["enforce_unique"] = o.EnforceUnique
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImportTargets != nil {
		toSerialize["import_targets"] = o.ImportTargets
	}
	if o.ExportTargets != nil {
		toSerialize["export_targets"] = o.ExportTargets
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
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if true {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if true {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritableVRF struct {
	value *WritableVRF
	isSet bool
}

func (v NullableWritableVRF) Get() *WritableVRF {
	return v.value
}

func (v *NullableWritableVRF) Set(val *WritableVRF) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVRF) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVRF(val *WritableVRF) *NullableWritableVRF {
	return &NullableWritableVRF{value: val, isSet: true}
}

func (v NullableWritableVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


