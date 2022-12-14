/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PatchedWritablePowerPortTemplateType - struct for PatchedWritablePowerPortTemplateType
type PatchedWritablePowerPortTemplateType struct {
	BlankEnum *BlankEnum
	PowerPortTypeChoices *PowerPortTypeChoices
}

// BlankEnumAsPatchedWritablePowerPortTemplateType is a convenience function that returns BlankEnum wrapped in PatchedWritablePowerPortTemplateType
func BlankEnumAsPatchedWritablePowerPortTemplateType(v *BlankEnum) PatchedWritablePowerPortTemplateType {
	return PatchedWritablePowerPortTemplateType{
		BlankEnum: v,
	}
}

// PowerPortTypeChoicesAsPatchedWritablePowerPortTemplateType is a convenience function that returns PowerPortTypeChoices wrapped in PatchedWritablePowerPortTemplateType
func PowerPortTypeChoicesAsPatchedWritablePowerPortTemplateType(v *PowerPortTypeChoices) PatchedWritablePowerPortTemplateType {
	return PatchedWritablePowerPortTemplateType{
		PowerPortTypeChoices: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritablePowerPortTemplateType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			match++
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into PowerPortTypeChoices
	err = newStrictDecoder(data).Decode(&dst.PowerPortTypeChoices)
	if err == nil {
		jsonPowerPortTypeChoices, _ := json.Marshal(dst.PowerPortTypeChoices)
		if string(jsonPowerPortTypeChoices) == "{}" { // empty struct
			dst.PowerPortTypeChoices = nil
		} else {
			match++
		}
	} else {
		dst.PowerPortTypeChoices = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.PowerPortTypeChoices = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PatchedWritablePowerPortTemplateType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PatchedWritablePowerPortTemplateType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritablePowerPortTemplateType) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.PowerPortTypeChoices != nil {
		return json.Marshal(&src.PowerPortTypeChoices)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritablePowerPortTemplateType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.PowerPortTypeChoices != nil {
		return obj.PowerPortTypeChoices
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritablePowerPortTemplateType struct {
	value *PatchedWritablePowerPortTemplateType
	isSet bool
}

func (v NullablePatchedWritablePowerPortTemplateType) Get() *PatchedWritablePowerPortTemplateType {
	return v.value
}

func (v *NullablePatchedWritablePowerPortTemplateType) Set(val *PatchedWritablePowerPortTemplateType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPortTemplateType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPortTemplateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPortTemplateType(val *PatchedWritablePowerPortTemplateType) *NullablePatchedWritablePowerPortTemplateType {
	return &NullablePatchedWritablePowerPortTemplateType{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPortTemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPortTemplateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


