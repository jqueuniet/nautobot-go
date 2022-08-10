/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PatchedManufacturer Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedManufacturer struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	DevicetypeCount *int32 `json:"devicetype_count,omitempty"`
	InventoryitemCount *int32 `json:"inventoryitem_count,omitempty"`
	PlatformCount *int32 `json:"platform_count,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedManufacturer instantiates a new PatchedManufacturer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedManufacturer() *PatchedManufacturer {
	this := PatchedManufacturer{}
	return &this
}

// NewPatchedManufacturerWithDefaults instantiates a new PatchedManufacturer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedManufacturerWithDefaults() *PatchedManufacturer {
	this := PatchedManufacturer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedManufacturer) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedManufacturer) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedManufacturer) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedManufacturer) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedManufacturer) SetDescription(v string) {
	o.Description = &v
}

// GetDevicetypeCount returns the DevicetypeCount field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetDevicetypeCount() int32 {
	if o == nil || o.DevicetypeCount == nil {
		var ret int32
		return ret
	}
	return *o.DevicetypeCount
}

// GetDevicetypeCountOk returns a tuple with the DevicetypeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetDevicetypeCountOk() (*int32, bool) {
	if o == nil || o.DevicetypeCount == nil {
		return nil, false
	}
	return o.DevicetypeCount, true
}

// HasDevicetypeCount returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasDevicetypeCount() bool {
	if o != nil && o.DevicetypeCount != nil {
		return true
	}

	return false
}

// SetDevicetypeCount gets a reference to the given int32 and assigns it to the DevicetypeCount field.
func (o *PatchedManufacturer) SetDevicetypeCount(v int32) {
	o.DevicetypeCount = &v
}

// GetInventoryitemCount returns the InventoryitemCount field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetInventoryitemCount() int32 {
	if o == nil || o.InventoryitemCount == nil {
		var ret int32
		return ret
	}
	return *o.InventoryitemCount
}

// GetInventoryitemCountOk returns a tuple with the InventoryitemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetInventoryitemCountOk() (*int32, bool) {
	if o == nil || o.InventoryitemCount == nil {
		return nil, false
	}
	return o.InventoryitemCount, true
}

// HasInventoryitemCount returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasInventoryitemCount() bool {
	if o != nil && o.InventoryitemCount != nil {
		return true
	}

	return false
}

// SetInventoryitemCount gets a reference to the given int32 and assigns it to the InventoryitemCount field.
func (o *PatchedManufacturer) SetInventoryitemCount(v int32) {
	o.InventoryitemCount = &v
}

// GetPlatformCount returns the PlatformCount field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetPlatformCount() int32 {
	if o == nil || o.PlatformCount == nil {
		var ret int32
		return ret
	}
	return *o.PlatformCount
}

// GetPlatformCountOk returns a tuple with the PlatformCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetPlatformCountOk() (*int32, bool) {
	if o == nil || o.PlatformCount == nil {
		return nil, false
	}
	return o.PlatformCount, true
}

// HasPlatformCount returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasPlatformCount() bool {
	if o != nil && o.PlatformCount != nil {
		return true
	}

	return false
}

// SetPlatformCount gets a reference to the given int32 and assigns it to the PlatformCount field.
func (o *PatchedManufacturer) SetPlatformCount(v int32) {
	o.PlatformCount = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedManufacturer) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PatchedManufacturer) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PatchedManufacturer) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedManufacturer) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedManufacturer) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedManufacturer) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedManufacturer) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedManufacturer) MarshalJSON() ([]byte, error) {
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
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DevicetypeCount != nil {
		toSerialize["devicetype_count"] = o.DevicetypeCount
	}
	if o.InventoryitemCount != nil {
		toSerialize["inventoryitem_count"] = o.InventoryitemCount
	}
	if o.PlatformCount != nil {
		toSerialize["platform_count"] = o.PlatformCount
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedManufacturer struct {
	value *PatchedManufacturer
	isSet bool
}

func (v NullablePatchedManufacturer) Get() *PatchedManufacturer {
	return v.value
}

func (v *NullablePatchedManufacturer) Set(val *PatchedManufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedManufacturer(val *PatchedManufacturer) *NullablePatchedManufacturer {
	return &NullablePatchedManufacturer{value: val, isSet: true}
}

func (v NullablePatchedManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


