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

// SubdeviceRoleEnum the model 'SubdeviceRoleEnum'
type SubdeviceRoleEnum string

// List of SubdeviceRoleEnum
const (
	PARENT SubdeviceRoleEnum = "parent"
	CHILD SubdeviceRoleEnum = "child"
)

// All allowed values of SubdeviceRoleEnum enum
var AllowedSubdeviceRoleEnumEnumValues = []SubdeviceRoleEnum{
	"parent",
	"child",
}

func (v *SubdeviceRoleEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubdeviceRoleEnum(value)
	for _, existing := range AllowedSubdeviceRoleEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubdeviceRoleEnum", value)
}

// NewSubdeviceRoleEnumFromValue returns a pointer to a valid SubdeviceRoleEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubdeviceRoleEnumFromValue(v string) (*SubdeviceRoleEnum, error) {
	ev := SubdeviceRoleEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubdeviceRoleEnum: valid values are %v", v, AllowedSubdeviceRoleEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubdeviceRoleEnum) IsValid() bool {
	for _, existing := range AllowedSubdeviceRoleEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubdeviceRoleEnum value
func (v SubdeviceRoleEnum) Ptr() *SubdeviceRoleEnum {
	return &v
}

type NullableSubdeviceRoleEnum struct {
	value *SubdeviceRoleEnum
	isSet bool
}

func (v NullableSubdeviceRoleEnum) Get() *SubdeviceRoleEnum {
	return v.value
}

func (v *NullableSubdeviceRoleEnum) Set(val *SubdeviceRoleEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSubdeviceRoleEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSubdeviceRoleEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubdeviceRoleEnum(val *SubdeviceRoleEnum) *NullableSubdeviceRoleEnum {
	return &NullableSubdeviceRoleEnum{value: val, isSet: true}
}

func (v NullableSubdeviceRoleEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubdeviceRoleEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

