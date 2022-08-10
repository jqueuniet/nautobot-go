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

// PatchedWritableDeviceType Extends ModelSerializer to render any CustomFields and their values associated with an object.
type PatchedWritableDeviceType struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Model *string `json:"model,omitempty"`
	Slug *string `json:"slug,omitempty"`
	// Discrete part number (optional)
	PartNumber *string `json:"part_number,omitempty"`
	UHeight *int32 `json:"u_height,omitempty"`
	// Device consumes both front and rear rack faces
	IsFullDepth *bool `json:"is_full_depth,omitempty"`
	SubdeviceRole *ParentChildStatus `json:"subdevice_role,omitempty"`
	FrontImage *string `json:"front_image,omitempty"`
	RearImage *string `json:"rear_image,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableDeviceType instantiates a new PatchedWritableDeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableDeviceType() *PatchedWritableDeviceType {
	this := PatchedWritableDeviceType{}
	return &this
}

// NewPatchedWritableDeviceTypeWithDefaults instantiates a new PatchedWritableDeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableDeviceTypeWithDefaults() *PatchedWritableDeviceType {
	this := PatchedWritableDeviceType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableDeviceType) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableDeviceType) SetUrl(v string) {
	o.Url = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *PatchedWritableDeviceType) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *PatchedWritableDeviceType) SetModel(v string) {
	o.Model = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PatchedWritableDeviceType) SetSlug(v string) {
	o.Slug = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *PatchedWritableDeviceType) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetUHeight() int32 {
	if o == nil || o.UHeight == nil {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetUHeightOk() (*int32, bool) {
	if o == nil || o.UHeight == nil {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasUHeight() bool {
	if o != nil && o.UHeight != nil {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *PatchedWritableDeviceType) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetIsFullDepth() bool {
	if o == nil || o.IsFullDepth == nil {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || o.IsFullDepth == nil {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasIsFullDepth() bool {
	if o != nil && o.IsFullDepth != nil {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *PatchedWritableDeviceType) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetSubdeviceRole() ParentChildStatus {
	if o == nil || o.SubdeviceRole == nil {
		var ret ParentChildStatus
		return ret
	}
	return *o.SubdeviceRole
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetSubdeviceRoleOk() (*ParentChildStatus, bool) {
	if o == nil || o.SubdeviceRole == nil {
		return nil, false
	}
	return o.SubdeviceRole, true
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasSubdeviceRole() bool {
	if o != nil && o.SubdeviceRole != nil {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given ParentChildStatus and assigns it to the SubdeviceRole field.
func (o *PatchedWritableDeviceType) SetSubdeviceRole(v ParentChildStatus) {
	o.SubdeviceRole = &v
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetFrontImage() string {
	if o == nil || o.FrontImage == nil {
		var ret string
		return ret
	}
	return *o.FrontImage
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetFrontImageOk() (*string, bool) {
	if o == nil || o.FrontImage == nil {
		return nil, false
	}
	return o.FrontImage, true
}

// HasFrontImage returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasFrontImage() bool {
	if o != nil && o.FrontImage != nil {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given string and assigns it to the FrontImage field.
func (o *PatchedWritableDeviceType) SetFrontImage(v string) {
	o.FrontImage = &v
}

// GetRearImage returns the RearImage field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetRearImage() string {
	if o == nil || o.RearImage == nil {
		var ret string
		return ret
	}
	return *o.RearImage
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetRearImageOk() (*string, bool) {
	if o == nil || o.RearImage == nil {
		return nil, false
	}
	return o.RearImage, true
}

// HasRearImage returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasRearImage() bool {
	if o != nil && o.RearImage != nil {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given string and assigns it to the RearImage field.
func (o *PatchedWritableDeviceType) SetRearImage(v string) {
	o.RearImage = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableDeviceType) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableDeviceType) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableDeviceType) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PatchedWritableDeviceType) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PatchedWritableDeviceType) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetDeviceCountOk() (*int32, bool) {
	if o == nil || o.DeviceCount == nil {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *PatchedWritableDeviceType) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableDeviceType) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableDeviceType) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceType) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableDeviceType) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableDeviceType) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableDeviceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.PartNumber != nil {
		toSerialize["part_number"] = o.PartNumber
	}
	if o.UHeight != nil {
		toSerialize["u_height"] = o.UHeight
	}
	if o.IsFullDepth != nil {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if o.SubdeviceRole != nil {
		toSerialize["subdevice_role"] = o.SubdeviceRole
	}
	if o.FrontImage != nil {
		toSerialize["front_image"] = o.FrontImage
	}
	if o.RearImage != nil {
		toSerialize["rear_image"] = o.RearImage
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
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
	if o.DeviceCount != nil {
		toSerialize["device_count"] = o.DeviceCount
	}
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableDeviceType struct {
	value *PatchedWritableDeviceType
	isSet bool
}

func (v NullablePatchedWritableDeviceType) Get() *PatchedWritableDeviceType {
	return v.value
}

func (v *NullablePatchedWritableDeviceType) Set(val *PatchedWritableDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableDeviceType(val *PatchedWritableDeviceType) *NullablePatchedWritableDeviceType {
	return &NullablePatchedWritableDeviceType{value: val, isSet: true}
}

func (v NullablePatchedWritableDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


