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

// PatchedImageAttachment Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedImageAttachment struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	ObjectId *string `json:"object_id,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty"`
	Name *string `json:"name,omitempty"`
	Image *string `json:"image,omitempty"`
	ImageHeight *int32 `json:"image_height,omitempty"`
	ImageWidth *int32 `json:"image_width,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedImageAttachment instantiates a new PatchedImageAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedImageAttachment() *PatchedImageAttachment {
	this := PatchedImageAttachment{}
	return &this
}

// NewPatchedImageAttachmentWithDefaults instantiates a new PatchedImageAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedImageAttachmentWithDefaults() *PatchedImageAttachment {
	this := PatchedImageAttachment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedImageAttachment) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedImageAttachment) SetUrl(v string) {
	o.Url = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedImageAttachment) SetContentType(v string) {
	o.ContentType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *PatchedImageAttachment) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetParent() map[string]interface{} {
	if o == nil || o.Parent == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetParentOk() (map[string]interface{}, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given map[string]interface{} and assigns it to the Parent field.
func (o *PatchedImageAttachment) SetParent(v map[string]interface{}) {
	o.Parent = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedImageAttachment) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *PatchedImageAttachment) SetImage(v string) {
	o.Image = &v
}

// GetImageHeight returns the ImageHeight field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetImageHeight() int32 {
	if o == nil || o.ImageHeight == nil {
		var ret int32
		return ret
	}
	return *o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetImageHeightOk() (*int32, bool) {
	if o == nil || o.ImageHeight == nil {
		return nil, false
	}
	return o.ImageHeight, true
}

// HasImageHeight returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasImageHeight() bool {
	if o != nil && o.ImageHeight != nil {
		return true
	}

	return false
}

// SetImageHeight gets a reference to the given int32 and assigns it to the ImageHeight field.
func (o *PatchedImageAttachment) SetImageHeight(v int32) {
	o.ImageHeight = &v
}

// GetImageWidth returns the ImageWidth field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetImageWidth() int32 {
	if o == nil || o.ImageWidth == nil {
		var ret int32
		return ret
	}
	return *o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetImageWidthOk() (*int32, bool) {
	if o == nil || o.ImageWidth == nil {
		return nil, false
	}
	return o.ImageWidth, true
}

// HasImageWidth returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasImageWidth() bool {
	if o != nil && o.ImageWidth != nil {
		return true
	}

	return false
}

// SetImageWidth gets a reference to the given int32 and assigns it to the ImageWidth field.
func (o *PatchedImageAttachment) SetImageWidth(v int32) {
	o.ImageWidth = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PatchedImageAttachment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedImageAttachment) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachment) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedImageAttachment) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedImageAttachment) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedImageAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.ImageHeight != nil {
		toSerialize["image_height"] = o.ImageHeight
	}
	if o.ImageWidth != nil {
		toSerialize["image_width"] = o.ImageWidth
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedImageAttachment struct {
	value *PatchedImageAttachment
	isSet bool
}

func (v NullablePatchedImageAttachment) Get() *PatchedImageAttachment {
	return v.value
}

func (v *NullablePatchedImageAttachment) Set(val *PatchedImageAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedImageAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedImageAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedImageAttachment(val *PatchedImageAttachment) *NullablePatchedImageAttachment {
	return &NullablePatchedImageAttachment{value: val, isSet: true}
}

func (v NullablePatchedImageAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedImageAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


