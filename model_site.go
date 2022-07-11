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

// Site Mixin to add `status` choice field to model serializers.
type Site struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	Status SiteStatus `json:"status"`
	Region NullableRegionParent `json:"region,omitempty"`
	Tenant NullableAggregateTenant `json:"tenant,omitempty"`
	// Local facility ID or description
	Facility *string `json:"facility,omitempty"`
	// 32-bit autonomous system number
	Asn NullableInt64 `json:"asn,omitempty"`
	TimeZone NullableString `json:"time_zone,omitempty"`
	Description *string `json:"description,omitempty"`
	PhysicalAddress *string `json:"physical_address,omitempty"`
	ShippingAddress *string `json:"shipping_address,omitempty"`
	// GPS coordinate (latitude)
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// GPS coordinate (longitude)
	Longitude NullableFloat64 `json:"longitude,omitempty"`
	ContactName *string `json:"contact_name,omitempty"`
	ContactPhone *string `json:"contact_phone,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	CircuitCount int32 `json:"circuit_count"`
	DeviceCount int32 `json:"device_count"`
	PrefixCount int32 `json:"prefix_count"`
	RackCount int32 `json:"rack_count"`
	VirtualmachineCount int32 `json:"virtualmachine_count"`
	VlanCount int32 `json:"vlan_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewSite instantiates a new Site object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSite(id string, url string, name string, status SiteStatus, created string, lastUpdated time.Time, circuitCount int32, deviceCount int32, prefixCount int32, rackCount int32, virtualmachineCount int32, vlanCount int32, display string) *Site {
	this := Site{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Status = status
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CircuitCount = circuitCount
	this.DeviceCount = deviceCount
	this.PrefixCount = prefixCount
	this.RackCount = rackCount
	this.VirtualmachineCount = virtualmachineCount
	this.VlanCount = vlanCount
	this.Display = display
	return &this
}

// NewSiteWithDefaults instantiates a new Site object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWithDefaults() *Site {
	this := Site{}
	return &this
}

// GetId returns the Id field value
func (o *Site) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Site) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Site) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Site) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Site) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Site) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *Site) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Site) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Site) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *Site) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *Site) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *Site) SetSlug(v string) {
	o.Slug = &v
}

// GetStatus returns the Status field value
func (o *Site) GetStatus() SiteStatus {
	if o == nil {
		var ret SiteStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Site) GetStatusOk() (*SiteStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Site) SetStatus(v SiteStatus) {
	o.Status = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetRegion() RegionParent {
	if o == nil || o.Region.Get() == nil {
		var ret RegionParent
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetRegionOk() (*RegionParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *Site) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableRegionParent and assigns it to the Region field.
func (o *Site) SetRegion(v RegionParent) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *Site) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *Site) UnsetRegion() {
	o.Region.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetTenant() AggregateTenant {
	if o == nil || o.Tenant.Get() == nil {
		var ret AggregateTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetTenantOk() (*AggregateTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Site) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableAggregateTenant and assigns it to the Tenant field.
func (o *Site) SetTenant(v AggregateTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Site) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Site) UnsetTenant() {
	o.Tenant.Unset()
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *Site) GetFacility() string {
	if o == nil || o.Facility == nil {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetFacilityOk() (*string, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *Site) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *Site) SetFacility(v string) {
	o.Facility = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetAsn() int64 {
	if o == nil || o.Asn.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// HasAsn returns a boolean if a field has been set.
func (o *Site) HasAsn() bool {
	if o != nil && o.Asn.IsSet() {
		return true
	}

	return false
}

// SetAsn gets a reference to the given NullableInt64 and assigns it to the Asn field.
func (o *Site) SetAsn(v int64) {
	o.Asn.Set(&v)
}
// SetAsnNil sets the value for Asn to be an explicit nil
func (o *Site) SetAsnNil() {
	o.Asn.Set(nil)
}

// UnsetAsn ensures that no value is present for Asn, not even an explicit nil
func (o *Site) UnsetAsn() {
	o.Asn.Unset()
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetTimeZone() string {
	if o == nil || o.TimeZone.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Site) HasTimeZone() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *Site) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *Site) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *Site) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Site) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Site) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Site) SetDescription(v string) {
	o.Description = &v
}

// GetPhysicalAddress returns the PhysicalAddress field value if set, zero value otherwise.
func (o *Site) GetPhysicalAddress() string {
	if o == nil || o.PhysicalAddress == nil {
		var ret string
		return ret
	}
	return *o.PhysicalAddress
}

// GetPhysicalAddressOk returns a tuple with the PhysicalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetPhysicalAddressOk() (*string, bool) {
	if o == nil || o.PhysicalAddress == nil {
		return nil, false
	}
	return o.PhysicalAddress, true
}

// HasPhysicalAddress returns a boolean if a field has been set.
func (o *Site) HasPhysicalAddress() bool {
	if o != nil && o.PhysicalAddress != nil {
		return true
	}

	return false
}

// SetPhysicalAddress gets a reference to the given string and assigns it to the PhysicalAddress field.
func (o *Site) SetPhysicalAddress(v string) {
	o.PhysicalAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *Site) GetShippingAddress() string {
	if o == nil || o.ShippingAddress == nil {
		var ret string
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetShippingAddressOk() (*string, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *Site) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given string and assigns it to the ShippingAddress field.
func (o *Site) SetShippingAddress(v string) {
	o.ShippingAddress = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLatitude() float64 {
	if o == nil || o.Latitude.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *Site) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *Site) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *Site) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *Site) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Site) GetLongitude() float64 {
	if o == nil || o.Longitude.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Site) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *Site) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *Site) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *Site) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *Site) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *Site) GetContactName() string {
	if o == nil || o.ContactName == nil {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactNameOk() (*string, bool) {
	if o == nil || o.ContactName == nil {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *Site) HasContactName() bool {
	if o != nil && o.ContactName != nil {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *Site) SetContactName(v string) {
	o.ContactName = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *Site) GetContactPhone() string {
	if o == nil || o.ContactPhone == nil {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactPhoneOk() (*string, bool) {
	if o == nil || o.ContactPhone == nil {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *Site) HasContactPhone() bool {
	if o != nil && o.ContactPhone != nil {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *Site) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Site) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Site) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Site) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Site) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Site) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Site) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Site) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Site) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *Site) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Site) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Site) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Site) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Site) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *Site) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Site) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Site) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Site) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Site) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Site) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetCircuitCount returns the CircuitCount field value
func (o *Site) GetCircuitCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetCircuitCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitCount, true
}

// SetCircuitCount sets field value
func (o *Site) SetCircuitCount(v int32) {
	o.CircuitCount = v
}

// GetDeviceCount returns the DeviceCount field value
func (o *Site) GetDeviceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetDeviceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCount, true
}

// SetDeviceCount sets field value
func (o *Site) SetDeviceCount(v int32) {
	o.DeviceCount = v
}

// GetPrefixCount returns the PrefixCount field value
func (o *Site) GetPrefixCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetPrefixCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixCount, true
}

// SetPrefixCount sets field value
func (o *Site) SetPrefixCount(v int32) {
	o.PrefixCount = v
}

// GetRackCount returns the RackCount field value
func (o *Site) GetRackCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetRackCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *Site) SetRackCount(v int32) {
	o.RackCount = v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value
func (o *Site) GetVirtualmachineCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualmachineCount, true
}

// SetVirtualmachineCount sets field value
func (o *Site) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = v
}

// GetVlanCount returns the VlanCount field value
func (o *Site) GetVlanCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *Site) GetVlanCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *Site) SetVlanCount(v int32) {
	o.VlanCount = v
}

// GetDisplay returns the Display field value
func (o *Site) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Site) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Site) SetDisplay(v string) {
	o.Display = v
}

func (o Site) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.Asn.IsSet() {
		toSerialize["asn"] = o.Asn.Get()
	}
	if o.TimeZone.IsSet() {
		toSerialize["time_zone"] = o.TimeZone.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PhysicalAddress != nil {
		toSerialize["physical_address"] = o.PhysicalAddress
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.ContactName != nil {
		toSerialize["contact_name"] = o.ContactName
	}
	if o.ContactPhone != nil {
		toSerialize["contact_phone"] = o.ContactPhone
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
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
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if true {
		toSerialize["rack_count"] = o.RackCount
	}
	if true {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if true {
		toSerialize["vlan_count"] = o.VlanCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableSite struct {
	value *Site
	isSet bool
}

func (v NullableSite) Get() *Site {
	return v.value
}

func (v *NullableSite) Set(val *Site) {
	v.value = val
	v.isSet = true
}

func (v NullableSite) IsSet() bool {
	return v.isSet
}

func (v *NullableSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSite(val *Site) *NullableSite {
	return &NullableSite{value: val, isSet: true}
}

func (v NullableSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


