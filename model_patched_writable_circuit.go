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

// PatchedWritableCircuit Mixin to add `status` choice field to model serializers.
type PatchedWritableCircuit struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	Cid *string `json:"cid,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Type *string `json:"type,omitempty"`
	Status *WritableCircuitStatusEnum `json:"status,omitempty"`
	Tenant NullableString `json:"tenant,omitempty"`
	InstallDate NullableString `json:"install_date,omitempty"`
	CommitRate NullableInt32 `json:"commit_rate,omitempty"`
	Description *string `json:"description,omitempty"`
	TerminationA *CircuitTerminationA `json:"termination_a,omitempty"`
	TerminationZ *CircuitTerminationA `json:"termination_z,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created *string `json:"created,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableCircuit instantiates a new PatchedWritableCircuit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCircuit() *PatchedWritableCircuit {
	this := PatchedWritableCircuit{}
	return &this
}

// NewPatchedWritableCircuitWithDefaults instantiates a new PatchedWritableCircuit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCircuitWithDefaults() *PatchedWritableCircuit {
	this := PatchedWritableCircuit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableCircuit) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableCircuit) SetUrl(v string) {
	o.Url = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetCid() string {
	if o == nil || o.Cid == nil {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetCidOk() (*string, bool) {
	if o == nil || o.Cid == nil {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasCid() bool {
	if o != nil && o.Cid != nil {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *PatchedWritableCircuit) SetCid(v string) {
	o.Cid = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PatchedWritableCircuit) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedWritableCircuit) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetStatus() WritableCircuitStatusEnum {
	if o == nil || o.Status == nil {
		var ret WritableCircuitStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetStatusOk() (*WritableCircuitStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WritableCircuitStatusEnum and assigns it to the Status field.
func (o *PatchedWritableCircuit) SetStatus(v WritableCircuitStatusEnum) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuit) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuit) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableString and assigns it to the Tenant field.
func (o *PatchedWritableCircuit) SetTenant(v string) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableCircuit) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableCircuit) UnsetTenant() {
	o.Tenant.Unset()
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuit) GetInstallDate() string {
	if o == nil || o.InstallDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstallDate.Get()
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuit) GetInstallDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallDate.Get(), o.InstallDate.IsSet()
}

// HasInstallDate returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasInstallDate() bool {
	if o != nil && o.InstallDate.IsSet() {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given NullableString and assigns it to the InstallDate field.
func (o *PatchedWritableCircuit) SetInstallDate(v string) {
	o.InstallDate.Set(&v)
}
// SetInstallDateNil sets the value for InstallDate to be an explicit nil
func (o *PatchedWritableCircuit) SetInstallDateNil() {
	o.InstallDate.Set(nil)
}

// UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
func (o *PatchedWritableCircuit) UnsetInstallDate() {
	o.InstallDate.Unset()
}

// GetCommitRate returns the CommitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCircuit) GetCommitRate() int32 {
	if o == nil || o.CommitRate.Get() == nil {
		var ret int32
		return ret
	}
	return *o.CommitRate.Get()
}

// GetCommitRateOk returns a tuple with the CommitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCircuit) GetCommitRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitRate.Get(), o.CommitRate.IsSet()
}

// HasCommitRate returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasCommitRate() bool {
	if o != nil && o.CommitRate.IsSet() {
		return true
	}

	return false
}

// SetCommitRate gets a reference to the given NullableInt32 and assigns it to the CommitRate field.
func (o *PatchedWritableCircuit) SetCommitRate(v int32) {
	o.CommitRate.Set(&v)
}
// SetCommitRateNil sets the value for CommitRate to be an explicit nil
func (o *PatchedWritableCircuit) SetCommitRateNil() {
	o.CommitRate.Set(nil)
}

// UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
func (o *PatchedWritableCircuit) UnsetCommitRate() {
	o.CommitRate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableCircuit) SetDescription(v string) {
	o.Description = &v
}

// GetTerminationA returns the TerminationA field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetTerminationA() CircuitTerminationA {
	if o == nil || o.TerminationA == nil {
		var ret CircuitTerminationA
		return ret
	}
	return *o.TerminationA
}

// GetTerminationAOk returns a tuple with the TerminationA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetTerminationAOk() (*CircuitTerminationA, bool) {
	if o == nil || o.TerminationA == nil {
		return nil, false
	}
	return o.TerminationA, true
}

// HasTerminationA returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasTerminationA() bool {
	if o != nil && o.TerminationA != nil {
		return true
	}

	return false
}

// SetTerminationA gets a reference to the given CircuitTerminationA and assigns it to the TerminationA field.
func (o *PatchedWritableCircuit) SetTerminationA(v CircuitTerminationA) {
	o.TerminationA = &v
}

// GetTerminationZ returns the TerminationZ field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetTerminationZ() CircuitTerminationA {
	if o == nil || o.TerminationZ == nil {
		var ret CircuitTerminationA
		return ret
	}
	return *o.TerminationZ
}

// GetTerminationZOk returns a tuple with the TerminationZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetTerminationZOk() (*CircuitTerminationA, bool) {
	if o == nil || o.TerminationZ == nil {
		return nil, false
	}
	return o.TerminationZ, true
}

// HasTerminationZ returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasTerminationZ() bool {
	if o != nil && o.TerminationZ != nil {
		return true
	}

	return false
}

// SetTerminationZ gets a reference to the given CircuitTerminationA and assigns it to the TerminationZ field.
func (o *PatchedWritableCircuit) SetTerminationZ(v CircuitTerminationA) {
	o.TerminationZ = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableCircuit) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableCircuit) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableCircuit) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PatchedWritableCircuit) SetCreated(v string) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PatchedWritableCircuit) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableCircuit) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableCircuit) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCircuit) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableCircuit) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableCircuit) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableCircuit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Cid != nil {
		toSerialize["cid"] = o.Cid
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.InstallDate.IsSet() {
		toSerialize["install_date"] = o.InstallDate.Get()
	}
	if o.CommitRate.IsSet() {
		toSerialize["commit_rate"] = o.CommitRate.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TerminationA != nil {
		toSerialize["termination_a"] = o.TerminationA
	}
	if o.TerminationZ != nil {
		toSerialize["termination_z"] = o.TerminationZ
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
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableCircuit struct {
	value *PatchedWritableCircuit
	isSet bool
}

func (v NullablePatchedWritableCircuit) Get() *PatchedWritableCircuit {
	return v.value
}

func (v *NullablePatchedWritableCircuit) Set(val *PatchedWritableCircuit) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCircuit) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCircuit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCircuit(val *PatchedWritableCircuit) *NullablePatchedWritableCircuit {
	return &NullablePatchedWritableCircuit{value: val, isSet: true}
}

func (v NullablePatchedWritableCircuit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCircuit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


