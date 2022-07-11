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

// PatchedCustomLink Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedCustomLink struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Jinja2 template code for link URL. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.slug }}</code>.
	TargetUrl *string `json:"target_url,omitempty"`
	Name *string `json:"name,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	// Jinja2 template code for link text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.slug }}</code>. Links which render as empty text will not be displayed.
	Text *string `json:"text,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	// Links with the same group will appear as a dropdown menu
	GroupName *string `json:"group_name,omitempty"`
	// The class of the first link in a group will be used for the dropdown button
	ButtonClass NullableButtonClassEnum `json:"button_class,omitempty"`
	// Force link to open in a new window
	NewWindow *bool `json:"new_window,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedCustomLink instantiates a new PatchedCustomLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCustomLink() *PatchedCustomLink {
	this := PatchedCustomLink{}
	return &this
}

// NewPatchedCustomLinkWithDefaults instantiates a new PatchedCustomLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCustomLinkWithDefaults() *PatchedCustomLink {
	this := PatchedCustomLink{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedCustomLink) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedCustomLink) SetUrl(v string) {
	o.Url = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *PatchedCustomLink) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCustomLink) SetName(v string) {
	o.Name = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedCustomLink) SetContentType(v string) {
	o.ContentType = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PatchedCustomLink) SetText(v string) {
	o.Text = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedCustomLink) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *PatchedCustomLink) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCustomLink) GetButtonClass() ButtonClassEnum {
	if o == nil || o.ButtonClass.Get() == nil {
		var ret ButtonClassEnum
		return ret
	}
	return *o.ButtonClass.Get()
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCustomLink) GetButtonClassOk() (*ButtonClassEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.ButtonClass.Get(), o.ButtonClass.IsSet()
}

// HasButtonClass returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasButtonClass() bool {
	if o != nil && o.ButtonClass.IsSet() {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given NullableButtonClassEnum and assigns it to the ButtonClass field.
func (o *PatchedCustomLink) SetButtonClass(v ButtonClassEnum) {
	o.ButtonClass.Set(&v)
}
// SetButtonClassNil sets the value for ButtonClass to be an explicit nil
func (o *PatchedCustomLink) SetButtonClassNil() {
	o.ButtonClass.Set(nil)
}

// UnsetButtonClass ensures that no value is present for ButtonClass, not even an explicit nil
func (o *PatchedCustomLink) UnsetButtonClass() {
	o.ButtonClass.Unset()
}

// GetNewWindow returns the NewWindow field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetNewWindow() bool {
	if o == nil || o.NewWindow == nil {
		var ret bool
		return ret
	}
	return *o.NewWindow
}

// GetNewWindowOk returns a tuple with the NewWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetNewWindowOk() (*bool, bool) {
	if o == nil || o.NewWindow == nil {
		return nil, false
	}
	return o.NewWindow, true
}

// HasNewWindow returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasNewWindow() bool {
	if o != nil && o.NewWindow != nil {
		return true
	}

	return false
}

// SetNewWindow gets a reference to the given bool and assigns it to the NewWindow field.
func (o *PatchedCustomLink) SetNewWindow(v bool) {
	o.NewWindow = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedCustomLink) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCustomLink) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedCustomLink) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedCustomLink) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedCustomLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.TargetUrl != nil {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if o.ButtonClass.IsSet() {
		toSerialize["button_class"] = o.ButtonClass.Get()
	}
	if o.NewWindow != nil {
		toSerialize["new_window"] = o.NewWindow
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedCustomLink struct {
	value *PatchedCustomLink
	isSet bool
}

func (v NullablePatchedCustomLink) Get() *PatchedCustomLink {
	return v.value
}

func (v *NullablePatchedCustomLink) Set(val *PatchedCustomLink) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCustomLink) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCustomLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCustomLink(val *PatchedCustomLink) *NullablePatchedCustomLink {
	return &NullablePatchedCustomLink{value: val, isSet: true}
}

func (v NullablePatchedCustomLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCustomLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


