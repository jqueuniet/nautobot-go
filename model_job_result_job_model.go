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

// JobResultJobModel struct for JobResultJobModel
type JobResultJobModel struct {
	Id string `json:"id"`
	Url string `json:"url"`
	// Source of the Python code for this job - local, Git repository, or plugins
	Source string `json:"source"`
	// Dotted name of the Python module providing this job
	ModuleName string `json:"module_name"`
	// Name of the Python class providing this job
	JobClassName string `json:"job_class_name"`
	// Human-readable grouping that this job belongs to
	Grouping string `json:"grouping"`
	// Human-readable name of this job
	Name string `json:"name"`
	Slug *string `json:"slug,omitempty"`
}

// NewJobResultJobModel instantiates a new JobResultJobModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResultJobModel(id string, url string, source string, moduleName string, jobClassName string, grouping string, name string) *JobResultJobModel {
	this := JobResultJobModel{}
	this.Id = id
	this.Url = url
	this.Source = source
	this.ModuleName = moduleName
	this.JobClassName = jobClassName
	this.Grouping = grouping
	this.Name = name
	return &this
}

// NewJobResultJobModelWithDefaults instantiates a new JobResultJobModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResultJobModelWithDefaults() *JobResultJobModel {
	this := JobResultJobModel{}
	return &this
}

// GetId returns the Id field value
func (o *JobResultJobModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobResultJobModel) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *JobResultJobModel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JobResultJobModel) SetUrl(v string) {
	o.Url = v
}

// GetSource returns the Source field value
func (o *JobResultJobModel) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *JobResultJobModel) SetSource(v string) {
	o.Source = v
}

// GetModuleName returns the ModuleName field value
func (o *JobResultJobModel) GetModuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModuleName
}

// GetModuleNameOk returns a tuple with the ModuleName field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetModuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleName, true
}

// SetModuleName sets field value
func (o *JobResultJobModel) SetModuleName(v string) {
	o.ModuleName = v
}

// GetJobClassName returns the JobClassName field value
func (o *JobResultJobModel) GetJobClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobClassName
}

// GetJobClassNameOk returns a tuple with the JobClassName field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetJobClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobClassName, true
}

// SetJobClassName sets field value
func (o *JobResultJobModel) SetJobClassName(v string) {
	o.JobClassName = v
}

// GetGrouping returns the Grouping field value
func (o *JobResultJobModel) GetGrouping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetGroupingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grouping, true
}

// SetGrouping sets field value
func (o *JobResultJobModel) SetGrouping(v string) {
	o.Grouping = v
}

// GetName returns the Name field value
func (o *JobResultJobModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobResultJobModel) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *JobResultJobModel) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResultJobModel) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *JobResultJobModel) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *JobResultJobModel) SetSlug(v string) {
	o.Slug = &v
}

func (o JobResultJobModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["module_name"] = o.ModuleName
	}
	if true {
		toSerialize["job_class_name"] = o.JobClassName
	}
	if true {
		toSerialize["grouping"] = o.Grouping
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	return json.Marshal(toSerialize)
}

type NullableJobResultJobModel struct {
	value *JobResultJobModel
	isSet bool
}

func (v NullableJobResultJobModel) Get() *JobResultJobModel {
	return v.value
}

func (v *NullableJobResultJobModel) Set(val *JobResultJobModel) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResultJobModel) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResultJobModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResultJobModel(val *JobResultJobModel) *NullableJobResultJobModel {
	return &NullableJobResultJobModel{value: val, isSet: true}
}

func (v NullableJobResultJobModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResultJobModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


