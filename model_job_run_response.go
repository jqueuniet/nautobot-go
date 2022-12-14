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

// JobRunResponse Serializer representing responses from the JobModelViewSet.run() POST endpoint.
type JobRunResponse struct {
	Schedule JobResultSchedule `json:"schedule"`
	JobResult JobRunResponseJobResult `json:"job_result"`
}

// NewJobRunResponse instantiates a new JobRunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRunResponse(schedule JobResultSchedule, jobResult JobRunResponseJobResult) *JobRunResponse {
	this := JobRunResponse{}
	this.Schedule = schedule
	this.JobResult = jobResult
	return &this
}

// NewJobRunResponseWithDefaults instantiates a new JobRunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobRunResponseWithDefaults() *JobRunResponse {
	this := JobRunResponse{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *JobRunResponse) GetSchedule() JobResultSchedule {
	if o == nil {
		var ret JobResultSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *JobRunResponse) GetScheduleOk() (*JobResultSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *JobRunResponse) SetSchedule(v JobResultSchedule) {
	o.Schedule = v
}

// GetJobResult returns the JobResult field value
func (o *JobRunResponse) GetJobResult() JobRunResponseJobResult {
	if o == nil {
		var ret JobRunResponseJobResult
		return ret
	}

	return o.JobResult
}

// GetJobResultOk returns a tuple with the JobResult field value
// and a boolean to check if the value has been set.
func (o *JobRunResponse) GetJobResultOk() (*JobRunResponseJobResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobResult, true
}

// SetJobResult sets field value
func (o *JobRunResponse) SetJobResult(v JobRunResponseJobResult) {
	o.JobResult = v
}

func (o JobRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["job_result"] = o.JobResult
	}
	return json.Marshal(toSerialize)
}

type NullableJobRunResponse struct {
	value *JobRunResponse
	isSet bool
}

func (v NullableJobRunResponse) Get() *JobRunResponse {
	return v.value
}

func (v *NullableJobRunResponse) Set(val *JobRunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRunResponse(val *JobRunResponse) *NullableJobRunResponse {
	return &NullableJobRunResponse{value: val, isSet: true}
}

func (v NullableJobRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


