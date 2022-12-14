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

// ComputedField Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ComputedField struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Internal field name
	Slug *string `json:"slug,omitempty"`
	// Name of the field as displayed to users
	Label string `json:"label"`
	Description *string `json:"description,omitempty"`
	ContentType string `json:"content_type"`
	// Jinja2 template code for field value
	Template string `json:"template"`
	// Fallback value (if any) to be output for the field in the case of a template rendering error.
	FallbackValue *string `json:"fallback_value,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewComputedField instantiates a new ComputedField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputedField(id string, url string, label string, contentType string, template string, display string) *ComputedField {
	this := ComputedField{}
	this.Id = id
	this.Url = url
	this.Label = label
	this.ContentType = contentType
	this.Template = template
	this.Display = display
	return &this
}

// NewComputedFieldWithDefaults instantiates a new ComputedField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputedFieldWithDefaults() *ComputedField {
	this := ComputedField{}
	return &this
}

// GetId returns the Id field value
func (o *ComputedField) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ComputedField) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ComputedField) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ComputedField) SetUrl(v string) {
	o.Url = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ComputedField) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputedField) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ComputedField) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ComputedField) SetSlug(v string) {
	o.Slug = &v
}

// GetLabel returns the Label field value
func (o *ComputedField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ComputedField) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ComputedField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputedField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ComputedField) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ComputedField) SetDescription(v string) {
	o.Description = &v
}

// GetContentType returns the ContentType field value
func (o *ComputedField) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ComputedField) SetContentType(v string) {
	o.ContentType = v
}

// GetTemplate returns the Template field value
func (o *ComputedField) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *ComputedField) SetTemplate(v string) {
	o.Template = v
}

// GetFallbackValue returns the FallbackValue field value if set, zero value otherwise.
func (o *ComputedField) GetFallbackValue() string {
	if o == nil || o.FallbackValue == nil {
		var ret string
		return ret
	}
	return *o.FallbackValue
}

// GetFallbackValueOk returns a tuple with the FallbackValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputedField) GetFallbackValueOk() (*string, bool) {
	if o == nil || o.FallbackValue == nil {
		return nil, false
	}
	return o.FallbackValue, true
}

// HasFallbackValue returns a boolean if a field has been set.
func (o *ComputedField) HasFallbackValue() bool {
	if o != nil && o.FallbackValue != nil {
		return true
	}

	return false
}

// SetFallbackValue gets a reference to the given string and assigns it to the FallbackValue field.
func (o *ComputedField) SetFallbackValue(v string) {
	o.FallbackValue = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ComputedField) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputedField) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ComputedField) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ComputedField) SetWeight(v int32) {
	o.Weight = &v
}

// GetDisplay returns the Display field value
func (o *ComputedField) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ComputedField) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ComputedField) SetDisplay(v string) {
	o.Display = v
}

func (o ComputedField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
		toSerialize["template"] = o.Template
	}
	if o.FallbackValue != nil {
		toSerialize["fallback_value"] = o.FallbackValue
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableComputedField struct {
	value *ComputedField
	isSet bool
}

func (v NullableComputedField) Get() *ComputedField {
	return v.value
}

func (v *NullableComputedField) Set(val *ComputedField) {
	v.value = val
	v.isSet = true
}

func (v NullableComputedField) IsSet() bool {
	return v.isSet
}

func (v *NullableComputedField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputedField(val *ComputedField) *NullableComputedField {
	return &NullableComputedField{value: val, isSet: true}
}

func (v NullableComputedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputedField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


