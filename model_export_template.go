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

// ExportTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ExportTemplate struct {
	Id string `json:"id"`
	Url string `json:"url"`
	ContentType string `json:"content_type"`
	OwnerContentType NullableString `json:"owner_content_type,omitempty"`
	OwnerObjectId NullableString `json:"owner_object_id,omitempty"`
	Owner map[string]interface{} `json:"owner"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode string `json:"template_code"`
	// Defaults to <code>text/plain</code>
	MimeType *string `json:"mime_type,omitempty"`
	// Extension to append to the rendered filename
	FileExtension *string `json:"file_extension,omitempty"`
	// Human friendly display value
	Display string `json:"display"`
}

// NewExportTemplate instantiates a new ExportTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportTemplate(id string, url string, contentType string, owner map[string]interface{}, name string, templateCode string, display string) *ExportTemplate {
	this := ExportTemplate{}
	this.Id = id
	this.Url = url
	this.ContentType = contentType
	this.Owner = owner
	this.Name = name
	this.TemplateCode = templateCode
	this.Display = display
	return &this
}

// NewExportTemplateWithDefaults instantiates a new ExportTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportTemplateWithDefaults() *ExportTemplate {
	this := ExportTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *ExportTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportTemplate) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ExportTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExportTemplate) SetUrl(v string) {
	o.Url = v
}

// GetContentType returns the ContentType field value
func (o *ExportTemplate) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ExportTemplate) SetContentType(v string) {
	o.ContentType = v
}

// GetOwnerContentType returns the OwnerContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportTemplate) GetOwnerContentType() string {
	if o == nil || o.OwnerContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerContentType.Get()
}

// GetOwnerContentTypeOk returns a tuple with the OwnerContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetOwnerContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerContentType.Get(), o.OwnerContentType.IsSet()
}

// HasOwnerContentType returns a boolean if a field has been set.
func (o *ExportTemplate) HasOwnerContentType() bool {
	if o != nil && o.OwnerContentType.IsSet() {
		return true
	}

	return false
}

// SetOwnerContentType gets a reference to the given NullableString and assigns it to the OwnerContentType field.
func (o *ExportTemplate) SetOwnerContentType(v string) {
	o.OwnerContentType.Set(&v)
}
// SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil
func (o *ExportTemplate) SetOwnerContentTypeNil() {
	o.OwnerContentType.Set(nil)
}

// UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
func (o *ExportTemplate) UnsetOwnerContentType() {
	o.OwnerContentType.Unset()
}

// GetOwnerObjectId returns the OwnerObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportTemplate) GetOwnerObjectId() string {
	if o == nil || o.OwnerObjectId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerObjectId.Get()
}

// GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetOwnerObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerObjectId.Get(), o.OwnerObjectId.IsSet()
}

// HasOwnerObjectId returns a boolean if a field has been set.
func (o *ExportTemplate) HasOwnerObjectId() bool {
	if o != nil && o.OwnerObjectId.IsSet() {
		return true
	}

	return false
}

// SetOwnerObjectId gets a reference to the given NullableString and assigns it to the OwnerObjectId field.
func (o *ExportTemplate) SetOwnerObjectId(v string) {
	o.OwnerObjectId.Set(&v)
}
// SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil
func (o *ExportTemplate) SetOwnerObjectIdNil() {
	o.OwnerObjectId.Set(nil)
}

// UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
func (o *ExportTemplate) UnsetOwnerObjectId() {
	o.OwnerObjectId.Unset()
}

// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ExportTemplate) GetOwner() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportTemplate) GetOwnerOk() (map[string]interface{}, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// SetOwner sets field value
func (o *ExportTemplate) SetOwner(v map[string]interface{}) {
	o.Owner = v
}

// GetName returns the Name field value
func (o *ExportTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportTemplate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportTemplate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateCode returns the TemplateCode field value
func (o *ExportTemplate) GetTemplateCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetTemplateCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCode, true
}

// SetTemplateCode sets field value
func (o *ExportTemplate) SetTemplateCode(v string) {
	o.TemplateCode = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *ExportTemplate) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *ExportTemplate) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *ExportTemplate) SetMimeType(v string) {
	o.MimeType = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *ExportTemplate) GetFileExtension() string {
	if o == nil || o.FileExtension == nil {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetFileExtensionOk() (*string, bool) {
	if o == nil || o.FileExtension == nil {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *ExportTemplate) HasFileExtension() bool {
	if o != nil && o.FileExtension != nil {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *ExportTemplate) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetDisplay returns the Display field value
func (o *ExportTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ExportTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ExportTemplate) SetDisplay(v string) {
	o.Display = v
}

func (o ExportTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if o.OwnerContentType.IsSet() {
		toSerialize["owner_content_type"] = o.OwnerContentType.Get()
	}
	if o.OwnerObjectId.IsSet() {
		toSerialize["owner_object_id"] = o.OwnerObjectId.Get()
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["template_code"] = o.TemplateCode
	}
	if o.MimeType != nil {
		toSerialize["mime_type"] = o.MimeType
	}
	if o.FileExtension != nil {
		toSerialize["file_extension"] = o.FileExtension
	}
	if true {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullableExportTemplate struct {
	value *ExportTemplate
	isSet bool
}

func (v NullableExportTemplate) Get() *ExportTemplate {
	return v.value
}

func (v *NullableExportTemplate) Set(val *ExportTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableExportTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableExportTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportTemplate(val *ExportTemplate) *NullableExportTemplate {
	return &NullableExportTemplate{value: val, isSet: true}
}

func (v NullableExportTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


