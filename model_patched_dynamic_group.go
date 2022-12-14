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

// PatchedDynamicGroup Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedDynamicGroup struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Dynamic Group name
	Name *string `json:"name,omitempty"`
	// Unique slug
	Slug *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	// A JSON-encoded dictionary of filter parameters for group membership
	Filter map[string]interface{} `json:"filter,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedDynamicGroup instantiates a new PatchedDynamicGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDynamicGroup() *PatchedDynamicGroup {
	this := PatchedDynamicGroup{}
	return &this
}

// NewPatchedDynamicGroupWithDefaults instantiates a new PatchedDynamicGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDynamicGroupWithDefaults() *PatchedDynamicGroup {
	this := PatchedDynamicGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedDynamicGroup) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedDynamicGroup) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDynamicGroup) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedDynamicGroup) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedDynamicGroup) SetDescription(v string) {
	o.Description = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedDynamicGroup) SetContentType(v string) {
	o.ContentType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetFilter() map[string]interface{} {
	if o == nil || o.Filter == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *PatchedDynamicGroup) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedDynamicGroup) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroup) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedDynamicGroup) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedDynamicGroup) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedDynamicGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDynamicGroup struct {
	value *PatchedDynamicGroup
	isSet bool
}

func (v NullablePatchedDynamicGroup) Get() *PatchedDynamicGroup {
	return v.value
}

func (v *NullablePatchedDynamicGroup) Set(val *PatchedDynamicGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDynamicGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDynamicGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDynamicGroup(val *PatchedDynamicGroup) *NullablePatchedDynamicGroup {
	return &NullablePatchedDynamicGroup{value: val, isSet: true}
}

func (v NullablePatchedDynamicGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDynamicGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


