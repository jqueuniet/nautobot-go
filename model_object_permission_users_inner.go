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

// ObjectPermissionUsersInner struct for ObjectPermissionUsersInner
type ObjectPermissionUsersInner struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewObjectPermissionUsersInner instantiates a new ObjectPermissionUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPermissionUsersInner(id string, url string, username string, display string) *ObjectPermissionUsersInner {
	this := ObjectPermissionUsersInner{}
	this.Id = id
	this.Url = url
	this.Username = username
	this.Display = display
	return &this
}

// NewObjectPermissionUsersInnerWithDefaults instantiates a new ObjectPermissionUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPermissionUsersInnerWithDefaults() *ObjectPermissionUsersInner {
	this := ObjectPermissionUsersInner{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectPermissionUsersInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionUsersInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectPermissionUsersInner) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ObjectPermissionUsersInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionUsersInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ObjectPermissionUsersInner) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *ObjectPermissionUsersInner) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionUsersInner) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ObjectPermissionUsersInner) SetUsername(v string) {
	o.Username = v
}

// GetDisplay returns the Display field value
func (o *ObjectPermissionUsersInner) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ObjectPermissionUsersInner) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ObjectPermissionUsersInner) SetDisplay(v string) {
	o.Display = v
}

func (o ObjectPermissionUsersInner) MarshalJSON() ([]byte, error) {
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

type NullableObjectPermissionUsersInner struct {
	value *ObjectPermissionUsersInner
	isSet bool
}

func (v NullableObjectPermissionUsersInner) Get() *ObjectPermissionUsersInner {
	return v.value
}

func (v *NullableObjectPermissionUsersInner) Set(val *ObjectPermissionUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPermissionUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPermissionUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPermissionUsersInner(val *ObjectPermissionUsersInner) *NullableObjectPermissionUsersInner {
	return &NullableObjectPermissionUsersInner{value: val, isSet: true}
}

func (v NullableObjectPermissionUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPermissionUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


