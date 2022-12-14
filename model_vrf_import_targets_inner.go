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

// VRFImportTargetsInner struct for VRFImportTargetsInner
type VRFImportTargetsInner struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Route target value (formatted in accordance with RFC 4360)
	Name string `json:"name"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewVRFImportTargetsInner instantiates a new VRFImportTargetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVRFImportTargetsInner(id string, url string, name string, display string) *VRFImportTargetsInner {
	this := VRFImportTargetsInner{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.Display = display
	return &this
}

// NewVRFImportTargetsInnerWithDefaults instantiates a new VRFImportTargetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVRFImportTargetsInnerWithDefaults() *VRFImportTargetsInner {
	this := VRFImportTargetsInner{}
	return &this
}

// GetId returns the Id field value
func (o *VRFImportTargetsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VRFImportTargetsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VRFImportTargetsInner) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VRFImportTargetsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VRFImportTargetsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VRFImportTargetsInner) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *VRFImportTargetsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VRFImportTargetsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VRFImportTargetsInner) SetName(v string) {
	o.Name = v
}

// GetDisplay returns the Display field value
func (o *VRFImportTargetsInner) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VRFImportTargetsInner) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VRFImportTargetsInner) SetDisplay(v string) {
	o.Display = v
}

func (o VRFImportTargetsInner) MarshalJSON() ([]byte, error) {
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
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableVRFImportTargetsInner struct {
	value *VRFImportTargetsInner
	isSet bool
}

func (v NullableVRFImportTargetsInner) Get() *VRFImportTargetsInner {
	return v.value
}

func (v *NullableVRFImportTargetsInner) Set(val *VRFImportTargetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVRFImportTargetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVRFImportTargetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVRFImportTargetsInner(val *VRFImportTargetsInner) *NullableVRFImportTargetsInner {
	return &NullableVRFImportTargetsInner{value: val, isSet: true}
}

func (v NullableVRFImportTargetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVRFImportTargetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


