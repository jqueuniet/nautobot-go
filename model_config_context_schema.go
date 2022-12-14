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

// ConfigContextSchema struct for ConfigContextSchema
type ConfigContextSchema struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewConfigContextSchema instantiates a new ConfigContextSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigContextSchema(id string, url string, name string, display string) *ConfigContextSchema {
	this := ConfigContextSchema{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Display = display
	return &this
}

// NewConfigContextSchemaWithDefaults instantiates a new ConfigContextSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigContextSchemaWithDefaults() *ConfigContextSchema {
	this := ConfigContextSchema{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigContextSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigContextSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigContextSchema) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ConfigContextSchema) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ConfigContextSchema) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ConfigContextSchema) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *ConfigContextSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigContextSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigContextSchema) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ConfigContextSchema) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextSchema) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ConfigContextSchema) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ConfigContextSchema) SetSlug(v string) {
	o.Slug = &v
}

// GetDisplay returns the Display field value
func (o *ConfigContextSchema) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ConfigContextSchema) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ConfigContextSchema) SetDisplay(v string) {
	o.Display = v
}

func (o ConfigContextSchema) MarshalJSON() ([]byte, error) {
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
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableConfigContextSchema struct {
	value *ConfigContextSchema
	isSet bool
}

func (v NullableConfigContextSchema) Get() *ConfigContextSchema {
	return v.value
}

func (v *NullableConfigContextSchema) Set(val *ConfigContextSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigContextSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigContextSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigContextSchema(val *ConfigContextSchema) *NullableConfigContextSchema {
	return &NullableConfigContextSchema{value: val, isSet: true}
}

func (v NullableConfigContextSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigContextSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


