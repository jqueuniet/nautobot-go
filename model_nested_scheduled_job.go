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

// NestedScheduledJob struct for NestedScheduledJob
type NestedScheduledJob struct {
	Name *string `json:"name,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	Interval IntervalEnum `json:"interval"`
	// Cronjob syntax string for custom scheduling
	Crontab *string `json:"crontab,omitempty"`
}

// NewNestedScheduledJob instantiates a new NestedScheduledJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedScheduledJob(interval IntervalEnum) *NestedScheduledJob {
	this := NestedScheduledJob{}
	this.Interval = interval
	return &this
}

// NewNestedScheduledJobWithDefaults instantiates a new NestedScheduledJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedScheduledJobWithDefaults() *NestedScheduledJob {
	this := NestedScheduledJob{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NestedScheduledJob) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedScheduledJob) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NestedScheduledJob) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NestedScheduledJob) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *NestedScheduledJob) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedScheduledJob) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *NestedScheduledJob) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *NestedScheduledJob) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetInterval returns the Interval field value
func (o *NestedScheduledJob) GetInterval() IntervalEnum {
	if o == nil {
		var ret IntervalEnum
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *NestedScheduledJob) GetIntervalOk() (*IntervalEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *NestedScheduledJob) SetInterval(v IntervalEnum) {
	o.Interval = v
}

// GetCrontab returns the Crontab field value if set, zero value otherwise.
func (o *NestedScheduledJob) GetCrontab() string {
	if o == nil || o.Crontab == nil {
		var ret string
		return ret
	}
	return *o.Crontab
}

// GetCrontabOk returns a tuple with the Crontab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedScheduledJob) GetCrontabOk() (*string, bool) {
	if o == nil || o.Crontab == nil {
		return nil, false
	}
	return o.Crontab, true
}

// HasCrontab returns a boolean if a field has been set.
func (o *NestedScheduledJob) HasCrontab() bool {
	if o != nil && o.Crontab != nil {
		return true
	}

	return false
}

// SetCrontab gets a reference to the given string and assigns it to the Crontab field.
func (o *NestedScheduledJob) SetCrontab(v string) {
	o.Crontab = &v
}

func (o NestedScheduledJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if o.Crontab != nil {
		toSerialize["crontab"] = o.Crontab
	}
	return json.Marshal(toSerialize)
}

type NullableNestedScheduledJob struct {
	value *NestedScheduledJob
	isSet bool
}

func (v NullableNestedScheduledJob) Get() *NestedScheduledJob {
	return v.value
}

func (v *NullableNestedScheduledJob) Set(val *NestedScheduledJob) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedScheduledJob) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedScheduledJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedScheduledJob(val *NestedScheduledJob) *NullableNestedScheduledJob {
	return &NullableNestedScheduledJob{value: val, isSet: true}
}

func (v NullableNestedScheduledJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedScheduledJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


