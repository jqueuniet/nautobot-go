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

// PatchedWritableConsolePortTemplateType - struct for PatchedWritableConsolePortTemplateType
type PatchedWritableConsolePortTemplateType struct {
	BlankEnum *BlankEnum
	ConsolePortTypeChoices *ConsolePortTypeChoices
}

// BlankEnumAsPatchedWritableConsolePortTemplateType is a convenience function that returns BlankEnum wrapped in PatchedWritableConsolePortTemplateType
func BlankEnumAsPatchedWritableConsolePortTemplateType(v *BlankEnum) PatchedWritableConsolePortTemplateType {
	return PatchedWritableConsolePortTemplateType{
		BlankEnum: v,
	}
}

// ConsolePortTypeChoicesAsPatchedWritableConsolePortTemplateType is a convenience function that returns ConsolePortTypeChoices wrapped in PatchedWritableConsolePortTemplateType
func ConsolePortTypeChoicesAsPatchedWritableConsolePortTemplateType(v *ConsolePortTypeChoices) PatchedWritableConsolePortTemplateType {
	return PatchedWritableConsolePortTemplateType{
		ConsolePortTypeChoices: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableConsolePortTemplateType) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ConsolePortTypeChoices
	err = newStrictDecoder(data).Decode(&dst.ConsolePortTypeChoices)
	if err == nil {
		jsonConsolePortTypeChoices, _ := json.Marshal(dst.ConsolePortTypeChoices)
		if string(jsonConsolePortTypeChoices) == "{}" { // empty struct
			dst.ConsolePortTypeChoices = nil
		} else {
			match++
		}
	} else {
		dst.ConsolePortTypeChoices = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.ConsolePortTypeChoices = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PatchedWritableConsolePortTemplateType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PatchedWritableConsolePortTemplateType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableConsolePortTemplateType) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.ConsolePortTypeChoices != nil {
		return json.Marshal(&src.ConsolePortTypeChoices)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableConsolePortTemplateType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.ConsolePortTypeChoices != nil {
		return obj.ConsolePortTypeChoices
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableConsolePortTemplateType struct {
	value *PatchedWritableConsolePortTemplateType
	isSet bool
}

func (v NullablePatchedWritableConsolePortTemplateType) Get() *PatchedWritableConsolePortTemplateType {
	return v.value
}

func (v *NullablePatchedWritableConsolePortTemplateType) Set(val *PatchedWritableConsolePortTemplateType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsolePortTemplateType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsolePortTemplateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsolePortTemplateType(val *PatchedWritableConsolePortTemplateType) *NullablePatchedWritableConsolePortTemplateType {
	return &NullablePatchedWritableConsolePortTemplateType{value: val, isSet: true}
}

func (v NullablePatchedWritableConsolePortTemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsolePortTemplateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


