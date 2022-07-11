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

// WritableVirtualChassis Extends ModelSerializer to render any CustomFields and their values associated with an object.
type WritableVirtualChassis struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Domain *string `json:"domain,omitempty"`
	Master NullableString `json:"master,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	MemberCount int32 `json:"member_count"`
	ComputedFields map[string]interface{} `json:"computed_fields"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableVirtualChassis instantiates a new WritableVirtualChassis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableVirtualChassis(id string, url string, name string, memberCount int32, computedFields map[string]interface{}, display string) *WritableVirtualChassis {
	this := WritableVirtualChassis{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.MemberCount = memberCount
	this.ComputedFields = computedFields
	this.Display = display
	return &this
}

// NewWritableVirtualChassisWithDefaults instantiates a new WritableVirtualChassis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableVirtualChassisWithDefaults() *WritableVirtualChassis {
	this := WritableVirtualChassis{}
	return &this
}

// GetId returns the Id field value
func (o *WritableVirtualChassis) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableVirtualChassis) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableVirtualChassis) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableVirtualChassis) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *WritableVirtualChassis) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableVirtualChassis) SetName(v string) {
	o.Name = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *WritableVirtualChassis) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *WritableVirtualChassis) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *WritableVirtualChassis) SetDomain(v string) {
	o.Domain = &v
}

// GetMaster returns the Master field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableVirtualChassis) GetMaster() string {
	if o == nil || o.Master.Get() == nil {
		var ret string
		return ret
	}
	return *o.Master.Get()
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableVirtualChassis) GetMasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Master.Get(), o.Master.IsSet()
}

// HasMaster returns a boolean if a field has been set.
func (o *WritableVirtualChassis) HasMaster() bool {
	if o != nil && o.Master.IsSet() {
		return true
	}

	return false
}

// SetMaster gets a reference to the given NullableString and assigns it to the Master field.
func (o *WritableVirtualChassis) SetMaster(v string) {
	o.Master.Set(&v)
}
// SetMasterNil sets the value for Master to be an explicit nil
func (o *WritableVirtualChassis) SetMasterNil() {
	o.Master.Set(nil)
}

// UnsetMaster ensures that no value is present for Master, not even an explicit nil
func (o *WritableVirtualChassis) UnsetMaster() {
	o.Master.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableVirtualChassis) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableVirtualChassis) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *WritableVirtualChassis) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableVirtualChassis) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableVirtualChassis) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableVirtualChassis) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetMemberCount returns the MemberCount field value
func (o *WritableVirtualChassis) GetMemberCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetMemberCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *WritableVirtualChassis) SetMemberCount(v int32) {
	o.MemberCount = v
}

// GetComputedFields returns the ComputedFields field value
func (o *WritableVirtualChassis) GetComputedFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// SetComputedFields sets field value
func (o *WritableVirtualChassis) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value
func (o *WritableVirtualChassis) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableVirtualChassis) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableVirtualChassis) SetDisplay(v string) {
	o.Display = v
}

func (o WritableVirtualChassis) MarshalJSON() ([]byte, error) {
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
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Master.IsSet() {
		toSerialize["master"] = o.Master.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["member_count"] = o.MemberCount
	}
	if true {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritableVirtualChassis struct {
	value *WritableVirtualChassis
	isSet bool
}

func (v NullableWritableVirtualChassis) Get() *WritableVirtualChassis {
	return v.value
}

func (v *NullableWritableVirtualChassis) Set(val *WritableVirtualChassis) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableVirtualChassis) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableVirtualChassis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableVirtualChassis(val *WritableVirtualChassis) *NullableWritableVirtualChassis {
	return &NullableWritableVirtualChassis{value: val, isSet: true}
}

func (v NullableWritableVirtualChassis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableVirtualChassis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


