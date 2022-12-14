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

// VRF Extends ModelSerializer to render any CustomFields and their values associated with an object.
type VRF struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd NullableString `json:"rd,omitempty"`
	Tenant NullableAggregateTenant `json:"tenant,omitempty"`
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique *bool `json:"enforce_unique,omitempty"`
	Description *string `json:"description,omitempty"`
	ImportTargets []VRFImportTargetsInner `json:"import_targets,omitempty"`
	ExportTargets []VRFImportTargetsInner `json:"export_targets,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	IpaddressCount int32 `json:"ipaddress_count"`
	PrefixCount int32 `json:"prefix_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewVRF instantiates a new VRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVRF(id string, url string, name string, created string, lastUpdated time.Time, ipaddressCount int32, prefixCount int32, display string) *VRF {
	this := VRF{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.IpaddressCount = ipaddressCount
	this.PrefixCount = prefixCount
	this.Display = display
	return &this
}

// NewVRFWithDefaults instantiates a new VRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVRFWithDefaults() *VRF {
	this := VRF{}
	return &this
}

// GetId returns the Id field value
func (o *VRF) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VRF) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VRF) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VRF) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VRF) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VRF) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *VRF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VRF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VRF) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetRd() string {
	if o == nil || o.Rd.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *VRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *VRF) SetRd(v string) {
	o.Rd.Set(&v)
}
// SetRdNil sets the value for Rd to be an explicit nil
func (o *VRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *VRF) UnsetRd() {
	o.Rd.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VRF) GetTenant() AggregateTenant {
	if o == nil || o.Tenant.Get() == nil {
		var ret AggregateTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VRF) GetTenantOk() (*AggregateTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VRF) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableAggregateTenant and assigns it to the Tenant field.
func (o *VRF) SetTenant(v AggregateTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VRF) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VRF) UnsetTenant() {
	o.Tenant.Unset()
}

// GetEnforceUnique returns the EnforceUnique field value if set, zero value otherwise.
func (o *VRF) GetEnforceUnique() bool {
	if o == nil || o.EnforceUnique == nil {
		var ret bool
		return ret
	}
	return *o.EnforceUnique
}

// GetEnforceUniqueOk returns a tuple with the EnforceUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetEnforceUniqueOk() (*bool, bool) {
	if o == nil || o.EnforceUnique == nil {
		return nil, false
	}
	return o.EnforceUnique, true
}

// HasEnforceUnique returns a boolean if a field has been set.
func (o *VRF) HasEnforceUnique() bool {
	if o != nil && o.EnforceUnique != nil {
		return true
	}

	return false
}

// SetEnforceUnique gets a reference to the given bool and assigns it to the EnforceUnique field.
func (o *VRF) SetEnforceUnique(v bool) {
	o.EnforceUnique = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VRF) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VRF) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VRF) SetDescription(v string) {
	o.Description = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *VRF) GetImportTargets() []VRFImportTargetsInner {
	if o == nil || o.ImportTargets == nil {
		var ret []VRFImportTargetsInner
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetImportTargetsOk() ([]VRFImportTargetsInner, bool) {
	if o == nil || o.ImportTargets == nil {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *VRF) HasImportTargets() bool {
	if o != nil && o.ImportTargets != nil {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []VRFImportTargetsInner and assigns it to the ImportTargets field.
func (o *VRF) SetImportTargets(v []VRFImportTargetsInner) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *VRF) GetExportTargets() []VRFImportTargetsInner {
	if o == nil || o.ExportTargets == nil {
		var ret []VRFImportTargetsInner
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetExportTargetsOk() ([]VRFImportTargetsInner, bool) {
	if o == nil || o.ExportTargets == nil {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *VRF) HasExportTargets() bool {
	if o != nil && o.ExportTargets != nil {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []VRFImportTargetsInner and assigns it to the ExportTargets field.
func (o *VRF) SetExportTargets(v []VRFImportTargetsInner) {
	o.ExportTargets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VRF) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VRF) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *VRF) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VRF) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VRF) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VRF) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VRF) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *VRF) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *VRF) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *VRF) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *VRF) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *VRF) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *VRF) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetIpaddressCount returns the IpaddressCount field value
func (o *VRF) GetIpaddressCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value
// and a boolean to check if the value has been set.
func (o *VRF) GetIpaddressCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpaddressCount, true
}

// SetIpaddressCount sets field value
func (o *VRF) SetIpaddressCount(v int32) {
	o.IpaddressCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *VRF) GetPrefixCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *VRF) GetPrefixCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *VRF) SetPrefixCount(v int32) {
	o.PrefixCount = v
}

// GetDisplay returns the Display field value
func (o *VRF) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VRF) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VRF) SetDisplay(v string) {
	o.Display = v
}

func (o VRF) MarshalJSON() ([]byte, error) {
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
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableVRF struct {
	value *VRF
	isSet bool
}

func (v NullableVRF) Get() *VRF {
	return v.value
}

func (v *NullableVRF) Set(val *VRF) {
	v.value = val
	v.isSet = true
}

func (v NullableVRF) IsSet() bool {
	return v.isSet
}

func (v *NullableVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVRF(val *VRF) *NullableVRF {
	return &NullableVRF{value: val, isSet: true}
}

func (v NullableVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


