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

// WritableObjectPermission Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableObjectPermission struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types"`
	Groups []int32 `json:"groups,omitempty"`
	Users []string `json:"users,omitempty"`
	// The list of actions granted by this permission
	Actions map[string]interface{} `json:"actions"`
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints map[string]interface{} `json:"constraints,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewWritableObjectPermission instantiates a new WritableObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableObjectPermission(id string, url string, name string, objectTypes []string, actions map[string]interface{}, display string) *WritableObjectPermission {
	this := WritableObjectPermission{}
	this.Id = id
	this.Url = url
	this.Name = name
	this.ObjectTypes = objectTypes
	this.Actions = actions
	this.Display = display
	return &this
}

// NewWritableObjectPermissionWithDefaults instantiates a new WritableObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableObjectPermissionWithDefaults() *WritableObjectPermission {
	this := WritableObjectPermission{}
	return &this
}

// GetId returns the Id field value
func (o *WritableObjectPermission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WritableObjectPermission) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WritableObjectPermission) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WritableObjectPermission) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *WritableObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableObjectPermission) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableObjectPermission) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableObjectPermission) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableObjectPermission) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableObjectPermission) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableObjectPermission) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableObjectPermission) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *WritableObjectPermission) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *WritableObjectPermission) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *WritableObjectPermission) GetGroups() []int32 {
	if o == nil || o.Groups == nil {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetGroupsOk() ([]int32, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *WritableObjectPermission) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *WritableObjectPermission) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *WritableObjectPermission) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *WritableObjectPermission) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *WritableObjectPermission) SetUsers(v []string) {
	o.Users = v
}

// GetActions returns the Actions field value
func (o *WritableObjectPermission) GetActions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetActionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *WritableObjectPermission) SetActions(v map[string]interface{}) {
	o.Actions = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableObjectPermission) GetConstraints() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableObjectPermission) GetConstraintsOk() (map[string]interface{}, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *WritableObjectPermission) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given map[string]interface{} and assigns it to the Constraints field.
func (o *WritableObjectPermission) SetConstraints(v map[string]interface{}) {
	o.Constraints = v
}

// GetDisplay returns the Display field value
func (o *WritableObjectPermission) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WritableObjectPermission) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WritableObjectPermission) SetDisplay(v string) {
	o.Display = v
}

func (o WritableObjectPermission) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableWritableObjectPermission struct {
	value *WritableObjectPermission
	isSet bool
}

func (v NullableWritableObjectPermission) Get() *WritableObjectPermission {
	return v.value
}

func (v *NullableWritableObjectPermission) Set(val *WritableObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableObjectPermission(val *WritableObjectPermission) *NullableWritableObjectPermission {
	return &NullableWritableObjectPermission{value: val, isSet: true}
}

func (v NullableWritableObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


