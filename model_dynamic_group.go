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

// DynamicGroup Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type DynamicGroup struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Dynamic Group name
	Name string `json:"name"`
	// Unique slug
	Slug *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentType string `json:"content_type"`
	// A JSON-encoded dictionary of filter parameters for group membership
	Filter map[string]interface{} `json:"filter"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewDynamicGroup instantiates a new DynamicGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicGroup(id string, url string, name string, contentType string, filter map[string]interface{}, display string) *DynamicGroup {
	this := DynamicGroup{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.ContentType = contentType
	this.Filter = filter
	this.Display = display
	return &this
}

// NewDynamicGroupWithDefaults instantiates a new DynamicGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicGroupWithDefaults() *DynamicGroup {
	this := DynamicGroup{}
	return &this
}

// GetId returns the Id field value
func (o *DynamicGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DynamicGroup) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DynamicGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DynamicGroup) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *DynamicGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DynamicGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *DynamicGroup) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *DynamicGroup) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *DynamicGroup) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicGroup) SetDescription(v string) {
	o.Description = &v
}

// GetContentType returns the ContentType field value
func (o *DynamicGroup) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *DynamicGroup) SetContentType(v string) {
	o.ContentType = v
}

// GetFilter returns the Filter field value
func (o *DynamicGroup) GetFilter() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter, true
}

// SetFilter sets field value
func (o *DynamicGroup) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetDisplay returns the Display field value
func (o *DynamicGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DynamicGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DynamicGroup) SetDisplay(v string) {
	o.Display = v
}

func (o DynamicGroup) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicGroup struct {
	value *DynamicGroup
	isSet bool
}

func (v NullableDynamicGroup) Get() *DynamicGroup {
	return v.value
}

func (v *NullableDynamicGroup) Set(val *DynamicGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicGroup(val *DynamicGroup) *NullableDynamicGroup {
	return &NullableDynamicGroup{value: val, isSet: true}
}

func (v NullableDynamicGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


