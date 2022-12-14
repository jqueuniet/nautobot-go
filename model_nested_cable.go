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

// NestedCable This base serializer implements common fields and logic for all ModelSerializers. Namely it defines the `display` field which exposes a human friendly value for the given object.
type NestedCable struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Label *string `json:"label,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedCable instantiates a new NestedCable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCable(id string, url string, display string) *NestedCable {
	this := NestedCable{}
	this.Id = id
	this.Url = url
	this.Display = display
	return &this
}

// NewNestedCableWithDefaults instantiates a new NestedCable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCableWithDefaults() *NestedCable {
	this := NestedCable{}
	return &this
}

// GetId returns the Id field value
func (o *NestedCable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedCable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedCable) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedCable) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedCable) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedCable) SetUrl(v string) {
	o.Url = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NestedCable) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedCable) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NestedCable) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NestedCable) SetLabel(v string) {
	o.Label = &v
}

// GetDisplay returns the Display field value
func (o *NestedCable) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedCable) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedCable) SetDisplay(v string) {
	o.Display = v
}

func (o NestedCable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedCable struct {
	value *NestedCable
	isSet bool
}

func (v NullableNestedCable) Get() *NestedCable {
	return v.value
}

func (v *NullableNestedCable) Set(val *NestedCable) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCable) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCable(val *NestedCable) *NullableNestedCable {
	return &NullableNestedCable{value: val, isSet: true}
}

func (v NullableNestedCable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


