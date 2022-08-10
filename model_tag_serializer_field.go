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

// TagSerializerField NestedSerializer field for `Tag` object fields.
type TagSerializerField struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color *string `json:"color,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewTagSerializerField instantiates a new TagSerializerField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagSerializerField(id string, url string, name string, slug string, display string) *TagSerializerField {
	this := TagSerializerField{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Slug = slug
	this.Display = display
	return &this
}

// NewTagSerializerFieldWithDefaults instantiates a new TagSerializerField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagSerializerFieldWithDefaults() *TagSerializerField {
	this := TagSerializerField{}
	return &this
}

// GetId returns the Id field value
func (o *TagSerializerField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagSerializerField) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *TagSerializerField) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TagSerializerField) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *TagSerializerField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagSerializerField) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *TagSerializerField) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *TagSerializerField) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *TagSerializerField) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *TagSerializerField) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *TagSerializerField) SetColor(v string) {
	o.Color = &v
}

// GetDisplay returns the Display field value
func (o *TagSerializerField) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *TagSerializerField) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *TagSerializerField) SetDisplay(v string) {
	o.Display = v
}

func (o TagSerializerField) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableTagSerializerField struct {
	value *TagSerializerField
	isSet bool
}

func (v NullableTagSerializerField) Get() *TagSerializerField {
	return v.value
}

func (v *NullableTagSerializerField) Set(val *TagSerializerField) {
	v.value = val
	v.isSet = true
}

func (v NullableTagSerializerField) IsSet() bool {
	return v.isSet
}

func (v *NullableTagSerializerField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagSerializerField(val *TagSerializerField) *NullableTagSerializerField {
	return &NullableTagSerializerField{value: val, isSet: true}
}

func (v NullableTagSerializerField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagSerializerField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


