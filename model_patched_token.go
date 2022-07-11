/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PatchedToken Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedToken struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Expires NullableTime `json:"expires,omitempty"`
	Key *string `json:"key,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled *bool `json:"write_enabled,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewPatchedToken instantiates a new PatchedToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedToken() *PatchedToken {
	this := PatchedToken{}
	return &this
}

// NewPatchedTokenWithDefaults instantiates a new PatchedToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTokenWithDefaults() *PatchedToken {
	this := PatchedToken{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedToken) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedToken) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedToken) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedToken) SetUrl(v string) {
	o.Url = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedToken) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedToken) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedToken) SetDisplay(v string) {
	o.Display = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedToken) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedToken) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PatchedToken) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedToken) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedToken) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *PatchedToken) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *PatchedToken) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *PatchedToken) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *PatchedToken) UnsetExpires() {
	o.Expires.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PatchedToken) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PatchedToken) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PatchedToken) SetKey(v string) {
	o.Key = &v
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *PatchedToken) GetWriteEnabled() bool {
	if o == nil || o.WriteEnabled == nil {
		var ret bool
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetWriteEnabledOk() (*bool, bool) {
	if o == nil || o.WriteEnabled == nil {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *PatchedToken) HasWriteEnabled() bool {
	if o != nil && o.WriteEnabled != nil {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given bool and assigns it to the WriteEnabled field.
func (o *PatchedToken) SetWriteEnabled(v bool) {
	o.WriteEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedToken) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToken) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedToken) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedToken) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.WriteEnabled != nil {
		toSerialize["write_enabled"] = o.WriteEnabled
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedToken struct {
	value *PatchedToken
	isSet bool
}

func (v NullablePatchedToken) Get() *PatchedToken {
	return v.value
}

func (v *NullablePatchedToken) Set(val *PatchedToken) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedToken) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedToken(val *PatchedToken) *NullablePatchedToken {
	return &NullablePatchedToken{value: val, isSet: true}
}

func (v NullablePatchedToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


