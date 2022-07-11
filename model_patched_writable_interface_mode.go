/*
API Documentation

Source of truth and network automation platform

API version: 1.3.7 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PatchedWritableInterfaceMode - struct for PatchedWritableInterfaceMode
type PatchedWritableInterfaceMode struct {
	BlankEnum *BlankEnum
	ModeEnum *ModeEnum
}

// BlankEnumAsPatchedWritableInterfaceMode is a convenience function that returns BlankEnum wrapped in PatchedWritableInterfaceMode
func BlankEnumAsPatchedWritableInterfaceMode(v *BlankEnum) PatchedWritableInterfaceMode {
	return PatchedWritableInterfaceMode{
		BlankEnum: v,
	}
}

// ModeEnumAsPatchedWritableInterfaceMode is a convenience function that returns ModeEnum wrapped in PatchedWritableInterfaceMode
func ModeEnumAsPatchedWritableInterfaceMode(v *ModeEnum) PatchedWritableInterfaceMode {
	return PatchedWritableInterfaceMode{
		ModeEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableInterfaceMode) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ModeEnum
	err = newStrictDecoder(data).Decode(&dst.ModeEnum)
	if err == nil {
		jsonModeEnum, _ := json.Marshal(dst.ModeEnum)
		if string(jsonModeEnum) == "{}" { // empty struct
			dst.ModeEnum = nil
		} else {
			match++
		}
	} else {
		dst.ModeEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.ModeEnum = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PatchedWritableInterfaceMode)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PatchedWritableInterfaceMode)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableInterfaceMode) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.ModeEnum != nil {
		return json.Marshal(&src.ModeEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableInterfaceMode) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.ModeEnum != nil {
		return obj.ModeEnum
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableInterfaceMode struct {
	value *PatchedWritableInterfaceMode
	isSet bool
}

func (v NullablePatchedWritableInterfaceMode) Get() *PatchedWritableInterfaceMode {
	return v.value
}

func (v *NullablePatchedWritableInterfaceMode) Set(val *PatchedWritableInterfaceMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableInterfaceMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableInterfaceMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableInterfaceMode(val *PatchedWritableInterfaceMode) *NullablePatchedWritableInterfaceMode {
	return &NullablePatchedWritableInterfaceMode{value: val, isSet: true}
}

func (v NullablePatchedWritableInterfaceMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableInterfaceMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


