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

// FaceEnum the model 'FaceEnum'
type FaceEnum string

// List of FaceEnum
const (
	FRONT FaceEnum = "front"
	REAR FaceEnum = "rear"
)

// All allowed values of FaceEnum enum
var AllowedFaceEnumEnumValues = []FaceEnum{
	"front",
	"rear",
}

func (v *FaceEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FaceEnum(value)
	for _, existing := range AllowedFaceEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FaceEnum", value)
}

// NewFaceEnumFromValue returns a pointer to a valid FaceEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFaceEnumFromValue(v string) (*FaceEnum, error) {
	ev := FaceEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FaceEnum: valid values are %v", v, AllowedFaceEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceEnum) IsValid() bool {
	for _, existing := range AllowedFaceEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FaceEnum value
func (v FaceEnum) Ptr() *FaceEnum {
	return &v
}

type NullableFaceEnum struct {
	value *FaceEnum
	isSet bool
}

func (v NullableFaceEnum) Get() *FaceEnum {
	return v.value
}

func (v *NullableFaceEnum) Set(val *FaceEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFaceEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFaceEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaceEnum(val *FaceEnum) *NullableFaceEnum {
	return &NullableFaceEnum{value: val, isSet: true}
}

func (v NullableFaceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaceEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

