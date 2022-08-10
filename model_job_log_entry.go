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

// JobLogEntry struct for JobLogEntry
type JobLogEntry struct {
	Id string `json:"id"`
	Url string `json:"url"`
	AbsoluteUrl NullableString `json:"absolute_url,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Grouping *string `json:"grouping,omitempty"`
	JobResult string `json:"job_result"`
	LogLevel *LogLevelEnum `json:"log_level,omitempty"`
	LogObject NullableString `json:"log_object,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewJobLogEntry instantiates a new JobLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobLogEntry(id string, url string, jobResult string) *JobLogEntry {
	this := JobLogEntry{}
	this.Id = id
	this.Url = url
	this.JobResult = jobResult
	return &this
}

// NewJobLogEntryWithDefaults instantiates a new JobLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobLogEntryWithDefaults() *JobLogEntry {
	this := JobLogEntry{}
	return &this
}

// GetId returns the Id field value
func (o *JobLogEntry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobLogEntry) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *JobLogEntry) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JobLogEntry) SetUrl(v string) {
	o.Url = v
}

// GetAbsoluteUrl returns the AbsoluteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobLogEntry) GetAbsoluteUrl() string {
	if o == nil || o.AbsoluteUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AbsoluteUrl.Get()
}

// GetAbsoluteUrlOk returns a tuple with the AbsoluteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobLogEntry) GetAbsoluteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbsoluteUrl.Get(), o.AbsoluteUrl.IsSet()
}

// HasAbsoluteUrl returns a boolean if a field has been set.
func (o *JobLogEntry) HasAbsoluteUrl() bool {
	if o != nil && o.AbsoluteUrl.IsSet() {
		return true
	}

	return false
}

// SetAbsoluteUrl gets a reference to the given NullableString and assigns it to the AbsoluteUrl field.
func (o *JobLogEntry) SetAbsoluteUrl(v string) {
	o.AbsoluteUrl.Set(&v)
}
// SetAbsoluteUrlNil sets the value for AbsoluteUrl to be an explicit nil
func (o *JobLogEntry) SetAbsoluteUrlNil() {
	o.AbsoluteUrl.Set(nil)
}

// UnsetAbsoluteUrl ensures that no value is present for AbsoluteUrl, not even an explicit nil
func (o *JobLogEntry) UnsetAbsoluteUrl() {
	o.AbsoluteUrl.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JobLogEntry) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JobLogEntry) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *JobLogEntry) SetCreated(v time.Time) {
	o.Created = &v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *JobLogEntry) GetGrouping() string {
	if o == nil || o.Grouping == nil {
		var ret string
		return ret
	}
	return *o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetGroupingOk() (*string, bool) {
	if o == nil || o.Grouping == nil {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *JobLogEntry) HasGrouping() bool {
	if o != nil && o.Grouping != nil {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given string and assigns it to the Grouping field.
func (o *JobLogEntry) SetGrouping(v string) {
	o.Grouping = &v
}

// GetJobResult returns the JobResult field value
func (o *JobLogEntry) GetJobResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobResult
}

// GetJobResultOk returns a tuple with the JobResult field value
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetJobResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobResult, true
}

// SetJobResult sets field value
func (o *JobLogEntry) SetJobResult(v string) {
	o.JobResult = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *JobLogEntry) GetLogLevel() LogLevelEnum {
	if o == nil || o.LogLevel == nil {
		var ret LogLevelEnum
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetLogLevelOk() (*LogLevelEnum, bool) {
	if o == nil || o.LogLevel == nil {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *JobLogEntry) HasLogLevel() bool {
	if o != nil && o.LogLevel != nil {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given LogLevelEnum and assigns it to the LogLevel field.
func (o *JobLogEntry) SetLogLevel(v LogLevelEnum) {
	o.LogLevel = &v
}

// GetLogObject returns the LogObject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobLogEntry) GetLogObject() string {
	if o == nil || o.LogObject.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogObject.Get()
}

// GetLogObjectOk returns a tuple with the LogObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobLogEntry) GetLogObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogObject.Get(), o.LogObject.IsSet()
}

// HasLogObject returns a boolean if a field has been set.
func (o *JobLogEntry) HasLogObject() bool {
	if o != nil && o.LogObject.IsSet() {
		return true
	}

	return false
}

// SetLogObject gets a reference to the given NullableString and assigns it to the LogObject field.
func (o *JobLogEntry) SetLogObject(v string) {
	o.LogObject.Set(&v)
}
// SetLogObjectNil sets the value for LogObject to be an explicit nil
func (o *JobLogEntry) SetLogObjectNil() {
	o.LogObject.Set(nil)
}

// UnsetLogObject ensures that no value is present for LogObject, not even an explicit nil
func (o *JobLogEntry) UnsetLogObject() {
	o.LogObject.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *JobLogEntry) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobLogEntry) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *JobLogEntry) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *JobLogEntry) SetMessage(v string) {
	o.Message = &v
}

func (o JobLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.AbsoluteUrl.IsSet() {
		toSerialize["absolute_url"] = o.AbsoluteUrl.Get()
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Grouping != nil {
		toSerialize["grouping"] = o.Grouping
	}
	if true {
		toSerialize["job_result"] = o.JobResult
	}
	if o.LogLevel != nil {
		toSerialize["log_level"] = o.LogLevel
	}
	if o.LogObject.IsSet() {
		toSerialize["log_object"] = o.LogObject.Get()
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableJobLogEntry struct {
	value *JobLogEntry
	isSet bool
}

func (v NullableJobLogEntry) Get() *JobLogEntry {
	return v.value
}

func (v *NullableJobLogEntry) Set(val *JobLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableJobLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableJobLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobLogEntry(val *JobLogEntry) *NullableJobLogEntry {
	return &NullableJobLogEntry{value: val, isSet: true}
}

func (v NullableJobLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


