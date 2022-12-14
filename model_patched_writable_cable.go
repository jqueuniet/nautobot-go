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

// PatchedWritableCable Mixin to add `status` choice field to model serializers.
type PatchedWritableCable struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	TerminationAType *string `json:"termination_a_type,omitempty"`
	TerminationAId *string `json:"termination_a_id,omitempty"`
	TerminationA map[string]interface{} `json:"termination_a,omitempty"`
	TerminationBType *string `json:"termination_b_type,omitempty"`
	TerminationBId *string `json:"termination_b_id,omitempty"`
	TerminationB map[string]interface{} `json:"termination_b,omitempty"`
	Type *CableType `json:"type,omitempty"`
	Status *WritableCableStatusEnum `json:"status,omitempty"`
	Label *string `json:"label,omitempty"`
	Color *string `json:"color,omitempty"`
	Length NullableInt32 `json:"length,omitempty"`
	LengthUnit *PatchedWritableCableLengthUnit `json:"length_unit,omitempty"`
	Tags []TagSerializerField `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	ComputedFields map[string]interface{} `json:"computed_fields,omitempty"`
	// Human friendly display value
	Display *string `json:"display,omitempty"`
}

// NewPatchedWritableCable instantiates a new PatchedWritableCable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableCable() *PatchedWritableCable {
	this := PatchedWritableCable{}
	return &this
}

// NewPatchedWritableCableWithDefaults instantiates a new PatchedWritableCable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableCableWithDefaults() *PatchedWritableCable {
	this := PatchedWritableCable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchedWritableCable) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedWritableCable) SetUrl(v string) {
	o.Url = &v
}

// GetTerminationAType returns the TerminationAType field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetTerminationAType() string {
	if o == nil || o.TerminationAType == nil {
		var ret string
		return ret
	}
	return *o.TerminationAType
}

// GetTerminationATypeOk returns a tuple with the TerminationAType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTerminationATypeOk() (*string, bool) {
	if o == nil || o.TerminationAType == nil {
		return nil, false
	}
	return o.TerminationAType, true
}

// HasTerminationAType returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationAType() bool {
	if o != nil && o.TerminationAType != nil {
		return true
	}

	return false
}

// SetTerminationAType gets a reference to the given string and assigns it to the TerminationAType field.
func (o *PatchedWritableCable) SetTerminationAType(v string) {
	o.TerminationAType = &v
}

// GetTerminationAId returns the TerminationAId field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetTerminationAId() string {
	if o == nil || o.TerminationAId == nil {
		var ret string
		return ret
	}
	return *o.TerminationAId
}

// GetTerminationAIdOk returns a tuple with the TerminationAId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTerminationAIdOk() (*string, bool) {
	if o == nil || o.TerminationAId == nil {
		return nil, false
	}
	return o.TerminationAId, true
}

// HasTerminationAId returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationAId() bool {
	if o != nil && o.TerminationAId != nil {
		return true
	}

	return false
}

// SetTerminationAId gets a reference to the given string and assigns it to the TerminationAId field.
func (o *PatchedWritableCable) SetTerminationAId(v string) {
	o.TerminationAId = &v
}

// GetTerminationA returns the TerminationA field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCable) GetTerminationA() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TerminationA
}

// GetTerminationAOk returns a tuple with the TerminationA field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCable) GetTerminationAOk() (map[string]interface{}, bool) {
	if o == nil || o.TerminationA == nil {
		return nil, false
	}
	return o.TerminationA, true
}

// HasTerminationA returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationA() bool {
	if o != nil && o.TerminationA != nil {
		return true
	}

	return false
}

// SetTerminationA gets a reference to the given map[string]interface{} and assigns it to the TerminationA field.
func (o *PatchedWritableCable) SetTerminationA(v map[string]interface{}) {
	o.TerminationA = v
}

// GetTerminationBType returns the TerminationBType field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetTerminationBType() string {
	if o == nil || o.TerminationBType == nil {
		var ret string
		return ret
	}
	return *o.TerminationBType
}

// GetTerminationBTypeOk returns a tuple with the TerminationBType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTerminationBTypeOk() (*string, bool) {
	if o == nil || o.TerminationBType == nil {
		return nil, false
	}
	return o.TerminationBType, true
}

// HasTerminationBType returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationBType() bool {
	if o != nil && o.TerminationBType != nil {
		return true
	}

	return false
}

// SetTerminationBType gets a reference to the given string and assigns it to the TerminationBType field.
func (o *PatchedWritableCable) SetTerminationBType(v string) {
	o.TerminationBType = &v
}

// GetTerminationBId returns the TerminationBId field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetTerminationBId() string {
	if o == nil || o.TerminationBId == nil {
		var ret string
		return ret
	}
	return *o.TerminationBId
}

// GetTerminationBIdOk returns a tuple with the TerminationBId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTerminationBIdOk() (*string, bool) {
	if o == nil || o.TerminationBId == nil {
		return nil, false
	}
	return o.TerminationBId, true
}

// HasTerminationBId returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationBId() bool {
	if o != nil && o.TerminationBId != nil {
		return true
	}

	return false
}

// SetTerminationBId gets a reference to the given string and assigns it to the TerminationBId field.
func (o *PatchedWritableCable) SetTerminationBId(v string) {
	o.TerminationBId = &v
}

// GetTerminationB returns the TerminationB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCable) GetTerminationB() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TerminationB
}

// GetTerminationBOk returns a tuple with the TerminationB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCable) GetTerminationBOk() (map[string]interface{}, bool) {
	if o == nil || o.TerminationB == nil {
		return nil, false
	}
	return o.TerminationB, true
}

// HasTerminationB returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTerminationB() bool {
	if o != nil && o.TerminationB != nil {
		return true
	}

	return false
}

// SetTerminationB gets a reference to the given map[string]interface{} and assigns it to the TerminationB field.
func (o *PatchedWritableCable) SetTerminationB(v map[string]interface{}) {
	o.TerminationB = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetType() CableType {
	if o == nil || o.Type == nil {
		var ret CableType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTypeOk() (*CableType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given CableType and assigns it to the Type field.
func (o *PatchedWritableCable) SetType(v CableType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetStatus() WritableCableStatusEnum {
	if o == nil || o.Status == nil {
		var ret WritableCableStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetStatusOk() (*WritableCableStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WritableCableStatusEnum and assigns it to the Status field.
func (o *PatchedWritableCable) SetStatus(v WritableCableStatusEnum) {
	o.Status = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableCable) SetLabel(v string) {
	o.Label = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *PatchedWritableCable) SetColor(v string) {
	o.Color = &v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableCable) GetLength() int32 {
	if o == nil || o.Length.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableCable) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableInt32 and assigns it to the Length field.
func (o *PatchedWritableCable) SetLength(v int32) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *PatchedWritableCable) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *PatchedWritableCable) UnsetLength() {
	o.Length.Unset()
}

// GetLengthUnit returns the LengthUnit field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetLengthUnit() PatchedWritableCableLengthUnit {
	if o == nil || o.LengthUnit == nil {
		var ret PatchedWritableCableLengthUnit
		return ret
	}
	return *o.LengthUnit
}

// GetLengthUnitOk returns a tuple with the LengthUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetLengthUnitOk() (*PatchedWritableCableLengthUnit, bool) {
	if o == nil || o.LengthUnit == nil {
		return nil, false
	}
	return o.LengthUnit, true
}

// HasLengthUnit returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasLengthUnit() bool {
	if o != nil && o.LengthUnit != nil {
		return true
	}

	return false
}

// SetLengthUnit gets a reference to the given PatchedWritableCableLengthUnit and assigns it to the LengthUnit field.
func (o *PatchedWritableCable) SetLengthUnit(v PatchedWritableCableLengthUnit) {
	o.LengthUnit = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetTags() []TagSerializerField {
	if o == nil || o.Tags == nil {
		var ret []TagSerializerField
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetTagsOk() ([]TagSerializerField, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSerializerField and assigns it to the Tags field.
func (o *PatchedWritableCable) SetTags(v []TagSerializerField) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetCustomFields() map[string]interface{} {
	if o == nil || o.CustomFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableCable) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetComputedFields returns the ComputedFields field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetComputedFields() map[string]interface{} {
	if o == nil || o.ComputedFields == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComputedFields
}

// GetComputedFieldsOk returns a tuple with the ComputedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetComputedFieldsOk() (map[string]interface{}, bool) {
	if o == nil || o.ComputedFields == nil {
		return nil, false
	}
	return o.ComputedFields, true
}

// HasComputedFields returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasComputedFields() bool {
	if o != nil && o.ComputedFields != nil {
		return true
	}

	return false
}

// SetComputedFields gets a reference to the given map[string]interface{} and assigns it to the ComputedFields field.
func (o *PatchedWritableCable) SetComputedFields(v map[string]interface{}) {
	o.ComputedFields = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *PatchedWritableCable) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableCable) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *PatchedWritableCable) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *PatchedWritableCable) SetDisplay(v string) {
	o.Display = &v
}

func (o PatchedWritableCable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.TerminationAType != nil {
		toSerialize["termination_a_type"] = o.TerminationAType
	}
	if o.TerminationAId != nil {
		toSerialize["termination_a_id"] = o.TerminationAId
	}
	if o.TerminationA != nil {
		toSerialize["termination_a"] = o.TerminationA
	}
	if o.TerminationBType != nil {
		toSerialize["termination_b_type"] = o.TerminationBType
	}
	if o.TerminationBId != nil {
		toSerialize["termination_b_id"] = o.TerminationBId
	}
	if o.TerminationB != nil {
		toSerialize["termination_b"] = o.TerminationB
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if o.LengthUnit != nil {
		toSerialize["length_unit"] = o.LengthUnit
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.ComputedFields != nil {
		toSerialize["computed_fields"] = o.ComputedFields
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWritableCable struct {
	value *PatchedWritableCable
	isSet bool
}

func (v NullablePatchedWritableCable) Get() *PatchedWritableCable {
	return v.value
}

func (v *NullablePatchedWritableCable) Set(val *PatchedWritableCable) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCable) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCable(val *PatchedWritableCable) *NullablePatchedWritableCable {
	return &NullablePatchedWritableCable{value: val, isSet: true}
}

func (v NullablePatchedWritableCable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


