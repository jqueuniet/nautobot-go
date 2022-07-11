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

// NestedTenantGroup Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedTenantGroup struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	TenantCount int32 `json:"tenant_count"`
	Depth int32 `json:"_depth"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedTenantGroup instantiates a new NestedTenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTenantGroup(id string, url string, name string, tenantCount int32, depth int32, display string) *NestedTenantGroup {
	this := NestedTenantGroup{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.TenantCount = tenantCount
	this.Depth = depth
	this.Display = display
	return &this
}

// NewNestedTenantGroupWithDefaults instantiates a new NestedTenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTenantGroupWithDefaults() *NestedTenantGroup {
	this := NestedTenantGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedTenantGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedTenantGroup) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedTenantGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedTenantGroup) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *NestedTenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *NestedTenantGroup) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *NestedTenantGroup) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *NestedTenantGroup) SetSlug(v string) {
	o.Slug = &v
}

// GetTenantCount returns the TenantCount field value
func (o *NestedTenantGroup) GetTenantCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantCount
}

// GetTenantCountOk returns a tuple with the TenantCount field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetTenantCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantCount, true
}

// SetTenantCount sets field value
func (o *NestedTenantGroup) SetTenantCount(v int32) {
	o.TenantCount = v
}

// GetDepth returns the Depth field value
func (o *NestedTenantGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *NestedTenantGroup) SetDepth(v int32) {
	o.Depth = v
}

// GetDisplay returns the Display field value
func (o *NestedTenantGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedTenantGroup) SetDisplay(v string) {
	o.Display = v
}

func (o NestedTenantGroup) MarshalJSON() ([]byte, error) {
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
		toSerialize["tenant_count"] = o.TenantCount
	}
	if true {
		toSerialize["_depth"] = o.Depth
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedTenantGroup struct {
	value *NestedTenantGroup
	isSet bool
}

func (v NullableNestedTenantGroup) Get() *NestedTenantGroup {
	return v.value
}

func (v *NullableNestedTenantGroup) Set(val *NestedTenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTenantGroup(val *NestedTenantGroup) *NullableNestedTenantGroup {
	return &NullableNestedTenantGroup{value: val, isSet: true}
}

func (v NullableNestedTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


