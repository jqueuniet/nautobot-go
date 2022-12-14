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

// LengthUnitEnum the model 'LengthUnitEnum'
type LengthUnitEnum string

// List of LengthUnitEnum
const (
	M  LengthUnitEnum = "m"
	CM LengthUnitEnum = "cm"
	FT LengthUnitEnum = "ft"
	//IN LengthUnitEnum = "in"
)

// All allowed values of LengthUnitEnum enum
var AllowedLengthUnitEnumEnumValues = []LengthUnitEnum{
	"m",
	"cm",
	"ft",
	"in",
}

func (v *LengthUnitEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LengthUnitEnum(value)
	for _, existing := range AllowedLengthUnitEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LengthUnitEnum", value)
}

// NewLengthUnitEnumFromValue returns a pointer to a valid LengthUnitEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLengthUnitEnumFromValue(v string) (*LengthUnitEnum, error) {
	ev := LengthUnitEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LengthUnitEnum: valid values are %v", v, AllowedLengthUnitEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LengthUnitEnum) IsValid() bool {
	for _, existing := range AllowedLengthUnitEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LengthUnitEnum value
func (v LengthUnitEnum) Ptr() *LengthUnitEnum {
	return &v
}

type NullableLengthUnitEnum struct {
	value *LengthUnitEnum
	isSet bool
}

func (v NullableLengthUnitEnum) Get() *LengthUnitEnum {
	return v.value
}

func (v *NullableLengthUnitEnum) Set(val *LengthUnitEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLengthUnitEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLengthUnitEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLengthUnitEnum(val *LengthUnitEnum) *NullableLengthUnitEnum {
	return &NullableLengthUnitEnum{value: val, isSet: true}
}

func (v NullableLengthUnitEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLengthUnitEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
