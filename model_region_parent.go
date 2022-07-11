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

// RegionParent struct for RegionParent
type RegionParent struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	SiteCount int32 `json:"site_count"`
	Depth int32 `json:"_depth"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewRegionParent instantiates a new RegionParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionParent(id string, url string, name string, siteCount int32, depth int32, display string) *RegionParent {
	this := RegionParent{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.SiteCount = siteCount
	this.Depth = depth
	this.Display = display
	return &this
}

// NewRegionParentWithDefaults instantiates a new RegionParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionParentWithDefaults() *RegionParent {
	this := RegionParent{}
	return &this
}

// GetId returns the Id field value
func (o *RegionParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegionParent) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *RegionParent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RegionParent) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *RegionParent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegionParent) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *RegionParent) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionParent) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *RegionParent) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *RegionParent) SetSlug(v string) {
	o.Slug = &v
}

// GetSiteCount returns the SiteCount field value
func (o *RegionParent) GetSiteCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetSiteCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteCount, true
}

// SetSiteCount sets field value
func (o *RegionParent) SetSiteCount(v int32) {
	o.SiteCount = v
}

// GetDepth returns the Depth field value
func (o *RegionParent) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *RegionParent) SetDepth(v int32) {
	o.Depth = v
}

// GetDisplay returns the Display field value
func (o *RegionParent) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *RegionParent) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *RegionParent) SetDisplay(v string) {
	o.Display = v
}

func (o RegionParent) MarshalJSON() ([]byte, error) {
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
		toSerialize["site_count"] = o.SiteCount
	}
	if true {
		toSerialize["_depth"] = o.Depth
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableRegionParent struct {
	value *RegionParent
	isSet bool
}

func (v NullableRegionParent) Get() *RegionParent {
	return v.value
}

func (v *NullableRegionParent) Set(val *RegionParent) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionParent) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionParent(val *RegionParent) *NullableRegionParent {
	return &NullableRegionParent{value: val, isSet: true}
}

func (v NullableRegionParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


