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

// InterfaceTaggedVlansInner struct for InterfaceTaggedVlansInner
type InterfaceTaggedVlansInner struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewInterfaceTaggedVlansInner instantiates a new InterfaceTaggedVlansInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTaggedVlansInner(id string, url string, vid int32, name string, display string) *InterfaceTaggedVlansInner {
	this := InterfaceTaggedVlansInner{}
	this.Id = id
	this.Url = url
	this.Vid = vid
	this.Name = name
	this.Display = display
	return &this
}

// NewInterfaceTaggedVlansInnerWithDefaults instantiates a new InterfaceTaggedVlansInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTaggedVlansInnerWithDefaults() *InterfaceTaggedVlansInner {
	this := InterfaceTaggedVlansInner{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceTaggedVlansInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceTaggedVlansInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceTaggedVlansInner) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *InterfaceTaggedVlansInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InterfaceTaggedVlansInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InterfaceTaggedVlansInner) SetUrl(v string) {
	o.Url = v
}

// GetVid returns the Vid field value
func (o *InterfaceTaggedVlansInner) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *InterfaceTaggedVlansInner) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *InterfaceTaggedVlansInner) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *InterfaceTaggedVlansInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterfaceTaggedVlansInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterfaceTaggedVlansInner) SetName(v string) {
	o.Name = v
}

// GetDisplay returns the Display field value
func (o *InterfaceTaggedVlansInner) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InterfaceTaggedVlansInner) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InterfaceTaggedVlansInner) SetDisplay(v string) {
	o.Display = v
}

func (o InterfaceTaggedVlansInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["vid"] = o.Vid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceTaggedVlansInner struct {
	value *InterfaceTaggedVlansInner
	isSet bool
}

func (v NullableInterfaceTaggedVlansInner) Get() *InterfaceTaggedVlansInner {
	return v.value
}

func (v *NullableInterfaceTaggedVlansInner) Set(val *InterfaceTaggedVlansInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTaggedVlansInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTaggedVlansInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTaggedVlansInner(val *InterfaceTaggedVlansInner) *NullableInterfaceTaggedVlansInner {
	return &NullableInterfaceTaggedVlansInner{value: val, isSet: true}
}

func (v NullableInterfaceTaggedVlansInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTaggedVlansInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


