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

// PatchedWritableRackType - struct for PatchedWritableRackType
type PatchedWritableRackType struct {
	BlankEnum *BlankEnum
	RackTypeChoices *RackTypeChoices
}

// BlankEnumAsPatchedWritableRackType is a convenience function that returns BlankEnum wrapped in PatchedWritableRackType
func BlankEnumAsPatchedWritableRackType(v *BlankEnum) PatchedWritableRackType {
	return PatchedWritableRackType{
		BlankEnum: v,
	}
}

// RackTypeChoicesAsPatchedWritableRackType is a convenience function that returns RackTypeChoices wrapped in PatchedWritableRackType
func RackTypeChoicesAsPatchedWritableRackType(v *RackTypeChoices) PatchedWritableRackType {
	return PatchedWritableRackType{
		RackTypeChoices: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableRackType) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into RackTypeChoices
	err = newStrictDecoder(data).Decode(&dst.RackTypeChoices)
	if err == nil {
		jsonRackTypeChoices, _ := json.Marshal(dst.RackTypeChoices)
		if string(jsonRackTypeChoices) == "{}" { // empty struct
			dst.RackTypeChoices = nil
		} else {
			match++
		}
	} else {
		dst.RackTypeChoices = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.RackTypeChoices = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(PatchedWritableRackType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(PatchedWritableRackType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableRackType) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.RackTypeChoices != nil {
		return json.Marshal(&src.RackTypeChoices)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableRackType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.RackTypeChoices != nil {
		return obj.RackTypeChoices
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableRackType struct {
	value *PatchedWritableRackType
	isSet bool
}

func (v NullablePatchedWritableRackType) Get() *PatchedWritableRackType {
	return v.value
}

func (v *NullablePatchedWritableRackType) Set(val *PatchedWritableRackType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackType(val *PatchedWritableRackType) *NullablePatchedWritableRackType {
	return &NullablePatchedWritableRackType{value: val, isSet: true}
}

func (v NullablePatchedWritableRackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


