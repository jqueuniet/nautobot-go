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

// PatchedWritableObjectPermission Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableObjectPermission struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
	Groups []int32 `json:"groups,omitempty"`
	Users []string `json:"users,omitempty"`
	// The list of actions granted by this permission
	Actions map[string]interface{} `json:"actions,omitempty"`
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints map[string]interface{} `json:"constraints,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableObjectPermission instantiates a new PatchedWritableObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableObjectPermission() *PatchedWritableObjectPermission {
	this := PatchedWritableObjectPermission{}
	return &this
}

// NewPatchedWritableObjectPermissionWithDefaults instantiates a new PatchedWritableObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableObjectPermissionWithDefaults() *PatchedWritableObjectPermission {
	this := PatchedWritableObjectPermission{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableObjectPermission) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableObjectPermission) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableObjectPermission) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableObjectPermission) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedWritableObjectPermission) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetObjectTypes returns the ObjectTypes field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetObjectTypes() []string {
	if o == nil || o.ObjectTypes == nil {
		var ret []string
		return ret
	}
	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetObjectTypesOk() ([]string, bool) {
	if o == nil || o.ObjectTypes == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// HasObjectTypes returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasObjectTypes() bool {
	if o != nil && o.ObjectTypes != nil {
		return true
	}

	return false
}

// SetObjectTypes gets a reference to the given []string and assigns it to the ObjectTypes field.
func (o *PatchedWritableObjectPermission) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetGroups() []int32 {
	if o == nil || o.Groups == nil {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetGroupsOk() ([]int32, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *PatchedWritableObjectPermission) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *PatchedWritableObjectPermission) SetUsers(v []string) {
	o.Users = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetActions() map[string]interface{} {
	if o == nil || o.Actions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetActionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given map[string]interface{} and assigns it to the Actions field.
func (o *PatchedWritableObjectPermission) SetActions(v map[string]interface{}) {
	o.Actions = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableObjectPermission) GetConstraints() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableObjectPermission) GetConstraintsOk() (map[string]interface{}, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given map[string]interface{} and assigns it to the Constraints field.
func (o *PatchedWritableObjectPermission) SetConstraints(v map[string]interface{}) {
	o.Constraints = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableObjectPermission) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableObjectPermission) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableObjectPermission) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableObjectPermission) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableObjectPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ObjectTypes != nil {
		toSerialize["object_types"] = o.ObjectTypes
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableObjectPermission struct {
	value *PatchedWritableObjectPermission
	isSet bool
}

func (v NullablePatchedWritableObjectPermission) Get() *PatchedWritableObjectPermission {
	return v.value
}

func (v *NullablePatchedWritableObjectPermission) Set(val *PatchedWritableObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableObjectPermission(val *PatchedWritableObjectPermission) *NullablePatchedWritableObjectPermission {
	return &NullablePatchedWritableObjectPermission{value: val, isSet: true}
}

func (v NullablePatchedWritableObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


