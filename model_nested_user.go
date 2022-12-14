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

// NestedUser Returns a nested representation of an object on read, but accepts either the nested representation or the primary key value on write operations.
type NestedUser struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewNestedUser instantiates a new NestedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedUser(id string, url string, username string, display string) *NestedUser {
	this := NestedUser{}
	this.Id = id
	this.Url = url
	this.Username = username
	this.Display = display
	return &this
}

// NewNestedUserWithDefaults instantiates a new NestedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedUserWithDefaults() *NestedUser {
	this := NestedUser{}
	return &this
}

// GetId returns the Id field value
func (o *NestedUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedUser) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedUser) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedUser) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *NestedUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *NestedUser) SetUsername(v string) {
	o.Username = v
}

// GetDisplay returns the Display field value
func (o *NestedUser) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedUser) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedUser) SetDisplay(v string) {
	o.Display = v
}

func (o NestedUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableNestedUser struct {
	value *NestedUser
	isSet bool
}

func (v NullableNestedUser) Get() *NestedUser {
	return v.value
}

func (v *NullableNestedUser) Set(val *NestedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedUser(val *NestedUser) *NullableNestedUser {
	return &NullableNestedUser{value: val, isSet: true}
}

func (v NullableNestedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


