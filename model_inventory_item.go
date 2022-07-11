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

// InventoryItem Extends ModelSerializer to render any CustomFields and their values associated with an object.
type InventoryItem struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Device NestedDevice `json:"device"`
	Parent NullableString `json:"parent,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Manufacturer NullableInventoryItemManufacturer `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this item
	AssetTag NullableString `json:"asset_tag,omitempty"`
	// This item was automatically discovered
	Discovered *bool `json:"discovered,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	Depth int32 `json:"_depth"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewInventoryItem instantiates a new InventoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItem(id string, url string, device NestedDevice, name string, depth int32, display string) *InventoryItem {
	this := InventoryItem{}
	this.Id = id
	this.Url = url
	this.Device = device
	this.Name = name
	this.Depth = depth
	this.Display = display
	return &this
}

// NewInventoryItemWithDefaults instantiates a new InventoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemWithDefaults() *InventoryItem {
	this := InventoryItem{}
	return &this
}

// GetId returns the Id field value
func (o *InventoryItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryItem) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *InventoryItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InventoryItem) SetUrl(v string) {
	o.Url = v
}

// GetDevice returns the Device field value
func (o *InventoryItem) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *InventoryItem) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItem) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItem) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *InventoryItem) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *InventoryItem) SetParent(v string) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *InventoryItem) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *InventoryItem) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value
func (o *InventoryItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryItem) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InventoryItem) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InventoryItem) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InventoryItem) SetLabel(v string) {
	o.Label = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItem) GetManufacturer() InventoryItemManufacturer {
	if o == nil || o.Manufacturer.Get() == nil {
		var ret InventoryItemManufacturer
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItem) GetManufacturerOk() (*InventoryItemManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InventoryItem) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableInventoryItemManufacturer and assigns it to the Manufacturer field.
func (o *InventoryItem) SetManufacturer(v InventoryItemManufacturer) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *InventoryItem) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *InventoryItem) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *InventoryItem) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *InventoryItem) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *InventoryItem) SetPartId(v string) {
	o.PartId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InventoryItem) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InventoryItem) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InventoryItem) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItem) GetAssetTag() string {
	if o == nil || o.AssetTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItem) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *InventoryItem) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *InventoryItem) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *InventoryItem) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *InventoryItem) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *InventoryItem) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *InventoryItem) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *InventoryItem) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InventoryItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InventoryItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InventoryItem) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InventoryItem) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InventoryItem) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *InventoryItem) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetDepth returns the Depth field value
func (o *InventoryItem) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *InventoryItem) SetDepth(v int32) {
	o.Depth = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InventoryItem) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InventoryItem) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InventoryItem) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetDisplay returns the Display field value
func (o *InventoryItem) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InventoryItem) SetDisplay(v string) {
	o.Display = v
}

func (o InventoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["device"] = o.Device
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.PartId != nil {
		toSerialize["part_id"] = o.PartId
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if o.Discovered != nil {
		toSerialize["discovered"] = o.Discovered
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["_depth"] = o.Depth
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryItem struct {
	value *InventoryItem
	isSet bool
}

func (v NullableInventoryItem) Get() *InventoryItem {
	return v.value
}

func (v *NullableInventoryItem) Set(val *InventoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItem(val *InventoryItem) *NullableInventoryItem {
	return &NullableInventoryItem{value: val, isSet: true}
}

func (v NullableInventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


