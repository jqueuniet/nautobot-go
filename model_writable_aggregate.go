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

// WritableAggregate Extends ModelSerializer to render any CustomFields and their values associated with an object.
type WritableAggregate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Family NullableFamilyEnum `json:"family"`
	Prefix string `json:"prefix"`
	Rir string `json:"rir"`
	Tenant NullableString `json:"tenant,omitempty"`
	DateAdded NullableString `json:"date_added,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableAggregate instantiates a new WritableAggregate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableAggregate(id string, url string, family NullableFamilyEnum, prefix string, rir string, created string, lastUpdated time.Time, computedFields map[string]interface{}, display string) *WritableAggregate {
	this := WritableAggregate{}
	this.Id = id
	this.Url = url
	this.Family = family
	this.Prefix = prefix
	this.Rir = rir
	this.Created = created
	this.LastUpdated = lastUpdated
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewWritableAggregateWithDefaults instantiates a new WritableAggregate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableAggregateWithDefaults() *WritableAggregate {
	this := WritableAggregate{}
	return &this
}

// GetId returns the Id field value
func (o *WritableAggregate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableAggregate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableAggregate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableAggregate) SetUrl(v string) {
	o.Url = v
}

// GetFamily returns the Family field value
// If the value is explicit nil, the zero value for FamilyEnum will be returned
func (o *WritableAggregate) GetFamily() FamilyEnum {
	if o == nil || o.Family.Get() == nil {
		var ret FamilyEnum
		return ret
	}

	return *o.Family.Get()
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableAggregate) GetFamilyOk() (*FamilyEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Family.Get(), o.Family.IsSet()
}

// SetFamily sets field value
func (o *WritableAggregate) SetFamily(v FamilyEnum) {
	o.Family.Set(&v)
}

// GetPrefix returns the Prefix field value
func (o *WritableAggregate) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *WritableAggregate) SetPrefix(v string) {
	o.Prefix = v
}

// GetRir returns the Rir field value
func (o *WritableAggregate) GetRir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rir
}

// GetRirOk returns a tuple with the Rir field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetRirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rir, true
}

// SetRir sets field value
func (o *WritableAggregate) SetRir(v string) {
	o.Rir = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableAggregate) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableAggregate) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableAggregate) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *WritableAggregate) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableAggregate) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableAggregate) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableAggregate) GetDateAdded() string {
	if o == nil || o.DateAdded.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateAdded.Get()
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableAggregate) GetDateAddedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAdded.Get(), o.DateAdded.IsSet()
}

// HasDateAdded returns a boolean if a field has been set.
func (o *WritableAggregate) HasDateAdded() bool {
	if o != nil && o.DateAdded.IsSet() {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given NullableString and assigns it to the DateAdded field.
func (o *WritableAggregate) SetDateAdded(v string) {
	o.DateAdded.Set(&v)
}
// SetDateAddedNil sets the value for DateAdded to be an explicit nil
func (o *WritableAggregate) SetDateAddedNil() {
	o.DateAdded.Set(nil)
}

// UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
func (o *WritableAggregate) UnsetDateAdded() {
	o.DateAdded.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableAggregate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableAggregate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableAggregate) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableAggregate) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableAggregate) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritableAggregate) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableAggregate) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableAggregate) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableAggregate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
func (o *WritableAggregate) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *WritableAggregate) SetCreated(v string) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *WritableAggregate) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *WritableAggregate) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetComputedFields returns the ComputedFields field value
func (o *WritableAggregate) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *WritableAggregate) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *WritableAggregate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableAggregate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableAggregate) SetDisplay(v string) {
	o.Display = v
}

func (o WritableAggregate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["family"] = o.Family.Get()
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if true {
		toSerialize["rir"] = o.Rir
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.DateAdded.IsSet() {
		toSerialize["date_added"] = o.DateAdded.Get()
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

type NullableWritableAggregate struct {
	value *WritableAggregate
	isSet bool
}

func (v NullableWritableAggregate) Get() *WritableAggregate {
	return v.value
}

func (v *NullableWritableAggregate) Set(val *WritableAggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableAggregate(val *WritableAggregate) *NullableWritableAggregate {
	return &NullableWritableAggregate{value: val, isSet: true}
}

func (v NullableWritableAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


