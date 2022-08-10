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

// JobInput struct for JobInput
type JobInput struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Commit *bool `json:"commit,omitempty"`
	Schedule *NestedScheduledJob `json:"schedule,omitempty"`
}

// NewJobInput instantiates a new JobInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobInput() *JobInput {
	this := JobInput{}
	return &this
}

// NewJobInputWithDefaults instantiates a new JobInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobInputWithDefaults() *JobInput {
	this := JobInput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *JobInput) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInput) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *JobInput) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *JobInput) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *JobInput) GetCommit() bool {
	if o == nil || o.Commit == nil {
		var ret bool
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInput) GetCommitOk() (*bool, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *JobInput) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given bool and assigns it to the Commit field.
func (o *JobInput) SetCommit(v bool) {
	o.Commit = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *JobInput) GetSchedule() NestedScheduledJob {
	if o == nil || o.Schedule == nil {
		var ret NestedScheduledJob
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobInput) GetScheduleOk() (*NestedScheduledJob, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *JobInput) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NestedScheduledJob and assigns it to the Schedule field.
func (o *JobInput) SetSchedule(v NestedScheduledJob) {
	o.Schedule = &v
}

func (o JobInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableJobInput struct {
	value *JobInput
	isSet bool
}

func (v NullableJobInput) Get() *JobInput {
	return v.value
}

func (v *NullableJobInput) Set(val *JobInput) {
	v.value = val
	v.isSet = true
}

func (v NullableJobInput) IsSet() bool {
	return v.isSet
}

func (v *NullableJobInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobInput(val *JobInput) *NullableJobInput {
	return &NullableJobInput{value: val, isSet: true}
}

func (v NullableJobInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


