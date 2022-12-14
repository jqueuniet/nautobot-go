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

// Prefix Mixin to add `status` choice field to model serializers.
type Prefix struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Family AggregateFamily `json:"family"`
	Prefix string `json:"prefix"`
	Site NullableCircuitTerminationSite `json:"site,omitempty"`
	Vrf NullableIPAddressVrf `json:"vrf,omitempty"`
	Tenant NullableAggregateTenant `json:"tenant,omitempty"`
	Vlan NullableInterfaceUntaggedVlan `json:"vlan,omitempty"`
	Status PrefixStatus `json:"status"`
	Role NullablePrefixRole `json:"role,omitempty"`
	// All IP addresses within this prefix are considered usable
	IsPool *bool `json:"is_pool,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewPrefix instantiates a new Prefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrefix(id string, url string, family AggregateFamily, prefix string, status PrefixStatus, created string, lastUpdated time.Time, display string) *Prefix {
	this := Prefix{}
	this.Id = id
	this.Url = url
	this.Family = family
	this.Prefix = prefix
	this.Status = status
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Display = display
	return &this
}

// NewPrefixWithDefaults instantiates a new Prefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrefixWithDefaults() *Prefix {
	this := Prefix{}
	return &this
}

// GetId returns the Id field value
func (o *Prefix) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Prefix) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Prefix) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Prefix) SetUrl(v string) {
	o.Url = v
}

// GetFamily returns the Family field value
func (o *Prefix) GetFamily() AggregateFamily {
	if o == nil {
		var ret AggregateFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetFamilyOk() (*AggregateFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *Prefix) SetFamily(v AggregateFamily) {
	o.Family = v
}

// GetPrefix returns the Prefix field value
func (o *Prefix) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *Prefix) SetPrefix(v string) {
	o.Prefix = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetSite() CircuitTerminationSite {
	if o == nil || o.Site.Get() == nil {
		var ret CircuitTerminationSite
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetSiteOk() (*CircuitTerminationSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *Prefix) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableCircuitTerminationSite and assigns it to the Site field.
func (o *Prefix) SetSite(v CircuitTerminationSite) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *Prefix) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *Prefix) UnsetSite() {
	o.Site.Unset()
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetVrf() IPAddressVrf {
	if o == nil || o.Vrf.Get() == nil {
		var ret IPAddressVrf
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetVrfOk() (*IPAddressVrf, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *Prefix) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableIPAddressVrf and assigns it to the Vrf field.
func (o *Prefix) SetVrf(v IPAddressVrf) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *Prefix) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *Prefix) UnsetVrf() {
	o.Vrf.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetTenant() AggregateTenant {
	if o == nil || o.Tenant.Get() == nil {
		var ret AggregateTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetTenantOk() (*AggregateTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Prefix) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableAggregateTenant and assigns it to the Tenant field.
func (o *Prefix) SetTenant(v AggregateTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Prefix) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Prefix) UnsetTenant() {
	o.Tenant.Unset()
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetVlan() InterfaceUntaggedVlan {
	if o == nil || o.Vlan.Get() == nil {
		var ret InterfaceUntaggedVlan
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetVlanOk() (*InterfaceUntaggedVlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *Prefix) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInterfaceUntaggedVlan and assigns it to the Vlan field.
func (o *Prefix) SetVlan(v InterfaceUntaggedVlan) {
	o.Vlan.Set(&v)
}
// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *Prefix) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *Prefix) UnsetVlan() {
	o.Vlan.Unset()
}

// GetStatus returns the Status field value
func (o *Prefix) GetStatus() PrefixStatus {
	if o == nil {
		var ret PrefixStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetStatusOk() (*PrefixStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Prefix) SetStatus(v PrefixStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prefix) GetRole() PrefixRole {
	if o == nil || o.Role.Get() == nil {
		var ret PrefixRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prefix) GetRoleOk() (*PrefixRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *Prefix) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullablePrefixRole and assigns it to the Role field.
func (o *Prefix) SetRole(v PrefixRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *Prefix) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *Prefix) UnsetRole() {
	o.Role.Unset()
}

// GetIsPool returns the IsPool field value if set, zero value otherwise.
func (o *Prefix) GetIsPool() bool {
	if o == nil || o.IsPool == nil {
		var ret bool
		return ret
	}
	return *o.IsPool
}

// GetIsPoolOk returns a tuple with the IsPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetIsPoolOk() (*bool, bool) {
	if o == nil || o.IsPool == nil {
		return nil, false
	}
	return o.IsPool, true
}

// HasIsPool returns a boolean if a field has been set.
func (o *Prefix) HasIsPool() bool {
	if o != nil && o.IsPool != nil {
		return true
	}

	return false
}

// SetIsPool gets a reference to the given bool and assigns it to the IsPool field.
func (o *Prefix) SetIsPool(v bool) {
	o.IsPool = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Prefix) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Prefix) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Prefix) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Prefix) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Prefix) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *Prefix) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Prefix) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prefix) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Prefix) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Prefix) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *Prefix) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Prefix) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Prefix) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Prefix) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetDisplay returns the Display field value
func (o *Prefix) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Prefix) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Prefix) SetDisplay(v string) {
	o.Display = v
}

func (o Prefix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.IsPool != nil {
		toSerialize["is_pool"] = o.IsPool
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
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePrefix struct {
	value *Prefix
	isSet bool
}

func (v NullablePrefix) Get() *Prefix {
	return v.value
}

func (v *NullablePrefix) Set(val *Prefix) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefix) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefix(val *Prefix) *NullablePrefix {
	return &NullablePrefix{value: val, isSet: true}
}

func (v NullablePrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


