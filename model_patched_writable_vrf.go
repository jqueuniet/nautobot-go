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

// PatchedWritableVRF Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritableVRF struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
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
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableVRF instantiates a new PatchedWritableVRF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableVRF() *PatchedWritableVRF {
	this := PatchedWritableVRF{}
	return &this
}

// NewPatchedWritableVRFWithDefaults instantiates a new PatchedWritableVRF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableVRFWithDefaults() *PatchedWritableVRF {
	this := PatchedWritableVRF{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableVRF) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableVRF) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableVRF) SetName(v string) {
	o.Name = &v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVRF) GetRd() string {
	if o == nil || o.Rd.Get() == nil {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVRF) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *PatchedWritableVRF) SetRd(v string) {
	o.Rd.Set(&v)
}
// SetRdNil sets the value for Rd to be an explicit nil
func (o *PatchedWritableVRF) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *PatchedWritableVRF) UnsetRd() {
	o.Rd.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableVRF) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableVRF) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *PatchedWritableVRF) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableVRF) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableVRF) UnsetTenant() {
	o.Tenant.Unset()
}

// GetEnforceUnique returns the EnforceUnique field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetEnforceUnique() bool {
	if o == nil || o.EnforceUnique == nil {
		var ret bool
		return ret
	}
	return *o.EnforceUnique
}

// GetEnforceUniqueOk returns a tuple with the EnforceUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetEnforceUniqueOk() (*bool, bool) {
	if o == nil || o.EnforceUnique == nil {
		return nil, false
	}
	return o.EnforceUnique, true
}

// HasEnforceUnique returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasEnforceUnique() bool {
	if o != nil && o.EnforceUnique != nil {
		return true
	}

	return false
}

// SetEnforceUnique gets a reference to the given bool and assigns it to the EnforceUnique field.
func (o *PatchedWritableVRF) SetEnforceUnique(v bool) {
	o.EnforceUnique = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableVRF) SetDescription(v string) {
	o.Description = &v
}

// GetImportTargets returns the ImportTargets field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetImportTargets() []string {
	if o == nil || o.ImportTargets == nil {
		var ret []string
		return ret
	}
	return o.ImportTargets
}

// GetImportTargetsOk returns a tuple with the ImportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetImportTargetsOk() ([]string, bool) {
	if o == nil || o.ImportTargets == nil {
		return nil, false
	}
	return o.ImportTargets, true
}

// HasImportTargets returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasImportTargets() bool {
	if o != nil && o.ImportTargets != nil {
		return true
	}

	return false
}

// SetImportTargets gets a reference to the given []string and assigns it to the ImportTargets field.
func (o *PatchedWritableVRF) SetImportTargets(v []string) {
	o.ImportTargets = v
}

// GetExportTargets returns the ExportTargets field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetExportTargets() []string {
	if o == nil || o.ExportTargets == nil {
		var ret []string
		return ret
	}
	return o.ExportTargets
}

// GetExportTargetsOk returns a tuple with the ExportTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetExportTargetsOk() ([]string, bool) {
	if o == nil || o.ExportTargets == nil {
		return nil, false
	}
	return o.ExportTargets, true
}

// HasExportTargets returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasExportTargets() bool {
	if o != nil && o.ExportTargets != nil {
		return true
	}

	return false
}

// SetExportTargets gets a reference to the given []string and assigns it to the ExportTargets field.
func (o *PatchedWritableVRF) SetExportTargets(v []string) {
	o.ExportTargets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableVRF) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableVRF) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PatchedWritableVRF) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PatchedWritableVRF) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetIpaddressCount returns the IpaddressCount field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetIpaddressCount() int32 {
	if o == nil || o.IpaddressCount == nil {
		var ret int32
		return ret
	}
	return *o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetIpaddressCountOk() (*int32, bool) {
	if o == nil || o.IpaddressCount == nil {
		return nil, false
	}
	return o.IpaddressCount, true
}

// HasIpaddressCount returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasIpaddressCount() bool {
	if o != nil && o.IpaddressCount != nil {
		return true
	}

	return false
}

// SetIpaddressCount gets a reference to the given int32 and assigns it to the IpaddressCount field.
func (o *PatchedWritableVRF) SetIpaddressCount(v int32) {
	o.IpaddressCount = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetPrefixCount() int32 {
	if o == nil || o.PrefixCount == nil {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetPrefixCountOk() (*int32, bool) {
	if o == nil || o.PrefixCount == nil {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasPrefixCount() bool {
	if o != nil && o.PrefixCount != nil {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *PatchedWritableVRF) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableVRF) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableVRF) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableVRF) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableVRF) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableVRF) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableVRF) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Name != nil {
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
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.IpaddressCount != nil {
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if o.PrefixCount != nil {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableVRF struct {
	value *PatchedWritableVRF
	isSet bool
}

func (v NullablePatchedWritableVRF) Get() *PatchedWritableVRF {
	return v.value
}

func (v *NullablePatchedWritableVRF) Set(val *PatchedWritableVRF) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableVRF) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableVRF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableVRF(val *PatchedWritableVRF) *NullablePatchedWritableVRF {
	return &NullablePatchedWritableVRF{value: val, isSet: true}
}

func (v NullablePatchedWritableVRF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableVRF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


