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

// NestedClusterType Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedClusterType struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	ClusterCount int32 `json:"cluster_count"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedClusterType instantiates a new NestedClusterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedClusterType(id string, url string, name string, clusterCount int32, display string) *NestedClusterType {
	this := NestedClusterType{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.ClusterCount = clusterCount
	this.Display = display
	return &this
}

// NewNestedClusterTypeWithDefaults instantiates a new NestedClusterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterTypeWithDefaults() *NestedClusterType {
	this := NestedClusterType{}
	return &this
}

// GetId returns the Id field value
func (o *NestedClusterType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedClusterType) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedClusterType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedClusterType) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *NestedClusterType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedClusterType) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *NestedClusterType) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *NestedClusterType) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *NestedClusterType) SetSlug(v string) {
	o.Slug = &v
}

// GetClusterCount returns the ClusterCount field value
func (o *NestedClusterType) GetClusterCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetClusterCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterCount, true
}

// SetClusterCount sets field value
func (o *NestedClusterType) SetClusterCount(v int32) {
	o.ClusterCount = v
}

// GetDisplay returns the Display field value
func (o *NestedClusterType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedClusterType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedClusterType) SetDisplay(v string) {
	o.Display = v
}

func (o NestedClusterType) MarshalJSON() ([]byte, error) {
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
		toSerialize["cluster_count"] = o.ClusterCount
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedClusterType struct {
	value *NestedClusterType
	isSet bool
}

func (v NullableNestedClusterType) Get() *NestedClusterType {
	return v.value
}

func (v *NullableNestedClusterType) Set(val *NestedClusterType) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedClusterType) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedClusterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedClusterType(val *NestedClusterType) *NullableNestedClusterType {
	return &NullableNestedClusterType{value: val, isSet: true}
}

func (v NullableNestedClusterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedClusterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


