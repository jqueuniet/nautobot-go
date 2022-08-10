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

// JobClassDetail struct for JobClassDetail
type JobClassDetail struct {
	Url string `json:"url"`
	Id string `json:"id"`
	Pk NullableString `json:"pk"`
	Name string `json:"name"`
	Description string `json:"description"`
	TestMethods []string `json:"test_methods"`
	Vars map[string]interface{} `json:"vars"`
	Result *JobResult `json:"result,omitempty"`
}

// NewJobClassDetail instantiates a new JobClassDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobClassDetail(url string, id string, pk NullableString, name string, description string, testMethods []string, vars map[string]interface{}) *JobClassDetail {
	this := JobClassDetail{}
	this.Url = url
	this.Id = id
	this.Pk = pk
	this.Name = name
	this.Description = description
	this.TestMethods = testMethods
	this.Vars = vars
	return &this
}

// NewJobClassDetailWithDefaults instantiates a new JobClassDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobClassDetailWithDefaults() *JobClassDetail {
	this := JobClassDetail{}
	return &this
}

// GetUrl returns the Url field value
func (o *JobClassDetail) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JobClassDetail) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *JobClassDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobClassDetail) SetId(v string) {
	o.Id = v
}

// GetPk returns the Pk field value
// If the value is explicit nil, the zero value for string will be returned
func (o *JobClassDetail) GetPk() string {
	if o == nil || o.Pk.Get() == nil {
		var ret string
		return ret
	}

	return *o.Pk.Get()
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobClassDetail) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pk.Get(), o.Pk.IsSet()
}

// SetPk sets field value
func (o *JobClassDetail) SetPk(v string) {
	o.Pk.Set(&v)
}

// GetName returns the Name field value
func (o *JobClassDetail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobClassDetail) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *JobClassDetail) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *JobClassDetail) SetDescription(v string) {
	o.Description = v
}

// GetTestMethods returns the TestMethods field value
func (o *JobClassDetail) GetTestMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TestMethods
}

// GetTestMethodsOk returns a tuple with the TestMethods field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetTestMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestMethods, true
}

// SetTestMethods sets field value
func (o *JobClassDetail) SetTestMethods(v []string) {
	o.TestMethods = v
}

// GetVars returns the Vars field value
func (o *JobClassDetail) GetVars() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetVarsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vars, true
}

// SetVars sets field value
func (o *JobClassDetail) SetVars(v map[string]interface{}) {
	o.Vars = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *JobClassDetail) GetResult() JobResult {
	if o == nil || o.Result == nil {
		var ret JobResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobClassDetail) GetResultOk() (*JobResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *JobClassDetail) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given JobResult and assigns it to the Result field.
func (o *JobClassDetail) SetResult(v JobResult) {
	o.Result = &v
}

func (o JobClassDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["pk"] = o.Pk.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["test_methods"] = o.TestMethods
	}
	if true {
		toSerialize["vars"] = o.Vars
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableJobClassDetail struct {
	value *JobClassDetail
	isSet bool
}

func (v NullableJobClassDetail) Get() *JobClassDetail {
	return v.value
}

func (v *NullableJobClassDetail) Set(val *JobClassDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableJobClassDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableJobClassDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobClassDetail(val *JobClassDetail) *NullableJobClassDetail {
	return &NullableJobClassDetail{value: val, isSet: true}
}

func (v NullableJobClassDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobClassDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


