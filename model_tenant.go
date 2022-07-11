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

// Tenant Extends ModelSerializer to render any CustomFields and their values associated with an object.
type Tenant struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	Group *NestedTenantGroup `json:"group,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	CircuitCount int32 `json:"circuit_count"`
	DeviceCount int32 `json:"device_count"`
	IpaddressCount int32 `json:"ipaddress_count"`
	PrefixCount int32 `json:"prefix_count"`
	RackCount int32 `json:"rack_count"`
	SiteCount int32 `json:"site_count"`
	VirtualmachineCount int32 `json:"virtualmachine_count"`
	VlanCount int32 `json:"vlan_count"`
	VrfCount int32 `json:"vrf_count"`
	ClusterCount int32 `json:"cluster_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(id string, url string, name string, created string, lastUpdated time.Time, circuitCount int32, deviceCount int32, ipaddressCount int32, prefixCount int32, rackCount int32, siteCount int32, virtualmachineCount int32, vlanCount int32, vrfCount int32, clusterCount int32, display string) *Tenant {
	this := Tenant{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CircuitCount = circuitCount
	this.DeviceCount = deviceCount
	this.IpaddressCount = ipaddressCount
	this.PrefixCount = prefixCount
	this.RackCount = rackCount
	this.SiteCount = siteCount
	this.VirtualmachineCount = virtualmachineCount
	this.VlanCount = vlanCount
	this.VrfCount = vrfCount
	this.ClusterCount = clusterCount
	this.Display = display
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value
func (o *Tenant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tenant) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Tenant) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tenant) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Tenant) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Tenant) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Tenant) SetSlug(v string) {
	o.Slug = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Tenant) GetGroup() NestedTenantGroup {
	if o == nil || o.Group == nil {
		var ret NestedTenantGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetGroupOk() (*NestedTenantGroup, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Tenant) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NestedTenantGroup and assigns it to the Group field.
func (o *Tenant) SetGroup(v NestedTenantGroup) {
	o.Group = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tenant) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Tenant) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Tenant) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Tenant) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Tenant) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Tenant) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *Tenant) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Tenant) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Tenant) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Tenant) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *Tenant) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Tenant) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Tenant) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Tenant) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetCircuitCount returns the CircuitCount field value
func (o *Tenant) GetCircuitCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetCircuitCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitCount, true
}

// SetCircuitCount sets field value
func (o *Tenant) SetCircuitCount(v int32) {
	o.CircuitCount = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *Tenant) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Tenant) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetIpaddressCount returns the IpaddressCount field value
func (o *Tenant) GetIpaddressCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetIpaddressCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpaddressCount, true
}

// SetIpaddressCount sets field value
func (o *Tenant) SetIpaddressCount(v int32) {
	o.IpaddressCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *Tenant) GetPrefixCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetPrefixCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *Tenant) SetPrefixCount(v int32) {
	o.PrefixCount = v
}

// GetRackCount returns the RackCount field value
func (o *Tenant) GetRackCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetRackCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *Tenant) SetRackCount(v int32) {
	o.RackCount = v
}

// GetSiteCount returns the SiteCount field value
func (o *Tenant) GetSiteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetSiteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteCount, true
}

// SetSiteCount sets field value
func (o *Tenant) SetSiteCount(v int32) {
	o.SiteCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Tenant) GetVirtualmachineCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Tenant) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = v
}

// GetVlanCount returns the VlanCount field value
func (o *Tenant) GetVlanCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVlanCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *Tenant) SetVlanCount(v int32) {
	o.VlanCount = v
}

// GetVrfCount returns the VrfCount field value
func (o *Tenant) GetVrfCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VrfCount
}

// GetVrfCountOk returns a tuple with the VrfCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetVrfCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrfCount, true
}

// SetVrfCount sets field value
func (o *Tenant) SetVrfCount(v int32) {
	o.VrfCount = v
}

// GetClusterCount returns the ClusterCount field value
func (o *Tenant) GetClusterCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetClusterCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterCount, true
}

// SetClusterCount sets field value
func (o *Tenant) SetClusterCount(v int32) {
	o.ClusterCount = v
}

// GetDisplay returns the Display field value
func (o *Tenant) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Tenant) SetDisplay(v string) {
	o.Display = v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
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
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
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
		toSerialize["circuit_count"] = o.CircuitCount
	}
	if true {
		toSerialize["device_count"] = o.DeviceCount
	}
	if true {
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if true {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if true {
		toSerialize["rack_count"] = o.RackCount
	}
	if true {
		toSerialize["site_count"] = o.SiteCount
	}
	if true {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if true {
		toSerialize["vlan_count"] = o.VlanCount
	}
	if true {
		toSerialize["vrf_count"] = o.VrfCount
	}
	if true {
		toSerialize["cluster_count"] = o.ClusterCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


