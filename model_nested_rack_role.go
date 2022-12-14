/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NestedRackRole Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedRackRole struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	RackCount int32 `json:"rack_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedRackRole instantiates a new NestedRackRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRackRole(id string, url string, name string, rackCount int32, display string) *NestedRackRole {
	this := NestedRackRole{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.RackCount = rackCount
	this.Display = display
	return &this
}

// NewNestedRackRoleWithDefaults instantiates a new NestedRackRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRackRoleWithDefaults() *NestedRackRole {
	this := NestedRackRole{}
	return &this
}

// GetId returns the Id field value
func (o *NestedRackRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedRackRole) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedRackRole) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedRackRole) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *NestedRackRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRackRole) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *NestedRackRole) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *NestedRackRole) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *NestedRackRole) SetSlug(v string) {
	o.Slug = &v
}

// GetRackCount returns the RackCount field value
func (o *NestedRackRole) GetRackCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetRackCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RackCount, true
}

// SetRackCount sets field value
func (o *NestedRackRole) SetRackCount(v int32) {
	o.RackCount = v
}

// GetDisplay returns the Display field value
func (o *NestedRackRole) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedRackRole) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedRackRole) SetDisplay(v string) {
	o.Display = v
}

func (o NestedRackRole) MarshalJSON() ([]byte, error) {
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
		toSerialize["rack_count"] = o.RackCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedRackRole struct {
	value *NestedRackRole
	isSet bool
}

func (v NullableNestedRackRole) Get() *NestedRackRole {
	return v.value
}

func (v *NullableNestedRackRole) Set(val *NestedRackRole) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRackRole) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRackRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRackRole(val *NestedRackRole) *NullableNestedRackRole {
	return &NullableNestedRackRole{value: val, isSet: true}
}

func (v NullableNestedRackRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRackRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


